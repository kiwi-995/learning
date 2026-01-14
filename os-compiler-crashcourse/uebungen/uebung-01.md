# Übung 1: C Foundations

**Ziel**: Praktische C-Kenntnisse durch eigene Implementierung erwerben.

---

## 🛠️ Setup

Erstelle ein Arbeitsverzeichnis:

```bash
mkdir -p ~/os-dev/day1
cd ~/os-dev/day1
```

Kompilieren:
```bash
# Mit Debug-Symbolen und allen Warnungen
gcc -Wall -Wextra -g -o programm datei.c

# Ausführen
./programm
```

---

## Aufgabe 1: Hello C (⏱️ 15 min)

Schreibe dein erstes C-Programm.

**Datei**: `hello.c`

```c
// TODO: Include für printf

// TODO: main-Funktion die "Hello, Systems Programming!" ausgibt
// und 0 zurückgibt
```

**Test**:
```bash
gcc -Wall -o hello hello.c && ./hello
# Erwartete Ausgabe: Hello, Systems Programming!
```

---

## Aufgabe 2: Pointer Explorer (⏱️ 30 min)

Implementiere Funktionen die Pointer manipulieren.

**Datei**: `pointers.c`

```c
#include <stdio.h>
#include <stdint.h>

// Aufgabe 2a: Tausche zwei Integer-Werte über Pointer
void swap(int *a, int *b) {
    // TODO: Implementiere!
}

// Aufgabe 2b: Gib die Summe eines Arrays zurück
// Benutze Pointer-Arithmetik (NICHT array[i])!
int array_sum(int *arr, int length) {
    // TODO: Implementiere!
    return 0;
}

// Aufgabe 2c: Finde den maximalen Wert und seine Adresse
void find_max(int *arr, int length, int **max_addr) {
    // TODO: *max_addr soll auf das Maximum zeigen
}

int main() {
    // Test 2a
    int x = 5, y = 10;
    swap(&x, &y);
    printf("swap: x=%d, y=%d (erwartet: 10, 5)\n", x, y);
    
    // Test 2b
    int nums[] = {1, 2, 3, 4, 5};
    int sum = array_sum(nums, 5);
    printf("sum: %d (erwartet: 15)\n", sum);
    
    // Test 2c
    int *max_ptr;
    find_max(nums, 5, &max_ptr);
    printf("max: %d @ %p (erwartet: 5)\n", *max_ptr, (void*)max_ptr);
    
    return 0;
}
```

**Erwartete Ausgabe**:
```
swap: x=10, y=5 (erwartet: 10, 5)
sum: 15 (erwartet: 15)
max: 5 @ 0x... (erwartet: 5)
```

---

## Aufgabe 3: Memory Layout (⏱️ 30 min)

Untersuche das Speicherlayout eines Programms.

**Datei**: `memory_layout.c`

```c
#include <stdio.h>
#include <stdlib.h>

// Globale Variablen
int global_initialized = 42;       // → .data Segment
int global_uninitialized;          // → .bss Segment
const char *string_literal = "Hello";  // String im .rodata

void print_addresses() {
    // Lokale Variable (Stack)
    int local = 100;
    
    // Dynamisch allokiert (Heap)
    int *heap_var = malloc(sizeof(int));
    *heap_var = 200;
    
    // TODO: Gib die Adressen aller Variablen aus
    // Formatiere mit %p für Pointer
    // Ordne sie von niedrig nach hoch
    
    printf("=== Memory Layout ===\n");
    // printf("Code (.text):     %p (Adresse von print_addresses)\n", ...);
    // printf(".rodata:          %p (string_literal zeigt auf)\n", ...);
    // printf(".data:            %p (global_initialized)\n", ...);
    // printf(".bss:             %p (global_uninitialized)\n", ...);
    // printf("Heap:             %p (heap_var zeigt auf)\n", ...);
    // printf("Stack:            %p (local)\n", ...);
    
    free(heap_var);
}

int main() {
    print_addresses();
    
    // Bonus: Welches Segment liegt am niedrigsten?
    // Bonus: In welche Richtung wächst der Stack?
    
    return 0;
}
```

**Hinweis**: Die exakten Adressen variieren, aber die relative Ordnung sollte sein:
`.text` < `.rodata` < `.data` < `.bss` < `Heap` < ... < `Stack`

---

## Aufgabe 4: Struct als Hardware-Register (⏱️ 45 min)

Simuliere ein Hardware-Register mit Structs.

**Datei**: `fake_uart.c`

```c
#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>

// Simuliertes "UART" Hardware-Register
// In echter Hardware wäre das an einer festen Speicheradresse
struct UART {
    uint8_t tx_data;      // Zu sendendes Byte
    uint8_t rx_data;      // Empfangenes Byte
    uint8_t status;       // Bit 0: TX ready, Bit 1: RX available
    uint8_t control;      // Bit 0: TX enable, Bit 1: RX enable
};

// Globaler "Hardware" Buffer (simuliert)
static char output_buffer[256];
static int output_pos = 0;

// TODO: Implementiere diese Funktionen

// Initialisiere UART (setze control bits)
void uart_init(struct UART *uart) {
    // TODO
}

// Prüfe ob senden möglich (TX ready bit)
bool uart_tx_ready(struct UART *uart) {
    // TODO
    return false;
}

// Sende ein Zeichen
void uart_send(struct UART *uart, char c) {
    // TODO: Warte bis ready, dann schreibe in tx_data
    // Simuliere Output indem du output_buffer befüllst
}

// Sende einen String
void uart_print(struct UART *uart, const char *str) {
    // TODO
}

int main() {
    struct UART my_uart = {0};
    
    uart_init(&my_uart);
    uart_print(&my_uart, "Hello UART!\n");
    
    // Überprüfe Ergebnis
    printf("Output buffer: %s", output_buffer);
    printf("Erwartet:      Hello UART!\n");
    
    return 0;
}
```

---

## Aufgabe 5: Ringbuffer Implementierung (⏱️ 60 min)

Implementiere einen kompletten Ringbuffer - eine essentielle Datenstruktur für I/O.

**Datei**: `ringbuffer.c`

```c
#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <assert.h>

#define BUFFER_SIZE 8  // Klein für einfaches Testen

struct RingBuffer {
    uint8_t data[BUFFER_SIZE];
    // TODO: Welche zusätzlichen Felder brauchst du?
};

// Initialisiere Buffer
void ring_init(struct RingBuffer *rb) {
    // TODO
}

// Ist der Buffer leer?
bool ring_empty(struct RingBuffer *rb) {
    // TODO
    return true;
}

// Ist der Buffer voll?
bool ring_full(struct RingBuffer *rb) {
    // TODO
    return false;
}

// Wie viele Elemente sind drin?
uint32_t ring_count(struct RingBuffer *rb) {
    // TODO
    return 0;
}

// Füge Byte hinzu, return false wenn voll
bool ring_push(struct RingBuffer *rb, uint8_t byte) {
    // TODO
    return false;
}

// Entferne Byte, return false wenn leer
bool ring_pop(struct RingBuffer *rb, uint8_t *byte) {
    // TODO
    return false;
}

// Schaue nächstes Byte an ohne zu entfernen
bool ring_peek(struct RingBuffer *rb, uint8_t *byte) {
    // TODO
    return false;
}

// === TESTS ===
void test_basic() {
    struct RingBuffer rb;
    ring_init(&rb);
    
    assert(ring_empty(&rb));
    assert(!ring_full(&rb));
    assert(ring_count(&rb) == 0);
    
    printf("✓ Basic tests passed\n");
}

void test_push_pop() {
    struct RingBuffer rb;
    ring_init(&rb);
    uint8_t val;
    
    // Push und Pop
    assert(ring_push(&rb, 42));
    assert(ring_count(&rb) == 1);
    assert(ring_pop(&rb, &val));
    assert(val == 42);
    assert(ring_empty(&rb));
    
    // Mehrere Werte
    for (int i = 0; i < 5; i++) {
        assert(ring_push(&rb, i * 10));
    }
    assert(ring_count(&rb) == 5);
    
    for (int i = 0; i < 5; i++) {
        assert(ring_pop(&rb, &val));
        assert(val == i * 10);
    }
    
    printf("✓ Push/Pop tests passed\n");
}

void test_wraparound() {
    struct RingBuffer rb;
    ring_init(&rb);
    uint8_t val;
    
    // Fülle komplett
    for (int i = 0; i < BUFFER_SIZE; i++) {
        assert(ring_push(&rb, i));
    }
    assert(ring_full(&rb));
    assert(!ring_push(&rb, 99));  // Sollte fehlschlagen
    
    // Halb leeren
    for (int i = 0; i < BUFFER_SIZE / 2; i++) {
        assert(ring_pop(&rb, &val));
    }
    
    // Wieder auffüllen (wrap-around!)
    for (int i = 0; i < BUFFER_SIZE / 2; i++) {
        assert(ring_push(&rb, 100 + i));
    }
    assert(ring_full(&rb));
    
    printf("✓ Wraparound tests passed\n");
}

void test_peek() {
    struct RingBuffer rb;
    ring_init(&rb);
    uint8_t val;
    
    assert(!ring_peek(&rb, &val));  // Leer, sollte false sein
    
    ring_push(&rb, 42);
    assert(ring_peek(&rb, &val));
    assert(val == 42);
    assert(ring_count(&rb) == 1);  // Peek entfernt nicht!
    
    printf("✓ Peek tests passed\n");
}

int main() {
    test_basic();
    test_push_pop();
    test_wraparound();
    test_peek();
    
    printf("\n🎉 Alle Tests bestanden!\n");
    return 0;
}
```

---

## Aufgabe 6: Bitmanipulation (⏱️ 30 min)

Hardware-Register werden über Bits gesteuert. Übe Bitoperationen.

**Datei**: `bits.c`

```c
#include <stdio.h>
#include <stdint.h>
#include <assert.h>

// Setze Bit n (0-indexed)
uint32_t set_bit(uint32_t value, int n) {
    // TODO: Benutze | und <<
    return 0;
}

// Lösche Bit n
uint32_t clear_bit(uint32_t value, int n) {
    // TODO: Benutze & und ~
    return 0;
}

// Toggle Bit n
uint32_t toggle_bit(uint32_t value, int n) {
    // TODO: Benutze ^
    return 0;
}

// Prüfe ob Bit n gesetzt ist
int test_bit(uint32_t value, int n) {
    // TODO: Return 0 oder 1
    return 0;
}

// Extrahiere Bits von start bis end (inclusive)
uint32_t extract_bits(uint32_t value, int start, int end) {
    // TODO: z.B. extract_bits(0b11010110, 2, 5) = 0b0101 = 5
    return 0;
}

// Setze ein Feld von Bits
uint32_t set_field(uint32_t value, int start, int width, uint32_t field_value) {
    // TODO: Ersetze die Bits von start mit width Bits aus field_value
    return 0;
}

int main() {
    // Tests für set_bit
    assert(set_bit(0, 0) == 1);         // 0b0001
    assert(set_bit(0, 3) == 8);         // 0b1000
    assert(set_bit(1, 3) == 9);         // 0b1001
    
    // Tests für clear_bit
    assert(clear_bit(0xFF, 0) == 0xFE); // 0b11111110
    assert(clear_bit(0xFF, 7) == 0x7F); // 0b01111111
    
    // Tests für toggle_bit
    assert(toggle_bit(0, 2) == 4);      // 0b0100
    assert(toggle_bit(4, 2) == 0);      // 0b0000
    
    // Tests für test_bit
    assert(test_bit(0b1010, 1) == 1);
    assert(test_bit(0b1010, 0) == 0);
    
    // Tests für extract_bits
    assert(extract_bits(0b11010110, 1, 4) == 0b1011);  // Bits 1-4
    
    // Tests für set_field
    assert(set_field(0b11110000, 0, 4, 0b1010) == 0b11111010);
    
    printf("🎉 Alle Bit-Tests bestanden!\n");
    return 0;
}
```

---

## 🎯 Bonusaufgaben

### Bonus A: Einfacher String-Parser

Schreibe eine Funktion die einen String wie `"x=42,y=100"` parsed und die Werte extrahiert.

### Bonus B: Hexdump

Schreibe eine `hexdump(void *addr, int len)` Funktion die Speicher als Hex darstellt.

### Bonus C: Linked List

Implementiere eine einfach verkettete Liste mit `push_front`, `push_back`, `pop_front`.

---

## ✅ Zusammenfassung

Nach dieser Übung solltest du:

- [ ] C-Programme kompilieren und ausführen können
- [ ] Pointer verstehen und manipulieren können
- [ ] Das Speicherlayout kennen
- [ ] Structs für Hardware-Register nutzen können
- [ ] Einen funktionierenden Ringbuffer haben
- [ ] Bitoperationen beherrschen

**Morgen**: Diese Kenntnisse nutzen wir, um auf echter RISC-V Hardware zu booten!

---

*Weiter zu Tag 2 → `vorlesungen/tag-02-bare-metal.md`*
