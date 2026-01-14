# Exercise Sheet 01: Rust Foundations for GUI Programming

> **Estimated Time:** 1.5-2 hours  
> **Difficulty:** ⭐⭐ (Intermediate Rust)

---

## Exercise 1: Implement a Widget Trait (30 min)

Create a basic widget system with the following requirements:

### 1.1 Define the Core Types

```rust
// TODO: Define these types
struct Size { /* width, height */ }
struct Rect { /* x, y, width, height */ }
struct Color { /* r, g, b, a as u8 */ }

// A simple canvas for drawing
struct Canvas {
    commands: Vec<DrawCommand>,
}

enum DrawCommand {
    Rect { bounds: Rect, color: Color },
    Text { text: String, position: (f32, f32) },
}
```

### 1.2 Define the Widget Trait

```rust
trait Widget {
    /// Returns the minimum size this widget needs
    fn min_size(&self) -> Size;
    
    /// Renders the widget to the canvas
    fn render(&self, canvas: &mut Canvas);
}
```

### 1.3 Implement Three Widgets

Implement the `Widget` trait for:

1. **`Label`**: Displays text
   - min_size: 100x20
   - render: outputs a `Text` command

2. **`ColorBox`**: A colored rectangle
   - min_size: 50x50  
   - render: outputs a `Rect` command

3. **`Spacer`**: Empty space
   - min_size: configurable width/height
   - render: does nothing

### Expected Output

```rust
fn main() {
    let label = Label::new("Hello, GUI!");
    let box_ = ColorBox::new(Color::rgb(255, 0, 0));
    let spacer = Spacer::new(10.0, 10.0);
    
    let mut canvas = Canvas::new();
    label.render(&mut canvas);
    box_.render(&mut canvas);
    spacer.render(&mut canvas);
    
    println!("Commands: {:?}", canvas.commands);
    // Should show: [Text { text: "Hello, GUI!", ... }, Rect { ... }]
}
```

---

## Exercise 2: Build a Tree Structure (30 min)

Create a widget tree that supports parent-child relationships.

### 2.1 Container Widget

```rust
struct Container {
    children: Vec<Box<dyn Widget>>,
    padding: f32,
}

impl Container {
    fn new() -> Self { /* ... */ }
    fn add<W: Widget + 'static>(&mut self, child: W) { /* ... */ }
}

impl Widget for Container {
    fn min_size(&self) -> Size {
        // Calculate based on children sizes + padding
        // Stack children vertically
    }
    
    fn render(&self, canvas: &mut Canvas) {
        // Render all children, offset by position
    }
}
```

### 2.2 Test Your Tree

```rust
fn main() {
    let mut root = Container::new();
    root.add(Label::new("Title"));
    root.add(ColorBox::new(Color::rgb(0, 128, 255)));
    root.add(Label::new("Footer"));
    
    println!("Total size: {:?}", root.min_size());
    // Should calculate combined height + padding
}
```

---

## Exercise 3: Event Propagation (30 min)

Add event handling with bubbling/capturing.

### 3.1 Define Events

```rust
#[derive(Debug, Clone)]
enum Event {
    MouseDown { x: f32, y: f32 },
    MouseUp { x: f32, y: f32 },
    MouseMove { x: f32, y: f32 },
    KeyPress { key: char },
}

#[derive(Debug, Clone, Copy, PartialEq)]
enum EventResult {
    Handled,    // Stop propagation
    Ignored,    // Continue to next widget
}
```

### 3.2 Add Event Handling to Widget Trait

```rust
trait Widget {
    fn min_size(&self) -> Size;
    fn render(&self, canvas: &mut Canvas);
    fn handle_event(&mut self, event: &Event) -> EventResult;
}
```

### 3.3 Implement a Button

```rust
struct Button {
    label: String,
    bounds: Rect,
    is_hovered: bool,
    is_pressed: bool,
}

impl Widget for Button {
    fn handle_event(&mut self, event: &Event) -> EventResult {
        match event {
            Event::MouseMove { x, y } => {
                self.is_hovered = self.bounds.contains(*x, *y);
                EventResult::Ignored
            }
            Event::MouseDown { x, y } if self.bounds.contains(*x, *y) => {
                self.is_pressed = true;
                println!("Button '{}' pressed!", self.label);
                EventResult::Handled
            }
            Event::MouseUp { .. } => {
                if self.is_pressed {
                    println!("Button '{}' clicked!", self.label);
                }
                self.is_pressed = false;
                EventResult::Handled
            }
            _ => EventResult::Ignored
        }
    }
    // ... other methods
}
```

### 3.4 Implement Event Propagation in Container

The container should:
1. Pass events to children in reverse order (top-most first)
2. Stop if any child returns `EventResult::Handled`

---

## Exercise 4: Shared State with Callbacks (30 min)

Implement a reactive counter using `Rc<RefCell<T>>`.

### 4.1 The Counter App

```rust
use std::rc::Rc;
use std::cell::RefCell;

struct CounterApp {
    count: Rc<RefCell<i32>>,
    label: Rc<RefCell<Label>>,
    increment_btn: Button,
    decrement_btn: Button,
}

impl CounterApp {
    fn new() -> Self {
        let count = Rc::new(RefCell::new(0));
        let label = Rc::new(RefCell::new(Label::new("Count: 0")));
        
        // Create buttons with callbacks that update the shared state
        let mut increment_btn = Button::new("+");
        let mut decrement_btn = Button::new("-");
        
        // TODO: Set up callbacks
        // When increment is clicked: count += 1, update label
        // When decrement is clicked: count -= 1, update label
        
        // Hint: Clone Rc's into the closures
        
        Self { count, label, increment_btn, decrement_btn }
    }
}
```

### Expected Behavior

```rust
fn main() {
    let mut app = CounterApp::new();
    
    // Simulate clicks
    app.increment_btn.click();
    // Output: "Count: 1"
    
    app.increment_btn.click();
    // Output: "Count: 2"
    
    app.decrement_btn.click();
    // Output: "Count: 1"
}
```

---

## Theory Problems

### Problem T1: Ownership Analysis

Consider this code:

```rust
struct Widget {
    id: u32,
}

fn process(widgets: Vec<Widget>) -> Vec<Widget> {
    let first = widgets[0]; // Line A
    widgets // Line B
}
```

1. Why does Line A fail to compile?
2. How would you fix it to:
   a) Take ownership of the first element?
   b) Borrow the first element?
3. Draw the ownership flow diagram.

---

### Problem T2: Lifetime Puzzle

```rust
struct App<'a> {
    widgets: Vec<&'a Widget>,
}

impl<'a> App<'a> {
    fn add(&mut self, widget: &'a Widget) {
        self.widgets.push(widget);
    }
    
    fn get_first(&self) -> &Widget {
        self.widgets[0]
    }
}
```

1. What lifetime does `get_first` return?
2. Can this code compile? If not, what's the fix?
3. What happens if we try: `app.add(&Widget { id: 1 })`?

---

### Problem T3: Reference Counting Proof

**Prove that `Rc<T>` prevents use-after-free.**

Given:
- `Rc::new(v)` creates an Rc with count = 1
- `Rc::clone(&rc)` increments count
- Dropping an Rc decrements count
- Value is dropped when count reaches 0

**Prove:** If you hold an `Rc<T>`, the `T` is always valid.

*Hint: Use induction on the sequence of clone/drop operations.*

---

### Problem T4: Tree Memory Layout (Advanced)

For this widget tree:

```
Container
├── Label("Hello")
├── Container
│   ├── Button("OK")
│   └── Button("Cancel")
└── Label("Goodbye")
```

Using `Box<dyn Widget>`:
1. Draw the memory layout (stack vs heap)
2. Count the total number of heap allocations
3. Calculate the memory overhead of Box<dyn> vs direct storage

---

## Challenge: Mini UI Framework (Optional, 45 min)

Combine everything to build a minimal but working UI description:

```rust
fn main() {
    let app_state = Rc::new(RefCell::new(AppState::default()));
    
    let ui = vstack![
        label("Welcome to RustUI"),
        hstack![
            button("Increment", {
                let state = Rc::clone(&app_state);
                move || state.borrow_mut().count += 1
            }),
            button("Decrement", {
                let state = Rc::clone(&app_state);
                move || state.borrow_mut().count -= 1
            }),
        ],
        dynamic_label({
            let state = Rc::clone(&app_state);
            move || format!("Count: {}", state.borrow().count)
        }),
    ];
    
    println!("{:?}", ui);
}
```

Implement the macros and types to make this work!

---

## Submission Checklist

- [ ] Exercise 1: Widget trait with 3 implementations
- [ ] Exercise 2: Container with child management
- [ ] Exercise 3: Event handling with propagation
- [ ] Exercise 4: Callback-based counter
- [ ] Theory T1-T2: Written answers
- [ ] Theory T3: Proof sketch
- [ ] (Optional) Challenge completed

---

## Hints

<details>
<summary>Exercise 1 Hint</summary>

```rust
impl Widget for Label {
    fn min_size(&self) -> Size {
        Size { width: self.text.len() as f32 * 8.0, height: 20.0 }
    }
    // ...
}
```
</details>

<details>
<summary>Exercise 4 Hint</summary>

```rust
let count_clone = Rc::clone(&count);
let label_clone = Rc::clone(&label);
increment_btn.set_on_click(move || {
    *count_clone.borrow_mut() += 1;
    label_clone.borrow_mut().set_text(&format!("Count: {}", count_clone.borrow()));
});
```
</details>
