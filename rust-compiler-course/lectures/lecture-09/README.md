# Lecture 9: Bytecode Virtual Machine

> *"Bytecode is the sweet spot between interpretation and compilation."*

## 🎯 Learning Objectives

- Design a bytecode instruction set
- Implement a stack-based virtual machine
- Compile AST to bytecode
- Implement garbage collection basics

## 📖 Topics

### 9.1 Why Bytecode?

- Faster than tree-walking (compact, cache-friendly)
- Portable across platforms
- Foundation for JIT compilation

### 9.2 Instruction Set

```rust
#[derive(Debug, Clone, Copy)]
pub enum OpCode {
    // Stack
    Constant(u16),  // Push constant from pool
    Pop,
    
    // Arithmetic
    Add, Sub, Mul, Div, Neg,
    
    // Comparison
    Equal, NotEqual, Less, Greater,
    
    // Control flow
    Jump(u16),
    JumpIfFalse(u16),
    
    // Variables
    GetLocal(u16),
    SetLocal(u16),
    GetGlobal(u16),
    SetGlobal(u16),
    
    // Functions
    Call(u8),
    Return,
}
```

### 9.3 The VM

```rust
pub struct VM {
    chunk: Chunk,
    ip: usize,
    stack: Vec<Value>,
    globals: HashMap<String, Value>,
}

impl VM {
    fn run(&mut self) -> Result<Value, RuntimeError> {
        loop {
            let op = self.chunk.code[self.ip];
            self.ip += 1;
            
            match op {
                OpCode::Constant(idx) => {
                    let value = self.chunk.constants[idx as usize].clone();
                    self.stack.push(value);
                }
                OpCode::Add => {
                    let b = self.pop();
                    let a = self.pop();
                    self.stack.push(a + b);
                }
                OpCode::Return => {
                    return Ok(self.pop());
                }
                // ...
            }
        }
    }
}
```

### 9.4 Compilation

Transform AST to bytecode:

```rust
fn compile_expr(&mut self, expr: &Expr) {
    match expr {
        Expr::Integer(n) => {
            let idx = self.add_constant(Value::Integer(*n));
            self.emit(OpCode::Constant(idx));
        }
        Expr::Binary { left, op, right } => {
            self.compile_expr(left);
            self.compile_expr(right);
            self.emit(match op {
                BinaryOp::Add => OpCode::Add,
                BinaryOp::Sub => OpCode::Sub,
                // ...
            });
        }
        // ...
    }
}
```

---

## 🧠 Theory: Stack Machines vs Register Machines

- **Stack machine**: Operations pop/push stack (simpler, more instructions)
- **Register machine**: Operations use virtual registers (fewer instructions)

---

**Exercises:** [Exercise Sheet 9](exercises.md)
**Next:** [Lecture 10: The Road to Self-Hosting](../lecture-10/README.md)
