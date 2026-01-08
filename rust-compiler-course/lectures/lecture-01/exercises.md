# Exercise Sheet 1: Getting Started with Rust

Complete these exercises to practice the concepts from Lecture 1. Each exercise builds skills you'll use in the compiler.

---

## Exercise 1.1: Environment Setup ✅

**Objective:** Verify your Rust installation.

1. Run the following commands and record the output:
```bash
rustc --version
cargo --version
rustup --version
```

2. Create a new project:
```bash
cargo new exercise_1_1
cd exercise_1_1
cargo run
```

**Expected output:** `Hello, world!`

---

## Exercise 1.2: Calculator Warmup 🧮

**Objective:** Practice variables, functions, and expressions.

Create a simple calculator that performs basic arithmetic:

```rust
// src/main.rs

fn add(a: i64, b: i64) -> i64 {
    // TODO: Implement
    todo!()
}

fn subtract(a: i64, b: i64) -> i64 {
    // TODO: Implement
    todo!()
}

fn multiply(a: i64, b: i64) -> i64 {
    // TODO: Implement
    todo!()
}

fn divide(a: i64, b: i64) -> i64 {
    // TODO: Implement
    // Note: Integer division (rounds toward zero)
    todo!()
}

fn main() {
    println!("10 + 5 = {}", add(10, 5));
    println!("10 - 5 = {}", subtract(10, 5));
    println!("10 * 5 = {}", multiply(10, 5));
    println!("10 / 5 = {}", divide(10, 5));
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
        assert_eq!(add(-1, 1), 0);
        assert_eq!(add(0, 0), 0);
    }
    
    #[test]
    fn test_subtract() {
        assert_eq!(subtract(5, 3), 2);
        assert_eq!(subtract(3, 5), -2);
    }
    
    #[test]
    fn test_multiply() {
        assert_eq!(multiply(3, 4), 12);
        assert_eq!(multiply(-2, 3), -6);
    }
    
    #[test]
    fn test_divide() {
        assert_eq!(divide(10, 3), 3);
        assert_eq!(divide(-7, 2), -3);
    }
}
```

Run tests with: `cargo test`

---

## Exercise 1.3: FizzBuzz Variations 🎲

**Objective:** Practice control flow with if/else.

### Part A: Classic FizzBuzz
Print numbers 1-100, but:
- For multiples of 3, print "Fizz"
- For multiples of 5, print "Buzz"  
- For multiples of both, print "FizzBuzz"

```rust
fn fizzbuzz(n: i32) -> String {
    // TODO: Implement
    todo!()
}

fn main() {
    for i in 1..=100 {
        println!("{}", fizzbuzz(i));
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_fizzbuzz() {
        assert_eq!(fizzbuzz(1), "1");
        assert_eq!(fizzbuzz(3), "Fizz");
        assert_eq!(fizzbuzz(5), "Buzz");
        assert_eq!(fizzbuzz(15), "FizzBuzz");
        assert_eq!(fizzbuzz(7), "7");
    }
}
```

### Part B: Generalized FizzBuzz 🌟
Modify the function to take the two divisors as parameters:

```rust
fn generalized_fizzbuzz(n: i32, fizz_div: i32, buzz_div: i32) -> String {
    // TODO: Implement
    todo!()
}
```

---

## Exercise 1.4: Number Guessing Game 🎮

**Objective:** Practice loops, conditionals, and user input.

Build an interactive number guessing game:

```rust
use std::io;

fn main() {
    let secret = 42;  // In a real game, this would be random
    let mut attempts = 0;
    
    println!("Guess the number (1-100)!");
    
    loop {
        println!("Enter your guess:");
        
        let mut input = String::new();
        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line");
        
        // Parse the input to a number
        let guess: i32 = match input.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Please enter a valid number!");
                continue;
            }
        };
        
        attempts += 1;
        
        // TODO: Compare guess with secret
        // - If too low, print "Too low!"
        // - If too high, print "Too high!"
        // - If correct, print "You got it in {attempts} attempts!" and break
        todo!()
    }
}
```

**Bonus:** Use random numbers:
```rust
// Add to Cargo.toml: rand = "0.8"
use rand::Rng;

let secret = rand::thread_rng().gen_range(1..=100);
```

---

## Exercise 1.5: Temperature Converter 🌡️

**Objective:** Practice functions and expressions.

```rust
fn celsius_to_fahrenheit(c: f64) -> f64 {
    // Formula: F = C * 9/5 + 32
    todo!()
}

fn fahrenheit_to_celsius(f: f64) -> f64 {
    // Formula: C = (F - 32) * 5/9
    todo!()
}

fn main() {
    let temps_c = [0.0, 20.0, 37.0, 100.0];
    
    println!("Celsius → Fahrenheit:");
    for c in temps_c {
        println!("  {:.1}°C = {:.1}°F", c, celsius_to_fahrenheit(c));
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_freezing() {
        assert!((celsius_to_fahrenheit(0.0) - 32.0).abs() < 0.01);
    }
    
    #[test]
    fn test_boiling() {
        assert!((celsius_to_fahrenheit(100.0) - 212.0).abs() < 0.01);
    }
    
    #[test]
    fn test_body_temp() {
        assert!((celsius_to_fahrenheit(37.0) - 98.6).abs() < 0.1);
    }
    
    #[test]
    fn test_roundtrip() {
        let original = 25.0;
        let result = fahrenheit_to_celsius(celsius_to_fahrenheit(original));
        assert!((result - original).abs() < 0.01);
    }
}
```

---

## Exercise 1.6: Fibonacci Sequence 🐚

**Objective:** Practice loops and mutable state.

### Part A: Iterative Fibonacci
```rust
fn fibonacci_iterative(n: u32) -> u64 {
    // Return the nth Fibonacci number (0-indexed)
    // F(0) = 0, F(1) = 1, F(n) = F(n-1) + F(n-2)
    todo!()
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_fibonacci() {
        assert_eq!(fibonacci_iterative(0), 0);
        assert_eq!(fibonacci_iterative(1), 1);
        assert_eq!(fibonacci_iterative(2), 1);
        assert_eq!(fibonacci_iterative(10), 55);
        assert_eq!(fibonacci_iterative(20), 6765);
    }
}
```

### Part B: Recursive Fibonacci (preview of future topics)
```rust
fn fibonacci_recursive(n: u32) -> u64 {
    // Same result, but using recursion
    // Warning: This is inefficient! We'll optimize it later.
    todo!()
}
```

**Question:** Why is the recursive version slow for large `n`? (We'll revisit this when we implement recursion in our compiler!)

---

## Exercise 1.7: Prime Numbers 🔢

**Objective:** Practice nested loops and early returns.

```rust
fn is_prime(n: u64) -> bool {
    // Return true if n is prime, false otherwise
    // Hint: Check divisibility from 2 to sqrt(n)
    todo!()
}

fn nth_prime(n: usize) -> u64 {
    // Return the nth prime (1-indexed)
    // nth_prime(1) = 2, nth_prime(2) = 3, nth_prime(3) = 5, ...
    todo!()
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_is_prime() {
        assert!(!is_prime(0));
        assert!(!is_prime(1));
        assert!(is_prime(2));
        assert!(is_prime(3));
        assert!(!is_prime(4));
        assert!(is_prime(5));
        assert!(is_prime(97));
        assert!(!is_prime(100));
    }
    
    #[test]
    fn test_nth_prime() {
        assert_eq!(nth_prime(1), 2);
        assert_eq!(nth_prime(6), 13);
        assert_eq!(nth_prime(100), 541);
    }
}
```

---

## 📊 Theory Exercise 1.8: Language Cardinality

**Objective:** Get comfortable thinking about languages as sets.

Consider the alphabet Σ = {a, b}.

1. How many strings of length exactly 3 are there over Σ?

2. How many strings of length ≤ 3 are there over Σ?

3. Let L = {strings where the number of 'a's equals the number of 'b's}
   - List all strings in L of length ≤ 4
   - Is L finite or infinite?
   - Is `"aabb"` in L? Is `"abba"` in L?

4. **Challenge:** Consider the language of all valid Rust variable names (identifiers). 
   - Describe this language informally.
   - Is it finite or infinite?
   - Could you describe it with a simple pattern?

---

## 🏆 Bonus Challenge: Collatz Conjecture

The Collatz sequence starting from `n`:
- If n is even: next = n / 2
- If n is odd: next = 3n + 1
- Stop when you reach 1

```rust
fn collatz_length(n: u64) -> u32 {
    // Return the number of steps to reach 1
    // collatz_length(1) = 0
    // collatz_length(2) = 1  (2 → 1)
    // collatz_length(3) = 7  (3 → 10 → 5 → 16 → 8 → 4 → 2 → 1)
    todo!()
}

fn longest_collatz(max: u64) -> (u64, u32) {
    // Return (starting_number, length) for the longest sequence
    // among all starting values from 1 to max
    todo!()
}
```

**Fun fact:** It's an open problem whether the Collatz sequence reaches 1 for all starting values!

---

## 📝 Submission Checklist

- [ ] Exercise 1.1: Environment verified
- [ ] Exercise 1.2: Calculator passes all tests
- [ ] Exercise 1.3: FizzBuzz works correctly
- [ ] Exercise 1.4: Guessing game is playable
- [ ] Exercise 1.5: Temperature converter passes all tests
- [ ] Exercise 1.6: Both Fibonacci implementations work
- [ ] Exercise 1.7: Prime functions pass all tests
- [ ] Exercise 1.8: Theory questions answered
- [ ] Bonus: Collatz functions work (optional)

---

**Solutions:** [solutions/](solutions/)

**Next Lecture:** [Lecture 2: Ownership, Borrowing & Memory Safety](../lecture-02/README.md)
