# Tag 3: Kernel Basics – Memory Management & Interrupts

**Ziel**: Implementiere Speicherverwaltung und lerne wie Hardware mit deinem Kernel kommuniziert.

---

## 📚 Theorie

### 3.1 Warum Memory Management?

```
┌────────────────────────────────────────────────────────────┐
│  Ohne Memory Management                                    │
├────────────────────────────────────────────────────────────┤
│  - Jeder kann überall schreiben                            │
│  - Kernel-Crash durch fehlerhafte Programme                │
│  - Kein Schutz zwischen Prozessen                          │
│  - Erschöpfung des Speichers ohne Warnung                  │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│  Mit Memory Management                                     │
├────────────────────────────────────────────────────────────┤
│  - Isolation: Programme können sich nicht gegenseitig      │
│    überschreiben                                           │
│  - Schutz: Kernel-Speicher ist geschützt                   │
│  - Virtueller Speicher: Jeder Prozess sieht "seinen"       │
│    eigenen Adressraum                                      │
│  - Allokation: Kontrollierte Speichervergabe               │
└────────────────────────────────────────────────────────────┘
```

### 3.2 Physische vs. Virtuelle Adressen

```
        Prozess A               Prozess B
     ┌────────────┐          ┌────────────┐
     │  0x0000    │          │  0x0000    │
     │    ↓       │          │    ↓       │
     │  Code      │          │  Code      │
     │  Data      │          │  Data      │
     │  Stack     │          │  Stack     │
     └─────┬──────┘          └─────┬──────┘
           │                       │
           │   Page Table A        │   Page Table B
           │   ┌─────────┐         │   ┌─────────┐
           └──▶│0x0→0x100│         └──▶│0x0→0x200│
               │0x1→0x101│             │0x1→0x201│
               └────┬────┘             └────┬────┘
                    │                       │
                    ▼                       ▼
              ┌─────────────────────────────────┐
              │  Physischer RAM                 │
              │  0x100: Code A                  │
              │  0x101: Data A                  │
              │  0x200: Code B                  │
              │  0x201: Data B                  │
              └─────────────────────────────────┘
```

### 3.3 RISC-V Paging (Sv39)

RISC-V verwendet **Sv39**: 39-bit virtuelle Adressen, 3-Level Page Tables.

```
┌──────────────────────────────────────────────────────────────────┐
│  Sv39 Virtual Address (39 bits)                                  │
├─────────────┬─────────────┬─────────────┬────────────────────────┤
│  VPN[2]     │  VPN[1]     │  VPN[0]     │  Page Offset           │
│  (9 bits)   │  (9 bits)   │  (9 bits)   │  (12 bits)             │
│  [38:30]    │  [29:21]    │  [20:12]    │  [11:0]                │
└─────────────┴─────────────┴─────────────┴────────────────────────┘

Page Table Entry (PTE):
┌───────────────────────────────────────────────────────────────────┐
│  Reserved │  PPN[2]   │  PPN[1]   │  PPN[0]   │ RSW │ D│A│G│U│X│W│R│V│
│  [63:54]  │  [53:28]  │  [27:19]  │  [18:10]  │[9:8]│ 7│6│5│4│3│2│1│0│
└───────────────────────────────────────────────────────────────────┘

V = Valid       R = Readable    W = Writable
X = Executable  U = User Mode   A = Accessed
D = Dirty       G = Global
```

### 3.4 Page Table Walk

```
                         SATP Register
                              │
                              ▼
              ┌───────────────────────────┐
              │  Root Page Table (Level 2)│
              └──────────────┬────────────┘
                             │ VPN[2] als Index
                             ▼
              ┌───────────────────────────┐
              │  Page Table (Level 1)     │
              └──────────────┬────────────┘
                             │ VPN[1] als Index  
                             ▼
              ┌───────────────────────────┐
              │  Page Table (Level 0)     │
              └──────────────┬────────────┘
                             │ VPN[0] als Index
                             ▼
              ┌───────────────────────────┐
              │  Physical Page Frame      │
              │  + Page Offset            │
              └───────────────────────────┘
```

### 3.5 Simple Page Frame Allocator

Für einen minimalen Kernel: **Bitmap-basierter Allocator**

```
RAM:  0x80000000 ────────────────────────── 0x88000000
      [Kernel Code/Data][Free Frames......................]
      
Bitmap:  1 1 1 1 1 0 0 0 1 0 0 0 0 0 0 0 ...
         │ │ │ │ │     │
         │ │ │ │ │     Frame 8: belegt (Kernel Stack)
         │ │ │ │ Frame 5-7: frei
         Frames 0-4: Kernel Code (belegt)

Jedes Bit = 1 Page (4KB)
```

### 3.6 Interrupts und Exceptions

Hardware muss den Kernel informieren können:

```
┌─────────────────────────────────────────────────────────────┐
│  Exceptions (Synchron - durch CPU ausgelöst)                │
├─────────────────────────────────────────────────────────────┤
│  - Illegal Instruction                                      │
│  - Load/Store Access Fault                                  │
│  - Environment Call (syscall)                               │
│  - Page Fault                                               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│  Interrupts (Asynchron - durch Hardware ausgelöst)          │
├─────────────────────────────────────────────────────────────┤
│  - Timer Interrupt                                          │
│  - External Interrupt (UART, Keyboard, etc.)                │
│  - Software Interrupt                                       │
└─────────────────────────────────────────────────────────────┘
```

### 3.7 RISC-V Trap Handling

```
┌─────────────────────────────────────────────────────────────┐
│  1. Exception/Interrupt tritt auf                           │
│  2. CPU speichert:                                          │
│     - mepc: Adresse der unterbrochenen Instruktion          │
│     - mcause: Grund für den Trap                            │
│     - mtval: Zusätzliche Info (z.B. fehlgeschlagene Adresse)│
│  3. CPU springt zu mtvec (Trap Vector)                      │
│  4. Trap Handler wird ausgeführt                            │
│  5. mret kehrt zum normalen Code zurück                     │
└─────────────────────────────────────────────────────────────┘
```

### 3.8 mcause Register

```
┌─────────────────────────────────────────────────────────────┐
│  Bit 63 (Interrupt Bit)                                     │
│    0 = Exception                                            │
│    1 = Interrupt                                            │
├─────────────────────────────────────────────────────────────┤
│  Exception Codes (Bit 63 = 0):                              │
│    0  = Instruction Address Misaligned                      │
│    1  = Instruction Access Fault                            │
│    2  = Illegal Instruction                                 │
│    5  = Load Access Fault                                   │
│    7  = Store Access Fault                                  │
│    8  = Environment Call from U-mode                        │
│    11 = Environment Call from M-mode                        │
│    12 = Instruction Page Fault                              │
│    13 = Load Page Fault                                     │
│    15 = Store Page Fault                                    │
├─────────────────────────────────────────────────────────────┤
│  Interrupt Codes (Bit 63 = 1):                              │
│    3 = Machine Software Interrupt                           │
│    7 = Machine Timer Interrupt                              │
│    11 = Machine External Interrupt                          │
└─────────────────────────────────────────────────────────────┘
```

---

## 💻 Code-Beispiele

### Page Frame Allocator

```c
/* palloc.c - Physical Page Allocator */

#include <stdint.h>
#include <stdbool.h>

#define PAGE_SIZE 4096
#define RAM_START 0x80000000
#define RAM_SIZE  (128 * 1024 * 1024)  /* 128 MB */
#define NUM_PAGES (RAM_SIZE / PAGE_SIZE)

/* Bitmap: 1 bit per page */
static uint8_t page_bitmap[NUM_PAGES / 8];

/* Ende des Kernels (vom Linker) */
extern char __kernel_end;

void palloc_init(void) {
    /* Alle Pages als frei markieren */
    for (int i = 0; i < NUM_PAGES / 8; i++) {
        page_bitmap[i] = 0;
    }
    
    /* Kernel-Pages als belegt markieren */
    uint64_t kernel_end = (uint64_t)&__kernel_end;
    uint64_t kernel_pages = (kernel_end - RAM_START + PAGE_SIZE - 1) / PAGE_SIZE;
    
    for (uint64_t i = 0; i < kernel_pages; i++) {
        uint64_t byte_idx = i / 8;
        uint8_t bit_idx = i % 8;
        page_bitmap[byte_idx] |= (1 << bit_idx);
    }
}

/* Alloziere eine physische Page, return Adresse oder 0 bei Fehler */
uint64_t palloc(void) {
    for (uint64_t i = 0; i < NUM_PAGES; i++) {
        uint64_t byte_idx = i / 8;
        uint8_t bit_idx = i % 8;
        
        if ((page_bitmap[byte_idx] & (1 << bit_idx)) == 0) {
            /* Frei! Markiere als belegt */
            page_bitmap[byte_idx] |= (1 << bit_idx);
            return RAM_START + (i * PAGE_SIZE);
        }
    }
    return 0;  /* Out of memory */
}

/* Gib eine Page frei */
void pfree(uint64_t addr) {
    if (addr < RAM_START || addr >= RAM_START + RAM_SIZE) {
        return;  /* Ungültige Adresse */
    }
    
    uint64_t page_num = (addr - RAM_START) / PAGE_SIZE;
    uint64_t byte_idx = page_num / 8;
    uint8_t bit_idx = page_num % 8;
    
    page_bitmap[byte_idx] &= ~(1 << bit_idx);
}

/* Zeige Statistiken */
void palloc_stats(void) {
    uint64_t used = 0, free_pages = 0;
    
    for (uint64_t i = 0; i < NUM_PAGES; i++) {
        uint64_t byte_idx = i / 8;
        uint8_t bit_idx = i % 8;
        
        if (page_bitmap[byte_idx] & (1 << bit_idx)) {
            used++;
        } else {
            free_pages++;
        }
    }
    
    /* Ausgabe über uart_puts/print_int */
    uart_puts("Memory: ");
    print_int(used * 4);
    uart_puts(" KB used, ");
    print_int(free_pages * 4);
    uart_puts(" KB free\n");
}
```

### Trap Handler

```c
/* trap.c - Interrupt and Exception Handling */

#include <stdint.h>

/* CSR Lese-Makros */
#define read_csr(csr) ({ \
    uint64_t val; \
    asm volatile ("csrr %0, " #csr : "=r"(val)); \
    val; \
})

#define write_csr(csr, val) ({ \
    asm volatile ("csrw " #csr ", %0" :: "r"(val)); \
})

/* Exception Namen */
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

/* Trap Handler - wird von Assembly aufgerufen */
void trap_handler(void) {
    uint64_t mcause = read_csr(mcause);
    uint64_t mepc = read_csr(mepc);
    uint64_t mtval = read_csr(mtval);
    
    int is_interrupt = (mcause >> 63) & 1;
    uint64_t code = mcause & 0x7FFFFFFFFFFFFFFF;
    
    if (is_interrupt) {
        /* Interrupt */
        switch (code) {
            case 7:  /* Timer */
                handle_timer_interrupt();
                break;
            case 11: /* External */
                handle_external_interrupt();
                break;
            default:
                uart_puts("Unknown interrupt: ");
                print_int(code);
                uart_puts("\n");
        }
    } else {
        /* Exception */
        uart_puts("\n!!! EXCEPTION !!!\n");
        uart_puts("Type: ");
        if (code < 16) {
            uart_puts(exception_names[code]);
        } else {
            uart_puts("Unknown");
        }
        uart_puts("\n");
        
        uart_puts("mepc:  ");
        print_hex(mepc);
        uart_puts("\n");
        
        uart_puts("mtval: ");
        print_hex(mtval);
        uart_puts("\n");
        
        /* Bei Exception: Endlosschleife */
        while (1) {}
    }
}

/* Timer Interrupt Handler */
void handle_timer_interrupt(void) {
    /* Setze nächsten Timer */
    #define CLINT_MTIMECMP(hart) (*(volatile uint64_t *)(0x02004000 + 8*(hart)))
    #define CLINT_MTIME (*(volatile uint64_t *)0x0200BFF8)
    
    CLINT_MTIMECMP(0) = CLINT_MTIME + 10000000;  /* 1 Sekunde bei 10MHz */
    
    /* Debug-Ausgabe */
    uart_puts("TICK\n");
}

/* External Interrupt Handler */
void handle_external_interrupt(void) {
    /* TODO: PLIC abfragen und entsprechendes Gerät behandeln */
}

/* Initialisiere Trap Handling */
void trap_init(void) {
    /* Setze Trap Vector */
    extern void trap_entry(void);
    write_csr(mtvec, (uint64_t)trap_entry);
    
    /* Aktiviere Machine Interrupts */
    /* mstatus.MIE = 1 */
    uint64_t mstatus = read_csr(mstatus);
    mstatus |= (1 << 3);  /* MIE bit */
    write_csr(mstatus, mstatus);
    
    /* Aktiviere Timer Interrupt */
    /* mie.MTIE = 1 */
    uint64_t mie = read_csr(mie);
    mie |= (1 << 7);  /* MTIE bit */
    write_csr(mie, mie);
}
```

### Trap Entry (Assembly)

```asm
# trap_entry.S - Low-level trap handling

.section .text
.global trap_entry
.align 4

trap_entry:
    # Speichere alle Register auf dem Stack
    addi sp, sp, -256
    
    sd ra,   0(sp)
    sd sp,   8(sp)
    sd gp,  16(sp)
    sd tp,  24(sp)
    sd t0,  32(sp)
    sd t1,  40(sp)
    sd t2,  48(sp)
    sd s0,  56(sp)
    sd s1,  64(sp)
    sd a0,  72(sp)
    sd a1,  80(sp)
    sd a2,  88(sp)
    sd a3,  96(sp)
    sd a4, 104(sp)
    sd a5, 112(sp)
    sd a6, 120(sp)
    sd a7, 128(sp)
    sd s2, 136(sp)
    sd s3, 144(sp)
    sd s4, 152(sp)
    sd s5, 160(sp)
    sd s6, 168(sp)
    sd s7, 176(sp)
    sd s8, 184(sp)
    sd s9, 192(sp)
    sd s10, 200(sp)
    sd s11, 208(sp)
    sd t3, 216(sp)
    sd t4, 224(sp)
    sd t5, 232(sp)
    sd t6, 240(sp)
    
    # Rufe C Handler auf
    call trap_handler
    
    # Stelle Register wieder her
    ld ra,   0(sp)
    ld sp,   8(sp)
    ld gp,  16(sp)
    ld tp,  24(sp)
    ld t0,  32(sp)
    ld t1,  40(sp)
    ld t2,  48(sp)
    ld s0,  56(sp)
    ld s1,  64(sp)
    ld a0,  72(sp)
    ld a1,  80(sp)
    ld a2,  88(sp)
    ld a3,  96(sp)
    ld a4, 104(sp)
    ld a5, 112(sp)
    ld a6, 120(sp)
    ld a7, 128(sp)
    ld s2, 136(sp)
    ld s3, 144(sp)
    ld s4, 152(sp)
    ld s5, 160(sp)
    ld s6, 168(sp)
    ld s7, 176(sp)
    ld s8, 184(sp)
    ld s9, 192(sp)
    ld s10, 200(sp)
    ld s11, 208(sp)
    ld t3, 216(sp)
    ld t4, 224(sp)
    ld t5, 232(sp)
    ld t6, 240(sp)
    
    addi sp, sp, 256
    
    # Kehre von Trap zurück
    mret
```

---

## 📖 Weiterführende Ressourcen

### Memory Management
- **RISC-V Privileged Spec** - Kapitel über Virtuellen Speicher
- **Operating Systems: Three Easy Pieces** - [ostep.org](https://pages.cs.wisc.edu/~remzi/OSTEP/) (FREE)

### Interrupts
- **RISC-V OS in 1000 Lines** - Kapitel Interrupts
- **xv6** - Einfaches Lehr-OS mit vollem Quellcode

---

## 🧠 Zusammenfassung

| Konzept | Was du gelernt hast |
|---------|---------------------|
| Physical Pages | RAM wird in 4KB Blöcke aufgeteilt |
| Page Allocator | Bitmap trackt belegte/freie Pages |
| Virtual Memory | Sv39: 3-Level Page Tables für Adress-Translation |
| Traps | Exceptions (synchron) vs Interrupts (asynchron) |
| CSRs | mcause, mepc, mtvec, mstatus, mie |

---

*Weiter zu den Übungen → `uebungen/uebung-03.md`*
