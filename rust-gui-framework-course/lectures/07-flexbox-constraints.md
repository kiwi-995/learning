# Lecture 07: Layout Engine Design — Part 2: Flexbox & Constraints

> **Duration:** ~1.5-2 hours  
> **Prerequisites:** Lecture 06  
> **Goal:** Implement Flexbox layout and understand constraint-based systems

---

## Overview

We'll implement the Flexbox algorithm and explore constraint-based layout with Cassowary.

---

## Part 1: Flexbox Concepts (20 min)

### Main Axis vs Cross Axis

```
flex-direction: row
┌─────────────────────────────────────┐
│ ┌───┐  ┌───┐  ┌───┐                 │
│ │ 1 │→ │ 2 │→ │ 3 │   Main Axis →   │
│ └───┘  └───┘  └───┘                 │
│   ↓ Cross Axis                      │
└─────────────────────────────────────┘

flex-direction: column
┌─────────────────┐
│ ┌─────────────┐ │
│ │      1      │ │ ↓ Main Axis
│ └─────────────┘ │
│ ┌─────────────┐ │
│ │      2      │ │
│ └─────────────┘ │
│  → Cross Axis   │
└─────────────────┘
```

### Flex Properties

```rust
pub struct FlexItem {
    pub flex_grow: f32,    // How much to grow (0 = don't grow)
    pub flex_shrink: f32,  // How much to shrink (1 = normal shrink)
    pub flex_basis: FlexBasis, // Initial size
}

pub enum FlexBasis {
    Auto,       // Use content size
    Fixed(f32), // Explicit size
}
```

---

## Part 2: Flexbox Algorithm (30 min)

### The Algorithm

```rust
impl FlexContainer {
    pub fn layout(&mut self, available: Rect) {
        // 1. Determine main axis size of each child
        let mut main_sizes: Vec<f32> = self.children.iter()
            .map(|c| match c.flex_basis {
                FlexBasis::Fixed(s) => s,
                FlexBasis::Auto => c.widget.measure(&ctx).preferred_main_size(),
            })
            .collect();
        
        // 2. Calculate total and remaining space
        let total_basis: f32 = main_sizes.iter().sum();
        let available_main = self.main_axis_size(available);
        let remaining = available_main - total_basis;
        
        // 3. Distribute remaining space
        if remaining > 0.0 {
            // Grow items with flex_grow > 0
            let total_grow: f32 = self.children.iter().map(|c| c.flex_grow).sum();
            if total_grow > 0.0 {
                for (i, child) in self.children.iter().enumerate() {
                    main_sizes[i] += remaining * (child.flex_grow / total_grow);
                }
            }
        } else if remaining < 0.0 {
            // Shrink items with flex_shrink > 0
            let total_shrink: f32 = self.children.iter().map(|c| c.flex_shrink).sum();
            if total_shrink > 0.0 {
                for (i, child) in self.children.iter().enumerate() {
                    main_sizes[i] += remaining * (child.flex_shrink / total_shrink);
                }
            }
        }
        
        // 4. Position children
        let mut main_offset = 0.0;
        for (i, child) in self.children.iter_mut().enumerate() {
            let bounds = self.child_bounds(available, main_offset, main_sizes[i]);
            child.widget.layout(bounds, &mut ctx);
            main_offset += main_sizes[i] + self.gap;
        }
    }
}
```

---

## Part 3: Alignment (20 min)

### justify-content (Main Axis)

```rust
pub enum JustifyContent {
    FlexStart,    // Pack at start
    FlexEnd,      // Pack at end
    Center,       // Center
    SpaceBetween, // Even space between items
    SpaceAround,  // Even space around items
    SpaceEvenly,  // Even space everywhere
}

fn calculate_main_positions(
    &self,
    sizes: &[f32],
    available: f32,
) -> Vec<f32> {
    let total: f32 = sizes.iter().sum();
    let remaining = available - total;
    let count = sizes.len();
    
    match self.justify_content {
        JustifyContent::FlexStart => cumulative_positions(sizes, 0.0, 0.0),
        JustifyContent::FlexEnd => cumulative_positions(sizes, remaining, 0.0),
        JustifyContent::Center => cumulative_positions(sizes, remaining / 2.0, 0.0),
        JustifyContent::SpaceBetween => {
            let gap = remaining / (count - 1).max(1) as f32;
            cumulative_positions(sizes, 0.0, gap)
        }
        // ... etc
    }
}
```

### align-items (Cross Axis)

```rust
pub enum AlignItems {
    FlexStart, Center, FlexEnd, Stretch, Baseline,
}
```

---

## Part 4: Constraint-Based Layout (25 min)

### Cassowary Algorithm

Express layout as linear equations:
```
button.left = container.left + 20
button.right = container.right - 20
label.top = button.bottom + 10
label.centerX = container.centerX
```

### Constraint Types

```rust
pub enum Constraint {
    Equal(Variable, Variable),           // a = b
    Offset(Variable, Variable, f32),     // a = b + c
    Proportional(Variable, Variable, f32), // a = b * c
}

pub enum Strength {
    Required,  // Must satisfy
    Strong,    // Try hard
    Medium,    // Moderate preference
    Weak,      // Nice to have
}
```

### Using `taffy` (Production Layout)

```rust
use taffy::prelude::*;

let mut taffy = TaffyTree::new();

let child = taffy.new_leaf(Style {
    size: Size { width: points(100.0), height: points(50.0) },
    ..Default::default()
})?;

let container = taffy.new_with_children(
    Style {
        display: Display::Flex,
        flex_direction: FlexDirection::Row,
        justify_content: JustifyContent::Center,
        ..Default::default()
    },
    &[child],
)?;

taffy.compute_layout(container, Size::MAX_CONTENT)?;
let layout = taffy.layout(child)?;
```

---

## Part 5: Comparison (15 min)

| Feature | Flexbox | Constraints |
|---------|---------|-------------|
| Mental Model | Linear flow | Relationship graph |
| Flexibility | 1D | 2D |
| Performance | O(n) | O(n²) typically |
| Use Case | Most layouts | Complex relationships |

---

## Key Takeaways

1. **Flexbox works in one dimension** at a time
2. **flex-grow/shrink** controls space distribution
3. **Constraints** are more powerful but complex
4. **Use `taffy`** for production Flexbox

---

## Up Next

**Lecture 08**: Declarative UI and the Elm Architecture.
