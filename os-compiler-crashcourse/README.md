# 🚀 7-Day OS & Compiler Crash Course

**From Nand2Tetris to Real Hardware: Build Your Own OS and Compiler**

> A intensive 7-day crash course to learn OS development, compiler design, and systems programming languages (C, C++, Rust, Ruby) from the ground up.

---

## 📋 Prerequisites

- **Completed or familiar with**: Nand2Tetris (or similar hardware-to-software understanding)
- **Programming experience**: Any language (you know Python)
- **Hardware**: Mac with QEMU/UTM installed
- **Time**: ~8-12 hours per day for 7 days

---

## 🎯 Learning Objectives

By the end of this course, you will:

1. **Understand** how computers boot and execute code at the lowest level
2. **Build** a minimal RISC-V kernel that boots in QEMU
3. **Implement** a terminal-based OS with keyboard input and screen output
4. **Create** a simple expression compiler/interpreter
5. **Know** the basics of C, C++, Rust, and Ruby
6. **Understand** framebuffer graphics for basic GUI

---

## 📅 Course Schedule

| Day | Topic | Languages | Key Deliverable |
|-----|-------|-----------|-----------------|
| **1** | Foundations: C & Hardware Bridge | C | "Hello World" in C, memory model understanding |
| **2** | Bare Metal RISC-V | C, Assembly | First boot in QEMU, UART output |
| **3** | Kernel Basics | C | Memory management, interrupt handling |
| **4** | Terminal OS | C | Keyboard input, shell prompt |
| **5** | Compiler Fundamentals | Ruby | Lexer and parser for expressions |
| **6** | Code Generation | C++, Ruby | AST evaluation, bytecode intro |
| **7** | GUI & Rust Introduction | Rust, C | Framebuffer graphics, Rust basics |

---

## 📁 Course Structure

```
os-compiler-crashcourse/
├── README.md                 # This file
├── vorlesungen/              # Lecture notes (theory + code examples)
│   ├── tag-01-foundations.md
│   ├── tag-02-bare-metal.md
│   ├── tag-03-kernel.md
│   ├── tag-04-terminal.md
│   ├── tag-05-compiler.md
│   ├── tag-06-codegen.md
│   └── tag-07-gui-rust.md
├── uebungen/                 # Exercise sheets with tests
│   ├── uebung-01.md
│   ├── uebung-02.md
│   └── ...
├── projekte/                 # Project templates and starter code
│   ├── mini-kernel/
│   ├── expression-compiler/
│   └── framebuffer-gui/
└── ressourcen/               # Additional resources and references
    └── lernmaterialien.md
```

---

## 🛠️ Required Tools

### Install Before Starting

```bash
# QEMU for RISC-V emulation
brew install qemu

# RISC-V toolchain
brew tap riscv-software-src/riscv
brew install riscv-tools

# Rust (for Day 7)
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Ruby (usually pre-installed on Mac)
ruby --version

# C/C++ compiler (Xcode Command Line Tools)
xcode-select --install
```

### Optional: UTM
If you prefer a GUI for QEMU, install [UTM](https://mac.getutm.app/) - it uses QEMU under the hood.

---

## 📚 Core Reference Materials

### OS Development
- **RISC-V OS in 1000 Lines** - [operating-system-in-1000-lines.vercel.app](https://operating-system-in-1000-lines.vercel.app/)
- **ARM Bare Metal** - [github.com/umanovskis/baremetal-arm](https://github.com/umanovskis/baremetal-arm)
- **OSDev Wiki** - [wiki.osdev.org](https://wiki.osdev.org/)

### Compiler Design
- **Crafting Interpreters** - [craftinginterpreters.com](https://craftinginterpreters.com/) (FREE)

### Language Learning
- **C**: [learn-c.org](https://learn-c.org/) | Beej's Guide to C
- **Rust**: [The Rust Book](https://doc.rust-lang.org/book/) | Rustlings
- **Ruby**: [learnrubyonline.org](https://www.learnrubyonline.org/)
- **C++**: [learncpp.com](https://www.learncpp.com/)

---

## ⚡ Quick Start

1. Clone/download this course
2. Install required tools (see above)
3. Start with `vorlesungen/tag-01-foundations.md`
4. Complete the exercises in `uebungen/uebung-01.md`
5. Proceed to the next day

---

## 🎓 Philosophy

This course follows the **Nand2Tetris approach**:

1. **Theory + Practice**: Every concept is immediately applied
2. **Bottom-up**: Start from the lowest level and build up
3. **Write everything yourself**: Code examples show you what to build, but YOU write the code
4. **Minimal viable features**: Get something working, then improve

> **Important**: The lectures contain code snippets and explanations, but exercises require you to write all code yourself. Tests are provided to verify your implementations.

---

## 📖 How to Use This Course

### Each Day Includes:

1. **📖 Lecture** (`vorlesungen/tag-XX.md`)
   - Theory and concepts explained
   - Code examples showing the approach
   - Architecture diagrams and visualizations

2. **✍️ Exercises** (`uebungen/uebung-XX.md`)
   - Practical tasks to implement yourself
   - Test cases to verify your work
   - Hints and debugging tips

3. **🔧 Project Work** (`projekte/`)
   - Starter templates and Makefiles
   - Your code goes here!

### Suggested Daily Schedule

| Time | Activity |
|------|----------|
| 2h | Read lecture notes, understand theory |
| 1h | Study code examples, trace through them |
| 4-6h | Complete exercises (the bulk of learning!) |
| 1-2h | Explore further, read additional resources |

---

## 🏆 Final Project Ideas

After completing all 7 days, you'll have the foundation to:

1. **Extend your OS**: Add a filesystem, multitasking, or network stack
2. **Build a full compiler**: Add variables, functions, control flow
3. **Create a GUI**: Build a window manager with mouse support
4. **Port to real hardware**: Run on a real RISC-V board

---

*Let's build something amazing! Start with Day 1 →*
