# Exercise Sheet 07: Flexbox & Constraints

> **Estimated Time:** 2 hours  
> **Difficulty:** ⭐⭐⭐⭐

---

## Exercise 1: Flexbox Core (35 min)

Implement the flex sizing algorithm:
```rust
fn distribute_flex_space(
    items: &[FlexItem],
    available: f32,
) -> Vec<f32>;
```

Test: 3 items with flex_grow 1, 2, 1 in 400px space with basis 50px each.

---

## Exercise 2: justify-content (25 min)

Implement all justify-content modes:
- FlexStart, FlexEnd, Center
- SpaceBetween, SpaceAround, SpaceEvenly

Visualize with a row of colored boxes.

---

## Exercise 3: Complete Flexbox (30 min)

Build a flex container with:
- Row and column direction
- justify-content and align-items
- Gap between items
- Nested flex containers

Recreate a navigation bar: logo left, links center, profile right.

---

## Exercise 4: taffy Integration (30 min)

Replace your layout engine with `taffy`:
1. Add taffy to dependencies
2. Create layout tree matching your widget tree
3. Compute layout and apply positions to widgets

---

## Theory Problems

### T1: Flexbox Space Distribution
3 items: flex-grow 2, 1, 1, flex-basis 100px each
Available: 500px  
1. Calculate each item's final size
2. What if available is only 250px with flex-shrink 1, 2, 1?

### T2: Constraint Satisfaction
Given constraints: A.right = B.left, B.width = 100, Container.width = 400
1. Is there a unique solution?
2. What additional constraint determines A.width?

### T3: Algorithm Complexity (Advanced)
Prove: Flexbox layout is O(n) for n items in a single container.

---

## Challenge: Grid Layout

Implement CSS Grid-like layout:
```rust
struct Grid {
    columns: Vec<TrackSize>,  // fr, px, auto
    rows: Vec<TrackSize>,
    gap: f32,
}
```

Handle `fr` units (fractional) and `auto` sizing.
