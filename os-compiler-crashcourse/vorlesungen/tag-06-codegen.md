# Tag 6: Code Generation – AST zu Bytecode in C++

**Ziel**: Erweitere deinen Compiler um Bytecode-Erzeugung und eine Stack-basierte VM.

---

## 📚 Theorie

### 6.1 Von AST zu Ausführung

```
┌─────────────────────────────────────────────────────────────────┐
│  Ausführungsstrategien                                          │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. Tree-Walking Interpreter (was wir in Ruby gemacht haben)    │
│     + Einfach zu implementieren                                 │
│     - Langsam (Pointer-Chasing, Method Calls)                   │
│                                                                 │
│  2. Bytecode Compiler + Virtual Machine                         │
│     + Schneller (lineare Ausführung)                            │
│     + Portabel (Bytecode läuft auf jeder VM)                    │
│     - Mehr Aufwand (zwei Phasen)                                │
│                                                                 │
│  3. Native Code Compiler (z.B. LLVM)                            │
│     + Maximale Geschwindigkeit                                  │
│     - Am komplexesten                                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### 6.2 Stack-basierte Virtual Machine

```
Bytecode:    PUSH 4 | PUSH 7 | MUL | PUSH 1 | SUB

Ausführung:
  PUSH 4    Stack: [4]
  PUSH 7    Stack: [4, 7]
  MUL       Stack: [28]      ← 4 * 7 = 28
  PUSH 1    Stack: [28, 1]
  SUB       Stack: [27]      ← 28 - 1 = 27

Ergebnis: 27
```

### 6.3 Bytecode Instruktionen

```
┌──────────┬────────────────────────────────────────┐
│ Opcode   │ Beschreibung                           │
├──────────┼────────────────────────────────────────┤
│ PUSH n   │ Push Konstante n auf Stack             │
│ POP      │ Entferne oberstes Element              │
│ ADD      │ Pop 2 Werte, push Summe                │
│ SUB      │ Pop 2 Werte, push Differenz (a-b)      │
│ MUL      │ Pop 2 Werte, push Produkt              │
│ DIV      │ Pop 2 Werte, push Quotient             │
│ NEG      │ Negiere oberstes Element               │
│ HALT     │ Beende Programm                        │
└──────────┴────────────────────────────────────────┘
```

### 6.4 AST zu Bytecode

Post-Order Traversierung des AST:

```
      (-)
     /   \
   (*)    1
   / \
  4   7

Traversierung:
1. Besuche links: (*)
   1a. Besuche links: 4   → PUSH 4
   1b. Besuche rechts: 7  → PUSH 7
   1c. Emit MUL
2. Besuche rechts: 1      → PUSH 1
3. Emit SUB

Bytecode: PUSH 4, PUSH 7, MUL, PUSH 1, SUB
```

### 6.5 C++ Crashkurs

C++ ist C mit Klassen und mehr:

```cpp
#include <iostream>
#include <vector>
#include <memory>
#include <string>

// Klassen
class Token {
public:
    enum class Type { INTEGER, PLUS, MINUS, STAR, SLASH };
    
    Type type;
    int value;  // für INTEGER
    
    Token(Type t, int v = 0) : type(t), value(v) {}
};

// Vererbung
class ASTNode {
public:
    virtual ~ASTNode() = default;
    virtual int evaluate() = 0;  // pure virtual (abstrakt)
};

class IntLiteral : public ASTNode {
    int value;
public:
    IntLiteral(int v) : value(v) {}
    int evaluate() override { return value; }
};

// Smart Pointers (kein manuelles delete!)
std::unique_ptr<ASTNode> node = std::make_unique<IntLiteral>(42);

// Vectors (dynamische Arrays)
std::vector<int> numbers = {1, 2, 3, 4, 5};
for (int n : numbers) {
    std::cout << n << std::endl;
}

// Optional (für nullable Werte)
#include <optional>
std::optional<int> find_value();

// Variant (union aber typesafe)
#include <variant>
std::variant<int, std::string> data = 42;
```

---

## 💻 Code-Beispiele

### Bytecode Definition (C++)

```cpp
// bytecode.hpp

#ifndef BYTECODE_HPP
#define BYTECODE_HPP

#include <cstdint>
#include <vector>
#include <string>

// Opcode Definitionen
enum class Opcode : uint8_t {
    PUSH,   // PUSH <value>
    POP,
    ADD,
    SUB,
    MUL,
    DIV,
    NEG,
    HALT
};

// Eine Instruktion
struct Instruction {
    Opcode op;
    int32_t operand;  // für PUSH
    
    Instruction(Opcode o, int32_t val = 0) : op(o), operand(val) {}
};

// Bytecode Chunk
class Chunk {
public:
    std::vector<Instruction> code;
    
    void emit(Opcode op) {
        code.push_back(Instruction(op));
    }
    
    void emit(Opcode op, int32_t value) {
        code.push_back(Instruction(op, value));
    }
    
    void disassemble() const;
};

// Instruction Namen für Debug
inline const char* opcode_name(Opcode op) {
    switch (op) {
        case Opcode::PUSH: return "PUSH";
        case Opcode::POP:  return "POP";
        case Opcode::ADD:  return "ADD";
        case Opcode::SUB:  return "SUB";
        case Opcode::MUL:  return "MUL";
        case Opcode::DIV:  return "DIV";
        case Opcode::NEG:  return "NEG";
        case Opcode::HALT: return "HALT";
        default: return "???";
    }
}

#endif
```

### Bytecode Compiler

```cpp
// compiler.hpp

#ifndef COMPILER_HPP
#define COMPILER_HPP

#include "ast.hpp"
#include "bytecode.hpp"

class Compiler {
    Chunk chunk;
    
public:
    Chunk compile(ASTNode* node) {
        visit(node);
        chunk.emit(Opcode::HALT);
        return chunk;
    }
    
private:
    void visit(ASTNode* node) {
        if (auto* intLit = dynamic_cast<IntLiteral*>(node)) {
            chunk.emit(Opcode::PUSH, intLit->value);
            
        } else if (auto* binOp = dynamic_cast<BinaryOp*>(node)) {
            // Post-Order: erst Kinder, dann Operator
            visit(binOp->left.get());
            visit(binOp->right.get());
            
            switch (binOp->op) {
                case TokenType::PLUS:  chunk.emit(Opcode::ADD); break;
                case TokenType::MINUS: chunk.emit(Opcode::SUB); break;
                case TokenType::STAR:  chunk.emit(Opcode::MUL); break;
                case TokenType::SLASH: chunk.emit(Opcode::DIV); break;
            }
            
        } else if (auto* unOp = dynamic_cast<UnaryOp*>(node)) {
            visit(unOp->operand.get());
            
            if (unOp->op == TokenType::MINUS) {
                chunk.emit(Opcode::NEG);
            }
        }
    }
};

#endif
```

### Virtual Machine

```cpp
// vm.hpp

#ifndef VM_HPP
#define VM_HPP

#include "bytecode.hpp"
#include <stack>
#include <stdexcept>

class VM {
    std::stack<int32_t> stack;
    
public:
    int32_t execute(const Chunk& chunk) {
        size_t ip = 0;  // Instruction Pointer
        
        while (ip < chunk.code.size()) {
            const Instruction& instr = chunk.code[ip];
            
            switch (instr.op) {
                case Opcode::PUSH:
                    stack.push(instr.operand);
                    break;
                    
                case Opcode::POP:
                    if (stack.empty()) throw std::runtime_error("Stack underflow");
                    stack.pop();
                    break;
                    
                case Opcode::ADD: {
                    int32_t b = pop();
                    int32_t a = pop();
                    stack.push(a + b);
                    break;
                }
                
                case Opcode::SUB: {
                    int32_t b = pop();
                    int32_t a = pop();
                    stack.push(a - b);
                    break;
                }
                
                case Opcode::MUL: {
                    int32_t b = pop();
                    int32_t a = pop();
                    stack.push(a * b);
                    break;
                }
                
                case Opcode::DIV: {
                    int32_t b = pop();
                    int32_t a = pop();
                    if (b == 0) throw std::runtime_error("Division by zero");
                    stack.push(a / b);
                    break;
                }
                
                case Opcode::NEG: {
                    int32_t a = pop();
                    stack.push(-a);
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
        if (stack.empty()) throw std::runtime_error("Stack underflow");
        int32_t val = stack.top();
        stack.pop();
        return val;
    }
};

#endif
```

### Komplettes Programm

```cpp
// main.cpp

#include <iostream>
#include <string>
#include "lexer.hpp"
#include "parser.hpp"
#include "compiler.hpp"
#include "vm.hpp"

int main(int argc, char* argv[]) {
    std::string source;
    
    if (argc > 1) {
        // Argumente als Expression
        for (int i = 1; i < argc; i++) {
            if (i > 1) source += " ";
            source += argv[i];
        }
    } else {
        // REPL
        std::cout << "Expression Compiler (C++)" << std::endl;
        std::cout << "Type 'quit' to exit" << std::endl;
        
        while (true) {
            std::cout << "> ";
            if (!std::getline(std::cin, source) || source == "quit") {
                break;
            }
            if (source.empty()) continue;
            
            try {
                Lexer lexer(source);
                auto tokens = lexer.tokenize();
                
                Parser parser(tokens);
                auto ast = parser.parse();
                
                Compiler compiler;
                Chunk bytecode = compiler.compile(ast.get());
                
                std::cout << "Bytecode:" << std::endl;
                bytecode.disassemble();
                
                VM vm;
                int32_t result = vm.execute(bytecode);
                std::cout << "= " << result << std::endl;
                
            } catch (const std::exception& e) {
                std::cerr << "Error: " << e.what() << std::endl;
            }
        }
        return 0;
    }
    
    // Compile and run
    try {
        Lexer lexer(source);
        auto tokens = lexer.tokenize();
        
        Parser parser(tokens);
        auto ast = parser.parse();
        
        Compiler compiler;
        Chunk bytecode = compiler.compile(ast.get());
        
        VM vm;
        std::cout << vm.execute(bytecode) << std::endl;
        
    } catch (const std::exception& e) {
        std::cerr << "Error: " << e.what() << std::endl;
        return 1;
    }
    
    return 0;
}
```

---

## 🔧 Build Setup

### CMakeLists.txt

```cmake
cmake_minimum_required(VERSION 3.16)
project(calc CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(calc
    main.cpp
    lexer.cpp
    parser.cpp
    bytecode.cpp
    compiler.cpp
    vm.cpp
)

# Debug build
target_compile_options(calc PRIVATE -Wall -Wextra -g)
```

### Einfaches Makefile

```makefile
CXX = clang++
CXXFLAGS = -std=c++17 -Wall -Wextra -g

SRCS = main.cpp lexer.cpp parser.cpp bytecode.cpp
OBJS = $(SRCS:.cpp=.o)

calc: $(OBJS)
	$(CXX) $(CXXFLAGS) -o $@ $^

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<

clean:
	rm -f calc *.o

.PHONY: clean
```

---

## 📖 Weiterführende Ressourcen

### Bytecode VMs
- **Crafting Interpreters - Bytecode** - [craftinginterpreters.com/a-bytecode-virtual-machine.html](https://craftinginterpreters.com/a-bytecode-virtual-machine.html)
- **Lua VM** - Einfache Stack-basierte VM

### C++
- **Learn C++** - [learncpp.com](https://www.learncpp.com/)
- **A Tour of C++** - Bjarne Stroustrup
- **C++ Reference** - [cppreference.com](https://en.cppreference.com/)

### Optimierung
- **Constant Folding** - Berechne Konstanten zur Compile-Zeit
- **Dead Code Elimination** - Entferne unerreichbaren Code
- **Peephole Optimization** - Lokale Bytecode-Verbesserungen

---

## 🧠 Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Bytecode | Kompakte Instruktionsfolge |
| Stack VM | Operanden auf Stack, Operatoren manipulieren Stack |
| Compiler | AST → Bytecode (Post-Order Traversierung) |
| VM | Führt Bytecode aus (dispatch loop) |
| C++ | Klassen, Smart Pointers, Vererbung |

---

*Weiter zu den Übungen → `uebungen/uebung-06.md`*
