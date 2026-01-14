# Tag 1: Foundations – C Basics & The Hardware Bridge

**Ziel**: Verstehe wie C direkt mit Hardware kommuniziert und baue die Brücke von Nand2Tetris zur echten Systemprogrammierung.

---

## 📚 Theorie

### 1.1 Von Nand2Tetris zu echter Hardware

In Nand2Tetris hast du gelernt:
- **Hardware**: NAND-Gatter → ALU → CPU → Computer
- **Software**: Assembler → VM → Compiler → OS

Der entscheidende Unterschied zur echten Welt:

| Nand2Tetris | Echte Hardware |
|-------------|----------------|
| Simulierte 16-bit CPU | Echte 32/64-bit RISC-V/ARM/x64 |
| Memory-mapped I/O (Screen, Keyboard) | UART, Framebuffer, Interrupts |
| Einfaches OS-Modell | Paging, Syscalls, Privilegierte Modi |
| Jack-Sprache | C, C++, Rust |

### 1.2 Warum C die Sprache der Systemprogrammierung ist

C wurde buchstäblich erfunden, um UNIX zu schreiben. Es bietet:

```
┌─────────────────────────────────────────────────────┐
│  High-Level Abstractions (Functions, Structs)       │
├─────────────────────────────────────────────────────┤
│  Direct Memory Access (Pointers)                    │
├─────────────────────────────────────────────────────┤
│  Minimal Runtime (No GC, No VM)                     │
├─────────────────────────────────────────────────────┤
│  Hardware Registers (volatile, inline asm)          │
└─────────────────────────────────────────────────────┘
```

### 1.3 Das C Memory Model

Dies ist das wichtigste Konzept für OS-Entwicklung:

```
┌─────────────────────────┐ Hohe Adressen (z.B. 0xFFFFFFFF)
│         Stack           │ ← Lokale Variablen, wächst nach unten
│           ↓             │
├─────────────────────────┤
│                         │
│    (freier Speicher)    │
│                         │
├─────────────────────────┤
│           ↑             │
│          Heap           │ ← malloc(), wächst nach oben
├─────────────────────────┤
│         .bss            │ ← Uninitialisierte globale Variablen
├─────────────────────────┤
│         .data           │ ← Initialisierte globale Variablen
├─────────────────────────┤
│         .text           │ ← Programmcode (Maschinencode)
└─────────────────────────┘ Niedrige Adressen (z.B. 0x00000000)
```

### 1.4 C Grundlagen – Crashkurs

#### Variablen und Datentypen

```c
// Grundtypen
int a = 42;           // Integer (32 bit auf den meisten Systemen)
char c = 'A';         // Character (8 bit)
float f = 3.14f;      // Floating point (32 bit)
double d = 3.14159;   // Double precision (64 bit)

// Feste Größen (wichtig für Hardware!)
#include <stdint.h>
uint8_t  byte = 255;        // Unsigned 8-bit (0-255)
int32_t  signed_int = -100; // Signed 32-bit
uint64_t big = 0xDEADBEEF;  // Unsigned 64-bit
```

#### Pointer – Das Herzstück von C

```c
int value = 42;
int *ptr = &value;  // ptr enthält die ADRESSE von value

printf("Wert: %d\n", value);    // 42
printf("Adresse: %p\n", ptr);   // z.B. 0x7fff5fbff8ac
printf("Dereferenziert: %d\n", *ptr);  // 42 (Wert AN der Adresse)

// Pointer-Arithmetik
int array[5] = {10, 20, 30, 40, 50};
int *p = array;  // p zeigt auf array[0]
p++;             // p zeigt jetzt auf array[1]
printf("%d\n", *p);  // 20
```

#### Structs – Zusammengesetzte Datentypen

```c
struct Point {
    int x;
    int y;
};

struct Point p1 = {10, 20};
printf("x=%d, y=%d\n", p1.x, p1.y);

// Mit Pointer
struct Point *ptr = &p1;
printf("x=%d\n", ptr->x);  // Arrow-Operator für Pointer auf Structs
```

#### Funktionen

```c
// Deklaration (Prototyp)
int add(int a, int b);

// Definition
int add(int a, int b) {
    return a + b;
}

// Pointer als Parameter (Pass-by-Reference)
void increment(int *value) {
    (*value)++;
}

int main() {
    int x = 5;
    increment(&x);  // x ist jetzt 6
    return 0;
}
```

### 1.5 Hardware-nahes C

#### Memory-Mapped I/O

Genau wie in Nand2Tetris, aber mit echten Adressen:

```c
// In Nand2Tetris: SCREEN = 16384
// In echtem RISC-V QEMU: UART = 0x10000000

#define UART_BASE 0x10000000

// volatile = Compiler darf nicht wegoptimieren!
volatile char *uart = (volatile char *)UART_BASE;

void putchar(char c) {
    *uart = c;  // Schreibt direkt ins Hardware-Register
}
```

#### Inline Assembly

Für Operationen die C nicht direkt unterstützt:

```c
// RISC-V: Lese den Stack Pointer
uint64_t get_sp() {
    uint64_t sp;
    asm volatile ("mv %0, sp" : "=r"(sp));
    return sp;
}

// RISC-V: Aktiviere Interrupts
void enable_interrupts() {
    asm volatile ("csrsi mstatus, 0x8");
}
```

### 1.6 Der Build-Prozess

```
┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐
│  main.c  │    │  main.i  │    │  main.s  │    │  main.o  │
│ (Source) │───▶│(Preprocessed)│───▶│(Assembly)│───▶│ (Object) │
└──────────┘    └──────────┘    └──────────┘    └──────────┘
     │              cpp              cc1             as
     │
     │         ┌──────────┐    ┌──────────┐
     └────────▶│  main.o  │───▶│   a.out  │
               │  lib.o   │    │ (Binary) │
               └──────────┘    └──────────┘
                     ld (Linker)
```

Schlüsselkonzept: **Der Linker** kombiniert Object-Dateien und löst Symbole auf.

---

## 💻 Code-Beispiele

### Beispiel 1: Pointer visualisieren

```c
#include <stdio.h>
#include <stdint.h>

int main() {
    int numbers[4] = {10, 20, 30, 40};
    
    printf("Array-Startadresse: %p\n", (void*)numbers);
    
    for (int i = 0; i < 4; i++) {
        printf("numbers[%d] @ %p = %d\n", 
               i, 
               (void*)&numbers[i], 
               numbers[i]);
    }
    
    // Pointer-Arithmetik
    int *ptr = numbers;
    printf("\n*ptr     = %d (@ %p)\n", *ptr, (void*)ptr);
    printf("*(ptr+2) = %d (@ %p)\n", *(ptr+2), (void*)(ptr+2));
    
    return 0;
}
```

### Beispiel 2: Struct für Hardware-Register

```c
#include <stdint.h>

// Typisches UART-Register-Layout
struct UART_Registers {
    volatile uint32_t THR;   // Transmit Holding Register
    volatile uint32_t RBR;   // Receive Buffer Register  
    volatile uint32_t IER;   // Interrupt Enable Register
    volatile uint32_t IIR;   // Interrupt Identification Register
    volatile uint32_t LCR;   // Line Control Register
    volatile uint32_t MCR;   // Modem Control Register
    volatile uint32_t LSR;   // Line Status Register
    volatile uint32_t MSR;   // Modem Status Register
};

#define UART0 ((struct UART_Registers *)0x10000000)

void uart_putchar(char c) {
    // Warte bis Transmitter bereit (Bit 5 im LSR)
    while ((UART0->LSR & 0x20) == 0) {}
    UART0->THR = c;
}

void uart_print(const char *s) {
    while (*s) {
        uart_putchar(*s++);
    }
}
```

### Beispiel 3: Einfacher Ringbuffer

```c
#include <stdint.h>
#include <stdbool.h>

#define BUFFER_SIZE 64

struct RingBuffer {
    uint8_t data[BUFFER_SIZE];
    uint32_t head;  // Schreibposition
    uint32_t tail;  // Leseposition
    uint32_t count; // Anzahl Elemente
};

void ring_init(struct RingBuffer *rb) {
    rb->head = 0;
    rb->tail = 0;
    rb->count = 0;
}

bool ring_push(struct RingBuffer *rb, uint8_t byte) {
    if (rb->count >= BUFFER_SIZE) {
        return false;  // Buffer voll
    }
    rb->data[rb->head] = byte;
    rb->head = (rb->head + 1) % BUFFER_SIZE;
    rb->count++;
    return true;
}

bool ring_pop(struct RingBuffer *rb, uint8_t *byte) {
    if (rb->count == 0) {
        return false;  // Buffer leer
    }
    *byte = rb->data[rb->tail];
    rb->tail = (rb->tail + 1) % BUFFER_SIZE;
    rb->count--;
    return true;
}
```

---

## 🔗 Theorie-Vertiefung

### Das C-Laufzeitmodell

```
Beim Start eines C-Programms passiert:

1. OS lädt Binary in den Speicher
2. crt0.o (C Runtime Startup) wird ausgeführt:
   - Stack initialisieren
   - .bss auf 0 setzen
   - argc/argv vorbereiten
   - main() aufrufen
3. main() wird ausgeführt
4. Bei return: exit() aufrufen, Rückgabewert an OS

Bei Bare-Metal (OS-Entwicklung):
- KEIN OS das uns hilft!
- WIR müssen crt0 selbst schreiben
- WIR initialisieren Stack und .bss
- WIR rufen unsere kernel_main() auf
```

### Calling Conventions (RISC-V)

Wie werden Funktionsparameter übergeben?

```
┌────────────────────────────────────────────────────────────┐
│  Register  │  Verwendung                                   │
├────────────┼───────────────────────────────────────────────┤
│  a0-a7     │  Funktionsargumente und Rückgabewerte         │
│  t0-t6     │  Temporäre Register (caller-saved)            │
│  s0-s11    │  Saved Register (callee-saved)                │
│  sp        │  Stack Pointer                                │
│  ra        │  Return Address                               │
│  zero      │  Hardwired Zero                               │
└────────────────────────────────────────────────────────────┘
```

---

## 📖 Weiterführende Ressourcen

### C lernen
- **Beej's Guide to C** - [beej.us/guide/bgc](https://beej.us/guide/bgc/) (FREE, excellent)
- **learn-c.org** - [learn-c.org](https://learn-c.org) (Interaktiv)
- **The C Programming Language** (K&R) - Das Standardwerk

### Memory & Hardware
- **What Every Programmer Should Know About Memory** - Ulrich Drepper (PDF)
- **Computer Systems: A Programmer's Perspective** (CS:APP)

### RISC-V
- **RISC-V Reader** - Patterson & Waterman (Free PDF)
- **RISC-V Spec** - [riscv.org/specifications](https://riscv.org/specifications/)

---

## 🧠 Zusammenfassung

| Konzept | Nand2Tetris | Echtes C |
|---------|-------------|----------|
| Speicherzugriff | RAM[addr] = val | `*(uint32_t*)addr = val` |
| I/O | SCREEN[offset] | `volatile` pointer |
| Funktionsaufruf | push/call/pop | Calling convention + Stack |
| Typen | 16-bit alles | int8/16/32/64, signed/unsigned |

**Key Takeaways für morgen:**
1. C gibt dir direkten Speicherzugriff über Pointer
2. `volatile` verhindert Compiler-Optimierungen bei Hardware-Registern
3. Der Build-Prozess: Preprocess → Compile → Assemble → Link
4. Structs können Hardware-Register abbilden

---

*Weiter zu den Übungen → `uebungen/uebung-01.md`*
