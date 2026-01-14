# Exercise Sheet 05: Widget Architecture

> **Estimated Time:** 1.5-2 hours  
> **Difficulty:** ⭐⭐⭐

---

## Exercise 1: Immediate Mode Button (30 min)

Implement an immediate mode UI system:

```rust
struct ImContext { /* mouse state, hot/active widget */ }

impl ImContext {
    fn button(&mut self, id: u32, label: &str, x: f32, y: f32) -> bool;
    fn label(&mut self, text: &str, x: f32, y: f32);
    fn checkbox(&mut self, id: u32, checked: &mut bool, label: &str, x: f32, y: f32);
}
```

Build a counter app: +/- buttons and label showing count.

---

## Exercise 2: Widget Trait Implementation (35 min)

Define and implement:
```rust
trait Widget {
    fn measure(&self) -> Size;
    fn layout(&mut self, bounds: Rect);
    fn event(&mut self, event: &Event) -> bool;
    fn render(&self, canvas: &mut Canvas);
}
```

Implement: Label, Button, Checkbox, TextInput

---

## Exercise 3: Container Widgets (30 min)

Implement:
- `VStack`: Stack children vertically with spacing
- `HStack`: Stack children horizontally
- `ZStack`: Layer children (back to front)

Test with nested containers.

---

## Exercise 4: Widget Builder Pattern (25 min)

Create a fluent API:
```rust
Button::new("Click")
    .width(100.0)
    .height(40.0)
    .color(Color::BLUE)
    .on_click(|| println!("Clicked!"))
```

---

## Theory Problems

### T1: Compare Paradigms
For a color picker with RGB sliders and preview:
1. Sketch the immediate mode implementation
2. Sketch the retained mode implementation
3. Which requires more lines of code? Which is more maintainable?

### T2: Event Propagation
In a nested button inside a scrollable container:
1. Should scroll or click take priority?
2. Design an event propagation strategy

---

## Challenge: Hot Reload UI

Design (pseudocode) a system where widget definitions can be hot-reloaded at runtime without restarting the application.
