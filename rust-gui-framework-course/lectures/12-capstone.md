# Lecture 12: Capstone Project

> **Duration:** ~2+ hours (project work)  
> **Prerequisites:** Lectures 01-11  
> **Goal:** Build a complete application with your framework

---

## Overview

It's time to use everything you've learned! Choose one of three projects and build it from scratch using your GUI framework.

---

## Project Options

### Option A: Text Editor

Build a simple code/text editor with:
- Multi-line text editing
- Syntax highlighting (pick one language)
- Keyboard shortcuts (Ctrl+S, Ctrl+Z, Ctrl+C/V)
- Line numbers
- File open/save dialogs

**Core challenges:**
- Text input handling
- Efficient text storage (rope or gap buffer)
- Cursor positioning
- Selection handling

---

### Option B: Drawing Application  

Build a paint program with:
- Tools: brush, eraser, rectangle, circle, line
- Color picker
- Adjustable brush size
- Undo/redo (at least 20 levels)
- Export to PNG

**Core challenges:**
- Off-screen canvas rendering
- Tool state management
- Command pattern for undo
- Image export

---

### Option C: Dashboard Application

Build a data dashboard with:
- Multiple chart types (bar, line, pie)
- Data table view
- Configurable layout (drag to resize panels)
- Theme switching (light/dark)
- Mock data generation

**Core challenges:**
- Complex layout with resizable areas
- Chart rendering
- Theme system
- Data binding

---

## Project Requirements

Regardless of which project you choose:

### Must Have
- [ ] Works on at least one platform
- [ ] Responsive to window resize
- [ ] Keyboard accessible
- [ ] No crashes on normal usage
- [ ] Clean, readable code

### Should Have
- [ ] Web build works
- [ ] Animations for state changes
- [ ] Loading/error states
- [ ] Configuration persistence

### Nice to Have
- [ ] Full accessibility support
- [ ] Mobile-friendly
- [ ] Theme customization
- [ ] Performance optimizations

---

## Project Structure

```
my-app/
├── Cargo.toml
├── src/
│   ├── main.rs           # Entry point
│   ├── app.rs            # Application state & logic
│   ├── ui/
│   │   ├── mod.rs
│   │   ├── components/   # Reusable widgets
│   │   └── views/        # Screen compositions
│   ├── platform/         # Platform abstractions
│   └── utils/            # Helpers
├── assets/
│   ├── fonts/
│   └── icons/
└── README.md
```

---

## Milestones

### Week 1: Foundation
- [ ] Project setup with dependencies
- [ ] Window creation and rendering loop
- [ ] Basic input handling
- [ ] Core data structures

### Week 2: Core Features
- [ ] Main functionality implemented
- [ ] UI layout complete
- [ ] State management working
- [ ] File I/O (if applicable)

### Week 3: Polish
- [ ] Error handling
- [ ] Edge cases handled
- [ ] Performance optimization
- [ ] Documentation

---

## Deliverables

1. **Working Application**
   - Runs on at least one platform
   - Core functionality works

2. **Source Code**
   - Clean, commented code
   - Logical structure

3. **README**
   - How to build and run
   - Features list
   - Screenshots

4. **Reflection Document** (1-2 pages)
   - What worked well
   - What was challenging
   - Architecture decisions
   - What you'd do differently

---

## Evaluation Criteria

| Criteria | Weight | Description |
|----------|--------|-------------|
| Functionality | 30% | Does it work? |
| Code Quality | 25% | Clean, idiomatic Rust? |
| Architecture | 20% | Good separation of concerns? |
| UX/Polish | 15% | Is it pleasant to use? |
| Documentation | 10% | Can someone else build it? |

---

## Tips for Success

1. **Start small** - Get something working, then iterate
2. **Commit often** - Small, focused commits
3. **Test as you go** - Don't wait until the end
4. **Ask for help** - When stuck, take a break or ask
5. **Have fun** - You built a GUI framework!

---

## Resources

- Your work from Lectures 01-11
- [wgpu examples](https://github.com/gfx-rs/wgpu/tree/trunk/examples)
- [Iced examples](https://github.com/iced-rs/iced/tree/master/examples)
- [egui demos](https://www.egui.rs/)

---

## Congratulations! 🎉

You've completed the course! You now understand:
- How GPUs render graphics
- How layout engines work
- How declarative UI systems are built
- How to deploy cross-platform apps

Keep building, keep learning!
