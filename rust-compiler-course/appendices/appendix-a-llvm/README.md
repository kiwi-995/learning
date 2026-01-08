# Appendix A: Hands-On LLVM Backend

> *"LLVM is the Swiss Army knife of compiler backends."*

This appendix teaches you to replace the bytecode VM with an LLVM backend, generating native machine code!

## 🎯 Learning Objectives

- Understand LLVM's intermediate representation (IR)
- Use Inkwell for safe LLVM bindings in Rust
- Implement JIT and AOT compilation
- Generate standalone executables from Rusty programs

---

## A.1 Introduction to LLVM

**LLVM** (Low-Level Virtual Machine) is a mature compiler infrastructure used by:
- **Clang** (C/C++ compiler)
- **Rustc** (Rust compiler)  
- **Swift** compiler
- Many other language implementations

### What LLVM Provides

1. **Intermediate Representation (IR)**: A low-level, typed, SSA-based language
2. **Optimization passes**: Hundreds of optimizations automatically applied
3. **Code generation**: Targets x86, ARM, WebAssembly, and more
4. **JIT compilation**: Execute code at runtime

### LLVM IR Example

```llvm
define i64 @add(i64 %a, i64 %b) {
entry:
    %result = add i64 %a, %b
    ret i64 %result
}

define i64 @main() {
entry:
    %result = call i64 @add(i64 5, i64 3)
    ret i64 %result
}
```

---

## A.2 Setting Up Inkwell

[Inkwell](https://github.com/TheDan64/inkwell) provides safe Rust bindings to LLVM.

### Installation

1. Install LLVM (version 15-18):
```bash
# macOS
brew install llvm@17

# Ubuntu
sudo apt install llvm-17 llvm-17-dev

# Set environment variable
export LLVM_SYS_170_PREFIX=/usr/local/opt/llvm@17
```

2. Add to Cargo.toml:
```toml
[dependencies]
inkwell = { version = "0.4", features = ["llvm17-0"] }
```

### Hello LLVM

```rust
use inkwell::context::Context;
use inkwell::OptimizationLevel;

fn main() {
    let context = Context::create();
    let module = context.create_module("hello");
    let builder = context.create_builder();
    
    // Define i64 type
    let i64_type = context.i64_type();
    
    // Create function: fn add(a: i64, b: i64) -> i64
    let fn_type = i64_type.fn_type(&[i64_type.into(), i64_type.into()], false);
    let function = module.add_function("add", fn_type, None);
    
    // Create entry block
    let entry = context.append_basic_block(function, "entry");
    builder.position_at_end(entry);
    
    // Get parameters
    let a = function.get_nth_param(0).unwrap().into_int_value();
    let b = function.get_nth_param(1).unwrap().into_int_value();
    
    // Build add instruction
    let sum = builder.build_int_add(a, b, "sum").unwrap();
    
    // Build return
    builder.build_return(Some(&sum)).unwrap();
    
    // Print the IR
    println!("{}", module.print_to_string().to_string());
}
```

---

## A.3 Code Generation Basics

### Types

```rust
let i64_type = context.i64_type();
let f64_type = context.f64_type();
let bool_type = context.bool_type();
let void_type = context.void_type();
let string_type = context.ptr_type(AddressSpace::default());  // char*
```

### Basic Blocks and Control Flow

```rust
// Create blocks for if/else
let then_block = context.append_basic_block(function, "then");
let else_block = context.append_basic_block(function, "else");
let merge_block = context.append_basic_block(function, "merge");

// Conditional branch
builder.build_conditional_branch(condition, then_block, else_block);

// Build then block
builder.position_at_end(then_block);
let then_value = /* compute value */;
builder.build_unconditional_branch(merge_block);

// Build else block
builder.position_at_end(else_block);
let else_value = /* compute value */;
builder.build_unconditional_branch(merge_block);

// Build merge block with phi node
builder.position_at_end(merge_block);
let phi = builder.build_phi(i64_type, "result").unwrap();
phi.add_incoming(&[(&then_value, then_block), (&else_value, else_block)]);
```

---

## A.4 Compiling Rusty to LLVM IR

```rust
struct CodeGen<'ctx> {
    context: &'ctx Context,
    module: Module<'ctx>,
    builder: Builder<'ctx>,
    variables: HashMap<String, PointerValue<'ctx>>,
}

impl<'ctx> CodeGen<'ctx> {
    fn compile_expr(&mut self, expr: &Expr) -> IntValue<'ctx> {
        match expr {
            Expr::Integer(n) => self.context.i64_type().const_int(*n as u64, true),
            
            Expr::Binary { left, op, right } => {
                let l = self.compile_expr(left);
                let r = self.compile_expr(right);
                
                match op {
                    BinaryOp::Add => self.builder.build_int_add(l, r, "add").unwrap(),
                    BinaryOp::Sub => self.builder.build_int_sub(l, r, "sub").unwrap(),
                    BinaryOp::Mul => self.builder.build_int_mul(l, r, "mul").unwrap(),
                    BinaryOp::Div => self.builder.build_int_signed_div(l, r, "div").unwrap(),
                    // ...
                }
            }
            
            Expr::Identifier(name) => {
                let ptr = self.variables.get(name).expect("undefined variable");
                self.builder.build_load(self.context.i64_type(), *ptr, name).unwrap()
                    .into_int_value()
            }
            // ...
        }
    }
}
```

---

## A.5 JIT Compilation

Execute code at runtime without writing files:

```rust
use inkwell::execution_engine::JitFunction;

fn jit_compile_and_run() {
    let context = Context::create();
    let module = context.create_module("jit");
    let builder = context.create_builder();
    
    // Build the add function
    // ... (as before)
    
    // Create JIT execution engine
    let execution_engine = module
        .create_jit_execution_engine(OptimizationLevel::Aggressive)
        .unwrap();
    
    // Get function pointer
    type AddFunc = unsafe extern "C" fn(i64, i64) -> i64;
    let add: JitFunction<AddFunc> = unsafe {
        execution_engine.get_function("add").unwrap()
    };
    
    // Call it!
    let result = unsafe { add.call(5, 3) };
    println!("5 + 3 = {}", result);  // Prints: 5 + 3 = 8
}
```

---

## A.6 Ahead-of-Time Compilation

Generate standalone executables:

```rust
use inkwell::targets::{
    CodeModel, FileType, InitializationConfig, RelocMode, Target, TargetMachine,
};

fn compile_to_executable(module: &Module) {
    Target::initialize_native(&InitializationConfig::default()).unwrap();
    
    let target_triple = TargetMachine::get_default_triple();
    let target = Target::from_triple(&target_triple).unwrap();
    
    let target_machine = target
        .create_target_machine(
            &target_triple,
            "generic",
            "",
            OptimizationLevel::Aggressive,
            RelocMode::Default,
            CodeModel::Default,
        )
        .unwrap();
    
    // Write object file
    target_machine
        .write_to_file(&module, FileType::Object, Path::new("output.o"))
        .unwrap();
    
    // Link with system linker
    // cc output.o -o output
}
```

---

## Exercises

### Exercise A.1: Arithmetic
Generate LLVM IR for arithmetic expressions.

### Exercise A.2: JIT REPL
Build a JIT-powered REPL for Rusty.

### Exercise A.3: Executable
Compile a Rusty program to a standalone executable.

### Exercise A.4: Optimization
Add optimization passes and measure performance difference.

---

## Resources

- [LLVM Language Reference](https://llvm.org/docs/LangRef.html)
- [Inkwell Documentation](https://thedan64.github.io/inkwell/inkwell/)
- [LLVM Kaleidoscope Tutorial](https://llvm.org/docs/tutorial/)
- [Rust LLVM Tutorial](https://github.com/acolite-d/llvm-tutorial-in-rust-using-inkwell)
