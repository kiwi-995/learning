# Lecture 08: Declarative UI & The Elm Architecture

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lectures 01-07  
> **Goal:** Implement declarative UI with unidirectional data flow

---

## Overview

This is where GUI programming gets elegant. We'll implement the Elm Architecture pattern that powers React, Redux, and many modern frameworks.

---

## Part 1: The Elm Architecture (TEA) (25 min)

### Core Concepts

```
     ┌─────────────────────────────────────────┐
     │                                         │
     ▼                                         │
 ┌───────┐      ┌────────┐      ┌──────────┐  │
 │ Model │─────▶│  View  │─────▶│   UI     │  │
 └───────┘      └────────┘      └──────────┘  │
     ▲                               │        │
     │                               ▼        │
     │          ┌────────┐      ┌──────────┐  │
     └──────────│ Update │◀─────│ Message  │──┘
                └────────┘      └──────────┘
```

### The Three Components

```rust
/// Application state
struct Model {
    count: i32,
}

/// Events that can occur
enum Message {
    Increment,
    Decrement,
    Reset,
}

/// Pure function: old state + message → new state
fn update(model: &Model, msg: Message) -> Model {
    match msg {
        Message::Increment => Model { count: model.count + 1 },
        Message::Decrement => Model { count: model.count - 1 },
        Message::Reset => Model { count: 0 },
    }
}

/// Pure function: state → UI description
fn view(model: &Model) -> impl View {
    Column::new()
        .child(Text::new(format!("Count: {}", model.count)))
        .child(
            Row::new()
                .child(Button::new("-").on_click(Message::Decrement))
                .child(Button::new("+").on_click(Message::Increment))
        )
        .child(Button::new("Reset").on_click(Message::Reset))
}
```

---

## Part 2: Implementing the Runtime (30 min)

### The Application Trait

```rust
pub trait Application: Sized {
    type Message: Clone;
    
    fn new() -> Self;
    fn update(&mut self, message: Self::Message);
    fn view(&self) -> Element<Self::Message>;
}

pub fn run<A: Application + 'static>() {
    let mut app = A::new();
    
    event_loop.run(move |event, elwt| {
        match event {
            Event::WindowEvent { event: WindowEvent::RedrawRequested, .. } => {
                // 1. Build view from current state
                let view = app.view();
                
                // 2. Render the view
                view.render(&mut renderer);
            }
            Event::UserEvent(message) => {
                // 3. Update state with message
                app.update(message);
                
                // 4. Request redraw
                window.request_redraw();
            }
            // ... handle input events, convert to messages
        }
    });
}
```

### Element Tree

```rust
pub enum Element<M> {
    Text { content: String },
    Button { label: String, on_click: Option<M> },
    Column { children: Vec<Element<M>> },
    Row { children: Vec<Element<M>> },
}

impl<M: Clone> Element<M> {
    pub fn collect_messages(&self, input: &InputState) -> Vec<M> {
        let mut messages = Vec::new();
        self.collect_messages_recursive(input, &mut messages);
        messages
    }
}
```

---

## Part 3: View Diffing (25 min)

### Why Diff?

Instead of rebuilding the entire widget tree, detect changes:

```rust
enum ViewDiff {
    None,                    // No change
    Replace(Element),        // Completely different
    Update { /* fields */ }, // Partial update
    ChildDiffs(Vec<ViewDiff>),
}

fn diff<M>(old: &Element<M>, new: &Element<M>) -> ViewDiff {
    match (old, new) {
        (Element::Text { content: a }, Element::Text { content: b }) => {
            if a == b { ViewDiff::None } else { ViewDiff::Replace(new.clone()) }
        }
        (Element::Column { children: a }, Element::Column { children: b }) => {
            let child_diffs: Vec<_> = a.iter().zip(b.iter())
                .map(|(old, new)| diff(old, new))
                .collect();
            ViewDiff::ChildDiffs(child_diffs)
        }
        _ => ViewDiff::Replace(new.clone())
    }
}
```

---

## Part 4: Commands and Side Effects (20 min)

### Pure Updates + Commands

```rust
enum Command<M> {
    None,
    Batch(Vec<Command<M>>),
    Task(Box<dyn FnOnce() -> M>),
    Delay { duration: Duration, message: M },
}

fn update(&mut self, msg: Message) -> Command<Message> {
    match msg {
        Message::FetchData => {
            self.loading = true;
            Command::Task(Box::new(|| {
                let data = fetch_data(); // Async work
                Message::DataReceived(data)
            }))
        }
        Message::DataReceived(data) => {
            self.data = data;
            self.loading = false;
            Command::None
        }
    }
}
```

---

## Part 5: Complete Example (15 min)

### Todo App

```rust
struct TodoApp {
    todos: Vec<Todo>,
    input: String,
}

enum Message {
    InputChanged(String),
    AddTodo,
    ToggleTodo(usize),
    RemoveTodo(usize),
}

fn view(model: &TodoApp) -> Element<Message> {
    Column::new()
        .child(
            Row::new()
                .child(TextInput::new(&model.input)
                    .on_change(Message::InputChanged))
                .child(Button::new("Add").on_click(Message::AddTodo))
        )
        .children(
            model.todos.iter().enumerate().map(|(i, todo)| {
                Row::new()
                    .child(Checkbox::new(todo.completed)
                        .on_toggle(Message::ToggleTodo(i)))
                    .child(Text::new(&todo.text))
                    .child(Button::new("×")
                        .on_click(Message::RemoveTodo(i)))
            })
        )
}
```

---

## Summary

| Component | Purpose |
|-----------|---------|
| Model | Application state |
| Message | Describes what happened |
| Update | Pure: (model, msg) → model |
| View | Pure: model → UI |
| Command | Side effects |

---

## Key Takeaways

1. **Unidirectional data flow** simplifies reasoning
2. **Pure functions** for update and view
3. **Messages** are the only way to change state
4. **Diffing** enables efficient updates
5. **Commands** handle side effects cleanly

---

## Up Next

**Lecture 09**: Reactive state with signals for fine-grained updates.
