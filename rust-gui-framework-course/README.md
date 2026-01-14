# Building GUI Frameworks in Rust

A comprehensive, Nand2Tetris-style course teaching the construction of cross-platform GUI frameworks from first principles.

## Course Overview

**Goal:** Build a complete, cross-platform GUI framework from scratch while deeply understanding every layer—from GPU rendering to declarative UI patterns.

**Prerequisites:**
- Basic Rust knowledge (variables, functions, control flow)
- Programming experience in any language
- Curiosity about how UI systems work!

## Course Structure

| Lecture | Topic | Focus |
|---------|-------|-------|
| 01 | [Rust Foundations](lectures/01-rust-foundations.md) | Ownership, traits, smart pointers for UI |
| 02 | [GPU Pipeline & wgpu](lectures/02-wgpu-fundamentals.md) | Modern GPU rendering basics |
| 03 | [2D Rendering](lectures/03-2d-rendering.md) | Shapes, shaders, primitives |
| 04 | [Event Handling](lectures/04-event-handling.md) | Input, windowing, cross-platform |
| 05 | [Widget Architecture](lectures/05-widget-architecture.md) | Immediate vs retained mode |
| 06 | [Layout Engine I](lectures/06-layout-fundamentals.md) | Box model, sizing, constraints |
| 07 | [Layout Engine II](lectures/07-flexbox-constraints.md) | Flexbox, Cassowary algorithm |
| 08 | [Declarative UI](lectures/08-elm-architecture.md) | Elm Architecture, view diffing |
| 09 | [Reactive State](lectures/09-reactive-signals.md) | Signals, effects, fine-grained reactivity |
| 10 | [Text & Graphics](lectures/10-text-graphics.md) | Font rendering, custom drawing |
| 11 | [Cross-Platform](lectures/11-cross-platform.md) | Web, mobile, desktop deployment |
| 12 | [Capstone Project](lectures/12-capstone.md) | Build a complete application |

## Exercise Sheets

Each lecture has a corresponding exercise sheet in the `exercises/` folder with:
- Coding challenges (progressively harder)
- Theory problems (proofs, analysis)
- Mini-projects

## Getting Started

```bash
# Set up Rust (if needed)
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Create your project
cargo new my-gui-framework
cd my-gui-framework

# Follow along with Lecture 01!
```

## Key Resources

- [Learn Wgpu](https://sotrh.github.io/learn-wgpu/) - GPU rendering foundation
- [The Rust Book](https://doc.rust-lang.org/book/) - Language fundamentals
- [Elm Architecture Guide](https://guide.elm-lang.org/architecture/) - Declarative patterns

## Framework We'll Build

By the end of this course, you'll have built a framework supporting:
- ✅ GPU-accelerated 2D rendering
- ✅ Flexbox-style layout engine
- ✅ Declarative, reactive UI patterns
- ✅ Cross-platform deployment (Web, Desktop)
- ✅ Custom widgets and styling
