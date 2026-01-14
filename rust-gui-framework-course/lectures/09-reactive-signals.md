# Lecture 09: Reactive State Management & Signals

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lecture 08  
> **Goal:** Implement fine-grained reactivity with signals and effects

---

## Overview

Signals provide automatic, fine-grained reactivity. Instead of diffing entire views, we track exactly which parts of the UI depend on which data.

---

## Part 1: What Are Signals? (20 min)

### The Core Idea

```rust
// A signal is reactive state
let count = create_signal(0);

// Reading a signal tracks a dependency
let value = count.get();

// Writing a signal notifies dependents
count.set(5);

// Effects run when dependencies change
create_effect(|| {
    println!("Count is now: {}", count.get());
});
```

### Signal vs Elm Architecture

| Elm Architecture | Signals |
|------------------|---------|
| Coarse updates | Fine-grained updates |
| Diff entire view | Update only affected parts |
| Explicit messages | Automatic tracking |

---

## Part 2: Implementing Signals (30 min)

### Basic Signal

```rust
use std::cell::RefCell;
use std::rc::Rc;

type Subscriber = Rc<RefCell<dyn FnMut()>>;

pub struct Signal<T> {
    value: Rc<RefCell<T>>,
    subscribers: Rc<RefCell<Vec<Subscriber>>>,
}

impl<T: Clone> Signal<T> {
    pub fn new(value: T) -> Self {
        Self {
            value: Rc::new(RefCell::new(value)),
            subscribers: Rc::new(RefCell::new(Vec::new())),
        }
    }
    
    pub fn get(&self) -> T {
        // Track this read (simplified)
        CURRENT_EFFECT.with(|effect| {
            if let Some(subscriber) = effect.borrow().clone() {
                self.subscribers.borrow_mut().push(subscriber);
            }
        });
        self.value.borrow().clone()
    }
    
    pub fn set(&self, value: T) {
        *self.value.borrow_mut() = value;
        // Notify all subscribers
        for subscriber in self.subscribers.borrow().iter() {
            subscriber.borrow_mut()();
        }
    }
}

thread_local! {
    static CURRENT_EFFECT: RefCell<Option<Subscriber>> = RefCell::new(None);
}
```

### Creating Effects

```rust
pub fn create_effect<F: FnMut() + 'static>(mut f: F) {
    let subscriber: Subscriber = Rc::new(RefCell::new(move || f()));
    
    // Run the effect once to collect dependencies
    CURRENT_EFFECT.with(|current| {
        *current.borrow_mut() = Some(Rc::clone(&subscriber));
    });
    
    subscriber.borrow_mut()(); // Initial run
    
    CURRENT_EFFECT.with(|current| {
        *current.borrow_mut() = None;
    });
}
```

---

## Part 3: Computed Values (20 min)

### Derived State

```rust
pub struct Computed<T> {
    compute: Box<dyn Fn() -> T>,
    cached: RefCell<Option<T>>,
    dirty: Cell<bool>,
}

impl<T: Clone> Computed<T> {
    pub fn new<F: Fn() -> T + 'static>(compute: F) -> Self {
        Self {
            compute: Box::new(compute),
            cached: RefCell::new(None),
            dirty: Cell::new(true),
        }
    }
    
    pub fn get(&self) -> T {
        if self.dirty.get() {
            let value = (self.compute)();
            *self.cached.borrow_mut() = Some(value.clone());
            self.dirty.set(false);
            value
        } else {
            self.cached.borrow().clone().unwrap()
        }
    }
}
```

### Usage

```rust
let first_name = create_signal("John".to_string());
let last_name = create_signal("Doe".to_string());

let full_name = create_computed(|| {
    format!("{} {}", first_name.get(), last_name.get())
});

// full_name automatically updates when either signal changes
```

---

## Part 4: Reactive UI (25 min)

### Signal-Based View

```rust
fn counter_view(count: Signal<i32>) -> Element {
    Column::new()
        .child(
            // This text automatically updates when count changes
            reactive_text(|| format!("Count: {}", count.get()))
        )
        .child(
            Button::new("+").on_click(|| count.set(count.get() + 1))
        )
        .child(
            Button::new("-").on_click(|| count.set(count.get() - 1))
        )
}

fn reactive_text<F: Fn() -> String + 'static>(f: F) -> Element {
    let text_element = /* create text element */;
    
    create_effect(move || {
        let new_text = f();
        text_element.update_text(&new_text);
    });
    
    text_element
}
```

---

## Part 5: Ownership Challenges (15 min)

### The Rust Problem

```rust
// This doesn't work!
let count = Signal::new(0);
let button = Button::new("+").on_click(|| {
    count.set(count.get() + 1); // count moved into closure
});
let text = Text::reactive(|| count.get()); // can't use count again!
```

### Solutions

```rust
// 1. Signal is Clone (internally Rc)
impl<T> Clone for Signal<T> { /* clone the Rc */ }

let count = Signal::new(0);
let button = Button::new("+").on_click({
    let count = count.clone();
    move || count.set(count.get() + 1)
});
let text = Text::reactive({
    let count = count.clone();
    move || format!("{}", count.get())
});

// 2. Context-based (like React hooks)
fn view() -> Element {
    let count = use_signal(0); // Retrieved from context
    // ...
}
```

---

## Summary

| Concept | Purpose |
|---------|---------|
| Signal | Reactive primitive value |
| Effect | Side effect that tracks dependencies |
| Computed | Derived value, cached |
| Subscriber | Function called on change |

---

## Key Takeaways

1. **Automatic dependency tracking** via get() calls
2. **Fine-grained updates** - only affected parts re-render
3. **Computed values cache** derived state
4. **Clone signals** to share across closures
5. **Effects clean up** when scope ends

---

## Up Next

**Lecture 10**: Text rendering and custom graphics.
