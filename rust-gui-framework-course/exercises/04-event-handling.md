# Exercise Sheet 04: Event Handling & Windowing

> **Estimated Time:** 1.5 hours  
> **Difficulty:** ⭐⭐

---

## Exercise 1: Input State Tracker (30 min)

Implement a complete `InputState` struct that tracks:
- Pressed keys (with `just_pressed` and `just_released`)
- Mouse position, buttons, and scroll
- Per-frame clearing

Test with: pressing WASD moves a rectangle, mouse click changes color.

---

## Exercise 2: Drawing Paint App (30 min)

Build a simple paint program:
- Left mouse button draws circles at cursor
- Right mouse button erases
- Keyboard shortcuts: C clears, 1-5 changes brush size
- Scroll wheel changes color hue

---

## Exercise 3: Keyboard Shortcuts (20 min)

Implement modifier key detection:
```rust
fn is_ctrl_pressed(&self) -> bool
fn is_shift_pressed(&self) -> bool
fn check_shortcut(&self, ctrl: bool, shift: bool, key: KeyCode) -> bool
```

Test: Ctrl+S prints "Save", Ctrl+Z prints "Undo"

---

## Exercise 4: Hit Testing (20 min)

Implement:
```rust
fn point_in_rect(px: f32, py: f32, rect: &Rect) -> bool
fn point_in_circle(px: f32, py: f32, cx: f32, cy: f32, r: f32) -> bool
```

Create clickable buttons that change appearance on hover and press.

---

## Theory Problems

### T1: Event Ordering
Given events: MouseDown, MouseMove, MouseMove, MouseUp
What state should `just_pressed` and `just_released` have after each?

### T2: High-DPI
On a 2x scale display with 800×600 logical window:
1. What's the physical pixel count?
2. If mouse reports (400, 300) logical, what's the physical position?
3. Why render at physical resolution but use logical for layout?

---

## Challenge: Gesture Detection

Implement simple gesture recognition:
- Tap: touch start + end within 100ms and 10px movement
- Swipe: detect direction (up/down/left/right)
- Pinch: two-finger distance change
