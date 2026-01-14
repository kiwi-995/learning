# Exercise Sheet 03: 2D Rendering & Drawing Primitives

> **Estimated Time:** 2 hours  
> **Difficulty:** ⭐⭐⭐ (Intermediate-Advanced)

---

## Exercise 1: Colorful Triangle (25 min)

Render a triangle with interpolated vertex colors.

### Requirements
- Vertex shader: position + color input
- Fragment shader: output interpolated color
- Top vertex: Red, Bottom-left: Green, Bottom-right: Blue

---

## Exercise 2: Rectangle Renderer (30 min)

### 2.1 Implement `draw_rect()`
Convert screen coordinates to NDC, add 4 vertices and 6 indices.

### 2.2 Draw a Layout
```
┌────────────────────────────────┐
│           Header               │
├──────────┬─────────────────────┤
│ Sidebar  │      Content        │
└──────────┴─────────────────────┘
```

---

## Exercise 3: Circle Renderer (30 min)

### 3.1 Triangle Fan
Implement `draw_circle(cx, cy, radius, color, segments)`

### 3.2 Segment Comparison  
Render circles with 6, 12, 24, and 64 segments to see quality difference.

### 3.3 Challenge: Outlined Circle
Draw only the perimeter using a ring (inner + outer circle).

---

## Exercise 4: Shape DSL (25 min)

```rust
enum Shape {
    Rect { x: f32, y: f32, width: f32, height: f32, color: Color },
    Circle { cx: f32, cy: f32, radius: f32, color: Color },
}
```

Implement `draw_shape()` and create a scene (sky, sun, ground, house).

---

## Exercise 5: Uniform Buffers (30 min)

Pass window dimensions to shader instead of hardcoding NDC conversion:
1. Create uniform buffer with screen size
2. Create bind group and layout
3. Convert coords in vertex shader
4. Update buffer on resize

---

## Theory Problems

### T1: Rasterization
Triangle A=(-0.5,-0.5), B=(0.5,-0.5), C=(0.0,0.5) on 100×100 screen:
1. Convert to screen coordinates
2. Which pixels are inside?
3. Color at pixel (50, 60) with vertex colors R, G, B?

### T2: Memory Analysis
For N rectangles, compare memory with/without index buffers.

### T3: Circle Approximation
Derive error formula for N-segment circle. How many segments for <1px error on r=100?

---

## Hints

<details>
<summary>Ring for Outlined Circle</summary>
Create outer vertices at radius, inner at radius-thickness, connect with quads.
</details>
