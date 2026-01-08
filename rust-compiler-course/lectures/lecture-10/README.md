# Lecture 10: The Road to Self-Hosting

> *"A self-hosting compiler is the ultimate test of a language's expressiveness."*

## 🎯 Learning Objectives

- Understand what self-hosting means
- Port the Rusty compiler from Rust to Rusty
- Achieve the first successful self-compilation
- Verify correctness through bootstrapping

## 📖 Topics

### 10.1 What is Self-Hosting?

A self-hosting compiler:
1. Is written in the same language it compiles
2. Can compile its own source code
3. Proves the language is powerful enough for real programs

Famous self-hosting examples: C, Go, Rust, OCaml

### 10.2 The Bootstrap Problem

**Chicken and egg:** How do you compile the first version?

**Solution:** Bootstrapping
1. Write compiler v1 in existing language (Rust for us)
2. Use v1 to compile compiler v2 (written in Rusty)
3. Use v2 to compile v3 from same source
4. Verify v2 and v3 are identical

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│ Rusty Compiler  │────▶│ Rusty Compiler  │────▶│ Rusty Compiler  │
│ (Rust source)   │     │ (Rusty source)  │     │ (Rusty source)  │
│     v1          │     │      v2         │     │      v3         │
└────────┬────────┘     └────────┬────────┘     └─────────────────┘
         │                       │                        ↑
    Rust Compiler           v1 compiles            v2 compiles
                            Rusty source           same source
                                                   
                            v2 == v3 ✓ (self-hosting verified!)
```

### 10.3 Language Requirements for Self-Hosting

The Rusty language needs:
- File I/O (read source files)
- String manipulation (parsing)
- Data structures (AST, symbol tables)
- Recursion (tree traversal)
- Error handling

### 10.4 Porting Strategy

1. Start with the simplest module (token types)
2. Port the lexer (string manipulation)
3. Port the parser (recursive descent)
4. Port code generation
5. Replace Rust-specific idioms with Rusty equivalents

### 10.5 Verification

```bash
# Compile v1 (bootstrap compiler)
cargo build --release

# Use v1 to compile Rusty source to v2
./target/release/rusty compile src/main.rusty -o rusty-v2

# Use v2 to compile same source to v3
./rusty-v2 compile src/main.rusty -o rusty-v3

# Compare binaries
diff rusty-v2 rusty-v3 && echo "Self-hosting verified! 🎉"
```

### 10.6 Optimization Basics

Once self-hosted, add optimizations:
- Constant folding
- Dead code elimination
- Common subexpression elimination

---

## 🎉 Congratulations!

You've built a complete, self-hosting compiler! This is a significant achievement that demonstrates deep understanding of:
- Rust programming
- Language design
- Compiler architecture
- Formal language theory

---

**Exercises:** [Exercise Sheet 10](exercises.md)
**Appendices:** [Appendix A: LLVM](../../appendices/appendix-a-llvm/README.md) | [Appendix B: Theory](../../appendices/appendix-b-theory/README.md)
