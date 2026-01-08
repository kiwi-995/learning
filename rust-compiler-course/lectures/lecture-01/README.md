# Lecture 1: Getting Started with Rust

> *"The journey of a thousand miles begins with a single step."* — Lao Tzu

Welcome to the first lecture of your Rust and compiler construction journey! In this lecture, you'll set up your development environment, write your first Rust programs, and understand why Rust is an excellent choice for building compilers.

## 🎯 Learning Objectives

By the end of this lecture, you will be able to:
- Explain why Rust is ideal for compiler development
- Set up a complete Rust development environment
- Write, compile, and run Rust programs
- Use variables, basic types, functions, and control flow

## 📖 Topics

### 1.1 Why Rust for Compilers?

Compilers are among the most demanding software to write. They must be:
- **Correct** — A compiler bug can break every program it compiles
- **Fast** — Compile times matter for developer productivity
- **Memory-safe** — No buffer overflows, use-after-free, or data races

Rust excels at all three:

| Requirement | How Rust Helps |
|-------------|----------------|
| Correctness | Strong type system catches errors at compile time |
| Performance | Zero-cost abstractions, no garbage collector |
| Memory Safety | Ownership system prevents memory bugs without runtime overhead |

**Famous compilers written in Rust:**
- `rustc` — The Rust compiler itself (self-hosting!)
- `swc` — Super-fast JavaScript/TypeScript compiler
- `deno` — JavaScript runtime with TypeScript support
- `oxc` — JavaScript parser faster than swc

### 1.2 Installing Rust

Rust uses `rustup` for installation and version management.

**macOS/Linux:**
```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

**Windows:**
Download and run [rustup-init.exe](https://rustup.rs/)

**Verify installation:**
```bash
rustc --version    # The Rust compiler
cargo --version    # The Rust package manager and build tool
```

**Recommended IDE:** VS Code with the `rust-analyzer` extension.

### 1.3 Hello, World!

Create your first Rust program:

```bash
cargo new hello_world
cd hello_world
```

This creates:
```
hello_world/
├── Cargo.toml    # Project configuration
└── src/
    └── main.rs   # Your source code
```

**src/main.rs:**
```rust
fn main() {
    println!("Hello, world!");
}
```

**Run it:**
```bash
cargo run
```

**Anatomy of the program:**
- `fn main()` — The entry point function
- `println!` — A macro (the `!` indicates it's a macro, not a function)
- Statements end with semicolons `;`

### 1.4 Variables and Mutability

In Rust, variables are **immutable by default**. This is a key design decision for safety.

```rust
fn main() {
    let x = 5;           // Immutable
    // x = 6;            // ERROR: cannot assign twice to immutable variable
    
    let mut y = 5;       // Mutable (explicit opt-in)
    y = 6;               // OK
    
    const MAX_POINTS: u32 = 100_000;  // Constant (must have type annotation)
}
```

**Why immutability by default?**
- Easier to reason about code
- Enables compiler optimizations
- Prevents accidental mutations
- Makes concurrency safer

**Shadowing** — You can declare a new variable with the same name:
```rust
fn main() {
    let x = 5;
    let x = x + 1;      // New variable, shadows the old one
    let x = x * 2;      // x is now 12
    println!("{}", x);  // Prints: 12
}
```

### 1.5 Data Types

Rust is **statically typed** — every variable has a type known at compile time.

#### Scalar Types

| Type | Description | Examples |
|------|-------------|----------|
| `i32`, `i64`, etc. | Signed integers | `-5`, `42` |
| `u32`, `u64`, etc. | Unsigned integers | `0`, `255` |
| `f32`, `f64` | Floating point | `3.14`, `-0.001` |
| `bool` | Boolean | `true`, `false` |
| `char` | Unicode character | `'a'`, `'❤'`, `'日'` |

```rust
fn main() {
    let integer: i32 = 42;
    let float: f64 = 3.14159;
    let boolean: bool = true;
    let character: char = '🦀';
    
    // Type inference usually works
    let inferred = 42;  // Rust infers i32
}
```

#### Strings

Rust has two main string types:

```rust
fn main() {
    // String literal (immutable, stored in binary)
    let s1: &str = "Hello";
    
    // String (growable, heap-allocated)
    let s2: String = String::from("Hello");
    let s3 = "Hello".to_string();  // Same as above
}
```

We'll explore the difference deeply in Lecture 2 when we cover ownership.

### 1.6 Functions

Functions are declared with `fn`:

```rust
fn main() {
    let result = add(5, 3);
    println!("5 + 3 = {}", result);
}

fn add(a: i32, b: i32) -> i32 {
    a + b  // No semicolon = this is the return value (expression)
}
```

**Key points:**
- Parameter types must be declared
- Return type comes after `->`
- The last expression (without `;`) is implicitly returned
- You can also use `return` explicitly

**Expressions vs. Statements:**
```rust
fn main() {
    // Statement: performs action, doesn't return value
    let x = 5;
    
    // Expression: evaluates to a value
    let y = {
        let temp = 3;
        temp + 1  // No semicolon — this block evaluates to 4
    };  // y = 4
    
    // If as expression
    let max = if x > y { x } else { y };
}
```

### 1.7 Control Flow

#### If Expressions

```rust
fn main() {
    let number = 7;
    
    if number < 5 {
        println!("less than 5");
    } else if number < 10 {
        println!("between 5 and 9");
    } else {
        println!("10 or greater");
    }
    
    // If as expression (must have else, branches must return same type)
    let description = if number % 2 == 0 { "even" } else { "odd" };
}
```

#### Loops

```rust
fn main() {
    // Infinite loop (use break to exit)
    let mut counter = 0;
    let result = loop {
        counter += 1;
        if counter == 10 {
            break counter * 2;  // loop can return a value!
        }
    };  // result = 20
    
    // While loop
    let mut n = 3;
    while n != 0 {
        println!("{}", n);
        n -= 1;
    }
    
    // For loop (most common for iteration)
    for i in 1..=5 {  // 1 to 5 inclusive
        println!("{}", i);
    }
    
    // Iterating over collections
    let arr = [10, 20, 30];
    for element in arr {
        println!("{}", element);
    }
}
```

### 1.8 Comments and Documentation

```rust
// Single-line comment

/*
 * Multi-line
 * comment
 */

/// Documentation comment (generates HTML docs)
/// This function adds two numbers.
/// 
/// # Examples
/// ```
/// let result = add(2, 3);
/// assert_eq!(result, 5);
/// ```
fn add(a: i32, b: i32) -> i32 {
    a + b
}
```

Generate documentation with:
```bash
cargo doc --open
```

---

## 🧠 Theory: Formal Languages Preview

Before we build a compiler, let's briefly understand what we're working with.

**Definition:** A *language* is a (possibly infinite) set of strings over some alphabet.

Examples:
- The language of all valid Rust programs
- The language of all valid JSON documents
- The language `{aⁿbⁿ | n ≥ 0}` = {`""`, `"ab"`, `"aabb"`, `"aaabbb"`, ...}

**Question to ponder:** Why can't a simple pattern like regex match `{aⁿbⁿ}`? (We'll prove this in Lecture 4!)

---

## ✅ Summary

| Concept | Syntax | Notes |
|---------|--------|-------|
| Variable | `let x = 5;` | Immutable by default |
| Mutable variable | `let mut x = 5;` | Explicit mutability |
| Constant | `const X: i32 = 5;` | Must have type, SCREAMING_CASE |
| Function | `fn name(x: i32) -> i32` | Types required |
| If | `if cond { } else { }` | No parentheses around condition |
| Loop | `loop { break; }` | Infinite, use break |
| While | `while cond { }` | Conditional loop |
| For | `for x in iter { }` | Iterator-based |

---

## 📚 Further Reading

- [The Rust Programming Language, Chapter 1-3](https://doc.rust-lang.org/book/ch01-00-getting-started.html)
- [Rust By Example](https://doc.rust-lang.org/rust-by-example/)
- [Rustlings exercises](https://github.com/rust-lang/rustlings) — Start with `variables` and `functions`

---

**Next:** [Lecture 2: Ownership, Borrowing & Memory Safety](../lecture-02/README.md)

**Exercises:** [Exercise Sheet 1](exercises.md)
