# Lecture 2: Ownership, Borrowing & Memory Safety

> *"Rust's ownership system is the key to its safety and performance guarantees."*

This is arguably the most important lecture in the Rust portion of the course. Ownership is what makes Rust unique — and it's the concept that trips up most newcomers. Take your time here; understanding ownership deeply will pay dividends throughout your compiler journey.

## 🎯 Learning Objectives

By the end of this lecture, you will be able to:
- Explain the three rules of ownership
- Distinguish between moves and copies
- Use references to borrow values without taking ownership
- Apply lifetime annotations when needed
- Debug common ownership errors

## 📖 Topics

### 2.1 Memory Management: A Historical Perspective

Every programming language must decide how to manage memory:

| Approach | Languages | Pros | Cons |
|----------|-----------|------|------|
| Manual | C, C++ | Full control, fast | Bugs: leaks, use-after-free, double-free |
| Garbage Collection | Java, Python, Go, JS | Safe, convenient | Pauses, overhead, less predictable |
| **Ownership** | Rust | Safe AND fast | Learning curve |

**The Problem with Manual Memory:**
```c
// C code - spot the bugs!
char* create_greeting() {
    char* msg = malloc(20);
    strcpy(msg, "Hello");
    return msg;  // Caller must remember to free!
}

void use_greeting() {
    char* greeting = create_greeting();
    printf("%s\n", greeting);
    free(greeting);
    printf("%s\n", greeting);  // BUG: use-after-free!
}
```

**The Problem with GC:**
```python
# Python - safe but pays runtime cost
def create_list():
    return [i for i in range(1000000)]  # GC tracks this

# GC may pause at unpredictable times
```

**Rust's Solution:** Compile-time ownership rules that guarantee safety with zero runtime cost.

### 2.2 The Stack and the Heap

Understanding where data lives is crucial for ownership.

**Stack:**
- Fixed size, known at compile time
- LIFO (last in, first out)
- Very fast allocation/deallocation
- Examples: `i32`, `f64`, `bool`, `char`, fixed-size arrays

**Heap:**
- Dynamic size, determined at runtime
- Slower allocation (need to find space)
- Data persists until explicitly freed
- Examples: `String`, `Vec<T>`, `Box<T>`

```rust
fn main() {
    let x = 5;                    // Stack: 4 bytes for i32
    let y = 3.14;                 // Stack: 8 bytes for f64
    
    let s = String::from("hello"); // Stack: ptr + len + capacity (24 bytes)
                                   // Heap: "hello" (5 bytes)
}
```

**Diagram:**
```
Stack                          Heap
┌─────────────┐               ┌───────────────┐
│ x: 5        │               │ h e l l o     │
├─────────────┤               └───────────────┘
│ y: 3.14     │                      ↑
├─────────────┤                      │
│ s.ptr ──────┼──────────────────────┘
│ s.len: 5    │
│ s.cap: 5    │
└─────────────┘
```

### 2.3 Ownership Rules

**The Three Rules:**

1. **Each value has exactly one owner** (a variable)
2. **There can only be one owner at a time**
3. **When the owner goes out of scope, the value is dropped** (memory freed)

```rust
fn main() {
    {
        let s = String::from("hello");  // s owns the String
        // s is valid here
    }  // s goes out of scope, String is dropped, memory freed
    
    // s is NOT valid here
}
```

### 2.4 Move Semantics

When you assign a heap-allocated value to another variable, ownership **moves**:

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1;  // Ownership MOVES from s1 to s2
    
    // println!("{}", s1);  // ERROR: s1 is no longer valid
    println!("{}", s2);     // OK: s2 owns the String
}
```

**Why move instead of copy?**
If both `s1` and `s2` pointed to the same heap data, which one frees it? We'd get a double-free bug. Instead, Rust invalidates `s1`.

**Diagram:**
```
After: let s2 = s1;

Stack                          Heap
┌─────────────┐               ┌───────────────┐
│ s1 (invalid)│               │ h e l l o     │
├─────────────┤               └───────────────┘
│ s2.ptr ─────┼──────────────────────↑
│ s2.len: 5   │
│ s2.cap: 5   │
└─────────────┘
```

**Moves happen on:**
- Assignment: `let s2 = s1;`
- Function calls: `takes_ownership(s1);`
- Returns: `let s2 = returns_ownership();`

### 2.5 The Copy Trait

Some types are so cheap to copy that Rust copies them automatically:

```rust
fn main() {
    let x = 5;
    let y = x;  // x is COPIED, not moved
    
    println!("x = {}, y = {}", x, y);  // Both valid!
}
```

**Types that implement Copy:**
- All integers (`i32`, `u64`, etc.)
- Floating point (`f32`, `f64`)
- Booleans (`bool`)
- Characters (`char`)
- Tuples of `Copy` types: `(i32, bool)`

**Types that do NOT implement Copy (they move):**
- `String`
- `Vec<T>`
- Any type with heap allocation

### 2.6 Clone: Explicit Deep Copy

When you need a real copy of heap data, use `.clone()`:

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1.clone();  // Deep copy of heap data
    
    println!("{} and {}", s1, s2);  // Both valid!
}
```

**Warning:** `.clone()` can be expensive for large data. Use it intentionally.

### 2.7 Borrowing with References

What if you want to use a value without taking ownership? **Borrow** it with a reference:

```rust
fn main() {
    let s = String::from("hello");
    
    let len = calculate_length(&s);  // &s creates a reference (borrow)
    
    println!("The length of '{}' is {}", s, len);  // s still valid!
}

fn calculate_length(s: &String) -> usize {  // s is a reference to a String
    s.len()
}  // s goes out of scope, but since it doesn't own the data, nothing happens
```

**Diagram:**
```
Stack                          Heap
┌─────────────┐               ┌───────────────┐
│ s.ptr ──────┼──────────────→│ h e l l o     │
│ s.len: 5    │               └───────────────┘
│ s.cap: 5    │                      ↑
├─────────────┤                      │
│ &s (ref) ───┼──────────────────────┘
└─────────────┘
```

### 2.8 Mutable References

By default, references are immutable. To modify borrowed data, use `&mut`:

```rust
fn main() {
    let mut s = String::from("hello");
    
    add_world(&mut s);
    
    println!("{}", s);  // "hello, world!"
}

fn add_world(s: &mut String) {
    s.push_str(", world!");
}
```

**The Borrowing Rules:**

1. **You can have EITHER:**
   - One mutable reference (`&mut T`), OR
   - Any number of immutable references (`&T`)
   
2. **References must always be valid** (no dangling references)

```rust
fn main() {
    let mut s = String::from("hello");
    
    let r1 = &s;     // OK: first immutable borrow
    let r2 = &s;     // OK: second immutable borrow
    // let r3 = &mut s;  // ERROR: cannot borrow as mutable while immutable borrows exist
    
    println!("{} and {}", r1, r2);
    // r1 and r2 are no longer used after this point
    
    let r3 = &mut s;  // OK: immutable borrows are done
    r3.push_str("!");
}
```

**Why this rule?** It prevents data races at compile time!

### 2.9 Dangling References

Rust prevents dangling references (references to freed memory):

```rust
fn main() {
    let reference = dangle();  // ERROR!
}

fn dangle() -> &String {
    let s = String::from("hello");
    &s  // ERROR: s is dropped when function ends, reference would be invalid
}

// Fix: return the owned value instead
fn no_dangle() -> String {
    let s = String::from("hello");
    s  // Ownership is moved out
}
```

### 2.10 Lifetimes

Sometimes Rust needs help understanding how long references are valid:

```rust
// This won't compile - Rust doesn't know which input lifetime to use for output
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() { x } else { y }
}

// Fix: add lifetime annotations
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}
```

**What `'a` means:** "The returned reference will be valid as long as both input references are valid."

**Lifetime Elision Rules:**
Rust often infers lifetimes automatically. The rules are:
1. Each input reference gets its own lifetime
2. If there's exactly one input lifetime, it's used for output
3. If there's `&self` or `&mut self`, that lifetime is used for output

**You'll rarely need explicit lifetimes**, but understanding them helps debug errors.

### 2.11 Slices: References to Parts of Data

A slice is a reference to a contiguous sequence of elements:

```rust
fn main() {
    let s = String::from("hello world");
    
    let hello = &s[0..5];   // Slice of first 5 bytes
    let world = &s[6..11];  // Slice from index 6 to 10
    
    println!("{} {}", hello, world);
    
    // Shorthand
    let hello = &s[..5];    // Same as [0..5]
    let world = &s[6..];    // Same as [6..len]
    let whole = &s[..];     // Same as [0..len]
}
```

**String slices have type `&str`** (not `&String`). This is why string literals are `&str`:
```rust
let literal: &str = "hello";  // Points directly into the binary
```

---

## 🧠 Theory: Why Ownership Enables Optimization

Ownership isn't just about safety — it enables powerful compiler optimizations.

**Aliasing and Mutation:**
In most languages, the compiler can't know if two pointers refer to the same memory:
```c
void add_to_both(int* a, int* b) {
    *a += 1;
    *b += 1;
    // If a and b point to the same memory, this adds 2
    // If they're different, this adds 1 to each
    // Compiler can't optimize without knowing!
}
```

**In Rust:** The borrow checker guarantees no aliasing with mutation:
- If you have `&mut T`, you know there are no other references
- The compiler can optimize freely!

This is called **"Pointer Aliasing"** and it's a key reason Rust can match C performance while being safe.

---

## ✅ Summary

| Concept | Syntax | Notes |
|---------|--------|-------|
| Ownership | `let s = String::from("x")` | One owner at a time |
| Move | `let s2 = s1;` | s1 invalidated |
| Clone | `let s2 = s1.clone();` | Deep copy |
| Immutable borrow | `&s` | Multiple allowed |
| Mutable borrow | `&mut s` | Only one allowed |
| Lifetime | `'a` | How long references are valid |
| Slice | `&s[0..5]` | Reference to contiguous elements |

---

## 📚 Further Reading

- [The Rust Programming Language, Chapter 4](https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html)
- [Rust By Example: Ownership](https://doc.rust-lang.org/rust-by-example/scope.html)
- [The Rustonomicon](https://doc.rust-lang.org/nomicon/) — For deep unsafe Rust understanding

---

**Next:** [Lecture 3: Structs, Enums & Pattern Matching](../lecture-03/README.md)

**Exercises:** [Exercise Sheet 2](exercises.md)
