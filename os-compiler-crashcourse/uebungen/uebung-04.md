# Übung 4: Terminal OS

**Ziel**: Baue ein interaktives Terminal mit Shell.

---

## 🛠️ Setup

```bash
cd ~/os-dev/day2/mini-kernel
# Baue auf deinem bestehenden Kernel auf
```

---

## Aufgabe 1: UART Input (⏱️ 45 min)

Erweitere UART um Eingabe-Funktionalität.

**Datei**: Erweitere `uart.c`

```c
/* Füge zu deinem bestehenden uart.c hinzu */

/* Prüfe ob Daten zum Lesen bereit sind */
bool uart_has_data(void) {
    /* TODO: Prüfe LSR Bit 0 (Data Ready) */
    return false;
}

/* Lese ein Zeichen (polling, blockierend) */
char uart_getc_poll(void) {
    /* TODO: Warte auf Daten, dann lese von RBR */
    return '\0';
}

/* Einfacher Echo-Test */
void uart_echo_test(void) {
    uart_puts("Echo test. Type characters (Ctrl+C to exit):\n");
    
    while (1) {
        char c = uart_getc_poll();
        
        if (c == 0x03) {  /* Ctrl+C */
            uart_puts("\n^C\n");
            break;
        }
        
        /* Echo */
        if (c == '\r') {
            uart_puts("\r\n");
        } else {
            uart_putc(c);
        }
    }
}
```

**Test in kernel_main():**
```c
uart_echo_test();
```

---

## Aufgabe 2: Input Buffer (⏱️ 45 min)

Implementiere einen Ringbuffer für UART-Eingabe.

**Datei**: `input_buffer.c`

```c
#include <stdint.h>
#include <stdbool.h>

#define INPUT_BUFFER_SIZE 256

static char buffer[INPUT_BUFFER_SIZE];
static volatile int head = 0;  /* Schreibposition */
static volatile int tail = 0;  /* Leseposition */

/* Buffer initialisieren */
void input_buffer_init(void) {
    /* TODO */
}

/* Ist Buffer leer? */
bool input_buffer_empty(void) {
    /* TODO */
    return true;
}

/* Zeichen in Buffer einfügen */
bool input_buffer_push(char c) {
    /* TODO: Prüfe ob Platz, dann einfügen
     * Berechne next = (head + 1) % SIZE
     * Wenn next == tail: Buffer voll, return false
     * Sonst: buffer[head] = c; head = next
     */
    return false;
}

/* Zeichen aus Buffer lesen */
int input_buffer_pop(void) {
    /* TODO: Wenn leer, return -1
     * Sonst: c = buffer[tail]; tail = (tail + 1) % SIZE
     * Return c
     */
    return -1;
}

/* Wie viele Zeichen im Buffer? */
int input_buffer_count(void) {
    /* TODO */
    return 0;
}
```

---

## Aufgabe 3: UART Interrupt (⏱️ 60 min)

Nutze Interrupts statt Polling für UART.

### 3a: PLIC Setup

```c
/* plic.c - Platform Level Interrupt Controller */

#define PLIC_BASE 0x0C000000

/* PLIC Register */
#define PLIC_PRIORITY(n)  (*(volatile uint32_t *)(PLIC_BASE + (n) * 4))
#define PLIC_PENDING      (*(volatile uint32_t *)(PLIC_BASE + 0x1000))
#define PLIC_ENABLE(hart) (*(volatile uint32_t *)(PLIC_BASE + 0x2000 + (hart) * 0x100))
#define PLIC_THRESHOLD(hart) (*(volatile uint32_t *)(PLIC_BASE + 0x200000 + (hart) * 0x2000))
#define PLIC_CLAIM(hart)  (*(volatile uint32_t *)(PLIC_BASE + 0x200004 + (hart) * 0x2000))

/* QEMU virt: UART ist Interrupt 10 */
#define UART_IRQ 10

void plic_init(void) {
    /* TODO: Setze Priorität für UART IRQ */
    PLIC_PRIORITY(UART_IRQ) = 1;
    
    /* TODO: Aktiviere UART IRQ für Hart 0 */
    
    /* TODO: Setze Threshold auf 0 (alle Interrupts erlaubt) */
}

/* Hole nächsten Interrupt (Claim) */
uint32_t plic_claim(void) {
    /* TODO */
    return 0;
}

/* Signalisiere Interrupt-Ende (Complete) */
void plic_complete(uint32_t irq) {
    /* TODO */
}
```

### 3b: UART Interrupt Handler

```c
/* In uart.c erweitern */

void uart_enable_interrupts(void) {
    /* IER Bit 0: Enable Received Data Available Interrupt */
    UART_IER = 0x01;
}

/* Wird von trap_handler aufgerufen */
void uart_interrupt_handler(void) {
    /* TODO: Solange Daten verfügbar, lese und pushe in Buffer */
    while (uart_has_data()) {
        char c = UART_RBR;
        input_buffer_push(c);
    }
}
```

### 3c: Trap Handler erweitern

```c
/* In trap_handler, füge External Interrupt hinzu */
void trap_handler(void) {
    /* ... */
    
    if (is_interrupt) {
        switch (code) {
            case 7:  /* Timer */
                timer_tick();
                break;
            case 11: /* External */
                handle_external_interrupt();
                break;
        }
    }
}

void handle_external_interrupt(void) {
    uint32_t irq = plic_claim();
    
    if (irq == UART_IRQ) {
        uart_interrupt_handler();
    }
    
    if (irq > 0) {
        plic_complete(irq);
    }
}
```

---

## Aufgabe 4: Line Editor (⏱️ 60 min)

Implementiere einen Line Editor mit Backspace und Cursor.

**Datei**: `line.c`

```c
#include <stdint.h>
#include <stdbool.h>

extern void uart_putc(char c);
extern void uart_puts(const char *s);
extern int input_buffer_pop(void);

/* Non-blocking getc (return -1 wenn leer) */
int getc_nonblock(void) {
    return input_buffer_pop();
}

/* Blocking getc */
char getc_block(void) {
    int c;
    while ((c = getc_nonblock()) == -1) {
        asm volatile ("wfi");
    }
    return (char)c;
}

/* Lese eine Zeile mit Line Editing */
int readline(const char *prompt, char *buffer, int maxlen) {
    uart_puts(prompt);
    
    int len = 0;  /* Aktuelle Länge */
    
    while (1) {
        char c = getc_block();
        
        /* TODO: Behandle verschiedene Tasten:
         * 
         * Enter (\r oder \n):
         *   - Gib \r\n aus
         *   - Terminiere Buffer mit \0
         *   - Return len
         * 
         * Backspace (0x7F oder 0x08):
         *   - Wenn len > 0:
         *     - Gib "\b \b" aus (Cursor zurück, Leerzeichen, Cursor zurück)
         *     - Dekrementiere len
         * 
         * Ctrl+C (0x03):
         *   - Gib "^C\r\n" aus
         *   - Return -1
         * 
         * Druckbare Zeichen (32-126):
         *   - Wenn len < maxlen-1:
         *     - Speichere in Buffer
         *     - Inkrementiere len
         *     - Gib Zeichen aus (Echo)
         */
    }
    
    return len;
}
```

**Test:**
```c
void test_readline(void) {
    char line[128];
    
    while (1) {
        int len = readline("test> ", line, sizeof(line));
        
        if (len < 0) {
            uart_puts("Cancelled\n");
            continue;
        }
        
        uart_puts("You typed: '");
        uart_puts(line);
        uart_puts("'\n");
    }
}
```

---

## Aufgabe 5: Command Shell (⏱️ 90 min)

Baue eine einfache Shell mit eingebauten Befehlen.

**Datei**: `shell.c`

```c
#include <stdint.h>

/* Externe Funktionen */
extern void uart_puts(const char *s);
extern void print_int(int val);
extern void print_hex(uint64_t val);
extern int readline(const char *prompt, char *buffer, int maxlen);

/* String-Vergleich */
static int streq(const char *a, const char *b) {
    while (*a && *b) {
        if (*a++ != *b++) return 0;
    }
    return *a == *b;
}

/* Überspringe Whitespace */
static char *skip_space(char *s) {
    while (*s == ' ' || *s == '\t') s++;
    return s;
}

/* Tokenize: Finde erstes Wort */
static char *get_token(char *s, char **next) {
    s = skip_space(s);
    if (*s == '\0') {
        *next = s;
        return 0;
    }
    
    char *start = s;
    while (*s && *s != ' ' && *s != '\t') s++;
    
    if (*s) {
        *s++ = '\0';
    }
    *next = s;
    return start;
}

/* === Befehle === */

void cmd_help(void) {
    uart_puts("Available commands:\n");
    uart_puts("  help     - Show this help\n");
    uart_puts("  echo     - Echo arguments\n");
    uart_puts("  clear    - Clear screen\n");
    uart_puts("  mem      - Show memory info\n");
    uart_puts("  uptime   - Show uptime\n");
    uart_puts("  reboot   - Reboot system\n");
}

void cmd_echo(char *args) {
    /* TODO: Gib restliche Argumente aus */
}

void cmd_clear(void) {
    /* TODO: Sende ANSI Clear Screen */
}

void cmd_mem(void) {
    /* TODO: Rufe palloc_stats() auf */
}

void cmd_uptime(void) {
    /* TODO: Zeige get_ticks() */
}

void cmd_reboot(void) {
    uart_puts("Rebooting...\n");
    /* QEMU: Einfach neustarten mit wfi-loop oder Reset via SBI */
    while (1) { asm volatile ("wfi"); }
}

/* === Shell Main Loop === */

void shell_run(void) {
    char line[256];
    
    uart_puts("\n\x1b[36m");
    uart_puts("╔════════════════════════════════════╗\n");
    uart_puts("║       Mini OS Shell v0.1           ║\n");
    uart_puts("╚════════════════════════════════════╝\n");
    uart_puts("\x1b[0m\n");
    
    while (1) {
        int len = readline("\x1b[32mminios\x1b[0m> ", line, sizeof(line));
        
        if (len < 0) continue;  /* Cancelled */
        if (len == 0) continue; /* Leer */
        
        char *rest = line;
        char *cmd = get_token(rest, &rest);
        
        if (!cmd) continue;
        
        /* TODO: Dispatch zu richtigen Command Handler */
        if (streq(cmd, "help")) {
            cmd_help();
        } else if (streq(cmd, "echo")) {
            cmd_echo(rest);
        } else if (streq(cmd, "clear")) {
            cmd_clear();
        } else if (streq(cmd, "mem")) {
            cmd_mem();
        } else if (streq(cmd, "uptime")) {
            cmd_uptime();
        } else if (streq(cmd, "reboot")) {
            cmd_reboot();
        } else {
            uart_puts("Unknown command: ");
            uart_puts(cmd);
            uart_puts("\nType 'help' for available commands.\n");
        }
    }
}
```

**Test in kernel_main():**
```c
void kernel_main(void) {
    uart_puts("Mini Kernel v0.3\n");
    
    palloc_init();
    trap_init();
    timer_init();
    plic_init();
    uart_enable_interrupts();
    input_buffer_init();
    
    shell_run();
}
```

---

## Aufgabe 6: Weitere Befehle (⏱️ 45 min)

Füge diese nützlichen Befehle hinzu:

```c
/* hexdump <addr> <len> - Zeige Speicher */
void cmd_hexdump(char *args) {
    /* TODO: Parse addr und len, zeige Hex-Dump */
}

/* poke <addr> <value> - Schreibe Wert an Adresse */
void cmd_poke(char *args) {
    /* TODO: Vorsicht! */
}

/* peek <addr> - Lese Wert von Adresse */
void cmd_peek(char *args) {
    /* TODO */
}

/* color <on|off> - Toggle Farben */
void cmd_color(char *args) {
    /* TODO */
}
```

---

## 🎯 Bonusaufgaben

### Bonus A: Command History

Speichere die letzten 10 Befehle und navigiere mit Pfeiltasten.

### Bonus B: Tab Completion

Vervollständige Befehle automatisch bei Tab-Druck.

### Bonus C: ASCII Art

Füge ein lustiges Boot-Logo hinzu:
```
void show_logo(void) {
    uart_puts("\x1b[33m");
    uart_puts("    __  ____      _ ____  _____\n");
    uart_puts("   /  |/  (_)__  (_) __ \\/ ___/\n");
    /* usw. */
    uart_puts("\x1b[0m");
}
```

---

## ✅ Checkliste

- [ ] UART kann lesen (polling)
- [ ] Input Buffer mit Interrupts funktioniert
- [ ] Line Editing mit Backspace
- [ ] Shell läuft und erkennt Befehle
- [ ] Mindestens 5 Befehle implementiert

**Morgen**: Compiler-Grundlagen mit Ruby!

---

*Weiter zu Tag 5 → `vorlesungen/tag-05-compiler.md`*
