# Vorlesung 3: Kontrollstrukturen & Funktionen

## Lernziele

Nach dieser Vorlesung kannst du:
- if/else als Ausdrücke verwenden
- For-Comprehensions für Iteration nutzen
- Funktionen definieren und aufrufen
- Rekursion verstehen und anwenden
- Tail-Rekursion für Effizienz nutzen

---

## 3.1 If/Else als Ausdruck

In Scala ist `if/else` **kein Statement, sondern ein Ausdruck** – es hat immer einen Rückgabewert!

### Grundsyntax

```scala
val zahl = 10

val ergebnis = if zahl > 0 then "positiv"
               else if zahl < 0 then "negativ"
               else "null"

println(ergebnis)  // "positiv"
```

### Vergleich mit anderen Sprachen

```
┌─────────────────────────────────────────────────────────────┐
│                Java (Statement)                             │
├─────────────────────────────────────────────────────────────┤
│ String ergebnis;                                            │
│ if (zahl > 0) {                                             │
│     ergebnis = "positiv";                                   │
│ } else {                                                    │
│     ergebnis = "negativ";                                   │
│ }                                                           │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                Scala (Ausdruck)                             │
├─────────────────────────────────────────────────────────────┤
│ val ergebnis = if zahl > 0 then "positiv" else "negativ"    │
└─────────────────────────────────────────────────────────────┘
```

### Ohne else

Wenn kein `else` vorhanden ist, wird `Unit` durchgereicht:

```scala
val x = if (true) then 42
// x hat Typ: AnyVal (weil else-Fall Unit wäre)

// Besser explizit:
val y = if true then 42 else 0
// y hat Typ: Int
```

### Blöcke in if/else

```scala
val punkte = 85

val note = 
  if punkte >= 90 then
    println("Sehr gut!")
    "1.0"
  else if punkte >= 80 then
    println("Gut!")
    "2.0"
  else
    println("Bestanden")
    "3.0"
```

---

## 3.2 Match-Ausdrücke (Vorschau)

Für komplexere Fallunterscheidungen gibt es `match`:

```scala
val tag = 3

val wochentag = tag match
  case 1 => "Montag"
  case 2 => "Dienstag"
  case 3 => "Mittwoch"
  case 4 => "Donnerstag"
  case 5 => "Freitag"
  case 6 => "Samstag"
  case 7 => "Sonntag"
  case _ => "Ungültig"  // Default-Fall

println(wochentag)  // "Mittwoch"
```

> Pattern Matching wird ausführlich in **Vorlesung 7** behandelt!

---

## 3.3 For-Schleifen und For-Comprehensions

### Einfache For-Schleife

```scala
// Iteration über Range
for i <- 1 to 5 do
  println(i)
// Ausgabe: 1, 2, 3, 4, 5

// Mit 'until' (exklusive Obergrenze)
for i <- 1 until 5 do
  println(i)
// Ausgabe: 1, 2, 3, 4
```

### Über Collections iterieren

```scala
val namen = List("Alice", "Bob", "Charlie")

for name <- namen do
  println(s"Hallo $name!")
```

### Mit Filter (Guard)

```scala
for i <- 1 to 10 if i % 2 == 0 do
  println(i)
// Ausgabe: 2, 4, 6, 8, 10
```

### Mehrere Generatoren (verschachtelt)

```scala
for 
  x <- 1 to 3
  y <- 1 to 3
do
  println(s"($x, $y)")

// Ausgabe: (1,1), (1,2), (1,3), (2,1), (2,2), ...
```

### For-Comprehension mit yield

`yield` erzeugt eine **neue Collection**:

```scala
val zahlen = List(1, 2, 3, 4, 5)

val verdoppelt = for n <- zahlen yield n * 2
// verdoppelt: List(2, 4, 6, 8, 10)

val geradeQuadrate = 
  for 
    n <- zahlen 
    if n % 2 == 0
  yield n * n
// geradeQuadrate: List(4, 16)
```

### For-Comprehension vs. Methoden

```scala
// Diese zwei sind äquivalent:

// For-Comprehension
val result1 = for n <- zahlen if n > 2 yield n * 2

// Methodenketten
val result2 = zahlen.filter(n => n > 2).map(n => n * 2)
```

---

## 3.4 While-Schleifen (Vermeiden!)

Scala hat while-Schleifen, aber sie sind **selten nötig**:

```scala
var i = 0
while i < 5 do
  println(i)
  i += 1
```

> **Warnung**: `while` erfordert `var`, was gegen funktionalen Stil verstößt. Verwende stattdessen Rekursion oder Collection-Methoden!

---

## 3.5 Funktionen Definieren

### Grundsyntax

```scala
def funktionsname(parameter: Typ, ...): RückgabeTyp =
  // Funktionskörper (letzte Zeile = Rückgabewert)
```

### Beispiele

```scala
// Einfache Funktion
def quadrat(x: Int): Int = x * x

// Mehrere Parameter
def addiere(a: Int, b: Int): Int = a + b

// Mit Block
def begruessung(name: String): String =
  val nachricht = s"Hallo $name!"
  nachricht.toUpperCase  // Rückgabewert

// Aufruf
println(quadrat(5))           // 25
println(addiere(3, 4))        // 7
println(begruessung("Scala")) // "HALLO SCALA!"
```

### Einzeilige Funktionen

```scala
def verdopple(x: Int): Int = x * 2
def istGerade(n: Int): Boolean = n % 2 == 0
def max(a: Int, b: Int): Int = if a > b then a else b
```

### Funktionen ohne Parameter

```scala
def aktuelleZeit(): String = 
  java.time.LocalTime.now().toString

println(aktuelleZeit())  // z.B. "14:30:45.123"
```

### Funktionen mit Unit-Rückgabe

```scala
def sagHallo(name: String): Unit =
  println(s"Hallo $name!")

// Alternative Schreibweise (veraltet, aber noch gültig)
def sagTschuess(name: String) =
  println(s"Tschüss $name!")
```

---

## 3.6 Typinferenz bei Funktionen

Der Compiler kann Rückgabetypen oft erkennen:

```scala
// Typ erkannt: Int
def quadrat(x: Int) = x * x

// Typ erkannt: String
def format(n: Int) = s"Die Zahl ist $n"
```

**Best Practice:** Für **öffentliche** Funktionen immer Typen angeben!

```scala
// ✓ Gut: Typ dokumentiert
def berechneGehalt(stunden: Int, satz: Double): Double =
  stunden * satz

// ✗ Schlecht für öffentliche API
def berechneGehalt(stunden: Int, satz: Double) =
  stunden * satz
```

---

## 3.7 Default-Parameter und Named Arguments

### Default-Parameter

```scala
def begruessung(name: String, gruss: String = "Hallo"): String =
  s"$gruss $name!"

println(begruessung("Max"))           // "Hallo Max!"
println(begruessung("Max", "Hi"))     // "Hi Max!"
```

### Named Arguments

```scala
def erstelleAdresse(
  strasse: String, 
  nummer: Int, 
  plz: String, 
  stadt: String
): String =
  s"$strasse $nummer, $plz $stadt"

// Mit benannten Argumenten (Reihenfolge egal!)
val adresse = erstelleAdresse(
  stadt = "Tübingen",
  plz = "72074",
  strasse = "Wilhelmstr.",
  nummer = 7
)
```

---

## 3.8 Rekursion

### Was ist Rekursion?

Eine Funktion ruft **sich selbst** auf mit einem einfacheren Problem.

### Beispiel: Fakultät

```
n! = n × (n-1) × (n-2) × ... × 1
5! = 5 × 4 × 3 × 2 × 1 = 120
```

```scala
def fakultaet(n: Int): Int =
  if n <= 1 then 1
  else n * fakultaet(n - 1)

println(fakultaet(5))  // 120
```

### Ablauf visualisiert

```
fakultaet(5)
= 5 * fakultaet(4)
= 5 * (4 * fakultaet(3))
= 5 * (4 * (3 * fakultaet(2)))
= 5 * (4 * (3 * (2 * fakultaet(1))))
= 5 * (4 * (3 * (2 * 1)))
= 5 * (4 * (3 * 2))
= 5 * (4 * 6)
= 5 * 24
= 120
```

### Beispiel: Fibonacci

```scala
def fibonacci(n: Int): Int =
  if n <= 1 then n
  else fibonacci(n - 1) + fibonacci(n - 2)

// fib(0) = 0, fib(1) = 1
// fib(2) = 1, fib(3) = 2, fib(4) = 3, fib(5) = 5, ...
```

---

## 3.9 Tail-Rekursion

### Das Stack-Problem

Normale Rekursion baut einen großen Call-Stack auf:

```scala
fakultaet(10000)  // StackOverflowError!
```

### Tail-Rekursion als Lösung

Bei **Tail-Rekursion** ist der rekursive Aufruf die **letzte** Operation:

```scala
import scala.annotation.tailrec

def fakultaetTail(n: Int): BigInt =
  @tailrec
  def loop(current: Int, akkumulator: BigInt): BigInt =
    if current <= 1 then akkumulator
    else loop(current - 1, akkumulator * current)
  
  loop(n, 1)

println(fakultaetTail(10000))  // Funktioniert!
```

### Vergleich

```
┌─────────────────────────────────────────────────────────────┐
│            Normale Rekursion vs Tail-Rekursion              │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Normale Rekursion:                                         │
│  n * fakultaet(n-1)                                         │
│  ↑                                                          │
│  Nach dem rekursiven Aufruf muss noch multipliziert werden  │
│  → Stack wächst                                             │
│                                                             │
│  Tail-Rekursion:                                            │
│  loop(n-1, acc * n)                                         │
│  ↑                                                          │
│  Rekursiver Aufruf ist die LETZTE Operation                 │
│  → Compiler optimiert zu Schleife                           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### @tailrec Annotation

```scala
import scala.annotation.tailrec

@tailrec
def summe(zahlen: List[Int], akkumulator: Int = 0): Int =
  zahlen match
    case Nil => akkumulator
    case kopf :: rest => summe(rest, akkumulator + kopf)
```

Die `@tailrec`-Annotation prüft zur Compile-Zeit, ob die Funktion wirklich tail-rekursiv ist!

---

## 3.10 Lokale Funktionen

Funktionen können in Funktionen definiert werden:

```scala
def berechneStatistik(zahlen: List[Int]): (Int, Int, Double) =
  // Lokale Hilfsfunktionen
  def minimum(liste: List[Int]): Int = liste.min
  def maximum(liste: List[Int]): Int = liste.max
  def durchschnitt(liste: List[Int]): Double = 
    liste.sum.toDouble / liste.length
  
  (minimum(zahlen), maximum(zahlen), durchschnitt(zahlen))

val (min, max, avg) = berechneStatistik(List(1, 2, 3, 4, 5))
println(s"Min: $min, Max: $max, Avg: $avg")
```

---

## Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| `if/else` | Ausdruck mit Rückgabewert |
| `for...do` | Schleife für Seiteneffekte |
| `for...yield` | Transformation zu neuer Collection |
| `def` | Funktionsdefinition |
| Parameter | Mit Typ: `name: Typ` |
| Default-Parameter | `name: Typ = wert` |
| Named Arguments | `funk(param = wert)` |
| Rekursion | Funktion ruft sich selbst auf |
| Tail-Rekursion | Letzter Aufruf ist rekursiv → O(1) Stack |
| `@tailrec` | Compiler prüft Tail-Rekursion |

---

## Nächste Schritte

In **Vorlesung 4** lernen wir:
- Klassen und Objekte
- Konstruktoren
- Case Classes
- Companion Objects

➡️ [Weiter zu Vorlesung 4: Klassen & Objekte](./04-klassen-objekte.md)
