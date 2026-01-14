# Tag 4: Terminal OS – Input, Output & Shell

**Ziel**: Baue ein interaktives Terminal mit Keyboard-Input und einem einfachen Command-Shell.

---

## 📚 Theorie

### 4.1 UART Vollständig

Bisher haben wir nur geschrieben. Jetzt auch lesen!

```
┌─────────────────────────────────────────────────────────────┐
│  UART Register (16550 kompatibel)                           │
├───────────┬─────────────────────────────────────────────────┤
│  Offset   │  Register                                       │
├───────────┼─────────────────────────────────────────────────┤
│  0x00     │  THR (Write) / RBR (Read) - Transmit/Receive    │
│  0x01     │  IER - Interrupt Enable Register                │
│  0x02     │  IIR (Read) / FCR (Write)                       │
│  0x03     │  LCR - Line Control Register                    │
│  0x04     │  MCR - Modem Control Register                   │
│  0x05     │  LSR - Line Status Register                     │
│  0x06     │  MSR - Modem Status Register                    │
│  0x07     │  SCR - Scratch Register                         │
└───────────┴─────────────────────────────────────────────────┘

LSR Bits:
  Bit 0: Data Ready (RBR hat Daten)
  Bit 1: Overrun Error
  Bit 5: THR Empty (bereit zum Senden)
  Bit 6: Transmitter Empty
```

### 4.2 Polling vs Interrupts

```
┌────────────────────────────────────────────────────────────┐
│  Polling                                                   │
├────────────────────────────────────────────────────────────┤
│  while (1) {                                               │
│      if (uart_has_data()) {                                │
│          char c = uart_read();                             │
│          handle(c);                                        │
│      }                                                     │
│      do_other_work();  // <-- Blockiert wenn Daten kommen! │
│  }                                                         │
│                                                            │
│  ✓ Einfach zu implementieren                               │
│  ✗ CPU ist immer beschäftigt                               │
│  ✗ Kann Daten verlieren                                    │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│  Interrupts                                                │
├────────────────────────────────────────────────────────────┤
│  void uart_interrupt_handler() {                           │
│      while (uart_has_data()) {                             │
│          buffer_push(uart_read());                         │
│      }                                                     │
│  }                                                         │
│                                                            │
│  main() {                                                  │
│      while (1) {                                           │
│          wfi();  // CPU schläft bis Interrupt              │
│          while (!buffer_empty()) {                         │
│              handle(buffer_pop());                         │
│          }                                                 │
│      }                                                     │
│  }                                                         │
│                                                            │
│  ✓ CPU kann schlafen (spart Energie)                       │
│  ✓ Reagiert sofort auf Eingabe                             │
│  ✗ Komplexer (Interrupt Setup, Buffering)                  │
└────────────────────────────────────────────────────────────┘
```

### 4.3 Line Editing

Ein gutes Terminal braucht Line Editing:

```
Tasten:
  Enter (0x0D)    → Zeile abschicken
  Backspace (0x7F oder 0x08) → Zeichen löschen
  Ctrl+C (0x03)   → Abbrechen
  Ctrl+D (0x04)   → EOF
  
Escape Sequences (Cursor bewegen):
  ESC [ A         → Cursor hoch (History)
  ESC [ B         → Cursor runter
  ESC [ C         → Cursor rechts
  ESC [ D         → Cursor links
```

### 4.4 ANSI Escape Codes

Terminal-Steuerung über Escape Sequences:

```
Cursor:
  \x1b[H        Home (0,0)
  \x1b[2J       Clear Screen
  \x1b[<n>A     Cursor n Zeilen hoch
  \x1b[<n>B     Cursor n Zeilen runter
  \x1b[<n>C     Cursor n Spalten rechts
  \x1b[<n>D     Cursor n Spalten links
  \x1b[<r>;<c>H Cursor zu Zeile r, Spalte c

Farben (Foreground):
  \x1b[30m Black    \x1b[31m Red
  \x1b[32m Green    \x1b[33m Yellow
  \x1b[34m Blue     \x1b[35m Magenta
  \x1b[36m Cyan     \x1b[37m White
  \x1b[0m  Reset

Background:
  \x1b[40m - \x1b[47m (wie Foreground, +10)
```

### 4.5 Shell Architektur

```
┌─────────────────────────────────────────────────────────────┐
│  User Input                                                 │
│      ↓                                                      │
│  ┌─────────────────┐                                        │
│  │  Line Buffer    │  ← Sammelt Eingabe bis Enter           │
│  └────────┬────────┘                                        │
│           ↓                                                 │
│  ┌─────────────────┐                                        │
│  │  Tokenizer      │  ← Zerlegt in Wörter                   │
│  └────────┬────────┘                                        │
│           ↓                                                 │
│  ┌─────────────────┐                                        │
│  │  Command Lookup │  ← Findet passenden Befehl             │
│  └────────┬────────┘                                        │
│           ↓                                                 │
│  ┌─────────────────┐                                        │
│  │  Execute        │  ← Führt Befehl aus                    │
│  └────────┬────────┘                                        │
│           ↓                                                 │
│  Output                                                     │
└─────────────────────────────────────────────────────────────┘
```

---

## 💻 Code-Beispiele

### UART mit Input

```c
/* uart.c - Vollständige UART Implementierung */

#include <stdint.h>
#include <stdbool.h>

#define UART_BASE 0x10000000

/* Register Offsets */
#define UART_RBR (*(volatile uint8_t *)(UART_BASE + 0x00))  /* Receive */
#define UART_THR (*(volatile uint8_t *)(UART_BASE + 0x00))  /* Transmit */
#define UART_IER (*(volatile uint8_t *)(UART_BASE + 0x01))  /* Interrupt Enable */
#define UART_FCR (*(volatile uint8_t *)(UART_BASE + 0x02))  /* FIFO Control */
#define UART_LCR (*(volatile uint8_t *)(UART_BASE + 0x03))  /* Line Control */
#define UART_LSR (*(volatile uint8_t *)(UART_BASE + 0x05))  /* Line Status */

/* LSR Bits */
#define LSR_RX_READY 0x01
#define LSR_TX_EMPTY 0x20

/* IER Bits */
#define IER_RX_ENABLE 0x01
#define IER_TX_ENABLE 0x02

/* Ringbuffer für Eingabe */
#define INPUT_BUFFER_SIZE 256
static char input_buffer[INPUT_BUFFER_SIZE];
static volatile int input_head = 0;
static volatile int input_tail = 0;

void uart_init(void) {
    /* FIFO aktivieren */
    UART_FCR = 0x07;  /* Enable FIFO, clear buffers */
    
    /* 8 bits, no parity, 1 stop bit */
    UART_LCR = 0x03;
    
    /* Interrupts aktivieren */
    UART_IER = IER_RX_ENABLE;
}

bool uart_has_data(void) {
    return (UART_LSR & LSR_RX_READY) != 0;
}

char uart_getc_poll(void) {
    while (!uart_has_data()) {}
    return UART_RBR;
}

void uart_putc(char c) {
    while ((UART_LSR & LSR_TX_EMPTY) == 0) {}
    UART_THR = c;
}

void uart_puts(const char *s) {
    while (*s) {
        if (*s == '\n') uart_putc('\r');
        uart_putc(*s++);
    }
}

/* Interrupt Handler */
void uart_interrupt(void) {
    while (uart_has_data()) {
        char c = UART_RBR;
        int next = (input_head + 1) % INPUT_BUFFER_SIZE;
        if (next != input_tail) {
            input_buffer[input_head] = c;
            input_head = next;
        }
    }
}

/* Lese aus Buffer (non-blocking) */
int uart_getc(void) {
    if (input_head == input_tail) {
        return -1;  /* Kein Zeichen verfügbar */
    }
    char c = input_buffer[input_tail];
    input_tail = (input_tail + 1) % INPUT_BUFFER_SIZE;
    return c;
}

/* Lese aus Buffer (blocking) */
char uart_getc_blocking(void) {
    int c;
    while ((c = uart_getc()) == -1) {
        asm volatile ("wfi");
    }
    return (char)c;
}
```

### Line Editor

```c
/* line.c - Line Editing */

#include <stdint.h>
#include <stdbool.h>

#define MAX_LINE_LENGTH 256

extern void uart_putc(char c);
extern void uart_puts(const char *s);
extern char uart_getc_blocking(void);

/* Lese eine Zeile mit Editing */
int readline(const char *prompt, char *buffer, int max_len) {
    uart_puts(prompt);
    
    int pos = 0;
    int len = 0;
    
    while (1) {
        char c = uart_getc_blocking();
        
        if (c == '\r' || c == '\n') {
            /* Enter: Zeile fertig */
            uart_puts("\r\n");
            buffer[len] = '\0';
            return len;
            
        } else if (c == 0x7F || c == 0x08) {
            /* Backspace */
            if (pos > 0) {
                /* Cursor zurück, Zeichen löschen */
                uart_puts("\b \b");
                pos--;
                len--;
                /* Zeichen aus Buffer entfernen */
                for (int i = pos; i < len; i++) {
                    buffer[i] = buffer[i + 1];
                }
            }
            
        } else if (c == 0x03) {
            /* Ctrl+C: Abbrechen */
            uart_puts("^C\r\n");
            buffer[0] = '\0';
            return -1;
            
        } else if (c == 0x04) {
            /* Ctrl+D: EOF */
            if (len == 0) {
                return -2;  /* EOF */
            }
            
        } else if (c == 0x1B) {
            /* Escape Sequence */
            char seq1 = uart_getc_blocking();
            char seq2 = uart_getc_blocking();
            
            if (seq1 == '[') {
                switch (seq2) {
                    case 'A':  /* Up Arrow */
                        /* TODO: Command History */
                        break;
                    case 'B':  /* Down Arrow */
                        break;
                    case 'C':  /* Right Arrow */
                        if (pos < len) {
                            uart_puts("\x1b[C");
                            pos++;
                        }
                        break;
                    case 'D':  /* Left Arrow */
                        if (pos > 0) {
                            uart_puts("\x1b[D");
                            pos--;
                        }
                        break;
                }
            }
            
        } else if (c >= 32 && c < 127) {
            /* Druckbares Zeichen */
            if (len < max_len - 1) {
                /* Zeichen einfügen */
                for (int i = len; i > pos; i--) {
                    buffer[i] = buffer[i - 1];
                }
                buffer[pos] = c;
                pos++;
                len++;
                
                uart_putc(c);
            }
        }
    }
}
```

### Simple Shell

```c
/* shell.c - Command Shell */

#include <stdint.h>
#include <stdbool.h>

extern void uart_puts(const char *s);
extern void print_int(int val);
extern void print_hex(uint64_t val);
extern int readline(const char *prompt, char *buffer, int max_len);
extern uint64_t strlen(const char *s);
extern int strcmp(const char *s1, const char *s2);

/* Command Handler Typ */
typedef void (*cmd_func_t)(int argc, char **argv);

/* Command Eintrag */
struct command {
    const char *name;
    const char *help;
    cmd_func_t handler;
};

/* Prototypen */
void cmd_help(int argc, char **argv);
void cmd_echo(int argc, char **argv);
void cmd_clear(int argc, char **argv);
void cmd_mem(int argc, char **argv);
void cmd_uptime(int argc, char **argv);

/* Command Table */
static struct command commands[] = {
    {"help",   "Show available commands", cmd_help},
    {"echo",   "Echo arguments",          cmd_echo},
    {"clear",  "Clear screen",            cmd_clear},
    {"mem",    "Show memory info",        cmd_mem},
    {"uptime", "Show system uptime",      cmd_uptime},
    {0, 0, 0}  /* Terminator */
};

/* Tokenize: Zerlege String in Argumente */
int tokenize(char *line, char **argv, int max_args) {
    int argc = 0;
    char *p = line;
    
    while (*p && argc < max_args - 1) {
        /* Überspringe Whitespace */
        while (*p == ' ' || *p == '\t') p++;
        if (*p == '\0') break;
        
        /* Argument Start */
        argv[argc++] = p;
        
        /* Finde Ende */
        while (*p && *p != ' ' && *p != '\t') p++;
        if (*p) {
            *p++ = '\0';
        }
    }
    argv[argc] = 0;
    return argc;
}

/* Führe Command aus */
void execute(int argc, char **argv) {
    if (argc == 0) return;
    
    for (int i = 0; commands[i].name; i++) {
        if (strcmp(argv[0], commands[i].name) == 0) {
            commands[i].handler(argc, argv);
            return;
        }
    }
    
    uart_puts("Unknown command: ");
    uart_puts(argv[0]);
    uart_puts("\nType 'help' for available commands.\n");
}

/* Command Handlers */
void cmd_help(int argc, char **argv) {
    (void)argc; (void)argv;
    uart_puts("Available commands:\n");
    for (int i = 0; commands[i].name; i++) {
        uart_puts("  ");
        uart_puts(commands[i].name);
        uart_puts(" - ");
        uart_puts(commands[i].help);
        uart_puts("\n");
    }
}

void cmd_echo(int argc, char **argv) {
    for (int i = 1; i < argc; i++) {
        uart_puts(argv[i]);
        if (i < argc - 1) uart_puts(" ");
    }
    uart_puts("\n");
}

void cmd_clear(int argc, char **argv) {
    (void)argc; (void)argv;
    uart_puts("\x1b[2J\x1b[H");  /* Clear screen, cursor home */
}

void cmd_mem(int argc, char **argv) {
    (void)argc; (void)argv;
    extern void palloc_stats(void);
    palloc_stats();
}

void cmd_uptime(int argc, char **argv) {
    (void)argc; (void)argv;
    extern uint64_t get_ticks(void);
    uint64_t ticks = get_ticks();
    uart_puts("Uptime: ");
    print_int(ticks);
    uart_puts(" seconds\n");
}

/* Shell Main Loop */
void shell_run(void) {
    char line[256];
    char *argv[16];
    
    uart_puts("\n");
    uart_puts("\x1b[32m╔═══════════════════════════════════════╗\x1b[0m\n");
    uart_puts("\x1b[32m║     Mini OS Shell v0.1                ║\x1b[0m\n");
    uart_puts("\x1b[32m╚═══════════════════════════════════════╝\x1b[0m\n");
    uart_puts("\nType 'help' for available commands.\n\n");
    
    while (1) {
        int len = readline("\x1b[34mminios\x1b[0m$ ", line, sizeof(line));
        
        if (len == -2) {
            /* Ctrl+D: Exit */
            uart_puts("Goodbye!\n");
            return;
        }
        
        if (len > 0) {
            int argc = tokenize(line, argv, 16);
            execute(argc, argv);
        }
    }
}
```

---

## 📖 Weiterführende Ressourcen

### Terminal/TTY
- **The TTY Demystified** - [linusakesson.net/programming/tty](https://www.linusakesson.net/programming/tty/)
- **VT100 User Guide** - Für alle Escape Codes

### Shell Implementation
- **POSIX Shell** - [pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html)
- **Build Your Own Shell** - [github.com/tokenrove/build-your-own-shell](https://github.com/tokenrove/build-your-own-shell)

---

## 🧠 Zusammenfassung

| Konzept | Was du gelernt hast |
|---------|---------------------|
| UART Input | LSR.RX_READY prüfen, RBR lesen |
| Buffering | Ringbuffer für Input-Zeichen |
| Line Editing | Backspace, Cursor-Bewegung |
| ANSI Codes | Farben, Cursor-Steuerung |
| Shell | Tokenize → Lookup → Execute |

---

*Weiter zu den Übungen → `uebungen/uebung-04.md`*
