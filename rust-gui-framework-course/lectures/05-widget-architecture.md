# Lecture 05: Widget Architecture & Composition

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lectures 01-04  
> **Goal:** Understand GUI paradigms and build a widget system

---

## Overview

Now we design how widgets are structured. The architecture choices here will affect everything else.

**Key topics:**
- Immediate Mode vs Retained Mode GUI
- Widget lifecycle
- Composition patterns  
- Building core widgets

---

## Part 1: Immediate Mode GUI (25 min)

### The Immediate Mode Philosophy

Every frame, you "describe" the UI:

```rust
fn ui_frame(state: &mut AppState, ctx: &mut UiContext) {
    if ctx.button("Increment") {
        state.counter += 1;
    }
    ctx.label(&format!("Count: {}", state.counter));
}
```

**Characteristics:**
- No widget objects stored between frames
- UI is rebuilt every frame
- State lives in application, not UI

### Pros and Cons

| Pros | Cons |
|------|------|
| Simple mental model | CPU work every frame |
| No state sync issues | Layout computed every frame |
| Easy dynamic UI | Animation is tricky |
| Great for tools | Hard to style independently |

### Example: dear imgui Style

```rust
pub struct ImmediateContext {
    mouse_pos: (f32, f32),
    mouse_down: bool,
    hot_id: Option<usize>,     // Widget under cursor
    active_id: Option<usize>,  // Widget being interacted
}

impl ImmediateContext {
    pub fn button(&mut self, id: usize, label: &str, x: f32, y: f32) -> bool {
        let rect = Rect::new(x, y, 100.0, 30.0);
        let hovered = rect.contains(self.mouse_pos);
        
        if hovered {
            self.hot_id = Some(id);
        }
        
        let clicked = self.active_id == Some(id) && !self.mouse_down;
        if hovered && self.mouse_down {
            self.active_id = Some(id);
        }
        
        // Draw button with appropriate state
        let color = if self.active_id == Some(id) {
            Color::DARK_BLUE
        } else if hovered {
            Color::LIGHT_BLUE
        } else {
            Color::BLUE
        };
        
        self.draw_rect(rect, color);
        self.draw_text(label, rect.center());
        
        clicked
    }
}
```

---

## Part 2: Retained Mode GUI (25 min)

### The Retained Mode Approach

Widgets exist as objects in a tree:

```rust
struct Button {
    label: String,
    bounds: Rect,
    on_click: Option<Box<dyn Fn()>>,
}

struct Container {
    children: Vec<Box<dyn Widget>>,
}

let mut ui = Container::new();
ui.add(Button::new("Click me").on_click(|| println!("Clicked!")));
```

**Characteristics:**
- Widgets persist between frames
- Explicit update → layout → render cycle
- State in widgets OR application

### Pros and Cons

| Pros | Cons |
|------|------|
| Efficient (only update changed) | State sync complexity |
| Rich animations | More boilerplate |
| Designer-friendly | Widget lifecycle management |
| Standard pattern | "Stale widget" bugs |

---

## Part 3: The Widget Trait (20 min)

### Core Widget Interface

```rust
pub trait Widget {
    /// Unique identifier
    fn id(&self) -> WidgetId;
    
    /// Preferred, minimum, maximum sizes
    fn measure(&self, ctx: &MeasureContext) -> SizeConstraints;
    
    /// Position and size children
    fn layout(&mut self, bounds: Rect, ctx: &mut LayoutContext);
    
    /// Handle input event
    fn event(&mut self, event: &Event, ctx: &mut EventContext) -> EventResult;
    
    /// Draw the widget
    fn render(&self, ctx: &mut RenderContext);
    
    /// Children (for containers)
    fn children(&self) -> &[WidgetRef] { &[] }
    fn children_mut(&mut self) -> &mut [WidgetRef] { &mut [] }
}
```

### Widget Lifecycle

```
┌─────────┐
│ Create  │
└────┬────┘
     ▼
┌─────────┐     Input
│ Measure │ ◄───────┐
└────┬────┘         │
     ▼              │
┌─────────┐         │
│ Layout  │         │
└────┬────┘         │
     ▼              │
┌─────────┐         │
│  Event  │─────────┘
└────┬────┘
     ▼
┌─────────┐
│ Render  │
└────┬────┘
     ▼
┌─────────┐
│ Destroy │
└─────────┘
```

---

## Part 4: Core Widgets (25 min)

### Label

```rust
pub struct Label {
    id: WidgetId,
    text: String,
    font_size: f32,
    color: Color,
}

impl Widget for Label {
    fn measure(&self, _ctx: &MeasureContext) -> SizeConstraints {
        let width = self.text.len() as f32 * self.font_size * 0.6;
        let height = self.font_size * 1.5;
        SizeConstraints::fixed(width, height)
    }
    
    fn render(&self, ctx: &mut RenderContext) {
        ctx.draw_text(&self.text, ctx.bounds.origin(), self.font_size, self.color);
    }
}
```

### Button

```rust
pub struct Button {
    id: WidgetId,
    label: String,
    is_hovered: bool,
    is_pressed: bool,
    on_click: Option<Box<dyn Fn()>>,
}

impl Widget for Button {
    fn event(&mut self, event: &Event, ctx: &mut EventContext) -> EventResult {
        match event {
            Event::MouseMove(pos) => {
                self.is_hovered = ctx.bounds.contains(*pos);
                EventResult::Ignored
            }
            Event::MouseDown(pos) if ctx.bounds.contains(*pos) => {
                self.is_pressed = true;
                EventResult::Captured
            }
            Event::MouseUp(_) if self.is_pressed => {
                self.is_pressed = false;
                if let Some(ref callback) = self.on_click {
                    callback();
                }
                EventResult::Captured
            }
            _ => EventResult::Ignored
        }
    }
    
    fn render(&self, ctx: &mut RenderContext) {
        let color = match (self.is_pressed, self.is_hovered) {
            (true, _) => Color::DARK_BLUE,
            (false, true) => Color::LIGHT_BLUE,
            (false, false) => Color::BLUE,
        };
        ctx.draw_rect(ctx.bounds, color);
        ctx.draw_text_centered(&self.label, ctx.bounds);
    }
}
```

### Container

```rust
pub struct VStack {
    id: WidgetId,
    children: Vec<Box<dyn Widget>>,
    spacing: f32,
}

impl Widget for VStack {
    fn measure(&self, ctx: &MeasureContext) -> SizeConstraints {
        let mut total_height = 0.0;
        let mut max_width = 0.0;
        
        for child in &self.children {
            let child_size = child.measure(ctx);
            total_height += child_size.preferred.height + self.spacing;
            max_width = max_width.max(child_size.preferred.width);
        }
        
        SizeConstraints::preferred(max_width, total_height)
    }
    
    fn layout(&mut self, bounds: Rect, ctx: &mut LayoutContext) {
        let mut y = bounds.y;
        for child in &mut self.children {
            let child_size = child.measure(&ctx.measure);
            let child_bounds = Rect::new(bounds.x, y, bounds.width, child_size.preferred.height);
            child.layout(child_bounds, ctx);
            y += child_size.preferred.height + self.spacing;
        }
    }
}
```

---

## Part 5: "UI as Data" Philosophy (15 min)

### Declarative UI Description

Instead of imperatively building UI:
```rust
// Imperative
let mut container = VStack::new();
container.add(Label::new("Hello"));
container.add(Button::new("Click"));
```

...describe what the UI should look like:
```rust
// Declarative
fn view(state: &State) -> impl Widget {
    VStack::new()
        .child(Label::new(&state.message))
        .child(Button::new("Click").on_click(|s| s.clicked = true))
}
```

### Benefits
- UI directly reflects state
- Easy to reason about
- Enables diffing and efficient updates
- Foundation for reactive patterns

---

## Summary

| Approach | Best For |
|----------|----------|
| Immediate Mode | Tools, debug UIs, simple apps |
| Retained Mode | Full applications, rich UIs |
| Declarative | Modern patterns (we'll use this) |

---

## Key Takeaways

1. **Immediate mode is simpler** but less efficient
2. **Retained mode is standard** for production UIs
3. **Widget trait is the core abstraction**
4. **Lifecycle: measure → layout → event → render**
5. **"UI as data"** enables powerful patterns

---

## Up Next

In **Lecture 06**, we dive deep into layout—how widgets determine their sizes and positions.
