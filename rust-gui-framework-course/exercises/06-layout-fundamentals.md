# Exercise Sheet 06: Layout Fundamentals

> **Estimated Time:** 1.5-2 hours  
> **Difficulty:** ⭐⭐⭐

---

## Exercise 1: Box Model Implementation (30 min)

Implement the complete box model:
```rust
struct BoxModel {
    content: Size,
    padding: EdgeInsets,
    border: EdgeInsets,
    margin: EdgeInsets,
}

impl BoxModel {
    fn outer_size(&self) -> Size;
    fn content_rect(&self, origin: Point) -> Rect;
    fn padding_rect(&self, origin: Point) -> Rect;
    fn border_rect(&self, origin: Point) -> Rect;
}
```

Render a widget showing all four regions with different colors.

---

## Exercise 2: Size Constraints (25 min)

Implement and test size negotiation:
```rust
fn negotiate_size(
    requested: SizeConstraint,
    available: f32,
) -> f32;
```

Test cases:
- Fixed size that fits
- Fixed size larger than available
- Flexible with plenty of space
- Flexible with limited space

---

## Exercise 3: Stack Layout (35 min)

Implement complete VStack and HStack:
- Spacing between children
- Main axis alignment (start, center, end, space-between, space-around)
- Cross axis alignment (start, center, end, stretch)

Test with mixed fixed and flexible children.

---

## Exercise 4: Nested Layout (30 min)

Build this layout using your stack widgets:
```
┌────────────────────────────────────┐
│ ┌──────────────────────────────┐   │
│ │           Header             │   │
│ └──────────────────────────────┘   │
│ ┌───────────┐ ┌────────────────┐   │
│ │           │ │                │   │
│ │  Sidebar  │ │    Content     │   │
│ │           │ │                │   │
│ └───────────┘ └────────────────┘   │
│ ┌──────────────────────────────┐   │
│ │           Footer             │   │
│ └──────────────────────────────┘   │
└────────────────────────────────────┘
```

---

## Theory Problems

### T1: Constraint Analysis
Given: Parent has 400px width  
Child wants: min=100, preferred=500, max=800

1. What size does child get?
2. What if parent only has 80px?
3. Design a policy for this situation

### T2: Layout Complexity
For a tree of N widgets with max depth D:
1. What's the time complexity of measure phase?
2. What's the time complexity of layout phase?
3. How can we optimize for large widget trees?

---

## Challenge: Percentage Sizing

Implement percentage-based sizing:
```rust
enum Size {
    Fixed(f32),
    Percent(f32),
    Auto,
}
```

Handle: 50% of parent, nested percentages, circular dependencies.
