# Exercise Sheet 09: Reactive Signals

> **Estimated Time:** 2 hours  
> **Difficulty:** ⭐⭐⭐⭐

---

## Exercise 1: Basic Signal (25 min)

Implement `Signal<T>`:
```rust
impl<T: Clone> Signal<T> {
    fn new(value: T) -> Self;
    fn get(&self) -> T;
    fn set(&self, value: T);
    fn subscribe<F: FnMut(&T) + 'static>(&self, f: F);
}
```

Test: signal changes trigger subscriber.

---

## Exercise 2: Effects with Auto-Tracking (35 min)

Implement automatic dependency tracking:
```rust
fn create_effect<F: FnMut() + 'static>(f: F);
```

The effect should:
1. Run immediately
2. Re-run when any signal accessed during execution changes
3. Only track signals read in the most recent run

---

## Exercise 3: Computed Values (30 min)

Implement `Computed<T>`:
- Lazy evaluation (only compute when accessed)
- Caching (recompute only when dependencies change)
- Track dependencies automatically

```rust
let a = Signal::new(2);
let b = Signal::new(3);
let sum = computed(|| a.get() + b.get());
assert_eq!(sum.get(), 5);
a.set(10);
assert_eq!(sum.get(), 13);
```

---

## Exercise 4: Reactive Widget (30 min)

Connect signals to your UI framework:
```rust
fn reactive_label<F: Fn() -> String + 'static>(f: F) -> Label;
```

The label should update automatically when signals change.

Build: A form with first/last name inputs and a "Hello, Full Name!" label.

---

## Theory Problems

### T1: Dependency Graph
Given signals A, B, C and computed D = A + B, E = B + C, F = D + E:
1. Draw the dependency graph
2. If A changes, which nodes recompute?
3. In what order should recomputation happen?

### T2: Memory Safety
In a signal system with closures:
1. What prevents use-after-free?
2. How do we handle signal disposal?
3. What causes memory leaks and how to prevent them?

### T3: Performance Analysis
For N signals with M effects:
1. Time complexity of one signal update?
2. Space complexity of the subscription system?

---

## Challenge: Batched Updates

Implement transaction batching:
```rust
batch(|| {
    signal_a.set(1);
    signal_b.set(2); // Effects don't run yet
}); // All effects run once here
```

Avoid redundant computations when multiple signals change together.
