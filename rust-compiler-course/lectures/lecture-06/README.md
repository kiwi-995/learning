# Lecture 6: Parsing & Building ASTs

> *"The parser transforms a flat stream of tokens into a meaningful tree structure."*

## 🎯 Learning Objectives

- Understand recursive descent parsing
- Implement a parser using Pratt parsing for operator precedence
- Build Abstract Syntax Trees using Rust enums
- Handle parse errors gracefully

## 📖 Topics

### 6.1 From Tokens to Trees

The parser takes tokens from the lexer and builds an AST.

### 6.2 Recursive Descent Parsing

Each grammar rule becomes a function:

```rust
impl Parser {
    fn parse_statement(&mut self) -> Result<Stmt, ParseError> {
        if self.check(TokenKind::Let) {
            self.parse_let_statement()
        } else if self.check(TokenKind::If) {
            self.parse_if_statement()
        } else {
            self.parse_expression_statement()
        }
    }
}
```

### 6.3 Pratt Parsing for Expressions

Handle operator precedence with binding power:

```rust
fn parse_expression(&mut self, min_bp: u8) -> Result<Expr, ParseError> {
    let mut lhs = self.parse_prefix()?;
    
    while let Some(op) = self.peek_infix_op() {
        let (l_bp, r_bp) = infix_binding_power(op);
        if l_bp < min_bp {
            break;
        }
        self.advance();
        let rhs = self.parse_expression(r_bp)?;
        lhs = Expr::Binary { left: Box::new(lhs), op, right: Box::new(rhs) };
    }
    
    Ok(lhs)
}
```

### 6.4 AST Definitions

```rust
pub enum Expr {
    Integer(i64),
    Float(f64),
    String(String),
    Bool(bool),
    Identifier(String),
    Binary { left: Box<Expr>, op: BinaryOp, right: Box<Expr> },
    Unary { op: UnaryOp, operand: Box<Expr> },
    Call { callee: Box<Expr>, args: Vec<Expr> },
    If { condition: Box<Expr>, then_branch: Box<Expr>, else_branch: Option<Box<Expr>> },
    Match { scrutinee: Box<Expr>, arms: Vec<MatchArm> },
    Block(Vec<Stmt>, Option<Box<Expr>>),
}

pub enum Stmt {
    Let { name: String, mutable: bool, ty: Option<Type>, init: Expr },
    Expression(Expr),
    Return(Option<Expr>),
}
```

### 6.5 Error Recovery

Use synchronization to continue parsing after errors.

---

## 🧠 Theory: LL vs LR Parsing

- **LL(k)**: Top-down, left-to-right, k tokens lookahead
- **LR(k)**: Bottom-up, shift-reduce parsing
- **Recursive descent**: LL(1) implemented by hand

---

**Exercises:** [Exercise Sheet 6](exercises.md)
**Next:** [Lecture 7: Semantic Analysis & Type Checking](../lecture-07/README.md)
