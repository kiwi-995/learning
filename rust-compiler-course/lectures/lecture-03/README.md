# Lecture 3: Structs, Enums & Pattern Matching

> *"Sum types and pattern matching are the bread and butter of compilers."*

This lecture covers Rust's powerful type system features: structs for grouping data, enums for modeling variants, and pattern matching for control flow. These are *essential* for building compilers — our entire AST will be built with enums!

## 🎯 Learning Objectives

By the end of this lecture, you will be able to:
- Define and use structs to group related data
- Implement methods on structs with `impl` blocks
- Create enums with associated data
- Use pattern matching for exhaustive case handling
- Work with `Option<T>` and `Result<T, E>` for safe error handling

## 📖 Topics

### 3.1 Structs: Grouping Related Data

A struct is a custom type that groups named fields together:

```rust
struct Token {
    kind: TokenKind,
    text: String,
    line: usize,
    column: usize,
}

// Creating an instance
let token = Token {
    kind: TokenKind::Number,
    text: String::from("42"),
    line: 1,
    column: 5,
};

// Accessing fields
println!("Token at line {}: {}", token.line, token.text);
```

**Struct Update Syntax:**
```rust
let token2 = Token {
    text: String::from("3.14"),
    ..token  // Copy remaining fields from token
};
```

### 3.2 Tuple Structs and Unit Structs

```rust
// Tuple struct: named type, unnamed fields
struct Point(f64, f64);
struct Color(u8, u8, u8);

let origin = Point(0.0, 0.0);
let red = Color(255, 0, 0);

// Unit struct: no fields (useful for markers)
struct EOF;
```

### 3.3 Methods with `impl` Blocks

Methods are functions attached to a type:

```rust
struct Rectangle {
    width: f64,
    height: f64,
}

impl Rectangle {
    // Associated function (like static method) - no self
    fn new(width: f64, height: f64) -> Rectangle {
        Rectangle { width, height }
    }
    
    fn square(size: f64) -> Rectangle {
        Rectangle { width: size, height: size }
    }
    
    // Method - takes &self (immutable borrow)
    fn area(&self) -> f64 {
        self.width * self.height
    }
    
    // Method - takes &mut self (mutable borrow)
    fn scale(&mut self, factor: f64) {
        self.width *= factor;
        self.height *= factor;
    }
    
    // Method - takes self (moves ownership)
    fn into_square(self) -> Rectangle {
        let side = (self.width + self.height) / 2.0;
        Rectangle::square(side)
    }
}

fn main() {
    let r = Rectangle::new(10.0, 20.0);
    println!("Area: {}", r.area());
    
    let mut r2 = Rectangle::new(5.0, 5.0);
    r2.scale(2.0);
    
    let r3 = r.into_square();  // r is moved (consumed)
    // println!("{}", r.area());  // ERROR: r was moved
}
```

### 3.4 Enums: One of Many Variants

Enums represent a value that is one of several *variants*:

```rust
// Simple enum (like in other languages)
enum Direction {
    North,
    South,
    East,
    West,
}

let dir = Direction::North;

// Enums with data (powerful!)
enum Message {
    Quit,                         // No data
    Move { x: i32, y: i32 },     // Named fields (struct-like)
    Write(String),                // Unnamed field (tuple-like)
    ChangeColor(i32, i32, i32),  // Multiple unnamed fields
}

let msg1 = Message::Quit;
let msg2 = Message::Move { x: 10, y: 20 };
let msg3 = Message::Write(String::from("Hello"));
```

**This is incredibly powerful for compilers!** Here's a preview of how we'll model expressions:

```rust
enum Expr {
    Number(f64),
    String(String),
    Identifier(String),
    Binary {
        left: Box<Expr>,
        operator: BinaryOp,
        right: Box<Expr>,
    },
    Unary {
        operator: UnaryOp,
        operand: Box<Expr>,
    },
    Call {
        callee: Box<Expr>,
        arguments: Vec<Expr>,
    },
}

enum BinaryOp {
    Add, Sub, Mul, Div,
    Eq, Ne, Lt, Gt, Le, Ge,
    And, Or,
}
```

### 3.5 Pattern Matching with `match`

Pattern matching is how you work with enums:

```rust
fn describe_direction(dir: Direction) -> &'static str {
    match dir {
        Direction::North => "Going north!",
        Direction::South => "Going south!",
        Direction::East => "Going east!",
        Direction::West => "Going west!",
    }
}
```

**Match is exhaustive** — you must handle all variants:

```rust
fn process_message(msg: Message) {
    match msg {
        Message::Quit => {
            println!("Quitting");
        }
        Message::Move { x, y } => {
            println!("Moving to ({}, {})", x, y);
        }
        Message::Write(text) => {
            println!("Writing: {}", text);
        }
        Message::ChangeColor(r, g, b) => {
            println!("Changing color to RGB({}, {}, {})", r, g, b);
        }
    }
}
```

**Pattern matching features:**

```rust
fn analyze_number(n: i32) -> &'static str {
    match n {
        0 => "zero",
        1 => "one",
        2..=9 => "single digit",      // Range pattern
        10 | 20 | 30 => "round tens", // Multiple patterns
        n if n > 100 => "large",      // Match guard
        _ => "something else",         // Catch-all
    }
}

fn inspect_point(point: (i32, i32)) {
    match point {
        (0, 0) => println!("Origin"),
        (0, y) => println!("On y-axis at {}", y),
        (x, 0) => println!("On x-axis at {}", x),
        (x, y) => println!("At ({}, {})", x, y),
    }
}
```

### 3.6 The `Option<T>` Type

Rust has no `null`. Instead, use `Option<T>` to represent the presence or absence of a value:

```rust
enum Option<T> {
    Some(T),
    None,
}
```

```rust
fn find_first_even(numbers: &[i32]) -> Option<i32> {
    for &n in numbers {
        if n % 2 == 0 {
            return Some(n);
        }
    }
    None
}

fn main() {
    let numbers = vec![1, 3, 5, 6, 7];
    
    match find_first_even(&numbers) {
        Some(n) => println!("First even: {}", n),
        None => println!("No even numbers found"),
    }
    
    // Convenient methods
    let maybe_num = find_first_even(&numbers);
    
    // unwrap_or: provide default value
    let num = maybe_num.unwrap_or(0);
    
    // map: transform the inner value
    let doubled = maybe_num.map(|n| n * 2);
    
    // if let: convenient single-case match
    if let Some(n) = maybe_num {
        println!("Found: {}", n);
    }
}
```

### 3.7 The `Result<T, E>` Type

`Result` represents operations that can succeed or fail:

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

```rust
use std::fs;
use std::io;

fn read_username() -> Result<String, io::Error> {
    fs::read_to_string("username.txt")
}

fn main() {
    match read_username() {
        Ok(username) => println!("Hello, {}", username),
        Err(e) => println!("Failed to read username: {}", e),
    }
}
```

**The `?` Operator:**

The `?` operator propagates errors automatically:

```rust
fn read_and_parse() -> Result<i32, Box<dyn std::error::Error>> {
    let contents = fs::read_to_string("number.txt")?;  // Propagate IO error
    let num = contents.trim().parse()?;                 // Propagate parse error
    Ok(num)
}
```

### 3.8 Defining Custom Error Types

For our compiler, we'll define custom error types:

```rust
#[derive(Debug)]
enum CompileError {
    UnexpectedCharacter { char: char, line: usize, column: usize },
    UnterminatedString { start_line: usize },
    UnexpectedToken { expected: String, found: String },
    UndefinedVariable { name: String },
}

impl std::fmt::Display for CompileError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            CompileError::UnexpectedCharacter { char, line, column } => {
                write!(f, "Unexpected character '{}' at {}:{}", char, line, column)
            }
            CompileError::UnterminatedString { start_line } => {
                write!(f, "Unterminated string starting at line {}", start_line)
            }
            CompileError::UnexpectedToken { expected, found } => {
                write!(f, "Expected {}, found {}", expected, found)
            }
            CompileError::UndefinedVariable { name } => {
                write!(f, "Undefined variable '{}'", name)
            }
        }
    }
}

impl std::error::Error for CompileError {}
```

### 3.9 AST Preview: Modeling Language Constructs

Here's how we'll model our "Rusty" language AST:

```rust
/// A complete program
struct Program {
    statements: Vec<Stmt>,
}

/// Statements
enum Stmt {
    Let {
        name: String,
        mutable: bool,
        type_annotation: Option<Type>,
        initializer: Expr,
    },
    Expression(Expr),
    If {
        condition: Expr,
        then_branch: Box<Stmt>,
        else_branch: Option<Box<Stmt>>,
    },
    While {
        condition: Expr,
        body: Box<Stmt>,
    },
    Block(Vec<Stmt>),
    Function {
        name: String,
        params: Vec<(String, Type)>,
        return_type: Type,
        body: Box<Stmt>,
    },
    Return(Option<Expr>),
}

/// Expressions
enum Expr {
    // Literals
    Number(f64),
    String(String),
    Bool(bool),
    Identifier(String),
    
    // Operations
    Binary {
        left: Box<Expr>,
        op: BinaryOp,
        right: Box<Expr>,
    },
    Unary {
        op: UnaryOp,
        operand: Box<Expr>,
    },
    
    // Access and calls
    Call {
        callee: Box<Expr>,
        args: Vec<Expr>,
    },
    FieldAccess {
        object: Box<Expr>,
        field: String,
    },
    
    // Control flow as expressions
    If {
        condition: Box<Expr>,
        then_branch: Box<Expr>,
        else_branch: Box<Expr>,
    },
    Match {
        scrutinee: Box<Expr>,
        arms: Vec<MatchArm>,
    },
}

/// Types in our language
enum Type {
    I64,
    F64,
    Bool,
    String,
    Unit,
    Function { params: Vec<Type>, ret: Box<Type> },
    Struct(String),
    Enum(String),
}
```

Notice how `Box<Expr>` is used for recursive types — without the indirection, the struct would have infinite size!

---

## 🧠 Theory: Algebraic Data Types

Enums and structs are examples of *algebraic data types* (ADTs):

**Product Types (Structs):**
The number of possible values is the *product* of the fields:
```rust
struct Pair {
    a: bool,  // 2 values
    b: bool,  // 2 values
}
// Total: 2 × 2 = 4 possible values
```

**Sum Types (Enums):**
The number of possible values is the *sum* of the variants:
```rust
enum Either {
    Left(bool),   // 2 values
    Right(bool),  // 2 values
}
// Total: 2 + 2 = 4 possible values
```

**Combined:**
```rust
enum Option<bool> {
    Some(bool),  // 2 values
    None,        // 1 value
}
// Total: 2 + 1 = 3 possible values
```

This is why they're called "algebraic" — you can reason about them mathematically!

**Recursive Types:**
```rust
enum List<T> {
    Cons(T, Box<List<T>>),
    Nil,
}
// Infinite possible values! (but finite memory per node)
```

---

## ✅ Summary

| Concept | Syntax | Notes |
|---------|--------|-------|
| Struct | `struct Name { field: Type }` | Product type |
| Tuple struct | `struct Name(Type, Type)` | Named tuple |
| Enum | `enum Name { Variant }` | Sum type |
| Enum with data | `Variant(Type)` or `Variant { field: Type }` | Tagged union |
| Method | `fn method(&self)` | Takes self reference |
| Associated fn | `fn new() -> Self` | No self parameter |
| Match | `match val { Pattern => expr }` | Exhaustive |
| Option | `Some(val)` / `None` | No null |
| Result | `Ok(val)` / `Err(e)` | Error handling |
| ? operator | `expr?` | Propagate errors |

---

## 📚 Further Reading

- [The Rust Programming Language, Chapter 5-6](https://doc.rust-lang.org/book/ch05-00-structs.html)
- [Rust By Example: Custom Types](https://doc.rust-lang.org/rust-by-example/custom_types.html)
- [Error Handling in Rust](https://doc.rust-lang.org/book/ch09-00-error-handling.html)

---

**Next:** [Lecture 4: Language Design & Lexical Analysis](../lecture-04/README.md)

**Exercises:** [Exercise Sheet 3](exercises.md)
