# Vorlesung 2: Datentypen & Variablen

## Lernziele

Nach dieser Vorlesung kannst du:
- Alle grundlegenden Datentypen in Scala nutzen
- String-Interpolation anwenden
- Mit Tupeln arbeiten
- Die Scala-Typhierarchie verstehen
- Typkonvertierungen durchführen

---

## 2.1 Die Scala Typhierarchie

Scala hat ein einheitliches Typsystem – **alles ist ein Objekt**:

```
                    ┌─────────┐
                    │   Any   │  ← Supertyp aller Typen
                    └────┬────┘
           ┌─────────────┴─────────────┐
           ▼                           ▼
    ┌──────────────┐            ┌──────────────┐
    │   AnyVal     │            │   AnyRef     │
    │  (Werttypen) │            │ (Referenzen) │
    └──────┬───────┘            └──────┬───────┘
           │                           │
    ┌──────┴──────┐             ┌──────┴──────┐
    │ Int, Double │             │ String, List│
    │ Boolean ... │             │ eigene Kl.  │
    └─────────────┘             └──────┬──────┘
                                       │
                                ┌──────┴──────┐
                                │    Null     │
                                └──────┬──────┘
                                       │
                                ┌──────┴──────┐
                                │   Nothing   │ ← Subtyp aller Typen
                                └─────────────┘
```

### Wichtige Typen

| Typ | Beschreibung | Beispiel |
|-----|--------------|----------|
| `Any` | Supertyp aller Typen | – |
| `AnyVal` | Werttypen (primitiv) | Int, Double |
| `AnyRef` | Referenztypen (Objekte) | String, List |
| `Unit` | Kein Rückgabewert | `()` |
| `Null` | null-Referenz (vermeiden!) | `null` |
| `Nothing` | Kein Wert (z.B. Exception) | – |

---

## 2.2 Numerische Typen

### Ganzzahlen

```scala
val byte: Byte = 127           // 8 Bit:  -128 bis 127
val short: Short = 32767       // 16 Bit: -32.768 bis 32.767
val int: Int = 2147483647      // 32 Bit: ca. ±2 Milliarden
val long: Long = 9223372036854775807L  // 64 Bit (L am Ende!)
```

### Fließkommazahlen

```scala
val float: Float = 3.14f       // 32 Bit (f am Ende!)
val double: Double = 3.14159   // 64 Bit (Standard für Dezimalzahlen)
```

### Rechenoperationen

```scala
val summe = 10 + 5      // 15
val differenz = 10 - 5  // 5
val produkt = 10 * 5    // 50
val quotient = 10 / 3   // 3 (Integer-Division!)
val rest = 10 % 3       // 1 (Modulo)

// Vorsicht bei Integer-Division:
val falsch = 10 / 3     // = 3, nicht 3.33!
val richtig = 10.0 / 3  // = 3.333...
```

### Nützliche Methoden

```scala
val zahl = -42

zahl.abs          // 42 (Absolutwert)
zahl.sign         // -1 (Vorzeichen)
10.max(5)         // 10
10.min(5)         // 5

// Math-Funktionen
math.sqrt(16)     // 4.0
math.pow(2, 10)   // 1024.0
math.random()     // Zufallszahl zwischen 0.0 und 1.0
```

---

## 2.3 Boolean

```scala
val wahr: Boolean = true
val falsch: Boolean = false
```

### Logische Operatoren

```scala
val a = true
val b = false

// UND (beide müssen wahr sein)
a && b          // false

// ODER (mindestens einer wahr)
a || b          // true

// NICHT (Negation)
!a              // false

// XOR (genau einer wahr)
a ^ b           // true
```

### Vergleichsoperatoren

```scala
val x = 5
val y = 10

x == y          // false (Gleichheit)
x != y          // true  (Ungleichheit)
x < y           // true  (kleiner)
x <= y          // true  (kleiner oder gleich)
x > y           // false (größer)
x >= y          // false (größer oder gleich)
```

---

## 2.4 Char und String

### Char (einzelnes Zeichen)

```scala
val buchstabe: Char = 'A'
val ziffer: Char = '7'
val sonderzeichen: Char = '\n'  // Newline

// Char ist numerisch (Unicode)
buchstabe.toInt    // 65
('A' + 1).toChar   // 'B'
```

### String

```scala
val text = "Hallo Welt"
val leer = ""
val mehrzeilig = """
  Das ist ein
  mehrzeiliger String.
"""
```

### Wichtige String-Operationen

```scala
val s = "Scala ist toll"

s.length            // 14
s.toUpperCase       // "SCALA IST TOLL"
s.toLowerCase       // "scala ist toll"
s.charAt(0)         // 'S'
s.substring(0, 5)   // "Scala"
s.split(" ")        // Array("Scala", "ist", "toll")
s.contains("ist")   // true
s.startsWith("Sc")  // true
s.endsWith("ll")    // true
s.replace("toll", "super")  // "Scala ist super"
s.trim              // Entfernt Leerzeichen am Anfang/Ende
```

### String-Verkettung

```scala
val vorname = "Max"
val nachname = "Mustermann"

// Mit + Operator
val name1 = vorname + " " + nachname

// Mit concat
val name2 = vorname.concat(" ").concat(nachname)
```

---

## 2.5 String-Interpolation ⭐

Scala bietet elegante String-Interpolation – **sehr wichtig**!

### s-Interpolation (Standard)

```scala
val name = "Alice"
val alter = 25

// Variablen einbetten mit $
val nachricht = s"Hallo, ich bin $name und $alter Jahre alt."
// → "Hallo, ich bin Alice und 25 Jahre alt."

// Ausdrücke mit ${...}
val info = s"In 5 Jahren bin ich ${alter + 5} Jahre alt."
// → "In 5 Jahren bin ich 30 Jahre alt."
```

### f-Interpolation (Formatierung)

```scala
val pi = 3.14159265
val preis = 19.99

// Formatierung wie printf in C
val text1 = f"Pi ist ungefähr $pi%.2f"
// → "Pi ist ungefähr 3.14"

val text2 = f"Der Preis beträgt $preis%8.2f Euro"
// → "Der Preis beträgt    19.99 Euro"

val zahl = 42
val text3 = f"Zahl: $zahl%05d"  // Mit führenden Nullen
// → "Zahl: 00042"
```

**Formatspezifikationen:**
| Format | Bedeutung | Beispiel |
|--------|-----------|----------|
| `%d` | Integer | `f"$x%d"` |
| `%f` | Float/Double | `f"$x%f"` |
| `%.2f` | 2 Dezimalstellen | `f"$x%.2f"` |
| `%8d` | Mindestbreite 8 | `f"$x%8d"` |
| `%08d` | Mit führenden Nullen | `f"$x%08d"` |
| `%s` | String | `f"$x%s"` |

### raw-Interpolation (Escape-Zeichen ignorieren)

```scala
val pfad = raw"C:\users\benni\dokumente"
// → "C:\users\benni\dokumente" 
// (Backslashes bleiben erhalten)
```

---

## 2.6 Tupel

Tupel gruppieren mehrere Werte **verschiedener Typen**:

### Tupel erstellen

```scala
// Explizite Syntax
val person = ("Alice", 25, true)

// Mit Pfeil-Syntax für Paare
val paar = "name" -> "Alice"  // entspricht ("name", "Alice")
```

### Auf Elemente zugreifen

```scala
val person = ("Alice", 25, true)

// Mit _1, _2, _3, ... (1-basiert!)
person._1    // "Alice"
person._2    // 25
person._3    // true

// Mit Dekonstruktion (Pattern Matching)
val (name, alter, aktiv) = person
println(name)   // "Alice"
println(alter)  // 25
```

### Tupel-Typen

```scala
val t2: (String, Int) = ("Test", 42)
val t3: (Int, Int, Int) = (1, 2, 3)

// Scala unterstützt Tupel bis Größe 22
```

### Wann Tupel verwenden?

```
┌─────────────────────────────────────────────────────────────┐
│                    Tupel vs Case Class                      │
├──────────────────────┬──────────────────────────────────────┤
│ Tupel                │ Case Class                           │
├──────────────────────┼──────────────────────────────────────┤
│ Schnell, ad-hoc      │ Benannte Felder, selbstdokumentiert  │
│ Für temporäre Daten  │ Für wiederverwendbare Strukturen     │
│ _1, _2 unverständlich│ person.name ist klar                 │
└──────────────────────┴──────────────────────────────────────┘
```

**Faustregel:** Verwende Tupel nur für kurze, lokale Daten. Für alles andere Case Classes (kommt später).

---

## 2.7 Typkonvertierung

### Sichere Konvertierungen

```scala
val i: Int = 42
val l: Long = i        // Int → Long (automatisch, kein Datenverlust)
val d: Double = i      // Int → Double (automatisch)

val s: String = i.toString  // → "42"
```

### Explizite Konvertierungen

```scala
val d = 3.99
val i = d.toInt        // 3 (schneidet ab, rundet nicht!)

val s = "123"
val n = s.toInt        // 123
val f = s.toDouble     // 123.0

// Vorsicht: Exception bei ungültiger Eingabe!
// "abc".toInt  →  NumberFormatException
```

### Sichere Konvertierung mit Option

```scala
val s = "123"
val sicher: Option[Int] = s.toIntOption  // Some(123)

val ungueltig = "abc"
val nix: Option[Int] = ungueltig.toIntOption  // None
```

---

## 2.8 Spezielle Werte

### Unit

```scala
val nichts: Unit = ()  // Der einzige Wert vom Typ Unit

def sagHallo(): Unit =
  println("Hallo!")    // Rückgabe ist implizit ()
```

### Null (vermeiden!)

```scala
val leer: String = null  // Möglich, aber schlecht!

// Besser: Option verwenden (kommt später)
val sicher: Option[String] = None
```

### Nothing

`Nothing` ist der Subtyp aller Typen – Programme, die `Nothing` zurückgeben, terminieren nie normal:

```scala
def fehler(): Nothing = 
  throw new RuntimeException("Boom!")

def endlos(): Nothing =
  while true do ()
  throw new RuntimeException()  // Compiler braucht das
```

---

## 2.9 Type Aliases

Du kannst eigene Typnamen definieren:

```scala
type Name = String
type Alter = Int
type Person = (Name, Alter)

val alice: Person = ("Alice", 25)
```

---

## Zusammenfassung

| Typ-Kategorie | Typen | Beispiele |
|---------------|-------|-----------|
| Ganzzahlen | Byte, Short, Int, Long | `42`, `100L` |
| Fließkomma | Float, Double | `3.14f`, `3.14` |
| Zeichen | Char | `'A'` |
| Text | String | `"Hallo"` |
| Wahrheitswert | Boolean | `true`, `false` |
| Tupel | (A, B, C) | `(1, "text", true)` |
| Ohne Wert | Unit | `()` |

**Wichtige Konzepte:**
- String-Interpolation: `s"Hallo $name"`, `f"$pi%.2f"`
- Tupel-Zugriff: `tuple._1` oder Dekonstruktion
- Typinferenz: Compiler erkennt Typen automatisch
- Alles ist ein Objekt: `42.toString`, `true.&&(false)`

---

## Nächste Schritte

In **Vorlesung 3** lernen wir:
- Kontrollstrukturen: if/else, for, while
- Funktionen definieren und aufrufen
- Rekursion und Tail-Rekursion

➡️ [Weiter zu Vorlesung 3: Kontrollstrukturen & Funktionen](./03-kontrollstrukturen.md)
