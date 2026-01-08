# Exercise Sheet 4: Language Design & Lexical Analysis

These exercises combine practical language design with theoretical foundations.

---

## Exercise 4.1: Token Set Design 📋

**Objective:** Design a complete token set for Rusty.

Review the token kinds listed in the lecture and extend them to handle:

1. **Comments:**
   - Line comments: `// comment`
   - Block comments: `/* comment */`
   - Nested block comments: `/* outer /* inner */ */`

2. **Additional literals:**
   - Character literals: `'a'`, `'\n'`, `'\x41'`
   - Binary integers: `0b1010`
   - Octal integers: `0o755`
   - Hexadecimal integers: `0xDEADBEEF`

3. **Additional operators:**
   - Compound assignment: `+=`, `-=`, `*=`, `/=`
   - Range: `..`, `..=`
   - Question mark: `?`

Write out the complete `TokenKind` enum with all variants:

```rust
#[derive(Debug, Clone, PartialEq)]
pub enum TokenKind {
    // Literals
    // TODO: List all literal variants
    
    // Keywords
    // TODO: List all keywords
    
    // Operators
    // TODO: List all operators
    
    // Delimiters
    // TODO: List all delimiters
    
    // Punctuation
    // TODO: List all punctuation
    
    // Special
    Eof,
    Error(String),
}
```

---

## Exercise 4.2: Regular Expressions 🔤

**Objective:** Write regular expressions for token patterns.

Write a regex for each of the following:

1. **Identifier:** Starts with letter or underscore, followed by letters, digits, or underscores
   ```
   [a-zA-Z_][a-zA-Z0-9_]*
   ```

2. **Decimal integer:** Optional sign, one or more digits, optional underscores between digits
   ```
   Your answer: _____________
   Example matches: 42, -17, 1_000_000
   Example non-matches: _42, 42_, 1__2
   ```

3. **Hexadecimal integer:** `0x` or `0X` prefix, hex digits with optional underscores
   ```
   Your answer: _____________
   Example matches: 0xFF, 0x1a2b, 0X_DEAD_BEEF
   ```

4. **Float:** Integer part, decimal point, fractional part, optional exponent
   ```
   Your answer: _____________
   Example matches: 3.14, 0.5, 1e10, 6.022e23, 1.5e-4
   ```

5. **String literal:** Double quotes, any chars except unescaped quote/newline, escape sequences
   ```
   Your answer: _____________
   Example matches: "", "hello", "say \"hi\"", "line1\nline2"
   ```

6. **Line comment:** `//` followed by any characters until newline
   ```
   Your answer: _____________
   ```

7. **Block comment (non-nested):** `/*` ... `*/`
   ```
   Your answer: _____________
   Hint: Be careful not to match */ inside!
   ```

---

## Exercise 4.3: Grammar in EBNF 📝

**Objective:** Write formal grammar rules for Rusty constructs.

Write EBNF rules for:

### Part A: Expressions with Precedence

Write grammar rules that correctly handle operator precedence:
- Lowest: `||`
- Then: `&&`
- Then: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Then: `+`, `-`
- Then: `*`, `/`, `%`
- Then: unary `-`, `!`
- Then: function calls, field access
- Highest: literals, identifiers, parenthesized expressions

```ebnf
expression     = ??? ;
or_expr        = ??? ;
and_expr       = ??? ;
equality_expr  = ??? ;
comparison_expr = ??? ;
additive_expr  = ??? ;
multiplicative_expr = ??? ;
unary_expr     = ??? ;
postfix_expr   = ??? ;
primary_expr   = ??? ;
```

### Part B: Match Expression

```ebnf
match_expr     = ??? ;
match_arm      = ??? ;
pattern        = ??? ;
```

Example to parse:
```rust
match x {
    0 => "zero",
    1 | 2 => "small",
    n if n < 10 => "medium",
    _ => "large",
}
```

### Part C: Struct Definition and Literal

```ebnf
struct_def     = ??? ;
field_def      = ??? ;
struct_literal = ??? ;
field_init     = ??? ;
```

Example:
```rust
struct Point { x: i64, y: i64 }

let p = Point { x: 10, y: 20 };
```

---

## Exercise 4.4: DFA Construction 🔧

**Objective:** Build a DFA by hand for simple patterns.

### Part A: DFA for `if`, `in`, `impl`

Construct a DFA that recognizes exactly these three keywords. All other inputs should be rejected.

Draw the state diagram with:
- States labeled (q0, q1, etc.)
- Transitions labeled with characters
- Accept states double-circled
- A clear start state

### Part B: DFA for Decimal Integers

Construct a DFA that recognizes valid decimal integers:
- One or more digits
- Optional leading `-`
- No leading zeros (except for just `0`)
- Valid: `0`, `42`, `-17`, `100`
- Invalid: `00`, `007`, `-`, `42a`

### Part C: Combine DFAs

Explain (in words) how you would combine the DFAs from Parts A and B to create a lexer that recognizes both keywords AND integers.

---

## 📊 Theory Exercise 4.5: Pumping Lemma Applications

**Objective:** Use the pumping lemma to prove languages are not regular.

### Part A: Prove that L₁ = {aⁿbⁿcⁿ | n ≥ 0} is not regular.

1. Assume L₁ is regular with pumping length p
2. Choose a string s ∈ L₁ with |s| ≥ p
3. Show that for any decomposition s = xyz satisfying the conditions, some pumped string is not in L₁
4. Conclude by contradiction

### Part B: Prove that L₂ = {ww | w ∈ {a,b}*} is not regular.

This is the language of all strings that are a word repeated twice.
Examples: `""`, `abab`, `abaaba`, `bbaabbaa`

### Part C: Prove that L₃ = {0ⁿ1ⁿ | n ≥ 1} is not regular.

This is the language of n zeros followed by n ones, for n ≥ 1.

### Part D: Why This Matters

In 2-3 sentences, explain why it's important that some languages are not regular, from the perspective of compiler design.

---

## 📊 Theory Exercise 4.6: Regular Grammar Equivalence

**Objective:** Understand the equivalence of regular language representations.

A **regular grammar** is a grammar where every production has the form:
- `A → aB` (right-linear)
- `A → a`
- `A → ε`

Where `A`, `B` are non-terminals and `a` is a terminal.

### Part A: Convert Regex to Grammar

Convert this regex to a regular grammar:
```
(a|b)*abb
```

### Part B: Convert DFA to Grammar

Given this DFA:
```
States: {q0, q1, q2}
Alphabet: {0, 1}
Start: q0
Accept: {q2}
Transitions:
  q0 --0--> q1
  q0 --1--> q0
  q1 --0--> q1
  q1 --1--> q2
  q2 --0--> q1
  q2 --1--> q0
```

Write an equivalent regular grammar.

### Part C: What Language?

What language does the DFA in Part B recognize? Describe it in English.

---

## Exercise 4.7: Lexer Edge Cases 🎯

**Objective:** Identify tricky tokenization scenarios.

For each input, show the expected token sequence:

1. `x-y`
2. `x - y`
3. `x--y`
4. `x->y`
5. `42.0`
6. `42.foo`
7. `42..`
8. `42..=50`
9. `"hello\"world"`
10. `"line1\nline2"`
11. `///doc comment`
12. `/*/* /* */`
13. `let letlet = 1;`
14. `_`
15. `__reserved__`

---

## Exercise 4.8: Test Cases for Lexer 🧪

**Objective:** Write comprehensive test cases.

Write test cases (input → expected tokens) for:

### Valid Inputs
```rust
#[test]
fn test_simple_function() {
    let input = "fn add(a: i64, b: i64) -> i64 { a + b }";
    let expected = vec![
        // TODO: list expected tokens
    ];
    assert_eq!(tokenize(input), expected);
}

#[test]
fn test_struct_definition() {
    let input = "struct Point { x: i64, y: i64 }";
    // TODO
}

#[test]
fn test_match_expression() {
    let input = r#"match x { 0 => "zero", _ => "other" }"#;
    // TODO
}

#[test]
fn test_nested_comments() {
    let input = "/* outer /* inner */ still comment */";
    // TODO: What tokens should this produce?
}
```

### Error Cases
```rust
#[test]
fn test_unterminated_string() {
    let input = r#""hello"#;
    // TODO: What error?
}

#[test]
fn test_invalid_escape() {
    let input = r#""hello\z""#;
    // TODO: What error?
}

#[test]
fn test_invalid_number() {
    let input = "123abc";
    // TODO: What should happen?
}
```

---

## 🏆 Bonus Challenge: Lexer Generator

Design a simple lexer generator. Given a set of (regex, token_type) pairs, generate Rust code for a lexer.

```rust
// Input specification
let spec = vec![
    (r"[a-zA-Z_][a-zA-Z0-9_]*", "Identifier"),
    (r"[0-9]+", "Integer"),
    (r"\+", "Plus"),
    (r"-", "Minus"),
    // ...
];

// Output: Rust code for a lexer function
fn generate_lexer(spec: &[(& str, &str)]) -> String {
    // TODO: Generate Rust code that implements the lexer
}
```

This is how tools like `lex` and `flex` work!

---

## 📝 Submission Checklist

- [ ] Exercise 4.1: Complete TokenKind enum
- [ ] Exercise 4.2: All regular expressions written
- [ ] Exercise 4.3: EBNF grammar complete
- [ ] Exercise 4.4: DFA diagrams drawn
- [ ] Exercise 4.5: Pumping lemma proofs complete
- [ ] Exercise 4.6: Grammar equivalence exercises done
- [ ] Exercise 4.7: Edge cases analyzed
- [ ] Exercise 4.8: Test cases written
- [ ] Bonus: Lexer generator designed (optional)

---

**Solutions:** [solutions/](solutions/)

**Next Lecture:** [Lecture 5: Building a Lexer in Rust](../lecture-05/README.md)
