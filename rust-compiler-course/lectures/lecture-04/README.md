# Lecture 4: Language Design & Lexical Analysis

> *"A language that doesn't affect the way you think about programming is not worth knowing."* — Alan Perlis

Welcome to the compiler portion of the course! In this lecture, we design our target language "Rusty" and explore the theory of lexical analysis — the first stage of compilation.

## 🎯 Learning Objectives

By the end of this lecture, you will be able to:
- Explain the stages of a compiler pipeline
- Design a simple programming language grammar
- Understand formal language theory basics (regular languages, DFAs)
- Map tokens to Rust enums
- Identify edge cases in tokenization

## 📖 Topics

### 4.1 The Compiler Pipeline

A compiler transforms source code into executable code through several stages:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                          COMPILER PIPELINE                               │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│  Source Code                                                             │
│       │                                                                  │
│       ▼                                                                  │
│  ┌─────────┐    ┌─────────┐    ┌──────────┐    ┌────────────┐           │
│  │  Lexer  │───▶│ Parser  │───▶│ Semantic │───▶│    Code    │           │
│  │         │    │         │    │ Analysis │    │ Generation │           │
│  └─────────┘    └─────────┘    └──────────┘    └────────────┘           │
│       │              │              │                │                   │
│       ▼              ▼              ▼                ▼                   │
│   Tokens           AST        Validated AST    Bytecode/Machine Code    │
│                                                                          │
└─────────────────────────────────────────────────────────────────────────┘
```

| Stage | Input | Output | Purpose |
|-------|-------|--------|---------|
| **Lexer** | Source text | Tokens | Break text into meaningful units |
| **Parser** | Tokens | AST | Build tree structure, check syntax |
| **Semantic Analysis** | AST | Validated AST | Check types, resolve names |
| **Code Generation** | Validated AST | Bytecode/IR | Produce executable code |
| **Optimization** | IR | Optimized IR | Improve performance (optional) |

### 4.2 Our Target Language: Rusty

Let's design "Rusty" — a simplified, Rust-inspired language that's feasible to implement but powerful enough to be self-hosting.

#### Example Rusty Program

```rust
// Rusty: A simplified Rust-like language

struct Point {
    x: i64,
    y: i64,
}

enum Option<T> {
    Some(T),
    None,
}

impl Point {
    fn new(x: i64, y: i64) -> Point {
        Point { x, y }
    }
    
    fn distance(&self, other: &Point) -> f64 {
        let dx = (self.x - other.x) as f64;
        let dy = (self.y - other.y) as f64;
        (dx * dx + dy * dy).sqrt()
    }
}

fn fibonacci(n: i64) -> i64 {
    match n {
        0 => 0,
        1 => 1,
        _ => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

fn main() {
    let mut sum = 0;
    let mut i = 0;
    
    while i < 10 {
        sum = sum + fibonacci(i);
        i = i + 1;
    }
    
    print("Sum of first 10 Fibonacci numbers:");
    print(sum);
}
```

#### Rusty Language Features

**Core Features (Required for self-hosting):**
- ✅ Variables: `let` and `let mut`
- ✅ Basic types: `i64`, `f64`, `bool`, `String`, `()`
- ✅ Functions with type annotations
- ✅ Control flow: `if`, `while`, `loop`, `for`
- ✅ Structs
- ✅ Enums with data
- ✅ Pattern matching with `match`
- ✅ Method syntax with `impl`
- ✅ Basic ownership (move semantics, explicit clone)
- ✅ References (read-only, no mutable refs for simplicity)

**Simplified from Rust:**
- ❌ No lifetimes (use garbage collection)
- ❌ No borrow checker (just move semantics)
- ⚠️ No macros (except built-in `print!`)
- ⚠️ Simple generics (no trait bounds)
- ⚠️ Basic traits (no associated types)

### 4.3 Formal Grammar: EBNF

We describe Rusty's syntax using Extended Backus-Naur Form (EBNF):

```ebnf
(* Program structure *)
program        = { item } ;
item           = function_def | struct_def | enum_def | impl_block ;

(* Functions *)
function_def   = "fn" , identifier , "(" , [ param_list ] , ")" , 
                 [ "->" , type ] , block ;
param_list     = param , { "," , param } ;
param          = identifier , ":" , type ;

(* Structs *)
struct_def     = "struct" , identifier , "{" , [ field_list ] , "}" ;
field_list     = field , { "," , field } , [ "," ] ;
field          = identifier , ":" , type ;

(* Enums *)
enum_def       = "enum" , identifier , [ "<" , type_params , ">" ] , 
                 "{" , [ variant_list ] , "}" ;
variant_list   = variant , { "," , variant } , [ "," ] ;
variant        = identifier , [ "(" , type_list , ")" | "{" , field_list , "}" ] ;

(* Types *)
type           = "i64" | "f64" | "bool" | "String" | "()" 
               | identifier , [ "<" , type_list , ">" ]
               | "&" , type
               | "fn" , "(" , [ type_list ] , ")" , "->" , type ;
type_list      = type , { "," , type } ;

(* Statements *)
statement      = let_stmt | expr_stmt | return_stmt | block ;
let_stmt       = "let" , [ "mut" ] , identifier , [ ":" , type ] , 
                 "=" , expression , ";" ;
expr_stmt      = expression , ";" ;
return_stmt    = "return" , [ expression ] , ";" ;
block          = "{" , { statement } , [ expression ] , "}" ;

(* Expressions - simplified, precedence handled in parser *)
expression     = if_expr | match_expr | while_expr | loop_expr 
               | binary_expr | unary_expr | call_expr | primary ;

if_expr        = "if" , expression , block , [ "else" , ( block | if_expr ) ] ;
match_expr     = "match" , expression , "{" , { match_arm } , "}" ;
match_arm      = pattern , "=>" , expression , "," ;
while_expr     = "while" , expression , block ;
loop_expr      = "loop" , block ;

(* Primary expressions *)
primary        = literal | identifier | "(" , expression , ")" 
               | struct_literal | array_literal ;
literal        = integer | float | string | "true" | "false" ;

(* Patterns *)
pattern        = literal | identifier | "_" 
               | identifier , "(" , [ pattern_list ] , ")"
               | identifier , "{" , [ field_pattern_list ] , "}" ;
```

### 4.4 Lexical Analysis: What Is It?

**Lexical analysis** (or *lexing* / *tokenization*) breaks source code into a sequence of *tokens*:

```
Source:     let x = 42 + y;

Tokens:     [LET] [IDENT("x")] [ASSIGN] [INT(42)] [PLUS] [IDENT("y")] [SEMI]
```

The lexer:
1. Reads characters from input
2. Groups them into meaningful units (tokens)
3. Classifies each token by kind
4. Discards whitespace and comments
5. Reports errors for invalid input

### 4.5 Token Design

Our tokens for Rusty:

```rust
#[derive(Debug, Clone, PartialEq)]
pub enum TokenKind {
    // === Literals ===
    Integer(i64),        // 42, -17, 0xFF
    Float(f64),          // 3.14, 0.5, 1e10
    String(String),      // "hello\nworld"
    True,                // true
    False,               // false
    
    // === Identifiers ===
    Identifier(String),  // foo, _bar, x1
    
    // === Keywords ===
    Let, Mut, Fn, Return,
    If, Else, While, Loop, For, In,
    Match, Struct, Enum, Impl,
    Break, Continue,
    True, False,
    Self_, SelfType,     // self, Self
    
    // === Operators ===
    // Arithmetic
    Plus, Minus, Star, Slash, Percent,
    
    // Comparison
    EqEq, NotEq, Lt, Gt, LtEq, GtEq,
    
    // Logical
    AndAnd, OrOr, Not,
    
    // Assignment
    Eq,
    
    // === Delimiters ===
    LParen, RParen,      // ( )
    LBrace, RBrace,      // { }
    LBracket, RBracket,  // [ ]
    
    // === Punctuation ===
    Comma,               // ,
    Colon,               // :
    Semicolon,           // ;
    Dot,                 // .
    Arrow,               // ->
    FatArrow,            // =>
    DoubleColon,         // ::
    Ampersand,           // &
    
    // === Special ===
    Eof,
    Error(String),
}
```

### 4.6 Formal Language Theory: Regular Languages

**Key insight:** The set of valid tokens forms a *regular language*.

**Definition:** A language is *regular* if it can be described by:
- A **regular expression**
- A **finite automaton** (DFA or NFA)
- A **regular grammar**

These are all equivalent!

**Examples of regular languages:**
- All identifiers: `[a-zA-Z_][a-zA-Z0-9_]*`
- All integers: `-?[0-9]+`
- All keywords: `let|mut|fn|if|else|...`

**Examples of NON-regular languages:**
- Balanced parentheses: `( ( ) ) ( )`
- Matching brackets: not regular! (needs stack/memory)

This is why we have a separate parser — syntax/grammar is NOT regular!

### 4.7 Regular Expressions for Tokens

| Token Kind | Regular Expression |
|------------|-------------------|
| Identifier | `[a-zA-Z_][a-zA-Z0-9_]*` |
| Integer | `-?[0-9]+` or `0x[0-9a-fA-F]+` |
| Float | `-?[0-9]+\.[0-9]+([eE][+-]?[0-9]+)?` |
| String | `"([^"\\]|\\.)*"` |
| Whitespace | `[ \t\n\r]+` |
| Line comment | `//[^\n]*` |
| Block comment | `/\*([^*]|\*[^/])*\*/` |

### 4.8 Finite Automata: Theory

A **Deterministic Finite Automaton (DFA)** is:
- A finite set of states
- An alphabet (input characters)
- A transition function: state × char → state
- A start state
- A set of accepting states

**Example: DFA for identifiers**

```
         [a-zA-Z_]           [a-zA-Z0-9_]
    ┌───────────────┐    ┌─────────────────┐
    │               ▼    │                 │
  ──►  (start)  ────►  ((accept))  ◄──────┘
                              ▲
                              │ [a-zA-Z0-9_]
                              │
```

States: {start, accept}
Alphabet: all ASCII characters
Transitions:
- start + [a-zA-Z_] → accept
- accept + [a-zA-Z0-9_] → accept
- all other transitions → error/reject

### 4.9 Lexer Design Strategies

**Strategy 1: Hand-written lexer** (what we'll do)
- Full control over error messages
- Easy to understand and debug
- Good performance
- More code to write

**Strategy 2: Lexer generators** (lex, flex, logos)
- Specify patterns, generate code
- Less code to write
- Harder to customize errors

**Our approach:** Hand-written lexer using:
- `Peekable<Chars>` for lookahead
- `match` expressions for character classification
- Helper methods for each token type

### 4.10 Edge Cases in Lexing

**Number parsing:**
```
42        → Integer(42)
42.0      → Float(42.0)
42.foo    → Integer(42), Dot, Identifier("foo")
42..      → Integer(42), DotDot  (range operator?)
```

**Operator ambiguity:**
```
x-y       → Identifier, Minus, Identifier
x - y     → Identifier, Minus, Identifier  (same!)
x--       → Identifier, MinusMinus? or Identifier, Minus, Minus?
```

**Keyword vs. identifier:**
```
let       → Keyword(Let)
letter    → Identifier("letter")  (not a keyword!)
_let      → Identifier("_let")
```

**String escapes:**
```
"hello"       → String("hello")
"hello\n"     → String("hello\n")
"hello\"x"    → String("hello\"x")
"hello        → Error: unterminated string
```

---

## 🧠 Theory: Regular Languages and the Pumping Lemma

**Theorem (Pumping Lemma):** If L is a regular language, then there exists a "pumping length" p such that any string s in L with |s| ≥ p can be written as s = xyz where:
1. |y| > 0 (y is non-empty)
2. |xy| ≤ p
3. For all n ≥ 0, xy^n z ∈ L (we can "pump" y)

**Using the pumping lemma to prove a language is NOT regular:**

**Example:** Prove L = {aⁿbⁿ | n ≥ 0} is not regular.

*Proof by contradiction:*
1. Assume L is regular with pumping length p
2. Consider s = aᵖbᵖ ∈ L (|s| = 2p ≥ p)
3. Write s = xyz with |y| > 0 and |xy| ≤ p
4. Since |xy| ≤ p and s starts with p a's, y consists only of a's
5. Consider xy²z = aᵖ⁺ᵏbᵖ where k = |y| > 0
6. This has more a's than b's, so xy²z ∉ L
7. Contradiction! Therefore L is not regular.

**Why this matters for compilers:**
- Tokens (identifiers, numbers) ARE regular → lexer can handle them
- Nested structures (balanced parens, blocks) are NOT regular → need a parser

---

## 🧠 Theory: DFA Minimization (Preview)

For efficiency, we want the smallest DFA possible. Given any DFA, there's an equivalent minimal DFA (unique up to renaming).

**Minimization algorithm (outline):**
1. Remove unreachable states
2. Partition states into equivalence classes
3. Merge indistinguishable states

We won't implement this (hand-written is simpler), but it's foundational theory!

---

## ✅ Summary

| Concept | Description |
|---------|-------------|
| Compiler pipeline | Lexer → Parser → Semantic → Code Gen |
| Token | Meaningful unit: keyword, identifier, operator, literal |
| Regular language | Can be described by regex/DFA/NFA |
| DFA | Finite automaton, deterministic transitions |
| Pumping lemma | Tool to prove languages are NOT regular |
| Lexer | Converts source text to token stream |

---

## 📚 Further Reading

- [Crafting Interpreters, Chapter 4: Scanning](https://craftinginterpreters.com/scanning.html)
- [Introduction to the Theory of Computation](https://www.amazon.com/Introduction-Theory-Computation-Michael-Sipser/dp/113318779X) — Sipser (for formal language theory)
- [Dragon Book, Chapter 3: Lexical Analysis](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools)

---

**Next:** [Lecture 5: Building a Lexer in Rust](../lecture-05/README.md)

**Exercises:** [Exercise Sheet 4](exercises.md)
