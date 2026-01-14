# Übung 7: GUI Basics & Rust

**Ziel**: Implementiere Framebuffer-Grafik und schreibe deine ersten Rust-Programme.

---

## 🛠️ Setup

```bash
# GUI Projekt
mkdir -p ~/os-dev/day7/gui
cd ~/os-dev/day7/gui

# Rust installieren (falls noch nicht)
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source ~/.cargo/env

# Rust Bare Metal Setup
rustup target add riscv64gc-unknown-none-elf
```

---

## Teil A: Framebuffer Grafik (C)

### Aufgabe 1: Grundlegende Primitives (⏱️ 60 min)

**Datei**: `fb.h`

```c
#ifndef FB_H
#define FB_H

#include <stdint.h>

#define FB_WIDTH  800
#define FB_HEIGHT 600

/* Farben */
#define COLOR_BLACK   0x00000000
#define COLOR_WHITE   0x00FFFFFF
#define COLOR_RED     0x00FF0000
#define COLOR_GREEN   0x0000FF00
#define COLOR_BLUE    0x000000FF

/* RGB Makro */
#define RGB(r,g,b) (((r)<<16)|((g)<<8)|(b))

/* Funktionen */
void fb_init(uint32_t *buffer);
void fb_clear(uint32_t color);
void fb_put_pixel(int x, int y, uint32_t color);
void fb_fill_rect(int x, int y, int w, int h, uint32_t color);
void fb_rect(int x, int y, int w, int h, uint32_t color);
void fb_line(int x0, int y0, int x1, int y1, uint32_t color);
void fb_circle(int cx, int cy, int r, uint32_t color);

#endif
```

**Datei**: `fb.c`

```c
#include "fb.h"

static uint32_t *framebuffer;

void fb_init(uint32_t *buffer) {
    framebuffer = buffer;
}

void fb_clear(uint32_t color) {
    for (int i = 0; i < FB_WIDTH * FB_HEIGHT; i++) {
        framebuffer[i] = color;
    }
}

void fb_put_pixel(int x, int y, uint32_t color) {
    if (x >= 0 && x < FB_WIDTH && y >= 0 && y < FB_HEIGHT) {
        framebuffer[y * FB_WIDTH + x] = color;
    }
}

void fb_fill_rect(int x, int y, int w, int h, uint32_t color) {
    // TODO: Implementiere gefülltes Rechteck
}

void fb_rect(int x, int y, int w, int h, uint32_t color) {
    // TODO: Implementiere Rechteck-Rahmen
}

/* Bresenham's Line Algorithm */
void fb_line(int x0, int y0, int x1, int y1, uint32_t color) {
    // TODO: Implementiere Linien-Algorithmus
    // Hint: 
    // dx = abs(x1 - x0), dy = abs(y1 - y0)
    // sx = x0 < x1 ? 1 : -1
    // sy = y0 < y1 ? 1 : -1
    // err = dx - dy
    // while (true) {
    //     put_pixel(x0, y0, color)
    //     if (x0 == x1 && y0 == y1) break
    //     e2 = 2 * err
    //     if (e2 > -dy) { err -= dy; x0 += sx; }
    //     if (e2 < dx) { err += dx; y0 += sy; }
    // }
}

/* Midpoint Circle Algorithm */
void fb_circle(int cx, int cy, int r, uint32_t color) {
    // TODO: Implementiere Kreis-Algorithmus
    // Hint: Zeichne 8 symmetrische Punkte gleichzeitig
}
```

### Aufgabe 2: Bitmap Font (⏱️ 45 min)

**Datei**: `font.c`

```c
#include <stdint.h>
#include "fb.h"

/* Minimaler 8x8 Font (nur ASCII 32-127) */
/* Du kannst die Daten von hier kopieren:
   https://github.com/dhepper/font8x8/blob/master/font8x8_basic.h */

static const uint8_t font_8x8[128][8] = {
    // Space (32)
    [' '] = {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
    
    // Digits 0-9
    ['0'] = {0x3C, 0x66, 0x6E, 0x76, 0x66, 0x66, 0x3C, 0x00},
    ['1'] = {0x18, 0x38, 0x18, 0x18, 0x18, 0x18, 0x7E, 0x00},
    ['2'] = {0x3C, 0x66, 0x06, 0x1C, 0x30, 0x60, 0x7E, 0x00},
    // TODO: Füge 3-9 hinzu
    
    // Letters A-Z
    ['A'] = {0x18, 0x3C, 0x66, 0x66, 0x7E, 0x66, 0x66, 0x00},
    ['B'] = {0x7C, 0x66, 0x66, 0x7C, 0x66, 0x66, 0x7C, 0x00},
    // TODO: Füge C-Z hinzu
    
    // Lowercase a-z
    ['a'] = {0x00, 0x00, 0x3C, 0x06, 0x3E, 0x66, 0x3E, 0x00},
    // TODO: Füge b-z hinzu
};

void fb_putchar(int x, int y, char c, uint32_t fg, uint32_t bg) {
    // TODO: Zeichne Zeichen mit Font-Bitmap
    // Für jede Zeile (0-7):
    //   Für jede Spalte (0-7):
    //     if (font[c][row] & (0x80 >> col)) { pixel = fg } else { pixel = bg }
}

void fb_puts(int x, int y, const char *s, uint32_t fg, uint32_t bg) {
    // TODO: Zeichne String
}
```

### Aufgabe 3: GUI Widgets (⏱️ 60 min)

**Datei**: `widgets.c`

```c
#include "fb.h"

/* Button Widget */
struct Button {
    int x, y, w, h;
    const char *label;
    int hover;
    int pressed;
};

void button_draw(struct Button *btn) {
    uint32_t bg = btn->pressed ? 0x00505050 :
                  btn->hover ? 0x00A0A0A0 : 0x00808080;
    
    // TODO: Zeichne Button mit 3D-Effekt
    // 1. Hintergrund füllen
    // 2. Helle Linien oben/links
    // 3. Dunkle Linien unten/rechts
    // 4. Label zentriert
}

int button_contains(struct Button *btn, int mx, int my) {
    // TODO: Prüfe ob Punkt (mx, my) im Button liegt
    return 0;
}

/* Progress Bar */
struct ProgressBar {
    int x, y, w, h;
    int value;  /* 0-100 */
    uint32_t fg, bg;
};

void progress_draw(struct ProgressBar *pb) {
    // TODO: Zeichne Progress-Bar
    // 1. Hintergrund
    // 2. Gefüllter Teil proportional zu value
}

/* Text Input (einfach) */
struct TextInput {
    int x, y, w, h;
    char text[256];
    int cursor_pos;
    int focused;
};

void textinput_draw(struct TextInput *ti) {
    // TODO: Zeichne Textfeld mit blinkenden Cursor
}
```

---

## Teil B: Rust Grundlagen

### Aufgabe 4: Rust Basics (⏱️ 45 min)

**Datei**: `basics.rs`

```rust
// Führe aus mit: rustc basics.rs && ./basics

fn main() {
    // Aufgabe 4a: Variablen und Mutabilität
    let x = 5;
    // TODO: Mache x mutable und addiere 1
    
    println!("x = {}", x);
    
    // Aufgabe 4b: Ownership
    let s1 = String::from("hello");
    // TODO: Erstelle s2 als Kopie (nicht move!) von s1
    // Hint: s1.clone()
    
    println!("s1 = {}, s2 = ???", s1);
    
    // Aufgabe 4c: Borrowing
    let s = String::from("hello world");
    // TODO: Rufe print_length mit Referenz auf s auf
    
    println!("s is still valid: {}", s);
    
    // Aufgabe 4d: Slices
    let sentence = "The quick brown fox";
    // TODO: Extrahiere "quick" als Slice
    let word: &str = ???;
    println!("Word: {}", word);
    
    // Aufgabe 4e: Structs
    // TODO: Erstelle Point struct und berechne Distanz
}

fn print_length(s: &String) {
    println!("Length: {}", s.len());
}

// TODO: Implementiere Point struct
struct Point {
    // x: ???,
    // y: ???,
}

impl Point {
    fn new(x: i32, y: i32) -> Self {
        // TODO
    }
    
    fn distance(&self, other: &Point) -> f64 {
        // TODO
    }
}
```

### Aufgabe 5: Enums und Pattern Matching (⏱️ 30 min)

**Datei**: `enums.rs`

```rust
// Aufgabe 5a: Definiere Color enum
enum Color {
    // TODO: RGB(u8, u8, u8)
    // TODO: Named(&'static str)
    // TODO: Hex(u32)
}

impl Color {
    fn to_hex(&self) -> u32 {
        match self {
            // TODO: Implementiere für alle Varianten
        }
    }
    
    fn name(&self) -> String {
        match self {
            // TODO
        }
    }
}

// Aufgabe 5b: Token für Mini-Compiler
enum Token {
    Integer(i64),
    Plus,
    Minus,
    Star,
    Slash,
    LParen,
    RParen,
    Eof,
}

fn token_to_string(token: &Token) -> String {
    // TODO: Pattern matching
}

fn main() {
    let colors = vec![
        Color::RGB(255, 0, 0),
        Color::Named("blue"),
        Color::Hex(0x00FF00),
    ];
    
    for c in &colors {
        println!("{} = 0x{:06X}", c.name(), c.to_hex());
    }
}
```

### Aufgabe 6: Ownership & Lifetimes (⏱️ 45 min)

**Datei**: `ownership.rs`

```rust
// Aufgabe 6a: Finde den Fehler
fn broken_1() {
    let s = String::from("hello");
    takes_ownership(s);
    // println!("{}", s);  // <- Warum Fehler?
}

fn takes_ownership(s: String) {
    println!("{}", s);
}

// Aufgabe 6b: Repariere mit Borrowing
fn fixed_1() {
    let s = String::from("hello");
    // TODO: Ändere zu borrows_string(&s)
}

fn borrows_string(s: ???) {
    println!("{}", s);
}

// Aufgabe 6c: Mutable Borrowing
fn append_exclamation(s: &mut String) {
    // TODO: Füge "!" an s an
}

fn test_mut_borrow() {
    let mut greeting = String::from("Hello");
    // TODO: Rufe append_exclamation auf
    println!("{}", greeting);  // "Hello!"
}

// Aufgabe 6d: Mehrere Referenzen
fn multiple_refs() {
    let s = String::from("hello");
    
    let r1 = &s;
    let r2 = &s;
    println!("{} and {}", r1, r2);
    
    // TODO: Warum funktioniert folgendes NICHT?
    // let r3 = &mut s;
    // println!("{}", r3);
}

// Aufgabe 6e: Einfache Lifetimes
// TODO: Füge Lifetime-Annotation hinzu
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() { x } else { y }
}

fn main() {
    test_mut_borrow();
    
    let s1 = String::from("long string");
    let s2 = String::from("short");
    let result = longest(&s1, &s2);
    println!("Longest: {}", result);
}
```

### Aufgabe 7: Collections & Iteratoren (⏱️ 30 min)

**Datei**: `collections.rs`

```rust
fn main() {
    // Aufgabe 7a: Vectors
    let mut numbers: Vec<i32> = vec![1, 2, 3, 4, 5];
    
    // TODO: Füge 6, 7, 8 hinzu
    
    // TODO: Entferne das erste Element
    
    // TODO: Verdopple alle Werte mit map()
    let doubled: Vec<i32> = ???;
    
    println!("Doubled: {:?}", doubled);
    
    // Aufgabe 7b: HashMaps
    use std::collections::HashMap;
    
    let mut scores: HashMap<String, i32> = HashMap::new();
    
    // TODO: Füge "Alice" -> 100, "Bob" -> 85 hinzu
    
    // TODO: Gib Bob's Score aus (mit get())
    
    // TODO: Aktualisiere Alice's Score auf 105
    
    // Aufgabe 7c: Iterator Chains
    let data = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    
    // TODO: Berechne Summe aller geraden Zahlen
    let sum: i32 = data.iter()
        // .filter(???)
        // .sum()
    ;
    
    println!("Sum of evens: {}", sum);  // 30
    
    // TODO: Finde das erste Element > 5
    let first_big: Option<&i32> = ???;
    
    if let Some(n) = first_big {
        println!("First > 5: {}", n);
    }
}
```

### Aufgabe 8: Mini-Rechner in Rust (⏱️ 60 min)

Portiere den Expression-Rechner nach Rust.

**Datei**: `calc.rs`

```rust
// Token
#[derive(Debug, Clone)]
enum Token {
    Integer(i64),
    Plus, Minus, Star, Slash,
    LParen, RParen,
    Eof,
}

// Lexer
struct Lexer {
    input: Vec<char>,
    pos: usize,
}

impl Lexer {
    fn new(input: &str) -> Self {
        Lexer {
            input: input.chars().collect(),
            pos: 0,
        }
    }
    
    fn tokenize(&mut self) -> Vec<Token> {
        let mut tokens = Vec::new();
        // TODO: Implementiere
        tokens.push(Token::Eof);
        tokens
    }
    
    // TODO: Implementiere Helper-Methoden
}

// AST
enum Expr {
    Int(i64),
    BinOp(Box<Expr>, Token, Box<Expr>),
    UnaryOp(Token, Box<Expr>),
}

impl Expr {
    fn eval(&self) -> i64 {
        match self {
            Expr::Int(n) => *n,
            Expr::BinOp(left, op, right) => {
                let l = left.eval();
                let r = right.eval();
                match op {
                    Token::Plus => l + r,
                    Token::Minus => l - r,
                    Token::Star => l * r,
                    Token::Slash => l / r,
                    _ => panic!("Invalid operator"),
                }
            }
            Expr::UnaryOp(op, expr) => {
                match op {
                    Token::Minus => -expr.eval(),
                    _ => panic!("Invalid unary operator"),
                }
            }
        }
    }
}

// Parser
struct Parser {
    tokens: Vec<Token>,
    pos: usize,
}

impl Parser {
    fn new(tokens: Vec<Token>) -> Self {
        Parser { tokens, pos: 0 }
    }
    
    fn parse(&mut self) -> Expr {
        // TODO: Implementiere recursive descent
        Expr::Int(0)
    }
    
    // TODO: parse_expression, parse_term, parse_factor
}

fn main() {
    let input = "4 * 7 / (4 - 3)";
    
    let mut lexer = Lexer::new(input);
    let tokens = lexer.tokenize();
    
    println!("Tokens: {:?}", tokens);
    
    let mut parser = Parser::new(tokens);
    let ast = parser.parse();
    
    println!("Result: {}", ast.eval());  // 28
}
```

---

## 🎯 Bonusaufgaben

### Bonus A: Gradient

Zeichne einen Farbverlauf (Gradient) über den gesamten Bildschirm.

### Bonus B: Bouncing Ball

Animiere einen Ball der an den Bildschirmrändern abprallt.

### Bonus C: Rust Bare Metal

Portiere den UART-Output nach Rust mit `#![no_std]`.

---

## ✅ Checkliste

- [ ] Framebuffer-Primitives funktionieren
- [ ] Font-Rendering funktioniert
- [ ] Mindestens ein Widget implementiert
- [ ] Rust Basics verstanden
- [ ] Ownership/Borrowing verstanden
- [ ] Mini-Rechner in Rust funktioniert

---

## 🎉 Herzlichen Glückwunsch!

Du hast den 7-Tage-Crashkurs abgeschlossen!

**Was du gelernt hast:**
- OS-Entwicklung: Bare Metal → Kernel → Terminal
- Compiler-Bau: Lexer → Parser → Bytecode → VM
- Sprachen: C, Ruby, C++, Rust Grundlagen
- Grafik: Framebuffer, Fonts, GUI-Widgets

**Nächste Schritte:**
- Folge `os.phil-opp.com` für mehr Rust OS
- Lies "Crafting Interpreters" für vollständigen Compiler
- Implementiere Multitasking in deinem Kernel
- Probiere echte Hardware (Raspberry Pi, RISC-V Boards)

---

*Viel Erfolg auf deiner weiteren Reise!*
