# Übung 5: Expression Compiler in Ruby

**Ziel**: Baue einen vollständigen Expression-Rechner mit Lexer, Parser, und Evaluator.

---

## 🛠️ Setup

```bash
# Erstelle Arbeitsverzeichnis
mkdir -p ~/os-dev/day5/calculator
cd ~/os-dev/day5/calculator

# Prüfe Ruby
ruby --version  # Sollte 2.7+ sein

# Erstelle Dateien
touch token.rb lexer.rb ast.rb parser.rb calc.rb
```

---

## Aufgabe 1: Ruby Basics (⏱️ 30 min)

Wärme dich mit Ruby auf.

**Datei**: `ruby_warmup.rb`

```ruby
# === Aufgabe 1a: Klassen und Objekte ===
# Erstelle eine Klasse "Point" mit x und y Koordinaten
# und eine Methode "distance_to(other)" die die Distanz berechnet

class Point
  # TODO: Implementiere!
  # attr_reader :x, :y
  # def initialize(x, y)
  # def distance_to(other)
end

# Test
p1 = Point.new(0, 0)
p2 = Point.new(3, 4)
puts p1.distance_to(p2)  # Sollte 5.0 sein

# === Aufgabe 1b: Arrays und Blöcke ===
# Gegeben ist ein Array von Zahlen
numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

# TODO: Benutze Blöcke um folgende zu berechnen:
# - Summe aller Zahlen (inject)
sum = # ???
puts "Sum: #{sum}"  # 55

# - Nur gerade Zahlen (select)
evens = # ???
puts "Evens: #{evens}"  # [2, 4, 6, 8, 10]

# - Quadrate aller Zahlen (map)
squares = # ???
puts "Squares: #{squares}"  # [1, 4, 9, 16, ...]

# === Aufgabe 1c: Pattern Matching ===
# Schreibe eine Funktion die verschiedene Eingaben behandelt

def describe(value)
  case value
  # TODO: Pattern matching für verschiedene Typen
  # when Integer: "Integer: #{value}"
  # when String: "String with #{value.length} chars"
  # when Array: "Array with #{value.size} elements"
  # else: "Unknown"
  end
end

puts describe(42)
puts describe("hello")
puts describe([1, 2, 3])

# === Aufgabe 1d: Reguläre Ausdrücke ===
text = "The answer is 42 and 7 * 6 equals 42"

# TODO: Finde alle Zahlen im Text
numbers = # text.scan(???)
puts "Numbers found: #{numbers}"  # ["42", "7", "6", "42"]

# TODO: Ersetze Zahlen mit "X"
censored = # text.gsub(???, "X")
puts censored  # "The answer is X and X * X equals X"
```

---

## Aufgabe 2: Token Klasse (⏱️ 20 min)

**Datei**: `token.rb`

```ruby
# Token repräsentiert ein lexikalisches Element

class Token
  # TODO: Definiere Token-Typen als Konstante
  TYPES = %i[
    # Literale
    INTEGER
    # Operatoren
    # TODO: PLUS, MINUS, STAR, SLASH
    # Klammern
    # TODO: LPAREN, RPAREN
    # Ende
    EOF
  ]
  
  attr_reader :type, :value, :pos
  
  def initialize(type, value = nil, pos = 0)
    # TODO: Validiere dass type in TYPES ist
    # TODO: Speichere @type, @value, @pos
  end
  
  def to_s
    # TODO: Formatiere als "TYPE" oder "TYPE(value)"
  end
  
  # Für Debugging
  def inspect
    to_s
  end
end

# Tests
if __FILE__ == $0
  t1 = Token.new(:INTEGER, 42, 0)
  puts t1  # INTEGER(42)
  
  t2 = Token.new(:PLUS, nil, 2)
  puts t2  # PLUS
  
  begin
    Token.new(:UNKNOWN)
  rescue => e
    puts "Expected error: #{e.message}"
  end
end
```

---

## Aufgabe 3: Lexer (⏱️ 45 min)

**Datei**: `lexer.rb`

```ruby
require_relative 'token'

class Lexer
  def initialize(source)
    @source = source
    @pos = 0
  end
  
  # Tokenisiere gesamten Input
  def tokenize
    tokens = []
    # TODO: Sammle alle Tokens, ende mit EOF
    tokens
  end
  
  private
  
  # Aktuelles Zeichen (oder nil am Ende)
  def current_char
    # TODO
  end
  
  # Nächstes Zeichen anschauen ohne zu konsumieren
  def peek_char
    # TODO
  end
  
  # Zeichen konsumieren und @pos erhöhen
  def advance
    # TODO
  end
  
  # Überspringe Whitespace
  def skip_whitespace
    # TODO: Solange aktuelles Zeichen Whitespace ist, advance
  end
  
  # Lese eine Integer-Zahl
  def read_integer
    start_pos = @pos
    value = ""
    # TODO: Sammle Ziffern
    Token.new(:INTEGER, value.to_i, start_pos)
  end
  
  # Lese nächstes Token
  def next_token
    skip_whitespace
    
    return Token.new(:EOF, nil, @pos) if @pos >= @source.length
    
    start_pos = @pos
    char = current_char
    
    # TODO: Pattern matching für verschiedene Zeichen
    case char
    when /\d/
      read_integer
    when '+'
      advance
      Token.new(:PLUS, nil, start_pos)
    # TODO: Weitere Fälle: - * / ( )
    else
      raise "Unexpected character: '#{char}' at position #{@pos}"
    end
  end
end

# Tests
if __FILE__ == $0
  test_cases = [
    "42",
    "1 + 2",
    "4 * 7 / (4 - 3)",
    "((1 + 2) * 3)"
  ]
  
  test_cases.each do |source|
    lexer = Lexer.new(source)
    tokens = lexer.tokenize
    puts "#{source.inspect} => #{tokens.join(', ')}"
  end
end
```

**Erwartete Ausgabe:**
```
"42" => INTEGER(42), EOF
"1 + 2" => INTEGER(1), PLUS, INTEGER(2), EOF
"4 * 7 / (4 - 3)" => INTEGER(4), STAR, INTEGER(7), SLASH, LPAREN, INTEGER(4), MINUS, INTEGER(3), RPAREN, EOF
```

---

## Aufgabe 4: AST Knoten (⏱️ 30 min)

**Datei**: `ast.rb`

```ruby
# Basisklasse für AST-Knoten
class ASTNode
  # Jeder Knoten muss evaluate() haben
  def evaluate
    raise NotImplementedError, "#{self.class} must implement #evaluate"
  end
  
  # Für Debugging: Baum darstellen
  def to_tree(indent = 0)
    raise NotImplementedError
  end
end

# Integer-Literal (z.B. 42)
class IntLiteral < ASTNode
  attr_reader :value
  
  def initialize(value)
    @value = value
  end
  
  def evaluate
    # TODO: Gib einfach den Wert zurück
  end
  
  def to_s
    @value.to_s
  end
  
  def to_tree(indent = 0)
    "  " * indent + "Int(#{@value})"
  end
end

# Binärer Operator (z.B. 1 + 2)
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
    
    # TODO: Berechne basierend auf @op
    # :PLUS => l + r
    # :MINUS => l - r
    # :STAR => l * r
    # :SLASH => l / r (Integer-Division!)
  end
  
  def to_s
    "(#{@left} #{op_symbol} #{@right})"
  end
  
  def to_tree(indent = 0)
    lines = []
    lines << "  " * indent + "BinaryOp(#{op_symbol})"
    lines << @left.to_tree(indent + 1)
    lines << @right.to_tree(indent + 1)
    lines.join("\n")
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

# Unärer Operator (z.B. -42)
class UnaryOp < ASTNode
  attr_reader :op, :operand
  
  def initialize(op, operand)
    @op = op
    @operand = operand
  end
  
  def evaluate
    val = @operand.evaluate
    # TODO: Für :MINUS => -val
  end
  
  def to_s
    "(#{@op == :MINUS ? '-' : @op}#{@operand})"
  end
  
  def to_tree(indent = 0)
    lines = []
    lines << "  " * indent + "UnaryOp(#{@op})"
    lines << @operand.to_tree(indent + 1)
    lines.join("\n")
  end
end

# Tests
if __FILE__ == $0
  # 2 + 3
  ast1 = BinaryOp.new(:PLUS, IntLiteral.new(2), IntLiteral.new(3))
  puts ast1
  puts "= #{ast1.evaluate}"
  puts ast1.to_tree
  puts
  
  # 4 * 7 / 2
  ast2 = BinaryOp.new(:SLASH,
    BinaryOp.new(:STAR, IntLiteral.new(4), IntLiteral.new(7)),
    IntLiteral.new(2)
  )
  puts ast2
  puts "= #{ast2.evaluate}"
  puts ast2.to_tree
end
```

---

## Aufgabe 5: Parser (⏱️ 60 min)

**Datei**: `parser.rb`

```ruby
require_relative 'token'
require_relative 'ast'

class Parser
  def initialize(tokens)
    @tokens = tokens
    @pos = 0
  end
  
  # Starte Parsing
  def parse
    result = parse_expression
    expect(:EOF)
    result
  end
  
  private
  
  # Aktuelles Token
  def current_token
    @tokens[@pos]
  end
  
  # Token konsumieren und weitergehen
  def advance
    token = current_token
    @pos += 1
    token
  end
  
  # Erwarte bestimmten Token-Typ
  def expect(type)
    if current_token.type != type
      raise "Expected #{type} but got #{current_token.type} at position #{current_token.pos}"
    end
    advance
  end
  
  # Prüfe ob aktuelles Token einer der Typen ist
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
    # TODO: Analog zu parse_expression, aber mit STAR und SLASH
  end
  
  # factor → INTEGER | '-' factor | '(' expr ')'
  def parse_factor
    if match?(:INTEGER)
      # TODO: Integer-Literal zurückgeben
      
    elsif match?(:MINUS)
      # TODO: Unäres Minus
      
    elsif match?(:LPAREN)
      # TODO: Geklammerten Ausdruck parsen
      
    else
      raise "Unexpected token: #{current_token} at position #{current_token.pos}"
    end
  end
end

# Tests
if __FILE__ == $0
  require_relative 'lexer'
  
  test_cases = [
    "42",
    "1 + 2",
    "1 + 2 * 3",
    "4 * 7 / (4 - 3)",
    "-5 + 3",
    "((1 + 2) * (3 + 4))"
  ]
  
  test_cases.each do |source|
    puts "=" * 40
    puts "Source: #{source}"
    
    lexer = Lexer.new(source)
    tokens = lexer.tokenize
    puts "Tokens: #{tokens.join(', ')}"
    
    parser = Parser.new(tokens)
    ast = parser.parse
    puts "AST: #{ast}"
    puts ast.to_tree
    puts "Result: #{ast.evaluate}"
  end
end
```

---

## Aufgabe 6: Calculator REPL (⏱️ 30 min)

**Datei**: `calc.rb`

```ruby
#!/usr/bin/env ruby

require_relative 'lexer'
require_relative 'parser'

def evaluate(source)
  lexer = Lexer.new(source)
  tokens = lexer.tokenize
  parser = Parser.new(tokens)
  ast = parser.parse
  ast.evaluate
end

def repl
  puts "╔═════════════════════════════════════╗"
  puts "║   Expression Calculator v1.0        ║"
  puts "║   Type 'quit' or Ctrl+D to exit     ║"
  puts "╚═════════════════════════════════════╝"
  puts
  
  loop do
    print "calc> "
    input = gets&.strip
    
    break if input.nil? || input.downcase == 'quit'
    next if input.empty?
    
    begin
      result = evaluate(input)
      puts "= #{result}"
    rescue => e
      puts "Error: #{e.message}"
    end
    puts
  end
  
  puts "Goodbye!"
end

# Main
if __FILE__ == $0
  if ARGV.empty?
    repl
  else
    expression = ARGV.join(' ')
    puts evaluate(expression)
  end
end
```

**Test:**
```bash
ruby calc.rb "4 * 7 / (4 - 3)"
# 28

ruby calc.rb
# REPL startet
```

---

## Aufgabe 7: Erweiterungen (⏱️ 60 min)

### 7a: Floating Point Support

Erweitere den Lexer um Dezimalzahlen:

```ruby
# Im Lexer: read_number statt read_integer
def read_number
  value = ""
  has_dot = false
  
  while current_char && (current_char.match?(/\d/) || current_char == '.')
    if current_char == '.'
      break if has_dot  # Zweiter Punkt = Ende
      has_dot = true
    end
    value += advance
  end
  
  Token.new(has_dot ? :FLOAT : :INTEGER, 
            has_dot ? value.to_f : value.to_i, 
            @pos - value.length)
end
```

### 7b: Variables

Erweitere um Variablen-Zuweisung und -Verwendung:

```ruby
# Neue Tokens
:IDENT      # Identifier
:ASSIGN     # =

# Neue AST-Knoten
class Variable < ASTNode
  # ...
end

class Assignment < ASTNode
  # ...
end

# Context für Variablen
class Context
  def initialize
    @variables = {}
  end
  
  def get(name)
    @variables[name] || raise("Undefined variable: #{name}")
  end
  
  def set(name, value)
    @variables[name] = value
  end
end
```

### 7c: Modulo und Power

Füge `%` und `**` Operatoren hinzu.

```ruby
# New tokens: PERCENT, POWER

# Grammatik erweitern:
# power → factor ('**' power)?   (rechts-assoziativ!)
```

---

## 🎯 Bonusaufgaben

### Bonus A: Error Recovery

Bei Fehlern: Zeige Position im Quellcode an:
```
Error at position 5:
  1 + * 3
      ^
Unexpected token: STAR
```

### Bonus B: Pretty Printer

Formatiere AST als schönen Output:
```
1 + 2 * 3
↓
(+ 1 (* 2 3))
```

### Bonus C: AST Visualizer

Generiere GraphViz DOT-Output für den AST:
```ruby
def to_dot
  # ...generates .dot file that can be visualized
end
```

---

## ✅ Checkliste

- [ ] Ruby Basics verstanden
- [ ] Token-Klasse implementiert
- [ ] Lexer tokenisiert korrekt
- [ ] AST-Knoten evaluieren sich
- [ ] Parser erzeugt korrekten AST
- [ ] REPL funktioniert
- [ ] Mindestens eine Erweiterung

**Morgen**: Code Generation in C++!

---

*Weiter zu Tag 6 → `vorlesungen/tag-06-codegen.md`*
