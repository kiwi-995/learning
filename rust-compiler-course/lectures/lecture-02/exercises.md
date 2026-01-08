# Exercise Sheet 2: Ownership, Borrowing & Memory Safety

These exercises will help you internalize Rust's ownership model. Many exercises involve fixing broken code — this is intentional! Learning to read and fix ownership errors is a critical skill.

---

## Exercise 2.1: Ownership Errors 🔧

**Objective:** Identify and fix ownership errors.

Each snippet has an ownership bug. Fix it!

### Part A: Double Use After Move
```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1;
    println!("{}", s1);  // ERROR: value borrowed after move
}

// TODO: Fix by cloning OR by using references
```

### Part B: Function Takes Ownership
```rust
fn print_string(s: String) {
    println!("{}", s);
}

fn main() {
    let s = String::from("hello");
    print_string(s);
    print_string(s);  // ERROR: use after move
}

// TODO: Fix by changing the function signature to borrow
```

### Part C: Returning References to Local Data
```rust
fn create_greeting() -> &String {
    let greeting = String::from("Hello!");
    &greeting  // ERROR: returns a reference to data that will be dropped
}

// TODO: Fix by returning owned data
```

### Part D: Mixing Mutable and Immutable Borrows
```rust
fn main() {
    let mut v = vec![1, 2, 3];
    let first = &v[0];
    v.push(4);  // ERROR: cannot borrow as mutable while immutable borrow exists
    println!("{}", first);
}

// TODO: Fix by reordering operations
```

---

## Exercise 2.2: Implementing a Stack 📚

**Objective:** Practice ownership with a simple data structure.

Implement a stack that owns its elements:

```rust
struct Stack {
    items: Vec<String>,
}

impl Stack {
    /// Creates a new empty stack
    fn new() -> Stack {
        todo!()
    }
    
    /// Pushes an item onto the stack (takes ownership)
    fn push(&mut self, item: String) {
        todo!()
    }
    
    /// Pops an item from the stack (returns ownership)
    fn pop(&mut self) -> Option<String> {
        todo!()
    }
    
    /// Peeks at the top item without removing it (returns reference)
    fn peek(&self) -> Option<&String> {
        todo!()
    }
    
    /// Returns the number of items
    fn len(&self) -> usize {
        todo!()
    }
    
    /// Returns true if the stack is empty
    fn is_empty(&self) -> bool {
        todo!()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_push_pop() {
        let mut stack = Stack::new();
        stack.push(String::from("first"));
        stack.push(String::from("second"));
        
        assert_eq!(stack.pop(), Some(String::from("second")));
        assert_eq!(stack.pop(), Some(String::from("first")));
        assert_eq!(stack.pop(), None);
    }
    
    #[test]
    fn test_peek() {
        let mut stack = Stack::new();
        assert_eq!(stack.peek(), None);
        
        stack.push(String::from("hello"));
        assert_eq!(stack.peek(), Some(&String::from("hello")));
        
        // peek shouldn't remove the item
        assert_eq!(stack.len(), 1);
    }
}
```

---

## Exercise 2.3: String Parser (Borrowing Practice) 📝

**Objective:** Practice borrowing with string processing.

This exercise previews what we'll need for the lexer!

```rust
/// Counts how many times `needle` appears in `haystack`
/// Both parameters are borrowed (we don't need ownership)
fn count_occurrences(haystack: &str, needle: &str) -> usize {
    todo!()
    // Hint: Use haystack.matches(needle).count()
}

/// Finds all words in a string (separated by whitespace)
/// Returns a vector of string slices (references into the original string)
fn find_words(text: &str) -> Vec<&str> {
    todo!()
    // Hint: Use text.split_whitespace().collect()
}

/// Extracts the first word from a string
/// Returns None if the string is empty
fn first_word(text: &str) -> Option<&str> {
    todo!()
}

/// Checks if `text` starts with `prefix` (case-insensitive)
fn starts_with_ignore_case(text: &str, prefix: &str) -> bool {
    todo!()
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_count_occurrences() {
        assert_eq!(count_occurrences("hello world", "o"), 2);
        assert_eq!(count_occurrences("hello world", "ll"), 1);
        assert_eq!(count_occurrences("hello world", "x"), 0);
    }
    
    #[test]
    fn test_find_words() {
        let text = "hello world foo bar";
        let words = find_words(text);
        assert_eq!(words, vec!["hello", "world", "foo", "bar"]);
    }
    
    #[test]
    fn test_first_word() {
        assert_eq!(first_word("hello world"), Some("hello"));
        assert_eq!(first_word("   "), None);
        assert_eq!(first_word("single"), Some("single"));
    }
    
    #[test]
    fn test_starts_with_ignore_case() {
        assert!(starts_with_ignore_case("Hello World", "hello"));
        assert!(starts_with_ignore_case("RUST", "ru"));
        assert!(!starts_with_ignore_case("Rust", "python"));
    }
}
```

---

## Exercise 2.4: Token Structure (Compiler Preview) 🎯

**Objective:** Create a token structure that borrows from source code.

In a real compiler, tokens often reference the original source code to avoid copying:

```rust
/// A token that references a portion of the source code
#[derive(Debug, PartialEq)]
struct Token<'a> {
    kind: TokenKind,
    text: &'a str,  // Reference to the original source code
    line: usize,
    column: usize,
}

#[derive(Debug, PartialEq, Clone, Copy)]
enum TokenKind {
    Number,
    Plus,
    Minus,
    Star,
    Slash,
    LeftParen,
    RightParen,
    Unknown,
}

impl<'a> Token<'a> {
    fn new(kind: TokenKind, text: &'a str, line: usize, column: usize) -> Token<'a> {
        Token { kind, text, line, column }
    }
}

/// A very simple tokenizer for arithmetic expressions
/// Returns a vector of tokens that reference the original source
fn tokenize(source: &str) -> Vec<Token> {
    let mut tokens = Vec::new();
    let mut line = 1;
    let mut column = 1;
    let mut chars = source.char_indices().peekable();
    
    while let Some((i, c)) = chars.next() {
        let token = match c {
            ' ' | '\t' => {
                column += 1;
                continue;
            }
            '\n' => {
                line += 1;
                column = 1;
                continue;
            }
            '+' => Token::new(TokenKind::Plus, &source[i..i+1], line, column),
            '-' => Token::new(TokenKind::Minus, &source[i..i+1], line, column),
            '*' => Token::new(TokenKind::Star, &source[i..i+1], line, column),
            '/' => Token::new(TokenKind::Slash, &source[i..i+1], line, column),
            '(' => Token::new(TokenKind::LeftParen, &source[i..i+1], line, column),
            ')' => Token::new(TokenKind::RightParen, &source[i..i+1], line, column),
            '0'..='9' => {
                // TODO: Handle multi-digit numbers
                // Consume all consecutive digits
                // Create a Number token with the full number text
                todo!()
            }
            _ => Token::new(TokenKind::Unknown, &source[i..i+1], line, column),
        };
        
        column += token.text.len();
        tokens.push(token);
    }
    
    tokens
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_simple_expression() {
        let source = "1 + 2";
        let tokens = tokenize(source);
        
        assert_eq!(tokens.len(), 3);
        assert_eq!(tokens[0].kind, TokenKind::Number);
        assert_eq!(tokens[0].text, "1");
        assert_eq!(tokens[1].kind, TokenKind::Plus);
        assert_eq!(tokens[2].kind, TokenKind::Number);
        assert_eq!(tokens[2].text, "2");
    }
    
    #[test]
    fn test_multi_digit_numbers() {
        let source = "123 + 456";
        let tokens = tokenize(source);
        
        assert_eq!(tokens[0].text, "123");
        assert_eq!(tokens[2].text, "456");
    }
    
    #[test]
    fn test_complex_expression() {
        let source = "(1 + 2) * 3";
        let tokens = tokenize(source);
        
        assert_eq!(tokens[0].kind, TokenKind::LeftParen);
        assert_eq!(tokens[4].kind, TokenKind::RightParen);
        assert_eq!(tokens[5].kind, TokenKind::Star);
    }
}
```

**Questions to answer:**
1. Why does `Token` need a lifetime parameter `'a`?
2. What happens to the tokens if we drop the original `source` string?
3. Could we make `Token` own its text instead? What would be the tradeoff?

---

## Exercise 2.5: Lifetime Annotations 🧬

**Objective:** Practice explicit lifetime annotations.

### Part A: Fix the Compile Error
```rust
// This doesn't compile. Add lifetime annotations to fix it.
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let s1 = String::from("hello");
    let result;
    {
        let s2 = String::from("world!");
        result = longest(&s1, &s2);
    }
    println!("{}", result);  // Should this work? Why or why not?
}
```

### Part B: Struct with References
```rust
// Add lifetime annotations to make this compile
struct Excerpt {
    text: &str,
}

impl Excerpt {
    fn new(text: &str) -> Excerpt {
        Excerpt { text }
    }
    
    fn text(&self) -> &str {
        self.text
    }
}

fn main() {
    let novel = String::from("Call me Ishmael. Some years ago...");
    let excerpt = Excerpt::new(&novel[0..16]);
    println!("Excerpt: {}", excerpt.text());
}
```

### Part C: Multiple Lifetimes
```rust
// Sometimes you need different lifetimes for different references
// Fix this function to compile

struct Context<'a> {
    source: &'a str,
}

struct Lexer<'a> {
    context: &'a Context<'a>,
    position: usize,
}

impl<'a> Lexer<'a> {
    fn new(context: &Context) -> Lexer {
        Lexer { context, position: 0 }
    }
    
    fn current_char(&self) -> Option<char> {
        self.context.source.chars().nth(self.position)
    }
}
```

---

## 📊 Theory Exercise 2.6: Ownership and Linearity

**Objective:** Connect ownership to linear type theory.

Rust's ownership system is related to *linear types* in programming language theory.

**Definition:** A linear type is a type whose values must be used exactly once.

**Questions:**

1. In Rust, after `let s2 = s1;` (where `s1: String`), why can't we use `s1`? How does this relate to linear types?

2. The `Copy` trait allows values to be used multiple times. In type theory terms, `Copy` types are more like *non-linear* types. Explain the difference.

3. Consider this Rust code:
   ```rust
   fn take_and_give_back(s: String) -> String {
       s  // Just return it unchanged
   }
   ```
   Why might a Rust programmer write a function like this? (Hint: think about ownership transfer)

4. **Challenge:** In some languages with linear types, you must explicitly "consume" or "free" a value. In Rust, this is automatic. What mechanism in Rust handles this? (Hint: it's related to scope)

---

## 📊 Theory Exercise 2.7: Memory Safety Proofs (Preview)

**Objective:** Understand how the borrow checker enforces safety.

Consider these two invariants that Rust maintains:

**Invariant 1 (No Use After Free):** A reference `&T` or `&mut T` is always valid — the data it points to has not been freed.

**Invariant 2 (No Data Races):** At any point in time, either:
- There are any number of immutable references `&T`, OR
- There is exactly one mutable reference `&mut T`
(But never both simultaneously)

**Questions:**

1. Show how the following code would violate Invariant 1, and explain why Rust rejects it:
   ```rust
   fn bad() -> &String {
       let s = String::from("hello");
       &s
   }
   ```

2. Show how the following code would violate Invariant 2, and explain what bug it prevents:
   ```rust
   fn main() {
       let mut v = vec![1, 2, 3];
       let first = &v[0];
       v.push(4);  // This could reallocate the vector!
       println!("{}", first);  // first might point to freed memory
   }
   ```

3. **The "Aliasing XOR Mutation" Principle:** Explain in your own words why having either multiple readers OR one writer (but not both) prevents data races.

---

## 🏆 Bonus Challenge: Arena Allocator

An arena allocator is a common pattern in compilers. It allocates many objects with the same lifetime efficiently.

```rust
use std::cell::RefCell;

/// A simple arena that owns strings
struct StringArena {
    strings: RefCell<Vec<String>>,
}

impl StringArena {
    fn new() -> StringArena {
        StringArena {
            strings: RefCell::new(Vec::new()),
        }
    }
    
    /// Allocates a string in the arena and returns a reference to it
    /// The reference is valid as long as the arena exists
    fn alloc(&self, s: String) -> &str {
        // TODO: Implement this
        // Hint: This requires unsafe code or a different approach
        // For now, you can use a simpler version that returns String
        todo!()
    }
}
```

**Question:** Why is this pattern useful for compilers? (Think about token and AST lifetimes)

---

## 📝 Submission Checklist

- [ ] Exercise 2.1: All ownership errors fixed
- [ ] Exercise 2.2: Stack implementation passes all tests
- [ ] Exercise 2.3: String parser functions work
- [ ] Exercise 2.4: Token structure compiles and tests pass
- [ ] Exercise 2.5: All lifetime annotations added correctly
- [ ] Exercise 2.6: Theory questions answered
- [ ] Exercise 2.7: Memory safety analysis complete
- [ ] Bonus: Arena allocator attempted (optional)

---

**Solutions:** [solutions/](solutions/)

**Next Lecture:** [Lecture 3: Structs, Enums & Pattern Matching](../lecture-03/README.md)
