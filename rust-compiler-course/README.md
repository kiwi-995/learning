# Rust & Compiler Construction Course

> **Learn Rust by building a self-hosting compiler from scratch**

A hands-on, project-based course inspired by [Nand2Tetris](https://www.nand2tetris.org/) and [Crafting Interpreters](https://craftinginterpreters.com/). You'll learn Rust programming while building **"Rusty"** — a simplified, Rust-inspired programming language that can eventually compile itself.

## 🎯 What You'll Build

```
┌─────────────────────────────────────────────────────────────┐
│                    The Rusty Compiler                       │
├─────────────────────────────────────────────────────────────┤
│  Source Code (.rusty)                                       │
│       ↓                                                     │
│  [Lexer] → Tokens                                           │
│       ↓                                                     │
│  [Parser] → Abstract Syntax Tree                            │
│       ↓                                                     │
│  [Type Checker] → Validated AST                             │
│       ↓                                                     │
│  [Compiler] → Bytecode                                      │
│       ↓                                                     │
│  [Virtual Machine] → Execution                              │
└─────────────────────────────────────────────────────────────┘
```

**Ultimate Goal:** A self-hosting compiler — the Rusty compiler, written in Rusty, compiling itself!

## 📚 Course Structure

| Part | Lectures | Topics |
|------|----------|--------|
| **I. Rust Foundations** | 1-3 | Variables, ownership, borrowing, structs, enums, pattern matching |
| **II. Compiler Frontend** | 4-6 | Language design, lexical analysis, parsing, ASTs |
| **III. Semantic Analysis** | 7 | Symbol tables, scoping, type checking |
| **IV. Code Generation** | 8-9 | Tree-walk interpreter, bytecode VM |
| **V. Self-Hosting** | 10 | Porting the compiler to Rusty |

**Bonus Appendices:**
- **Appendix A:** Hands-on LLVM backend with Inkwell
- **Appendix B:** Compiler theory deep-dive (automata, parsing theory, type theory, semantics)

## 🚀 Getting Started

### Prerequisites
- Basic programming experience (Python is fine!)
- CS/Math background helpful but not required
- No Rust or compiler experience needed

### Setup
1. [Install Rust](https://rustup.rs/)
2. Clone this repository
3. Start with [Lecture 1](lectures/lecture-01/README.md)

```bash
# Install Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Verify installation
rustc --version
cargo --version

# Navigate to the course
cd rust-compiler-course

# Start the Rusty project
cd rusty
cargo build
```

## 📁 Directory Structure

```
rust-compiler-course/
├── README.md                 # You are here
├── lectures/
│   ├── lecture-01/           # Getting Started with Rust
│   ├── lecture-02/           # Ownership, Borrowing & Memory
│   ├── lecture-03/           # Structs, Enums & Pattern Matching
│   ├── lecture-04/           # Language Design & Lexical Analysis
│   ├── lecture-05/           # Building a Lexer
│   ├── lecture-06/           # Parsing & ASTs
│   ├── lecture-07/           # Semantic Analysis & Type Checking
│   ├── lecture-08/           # Tree-Walk Interpreter
│   ├── lecture-09/           # Bytecode Virtual Machine
│   └── lecture-10/           # The Road to Self-Hosting
├── appendices/
│   ├── appendix-a-llvm/      # LLVM backend tutorial
│   └── appendix-b-theory/    # Compiler theory deep-dive
├── rusty/                    # The compiler (written in Rust)
│   ├── Cargo.toml
│   └── src/
├── rusty-in-rusty/           # The self-hosted compiler (written in Rusty)
│   └── src/
└── resources/                # Additional materials and links
```

## 🦀 The "Rusty" Language

Rusty is a simplified, Rust-inspired language designed for learning:

```rust
// Rusty example
fn fibonacci(n: i64) -> i64 {
    match n {
        0 => 0,
        1 => 1,
        _ => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

fn main() {
    let mut i = 0;
    while i < 10 {
        print(fibonacci(i));
        i = i + 1;
    }
}
```

**Features:**
- ✅ Variables with `let` and `let mut`
- ✅ Basic types: `i64`, `f64`, `bool`, `String`
- ✅ Functions with explicit types
- ✅ Structs and enums with data
- ✅ Pattern matching with `match`
- ✅ Control flow: `if`, `while`, `for`, `loop`
- ✅ Simplified ownership (move + explicit clone)
- ⚠️ Basic generics and traits
- ❌ Lifetimes (uses GC instead)
- ❌ Macros

## ⏱️ Time Investment

- **Core Course:** 10 lectures × 1-2 hours each = ~15-20 hours
- **Exercises:** ~1-2 hours per lecture = ~10-20 hours
- **Appendix A (LLVM):** ~3-4 hours
- **Appendix B (Theory):** ~4-6 hours

**Total:** ~30-50 hours for comprehensive mastery

## 📖 Resources

This course synthesizes content from:

- [The Rust Programming Language](https://doc.rust-lang.org/book/) — The official Rust guide
- [Rustlings](https://github.com/rust-lang/rustlings) — Interactive Rust exercises
- [Crafting Interpreters](https://craftinginterpreters.com/) — Excellent compiler tutorial
- [Create Your Own Programming Language](https://createlang.rs/) — Rust-specific compiler patterns

## 📜 License

This course is for educational purposes. Feel free to use and modify for your own learning!

---

**Ready to start?** → [Lecture 1: Getting Started with Rust](lectures/lecture-01/README.md)
