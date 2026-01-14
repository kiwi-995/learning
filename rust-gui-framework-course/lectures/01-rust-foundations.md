# Lecture 01: Rust Foundations for GUI Programming

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Basic programming experience  
> **Goal:** Master the Rust concepts essential for building GUI frameworks

---

## Overview

Before we draw a single pixel, we need to understand how Rust's unique ownership system shapes GUI architecture. This lecture covers the Rust fundamentals that make GUI programming both challenging and elegant.

**Why this matters for GUIs:**
- UI trees have complex ownership (parent ↔ child relationships)
- Event handlers need to modify state (callbacks + ownership = tricky!)
- Performance matters: no garbage collector pauses during animations

---

## Part 1: Ownership in the Context of UI Trees (20 min)

### The Ownership Rules

```rust
// Rule 1: Each value has exactly one owner
let window = Window::new("My App");

// Rule 2: When the owner goes out of scope, the value is dropped
{
    let button = Button::new("Click me");
} // button is dropped here - memory freed!

// Rule 3: Ownership can be transferred (moved)
let container = Container::new();
container.add_child(button); // button moved INTO container
// button is no longer valid here!
```

### The GUI Ownership Challenge

In a typical UI tree:
```
Window
├── Header
│   └── Title
└── Content
    ├── Button
    └── Label
```

**Question:** Who "owns" the Button?
- The Content container owns it
- But what if we need to reference it from an event handler?

### References and Borrowing

```rust
// Immutable borrow: read-only access
fn render_widget(widget: &Widget) {
    println!("Rendering: {:?}", widget.bounds);
}

// Mutable borrow: read-write access
fn update_position(widget: &mut Widget, x: f32, y: f32) {
    widget.bounds.x = x;
    widget.bounds.y = y;
}

// The crucial rule: ONE mutable OR many immutable, never both
let mut button = Button::new("Click");
let r1 = &button;     // OK: immutable borrow
let r2 = &button;     // OK: multiple immutable borrows
// let r3 = &mut button; // ERROR! Can't borrow mutably while immutable borrows exist
```

---

## Part 2: Lifetimes for UI Components (20 min)

### What Are Lifetimes?

Lifetimes ensure references don't outlive the data they point to.

```rust
// This struct holds a reference - it needs a lifetime parameter
struct Label<'a> {
    text: &'a str,  // The label borrows text from somewhere
}

// The 'a says: "Label cannot outlive the string it references"
fn create_label<'a>(text: &'a str) -> Label<'a> {
    Label { text }
}

// Usage:
let message = String::from("Hello, GUI!");
let label = create_label(&message);
// label is valid as long as message exists
```

### Lifetime Elision (The Compiler Helps!)

Often, you don't need to write lifetimes explicitly:

```rust
// You write:
fn get_text(label: &Label) -> &str {
    label.text
}

// Compiler infers:
fn get_text<'a>(label: &'a Label) -> &'a str {
    label.text
}
```

### When Lifetimes Get Complex

```rust
// A widget that references its parent
struct ChildWidget<'parent> {
    parent: &'parent Container,
    name: String,
}

// Problem: What if parent is destroyed while child exists?
// Rust prevents this at compile time!
```

---

## Part 3: Traits for Widget Abstractions (20 min)

### Defining Widget Behavior

```rust
/// The core trait all widgets must implement
trait Widget {
    /// Calculate the minimum size this widget needs
    fn min_size(&self) -> Size;
    
    /// Draw the widget to the screen
    fn render(&self, canvas: &mut Canvas);
    
    /// Handle an input event, return true if consumed
    fn handle_event(&mut self, event: &Event) -> bool;
}

// Implementing for a concrete type
struct Button {
    label: String,
    bounds: Rect,
    is_pressed: bool,
}

impl Widget for Button {
    fn min_size(&self) -> Size {
        Size { width: 80.0, height: 30.0 }
    }
    
    fn render(&self, canvas: &mut Canvas) {
        let color = if self.is_pressed { DARK_BLUE } else { BLUE };
        canvas.draw_rect(self.bounds, color);
        canvas.draw_text(&self.label, self.bounds.center());
    }
    
    fn handle_event(&mut self, event: &Event) -> bool {
        match event {
            Event::MouseDown(pos) if self.bounds.contains(*pos) => {
                self.is_pressed = true;
                true // Event consumed
            }
            Event::MouseUp(_) => {
                self.is_pressed = false;
                true
            }
            _ => false // Event not consumed
        }
    }
}
```

### Trait Objects for Dynamic Dispatch

```rust
// A container that holds ANY widget
struct Container {
    children: Vec<Box<dyn Widget>>, // "dyn" = dynamic dispatch
}

impl Container {
    fn add<W: Widget + 'static>(&mut self, widget: W) {
        self.children.push(Box::new(widget));
    }
    
    fn render_all(&self, canvas: &mut Canvas) {
        for child in &self.children {
            child.render(canvas);
        }
    }
}

// Usage:
let mut container = Container { children: vec![] };
container.add(Button::new("OK"));
container.add(Label::new("Status: Ready"));
// Both stored as Box<dyn Widget>!
```

### Generics vs Trait Objects

| Approach | Pros | Cons |
|----------|------|------|
| Generics `<T: Widget>` | Zero-cost, inlined | Can't mix types |
| Trait Objects `dyn Widget` | Heterogeneous collections | Virtual dispatch overhead |

---

## Part 4: Smart Pointers for UI Graphs (25 min)

### The Problem with Trees

```rust
// This doesn't work - who owns what?
struct Widget {
    parent: ???,    // Reference to parent
    children: ???,  // Owned children
}
```

### `Rc<T>`: Reference Counting

```rust
use std::rc::Rc;

// Multiple owners of the same data
let shared_theme = Rc::new(Theme::dark());

let button1 = Button::with_theme(Rc::clone(&shared_theme));
let button2 = Button::with_theme(Rc::clone(&shared_theme));
// Both buttons share the same theme!
// Theme is dropped when ALL Rc's are dropped
```

### `RefCell<T>`: Interior Mutability

```rust
use std::cell::RefCell;

// Mutate through a shared reference
let counter = RefCell::new(0);

// Multiple places can mutate
*counter.borrow_mut() += 1;
*counter.borrow_mut() += 1;
println!("Count: {}", counter.borrow()); // 2
```

### The Classic Pattern: `Rc<RefCell<T>>`

```rust
use std::rc::Rc;
use std::cell::RefCell;

type WidgetRef = Rc<RefCell<dyn Widget>>;

struct Container {
    children: Vec<WidgetRef>,
}

impl Container {
    fn add_child(&mut self, child: WidgetRef) {
        self.children.push(child);
    }
    
    fn update_all(&self) {
        for child in &self.children {
            child.borrow_mut().update(); // Mutable access through Rc!
        }
    }
}

// Creating shared widgets
let button = Rc::new(RefCell::new(Button::new("Shared")));
container.add_child(Rc::clone(&button));
// We can still modify button elsewhere:
button.borrow_mut().set_label("Updated!");
```

### `Weak<T>`: Breaking Cycles

```rust
use std::rc::{Rc, Weak};

struct TreeNode {
    parent: Weak<RefCell<TreeNode>>,  // Weak reference UP
    children: Vec<Rc<RefCell<TreeNode>>>, // Strong references DOWN
}

// Parent → Child: Strong (Rc)
// Child → Parent: Weak (prevents cycles)
```

---

## Part 5: Callbacks and Closures (20 min)

### The Callback Challenge

```rust
// We want: when button is clicked, update a label
let label = Label::new("Not clicked");
let button = Button::new("Click me");

// How do we connect them?
button.on_click(|| {
    label.set_text("Clicked!"); // ERROR: can't capture label mutably
});
```

### Solution: Shared State

```rust
use std::rc::Rc;
use std::cell::RefCell;

let label = Rc::new(RefCell::new(Label::new("Not clicked")));
let label_clone = Rc::clone(&label);

button.on_click(move || {
    label_clone.borrow_mut().set_text("Clicked!");
});

// Storing the callback
struct Button {
    on_click: Option<Box<dyn Fn()>>,
}

impl Button {
    fn set_on_click<F: Fn() + 'static>(&mut self, callback: F) {
        self.on_click = Some(Box::new(callback));
    }
    
    fn click(&self) {
        if let Some(ref callback) = self.on_click {
            callback();
        }
    }
}
```

### The 'static Lifetime in Callbacks

```rust
// Why 'static? The callback might outlive the current scope
fn set_on_click<F: Fn() + 'static>(&mut self, callback: F)
//                        ^^^^^^^^
// Means: F cannot contain references to local variables
// Solution: Use owned data or Rc/Arc
```

---

## Summary: Patterns You'll Use Constantly

| Pattern | Use Case |
|---------|----------|
| `&T` / `&mut T` | Temporary access during rendering/events |
| `Box<dyn Trait>` | Owned, heap-allocated widgets |
| `Rc<T>` | Shared ownership (single-threaded) |
| `Arc<T>` | Shared ownership (multi-threaded) |
| `RefCell<T>` | Interior mutability |
| `Rc<RefCell<T>>` | Shared + mutable (the GUI workhorse) |
| `Weak<T>` | Parent references to break cycles |

---

## Key Takeaways

1. **Ownership flows down the tree**: Parents own children
2. **References flow up**: Children hold `Weak` refs to parents
3. **Callbacks need `'static`**: Use `Rc<RefCell<T>>` for shared state
4. **Traits enable polymorphism**: `dyn Widget` for mixed collections
5. **Interior mutability is essential**: `RefCell` lets us mutate through shared refs

---

## Up Next

In **Lecture 02**, we'll put these patterns to use by creating our first window and understanding how the GPU rendering pipeline works with wgpu!

---

## References

- [The Rust Book: Ownership](https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html)
- [The Rust Book: Smart Pointers](https://doc.rust-lang.org/book/ch15-00-smart-pointers.html)
- [Rust by Example: Traits](https://doc.rust-lang.org/rust-by-example/trait.html)
