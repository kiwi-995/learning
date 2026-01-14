# Lecture 06: Layout Engine Design — Part 1: Fundamentals

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lectures 01-05  
> **Goal:** Build a layout engine with box model and sizing constraints

---

## Overview

Layout is where math meets UI. We'll build the foundation for positioning widgets.

**Topics:**
- The box model
- Size constraints and intrinsic sizing
- Two-pass layout algorithm
- Parent-child negotiation

---

## Part 1: The Box Model (25 min)

### CSS Box Model (Inspiration)

```
┌─────────────────────────────────────────────┐
│                   MARGIN                    │
│  ┌───────────────────────────────────────┐  │
│  │                BORDER                 │  │
│  │  ┌─────────────────────────────────┐  │  │
│  │  │            PADDING              │  │  │
│  │  │  ┌───────────────────────────┐  │  │  │
│  │  │  │         CONTENT           │  │  │  │
│  │  │  └───────────────────────────┘  │  │  │
│  │  └─────────────────────────────────┘  │  │
│  └───────────────────────────────────────┘  │
└─────────────────────────────────────────────┘
```

### Implementation

```rust
#[derive(Clone, Copy, Default)]
pub struct EdgeInsets {
    pub top: f32,
    pub right: f32,
    pub bottom: f32,
    pub left: f32,
}

impl EdgeInsets {
    pub fn all(value: f32) -> Self {
        Self { top: value, right: value, bottom: value, left: value }
    }
    
    pub fn horizontal(&self) -> f32 { self.left + self.right }
    pub fn vertical(&self) -> f32 { self.top + self.bottom }
}

pub struct BoxConstraints {
    pub padding: EdgeInsets,
    pub margin: EdgeInsets,
    pub border_width: f32,
}
```

---

## Part 2: Size Constraints (25 min)

### Constraint Types

```rust
#[derive(Clone, Copy)]
pub struct SizeConstraint {
    pub min: f32,      // Minimum size
    pub preferred: f32, // Ideal size
    pub max: f32,      // Maximum (f32::INFINITY for unconstrained)
}

impl SizeConstraint {
    pub fn fixed(size: f32) -> Self {
        Self { min: size, preferred: size, max: size }
    }
    
    pub fn flexible(min: f32, preferred: f32, max: f32) -> Self {
        Self { min, preferred, max }
    }
    
    pub fn clamp(&self, value: f32) -> f32 {
        value.clamp(self.min, self.max)
    }
}

pub struct SizeConstraints {
    pub width: SizeConstraint,
    pub height: SizeConstraint,
}
```

### Intrinsic Sizing

```rust
pub enum IntrinsicSize {
    /// Shrink to fit content
    MinContent,
    /// Expand to fill available space
    MaxContent,
    /// Use preferred size
    FitContent,
    /// Explicit size
    Fixed(f32),
    /// Percentage of parent
    Percent(f32),
}
```

---

## Part 3: The Measure Phase (25 min)

### Measure Context

```rust
pub struct MeasureContext {
    /// Available space from parent
    pub available_width: f32,
    pub available_height: f32,
    /// Text measurement
    pub measure_text: Box<dyn Fn(&str, f32) -> (f32, f32)>,
}

impl Widget for Label {
    fn measure(&self, ctx: &MeasureContext) -> SizeConstraints {
        let (text_width, text_height) = (ctx.measure_text)(&self.text, self.font_size);
        SizeConstraints {
            width: SizeConstraint::fixed(text_width + self.padding.horizontal()),
            height: SizeConstraint::fixed(text_height + self.padding.vertical()),
        }
    }
}
```

### Container Measurement

```rust
impl Widget for VStack {
    fn measure(&self, ctx: &MeasureContext) -> SizeConstraints {
        let mut total_height = 0.0;
        let mut max_width = 0.0;
        let mut min_height = 0.0;
        
        for child in &self.children {
            let child_constraints = child.measure(ctx);
            
            min_height += child_constraints.height.min;
            total_height += child_constraints.height.preferred;
            max_width = max_width.max(child_constraints.width.preferred);
        }
        
        // Add spacing
        let spacing_total = self.spacing * (self.children.len().saturating_sub(1)) as f32;
        
        SizeConstraints {
            width: SizeConstraint::flexible(0.0, max_width, f32::INFINITY),
            height: SizeConstraint::flexible(
                min_height + spacing_total,
                total_height + spacing_total,
                f32::INFINITY
            ),
        }
    }
}
```

---

## Part 4: The Layout Phase (25 min)

### Layout Algorithm

```rust
pub struct LayoutContext {
    pub measure: MeasureContext,
}

impl Widget for VStack {
    fn layout(&mut self, bounds: Rect, ctx: &mut LayoutContext) {
        let mut y = bounds.y + self.padding.top;
        let available_width = bounds.width - self.padding.horizontal();
        
        // Calculate how much extra space we have
        let measured = self.measure(&ctx.measure);
        let extra_space = bounds.height - measured.height.preferred;
        
        for child in &mut self.children {
            let child_constraints = child.measure(&ctx.measure);
            
            // Determine child height
            let child_height = child_constraints.height.preferred;
            
            // Determine child width (expand or use preferred)
            let child_width = if self.cross_axis_alignment == CrossAxisAlignment::Stretch {
                available_width
            } else {
                child_constraints.width.preferred.min(available_width)
            };
            
            // Calculate x position based on alignment
            let child_x = match self.cross_axis_alignment {
                CrossAxisAlignment::Start => bounds.x + self.padding.left,
                CrossAxisAlignment::Center => bounds.x + (bounds.width - child_width) / 2.0,
                CrossAxisAlignment::End => bounds.x + bounds.width - child_width - self.padding.right,
                CrossAxisAlignment::Stretch => bounds.x + self.padding.left,
            };
            
            let child_bounds = Rect::new(child_x, y, child_width, child_height);
            child.layout(child_bounds, ctx);
            
            y += child_height + self.spacing;
        }
    }
}
```

---

## Part 5: Size Negotiation (15 min)

### Parent vs Child Control

```
Parent-Controlled:               Child-Controlled:
┌────────────────────┐           ┌──────┐
│    "You get this   │           │ "I   │
│     much space"    │           │ need │
│                    │           │ this │
└────────────────────┘           └──────┘

Negotiated:
Parent: "I have 400px available"
Child: "I need at least 100px, prefer 200px"
Result: Child gets 200px (or less if constrained)
```

### Constraint Propagation

```rust
fn layout_child(parent_bounds: Rect, child: &mut dyn Widget, ctx: &mut LayoutContext) {
    // Get child's wishes
    let child_constraints = child.measure(&ctx.measure);
    
    // Apply parent limits
    let final_width = child_constraints.width.clamp_to(parent_bounds.width);
    let final_height = child_constraints.height.clamp_to(parent_bounds.height);
    
    let child_bounds = Rect::new(
        parent_bounds.x,
        parent_bounds.y,
        final_width,
        final_height
    );
    
    child.layout(child_bounds, ctx);
}
```

---

## Summary

| Concept | Purpose |
|---------|---------|
| Box Model | Structural spacing (padding, margin, border) |
| Constraints | min/preferred/max sizing |
| Measure | Calculate desired size |
| Layout | Assign final positions |
| Negotiation | Parent and child agree on size |

---

## Key Takeaways

1. **Two-pass layout**: Measure first, layout second
2. **Constraints flow down**: Parent tells child available space
3. **Sizes flow up**: Child tells parent desired size
4. **Box model adds structure**: Padding, margin, border
5. **Flexibility is key**: min/preferred/max allows adaptation

---

## Up Next

In **Lecture 07**, we implement Flexbox and constraint-based layout systems.
