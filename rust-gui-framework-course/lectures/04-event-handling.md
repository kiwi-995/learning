# Lecture 04: Event Handling & Cross-Platform Windowing

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lectures 01-03  
> **Goal:** Build a robust event handling system for keyboard, mouse, and touch

---

## Overview

A GUI that can't respond to input isn't very useful! This lecture covers event handling, input state tracking, and building a cross-platform abstraction.

**By the end:**
- Understand the event loop deeply
- Handle keyboard, mouse, and touch input
- Build an input state tracker
- Implement high-DPI scaling

---

## Part 1: The Event Loop Revisited (20 min)

### winit's Event Flow

```rust
event_loop.run(move |event, elwt| {
    match event {
        Event::NewEvents(_) => {
            // Start of event batch
        }
        Event::WindowEvent { window_id, event } => {
            // Window-specific events
        }
        Event::DeviceEvent { device_id, event } => {
            // Raw device input (mouse motion, etc.)
        }
        Event::AboutToWait => {
            // All events processed, request redraw
            window.request_redraw();
        }
        Event::LoopExiting => {
            // Clean up
        }
        _ => {}
    }
});
```

### Window Events

```rust
match event {
    WindowEvent::CloseRequested => elwt.exit(),
    
    WindowEvent::Resized(size) => {
        // Handle resize
    }
    
    WindowEvent::Focused(focused) => {
        // Window gained/lost focus
    }
    
    WindowEvent::KeyboardInput { event, .. } => {
        // Keyboard
    }
    
    WindowEvent::MouseInput { button, state, .. } => {
        // Mouse button
    }
    
    WindowEvent::CursorMoved { position, .. } => {
        // Mouse position
    }
    
    WindowEvent::MouseWheel { delta, .. } => {
        // Scroll
    }
    
    WindowEvent::Touch(touch) => {
        // Touch input
    }
    
    WindowEvent::RedrawRequested => {
        // Render!
    }
    
    _ => {}
}
```

---

## Part 2: Keyboard Input (20 min)

### Physical vs Logical Keys

```rust
WindowEvent::KeyboardInput { event, .. } => {
    // Physical key (hardware position)
    let physical = event.physical_key;
    // Example: PhysicalKey::Code(KeyCode::KeyA)
    
    // Logical key (what it types, respects layout)
    let logical = &event.logical_key;
    // Example: Key::Character("a") or Key::Named(NamedKey::Enter)
    
    // Event state
    let pressed = event.state == ElementState::Pressed;
    
    // Key repeat?
    let repeat = event.repeat;
}
```

### Building a Keyboard State Tracker

```rust
use std::collections::HashSet;
use winit::keyboard::{KeyCode, PhysicalKey};

pub struct KeyboardState {
    pressed: HashSet<KeyCode>,
    just_pressed: HashSet<KeyCode>,
    just_released: HashSet<KeyCode>,
}

impl KeyboardState {
    pub fn new() -> Self {
        Self {
            pressed: HashSet::new(),
            just_pressed: HashSet::new(),
            just_released: HashSet::new(),
        }
    }
    
    pub fn process_event(&mut self, event: &KeyEvent) {
        if let PhysicalKey::Code(code) = event.physical_key {
            match event.state {
                ElementState::Pressed if !event.repeat => {
                    self.pressed.insert(code);
                    self.just_pressed.insert(code);
                }
                ElementState::Released => {
                    self.pressed.remove(&code);
                    self.just_released.insert(code);
                }
                _ => {}
            }
        }
    }
    
    pub fn is_pressed(&self, key: KeyCode) -> bool {
        self.pressed.contains(&key)
    }
    
    pub fn just_pressed(&self, key: KeyCode) -> bool {
        self.just_pressed.contains(&key)
    }
    
    pub fn clear_frame(&mut self) {
        self.just_pressed.clear();
        self.just_released.clear();
    }
}
```

---

## Part 3: Mouse Input (20 min)

### Mouse State Tracker

```rust
use winit::event::MouseButton;

pub struct MouseState {
    position: (f32, f32),
    buttons: HashSet<MouseButton>,
    just_pressed: HashSet<MouseButton>,
    just_released: HashSet<MouseButton>,
    scroll_delta: (f32, f32),
}

impl MouseState {
    pub fn process_event(&mut self, event: &WindowEvent) {
        match event {
            WindowEvent::CursorMoved { position, .. } => {
                self.position = (position.x as f32, position.y as f32);
            }
            WindowEvent::MouseInput { button, state, .. } => {
                match state {
                    ElementState::Pressed => {
                        self.buttons.insert(*button);
                        self.just_pressed.insert(*button);
                    }
                    ElementState::Released => {
                        self.buttons.remove(button);
                        self.just_released.insert(*button);
                    }
                }
            }
            WindowEvent::MouseWheel { delta, .. } => {
                match delta {
                    MouseScrollDelta::LineDelta(x, y) => {
                        self.scroll_delta = (*x, *y);
                    }
                    MouseScrollDelta::PixelDelta(pos) => {
                        self.scroll_delta = (pos.x as f32, pos.y as f32);
                    }
                }
            }
            _ => {}
        }
    }
    
    pub fn is_in_rect(&self, x: f32, y: f32, w: f32, h: f32) -> bool {
        self.position.0 >= x && self.position.0 < x + w &&
        self.position.1 >= y && self.position.1 < y + h
    }
}
```

---

## Part 4: High-DPI Handling (15 min)

### Understanding Scale Factor

```rust
// Physical pixels vs logical pixels
let scale_factor = window.scale_factor();  // e.g., 2.0 on Retina

// Physical size (actual pixels)
let physical: PhysicalSize<u32> = window.inner_size();

// Logical size (what the OS reports)
let logical: LogicalSize<f32> = physical.to_logical(scale_factor);

// Convert mouse position
fn to_physical(logical_pos: (f32, f32), scale: f64) -> (f32, f32) {
    (logical_pos.0 * scale as f32, logical_pos.1 * scale as f32)
}
```

### Handling Scale Changes

```rust
WindowEvent::ScaleFactorChanged { scale_factor, inner_size_writer } => {
    // Update rendering
    // Adjust UI element sizes
}
```

---

## Part 5: Touch Input (15 min)

### Touch Events

```rust
WindowEvent::Touch(touch) => {
    match touch.phase {
        TouchPhase::Started => {
            // Finger down
            let id = touch.id;
            let pos = touch.location;
        }
        TouchPhase::Moved => {
            // Finger moved
        }
        TouchPhase::Ended | TouchPhase::Cancelled => {
            // Finger up
        }
    }
}
```

### Multi-Touch Tracker

```rust
pub struct TouchState {
    touches: HashMap<u64, TouchPoint>,
}

pub struct TouchPoint {
    id: u64,
    start_position: (f32, f32),
    current_position: (f32, f32),
    phase: TouchPhase,
}
```

---

## Part 6: Unified Input System (15 min)

### Combined Input State

```rust
pub struct InputState {
    pub keyboard: KeyboardState,
    pub mouse: MouseState,
    pub touch: TouchState,
}

impl InputState {
    pub fn process_event(&mut self, event: &WindowEvent) {
        match event {
            WindowEvent::KeyboardInput { event, .. } => {
                self.keyboard.process_event(event);
            }
            e @ (WindowEvent::CursorMoved { .. }
                | WindowEvent::MouseInput { .. }
                | WindowEvent::MouseWheel { .. }) => {
                self.mouse.process_event(e);
            }
            WindowEvent::Touch(touch) => {
                self.touch.process_event(touch);
            }
            _ => {}
        }
    }
    
    pub fn end_frame(&mut self) {
        self.keyboard.clear_frame();
        self.mouse.clear_frame();
    }
}
```

---

## Summary

| Input Type | Key Methods |
|------------|-------------|
| Keyboard | `is_pressed()`, `just_pressed()`, physical vs logical keys |
| Mouse | position, buttons, scroll delta, hit testing |
| Touch | multi-touch tracking, gesture detection |
| Scale | physical vs logical pixels, DPI handling |

---

## Key Takeaways

1. **Physical keys for shortcuts** - KeyCode::KeyW is always W key position
2. **Logical keys for text** - Respects keyboard layout
3. **Track "just pressed"** - Essential for UI interactions
4. **Clear per-frame state** - `just_pressed` resets each frame
5. **Always handle scale factor** - Required for high-DPI displays

---

## Up Next

In **Lecture 05**, we'll design the widget architecture—immediate mode vs retained mode, widget traits, and composition patterns.
