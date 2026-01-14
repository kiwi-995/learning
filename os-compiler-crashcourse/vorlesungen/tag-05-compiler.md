# Tag 5: Compiler Fundamentals – Lexer & Parser in Ruby

**Ziel**: Lerne die Grundlagen des Compilerbaus und baue einen Expression-Parser in Ruby.

---

## 📚 Theorie

### 5.1 Was ist ein Compiler?

```
┌─────────────────────────────────────────────────────────────────────┐
│                        Compiler Pipeline                            │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Source Code     "4 * 7 / (4 - 3)"                                  │
│      │                                                              │
│      ▼                                                              │
│  ┌─────────┐                                                        │
│  │  Lexer  │  → Tokens: [INT(4), STAR, INT(7), SLASH, ...]         │
│  └────┬────┘                                                        │
│       ▼                                                             │
│  ┌─────────┐     ┌───────────────┐                                  │
│  │ Parser  │  →  │      AST      │                                  │
│  └────┬────┘     │    (/)        │                                  │
│       │          │   /   \       │                                  │
│       │          │ (*)   (-)     │                                  │
│       │          │ / \   / \     │                                  │
│       │          │ 4  7 4   3    │                                  │
│       ▼          └───────────────┘                                  │
│  ┌───────────────┐                                                  │
│  │ Code Generator│  → Bytecode oder Maschinencode                   │
│  └───────────────┘                                                  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### 5.2 Lexer (Tokenizer)

Der Lexer wandelt Text in eine Folge von **Tokens**:

```
Input:  "42 + 3 * x"

Tokens:
  1. Token(INT, 42)
  2. Token(PLUS)
  3. Token(INT, 3)
  4. Token(STAR)
  5. Token(IDENT, "x")
```

Reguläre Ausdrücke für Tokens:
```
INTEGER:    /[0-9]+/
IDENTIFIER: /[a-zA-Z_][a-zA-Z0-9_]*/
OPERATORS:  /[+\-*\/()]/
WHITESPACE: /\s+/  (wird übersprungen)
```

### 5.3 Parser und Grammatiken

Eine **Grammatik** beschreibt die Struktur der Sprache:

```
Expression Grammar (BNF-ähnlich):

expr     → term (('+' | '-') term)*
term     → factor (('*' | '/') factor)*
factor   → INTEGER | '(' expr ')'
```

Diese Grammatik definiert **Operator-Präzedenz**:
- `*` und `/` binden stärker als `+` und `-`
- Klammern überschreiben die Präzedenz

### 5.4 Recursive Descent Parser

Ein **Recursive Descent Parser** hat eine Funktion pro Grammatik-Regel:

```ruby
def parse_expr
  left = parse_term
  while current_token in ['+', '-']
    op = consume_token
    right = parse_term
    left = BinaryOp.new(op, left, right)
  end
  left
end

def parse_term
  left = parse_factor
  while current_token in ['*', '/']
    op = consume_token
    right = parse_factor
    left = BinaryOp.new(op, left, right)
  end
  left
end

def parse_factor
  if current_token.type == :INTEGER
    return IntLiteral.new(consume_token.value)
  elsif current_token.type == :LPAREN
    consume_token  # (
    expr = parse_expr
    expect(:RPAREN)
    return expr
  else
    error("Unexpected token")
  end
end
```

### 5.5 Abstract Syntax Tree (AST)

Der AST repräsentiert die **Struktur** des Programms:

```
"4 * 7 / (4 - 3)"

        BinaryOp(/)
       /          \
   BinaryOp(*)    BinaryOp(-)
   /      \       /       \
IntLit(4)  IntLit(7)  IntLit(4)  IntLit(3)
```

### 5.6 Ruby Crashkurs

Ruby ist ideal für Prototypen wegen seiner Ausdrucksstärke:

```ruby
# Variablen (kein Typ-System!)
x = 42
name = "Ruby"

# Arrays und Hashes
numbers = [1, 2, 3, 4, 5]
person = { name: "Alice", age: 30 }

# Klassen
class Token
  attr_reader :type, :value
  
  def initialize(type, value = nil)
    @type = type
    @value = value
  end
  
  def to_s
    "Token(#{@type}, #{@value})"
  end
end

# Blöcke und Iteratoren
numbers.each { |n| puts n * 2 }
numbers.map { |n| n ** 2 }
numbers.select { |n| n.even? }

# Pattern Matching (Ruby 3)
case token.type
when :INTEGER then parse_integer
when :PLUS, :MINUS then parse_binary_op
else raise "Unknown token"
end

# Reguläre Ausdrücke
"hello123".match?(/\d+/)  # => true
"123".scan(/\d/).map(&:to_i)  # => [1, 2, 3]
```

---

## 💻 Code-Beispiele

### Token Definition

```ruby
# token.rb

class Token
  attr_reader :type, :value, :pos
  
  TYPES = %i[
    INTEGER PLUS MINUS STAR SLASH
    LPAREN RPAREN
    EOF
  ]
  
  def initialize(type, value = nil, pos = 0)
    raise "Unknown token type: #{type}" unless TYPES.include?(type)
    @type = type
    @value = value
    @pos = pos
  end
  
  def to_s
    if @value
      "#{@type}(#{@value})"
    else
      @type.to_s
    end
  end
  
  def inspect
    to_s
  end
end
```

### Lexer

```ruby
# lexer.rb

class Lexer
  def initialize(source)
    @source = source
    @pos = 0
  end
  
  def tokenize
    tokens = []
    while @pos < @source.length
      tokens << next_token
    end
    tokens << Token.new(:EOF, nil, @pos)
    tokens
  end
  
  private
  
  def current_char
    @source[@pos]
  end
  
  def peek_char
    @source[@pos + 1]
  end
  
  def advance
    char = current_char
    @pos += 1
    char
  end
  
  def skip_whitespace
    while current_char && current_char.match?(/\s/)
      advance
    end
  end
  
  def read_integer
    start_pos = @pos
    value = ""
    while current_char && current_char.match?(/\d/)
      value += advance
    end
    Token.new(:INTEGER, value.to_i, start_pos)
  end
  
  def next_token
    skip_whitespace
    
    return Token.new(:EOF, nil, @pos) if @pos >= @source.length
    
    start_pos = @pos
    char = current_char
    
    case char
    when /\d/
      read_integer
    when '+'
      advance
      Token.new(:PLUS, nil, start_pos)
    when '-'
      advance
      Token.new(:MINUS, nil, start_pos)
    when '*'
      advance
      Token.new(:STAR, nil, start_pos)
    when '/'
      advance
      Token.new(:SLASH, nil, start_pos)
    when '('
      advance
      Token.new(:LPAREN, nil, start_pos)
    when ')'
      advance
      Token.new(:RPAREN, nil, start_pos)
    else
      raise "Unexpected character: '#{char}' at position #{@pos}"
    end
  end
end

# Test
if __FILE__ == $0
  source = "4 * 7 / (4 - 3)"
  lexer = Lexer.new(source)
  tokens = lexer.tokenize
  puts "Source: #{source}"
  puts "Tokens: #{tokens.map(&:to_s).join(', ')}"
end
```

### AST Nodes

```ruby
# ast.rb

# Basisklasse für alle AST-Knoten
class ASTNode
  def evaluate
    raise "Not implemented"
  end
  
  def to_s
    raise "Not implemented"
  end
end

# Integer-Literal
class IntLiteral < ASTNode
  attr_reader :value
  
  def initialize(value)
    @value = value
  end
  
  def evaluate
    @value
  end
  
  def to_s
    @value.to_s
  end
end

# Binärer Operator
class BinaryOp < ASTNode
  attr_reader :op, :left, :right
  
  def initialize(op, left, right)
    @op = op
    @left = left
    @right = right
  end
  
  def evaluate
    l = @left.evaluate
    r = @right.evaluate
    
    case @op
    when :PLUS  then l + r
    when :MINUS then l - r
    when :STAR  then l * r
    when :SLASH then l / r
    else raise "Unknown operator: #{@op}"
    end
  end
  
  def to_s
    "(#{@left} #{op_symbol} #{@right})"
  end
  
  private
  
  def op_symbol
    case @op
    when :PLUS  then '+'
    when :MINUS then '-'
    when :STAR  then '*'
    when :SLASH then '/'
    end
  end
end

# Unärer Operator (z.B. Negation)
class UnaryOp < ASTNode
  attr_reader :op, :operand
  
  def initialize(op, operand)
    @op = op
    @operand = operand
  end
  
  def evaluate
    val = @operand.evaluate
    case @op
    when :MINUS then -val
    else raise "Unknown unary operator: #{@op}"
    end
  end
  
  def to_s
    "(#{op_symbol}#{@operand})"
  end
  
  private
  
  def op_symbol
    @op == :MINUS ? '-' : @op.to_s
  end
end
```

### Parser

```ruby
# parser.rb

require_relative 'token'
require_relative 'ast'

class Parser
  def initialize(tokens)
    @tokens = tokens
    @pos = 0
  end
  
  def parse
    expr = parse_expression
    expect(:EOF)
    expr
  end
  
  private
  
  def current_token
    @tokens[@pos]
  end
  
  def peek_token
    @tokens[@pos + 1]
  end
  
  def advance
    token = current_token
    @pos += 1
    token
  end
  
  def expect(type)
    if current_token.type != type
      raise "Expected #{type} but got #{current_token.type} at position #{current_token.pos}"
    end
    advance
  end
  
  def match?(*types)
    types.include?(current_token.type)
  end
  
  # expr → term (('+' | '-') term)*
  def parse_expression
    left = parse_term
    
    while match?(:PLUS, :MINUS)
      op = advance.type
      right = parse_term
      left = BinaryOp.new(op, left, right)
    end
    
    left
  end
  
  # term → factor (('*' | '/') factor)*
  def parse_term
    left = parse_factor
    
    while match?(:STAR, :SLASH)
      op = advance.type
      right = parse_factor
      left = BinaryOp.new(op, left, right)
    end
    
    left
  end
  
  # factor → INTEGER | '-' factor | '(' expr ')'
  def parse_factor
    if match?(:INTEGER)
      IntLiteral.new(advance.value)
      
    elsif match?(:MINUS)
      advance
      operand = parse_factor
      UnaryOp.new(:MINUS, operand)
      
    elsif match?(:LPAREN)
      advance  # consume '('
      expr = parse_expression
      expect(:RPAREN)
      expr
      
    else
      raise "Unexpected token: #{current_token} at position #{current_token.pos}"
    end
  end
end

# Test
if __FILE__ == $0
  require_relative 'lexer'
  
  source = "4 * 7 / (4 - 3)"
  puts "Source: #{source}"
  
  lexer = Lexer.new(source)
  tokens = lexer.tokenize
  puts "Tokens: #{tokens.map(&:to_s).join(', ')}"
  
  parser = Parser.new(tokens)
  ast = parser.parse
  puts "AST: #{ast}"
  puts "Result: #{ast.evaluate}"
end
```

### Main Program

```ruby
# calc.rb - Expression Calculator

require_relative 'lexer'
require_relative 'parser'

def evaluate(source)
  lexer = Lexer.new(source)
  tokens = lexer.tokenize
  parser = Parser.new(tokens)
  ast = parser.parse
  ast.evaluate
end

# REPL (Read-Eval-Print Loop)
def repl
  puts "Expression Calculator"
  puts "Type 'quit' to exit"
  puts
  
  loop do
    print "> "
    input = gets&.strip
    
    break if input.nil? || input == 'quit'
    next if input.empty?
    
    begin
      result = evaluate(input)
      puts "= #{result}"
    rescue => e
      puts "Error: #{e.message}"
    end
  end
end

if __FILE__ == $0
  if ARGV.empty?
    repl
  else
    puts evaluate(ARGV.join(' '))
  end
end
```

---

## 📖 Weiterführende Ressourcen

### Compilerbau
- **Crafting Interpreters** - [craftinginterpreters.com](https://craftinginterpreters.com/) (FREE, excellent!)
- **Writing An Interpreter In Go** - Thorsten Ball
- **Dragon Book** - Compilers: Principles, Techniques, and Tools (Klassiker)

### Ruby
- **Ruby in 20 Minutes** - [ruby-lang.org/en/documentation/quickstart](https://www.ruby-lang.org/en/documentation/quickstart/)
- **Learn Ruby Online** - [learnrubyonline.org](https://www.learnrubyonline.org/)
- **Ruby Koans** - Interaktives Lernen

### Theoretische Grundlagen
- **Reguläre Ausdrücke** - [regexr.com](https://regexr.com/)
- **Formale Grammatiken** - Chomsky-Hierarchie
- **Recursive Descent Parsing** - [craftinginterpreters.com/parsing-expressions.html](https://craftinginterpreters.com/parsing-expressions.html)

---

## 🧠 Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Lexer | Text → Tokens (reguläre Ausdrücke) |
| Parser | Tokens → AST (Grammatikregeln) |
| AST | Baumstruktur des Programms |
| Recursive Descent | Eine Funktion pro Grammatikregel |
| Operator-Präzedenz | Durch Grammatikebenen definiert |

**Grammatik-Regel → Parser-Funktion:**
```
expr → term (('+' | '-') term)*    →  parse_expression()
term → factor (('*' | '/') factor)* →  parse_term()  
factor → INTEGER | '(' expr ')'    →  parse_factor()
```

---

*Weiter zu den Übungen → `uebungen/uebung-05.md`*
