# Tag 2: Bare Metal RISC-V – Dein erstes Bootup

**Ziel**: Boot deinen ersten eigenen Code auf RISC-V in QEMU und gib "Hello" über UART aus.

---

## 📚 Theorie

### 2.1 Was bedeutet "Bare Metal"?

```
┌─────────────────────────────────────────────────────────────┐
│  Normale Anwendung                                          │
├──────────────────┬──────────────────────────────────────────┤
│  Dein Programm   │  Andere Programme                        │
├──────────────────┴──────────────────────────────────────────┤
│  Betriebssystem (Linux, macOS, Windows)                     │
├─────────────────────────────────────────────────────────────┤
│  Hardware (CPU, RAM, Geräte)                                │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│  Bare Metal                                                 │
├─────────────────────────────────────────────────────────────┤
│  Dein Programm (IST das "OS")                               │
├─────────────────────────────────────────────────────────────┤
│  Hardware (CPU, RAM, Geräte)                                │
└─────────────────────────────────────────────────────────────┘
```

**Kein OS** bedeutet:
- Kein `printf()` (keine libc!)
- Kein Filesystem
- Kein Speicherschutz
- Kein Multitasking
- Direkte Hardware-Kontrolle

### 2.2 Die RISC-V Architektur

RISC-V ist eine offene Instruction Set Architecture (ISA):

```
┌────────────────────────────────────────────────────────────┐
│  RISC-V Privilege Levels                                   │
├──────────┬────────────────────────────────────────────────┤
│  U-Mode  │  User Mode (normale Programme)                 │
├──────────┼────────────────────────────────────────────────┤
│  S-Mode  │  Supervisor Mode (OS Kernel)                   │
├──────────┼────────────────────────────────────────────────┤
│  M-Mode  │  Machine Mode (Firmware, höchste Privilegien)  │
└──────────┴────────────────────────────────────────────────┘
```

Wir starten in **M-Mode** – volle Kontrolle über alles!

### 2.3 RISC-V Register

```
┌────────┬─────────┬───────────────────────────────────────┐
│ Nummer │  Name   │  Verwendung                           │
├────────┼─────────┼───────────────────────────────────────┤
│  x0    │  zero   │  Hardwired Zero (lesen=0, schreiben=noop) │
│  x1    │  ra     │  Return Address                       │
│  x2    │  sp     │  Stack Pointer                        │
│  x3    │  gp     │  Global Pointer                       │
│  x4    │  tp     │  Thread Pointer                       │
│  x5-7  │  t0-t2  │  Temporaries                          │
│  x8    │  s0/fp  │  Saved Register / Frame Pointer       │
│  x9    │  s1     │  Saved Register                       │
│  x10-11│  a0-a1  │  Function Arguments / Return Values   │
│  x12-17│  a2-a7  │  Function Arguments                   │
│  x18-27│  s2-s11 │  Saved Registers                      │
│  x28-31│  t3-t6  │  Temporaries                          │
└────────┴─────────┴───────────────────────────────────────┘
```

### 2.4 QEMU Virt Machine

QEMU emuliert eine "virt" Machine mit definiertem Memory Map:

```
┌──────────────────────────────────────────────────────────┐
│  Adresse          │  Gerät                               │
├───────────────────┼──────────────────────────────────────┤
│  0x00001000       │  Boot ROM                            │
│  0x02000000       │  CLINT (Core Local Interrupter)      │
│  0x0C000000       │  PLIC (Platform Level Interrupt)     │
│  0x10000000       │  UART0 (Serielle Konsole)            │
│  0x10001000       │  VirtIO Block Device                 │
│  0x80000000       │  RAM Start (wo unser Code landet!)   │
└───────────────────┴──────────────────────────────────────┘
```

### 2.5 Der Boot-Prozess

```
┌─────────────────────────────────────────────────────────┐
│ 1. QEMU startet                                         │
│    ↓                                                    │
│ 2. Lädt unseren Code an 0x80000000                      │
│    ↓                                                    │
│ 3. Setzt PC (Program Counter) auf 0x80000000            │
│    ↓                                                    │
│ 4. Beginnt Instruktionen auszuführen                    │
│    ↓                                                    │
│ 5. Unsere erste Instruktion läuft!                      │
└─────────────────────────────────────────────────────────┘
```

### 2.6 UART: Serielle Kommunikation

UART (Universal Asynchronous Receiver-Transmitter) ist der einfachste Weg für Ausgabe:

```
┌─────────────────────────────────────────────────────────┐
│  UART Register bei 0x10000000                           │
├─────────────────────────────────────────────────────────┤
│  Offset 0x00: THR (Transmit Holding Register)           │
│               - Schreibe hier um Zeichen zu senden      │
│  Offset 0x00: RBR (Receive Buffer Register)             │
│               - Lese hier um Zeichen zu empfangen       │
│  Offset 0x05: LSR (Line Status Register)                │
│               - Bit 0: Data Ready (zum Lesen)           │
│               - Bit 5: THR Empty (bereit zum Schreiben) │
└─────────────────────────────────────────────────────────┘
```

---

## 💻 Code-Beispiele

### Das minimale Boot-Setup

#### Linker Script (`kernel.ld`)

Der Linker Script sagt dem Linker wo was im Speicher liegt:

```ld
/* kernel.ld - Linker Script für RISC-V QEMU virt */

OUTPUT_ARCH(riscv)
ENTRY(_start)

MEMORY {
    /* RAM beginnt bei 0x80000000, wir nehmen 128MB */
    RAM (rwx) : ORIGIN = 0x80000000, LENGTH = 128M
}

SECTIONS {
    . = 0x80000000;

    .text : {
        *(.text.boot)    /* Boot-Code zuerst! */
        *(.text .text.*)
    } > RAM

    .rodata : {
        *(.rodata .rodata.*)
    } > RAM

    .data : {
        *(.data .data.*)
    } > RAM

    .bss : {
        __bss_start = .;
        *(.bss .bss.*)
        *(COMMON)
        __bss_end = .;
    } > RAM

    /* Stack am Ende des RAM */
    . = ALIGN(16);
    . = . + 0x10000;   /* 64KB Stack */
    __stack_top = .;
}
```

#### Boot-Code (`boot.S`)

Die allerersten Instruktionen müssen in Assembly sein:

```asm
# boot.S - RISC-V Boot-Code

.section .text.boot
.global _start

_start:
    # Setze Stack Pointer
    la sp, __stack_top
    
    # Lösche BSS Segment (uninitialisierte Variablen auf 0)
    la t0, __bss_start
    la t1, __bss_end
clear_bss:
    beq t0, t1, done_bss
    sd zero, 0(t0)
    addi t0, t0, 8
    j clear_bss
done_bss:
    
    # Springe zu C-Code
    call kernel_main
    
    # Falls kernel_main zurückkehrt, halte an
halt:
    wfi              # Wait For Interrupt
    j halt
```

#### Kernel Main (`kernel.c`)

```c
/* kernel.c - Minimaler Kernel */

#include <stdint.h>

/* UART Basisadresse für QEMU virt */
#define UART_BASE 0x10000000

/* UART Register */
#define UART_THR (*(volatile uint8_t *)(UART_BASE + 0x00)) /* Transmit */
#define UART_LSR (*(volatile uint8_t *)(UART_BASE + 0x05)) /* Status */

/* LSR Bits */
#define LSR_TX_EMPTY 0x20

void uart_putc(char c) {
    /* Warte bis Transmitter bereit */
    while ((UART_LSR & LSR_TX_EMPTY) == 0) {}
    UART_THR = c;
}

void uart_puts(const char *s) {
    while (*s) {
        if (*s == '\n') {
            uart_putc('\r');  /* Carriage Return vor Newline */
        }
        uart_putc(*s++);
    }
}

void kernel_main(void) {
    uart_puts("Hello from RISC-V!\n");
    uart_puts("My first bare-metal kernel is running!\n");
    
    /* Endlosschleife */
    while (1) {}
}
```

### Makefile

```makefile
# Makefile für RISC-V Kernel

CROSS = riscv64-unknown-elf-
CC = $(CROSS)gcc
AS = $(CROSS)as
LD = $(CROSS)ld
OBJCOPY = $(CROSS)objcopy

CFLAGS = -Wall -Wextra -O2 -ffreestanding -nostdlib -mcmodel=medany
ASFLAGS = -march=rv64imac -mabi=lp64

all: kernel.bin

boot.o: boot.S
	$(AS) $(ASFLAGS) -o $@ $<

kernel.o: kernel.c
	$(CC) $(CFLAGS) -c -o $@ $<

kernel.elf: boot.o kernel.o kernel.ld
	$(LD) -T kernel.ld -o $@ boot.o kernel.o

kernel.bin: kernel.elf
	$(OBJCOPY) -O binary $< $@

run: kernel.bin
	qemu-system-riscv64 -machine virt -bios none -kernel kernel.bin -nographic

clean:
	rm -f *.o *.elf *.bin

.PHONY: all run clean
```

---

## 🔧 QEMU Setup

### Installation (macOS)

```bash
# Via Homebrew
brew install qemu

# Teste Installation
qemu-system-riscv64 --version
```

### RISC-V Toolchain

```bash
# Option 1: Homebrew (empfohlen)
brew tap riscv-software-src/riscv
brew install riscv-gnu-toolchain

# Option 2: Fertige Binaries
# Download von: https://github.com/riscv-collab/riscv-gnu-toolchain/releases

# Teste Toolchain
riscv64-unknown-elf-gcc --version
```

### Alternative: Docker

Falls die Toolchain nicht funktioniert:

```bash
# Dockerfile
FROM ubuntu:22.04
RUN apt-get update && apt-get install -y \
    gcc-riscv64-unknown-elf \
    qemu-system-riscv64 \
    make
WORKDIR /kernel

# Baue und starte
docker build -t riscv-dev .
docker run -it -v $(pwd):/kernel riscv-dev make run
```

---

## 🔬 Assembly Crashkurs (RISC-V)

### Grundlegende Instruktionen

```asm
# Laden und Speichern
li  a0, 42          # Load Immediate: a0 = 42
la  a0, label       # Load Address: a0 = Adresse von label
lw  a0, 0(a1)       # Load Word: a0 = *(int32_t*)(a1 + 0)
ld  a0, 0(a1)       # Load Doubleword: a0 = *(int64_t*)(a1 + 0)
sw  a0, 0(a1)       # Store Word: *(int32_t*)(a1 + 0) = a0
sd  a0, 0(a1)       # Store Doubleword

# Arithmetik
add  a0, a1, a2     # a0 = a1 + a2
addi a0, a1, 10     # a0 = a1 + 10
sub  a0, a1, a2     # a0 = a1 - a2
mul  a0, a1, a2     # a0 = a1 * a2

# Logik
and  a0, a1, a2     # a0 = a1 & a2
or   a0, a1, a2     # a0 = a1 | a2
xor  a0, a1, a2     # a0 = a1 ^ a2
sll  a0, a1, a2     # a0 = a1 << a2 (Shift Left Logical)
srl  a0, a1, a2     # a0 = a1 >> a2 (Shift Right Logical)

# Vergleiche und Sprünge
beq  a0, a1, label  # Branch if Equal: if (a0 == a1) goto label
bne  a0, a1, label  # Branch if Not Equal
blt  a0, a1, label  # Branch if Less Than
bge  a0, a1, label  # Branch if Greater or Equal
j    label          # Jump (unconditional)

# Funktionen
call function       # Aufruf (setzt ra)
ret                 # Return (springt zu ra)
```

### C und Assembly kombinieren

```c
/* Inline Assembly in C */

/* Lese Machine Status Register */
uint64_t read_mstatus(void) {
    uint64_t value;
    asm volatile ("csrr %0, mstatus" : "=r"(value));
    return value;
}

/* Schreibe in CSR */
void write_mstatus(uint64_t value) {
    asm volatile ("csrw mstatus, %0" :: "r"(value));
}

/* Wait For Interrupt */
void wfi(void) {
    asm volatile ("wfi");
}
```

---

## 📊 Debugging mit QEMU

### Debug-Output

```bash
# Starte QEMU mit Debug-Info
qemu-system-riscv64 -machine virt -bios none -kernel kernel.bin \
    -nographic \
    -d in_asm,cpu_reset \
    -D qemu.log
```

### Mit GDB debuggen

```bash
# Terminal 1: QEMU mit GDB-Server
qemu-system-riscv64 -machine virt -bios none -kernel kernel.elf \
    -nographic -S -gdb tcp::1234

# Terminal 2: GDB verbinden
riscv64-unknown-elf-gdb kernel.elf
(gdb) target remote :1234
(gdb) break kernel_main
(gdb) continue
```

### Nützliche GDB-Befehle

```
# Breakpoints
break kernel_main    # Bei Funktion
break *0x80000000    # Bei Adresse

# Ausführung
continue             # Weiterlaufen
step                 # Einzelschritt (in Funktion)
next                 # Einzelschritt (über Funktion)
stepi                # Eine Instruktion

# Anzeige
info registers       # Alle Register
print $a0            # Register a0
x/10i $pc            # Nächste 10 Instruktionen
x/10x $sp            # Stack anzeigen
```

---

## 📖 Weiterführende Ressourcen

### RISC-V
- **RISC-V OS in 1000 Lines** - [operating-system-in-1000-lines.vercel.app](https://operating-system-in-1000-lines.vercel.app/)
- **RISC-V Reader** - Patterson & Waterman (Free PDF)
- **RISC-V ISA Manual** - [riscv.org/specifications](https://riscv.org/specifications/)

### QEMU
- **QEMU Dokumentation** - [qemu.org/docs](https://www.qemu.org/docs/master/)
- **Virt Machine** - [qemu-project.gitlab.io/qemu/system/riscv/virt.html](https://qemu-project.gitlab.io/qemu/system/riscv/virt.html)

### Assembly
- **RISC-V Assembly Programmer's Manual** - GitHub riscv/riscv-asm-manual

---

## 🧠 Zusammenfassung

| Konzept | Was du gelernt hast |
|---------|---------------------|
| Bare Metal | Code ohne OS, direkt auf Hardware |
| RISC-V | Offene ISA, 32 Register, 3 Privilege Levels |
| QEMU virt | UART @ 0x10000000, RAM @ 0x80000000 |
| Boot-Prozess | boot.S → Stack init → BSS clear → kernel_main |
| UART | Einfachste I/O, THR schreiben = Zeichen senden |

**Key Takeaways für morgen:**
1. Der Boot-Code muss den Stack initialisieren
2. Linker Script definiert das Memory Layout
3. UART ist der einfachste Weg für Debug-Output
4. QEMU + GDB ermöglicht vollständiges Debugging

---

*Weiter zu den Übungen → `uebungen/uebung-02.md`*
