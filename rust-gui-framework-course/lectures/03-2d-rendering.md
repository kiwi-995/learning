# Lecture 03: 2D Rendering & Drawing Primitives

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lecture 02 (wgpu Fundamentals)  
> **Goal:** Draw shapes using vertices, shaders, and render pipelines

---

## Overview

Now that we can clear the screen, let's draw something! This lecture covers the fundamentals of 2D rendering: vertices, shaders, and how to turn geometry into pixels.

**By the end of this lecture:**
- Understand coordinate systems and transformations
- Write vertex and fragment shaders in WGSL
- Render triangles, rectangles, and circles
- Build a reusable `draw_rect()` function

---

## Part 1: Coordinate Systems (15 min)

### Normalized Device Coordinates (NDC)

wgpu uses NDC where:
- X: -1 (left) to +1 (right)
- Y: -1 (bottom) to +1 (top)  
- Origin: center of screen

```
        (+1, +1)
    ┌─────────────┐
    │      │      │
    │------+------│ (0, 0) = center
    │      │      │
    └─────────────┘
(-1, -1)
```

### Screen Space Coordinates

We usually think in pixels:
- X: 0 (left) to width (right)
- Y: 0 (top) to height (bottom)

### Conversion

```rust
fn screen_to_ndc(x: f32, y: f32, width: f32, height: f32) -> [f32; 2] {
    [
        (x / width) * 2.0 - 1.0,           // X: [0, width] → [-1, 1]
        1.0 - (y / height) * 2.0,          // Y: [0, height] → [1, -1] (flip!)
    ]
}

// Example: center of 800x600 window
let ndc = screen_to_ndc(400.0, 300.0, 800.0, 600.0);
// Result: [0.0, 0.0]
```

---

## Part 2: Vertices and Buffers (20 min)

### Defining Vertex Data

```rust
#[repr(C)]
#[derive(Copy, Clone, Debug, bytemuck::Pod, bytemuck::Zeroable)]
struct Vertex {
    position: [f32; 2],
    color: [f32; 3],
}

impl Vertex {
    // Describe how vertex data is laid out in memory
    fn desc() -> wgpu::VertexBufferLayout<'static> {
        wgpu::VertexBufferLayout {
            array_stride: std::mem::size_of::<Vertex>() as wgpu::BufferAddress,
            step_mode: wgpu::VertexStepMode::Vertex,
            attributes: &[
                wgpu::VertexAttribute {
                    offset: 0,
                    shader_location: 0,
                    format: wgpu::VertexFormat::Float32x2,
                },
                wgpu::VertexAttribute {
                    offset: std::mem::size_of::<[f32; 2]>() as wgpu::BufferAddress,
                    shader_location: 1,
                    format: wgpu::VertexFormat::Float32x3,
                },
            ],
        }
    }
}
```

### Creating a Vertex Buffer

```rust
// A colorful triangle
const TRIANGLE_VERTICES: &[Vertex] = &[
    Vertex { position: [0.0, 0.5], color: [1.0, 0.0, 0.0] },    // Top (red)
    Vertex { position: [-0.5, -0.5], color: [0.0, 1.0, 0.0] },  // Bottom-left (green)
    Vertex { position: [0.5, -0.5], color: [0.0, 0.0, 1.0] },   // Bottom-right (blue)
];

let vertex_buffer = device.create_buffer_init(&wgpu::util::BufferInitDescriptor {
    label: Some("Vertex Buffer"),
    contents: bytemuck::cast_slice(TRIANGLE_VERTICES),
    usage: wgpu::BufferUsages::VERTEX,
});
```

Add to dependencies:
```toml
bytemuck = { version = "1.14", features = ["derive"] }
wgpu = { version = "0.19", features = ["util"] }
```

---

## Part 3: Writing Shaders in WGSL (25 min)

### Introduction to WGSL

WGSL (WebGPU Shading Language) is the shader language for wgpu.

**shader.wgsl:**
```wgsl
// Vertex shader input
struct VertexInput {
    @location(0) position: vec2<f32>,
    @location(1) color: vec3<f32>,
};

// Vertex shader output / Fragment shader input
struct VertexOutput {
    @builtin(position) clip_position: vec4<f32>,
    @location(0) color: vec3<f32>,
};

// Vertex shader: runs once per vertex
@vertex
fn vs_main(in: VertexInput) -> VertexOutput {
    var out: VertexOutput;
    out.clip_position = vec4<f32>(in.position, 0.0, 1.0);
    out.color = in.color;
    return out;
}

// Fragment shader: runs once per pixel
@fragment
fn fs_main(in: VertexOutput) -> @location(0) vec4<f32> {
    return vec4<f32>(in.color, 1.0);
}
```

### Loading Shaders in Rust

```rust
let shader = device.create_shader_module(wgpu::ShaderModuleDescriptor {
    label: Some("Shader"),
    source: wgpu::ShaderSource::Wgsl(include_str!("shader.wgsl").into()),
});
```

### WGSL Basics

```wgsl
// Types
let x: f32 = 1.0;           // Float
let v: vec2<f32> = vec2(1.0, 2.0);  // 2D vector
let c: vec4<f32> = vec4(1.0, 0.0, 0.0, 1.0);  // RGBA color

// Operations
let sum = v.x + v.y;
let scaled = v * 2.0;
let rgb = c.rgb;  // Swizzling: extract r, g, b

// Built-in functions
let len = length(v);
let norm = normalize(v);
let d = distance(v, vec2(0.0, 0.0));
```

---

## Part 4: Render Pipelines (20 min)

### Creating a Render Pipeline

```rust
let render_pipeline_layout = device.create_pipeline_layout(&wgpu::PipelineLayoutDescriptor {
    label: Some("Render Pipeline Layout"),
    bind_group_layouts: &[],
    push_constant_ranges: &[],
});

let render_pipeline = device.create_render_pipeline(&wgpu::RenderPipelineDescriptor {
    label: Some("Render Pipeline"),
    layout: Some(&render_pipeline_layout),
    vertex: wgpu::VertexState {
        module: &shader,
        entry_point: "vs_main",
        buffers: &[Vertex::desc()],
    },
    fragment: Some(wgpu::FragmentState {
        module: &shader,
        entry_point: "fs_main",
        targets: &[Some(wgpu::ColorTargetState {
            format: config.format,
            blend: Some(wgpu::BlendState::ALPHA_BLENDING),
            write_mask: wgpu::ColorWrites::ALL,
        })],
    }),
    primitive: wgpu::PrimitiveState {
        topology: wgpu::PrimitiveTopology::TriangleList,
        strip_index_format: None,
        front_face: wgpu::FrontFace::Ccw,
        cull_mode: Some(wgpu::Face::Back),
        polygon_mode: wgpu::PolygonMode::Fill,
        unclipped_depth: false,
        conservative: false,
    },
    depth_stencil: None,
    multisample: wgpu::MultisampleState {
        count: 1,
        mask: !0,
        alpha_to_coverage_enabled: false,
    },
    multiview: None,
});
```

### Drawing with the Pipeline

```rust
fn render(&mut self) -> Result<(), wgpu::SurfaceError> {
    let output = self.surface.get_current_texture()?;
    let view = output.texture.create_view(&Default::default());
    
    let mut encoder = self.device.create_command_encoder(&wgpu::CommandEncoderDescriptor {
        label: Some("Render Encoder"),
    });
    
    {
        let mut render_pass = encoder.begin_render_pass(&wgpu::RenderPassDescriptor {
            label: Some("Render Pass"),
            color_attachments: &[Some(wgpu::RenderPassColorAttachment {
                view: &view,
                resolve_target: None,
                ops: wgpu::Operations {
                    load: wgpu::LoadOp::Clear(wgpu::Color::BLACK),
                    store: wgpu::StoreOp::Store,
                },
            })],
            depth_stencil_attachment: None,
            ..Default::default()
        });
        
        render_pass.set_pipeline(&self.render_pipeline);
        render_pass.set_vertex_buffer(0, self.vertex_buffer.slice(..));
        render_pass.draw(0..3, 0..1);  // 3 vertices, 1 instance
    }
    
    self.queue.submit(std::iter::once(encoder.finish()));
    output.present();
    
    Ok(())
}
```

---

## Part 5: Drawing Rectangles with Index Buffers (20 min)

### Why Index Buffers?

A rectangle needs 2 triangles = 6 vertices. But we only have 4 corners!

```
Without indices:     With indices:
6 vertices           4 vertices + 6 indices

0───1  2             0───1
│ ╱ │╱ │             │ ╲ │
2  3───4             │  ╲│
                     2───3

Indices: [0, 2, 1, 1, 2, 3]
```

### Implementing Rectangles

```rust
const RECT_VERTICES: &[Vertex] = &[
    Vertex { position: [-0.5, 0.5], color: [1.0, 0.0, 0.0] },   // Top-left
    Vertex { position: [0.5, 0.5], color: [0.0, 1.0, 0.0] },    // Top-right
    Vertex { position: [-0.5, -0.5], color: [0.0, 0.0, 1.0] },  // Bottom-left
    Vertex { position: [0.5, -0.5], color: [1.0, 1.0, 0.0] },   // Bottom-right
];

const RECT_INDICES: &[u16] = &[
    0, 2, 1,  // First triangle
    1, 2, 3,  // Second triangle
];

// Create index buffer
let index_buffer = device.create_buffer_init(&wgpu::util::BufferInitDescriptor {
    label: Some("Index Buffer"),
    contents: bytemuck::cast_slice(RECT_INDICES),
    usage: wgpu::BufferUsages::INDEX,
});

// In render:
render_pass.set_index_buffer(index_buffer.slice(..), wgpu::IndexFormat::Uint16);
render_pass.draw_indexed(0..6, 0, 0..1);  // 6 indices
```

---

## Part 6: Drawing Circles (15 min)

### Circle as Triangle Fan

```rust
fn create_circle_vertices(center: [f32; 2], radius: f32, segments: u32, color: [f32; 3]) -> Vec<Vertex> {
    let mut vertices = Vec::with_capacity(segments as usize + 2);
    
    // Center vertex
    vertices.push(Vertex { position: center, color });
    
    // Perimeter vertices
    for i in 0..=segments {
        let angle = (i as f32 / segments as f32) * std::f32::consts::TAU;
        vertices.push(Vertex {
            position: [
                center[0] + radius * angle.cos(),
                center[1] + radius * angle.sin(),
            ],
            color,
        });
    }
    
    vertices
}

fn create_circle_indices(segments: u32) -> Vec<u16> {
    let mut indices = Vec::with_capacity(segments as usize * 3);
    
    for i in 0..segments {
        indices.push(0);                    // Center
        indices.push(i as u16 + 1);         // Current perimeter
        indices.push(i as u16 + 2);         // Next perimeter
    }
    
    indices
}
```

---

## Part 7: Building a Shape Renderer (15 min)

### A Reusable Abstraction

```rust
pub struct ShapeRenderer {
    pipeline: wgpu::RenderPipeline,
    vertices: Vec<Vertex>,
    indices: Vec<u16>,
    vertex_buffer: wgpu::Buffer,
    index_buffer: wgpu::Buffer,
    needs_update: bool,
}

impl ShapeRenderer {
    pub fn draw_rect(&mut self, x: f32, y: f32, width: f32, height: f32, color: [f32; 3]) {
        let base_index = self.vertices.len() as u16;
        
        self.vertices.extend_from_slice(&[
            Vertex { position: [x, y], color },
            Vertex { position: [x + width, y], color },
            Vertex { position: [x, y + height], color },
            Vertex { position: [x + width, y + height], color },
        ]);
        
        self.indices.extend_from_slice(&[
            base_index, base_index + 2, base_index + 1,
            base_index + 1, base_index + 2, base_index + 3,
        ]);
        
        self.needs_update = true;
    }
    
    pub fn clear(&mut self) {
        self.vertices.clear();
        self.indices.clear();
        self.needs_update = true;
    }
    
    pub fn flush(&mut self, device: &wgpu::Device, queue: &wgpu::Queue) {
        if self.needs_update && !self.vertices.is_empty() {
            // Recreate or update buffers
            self.vertex_buffer = device.create_buffer_init(&wgpu::util::BufferInitDescriptor {
                label: Some("Shape Vertex Buffer"),
                contents: bytemuck::cast_slice(&self.vertices),
                usage: wgpu::BufferUsages::VERTEX,
            });
            
            self.index_buffer = device.create_buffer_init(&wgpu::util::BufferInitDescriptor {
                label: Some("Shape Index Buffer"),
                contents: bytemuck::cast_slice(&self.indices),
                usage: wgpu::BufferUsages::INDEX,
            });
            
            self.needs_update = false;
        }
    }
}
```

---

## Summary

| Concept | Description |
|---------|-------------|
| NDC | Normalized Device Coordinates (-1 to 1) |
| Vertex | Point with position and attributes |
| Vertex Buffer | GPU storage for vertex data |
| Index Buffer | Reuse vertices for triangles |
| Shader | GPU program (vertex/fragment) |
| Render Pipeline | Complete draw configuration |

---

## Key Takeaways

1. **Everything is triangles** - Even circles are made of triangles
2. **Index buffers save memory** - Reuse vertices efficiently
3. **Shaders run in parallel** - One per vertex/pixel simultaneously
4. **NDC is the target space** - Convert screen coords → NDC
5. **Flush batched draws** - Minimize buffer updates

---

## Up Next

In **Lecture 04**, we'll handle user input properly—keyboard, mouse, and touch events with proper event propagation.

---

## References

- [Learn Wgpu - The Pipeline](https://sotrh.github.io/learn-wgpu/beginner/tutorial3-pipeline/)
- [WGSL Specification](https://www.w3.org/TR/WGSL/)
- [bytemuck Documentation](https://docs.rs/bytemuck)
