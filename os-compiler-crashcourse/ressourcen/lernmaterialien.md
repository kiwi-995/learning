# 📚 Lernmaterialien & Ressourcen

Eine kuratierte Sammlung der besten kostenlosen Ressourcen für diesen Kurs.

---

## 🖥️ OS-Entwicklung

### Primäre Ressourcen

| Ressource | Beschreibung | Link |
|-----------|--------------|------|
| **RISC-V OS in 1000 Lines** | Exzellentes Tutorial für RISC-V Kernel | [operating-system-in-1000-lines.vercel.app](https://operating-system-in-1000-lines.vercel.app/) |
| **OSDev Wiki** | Die Bibel der OS-Entwicklung | [wiki.osdev.org](https://wiki.osdev.org/) |
| **Writing an OS in Rust** | Moderne OS-Entwicklung mit Rust | [os.phil-opp.com](https://os.phil-opp.com/) |

### Weiterführend

- **xv6** - Einfaches UNIX-ähnliches Lehr-OS (MIT)
- **Bare Metal ARM** - [github.com/umanovskis/baremetal-arm](https://github.com/umanovskis/baremetal-arm)
- **Operating Systems: Three Easy Pieces** - [ostep.org](https://pages.cs.wisc.edu/~remzi/OSTEP/)

---

## 🏗️ Compiler & Interpreter

### Primäre Ressourcen

| Ressource | Beschreibung | Link |
|-----------|--------------|------|
| **Crafting Interpreters** | DAS Buch für Compiler-Anfänger | [craftinginterpreters.com](https://craftinginterpreters.com/) |
| **Build Your Own Lisp** | Lisp-Interpreter in C | [buildyourownlisp.com](http://www.buildyourownlisp.com/) |

### Weiterführend

- **Dragon Book** - Compilers: Principles, Techniques and Tools (Klassiker)
- **Engineering a Compiler** - Cooper & Torczon
- **Writing An Interpreter In Go** - Thorsten Ball

---

## 💻 Programmiersprachen

### C

| Ressource | Beschreibung | Link |
|-----------|--------------|------|
| **Beej's Guide to C** | Exzellentes C-Tutorial | [beej.us/guide/bgc](https://beej.us/guide/bgc/) |
| **learn-c.org** | Interaktives Tutorial | [learn-c.org](https://learn-c.org) |
| **Modern C** | Kostenloses Buch | [inria.hal.science](https://inria.hal.science/hal-02383654/document) |

### Rust

| Ressource | Beschreibung | Link |
|-----------|--------------|------|
| **The Rust Book** | Offizielles Tutorial | [doc.rust-lang.org/book](https://doc.rust-lang.org/book/) |
| **Rustlings** | Interaktive Übungen | [github.com/rust-lang/rustlings](https://github.com/rust-lang/rustlings) |
| **Rust by Example** | Code-basiertes Lernen | [doc.rust-lang.org/rust-by-example](https://doc.rust-lang.org/rust-by-example/) |

### C++

| Ressource | Beschreibung | Link |
|-----------|--------------|------|
| **Learn C++** | Umfassendes Tutorial | [learncpp.com](https://www.learncpp.com/) |
| **C++ Reference** | Referenz-Dokumentation | [cppreference.com](https://en.cppreference.com/) |

### Ruby

| Ressource | Beschreibung | Link |
|-----------|--------------|------|
| **Learn Ruby Online** | Interaktives Tutorial | [learnrubyonline.org](https://www.learnrubyonline.org/) |
| **Ruby in 20 Minutes** | Schnellstart | [ruby-lang.org/quickstart](https://www.ruby-lang.org/en/documentation/quickstart/) |

---

## 🔧 Tools & Setup

### RISC-V Toolchain

```bash
# macOS
brew tap riscv-software-src/riscv
brew install riscv-gnu-toolchain

# Alternative: Fertige Binaries
# github.com/riscv-collab/riscv-gnu-toolchain/releases
```

### QEMU

```bash
# macOS
brew install qemu

# Linux
sudo apt install qemu-system-riscv64
```

### Rust

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
rustup target add riscv64gc-unknown-none-elf
```

---

## 📖 Bücher & Papers

### Kostenlos Verfügbar

1. **Operating Systems: Three Easy Pieces** - Remzi Arpaci-Dusseau
2. **RISC-V Reader** - Patterson & Waterman
3. **Crafting Interpreters** - Robert Nystrom
4. **Modern C** - Jens Gustedt
5. **The Rust Book** - Steve Klabnik & Carol Nichols

### Klassiker (nicht kostenlos, aber empfehlenswert)

1. **The C Programming Language** - K&R
2. **Computer Systems: A Programmer's Perspective** - Bryant & O'Hallaron
3. **Compilers: Principles, Techniques and Tools** - Aho et al.
4. **Operating System Concepts** - Silberschatz

---

## 🎬 Video-Kurse

| Kurs | Thema | Link |
|------|-------|------|
| **Nand2Tetris** | Hardware → Software | [nand2tetris.org](https://www.nand2tetris.org/) |
| **CS50** | Intro CS mit C | [cs50.harvard.edu](https://cs50.harvard.edu/) |
| **LF110x** | RISC-V Intro | [Linux Foundation](https://training.linuxfoundation.org/training/introduction-to-riscv-lfd110x/) |

---

## 🌐 Communities

- **r/osdev** - Reddit OS Development Community
- **OSDev Forums** - [forum.osdev.org](https://forum.osdev.org/)
- **Rust Users Forum** - [users.rust-lang.org](https://users.rust-lang.org/)
- **RISC-V Discord** - [discord.gg/riscv](https://discord.gg/riscv)

---

## 📋 RISC-V Spezifikationen

| Dokument | Beschreibung |
|----------|--------------|
| **RISC-V ISA Manual** | Instruction Set Architecture |
| **RISC-V Privileged Spec** | Privilege Levels, CSRs, Traps |
| **QEMU Virt Machine** | Memory Map, Devices |

Alle verfügbar auf [riscv.org/specifications](https://riscv.org/specifications/)

---

## 🛠️ Nützliche GitHub Repos

```
# RISC-V
github.com/riscv/riscv-gnu-toolchain
github.com/riscv/riscv-isa-manual
github.com/mikeroyal/RISC-V-Guide

# OS Examples
github.com/mit-pdos/xv6-riscv
github.com/rust-embedded/rust-raspberrypi-OS-tutorials

# Compiler Examples
github.com/munificent/craftinginterpreters
```

---

*Viel Erfolg beim Lernen!*
