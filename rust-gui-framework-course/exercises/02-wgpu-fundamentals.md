# Exercise Sheet 02: GPU Pipeline & wgpu Fundamentals

> **Estimated Time:** 1.5-2 hours  
> **Difficulty:** ⭐⭐ (Intermediate)

---

## Exercise 1: Basic wgpu Setup (25 min)

Create a complete wgpu application that displays a window with a colored background.

### 1.1 Project Setup

```bash
cargo new wgpu-basics
cd wgpu-basics
```

Add dependencies to `Cargo.toml`:
```toml
[dependencies]
wgpu = "0.19"
winit = "0.29"
pollster = "0.3"
env_logger = "0.10"
log = "0.4"
```

### 1.2 Requirements

Your application should:
1. Create an 800x600 window titled "wgpu Exercise"
2. Initialize wgpu with proper error handling
3. Clear the screen with a dark blue color `(0.1, 0.1, 0.3, 1.0)`
4. Handle window close properly
5. Print the GPU adapter name on startup

### Expected Output
```
Using adapter: AMD Radeon Pro 5500M (or similar)
```

---

## Exercise 2: Interactive Color Control (30 min)

Extend your application to respond to user input.

### 2.1 Keyboard Controls

| Key | Action |
|-----|--------|
| R | Set clear color to red |
| G | Set clear color to green |
| B | Set clear color to blue |
| Space | Cycle through colors |
| Escape | Close window |

### 2.2 Mouse Control

Make the clear color change based on mouse position:
- X position controls red channel (0.0 at left, 1.0 at right)
- Y position controls green channel (0.0 at top, 1.0 at bottom)
- Blue stays at 0.5

### 2.3 Required Code Structure

```rust
impl State {
    fn handle_input(&mut self, event: &WindowEvent) -> bool {
        // Return true if event was handled
        match event {
            // TODO: Handle keyboard
            // TODO: Handle mouse
            _ => false
        }
    }
}
```

---

## Exercise 3: Frame Timing and FPS Counter (25 min)

### 3.1 Implement Frame Timing

Create a `FrameTimer` struct that:
1. Calculates delta time between frames
2. Counts frames per second
3. Updates the window title with FPS

```rust
struct FrameTimer {
    last_frame: std::time::Instant,
    frame_count: u32,
    last_fps_update: std::time::Instant,
    current_fps: u32,
}

impl FrameTimer {
    fn new() -> Self { /* ... */ }
    fn tick(&mut self) -> std::time::Duration { /* ... */ }
    fn fps(&self) -> u32 { /* ... */ }
}
```

### 3.2 Animate the Background

Use delta time to animate the background color:
- Smoothly cycle through hue over 5 seconds
- Use HSV to RGB conversion

```rust
fn hsv_to_rgb(h: f32, s: f32, v: f32) -> (f32, f32, f32) {
    // h: 0-360, s: 0-1, v: 0-1
    // TODO: Implement
}
```

---

## Exercise 4: Window Resize Handling (20 min)

### 4.1 Proper Resize

Ensure your application:
1. Reconfigures the surface on resize
2. Maintains correct aspect ratio
3. Handles minimize (size 0,0) gracefully
4. Prints new dimensions on resize

### 4.2 Debug Grid

Draw a "grid" in the console showing the new dimensions:
```
Resized to 1200 x 800
┌────────────────────────────────────────┐
│                                        │
│            1200 x 800                  │
│                                        │
└────────────────────────────────────────┘
```

---

## Theory Problems

### Problem T1: GPU Pipeline Stages

For each stage of the rendering pipeline, describe:
1. What data it receives as input
2. What processing it performs
3. What data it outputs

| Stage | Input | Processing | Output |
|-------|-------|------------|--------|
| Vertex Shader | | | |
| Primitive Assembly | | | |
| Rasterizer | | | |
| Fragment Shader | | | |
| Blending | | | |

---

### Problem T2: wgpu Object Lifetimes

Consider this code:
```rust
fn create_render_pass<'a>(
    encoder: &'a mut wgpu::CommandEncoder,
    view: &'a wgpu::TextureView,
) -> wgpu::RenderPass<'a> {
    encoder.begin_render_pass(&wgpu::RenderPassDescriptor {
        color_attachments: &[Some(wgpu::RenderPassColorAttachment {
            view,
            resolve_target: None,
            ops: wgpu::Operations {
                load: wgpu::LoadOp::Clear(wgpu::Color::BLACK),
                store: wgpu::StoreOp::Store,
            },
        })],
        ..Default::default()
    })
}
```

1. Why does `RenderPass` have a lifetime parameter `'a`?
2. What would happen if we tried to use the encoder after creating the render pass?
3. Why must the render pass be dropped before calling `encoder.finish()`?

---

### Problem T3: Present Modes Analysis

For each present mode, analyze:

| Mode | Latency | Tearing | Power | Best For |
|------|---------|---------|-------|----------|
| Fifo | | | | |
| Immediate | | | | |
| Mailbox | | | | |
| FifoRelaxed | | | | |

Which would you choose for:
1. A text editor?
2. A fast-paced action game?
3. A battery-powered laptop?

---

### Problem T4: GPU Memory Architecture (Advanced)

Research and explain:
1. What is "GPU memory bandwidth" and why does it matter?
2. What is "texture cache" and how does it affect performance?
3. Why do we prefer indexed rendering over raw vertex lists?

Calculate: If you have 1920×1080 pixels, each pixel is 4 bytes (RGBA), and you're running at 60 FPS, what is the minimum memory bandwidth required just for the framebuffer?

---

## Challenge: Multi-Window Support (Optional, 30 min)

Extend your application to support multiple windows:

### Requirements
1. Press 'N' to create a new window
2. Each window has independent clear color
3. Each window can be resized independently
4. Closing one window doesn't close others
5. Application exits when all windows are closed

### Hints
- Use `HashMap<WindowId, WindowState>`
- Each window needs its own Surface
- Device and Queue can be shared

```rust
struct App {
    device: wgpu::Device,
    queue: wgpu::Queue,
    windows: HashMap<WindowId, WindowState>,
}

struct WindowState {
    window: Window,
    surface: wgpu::Surface,
    config: wgpu::SurfaceConfiguration,
    clear_color: wgpu::Color,
}
```

---

## Submission Checklist

- [ ] Exercise 1: Basic wgpu window with colored background
- [ ] Exercise 2: Keyboard and mouse color control
- [ ] Exercise 3: FPS counter with animated hue
- [ ] Exercise 4: Proper resize handling
- [ ] Theory T1: Pipeline stages table
- [ ] Theory T2: Lifetime analysis
- [ ] Theory T3: Present modes analysis
- [ ] (Optional) Challenge: Multi-window support

---

## Hints

<details>
<summary>Exercise 3: HSV to RGB</summary>

```rust
fn hsv_to_rgb(h: f32, s: f32, v: f32) -> (f32, f32, f32) {
    let h = h % 360.0;
    let c = v * s;
    let x = c * (1.0 - ((h / 60.0) % 2.0 - 1.0).abs());
    let m = v - c;
    
    let (r, g, b) = match h as i32 {
        0..=59 => (c, x, 0.0),
        60..=119 => (x, c, 0.0),
        120..=179 => (0.0, c, x),
        180..=239 => (0.0, x, c),
        240..=299 => (x, 0.0, c),
        _ => (c, 0.0, x),
    };
    
    (r + m, g + m, b + m)
}
```
</details>

<details>
<summary>Challenge: Multi-window Event Loop</summary>

```rust
event_loop.run(move |event, elwt| {
    match event {
        Event::WindowEvent { window_id, event } => {
            if let Some(state) = app.windows.get_mut(&window_id) {
                match event {
                    WindowEvent::CloseRequested => {
                        app.windows.remove(&window_id);
                        if app.windows.is_empty() {
                            elwt.exit();
                        }
                    }
                    // ... other events
                }
            }
        }
        // ...
    }
}).unwrap();
```
</details>
