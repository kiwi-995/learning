# Übung 3: Kernel Basics – Memory & Interrupts

**Ziel**: Implementiere einen Page Allocator und Trap Handler.

---

## 🛠️ Setup

```bash
cd ~/os-dev/day2/mini-kernel
# Wir bauen auf dem Kernel von Tag 2 auf
```

---

## Aufgabe 1: Page Frame Allocator (⏱️ 90 min)

Implementiere einen einfachen physischen Speicher-Allocator.

### 1a: Header-Datei

**Datei**: `palloc.h`

```c
#ifndef PALLOC_H
#define PALLOC_H

#include <stdint.h>

#define PAGE_SIZE 4096

/* Initialisiere den Allocator */
void palloc_init(void);

/* Alloziere eine 4KB Page, return Adresse oder 0 */
uint64_t palloc(void);

/* Gib eine Page frei */
void pfree(uint64_t addr);

/* Zeige Speicherstatistik */
void palloc_stats(void);

/* Wie viele Pages sind frei? */
uint64_t palloc_free_count(void);

#endif
```

### 1b: Implementierung

**Datei**: `palloc.c`

```c
#include "palloc.h"
#include <stdint.h>

/* Memory Konfiguration */
#define RAM_START 0x80000000UL
#define RAM_SIZE  (128 * 1024 * 1024)  /* 128 MB */
#define NUM_PAGES (RAM_SIZE / PAGE_SIZE)

/* Bitmap: 1 bit pro page (0 = frei, 1 = belegt) */
static uint8_t page_bitmap[NUM_PAGES / 8];

/* Kernel-Ende (muss im Linker Script definiert werden!) */
extern char __kernel_end;

void palloc_init(void) {
    /* TODO: Implementiere!
     * 1. Setze alle Bits auf 0 (alles frei)
     * 2. Berechne wie viele Pages der Kernel belegt
     * 3. Markiere Kernel-Pages als belegt
     */
}

uint64_t palloc(void) {
    /* TODO: Implementiere!
     * 1. Suche erstes freies Bit in Bitmap
     * 2. Markiere als belegt
     * 3. Return physische Adresse
     * 4. Wenn nichts frei: return 0
     */
    return 0;
}

void pfree(uint64_t addr) {
    /* TODO: Implementiere!
     * 1. Prüfe ob Adresse gültig
     * 2. Berechne Page-Nummer
     * 3. Setze Bit auf 0
     */
}

uint64_t palloc_free_count(void) {
    /* TODO: Zähle freie Pages */
    return 0;
}

void palloc_stats(void) {
    /* TODO: Gib Statistik aus */
    /* Benutze uart_puts und print_int */
}
```

### 1c: Linker Script Update

Füge zu `kernel.ld` hinzu:

```ld
/* Am Ende der SECTIONS, vor dem Stack */
    . = ALIGN(4096);
    __kernel_end = .;
```

### Test

```c
/* In kernel_main() */
void kernel_main(void) {
    uart_puts("Initializing memory...\n");
    palloc_init();
    palloc_stats();
    
    uart_puts("\nAllocating 3 pages:\n");
    uint64_t p1 = palloc();
    uint64_t p2 = palloc();
    uint64_t p3 = palloc();
    
    uart_puts("Page 1: "); print_hex(p1); uart_puts("\n");
    uart_puts("Page 2: "); print_hex(p2); uart_puts("\n");
    uart_puts("Page 3: "); print_hex(p3); uart_puts("\n");
    
    palloc_stats();
    
    uart_puts("\nFreeing page 2...\n");
    pfree(p2);
    palloc_stats();
    
    uart_puts("\nAllocating 1 page:\n");
    uint64_t p4 = palloc();
    uart_puts("Page 4: "); print_hex(p4); uart_puts("\n");
    uart_puts("(Should reuse freed page 2's address)\n");
    
    while (1) {}
}
```

---

## Aufgabe 2: Trap Handler Setup (⏱️ 60 min)

Implementiere grundlegendes Exception Handling.

### 2a: CSR Makros

**Datei**: `riscv.h`

```c
#ifndef RISCV_H
#define RISCV_H

#include <stdint.h>

/* CSR Lesen */
#define read_csr(csr) ({ \
    uint64_t __val; \
    asm volatile ("csrr %0, " #csr : "=r"(__val)); \
    __val; \
})

/* CSR Schreiben */
#define write_csr(csr, val) ({ \
    asm volatile ("csrw " #csr ", %0" :: "r"(val)); \
})

/* CSR Bits setzen */
#define set_csr(csr, val) ({ \
    asm volatile ("csrs " #csr ", %0" :: "r"(val)); \
})

/* CSR Bits löschen */
#define clear_csr(csr, val) ({ \
    asm volatile ("csrc " #csr ", %0" :: "r"(val)); \
})

/* Wichtige CSR Bits */
#define MSTATUS_MIE (1 << 3)   /* Machine Interrupt Enable */
#define MSTATUS_MPIE (1 << 7)  /* Previous MIE */

#define MIE_MTIE (1 << 7)      /* Machine Timer Interrupt Enable */
#define MIE_MEIE (1 << 11)     /* Machine External Interrupt Enable */

#endif
```

### 2b: Trap Entry Assembly

**Datei**: `trap_entry.S`

```asm
.section .text
.global trap_entry
.align 4

trap_entry:
    # TODO: Speichere Register auf Stack
    # Hint: addi sp, sp, -256 für Platz
    # Hint: sd ra, 0(sp) für jedes Register
    
    # TODO: Rufe C trap_handler auf
    # Hint: call trap_handler
    
    # TODO: Stelle Register wieder her
    # Hint: ld ra, 0(sp) für jedes Register
    # Hint: addi sp, sp, 256
    
    # TODO: Return from trap
    # Hint: mret
    
    j trap_entry  # Placeholder, ersetze mit echtem Code
```

### 2c: C Trap Handler

**Datei**: `trap.c`

```c
#include <stdint.h>
#include "riscv.h"

/* Externe Funktionen */
extern void uart_puts(const char *s);
extern void print_hex(uint64_t val);
extern void print_int(int val);

/* Exception-Namen Tabelle */
static const char *exception_names[] = {
    "Instruction Address Misaligned",
    "Instruction Access Fault", 
    "Illegal Instruction",
    "Breakpoint",
    "Load Address Misaligned",
    "Load Access Fault",
    "Store Address Misaligned",
    "Store Access Fault",
    "Environment Call (U-mode)",
    "Environment Call (S-mode)",
    "Reserved",
    "Environment Call (M-mode)",
    "Instruction Page Fault",
    "Load Page Fault",
    "Reserved",
    "Store Page Fault"
};

/* Trap Handler - von Assembly aufgerufen */
void trap_handler(void) {
    uint64_t mcause = read_csr(mcause);
    uint64_t mepc = read_csr(mepc);
    uint64_t mtval = read_csr(mtval);
    
    /* TODO: Implementiere!
     * 1. Prüfe ob Interrupt (Bit 63) oder Exception
     * 2. Bei Interrupt: Handle Timer/External
     * 3. Bei Exception: Zeige Info und halte an
     */
    
    int is_interrupt = /* TODO: Extrahiere Bit 63 */;
    uint64_t code = /* TODO: Extrahiere Exception/Interrupt Code */;
    
    if (is_interrupt) {
        /* Interrupt */
        if (code == 7) {
            /* Timer Interrupt */
            uart_puts("[TIMER] ");
            /* TODO: Setze nächsten Timer */
        } else {
            uart_puts("Unknown interrupt\n");
        }
    } else {
        /* Exception */
        uart_puts("\n!!! EXCEPTION !!!\n");
        uart_puts("Type: ");
        if (code < 16) {
            uart_puts(exception_names[code]);
        }
        uart_puts("\nmepc: ");
        print_hex(mepc);
        uart_puts("\nmtval: ");
        print_hex(mtval);
        uart_puts("\n");
        
        /* Halt */
        while (1) {}
    }
}

/* Initialisiere Trap Handling */
void trap_init(void) {
    extern void trap_entry(void);
    
    /* TODO: Setze mtvec auf trap_entry Adresse */
    
    /* TODO: Aktiviere Machine Interrupts (mstatus.MIE) */
}
```

### Test

```c
void kernel_main(void) {
    uart_puts("Setting up traps...\n");
    trap_init();
    
    uart_puts("Testing illegal instruction...\n");
    
    /* Trigger Exception: Ungültige Instruktion */
    asm volatile (".word 0x00000000");  /* Ungültige Instruktion */
    
    uart_puts("This should not print!\n");
}
```

**Erwartete Ausgabe**:
```
Setting up traps...
Testing illegal instruction...

!!! EXCEPTION !!!
Type: Illegal Instruction
mepc: 0x80000XXX
mtval: 0x00000000
```

---

## Aufgabe 3: Timer Interrupt (⏱️ 45 min)

Implementiere einen periodischen Timer.

**Datei**: `timer.c`

```c
#include <stdint.h>
#include "riscv.h"

/* CLINT Register */
#define CLINT_MTIME     (*(volatile uint64_t *)0x0200BFF8)
#define CLINT_MTIMECMP  (*(volatile uint64_t *)0x02004000)

/* Timer Frequenz (QEMU virt: 10 MHz) */
#define TIMER_FREQ 10000000

/* Tick Counter */
static volatile uint64_t ticks = 0;

/* Initialisiere Timer */
void timer_init(void) {
    /* TODO: Setze ersten Timer Interrupt */
    /* CLINT_MTIMECMP = CLINT_MTIME + TIMER_FREQ */
    
    /* TODO: Aktiviere Timer Interrupt (MIE.MTIE) */
}

/* Timer Interrupt Handler (von trap_handler aufgerufen) */
void timer_tick(void) {
    ticks++;
    
    /* TODO: Setze nächsten Timer */
    /* CLINT_MTIMECMP = CLINT_MTIME + TIMER_FREQ */
    
    /* Debug: Zeige Tick alle 5 Sekunden */
    if (ticks % 5 == 0) {
        extern void uart_puts(const char *s);
        extern void print_int(int val);
        uart_puts("[");
        print_int(ticks);
        uart_puts(" sec] ");
    }
}

/* Hole aktuelle Ticks */
uint64_t get_ticks(void) {
    return ticks;
}

/* Warte n Ticks */
void wait_ticks(uint64_t n) {
    uint64_t target = ticks + n;
    while (ticks < target) {
        asm volatile ("wfi");  /* Wait for interrupt */
    }
}
```

### Update trap_handler

```c
/* In trap_handler, bei Interrupt code == 7: */
if (code == 7) {
    extern void timer_tick(void);
    timer_tick();
}
```

### Test

```c
void kernel_main(void) {
    uart_puts("Mini Kernel v0.2\n");
    
    palloc_init();
    trap_init();
    timer_init();
    
    uart_puts("Timer running. Waiting 10 seconds...\n");
    
    extern void wait_ticks(uint64_t);
    wait_ticks(10);
    
    uart_puts("\n10 seconds passed!\n");
    
    while (1) {
        asm volatile ("wfi");
    }
}
```

---

## Aufgabe 4: Einfacher Heap Allocator (⏱️ 60 min)

Baue auf dem Page Allocator auf für kleinere Allokationen.

**Datei**: `heap.h`

```c
#ifndef HEAP_H
#define HEAP_H

#include <stdint.h>

void heap_init(void);
void *kmalloc(uint64_t size);
void kfree(void *ptr);

#endif
```

**Datei**: `heap.c`

```c
#include "heap.h"
#include "palloc.h"

/* Einfacher Bump Allocator:
 * - Alloziert Pages vom palloc
 * - Bump Pointer innerhalb der Pages
 * - kfree ist ein No-Op (Speicher wird nicht zurückgegeben)
 * 
 * Für echtes OS: Implementiere freelist oder slab allocator
 */

static uint64_t heap_start = 0;
static uint64_t heap_end = 0;
static uint64_t heap_ptr = 0;

void heap_init(void) {
    /* Initiale Page allozieren */
    heap_start = palloc();
    if (heap_start == 0) {
        /* Out of memory */
        return;
    }
    heap_end = heap_start + PAGE_SIZE;
    heap_ptr = heap_start;
}

void *kmalloc(uint64_t size) {
    /* TODO: Implementiere!
     * 1. Alignment auf 8 Bytes
     * 2. Prüfe ob genug Platz
     * 3. Wenn nicht: Neue Page allozieren
     * 4. Bump heap_ptr
     * 5. Return alte heap_ptr Position
     */
    
    /* Alignment */
    size = (size + 7) & ~7UL;
    
    /* TODO: Rest implementieren */
    
    return 0;  /* NULL */
}

void kfree(void *ptr) {
    /* Bump allocator: kfree macht nichts */
    (void)ptr;
}
```

### Test

```c
void test_heap(void) {
    heap_init();
    
    int *numbers = kmalloc(10 * sizeof(int));
    if (numbers == 0) {
        uart_puts("kmalloc failed!\n");
        return;
    }
    
    for (int i = 0; i < 10; i++) {
        numbers[i] = i * i;
    }
    
    uart_puts("Squares: ");
    for (int i = 0; i < 10; i++) {
        print_int(numbers[i]);
        uart_puts(" ");
    }
    uart_puts("\n");
}
```

---

## Aufgabe 5: String Library (⏱️ 30 min)

Implementiere grundlegende String-Funktionen (ohne libc!).

**Datei**: `string.c`

```c
#include <stdint.h>

/* String-Länge */
uint64_t strlen(const char *s) {
    /* TODO */
    return 0;
}

/* String kopieren */
char *strcpy(char *dest, const char *src) {
    /* TODO */
    return dest;
}

/* String vergleichen */
int strcmp(const char *s1, const char *s2) {
    /* TODO: Return 0 wenn gleich, <0 wenn s1<s2, >0 wenn s1>s2 */
    return 0;
}

/* Speicher setzen */
void *memset(void *s, int c, uint64_t n) {
    /* TODO */
    return s;
}

/* Speicher kopieren */
void *memcpy(void *dest, const void *src, uint64_t n) {
    /* TODO */
    return dest;
}

/* Speicher vergleichen */
int memcmp(const void *s1, const void *s2, uint64_t n) {
    /* TODO */
    return 0;
}
```

---

## 🎯 Bonusaufgaben

### Bonus A: Kernel Panic

Implementiere eine `panic(const char *msg)` Funktion die eine Fehlermeldung ausgibt und den Kernel anhält.

### Bonus B: Stack Trace

Bei einer Exception, versuche den Call Stack auszugeben (folge Frame Pointern).

### Bonus C: Page Table Setup

Erstelle eine einfache Identity-Mapped Page Table (virtuelle Adresse = physische Adresse).

---

## ✅ Checkliste

- [ ] Page Frame Allocator funktioniert
- [ ] Exceptions werden gefangen und angezeigt
- [ ] Timer Interrupt funktioniert
- [ ] Heap Allocator (kmalloc) funktioniert
- [ ] String-Funktionen implementiert

**Morgen**: Terminal mit Keyboard-Input!

---

*Weiter zu Tag 4 → `vorlesungen/tag-04-terminal.md`*
