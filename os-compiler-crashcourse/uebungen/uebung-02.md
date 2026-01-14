# Übung 2: Bare Metal RISC-V

**Ziel**: Boot deinen ersten Kernel in QEMU und gib Text über UART aus.

---

## 🛠️ Setup

### Vorbereitung

```bash
# Erstelle Arbeitsverzeichnis
mkdir -p ~/os-dev/day2/mini-kernel
cd ~/os-dev/day2/mini-kernel

# Prüfe ob Toolchain verfügbar
riscv64-unknown-elf-gcc --version
qemu-system-riscv64 --version
```

Falls die Toolchain fehlt, siehe Vorlesung für Installation.

---

## Aufgabe 1: Hello Kernel (⏱️ 90 min)

Schreibe deinen ersten bootbaren Kernel.

### 1a: Erstelle den Linker Script

**Datei**: `kernel.ld`

```ld
/* TODO: Vervollständige den Linker Script */

OUTPUT_ARCH(riscv)
ENTRY(_start)

/* Definiere Speicherbereiche */
MEMORY {
    /* RAM beginnt bei 0x80000000 mit 128MB */
    /* TODO: RAM (rwx) : ORIGIN = ?, LENGTH = ? */
}

SECTIONS {
    . = 0x80000000;

    /* Text-Section: Code */
    .text : {
        /* TODO: Boot-Code muss zuerst kommen! */
        *(.text .text.*)
    } > RAM

    /* Read-only Data */
    .rodata : {
        *(.rodata .rodata.*)
    } > RAM

    /* Initialisierte Daten */
    .data : {
        *(.data .data.*)
    } > RAM

    /* Uninitialisierte Daten (BSS) */
    .bss : {
        /* TODO: Exportiere __bss_start Symbol */
        *(.bss .bss.*)
        *(COMMON)
        /* TODO: Exportiere __bss_end Symbol */
    } > RAM

    /* Stack */
    . = ALIGN(16);
    . = . + 0x10000;  /* 64KB Stack */
    /* TODO: Exportiere __stack_top Symbol */
}
```

**Hinweis**: Symbole exportiert man mit `symbolname = .;`

### 1b: Schreibe den Boot-Code

**Datei**: `boot.S`

```asm
# TODO: Vervollständige den Boot-Code

.section .text.boot
.global _start

_start:
    # 1. Setze den Stack Pointer auf __stack_top
    # TODO: la sp, __stack_top
    
    # 2. Lösche das BSS Segment (alle Bytes auf 0)
    # TODO: Load __bss_start in t0
    # TODO: Load __bss_end in t1
clear_bss:
    # TODO: Wenn t0 == t1, springe zu done_bss
    # TODO: Schreibe 0 an Adresse in t0 (sd zero, 0(t0))
    # TODO: Erhöhe t0 um 8
    # TODO: Springe zurück zu clear_bss
done_bss:
    
    # 3. Rufe kernel_main auf
    # TODO: call kernel_main
    
    # 4. Falls Return, halte an
halt:
    wfi
    j halt
```

### 1c: Implementiere den Kernel

**Datei**: `kernel.c`

```c
#include <stdint.h>

/* UART Konfiguration */
#define UART_BASE 0x10000000

/* TODO: Definiere UART Register als volatile Pointer */
/* Tipp: #define UART_THR (*(volatile uint8_t *)(UART_BASE + 0x00)) */

/* TODO: Definiere LSR Register (Offset 0x05) */

/* TODO: Definiere LSR_TX_EMPTY Bit (0x20) */

/* Sende ein Zeichen über UART */
void uart_putc(char c) {
    /* TODO: Warte bis THR leer ist (LSR & LSR_TX_EMPTY != 0) */
    /* TODO: Schreibe Zeichen in THR */
}

/* Sende einen String */
void uart_puts(const char *s) {
    /* TODO: Iteriere über String, sende jedes Zeichen */
    /* TODO: Bei '\n' auch '\r' senden (Carriage Return) */
}

/* Kernel Entry Point */
void kernel_main(void) {
    uart_puts("=========================\n");
    uart_puts("  Mini Kernel v0.1\n");
    uart_puts("  Running on RISC-V!\n");
    uart_puts("=========================\n");
    
    /* Endlosschleife */
    while (1) {}
}
```

### 1d: Erstelle das Makefile

**Datei**: `Makefile`

```makefile
# Cross-Compiler Prefix
CROSS = riscv64-unknown-elf-

CC = $(CROSS)gcc
AS = $(CROSS)as
LD = $(CROSS)ld
OBJCOPY = $(CROSS)objcopy

# Compiler Flags
# -ffreestanding: Kein Standard-Library Startup
# -nostdlib: Keine Standard-Library linken
# -mcmodel=medany: Erlaubt beliebige Adressen
CFLAGS = -Wall -Wextra -O2 -ffreestanding -nostdlib -mcmodel=medany

# Assembler Flags
ASFLAGS = -march=rv64imac -mabi=lp64

# Build Targets
all: kernel.bin

boot.o: boot.S
	$(AS) $(ASFLAGS) -o $@ $<

kernel.o: kernel.c
	$(CC) $(CFLAGS) -c -o $@ $<

kernel.elf: boot.o kernel.o kernel.ld
	$(LD) -T kernel.ld -o $@ boot.o kernel.o

kernel.bin: kernel.elf
	$(OBJCOPY) -O binary $< $@

# QEMU starten
run: kernel.bin
	qemu-system-riscv64 \
		-machine virt \
		-bios none \
		-kernel kernel.bin \
		-nographic

# Mit Debug
debug: kernel.elf
	qemu-system-riscv64 \
		-machine virt \
		-bios none \
		-kernel kernel.elf \
		-nographic \
		-S -gdb tcp::1234

clean:
	rm -f *.o *.elf *.bin

.PHONY: all run debug clean
```

### Test

```bash
make clean && make run
```

**Erwartete Ausgabe**:
```
=========================
  Mini Kernel v0.1
  Running on RISC-V!
=========================
```

Drücke `Ctrl+A` dann `X` um QEMU zu beenden.

---

## Aufgabe 2: Printf-Grundlagen (⏱️ 60 min)

Erweitere deinen Kernel um formatierte Ausgabe.

**Datei**: `printf.c`

```c
#include <stdint.h>

/* uart_putc und uart_puts aus kernel.c verwenden */
extern void uart_putc(char c);
extern void uart_puts(const char *s);

/* Konvertiere Integer zu String (Dezimal) */
void print_int(int value) {
    /* TODO: Handle negative numbers */
    /* TODO: Konvertiere digit by digit */
    /* TODO: Beachte: Ziffern kommen rückwärts! */
    /* Tipp: Speichere Ziffern in Buffer, dann rückwärts ausgeben */
    
    char buffer[20];
    int pos = 0;
    int negative = 0;
    
    if (value < 0) {
        negative = 1;
        value = -value;
    }
    
    if (value == 0) {
        uart_putc('0');
        return;
    }
    
    /* TODO: Ziffern extrahieren */
    
    /* TODO: Rückwärts ausgeben */
}

/* Konvertiere Integer zu Hex-String */
void print_hex(uint64_t value) {
    /* TODO: Konvertiere zu Hexadezimal */
    /* Format: 0x... mit führenden Nullen für 64-bit */
    
    const char *hex_digits = "0123456789abcdef";
    
    uart_puts("0x");
    
    /* TODO: 16 Nibbles ausgeben (64 bit = 16 hex digits) */
}

/* Einfaches printf (nur %d, %x, %s, %c) */
void kprintf(const char *fmt, ...) {
    /* TODO: Diese Funktion ist eine Bonus-Aufgabe! */
    /* Tipp: Benutze __builtin_va_list für variable Argumente */
}

/* Test-Funktion */
void test_printf(void) {
    uart_puts("Testing print functions:\n");
    
    uart_puts("print_int(123): ");
    print_int(123);
    uart_puts("\n");
    
    uart_puts("print_int(-456): ");
    print_int(-456);
    uart_puts("\n");
    
    uart_puts("print_hex(0xDEADBEEF): ");
    print_hex(0xDEADBEEF);
    uart_puts("\n");
    
    uart_puts("print_hex(0x80000000): ");
    print_hex(0x80000000);
    uart_puts("\n");
}
```

**Test**: Rufe `test_printf()` in `kernel_main()` auf.

---

## Aufgabe 3: Machine Info (⏱️ 45 min)

Lese RISC-V System-Register aus.

**Datei**: `machine.c`

```c
#include <stdint.h>

extern void uart_puts(const char *s);
extern void print_int(int value);
extern void print_hex(uint64_t value);

/* Lese Machine ISA Register (zeigt unterstützte Extensions) */
uint64_t read_misa(void) {
    uint64_t value;
    /* TODO: Inline Assembly um misa CSR zu lesen */
    /* asm volatile ("csrr %0, misa" : "=r"(value)); */
    return value;
}

/* Lese Machine Vendor ID */
uint64_t read_mvendorid(void) {
    uint64_t value;
    /* TODO: Lese mvendorid CSR */
    return value;
}

/* Lese Machine Architecture ID */
uint64_t read_marchid(void) {
    uint64_t value;
    /* TODO: Lese marchid CSR */
    return value;
}

/* Lese Hart ID (Hardware Thread ID) */
uint64_t read_mhartid(void) {
    uint64_t value;
    /* TODO: Lese mhartid CSR */
    return value;
}

/* Dekodiere MISA Extensions */
void decode_misa(uint64_t misa) {
    uart_puts("Extensions: ");
    
    /* Bits 0-25 repräsentieren Extensions A-Z */
    const char *ext_names = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    
    for (int i = 0; i < 26; i++) {
        if (misa & (1UL << i)) {
            uart_putc(ext_names[i]);
            uart_putc(' ');
        }
    }
    uart_puts("\n");
    
    /* MXL (Machine XLEN) in Bits 62-63 */
    int mxl = (misa >> 62) & 3;
    uart_puts("XLEN: ");
    if (mxl == 1) uart_puts("32-bit");
    else if (mxl == 2) uart_puts("64-bit");
    else if (mxl == 3) uart_puts("128-bit");
    uart_puts("\n");
}

/* Zeige alle Machine-Infos */
void print_machine_info(void) {
    uart_puts("\n=== Machine Information ===\n");
    
    uart_puts("Hart ID: ");
    print_int((int)read_mhartid());
    uart_puts("\n");
    
    uart_puts("Vendor ID: ");
    print_hex(read_mvendorid());
    uart_puts("\n");
    
    uart_puts("Architecture ID: ");
    print_hex(read_marchid());
    uart_puts("\n");
    
    uint64_t misa = read_misa();
    uart_puts("MISA: ");
    print_hex(misa);
    uart_puts("\n");
    
    decode_misa(misa);
    
    uart_puts("===========================\n");
}
```

**Erwartete Ausgabe** (ähnlich):
```
=== Machine Information ===
Hart ID: 0
Vendor ID: 0x0
Architecture ID: 0x0
MISA: 0x800000000014112d
Extensions: A C D F I M S U 
XLEN: 64-bit
===========================
```

---

## Aufgabe 4: Memory Map Explorer (⏱️ 30 min)

Erkunde den QEMU Memory Map durch Lesen von bekannten Adressen.

**Datei**: `memmap.c`

```c
#include <stdint.h>

extern void uart_puts(const char *s);
extern void print_hex(uint64_t value);

/* Versuche von einer Adresse zu lesen */
uint32_t safe_read32(uint64_t addr) {
    return *(volatile uint32_t *)addr;
}

/* Zeige interessante Speicherbereiche */
void explore_memory(void) {
    uart_puts("\n=== Memory Map Exploration ===\n");
    
    /* UART @ 0x10000000 */
    uart_puts("UART base (0x10000000): ");
    print_hex(safe_read32(0x10000000));
    uart_puts("\n");
    
    /* PLIC @ 0x0C000000 */
    uart_puts("PLIC base (0x0C000000): ");
    print_hex(safe_read32(0x0C000000));
    uart_puts("\n");
    
    /* RAM @ 0x80000000 (unser Code!) */
    uart_puts("RAM base (0x80000000): ");
    print_hex(safe_read32(0x80000000));
    uart_puts(" <- Our boot code!\n");
    
    /* TODO: Lies weitere interessante Adressen */
    
    uart_puts("==============================\n");
}
```

---

## Aufgabe 5: Timer Delay (⏱️ 45 min)

Implementiere eine simple Delay-Funktion.

**Datei**: `timer.c`

```c
#include <stdint.h>

/* CLINT (Core Local Interrupter) Memory Map */
#define CLINT_BASE 0x02000000
#define CLINT_MTIME (*(volatile uint64_t *)(CLINT_BASE + 0xBFF8))

/* QEMU virt verwendet 10MHz Timer */
#define TIMER_FREQ 10000000

/* Lese aktuelle Timer-Wert */
uint64_t get_time(void) {
    return CLINT_MTIME;
}

/* Warte für n Millisekunden */
void delay_ms(uint32_t ms) {
    /* TODO: Berechne Ziel-Zeit */
    /* TODO: Warte bis Zeit erreicht */
    
    uint64_t start = get_time();
    uint64_t ticks = (uint64_t)ms * TIMER_FREQ / 1000;
    
    /* TODO: Busy-wait loop */
}

/* Test-Funktion */
void test_timer(void) {
    extern void uart_puts(const char *s);
    extern void print_int(int value);
    
    uart_puts("Timer test: Counting 1-5 with 1s delay...\n");
    
    for (int i = 1; i <= 5; i++) {
        print_int(i);
        uart_puts("...\n");
        delay_ms(1000);  /* 1 Sekunde warten */
    }
    
    uart_puts("Done!\n");
}
```

---

## 🎯 Bonusaufgaben

### Bonus A: Memory Hexdump

Implementiere eine `hexdump(addr, len)` Funktion die Speicher formatiert ausgibt:
```
80000000: 97 02 00 00 93 82 02 20  13 05 00 00 17 15 00 00
80000010: ...
```

### Bonus B: UART Echo

Implementiere `uart_getc()` um Zeichen zu empfangen und ein Echo-Programm.

### Bonus C: Bunte Ausgabe

UART unterstützt ANSI Escape Codes. Implementiere farbige Ausgabe:
```c
uart_puts("\x1b[32mGrüner Text\x1b[0m");
```

---

## ✅ Checkliste

Nach dieser Übung solltest du:

- [ ] Einen bootbaren RISC-V Kernel haben
- [ ] UART-Ausgabe funktioniert
- [ ] Integer und Hex-Zahlen ausgeben können
- [ ] System-Register lesen können
- [ ] Timer/Delay funktioniert

**Morgen**: Memory Management und Interrupts!

---

*Weiter zu Tag 3 → `vorlesungen/tag-03-kernel.md`*
