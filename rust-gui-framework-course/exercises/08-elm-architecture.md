# Exercise Sheet 08: Elm Architecture

> **Estimated Time:** 2 hours  
> **Difficulty:** ⭐⭐⭐

---

## Exercise 1: Counter App (25 min)

Implement a complete counter with Elm Architecture:
- Model: count value
- Messages: Increment, Decrement, Reset, SetValue(i32)
- View: display + buttons

Run it with your renderer from previous lectures.

---

## Exercise 2: Todo Application (40 min)

Build a todo list:
- Add todos via text input
- Toggle completion status
- Remove todos
- Filter: All, Active, Completed
- "Clear completed" button

---

## Exercise 3: View Diffing (30 min)

Implement basic diffing:
```rust
fn diff(old: &View, new: &View) -> Vec<Patch>;
enum Patch {
    Replace(Path, View),
    UpdateText(Path, String),
    UpdateAttr(Path, String, String),
}
```

Print patches when view changes to verify.

---

## Exercise 4: Commands (25 min)

Add async commands to your runtime:
```rust
enum Command<M> {
    None,
    Delay { ms: u64, then: M },
    Task(Box<dyn FnOnce() -> M + Send>),
}
```

Test: click button → show "Loading..." → after 2s → show "Done!"

---

## Theory Problems

### T1: Time-Travel Debugging
With Elm Architecture, we can record all messages.
1. How would you implement "undo"?
2. How would you implement "replay"?
3. What's the memory cost of storing N actions?

### T2: Pure Functions
Prove: If `update` and `view` are pure, the UI is deterministic given the message history.

### T3: Diffing Complexity
For a tree of N nodes:
1. What's worst-case diff complexity?
2. How does React's key optimization help?

---

## Challenge: Component System

Design a component abstraction:
```rust
trait Component {
    type Model;
    type Message;
    fn update(&mut self, msg: Self::Message);
    fn view(&self) -> Element<Self::Message>;
}
```

Compose components: `Parent { child: Child }` with message mapping.
