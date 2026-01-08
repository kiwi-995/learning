# Exercise Sheet 3: Structs, Enums & Pattern Matching

These exercises build directly toward our compiler. You'll create the token and error types we'll use throughout the project!

---

## Exercise 3.1: Token Structure 🎫

**Objective:** Create a complete token structure for our lexer.

```rust
/// The kind of token
#[derive(Debug, Clone, PartialEq)]
pub enum TokenKind {
    // Literals
    Integer(i64),
    Float(f64),
    String(String),
    Bool(bool),
    
    // Identifiers and keywords
    Identifier(String),
    
    // Keywords
    Let,
    Mut,
    Fn,
    If,
    Else,
    While,
    For,
    Loop,
    Return,
    Match,
    Struct,
    Enum,
    Impl,
    True,
    False,
    
    // Operators
    Plus,        // +
    Minus,       // -
    Star,        // *
    Slash,       // /
    Percent,     // %
    
    // Comparison
    Eq,          // ==
    Ne,          // !=
    Lt,          // <
    Gt,          // >
    Le,          // <=
    Ge,          // >=
    
    // Logical
    And,         // &&
    Or,          // ||
    Not,         // !
    
    // Assignment
    Assign,      // =
    
    // Delimiters
    LeftParen,   // (
    RightParen,  // )
    LeftBrace,   // {
    RightBrace,  // }
    LeftBracket, // [
    RightBracket,// ]
    
    // Punctuation
    Comma,       // ,
    Colon,       // :
    Semicolon,   // ;
    Arrow,       // ->
    FatArrow,    // =>
    Dot,         // .
    DoubleColon, // ::
    
    // Special
    Eof,
}

/// The location of a token in source code
#[derive(Debug, Clone, Copy, PartialEq)]
pub struct Span {
    pub start: usize,  // Byte offset
    pub end: usize,    // Byte offset (exclusive)
    pub line: usize,
    pub column: usize,
}

/// A token with its kind and location
#[derive(Debug, Clone, PartialEq)]
pub struct Token {
    pub kind: TokenKind,
    pub span: Span,
}

impl Token {
    pub fn new(kind: TokenKind, span: Span) -> Token {
        todo!()
    }
    
    /// Returns true if this token is a keyword
    pub fn is_keyword(&self) -> bool {
        todo!()
        // Hint: Use matches! macro or match expression
    }
    
    /// Returns true if this token is a literal (number, string, bool)
    pub fn is_literal(&self) -> bool {
        todo!()
    }
    
    /// Returns true if this token is an operator
    pub fn is_operator(&self) -> bool {
        todo!()
    }
}

impl Span {
    pub fn new(start: usize, end: usize, line: usize, column: usize) -> Span {
        todo!()
    }
    
    /// Combines two spans into one that covers both
    pub fn merge(&self, other: &Span) -> Span {
        todo!()
    }
}

/// Converts a keyword string to its TokenKind, or returns None if not a keyword
pub fn keyword_to_token(s: &str) -> Option<TokenKind> {
    todo!()
    // Hint: Use a match on the string
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_keyword_detection() {
        assert_eq!(keyword_to_token("let"), Some(TokenKind::Let));
        assert_eq!(keyword_to_token("mut"), Some(TokenKind::Mut));
        assert_eq!(keyword_to_token("fn"), Some(TokenKind::Fn));
        assert_eq!(keyword_to_token("foo"), None);
    }
    
    #[test]
    fn test_is_keyword() {
        let token = Token::new(TokenKind::Let, Span::new(0, 3, 1, 1));
        assert!(token.is_keyword());
        
        let token = Token::new(TokenKind::Identifier("foo".to_string()), Span::new(0, 3, 1, 1));
        assert!(!token.is_keyword());
    }
    
    #[test]
    fn test_span_merge() {
        let span1 = Span::new(0, 5, 1, 1);
        let span2 = Span::new(10, 15, 1, 11);
        let merged = span1.merge(&span2);
        
        assert_eq!(merged.start, 0);
        assert_eq!(merged.end, 15);
    }
}
```

---

## Exercise 3.2: Compiler Error Type 🚨

**Objective:** Create a comprehensive error type for the compiler.

```rust
use std::fmt;

/// A location in the source code
#[derive(Debug, Clone, Copy)]
pub struct Location {
    pub line: usize,
    pub column: usize,
}

/// The kind of compiler error
#[derive(Debug, Clone)]
pub enum ErrorKind {
    // Lexer errors
    UnexpectedCharacter(char),
    UnterminatedString,
    InvalidNumber(String),
    InvalidEscape(char),
    
    // Parser errors
    UnexpectedToken { expected: String, found: String },
    UnexpectedEof,
    InvalidExpression,
    
    // Semantic errors
    UndefinedVariable(String),
    UndefinedFunction(String),
    TypeMismatch { expected: String, found: String },
    DuplicateDefinition(String),
    
    // Type errors
    CannotInferType,
    InvalidOperator { op: String, left: String, right: String },
}

/// A complete compiler error with location and context
#[derive(Debug, Clone)]
pub struct CompileError {
    pub kind: ErrorKind,
    pub location: Location,
    pub source_line: Option<String>,  // The actual source code line
}

impl CompileError {
    pub fn new(kind: ErrorKind, location: Location) -> CompileError {
        todo!()
    }
    
    pub fn with_source_line(mut self, line: String) -> CompileError {
        todo!()
    }
    
    /// Returns a user-friendly error code (e.g., "E001")
    pub fn error_code(&self) -> &'static str {
        todo!()
        // Assign different codes to different error kinds
    }
}

impl fmt::Display for CompileError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        // TODO: Format the error nicely
        // Example output:
        // error[E001]: unexpected character '@' at 5:10
        //    |
        //  5 | let x = @invalid
        //    |         ^
        todo!()
    }
}

impl std::error::Error for CompileError {}

/// A result type for compiler operations
pub type CompileResult<T> = Result<T, CompileError>;

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_error_display() {
        let error = CompileError::new(
            ErrorKind::UnexpectedCharacter('@'),
            Location { line: 5, column: 10 },
        ).with_source_line("let x = @invalid".to_string());
        
        let message = format!("{}", error);
        assert!(message.contains("unexpected character"));
        assert!(message.contains("@"));
    }
    
    #[test]
    fn test_type_mismatch() {
        let error = CompileError::new(
            ErrorKind::TypeMismatch {
                expected: "i64".to_string(),
                found: "String".to_string(),
            },
            Location { line: 10, column: 5 },
        );
        
        let message = format!("{}", error);
        assert!(message.contains("i64"));
        assert!(message.contains("String"));
    }
}
```

---

## Exercise 3.3: Expression AST 🌳

**Objective:** Implement an expression AST with evaluation.

```rust
#[derive(Debug, Clone, PartialEq)]
pub enum Expr {
    // Literals
    Number(f64),
    Bool(bool),
    
    // Binary operations
    Add(Box<Expr>, Box<Expr>),
    Sub(Box<Expr>, Box<Expr>),
    Mul(Box<Expr>, Box<Expr>),
    Div(Box<Expr>, Box<Expr>),
    
    // Comparison
    Eq(Box<Expr>, Box<Expr>),
    Lt(Box<Expr>, Box<Expr>),
    Gt(Box<Expr>, Box<Expr>),
    
    // Logical
    And(Box<Expr>, Box<Expr>),
    Or(Box<Expr>, Box<Expr>),
    Not(Box<Expr>),
    
    // Conditional
    If {
        condition: Box<Expr>,
        then_branch: Box<Expr>,
        else_branch: Box<Expr>,
    },
}

/// Runtime value
#[derive(Debug, Clone, PartialEq)]
pub enum Value {
    Number(f64),
    Bool(bool),
}

/// Evaluation error
#[derive(Debug)]
pub enum EvalError {
    TypeError(String),
    DivisionByZero,
}

impl Expr {
    /// Evaluates the expression to a value
    pub fn eval(&self) -> Result<Value, EvalError> {
        match self {
            Expr::Number(n) => Ok(Value::Number(*n)),
            Expr::Bool(b) => Ok(Value::Bool(*b)),
            
            Expr::Add(left, right) => {
                todo!()
                // Evaluate both sides, check they're numbers, add them
            }
            
            Expr::Sub(left, right) => {
                todo!()
            }
            
            Expr::Mul(left, right) => {
                todo!()
            }
            
            Expr::Div(left, right) => {
                todo!()
                // Don't forget to check for division by zero!
            }
            
            Expr::Eq(left, right) => {
                todo!()
            }
            
            Expr::Lt(left, right) => {
                todo!()
            }
            
            Expr::Gt(left, right) => {
                todo!()
            }
            
            Expr::And(left, right) => {
                todo!()
                // Short-circuit evaluation!
            }
            
            Expr::Or(left, right) => {
                todo!()
                // Short-circuit evaluation!
            }
            
            Expr::Not(operand) => {
                todo!()
            }
            
            Expr::If { condition, then_branch, else_branch } => {
                todo!()
            }
        }
    }
    
    /// Pretty prints the expression
    pub fn pretty_print(&self) -> String {
        match self {
            Expr::Number(n) => n.to_string(),
            Expr::Bool(b) => b.to_string(),
            Expr::Add(l, r) => format!("({} + {})", l.pretty_print(), r.pretty_print()),
            Expr::Sub(l, r) => format!("({} - {})", l.pretty_print(), r.pretty_print()),
            Expr::Mul(l, r) => format!("({} * {})", l.pretty_print(), r.pretty_print()),
            Expr::Div(l, r) => format!("({} / {})", l.pretty_print(), r.pretty_print()),
            Expr::Eq(l, r) => format!("({} == {})", l.pretty_print(), r.pretty_print()),
            Expr::Lt(l, r) => format!("({} < {})", l.pretty_print(), r.pretty_print()),
            Expr::Gt(l, r) => format!("({} > {})", l.pretty_print(), r.pretty_print()),
            Expr::And(l, r) => format!("({} && {})", l.pretty_print(), r.pretty_print()),
            Expr::Or(l, r) => format!("({} || {})", l.pretty_print(), r.pretty_print()),
            Expr::Not(e) => format!("!{}", e.pretty_print()),
            Expr::If { condition, then_branch, else_branch } => {
                format!("if {} then {} else {}", 
                    condition.pretty_print(),
                    then_branch.pretty_print(),
                    else_branch.pretty_print())
            }
        }
    }
}

// Helper constructors
impl Expr {
    pub fn num(n: f64) -> Expr {
        Expr::Number(n)
    }
    
    pub fn bool(b: bool) -> Expr {
        Expr::Bool(b)
    }
    
    pub fn add(left: Expr, right: Expr) -> Expr {
        Expr::Add(Box::new(left), Box::new(right))
    }
    
    // TODO: Add more constructors for other variants
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_simple_arithmetic() {
        // (1 + 2) * 3 = 9
        let expr = Expr::mul(
            Expr::add(Expr::num(1.0), Expr::num(2.0)),
            Expr::num(3.0),
        );
        
        assert_eq!(expr.eval().unwrap(), Value::Number(9.0));
    }
    
    #[test]
    fn test_comparison() {
        let expr = Expr::Lt(
            Box::new(Expr::num(5.0)),
            Box::new(Expr::num(10.0)),
        );
        
        assert_eq!(expr.eval().unwrap(), Value::Bool(true));
    }
    
    #[test]
    fn test_if_expression() {
        // if 5 < 10 then 1 else 2
        let expr = Expr::If {
            condition: Box::new(Expr::Lt(
                Box::new(Expr::num(5.0)),
                Box::new(Expr::num(10.0)),
            )),
            then_branch: Box::new(Expr::num(1.0)),
            else_branch: Box::new(Expr::num(2.0)),
        };
        
        assert_eq!(expr.eval().unwrap(), Value::Number(1.0));
    }
    
    #[test]
    fn test_short_circuit_and() {
        // false && (1/0) should NOT cause division by zero
        let expr = Expr::And(
            Box::new(Expr::bool(false)),
            Box::new(Expr::Div(
                Box::new(Expr::num(1.0)),
                Box::new(Expr::num(0.0)),
            )),
        );
        
        assert_eq!(expr.eval().unwrap(), Value::Bool(false));
    }
    
    #[test]
    fn test_division_by_zero() {
        let expr = Expr::Div(
            Box::new(Expr::num(1.0)),
            Box::new(Expr::num(0.0)),
        );
        
        match expr.eval() {
            Err(EvalError::DivisionByZero) => (),
            _ => panic!("Expected division by zero error"),
        }
    }
}
```

---

## Exercise 3.4: Statement AST 📝

**Objective:** Extend the AST to include statements.

```rust
use std::collections::HashMap;

#[derive(Debug, Clone)]
pub enum Stmt {
    /// Let binding: `let x = expr;` or `let mut x = expr;`
    Let {
        name: String,
        mutable: bool,
        initializer: Expr,
    },
    
    /// Assignment: `x = expr;`
    Assign {
        name: String,
        value: Expr,
    },
    
    /// Expression statement: `expr;`
    Expression(Expr),
    
    /// Print statement: `print(expr);`
    Print(Expr),
    
    /// Block: `{ stmt* }`
    Block(Vec<Stmt>),
    
    /// If statement: `if cond { ... } else { ... }`
    If {
        condition: Expr,
        then_branch: Box<Stmt>,
        else_branch: Option<Box<Stmt>>,
    },
    
    /// While loop: `while cond { ... }`
    While {
        condition: Expr,
        body: Box<Stmt>,
    },
}

/// An environment stores variable bindings
#[derive(Debug, Clone)]
pub struct Environment {
    values: HashMap<String, Value>,
    mutability: HashMap<String, bool>,
}

impl Environment {
    pub fn new() -> Environment {
        Environment {
            values: HashMap::new(),
            mutability: HashMap::new(),
        }
    }
    
    /// Defines a new variable
    pub fn define(&mut self, name: String, value: Value, mutable: bool) {
        todo!()
    }
    
    /// Gets a variable's value
    pub fn get(&self, name: &str) -> Option<&Value> {
        todo!()
    }
    
    /// Assigns a new value to an existing variable
    /// Returns Err if variable doesn't exist or isn't mutable
    pub fn assign(&mut self, name: &str, value: Value) -> Result<(), String> {
        todo!()
    }
}

/// Interpreter for statements
pub struct Interpreter {
    env: Environment,
    output: Vec<String>,  // Captured print output
}

impl Interpreter {
    pub fn new() -> Interpreter {
        Interpreter {
            env: Environment::new(),
            output: Vec::new(),
        }
    }
    
    /// Executes a statement
    pub fn execute(&mut self, stmt: &Stmt) -> Result<(), String> {
        match stmt {
            Stmt::Let { name, mutable, initializer } => {
                todo!()
            }
            
            Stmt::Assign { name, value } => {
                todo!()
            }
            
            Stmt::Expression(expr) => {
                todo!()
            }
            
            Stmt::Print(expr) => {
                todo!()
            }
            
            Stmt::Block(statements) => {
                todo!()
            }
            
            Stmt::If { condition, then_branch, else_branch } => {
                todo!()
            }
            
            Stmt::While { condition, body } => {
                todo!()
            }
        }
    }
    
    /// Returns captured output
    pub fn get_output(&self) -> &[String] {
        &self.output
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_let_and_print() {
        let mut interp = Interpreter::new();
        
        let program = vec![
            Stmt::Let {
                name: "x".to_string(),
                mutable: false,
                initializer: Expr::num(42.0),
            },
            Stmt::Print(Expr::Identifier("x".to_string())),
        ];
        
        for stmt in &program {
            interp.execute(stmt).unwrap();
        }
        
        assert_eq!(interp.get_output(), &["42"]);
    }
    
    #[test]
    fn test_mutable_assignment() {
        let mut interp = Interpreter::new();
        
        let program = vec![
            Stmt::Let {
                name: "x".to_string(),
                mutable: true,
                initializer: Expr::num(1.0),
            },
            Stmt::Assign {
                name: "x".to_string(),
                value: Expr::num(2.0),
            },
            Stmt::Print(Expr::Identifier("x".to_string())),
        ];
        
        for stmt in &program {
            interp.execute(stmt).unwrap();
        }
        
        assert_eq!(interp.get_output(), &["2"]);
    }
    
    #[test]
    fn test_immutable_assignment_fails() {
        let mut interp = Interpreter::new();
        
        interp.execute(&Stmt::Let {
            name: "x".to_string(),
            mutable: false,
            initializer: Expr::num(1.0),
        }).unwrap();
        
        let result = interp.execute(&Stmt::Assign {
            name: "x".to_string(),
            value: Expr::num(2.0),
        });
        
        assert!(result.is_err());
    }
}
```

---

## 📊 Theory Exercise 3.5: Grammar for Expressions

**Objective:** Practice writing formal grammars.

Write a grammar in BNF (Backus-Naur Form) for the expression language in Exercise 3.3.

**Example BNF:**
```
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
<number> ::= <digit> | <digit> <number>
```

**Your task:**

1. Write BNF rules for:
   - `<expr>` — a complete expression
   - `<bool>` — a boolean literal
   - `<number>` — a numeric literal
   - `<binary_op>` — a binary operator
   - `<if_expr>` — an if expression

2. Consider operator precedence. In the expression `1 + 2 * 3`:
   - How should this be parsed?
   - Modify your grammar to handle precedence correctly

3. **Challenge:** Show that your grammar is unambiguous by constructing the unique parse tree for `if true then 1 + 2 else 3 * 4`.

---

## 📊 Theory Exercise 3.6: Algebraic Data Types

**Objective:** Understand the algebraic properties of types.

1. **Counting values:** How many distinct values can each type hold?
   
   a) `bool`
   b) `Option<bool>`
   c) `Result<bool, bool>`
   d) `(bool, bool)`
   e) `(bool, Option<bool>)`

2. **Isomorphisms:** Two types are *isomorphic* if they have the same number of values. Show that these types are isomorphic:
   
   a) `Option<bool>` and `Result<bool, ()>`
   b) `Result<A, B>` and `Result<B, A>` (what function converts between them?)

3. **Recursive types:** Consider this definition:
   ```rust
   enum List<T> {
       Nil,
       Cons(T, Box<List<T>>),
   }
   ```
   
   a) What are all possible values of `List<()>` of length ≤ 3?
   b) Express the "size" of `List<T>` as a formula in terms of `|T|` (the size of T)
   
   Hint: If `L = 1 + T * L`, solve for L algebraically. What does this remind you of?

---

## 🏆 Bonus Challenge: Visitor Pattern

Implement the visitor pattern for the expression AST:

```rust
trait ExprVisitor<T> {
    fn visit_number(&mut self, n: f64) -> T;
    fn visit_bool(&mut self, b: bool) -> T;
    fn visit_add(&mut self, left: &Expr, right: &Expr) -> T;
    fn visit_sub(&mut self, left: &Expr, right: &Expr) -> T;
    fn visit_mul(&mut self, left: &Expr, right: &Expr) -> T;
    fn visit_div(&mut self, left: &Expr, right: &Expr) -> T;
    // ... etc
}

impl Expr {
    fn accept<T>(&self, visitor: &mut dyn ExprVisitor<T>) -> T {
        match self {
            Expr::Number(n) => visitor.visit_number(*n),
            Expr::Bool(b) => visitor.visit_bool(*b),
            Expr::Add(l, r) => visitor.visit_add(l, r),
            // ... etc
            todo!()
        }
    }
}

// Implement a PrettyPrinter using the visitor pattern
struct PrettyPrinter;

impl ExprVisitor<String> for PrettyPrinter {
    // TODO: Implement all methods
}

// Implement a ConstantFolder using the visitor pattern
struct ConstantFolder;

impl ExprVisitor<Expr> for ConstantFolder {
    // TODO: Fold constants where possible
    // e.g., Add(Number(1), Number(2)) -> Number(3)
}
```

---

## 📝 Submission Checklist

- [ ] Exercise 3.1: Token types complete with tests
- [ ] Exercise 3.2: Error type with Display implementation
- [ ] Exercise 3.3: Expression AST evaluates correctly
- [ ] Exercise 3.4: Statement interpreter works
- [ ] Exercise 3.5: Grammar written in BNF
- [ ] Exercise 3.6: Algebraic data type questions answered
- [ ] Bonus: Visitor pattern implemented (optional)

---

**Solutions:** [solutions/](solutions/)

**Next Lecture:** [Lecture 4: Language Design & Lexical Analysis](../lecture-04/README.md)
