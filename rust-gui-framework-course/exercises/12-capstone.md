# Exercise Sheet 12: Capstone Project

> **Estimated Time:** 10-20 hours  
> **Difficulty:** ⭐⭐⭐⭐⭐ (Full Project)

---

## Choose Your Project

Select ONE of these projects:

### A. Text Editor
- Multi-line editing
- Syntax highlighting
- Keyboard shortcuts
- File open/save

### B. Drawing App
- Multiple tools (brush, shapes)
- Color picker
- Undo/redo stack
- Export to PNG

### C. Dashboard
- Multiple chart types
- Data tables
- Resizable panels
- Theme switching

---

## Milestones

### Milestone 1: Skeleton (2-3 hours)
- [ ] Create project structure
- [ ] Window + render loop working
- [ ] Basic input handling
- [ ] Empty UI structure

### Milestone 2: Core Feature (4-6 hours)
- [ ] Main functionality works
- [ ] State management complete
- [ ] UI responds to input

### Milestone 3: Polish (3-4 hours)
- [ ] Edge cases handled
- [ ] Error handling
- [ ] Visual polish
- [ ] Performance acceptable

### Milestone 4: Documentation (1-2 hours)
- [ ] README with build instructions
- [ ] Code comments
- [ ] Reflection document

---

## Deliverables

1. **Working application** (runs without crashes)
2. **Source code** (clean, idiomatic Rust)
3. **README.md** (build instructions, screenshots)
4. **Reflection.md** (what worked, what was hard, lessons learned)

---

## Grading Rubric

| Area | Points | Criteria |
|------|--------|----------|
| Functionality | 30 | Does the core feature work? |
| Code Quality | 25 | Readable? Idiomatic? Well-structured? |
| Architecture | 20 | Separation of concerns? Extensible? |
| Polish | 15 | Edge cases? Error handling? Smooth UX? |
| Documentation | 10 | Could someone else run this? |

---

## Submission

Create a GitHub repository with:
```
my-capstone/
├── src/
├── assets/
├── Cargo.toml
├── README.md
└── Reflection.md
```

Include at least one screenshot in README.

---

## Hints

<details>
<summary>Text Editor: Data Structure</summary>
Consider using a rope or gap buffer for efficient text operations.
Start with a simple Vec<String> if needed.
</details>

<details>
<summary>Drawing App: Undo/Redo</summary>
Use the Command pattern: Vec<Command> for history, push on action, pop on undo.
</details>

<details>
<summary>Dashboard: Panel Layout</summary>
Store panel sizes as percentages. On drag, update both adjacent panels.
</details>

---

## Good luck! 🚀

You have all the skills needed. Take it one step at a time!
