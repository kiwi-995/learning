# Tag 7: GUI Basics & Rust Introduction

**Ziel**: Framebuffer-Grafik für dein OS und die ersten Schritte mit Rust.

---

## 📚 Theorie

### 7.1 Framebuffer Grundlagen

Ein Framebuffer ist ein Speicherbereich, der direkt die Pixel des Bildschirms enthält:

```
┌─────────────────────────────────────────────────────────────────┐
│  Framebuffer Layout (Linear)                                   │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  Pixel (x, y) = framebuffer[(y * width + x) * bytes_per_pixel] │
│                                                                 │
│  Beispiel: 800x600 mit 32-bit Farben (ARGB)                     │
│                                                                 │
│  fb[0..3]   = Pixel (0,0)    ← Oben links                       │
│  fb[4..7]   = Pixel (1,0)                                       │
│  ...                                                            │
│  fb[3196..3199] = Pixel (799,0) ← Ende erste Zeile              │
│  fb[3200..3203] = Pixel (0,1)   ← Anfang zweite Zeile           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Farb-Format (32-bit ARGB):
┌────────┬────────┬────────┬────────┐
│ Alpha  │  Red   │ Green  │  Blue  │
│ 8 bit  │ 8 bit  │ 8 bit  │ 8 bit  │
└────────┴────────┴────────┴────────┘
  Bits:  31-24    23-16    15-8     7-0
```

### 7.2 QEMU VirtIO GPU

QEMU bietet verschiedene Display-Optionen:

```
┌─────────────────────────────────────────────────────────────────┐
│  Option 1: Ramfb (Simple Framebuffer)                           │
│  - Einfachster Ansatz                                           │
│  - Feste Auflösung, direkter Speicherzugriff                    │
│  + Perfekt zum Lernen                                           │
├─────────────────────────────────────────────────────────────────┤
│  Option 2: VirtIO GPU                                           │
│  - Modernere Lösung                                             │
│  - Unterstützt 2D und 3D Beschleunigung                         │
│  - Komplexere Initialisierung                                   │
├─────────────────────────────────────────────────────────────────┤
│  Option 3: Bochs VBE                                            │
│  - Legacy BIOS Grafik                                           │
│  - Gut dokumentiert                                             │
│  - Für x86 typischer                                            │
└─────────────────────────────────────────────────────────────────┘
```

### 7.3 Einfache Grafikoperationen

```c
// Grundlegende Primitives

void put_pixel(int x, int y, uint32_t color) {
    framebuffer[y * width + x] = color;
}

void draw_line(int x0, int y0, int x1, int y1, uint32_t color) {
    // Bresenham's Line Algorithm
}

void draw_rect(int x, int y, int w, int h, uint32_t color) {
    for (int py = y; py < y + h; py++) {
        for (int px = x; px < x + w; px++) {
            put_pixel(px, py, color);
        }
    }
}

void fill_screen(uint32_t color) {
    for (int i = 0; i < width * height; i++) {
        framebuffer[i] = color;
    }
}
```

### 7.4 Font Rendering

Um Text anzuzeigen brauchen wir einen Bitmap-Font:

```
┌──────────────────────────────────────────────────────────────┐
│  Bitmap Font (8x16 pro Zeichen)                              │
├──────────────────────────────────────────────────────────────┤
│                                                              │
│  Zeichen 'A' (8x16 Pixel):                                   │
│                                                              │
│    00011000  = 0x18                                          │
│    00111100  = 0x3C                                          │
│    01100110  = 0x66                                          │
│    01100110  = 0x66                                          │
│    01111110  = 0x7E                                          │
│    01100110  = 0x66                                          │
│    01100110  = 0x66                                          │
│    00000000                                                  │
│    ...                                                       │
│                                                              │
│  Font als Array:                                             │
│    font[65] = {0x18, 0x3C, 0x66, ...}  // 'A'                │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

### 7.5 Rust Einführung

Rust bietet Memory Safety ohne Garbage Collector:

```rust
// Variablen (immutable by default!)
let x = 5;          // immutable
let mut y = 10;     // mutable
y += 1;             // OK

// Ownership
let s1 = String::from("hello");
let s2 = s1;        // s1 ist jetzt UNGÜLTIG (moved)
// println!("{}", s1);  // ERROR!

// Borrowing
let s = String::from("hello");
let len = calculate_length(&s);  // Borrows s
println!("{} has length {}", s, len);  // s ist noch gültig

fn calculate_length(s: &String) -> usize {
    s.len()
}

// Mutable Borrowing
fn append_world(s: &mut String) {
    s.push_str(" world");
}

// Structs
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn new(x: i32, y: i32) -> Self {
        Point { x, y }
    }
    
    fn distance(&self, other: &Point) -> f64 {
        let dx = (self.x - other.x) as f64;
        let dy = (self.y - other.y) as f64;
        (dx*dx + dy*dy).sqrt()
    }
}

// Enums
enum Color {
    Rgb(u8, u8, u8),
    Named(&'static str),
}

// Pattern Matching
match color {
    Color::Rgb(r, g, b) => println!("RGB: {},{},{}", r, g, b),
    Color::Named(name) => println!("Color: {}", name),
}

// Option und Result
fn divide(a: i32, b: i32) -> Option<i32> {
    if b == 0 { None } else { Some(a / b) }
}

if let Some(result) = divide(10, 2) {
    println!("Result: {}", result);
}
```

### 7.6 Bare Metal Rust

Rust für OS-Entwicklung:

```rust
#![no_std]  // Keine Standard Library
#![no_main] // Kein normaler main()

use core::panic::PanicInfo;

// Eigener Entry Point
#[no_mangle]
pub extern "C" fn kernel_main() -> ! {
    // Kernel code
    loop {}
}

// Panic Handler (required)
#[panic_handler]
fn panic(_info: &PanicInfo) -> ! {
    loop {}
}

// Unsafe für Hardware-Zugriff
fn uart_write(c: u8) {
    unsafe {
        let uart = 0x10000000 as *mut u8;
        core::ptr::write_volatile(uart, c);
    }
}
```

---

## 💻 Code-Beispiele

### Einfacher Framebuffer (C)

```c
/* framebuffer.c */

#include <stdint.h>

/* Framebuffer Konfiguration (für QEMU ramfb) */
#define FB_WIDTH  800
#define FB_HEIGHT 600
#define FB_BPP    32   /* Bits per pixel */

/* Framebuffer Adresse (variiert je nach Setup) */
static uint32_t *framebuffer = (uint32_t *)0x40000000;

/* Farben */
#define COLOR_BLACK   0x00000000
#define COLOR_WHITE   0x00FFFFFF
#define COLOR_RED     0x00FF0000
#define COLOR_GREEN   0x0000FF00
#define COLOR_BLUE    0x000000FF

/* RGB zu 32-bit Farbe */
static inline uint32_t rgb(uint8_t r, uint8_t g, uint8_t b) {
    return (r << 16) | (g << 8) | b;
}

/* Pixel setzen */
void fb_put_pixel(int x, int y, uint32_t color) {
    if (x >= 0 && x < FB_WIDTH && y >= 0 && y < FB_HEIGHT) {
        framebuffer[y * FB_WIDTH + x] = color;
    }
}

/* Rechteck zeichnen (gefüllt) */
void fb_fill_rect(int x, int y, int w, int h, uint32_t color) {
    for (int py = y; py < y + h && py < FB_HEIGHT; py++) {
        for (int px = x; px < x + w && px < FB_WIDTH; px++) {
            if (px >= 0 && py >= 0) {
                framebuffer[py * FB_WIDTH + px] = color;
            }
        }
    }
}

/* Bildschirm löschen */
void fb_clear(uint32_t color) {
    for (int i = 0; i < FB_WIDTH * FB_HEIGHT; i++) {
        framebuffer[i] = color;
    }
}

/* Horizontale Linie */
void fb_hline(int x, int y, int w, uint32_t color) {
    for (int i = 0; i < w && x + i < FB_WIDTH; i++) {
        fb_put_pixel(x + i, y, color);
    }
}

/* Vertikale Linie */
void fb_vline(int x, int y, int h, uint32_t color) {
    for (int i = 0; i < h && y + i < FB_HEIGHT; i++) {
        fb_put_pixel(x, y + i, color);
    }
}

/* Rechteck-Rahmen */
void fb_rect(int x, int y, int w, int h, uint32_t color) {
    fb_hline(x, y, w, color);           /* Oben */
    fb_hline(x, y + h - 1, w, color);   /* Unten */
    fb_vline(x, y, h, color);           /* Links */
    fb_vline(x + w - 1, y, h, color);   /* Rechts */
}
```

### Bitmap Font Renderer

```c
/* font.c */

#include <stdint.h>

/* Einfacher 8x8 Bitmap Font (nur Großbuchstaben + Ziffern) */
static const uint8_t font_8x8[128][8] = {
    /* 0-31: Control characters (leer) */
    [0 ... 31] = {0, 0, 0, 0, 0, 0, 0, 0},
    
    /* 32: Space */
    [32] = {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
    
    /* 48-57: Ziffern 0-9 */
    [48] = {0x3C, 0x66, 0x6E, 0x76, 0x66, 0x66, 0x3C, 0x00}, /* 0 */
    [49] = {0x18, 0x38, 0x18, 0x18, 0x18, 0x18, 0x7E, 0x00}, /* 1 */
    /* ... weitere Ziffern ... */
    
    /* 65-90: Großbuchstaben A-Z */
    [65] = {0x18, 0x3C, 0x66, 0x66, 0x7E, 0x66, 0x66, 0x00}, /* A */
    [66] = {0x7C, 0x66, 0x66, 0x7C, 0x66, 0x66, 0x7C, 0x00}, /* B */
    [67] = {0x3C, 0x66, 0x60, 0x60, 0x60, 0x66, 0x3C, 0x00}, /* C */
    /* ... weitere Buchstaben ... */
};

extern void fb_put_pixel(int x, int y, uint32_t color);

/* Zeichen zeichnen */
void fb_putchar(int x, int y, char c, uint32_t fg, uint32_t bg) {
    const uint8_t *glyph = font_8x8[(unsigned char)c];
    
    for (int row = 0; row < 8; row++) {
        uint8_t bits = glyph[row];
        for (int col = 0; col < 8; col++) {
            uint32_t color = (bits & (0x80 >> col)) ? fg : bg;
            fb_put_pixel(x + col, y + row, color);
        }
    }
}

/* String zeichnen */
void fb_puts(int x, int y, const char *s, uint32_t fg, uint32_t bg) {
    while (*s) {
        fb_putchar(x, y, *s, fg, bg);
        x += 8;
        if (x >= 800 - 8) {  /* Wrap around */
            x = 0;
            y += 8;
        }
        s++;
    }
}

/* Integer als String zeichnen */
void fb_print_int(int x, int y, int value, uint32_t fg, uint32_t bg) {
    char buffer[20];
    int i = 0;
    int negative = 0;
    
    if (value < 0) {
        negative = 1;
        value = -value;
    }
    
    if (value == 0) {
        buffer[i++] = '0';
    } else {
        while (value > 0) {
            buffer[i++] = '0' + (value % 10);
            value /= 10;
        }
    }
    
    if (negative) buffer[i++] = '-';
    
    /* Reverse */
    for (int j = 0; j < i / 2; j++) {
        char tmp = buffer[j];
        buffer[j] = buffer[i - 1 - j];
        buffer[i - 1 - j] = tmp;
    }
    buffer[i] = '\0';
    
    fb_puts(x, y, buffer, fg, bg);
}
```

### Einfache GUI-Elemente

```c
/* gui.c */

#include <stdint.h>

/* Button */
struct Button {
    int x, y, w, h;
    const char *label;
    uint32_t bg_color;
    uint32_t fg_color;
    int pressed;
};

void draw_button(struct Button *btn) {
    uint32_t bg = btn->pressed ? 0x00404040 : btn->bg_color;
    
    /* Hintergrund */
    fb_fill_rect(btn->x, btn->y, btn->w, btn->h, bg);
    
    /* Rahmen */
    uint32_t light = 0x00FFFFFF;
    uint32_t dark = 0x00404040;
    
    if (!btn->pressed) {
        /* 3D-Effekt: Links/Oben hell, Rechts/Unten dunkel */
        fb_hline(btn->x, btn->y, btn->w, light);
        fb_vline(btn->x, btn->y, btn->h, light);
        fb_hline(btn->x, btn->y + btn->h - 1, btn->w, dark);
        fb_vline(btn->x + btn->w - 1, btn->y, btn->h, dark);
    } else {
        /* Umgekehrt wenn gedrückt */
        fb_hline(btn->x, btn->y, btn->w, dark);
        fb_vline(btn->x, btn->y, btn->h, dark);
        fb_hline(btn->x, btn->y + btn->h - 1, btn->w, light);
        fb_vline(btn->x + btn->w - 1, btn->y, btn->h, light);
    }
    
    /* Label zentriert */
    int text_len = 0;
    const char *s = btn->label;
    while (*s++) text_len++;
    
    int text_x = btn->x + (btn->w - text_len * 8) / 2;
    int text_y = btn->y + (btn->h - 8) / 2;
    
    fb_puts(text_x, text_y, btn->label, btn->fg_color, bg);
}

/* Einfaches Window */
struct Window {
    int x, y, w, h;
    const char *title;
    uint32_t title_bg;
};

void draw_window(struct Window *win) {
    /* Hintergrund */
    fb_fill_rect(win->x, win->y, win->w, win->h, 0x00C0C0C0);
    
    /* Titelleiste */
    fb_fill_rect(win->x, win->y, win->w, 20, win->title_bg);
    fb_puts(win->x + 4, win->y + 6, win->title, 0x00FFFFFF, win->title_bg);
    
    /* Rahmen */
    fb_rect(win->x, win->y, win->w, win->h, 0x00000000);
}
```

### Rust Bare Metal Beispiel

```rust
// main.rs - Bare Metal Rust für RISC-V

#![no_std]
#![no_main]

use core::panic::PanicInfo;

// UART Adresse
const UART_BASE: usize = 0x10000000;

fn uart_putc(c: u8) {
    unsafe {
        let uart = UART_BASE as *mut u8;
        core::ptr::write_volatile(uart, c);
    }
}

fn uart_puts(s: &str) {
    for c in s.bytes() {
        if c == b'\n' {
            uart_putc(b'\r');
        }
        uart_putc(c);
    }
}

#[no_mangle]
pub extern "C" fn kernel_main() -> ! {
    uart_puts("Hello from Rust!\n");
    uart_puts("This is bare-metal Rust on RISC-V\n");
    
    loop {
        unsafe { core::arch::asm!("wfi"); }
    }
}

#[panic_handler]
fn panic(info: &PanicInfo) -> ! {
    uart_puts("PANIC: ");
    if let Some(location) = info.location() {
        uart_puts(location.file());
    }
    uart_puts("\n");
    
    loop {
        unsafe { core::arch::asm!("wfi"); }
    }
}
```

---

## 📖 Weiterführende Ressourcen

### Grafik
- **OSDev Wiki - Drawing In Protected Mode** - [wiki.osdev.org](https://wiki.osdev.org/Drawing_In_Protected_Mode)
- **Bresenham's Line Algorithm** - [Wikipedia](https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm)

### GUI Design
- **GUI From Scratch** - [Creating a GUI from scratch](https://handmade.network/)
- **Casey Muratori's Immediate Mode GUI** - [youtube.com](https://www.youtube.com/watch?v=Z1qyvQsjK5Y)

### Rust
- **The Rust Book** - [doc.rust-lang.org/book](https://doc.rust-lang.org/book/)
- **Rustlings** - Interactive exercises
- **Rust by Example** - [doc.rust-lang.org/rust-by-example](https://doc.rust-lang.org/rust-by-example/)
- **Writing an OS in Rust** - [os.phil-opp.com](https://os.phil-opp.com/) (x86, aber exzellent)

---

## 🧠 Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Framebuffer | Speicherbereich für Pixel-Daten |
| Pixel Format | ARGB 32-bit: (Alpha, Red, Green, Blue) |
| Bitmap Font | Array von Bit-Patterns pro Zeichen |
| Rust Ownership | Jeder Wert hat genau einen Owner |
| Borrowing | Referenzen (`&T`, `&mut T`) ohne Ownership-Transfer |
| `no_std` | Rust ohne Standard Library |

---

## 🎉 Kurs-Abschluss

Du hast in 7 Tagen gelernt:

1. **Tag 1**: C Grundlagen, Pointer, Memory Model
2. **Tag 2**: Bare Metal RISC-V, QEMU, Boot-Prozess
3. **Tag 3**: Memory Management, Interrupts, Traps
4. **Tag 4**: UART I/O, Terminal, Shell
5. **Tag 5**: Compiler-Grundlagen, Lexer/Parser in Ruby
6. **Tag 6**: Bytecode-Compiler, Stack-VM in C++
7. **Tag 7**: Framebuffer-Grafik, Rust Einführung

**Du kannst jetzt:**
- Bare Metal Code für RISC-V schreiben
- Einen minimalen Kernel mit Terminal bauen
- Einen Expression-Compiler entwickeln
- Grundlegende GUI-Elemente rendern
- Rust für Systems Programming nutzen

**Nächste Schritte:**
- [ ] Implementiere Paging und Virtual Memory
- [ ] Füge Multitasking hinzu
- [ ] Portiere nach ARM (Raspberry Pi)
- [ ] Erweitere den Compiler um Variablen und Funktionen
- [ ] Baue einen Window Manager

---

*Weiter zu den Übungen → `uebungen/uebung-07.md`*
