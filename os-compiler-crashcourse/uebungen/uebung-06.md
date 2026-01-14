# Übung 6: Bytecode Compiler in C++

**Ziel**: Portiere den Expression-Compiler nach C++ und füge Bytecode-Erzeugung hinzu.

---

## 🛠️ Setup

```bash
mkdir -p ~/os-dev/day6/compiler-cpp
cd ~/os-dev/day6/compiler-cpp

# Prüfe C++ Compiler
clang++ --version  # oder g++ --version
```

---

## Aufgabe 1: C++ Basics (⏱️ 30 min)

**Datei**: `cpp_warmup.cpp`

```cpp
#include <iostream>
#include <vector>
#include <string>
#include <memory>

// Aufgabe 1a: Einfache Klasse
class Rectangle {
    double width, height;
public:
    // TODO: Konstruktor
    Rectangle(double w, double h) : /* TODO: Initialisierungsliste */ {}
    
    // TODO: Getter
    double get_width() const { /* TODO */ }
    double get_height() const { /* TODO */ }
    
    // TODO: Fläche berechnen
    double area() const { /* TODO */ }
    
    // TODO: Umfang berechnen
    double perimeter() const { /* TODO */ }
};

// Aufgabe 1b: Vererbung
class Shape {
public:
    virtual ~Shape() = default;
    virtual double area() const = 0;  // pure virtual
    virtual std::string name() const = 0;
};

class Circle : public Shape {
    double radius;
public:
    Circle(double r) : radius(r) {}
    
    // TODO: Implementiere area() und name()
};

class Square : public Shape {
    double side;
public:
    // TODO: Implementiere
};

// Aufgabe 1c: Smart Pointers
void smart_pointer_demo() {
    // TODO: Erstelle unique_ptr auf Shape
    // std::unique_ptr<Shape> shape = std::make_unique<Circle>(5.0);
    
    // TODO: Gib Fläche aus
    
    // TODO: Vektor von Shapes
    std::vector<std::unique_ptr<Shape>> shapes;
    // shapes.push_back(std::make_unique<Circle>(3.0));
    // shapes.push_back(std::make_unique<Square>(4.0));
    
    // TODO: Iteriere und gib alle Flächen aus
}

// Aufgabe 1d: Standard Library
void stl_demo() {
    std::vector<int> nums = {5, 2, 8, 1, 9, 3};
    
    // TODO: Sortiere mit std::sort
    
    // TODO: Finde Maximum mit std::max_element
    
    // TODO: Summiere mit std::accumulate
    
    // TODO: Gib Ergebnisse aus
}

int main() {
    // Test Rectangle
    Rectangle r(3, 4);
    std::cout << "Rectangle: " << r.area() << " / " << r.perimeter() << std::endl;
    
    smart_pointer_demo();
    stl_demo();
    
    return 0;
}
```

---

## Aufgabe 2: Token & Lexer in C++ (⏱️ 45 min)

**Datei**: `token.hpp`

```cpp
#ifndef TOKEN_HPP
#define TOKEN_HPP

#include <string>
#include <variant>

enum class TokenType {
    INTEGER,
    PLUS, MINUS, STAR, SLASH,
    LPAREN, RPAREN,
    END_OF_FILE
};

struct Token {
    TokenType type;
    std::variant<std::monostate, int> value;  // int für INTEGER
    size_t pos;
    
    Token(TokenType t, size_t p = 0) : type(t), value(std::monostate{}), pos(p) {}
    Token(TokenType t, int v, size_t p = 0) : type(t), value(v), pos(p) {}
    
    std::string to_string() const;
};

const char* token_type_name(TokenType type);

#endif
```

**Datei**: `lexer.hpp`

```cpp
#ifndef LEXER_HPP
#define LEXER_HPP

#include <string>
#include <vector>
#include "token.hpp"

class Lexer {
    std::string source;
    size_t pos = 0;
    
public:
    explicit Lexer(const std::string& src) : source(src) {}
    
    std::vector<Token> tokenize();
    
private:
    char current() const;
    char peek() const;
    char advance();
    void skip_whitespace();
    Token read_integer();
    Token next_token();
};

#endif
```

**Datei**: `lexer.cpp`

```cpp
#include "lexer.hpp"
#include <stdexcept>
#include <cctype>

// TODO: Implementiere alle Methoden

char Lexer::current() const {
    // TODO: Return aktuelles Zeichen oder '\0'
}

char Lexer::peek() const {
    // TODO
}

char Lexer::advance() {
    // TODO
}

void Lexer::skip_whitespace() {
    // TODO
}

Token Lexer::read_integer() {
    size_t start = pos;
    int value = 0;
    // TODO: Parse Integer
    return Token(TokenType::INTEGER, value, start);
}

Token Lexer::next_token() {
    skip_whitespace();
    
    if (pos >= source.length()) {
        return Token(TokenType::END_OF_FILE, pos);
    }
    
    size_t start = pos;
    char c = current();
    
    if (std::isdigit(c)) {
        return read_integer();
    }
    
    // TODO: Handle +, -, *, /, (, )
    
    throw std::runtime_error("Unexpected character: " + std::string(1, c));
}

std::vector<Token> Lexer::tokenize() {
    std::vector<Token> tokens;
    // TODO: Sammle alle Tokens
    return tokens;
}
```

---

## Aufgabe 3: AST in C++ (⏱️ 30 min)

**Datei**: `ast.hpp`

```cpp
#ifndef AST_HPP
#define AST_HPP

#include <memory>
#include <string>
#include "token.hpp"

// Basis-Klasse
class ASTNode {
public:
    virtual ~ASTNode() = default;
    virtual int evaluate() const = 0;
    virtual std::string to_string() const = 0;
};

using ASTPtr = std::unique_ptr<ASTNode>;

// Integer Literal
class IntLiteral : public ASTNode {
public:
    int value;
    
    explicit IntLiteral(int v) : value(v) {}
    
    int evaluate() const override { return value; }
    std::string to_string() const override { return std::to_string(value); }
};

// Binary Operation
class BinaryOp : public ASTNode {
public:
    TokenType op;
    ASTPtr left;
    ASTPtr right;
    
    BinaryOp(TokenType o, ASTPtr l, ASTPtr r)
        : op(o), left(std::move(l)), right(std::move(r)) {}
    
    int evaluate() const override {
        int l = left->evaluate();
        int r = right->evaluate();
        // TODO: Implementiere basierend auf op
    }
    
    std::string to_string() const override {
        // TODO: "(left op right)"
    }
};

// Unary Operation
class UnaryOp : public ASTNode {
public:
    TokenType op;
    ASTPtr operand;
    
    // TODO: Konstruktor und Methoden
};

#endif
```

---

## Aufgabe 4: Parser in C++ (⏱️ 45 min)

**Datei**: `parser.hpp`

```cpp
#ifndef PARSER_HPP
#define PARSER_HPP

#include <vector>
#include "token.hpp"
#include "ast.hpp"

class Parser {
    std::vector<Token> tokens;
    size_t pos = 0;
    
public:
    explicit Parser(std::vector<Token> toks) : tokens(std::move(toks)) {}
    
    ASTPtr parse();
    
private:
    Token& current();
    Token advance();
    Token expect(TokenType type);
    bool match(TokenType type);
    bool match(std::initializer_list<TokenType> types);
    
    ASTPtr parse_expression();
    ASTPtr parse_term();
    ASTPtr parse_factor();
};

#endif
```

**Datei**: `parser.cpp`

```cpp
#include "parser.hpp"
#include <stdexcept>

Token& Parser::current() {
    return tokens[pos];
}

Token Parser::advance() {
    Token tok = current();
    if (pos < tokens.size() - 1) pos++;
    return tok;
}

Token Parser::expect(TokenType type) {
    if (current().type != type) {
        throw std::runtime_error("Unexpected token");
    }
    return advance();
}

bool Parser::match(TokenType type) {
    return current().type == type;
}

bool Parser::match(std::initializer_list<TokenType> types) {
    for (auto t : types) {
        if (current().type == t) return true;
    }
    return false;
}

ASTPtr Parser::parse() {
    auto expr = parse_expression();
    expect(TokenType::END_OF_FILE);
    return expr;
}

ASTPtr Parser::parse_expression() {
    auto left = parse_term();
    
    while (match({TokenType::PLUS, TokenType::MINUS})) {
        TokenType op = advance().type;
        auto right = parse_term();
        left = std::make_unique<BinaryOp>(op, std::move(left), std::move(right));
    }
    
    return left;
}

ASTPtr Parser::parse_term() {
    // TODO: Analog zu expression, aber mit STAR/SLASH
}

ASTPtr Parser::parse_factor() {
    // TODO: INTEGER, MINUS, LPAREN
}
```

---

## Aufgabe 5: Bytecode (⏱️ 30 min)

**Datei**: `bytecode.hpp`

```cpp
#ifndef BYTECODE_HPP
#define BYTECODE_HPP

#include <cstdint>
#include <vector>
#include <iostream>

enum class Opcode : uint8_t {
    PUSH,
    ADD, SUB, MUL, DIV,
    NEG,
    HALT
};

struct Instruction {
    Opcode op;
    int32_t operand;
    
    Instruction(Opcode o, int32_t v = 0) : op(o), operand(v) {}
};

class Chunk {
public:
    std::vector<Instruction> code;
    
    void emit(Opcode op) {
        code.emplace_back(op);
    }
    
    void emit(Opcode op, int32_t value) {
        code.emplace_back(op, value);
    }
    
    void disassemble() const {
        for (size_t i = 0; i < code.size(); i++) {
            const auto& instr = code[i];
            std::cout << i << ": ";
            // TODO: Formatierte Ausgabe
        }
    }
};

#endif
```

---

## Aufgabe 6: Compiler (⏱️ 45 min)

**Datei**: `compiler.hpp`

```cpp
#ifndef COMPILER_HPP
#define COMPILER_HPP

#include "ast.hpp"
#include "bytecode.hpp"

class Compiler {
    Chunk chunk;
    
public:
    Chunk compile(ASTNode* root) {
        emit_code(root);
        chunk.emit(Opcode::HALT);
        return std::move(chunk);
    }
    
private:
    void emit_code(ASTNode* node) {
        // TODO: Dynamisches Type-Checking mit dynamic_cast
        // Post-Order Traversierung für Binary/Unary Ops
        
        if (auto* intLit = dynamic_cast<IntLiteral*>(node)) {
            // TODO: PUSH
        }
        else if (auto* binOp = dynamic_cast<BinaryOp*>(node)) {
            // TODO: Erst Kinder, dann Operator
        }
        else if (auto* unOp = dynamic_cast<UnaryOp*>(node)) {
            // TODO: Erst Operand, dann NEG
        }
    }
};

#endif
```

---

## Aufgabe 7: Virtual Machine (⏱️ 45 min)

**Datei**: `vm.hpp`

```cpp
#ifndef VM_HPP
#define VM_HPP

#include "bytecode.hpp"
#include <stack>
#include <stdexcept>

class VM {
    std::stack<int32_t> stack;
    
public:
    int32_t execute(const Chunk& chunk) {
        size_t ip = 0;
        
        while (ip < chunk.code.size()) {
            const auto& instr = chunk.code[ip];
            
            switch (instr.op) {
                case Opcode::PUSH:
                    stack.push(instr.operand);
                    break;
                    
                case Opcode::ADD: {
                    // TODO
                    break;
                }
                
                case Opcode::SUB: {
                    // TODO: Beachte Reihenfolge! a - b
                    break;
                }
                
                case Opcode::MUL: {
                    // TODO
                    break;
                }
                
                case Opcode::DIV: {
                    // TODO: Division by zero check!
                    break;
                }
                
                case Opcode::NEG: {
                    // TODO
                    break;
                }
                
                case Opcode::HALT:
                    return stack.empty() ? 0 : stack.top();
            }
            
            ip++;
        }
        
        return stack.empty() ? 0 : stack.top();
    }
    
private:
    int32_t pop() {
        if (stack.empty()) {
            throw std::runtime_error("Stack underflow");
        }
        int32_t val = stack.top();
        stack.pop();
        return val;
    }
};

#endif
```

---

## Aufgabe 8: Main & Build (⏱️ 30 min)

**Datei**: `main.cpp`

```cpp
#include <iostream>
#include <string>
#include "lexer.hpp"
#include "parser.hpp"
#include "compiler.hpp"
#include "vm.hpp"

int main(int argc, char* argv[]) {
    // TODO: Wenn Argumente gegeben: Evaluiere und gib Ergebnis aus
    // TODO: Sonst: REPL starten
    
    std::cout << "C++ Expression Compiler" << std::endl;
    std::cout << "Type 'quit' to exit, 'debug' to toggle bytecode output" << std::endl;
    
    bool debug = false;
    std::string input;
    
    while (true) {
        std::cout << "> ";
        if (!std::getline(std::cin, input)) break;
        
        if (input == "quit") break;
        if (input == "debug") {
            debug = !debug;
            std::cout << "Debug: " << (debug ? "ON" : "OFF") << std::endl;
            continue;
        }
        if (input.empty()) continue;
        
        try {
            // TODO: Lexer → Parser → Compiler → VM
            
        } catch (const std::exception& e) {
            std::cerr << "Error: " << e.what() << std::endl;
        }
    }
    
    return 0;
}
```

**Datei**: `Makefile`

```makefile
CXX = clang++
CXXFLAGS = -std=c++17 -Wall -Wextra -g

SRCS = main.cpp lexer.cpp parser.cpp token.cpp
OBJS = $(SRCS:.cpp=.o)

calc: $(OBJS)
	$(CXX) $(CXXFLAGS) -o $@ $^

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<

clean:
	rm -f calc *.o

test: calc
	./calc "4 * 7 / (4 - 3)"
	./calc "1 + 2 + 3 + 4 + 5"
	./calc "-5 + 10"

.PHONY: clean test
```

---

## 🎯 Bonusaufgaben

### Bonus A: Optimierung - Constant Folding

Wenn beide Operanden Konstanten sind, berechne zur Compile-Zeit:

```cpp
// Statt:
PUSH 4
PUSH 7
MUL

// Optimiert:
PUSH 28
```

### Bonus B: Tracing Execution

Zeige Stack-Zustand bei jedem Schritt:

```
0: PUSH 4      Stack: [4]
1: PUSH 7      Stack: [4, 7]
2: MUL         Stack: [28]
3: HALT        Result: 28
```

### Bonus C: Bytecode Serialisierung

Speichere Bytecode in Datei, lade und führe aus.

---

## ✅ Checkliste

- [ ] C++ Basics verstanden (Klassen, Smart Pointers)
- [ ] Lexer tokenisiert korrekt
- [ ] Parser erzeugt AST
- [ ] Compiler generiert Bytecode
- [ ] VM führt Bytecode aus
- [ ] REPL funktioniert

**Morgen**: GUI-Basics und Rust!

---

*Weiter zu Tag 7 → `vorlesungen/tag-07-gui-rust.md`*
