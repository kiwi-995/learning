# Appendix B: Compiler Theory Deep-Dive

> *"Theory is the map; implementation is the journey."*

This appendix provides the theoretical foundations that underpin everything in the main course.

---

## B.1 Formal Language Theory

### Automata Theory

**Definition (Finite Automaton):** A DFA is a 5-tuple (Q, Σ, δ, q₀, F) where:
- Q = finite set of states
- Σ = alphabet (finite set of input symbols)
- δ: Q × Σ → Q = transition function
- q₀ ∈ Q = start state
- F ⊆ Q = set of accepting states

**Theorem:** Regular languages are exactly those recognized by DFAs.

### Regular Expressions

| Regex | Language |
|-------|----------|
| ∅ | Empty language |
| ε | {""} (empty string) |
| a | {"a"} |
| R₁R₂ | Concatenation |
| R₁\|R₂ | Union |
| R* | Kleene star (0 or more repetitions) |

**Theorem:** Regular expressions and finite automata have equivalent expressive power.

### The Pumping Lemma

**Theorem (Pumping Lemma for Regular Languages):**

If L is regular, ∃p ≥ 1 (the "pumping length") such that for any s ∈ L with |s| ≥ p, s can be written as s = xyz where:
1. |y| > 0
2. |xy| ≤ p  
3. ∀n ≥ 0: xyⁿz ∈ L

**Application:** Prove L = {aⁿbⁿ | n ≥ 0} is not regular.

*Proof:*
1. Assume L is regular with pumping length p
2. Choose s = aᵖbᵖ ∈ L
3. By conditions 1 and 2, y = aᵏ for some k > 0
4. Then xy²z = aᵖ⁺ᵏbᵖ ∉ L (more a's than b's)
5. Contradiction ∎

---

## B.2 Context-Free Languages

### Context-Free Grammars

**Definition:** A CFG is a 4-tuple (V, Σ, R, S) where:
- V = finite set of variables (non-terminals)
- Σ = finite set of terminals
- R = finite set of production rules V → (V ∪ Σ)*
- S ∈ V = start symbol

**Example:** Grammar for balanced parentheses:
```
S → (S) | SS | ε
```

### Parse Trees and Ambiguity

**Definition:** A grammar G is *ambiguous* if some string w ∈ L(G) has more than one parse tree.

**Example:** The grammar E → E + E | E * E | (E) | id is ambiguous.

String "id + id * id" has two parse trees:
```
      E                    E
    / | \                / | \
   E  +  E              E  *  E
   |   / | \          / | \   |
  id  E  *  E        E  +  E  id
      |     |        |     |
     id    id       id    id
```

### Chomsky Normal Form

**Definition:** A CFG is in CNF if every rule has the form:
- A → BC (where B, C are non-terminals)
- A → a (where a is a terminal)
- S → ε (only for start symbol if ε ∈ L)

**Theorem:** Every CFG can be converted to an equivalent CNF grammar.

---

## B.3 Parsing Theory

### LL Parsing

**First and Follow Sets:**

- FIRST(α) = set of terminals that can begin strings derived from α
- FOLLOW(A) = set of terminals that can appear after A in a derivation

**LL(1) Condition:** Grammar G is LL(1) if for each non-terminal A with productions A → α | β:
1. FIRST(α) ∩ FIRST(β) = ∅
2. If ε ∈ FIRST(α), then FIRST(β) ∩ FOLLOW(A) = ∅

### LR Parsing

**LR(0) Items:** A production with a dot indicating parsing progress.
- A → ·XYZ (about to parse XYZ)
- A → X·YZ (parsed X, about to parse YZ)
- A → XYZ· (finished parsing)

**Canonical LR(0) Collection:** Build states from items using closure and goto.

### Pratt Parsing

Elegant top-down operator-precedence parsing:

```
parse_expression(min_bp):
    lhs = parse_prefix()
    while next token is infix operator with binding power ≥ min_bp:
        op = consume operator
        rhs = parse_expression(right_binding_power(op))
        lhs = make_binary(lhs, op, rhs)
    return lhs
```

---

## B.4 Type Theory

### Simply Typed Lambda Calculus

**Typing Rules:**

```
    x : τ ∈ Γ
    ─────────── (Var)
    Γ ⊢ x : τ

    Γ, x : τ₁ ⊢ e : τ₂
    ─────────────────────── (Abs)
    Γ ⊢ (λx.e) : τ₁ → τ₂

    Γ ⊢ e₁ : τ₁ → τ₂    Γ ⊢ e₂ : τ₁
    ─────────────────────────────────── (App)
    Γ ⊢ (e₁ e₂) : τ₂
```

### Hindley-Milner Type Inference

**Algorithm W:**
1. Assign fresh type variables to untyped subexpressions
2. Generate constraints from typing rules
3. Solve constraints via unification
4. Substitute solution into type variables

**Principal Types:** HM guarantees the most general type.

### Unification

**Unification Algorithm:**
```
unify(τ₁, τ₂):
    if τ₁ = α (type variable):
        if α occurs in τ₂: error (infinite type)
        else: return {α ↦ τ₂}
    if τ₂ = β (type variable):
        return {β ↦ τ₁}
    if τ₁ = C(σ₁...σₙ) and τ₂ = C(ρ₁...ρₙ):
        return unify(σ₁,ρ₁) ∪ ... ∪ unify(σₙ,ρₙ)
    else: error (types don't match)
```

---

## B.5 Operational Semantics

### Small-Step Semantics

Defines single computation steps:

```
    ─────────────────────────── (E-Add)
    n₁ + n₂ → n₁ + n₂ (arithmetic)

    e₁ → e₁'
    ─────────────────── (E-Add-L)
    e₁ + e₂ → e₁' + e₂

    e₂ → e₂'
    ─────────────────── (E-Add-R)
    v₁ + e₂ → v₁ + e₂'
```

### Big-Step Semantics

Evaluates directly to final value:

```
    n ⇓ n (numbers evaluate to themselves)

    e₁ ⇓ n₁    e₂ ⇓ n₂    n = n₁ + n₂
    ───────────────────────────────────
    e₁ + e₂ ⇓ n
```

---

## B.6 Data Flow Analysis

### Control Flow Graphs

A CFG represents program control flow:
- Nodes = basic blocks (straight-line code)
- Edges = possible control transfers

### Reaching Definitions

A definition d reaches point p if there's a path from d to p where d is not redefined.

**Data flow equations:**
```
IN[B] = ⋃ OUT[P] for all predecessors P
OUT[B] = GEN[B] ∪ (IN[B] - KILL[B])
```

### Live Variable Analysis

Variable v is live at point p if v is used on some path from p before being redefined.

**Backward analysis:**
```
OUT[B] = ⋃ IN[S] for all successors S
IN[B] = USE[B] ∪ (OUT[B] - DEF[B])
```

---

## B.7 Garbage Collection

### Reference Counting

Each object has a reference count. When count reaches 0, free the object.

**Problem:** Cannot collect cycles (A → B → A).

### Mark-and-Sweep

1. **Mark phase:** Starting from roots, mark all reachable objects
2. **Sweep phase:** Free all unmarked objects

### Copying Collection (Cheney's Algorithm)

Divide heap into two semispaces:
1. Allocate in "from" space until full
2. Copy all reachable objects to "to" space  
3. Swap spaces

**Benefit:** Compaction is free!

### Generational GC

Based on the *generational hypothesis*: most objects die young.

- Divide heap into generations (young, old)
- Collect young generation frequently
- Promote survivors to older generations
- Collect old generation rarely

---

## B.8 Register Allocation

### Graph Coloring

1. Build interference graph (variables interfere if live simultaneously)
2. Color graph with k colors (k = number of registers)
3. Variables with same color share a register
4. If k-coloring fails, spill some variables to memory

### Linear Scan

Simpler alternative:
1. Compute live intervals for each variable
2. Scan intervals in order
3. Allocate registers greedily, spill when necessary

---

## Theory Exercises

### Exercise B.1: Pumping Lemma
Prove that {ww | w ∈ {a,b}*} is not regular.

### Exercise B.2: First/Follow Sets
Compute First and Follow for this grammar:
```
E → E + T | T
T → T * F | F
F → (E) | id
```

### Exercise B.3: Type Derivation
Derive the type of: λf. λx. f (f x)

### Exercise B.4: Operational Semantics
Write small-step rules for if-expressions.

### Exercise B.5: Data Flow
Compute reaching definitions for:
```
1: a = 1
2: b = 2
3: c = a + b
4: if c > 0 goto 6
5: a = 3
6: d = a + c
```

---

## Further Reading

### Textbooks
- **Compilers: Principles, Techniques, and Tools** (Dragon Book) — Aho, Lam, Sethi, Ullman
- **Modern Compiler Implementation in ML** — Andrew Appel
- **Engineering a Compiler** — Cooper & Torczon
- **Types and Programming Languages** — Benjamin Pierce
- **Introduction to the Theory of Computation** — Michael Sipser

### Papers
- "A Unified Theory of Garbage Collection" — Bacon, Cheng, Rajan
- "Linear Scan Register Allocation" — Poletto & Sarkar
- "Simple and Efficient Construction of SSA Form" — Braun et al.
- "Parsing Techniques" — Grune & Jacobs
