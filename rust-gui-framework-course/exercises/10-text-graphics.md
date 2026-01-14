# Exercise Sheet 10: Text & Custom Graphics

> **Estimated Time:** 1.5-2 hours  
> **Difficulty:** ⭐⭐⭐

---

## Exercise 1: Basic Text Rendering (30 min)

Using fontdue:
1. Load a TTF font
2. Rasterize individual characters
3. Upload to GPU as textures
4. Render "Hello, World!" on screen

---

## Exercise 2: Glyph Atlas (35 min)

Implement a glyph cache:
```rust
struct GlyphAtlas {
    fn get_or_insert(&mut self, char: char, size: u32) -> &GlyphInfo;
    fn render_text(&self, text: &str, x: f32, y: f32) -> Vec<Quad>;
}
```

Draw a paragraph with word wrapping.

---

## Exercise 3: Path Drawing (30 min)

Using your shape renderer, add path support:
1. Implement `Path` builder with move/line/curve/close
2. Create a `draw_path()` function
3. Draw: rounded rectangles, stars, hearts

---

## Exercise 4: Simple Chart (25 min)

Build a bar chart widget:
- Takes data: `&[(String, f32)]`
- Auto-scales to fit container
- Labels on x-axis
- Values on y-axis

---

## Theory Problems

### T1: Font Hinting
1. What is font hinting?
2. Why is it more important at small sizes?
3. How does subpixel rendering improve clarity?

### T2: Bezier Curves
For cubic Bezier P0, P1, P2, P3:
1. Write the parametric equation B(t)
2. How many line segments for visually smooth curve at 100px length?

---

## Challenge: SVG Parser

Parse simple SVG paths:
```
M 10 10 L 100 10 L 100 100 Z
```

Convert to your Path type and render.
