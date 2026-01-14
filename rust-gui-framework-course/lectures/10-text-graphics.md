# Lecture 10: Text Rendering & Custom Graphics

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lectures 01-09  
> **Goal:** Render text and custom vector graphics

---

## Overview

Text is the most important UI element. We'll learn how to render it properly and add custom graphics capabilities.

---

## Part 1: Font Basics (20 min)

### Font Anatomy

```
         ┌─────────────── Ascender
         │   ┌───────────  Cap Height
         │   │   ┌─────── x-Height
    h  T │ x │   │
    ─────┼───┼───│──────── Baseline
         │   │   │ g  y
         │   │   └─────── Descender
         │   │
```

### Fonts in Rust

```rust
// Using fontdue (simple, fast)
use fontdue::{Font, FontSettings};

let font = Font::from_bytes(
    include_bytes!("../fonts/Roboto-Regular.ttf") as &[u8],
    FontSettings::default()
)?;

// Rasterize a glyph
let (metrics, bitmap) = font.rasterize('A', 32.0);
// bitmap is a Vec<u8> with coverage values
```

---

## Part 2: Text Layout (25 min)

### Measuring Text

```rust
pub struct TextMetrics {
    pub width: f32,
    pub height: f32,
    pub ascent: f32,
    pub descent: f32,
}

fn measure_text(font: &Font, text: &str, size: f32) -> TextMetrics {
    let mut width = 0.0;
    let mut max_ascent = 0.0;
    let mut max_descent = 0.0;
    
    for c in text.chars() {
        let metrics = font.metrics(c, size);
        width += metrics.advance_width;
        max_ascent = max_ascent.max(metrics.bounds.ymin.abs());
        max_descent = max_descent.max(metrics.bounds.height - metrics.bounds.ymin.abs());
    }
    
    TextMetrics {
        width,
        height: max_ascent + max_descent,
        ascent: max_ascent,
        descent: max_descent,
    }
}
```

### Line Wrapping

```rust
fn wrap_text(font: &Font, text: &str, size: f32, max_width: f32) -> Vec<String> {
    let mut lines = Vec::new();
    let mut current_line = String::new();
    let mut current_width = 0.0;
    
    for word in text.split_whitespace() {
        let word_width = measure_text(font, word, size).width;
        let space_width = measure_text(font, " ", size).width;
        
        if current_width + word_width > max_width && !current_line.is_empty() {
            lines.push(current_line.trim().to_string());
            current_line = word.to_string();
            current_width = word_width;
        } else {
            if !current_line.is_empty() {
                current_line.push(' ');
                current_width += space_width;
            }
            current_line.push_str(word);
            current_width += word_width;
        }
    }
    
    if !current_line.is_empty() {
        lines.push(current_line);
    }
    lines
}
```

---

## Part 3: GPU Text Rendering (25 min)

### Glyph Atlas

```rust
pub struct GlyphAtlas {
    texture: wgpu::Texture,
    cache: HashMap<(char, u32), GlyphInfo>,
    packer: AtlasPacker,
}

struct GlyphInfo {
    uv_rect: Rect,  // Position in atlas
    size: (u32, u32),
    offset: (f32, f32),
    advance: f32,
}

impl GlyphAtlas {
    fn get_glyph(&mut self, font: &Font, char: char, size: u32) -> &GlyphInfo {
        let key = (char, size);
        if !self.cache.contains_key(&key) {
            let (metrics, bitmap) = font.rasterize(char, size as f32);
            let position = self.packer.pack(metrics.width, metrics.height);
            // Upload bitmap to texture at position
            self.upload_glyph(&bitmap, position, metrics.width, metrics.height);
            // Cache the info
            self.cache.insert(key, GlyphInfo { /* ... */ });
        }
        &self.cache[&key]
    }
}
```

---

## Part 4: Custom Vector Graphics (25 min)

### Path API

```rust
pub struct Path {
    commands: Vec<PathCommand>,
}

pub enum PathCommand {
    MoveTo(f32, f32),
    LineTo(f32, f32),
    QuadTo(f32, f32, f32, f32),       // Control, end
    CubicTo(f32, f32, f32, f32, f32, f32), // Control1, control2, end
    Close,
}

impl Path {
    pub fn new() -> Self { Self { commands: Vec::new() } }
    pub fn move_to(&mut self, x: f32, y: f32) -> &mut Self;
    pub fn line_to(&mut self, x: f32, y: f32) -> &mut Self;
    pub fn quad_to(&mut self, cx: f32, cy: f32, x: f32, y: f32) -> &mut Self;
    pub fn close(&mut self) -> &mut Self;
}

// Usage
let mut path = Path::new();
path.move_to(0.0, 0.0)
    .line_to(100.0, 0.0)
    .quad_to(150.0, 50.0, 100.0, 100.0)
    .line_to(0.0, 100.0)
    .close();
```

### Rendering Paths

For GPU rendering, tessellate paths to triangles using crates like `lyon`:

```rust
use lyon::tessellation::*;

let mut geometry = VertexBuffers::new();
let mut tessellator = FillTessellator::new();

tessellator.tessellate_path(
    &path,
    &FillOptions::default(),
    &mut BuffersBuilder::new(&mut geometry, |vertex: FillVertex| {
        Vertex {
            position: vertex.position().to_array(),
            color: [1.0, 0.0, 0.0],
        }
    }),
)?;
```

---

## Part 5: Putting It Together (15 min)

### Text Renderer

```rust
pub struct TextRenderer {
    atlas: GlyphAtlas,
    vertices: Vec<Vertex>,
    indices: Vec<u16>,
}

impl TextRenderer {
    pub fn draw_text(&mut self, text: &str, x: f32, y: f32, size: f32, color: Color) {
        let mut cursor_x = x;
        
        for c in text.chars() {
            let glyph = self.atlas.get_glyph(&self.font, c, size as u32);
            
            // Add quad for this glyph
            self.add_glyph_quad(
                cursor_x + glyph.offset.0,
                y + glyph.offset.1,
                glyph.size.0 as f32,
                glyph.size.1 as f32,
                glyph.uv_rect,
                color,
            );
            
            cursor_x += glyph.advance;
        }
    }
}
```

---

## Key Takeaways

1. **Fonts are complex** - use fontdue or cosmic-text
2. **Glyph atlases** cache rasterized glyphs on GPU
3. **Text layout** involves line wrapping, alignment
4. **Path tessellation** converts curves to triangles
5. **lyon crate** is excellent for vector graphics

---

## Up Next

**Lecture 11**: Cross-platform deployment (web, desktop, mobile).
