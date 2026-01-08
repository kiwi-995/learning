# Lecture 7: Semantic Analysis & Type Checking

> *"The type checker is the compiler's most helpful critic."*

## 🎯 Learning Objectives

- Implement variable resolution and scoping
- Build a symbol table with HashMap
- Implement a basic type checker
- Detect semantic errors (undefined variables, type mismatches)

## 📖 Topics

### 7.1 Beyond Syntax

The parser checks syntax; semantic analysis checks meaning:
- Are all variables defined before use?
- Do types match in operations?
- Are function calls correct?

### 7.2 Symbol Tables

```rust
struct SymbolTable {
    scopes: Vec<HashMap<String, Symbol>>,
}

struct Symbol {
    name: String,
    ty: Type,
    mutable: bool,
    defined_at: Span,
}
```

### 7.3 Name Resolution

Resolve identifiers to their declarations:

```rust
fn resolve_expr(&mut self, expr: &Expr) -> Result<ResolvedExpr, SemanticError> {
    match expr {
        Expr::Identifier(name) => {
            if let Some(symbol) = self.lookup(name) {
                Ok(ResolvedExpr::Variable(symbol.clone()))
            } else {
                Err(SemanticError::UndefinedVariable(name.clone()))
            }
        }
        // ... handle other cases
    }
}
```

### 7.4 Type Checking

```rust
fn check_binary(&mut self, left: &Expr, op: BinaryOp, right: &Expr) -> Result<Type, TypeError> {
    let left_ty = self.infer_type(left)?;
    let right_ty = self.infer_type(right)?;
    
    match op {
        BinaryOp::Add | BinaryOp::Sub | BinaryOp::Mul | BinaryOp::Div => {
            if left_ty == right_ty && (left_ty == Type::I64 || left_ty == Type::F64) {
                Ok(left_ty)
            } else {
                Err(TypeError::InvalidOperator { op, left: left_ty, right: right_ty })
            }
        }
        // ...
    }
}
```

---

## 🧠 Theory: Type Systems

- **Sound**: Well-typed programs don't have type errors
- **Complete**: All correct programs type-check
- **Decidable**: Type checking always terminates

---

**Exercises:** [Exercise Sheet 7](exercises.md)
**Next:** [Lecture 8: Tree-Walk Interpreter](../lecture-08/README.md)
