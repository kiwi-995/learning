# Lecture 8: Tree-Walk Interpreter

> *"The simplest way to execute code: walk the tree."*

## 🎯 Learning Objectives

- Implement a tree-walk interpreter
- Handle runtime values with Rust enums
- Build environments for variable storage
- Implement functions and closures

## 📖 Topics

### 8.1 Interpretation Strategies

- **Tree-walk**: Traverse and evaluate AST directly (simple, slow)
- **Bytecode VM**: Compile to bytecode, interpret (faster)
- **JIT**: Compile to native code at runtime (fastest)

### 8.2 Runtime Values

```rust
#[derive(Debug, Clone)]
pub enum Value {
    Integer(i64),
    Float(f64),
    Bool(bool),
    String(String),
    Unit,
    Function(Rc<Function>),
    Struct { name: String, fields: HashMap<String, Value> },
}
```

### 8.3 Environment

```rust
pub struct Environment {
    values: HashMap<String, Value>,
    parent: Option<Rc<RefCell<Environment>>>,
}
```

### 8.4 The Interpreter Loop

```rust
fn evaluate(&mut self, expr: &Expr) -> Result<Value, RuntimeError> {
    match expr {
        Expr::Integer(n) => Ok(Value::Integer(*n)),
        Expr::Binary { left, op, right } => {
            let l = self.evaluate(left)?;
            let r = self.evaluate(right)?;
            self.apply_binary(op, l, r)
        }
        Expr::Call { callee, args } => {
            let func = self.evaluate(callee)?;
            let arg_values: Vec<Value> = args.iter()
                .map(|a| self.evaluate(a))
                .collect::<Result<_, _>>()?;
            self.call_function(func, arg_values)
        }
        // ...
    }
}
```

---

**Exercises:** [Exercise Sheet 8](exercises.md)
**Next:** [Lecture 9: Bytecode Virtual Machine](../lecture-09/README.md)
