# Lecture 02: The GPU Pipeline & wgpu Fundamentals

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lecture 01 (Rust Foundations)  
> **Goal:** Understand GPU rendering and create your first wgpu application

---

## Overview

Modern GUIs are GPU-accelerated. Before we can render our first widget, we need to understand how graphics hardware works and how wgpu abstracts it across platforms.

**By the end of this lecture:**
- Understand the GPU rendering pipeline
- Set up a wgpu project with windowing
- Render a solid color to the screen
- Handle the event loop and frame timing

---

## Part 1: How GPUs Work (25 min)

### CPU vs GPU Architecture

```
CPU: Few powerful cores, optimized for sequential tasks
┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐
│Core1│ │Core2│ │Core3│ │Core4│   (4-16 cores)
└─────┘ └─────┘ └─────┘ └─────┘

GPU: Many simple cores, optimized for parallel tasks
┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐...  (1000s of cores)
└─┘└─┘└─┘└─┘└─┘└─┘└─┘└─┘└─┘└─┘└─┘└─┘
```

**Why GPUs for UI?**
- Drawing millions of pixels is embarrassingly parallel
- Each pixel can be computed independently
- 60 FPS = 16.6ms per frame (GPU handles this easily)

### The Graphics Pipeline

```
┌──────────────┐
│ Vertex Data  │  Points that define shapes
└──────┬───────┘
       ▼
┌──────────────┐
│Vertex Shader │  Transform positions (GPU program)
└──────┬───────┘
       ▼
┌──────────────┐
│  Rasterizer  │  Convert triangles to pixels
└──────┬───────┘
       ▼
┌──────────────┐
│Fragment Shader│ Calculate pixel colors (GPU program)
└──────┬───────┘
       ▼
┌──────────────┐
│ Framebuffer  │  Final image to display
└──────────────┘
```

### Key Concepts

**Vertices:** Points in space with attributes (position, color, texture coords)
```rust
struct Vertex {
    position: [f32; 2],  // x, y
    color: [f32; 3],     // r, g, b
}
```

**Shaders:** Programs that run on the GPU
- Vertex shader: Runs once per vertex
- Fragment shader: Runs once per pixel

**Buffers:** Data storage on the GPU (vertex buffers, index buffers, uniform buffers)

**Render Pipeline:** Configuration of how to draw (shaders, blending, depth testing)

---

## Part 2: wgpu Architecture (20 min)

### What is wgpu?

wgpu is a cross-platform graphics API that abstracts:
- **Vulkan** (Linux, Windows, Android)
- **Metal** (macOS, iOS)  
- **DirectX 12** (Windows)
- **WebGPU** (Browsers via WASM)

```rust
// One codebase, all platforms!
```

### The wgpu Object Hierarchy

```
Instance
    └── Adapter (represents a GPU)
          └── Device (logical connection to GPU)
                ├── Queue (command submission)
                ├── Buffer (GPU memory)
                ├── Texture (images)
                └── RenderPipeline (drawing config)
```

### Key wgpu Types

```rust
// Instance: Entry point to wgpu
let instance = wgpu::Instance::default();

// Adapter: A handle to a physical GPU
let adapter = instance.request_adapter(&options).await;

// Device + Queue: The main interface for GPU work
let (device, queue) = adapter.request_device(&desc, None).await;

// Surface: Where we render to (usually a window)
let surface = instance.create_surface(&window)?;

// RenderPipeline: How to render (shaders, blending, etc.)
let pipeline = device.create_render_pipeline(&desc);
```

---

## Part 3: Setting Up Your First wgpu Project (30 min)

### Project Setup

```bash
cargo new gpu-hello
cd gpu-hello
```

**Cargo.toml:**
```toml
[package]
name = "gpu-hello"
version = "0.1.0"
edition = "2021"

[dependencies]
wgpu = "0.19"
winit = "0.29"
pollster = "0.3"  # For blocking on async
env_logger = "0.10"
log = "0.4"
```

### The Application Structure

```rust
use winit::{
    event::{Event, WindowEvent},
    event_loop::EventLoop,
    window::WindowBuilder,
};

fn main() {
    env_logger::init();
    
    // 1. Create event loop and window
    let event_loop = EventLoop::new().unwrap();
    let window = WindowBuilder::new()
        .with_title("Hello wgpu!")
        .with_inner_size(winit::dpi::LogicalSize::new(800, 600))
        .build(&event_loop)
        .unwrap();
    
    // 2. Initialize wgpu (async, but we block with pollster)
    let mut state = pollster::block_on(State::new(&window));
    
    // 3. Run the event loop
    event_loop.run(move |event, elwt| {
        match event {
            Event::WindowEvent { event, .. } => match event {
                WindowEvent::CloseRequested => elwt.exit(),
                WindowEvent::Resized(size) => state.resize(size),
                WindowEvent::RedrawRequested => {
                    state.update();
                    state.render().unwrap();
                }
                _ => {}
            }
            Event::AboutToWait => {
                window.request_redraw();
            }
            _ => {}
        }
    }).unwrap();
}
```

### The State Struct

```rust
struct State<'a> {
    surface: wgpu::Surface<'a>,
    device: wgpu::Device,
    queue: wgpu::Queue,
    config: wgpu::SurfaceConfiguration,
    size: winit::dpi::PhysicalSize<u32>,
    clear_color: wgpu::Color,
}

impl<'a> State<'a> {
    async fn new(window: &'a winit::window::Window) -> Self {
        let size = window.inner_size();
        
        // Create instance
        let instance = wgpu::Instance::new(wgpu::InstanceDescriptor {
            backends: wgpu::Backends::all(),
            ..Default::default()
        });
        
        // Create surface
        let surface = instance.create_surface(window).unwrap();
        
        // Request adapter
        let adapter = instance
            .request_adapter(&wgpu::RequestAdapterOptions {
                power_preference: wgpu::PowerPreference::default(),
                compatible_surface: Some(&surface),
                force_fallback_adapter: false,
            })
            .await
            .unwrap();
        
        // Request device and queue
        let (device, queue) = adapter
            .request_device(
                &wgpu::DeviceDescriptor {
                    label: None,
                    required_features: wgpu::Features::empty(),
                    required_limits: wgpu::Limits::default(),
                },
                None,
            )
            .await
            .unwrap();
        
        // Configure surface
        let surface_caps = surface.get_capabilities(&adapter);
        let surface_format = surface_caps.formats.iter()
            .find(|f| f.is_srgb())
            .copied()
            .unwrap_or(surface_caps.formats[0]);
            
        let config = wgpu::SurfaceConfiguration {
            usage: wgpu::TextureUsages::RENDER_ATTACHMENT,
            format: surface_format,
            width: size.width,
            height: size.height,
            present_mode: wgpu::PresentMode::Fifo,
            alpha_mode: surface_caps.alpha_modes[0],
            view_formats: vec![],
            desired_maximum_frame_latency: 2,
        };
        surface.configure(&device, &config);
        
        Self {
            surface,
            device,
            queue,
            config,
            size,
            clear_color: wgpu::Color {
                r: 0.1,
                g: 0.2,
                b: 0.3,
                a: 1.0,
            },
        }
    }
    
    fn resize(&mut self, new_size: winit::dpi::PhysicalSize<u32>) {
        if new_size.width > 0 && new_size.height > 0 {
            self.size = new_size;
            self.config.width = new_size.width;
            self.config.height = new_size.height;
            self.surface.configure(&self.device, &self.config);
        }
    }
    
    fn update(&mut self) {
        // Animation goes here
    }
    
    fn render(&mut self) -> Result<(), wgpu::SurfaceError> {
        // Get the current texture to render to
        let output = self.surface.get_current_texture()?;
        let view = output.texture.create_view(&Default::default());
        
        // Create command encoder
        let mut encoder = self.device.create_command_encoder(
            &wgpu::CommandEncoderDescriptor { label: Some("Render Encoder") }
        );
        
        // Begin render pass
        {
            let _render_pass = encoder.begin_render_pass(&wgpu::RenderPassDescriptor {
                label: Some("Render Pass"),
                color_attachments: &[Some(wgpu::RenderPassColorAttachment {
                    view: &view,
                    resolve_target: None,
                    ops: wgpu::Operations {
                        load: wgpu::LoadOp::Clear(self.clear_color),
                        store: wgpu::StoreOp::Store,
                    },
                })],
                depth_stencil_attachment: None,
                ..Default::default()
            });
        }
        
        // Submit commands
        self.queue.submit(std::iter::once(encoder.finish()));
        output.present();
        
        Ok(())
    }
}
```

---

## Part 4: Understanding the Render Loop (15 min)

### The Event Loop Model

```
┌─────────────────────────────────────────┐
│              Event Loop                  │
│  ┌─────────────────────────────────────┐│
│  │ Receive Events (input, window, etc) ││
│  └────────────────┬────────────────────┘│
│                   ▼                      │
│  ┌─────────────────────────────────────┐│
│  │         Update State                 ││
│  └────────────────┬────────────────────┘│
│                   ▼                      │
│  ┌─────────────────────────────────────┐│
│  │          Render Frame               ││
│  └────────────────┬────────────────────┘│
│                   ▼                      │
│  ┌─────────────────────────────────────┐│
│  │         Present to Screen           ││
│  └─────────────────────────────────────┘│
└─────────────────────────────────────────┘
```

### Frame Timing

```rust
use std::time::{Duration, Instant};

struct FrameTimer {
    last_frame: Instant,
    frame_count: u32,
    fps_timer: Instant,
}

impl FrameTimer {
    fn new() -> Self {
        Self {
            last_frame: Instant::now(),
            frame_count: 0,
            fps_timer: Instant::now(),
        }
    }
    
    fn delta(&mut self) -> Duration {
        let now = Instant::now();
        let delta = now - self.last_frame;
        self.last_frame = now;
        
        self.frame_count += 1;
        if self.fps_timer.elapsed() >= Duration::from_secs(1) {
            println!("FPS: {}", self.frame_count);
            self.frame_count = 0;
            self.fps_timer = now;
        }
        
        delta
    }
}
```

### VSync and Present Modes

```rust
// PresentMode options:
wgpu::PresentMode::Fifo       // VSync on, guaranteed no tearing
wgpu::PresentMode::Immediate  // VSync off, lowest latency, may tear
wgpu::PresentMode::Mailbox    // Low latency, no tearing, higher GPU usage
```

---

## Part 5: Interactive Demo (15 min)

### Changing Color on Input

```rust
fn handle_input(&mut self, event: &WindowEvent) {
    match event {
        WindowEvent::KeyboardInput { event, .. } => {
            if event.state == winit::event::ElementState::Pressed {
                match event.physical_key {
                    winit::keyboard::PhysicalKey::Code(winit::keyboard::KeyCode::KeyR) => {
                        self.clear_color = wgpu::Color::RED;
                    }
                    winit::keyboard::PhysicalKey::Code(winit::keyboard::KeyCode::KeyG) => {
                        self.clear_color = wgpu::Color::GREEN;
                    }
                    winit::keyboard::PhysicalKey::Code(winit::keyboard::KeyCode::KeyB) => {
                        self.clear_color = wgpu::Color::BLUE;
                    }
                    _ => {}
                }
            }
        }
        WindowEvent::CursorMoved { position, .. } => {
            // Color based on mouse position
            self.clear_color = wgpu::Color {
                r: position.x / self.size.width as f64,
                g: position.y / self.size.height as f64,
                b: 0.5,
                a: 1.0,
            };
        }
        _ => {}
    }
}
```

---

## Summary

| Concept | What It Does |
|---------|-------------|
| Instance | Entry point to wgpu |
| Adapter | Represents a physical GPU |
| Device | Logical connection for creating resources |
| Queue | Submits commands to the GPU |
| Surface | Target for rendering (usually a window) |
| RenderPass | A set of draw commands |
| CommandEncoder | Records commands for the GPU |

---

## Key Takeaways

1. **GPUs are massively parallel** - perfect for rendering
2. **wgpu abstracts all major graphics APIs** - write once, run everywhere
3. **The render loop**: gather input → update state → render → present
4. **Surface configuration** is essential for proper display
5. **VSync (Fifo)** is usually what you want for UI

---

## Up Next

In **Lecture 03**, we'll draw actual shapes! We'll learn about vertices, shaders, and how to render 2D primitives like rectangles and circles.

---

## References

- [Learn Wgpu](https://sotrh.github.io/learn-wgpu/) - The primary wgpu tutorial
- [wgpu Documentation](https://docs.rs/wgpu)
- [winit Documentation](https://docs.rs/winit)
- [WebGPU Specification](https://www.w3.org/TR/webgpu/)
