# Exercise Sheet 5: Building a Lexer

Implement and test the complete Rusty lexer.

---

## Exercise 5.1: Complete the Lexer

Using the code from the lecture as a starting point, implement:

1. `scan_char()` - Character literal scanning with escapes
2. `scan_binary()` - Binary literal scanning (0b1010)
3. `scan_octal()` - Octal literal scanning (0o755)

---

## Exercise 5.2: Comprehensive Tests

Write tests for:
- All operators and punctuation
- All keywords
- Edge cases (empty input, only whitespace, only comments)
- Error cases (unterminated strings, invalid escapes)

---

## Exercise 5.3: Error Recovery

Modify the lexer to continue after errors and collect multiple errors.

---

## Exercise 5.4: REPL

Build a simple REPL that tokenizes input lines:

```rust
fn main() {
    loop {
        print!("> ");
        let mut input = String::new();
        std::io::stdin().read_line(&mut input).unwrap();
        
        let lexer = Lexer::new(&input);
        match lexer.tokenize() {
            Ok(tokens) => {
                for token in tokens {
                    println!("{:?}", token);
                }
            }
            Err(errors) => {
                for err in errors {
                    eprintln!("Error: {}", err.message);
                }
            }
        }
    }
}
```

---

**Solutions:** [solutions/](solutions/)

**Next Lecture:** [Lecture 6: Parsing & Building ASTs](../lecture-06/README.md)
