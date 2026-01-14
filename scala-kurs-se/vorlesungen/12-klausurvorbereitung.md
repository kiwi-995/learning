# Vorlesung 12: Klausurvorbereitung

## Lernziele

Diese Vorlesung bereitet dich optimal auf die Programmierklausur vor:
- Wichtige Konzepte wiederholen
- Typische Klausuraufgaben üben
- Häufige Fehler vermeiden
- Zeitmanagement-Tipps

---

## 12.1 Wichtige Scala-Konzepte

### val vs var

```scala
val unveraenderlich = 42    // Kann nicht neu zugewiesen werden
var veraenderlich = 42      // Kann geändert werden
veraenderlich = 100         // OK

// IMMER val bevorzugen!
```

### Funktionen definieren

```scala
// Mit Rückgabetyp (empfohlen für Klausur!)
def addiere(a: Int, b: Int): Int = a + b

// Rekursiv
def fakultaet(n: Int): Int =
  if n <= 1 then 1
  else n * fakultaet(n - 1)

// Tail-rekursiv
import scala.annotation.tailrec
@tailrec
def summe(n: Int, acc: Int = 0): Int =
  if n <= 0 then acc
  else summe(n - 1, acc + n)
```

### Pattern Matching

```scala
x match
  case 0 => "Null"
  case n if n > 0 => "Positiv"
  case _ => "Negativ"

// Auf Case Classes
person match
  case Person(name, alter) if alter >= 18 => s"$name ist erwachsen"
  case Person(name, _) => s"$name ist minderjährig"

// Auf Listen
liste match
  case Nil => "Leer"
  case kopf :: rest => s"Erstes: $kopf"
```

### Case Classes

```scala
case class Person(name: String, alter: Int)

val p = Person("Max", 25)
p.name                    // "Max"
p.copy(alter = 26)        // Person("Max", 26)
p == Person("Max", 25)    // true (strukturell)
```

### Sealed Traits (ADTs)

```scala
sealed trait Ausdruck
case class Zahl(n: Int) extends Ausdruck
case class Plus(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Mal(l: Ausdruck, r: Ausdruck) extends Ausdruck

def auswerten(e: Ausdruck): Int = e match
  case Zahl(n) => n
  case Plus(l, r) => auswerten(l) + auswerten(r)
  case Mal(l, r) => auswerten(l) * auswerten(r)
```

---

## 12.2 Wichtige Collection-Operationen

```scala
val zahlen = List(1, 2, 3, 4, 5)

// map - transformieren
zahlen.map(_ * 2)              // List(2, 4, 6, 8, 10)

// filter - filtern
zahlen.filter(_ % 2 == 0)      // List(2, 4)

// fold - reduzieren
zahlen.fold(0)(_ + _)          // 15

// flatMap - map + flatten
List(1, 2).flatMap(x => List(x, x * 10))  // List(1, 10, 2, 20)

// For-Comprehension
for
  x <- zahlen
  if x > 2
yield x * 2                    // List(6, 8, 10)
```

---

## 12.3 Option, Either, Try

```scala
// Option
val opt: Option[Int] = Some(42)
opt.getOrElse(0)              // 42
opt.map(_ * 2)                // Some(84)

// Either
val result: Either[String, Int] = Right(42)
for
  x <- result
yield x * 2                   // Right(84)

// Try
import scala.util.Try
Try("42".toInt).getOrElse(0)  // 42
```

---

## 12.4 Git-Befehle für die Klausur

```bash
# Repository
git init
git clone <url>

# Änderungen
git status
git add <datei>
git commit -m "Message"

# Branches
git branch <name>
git checkout <name>
git merge <name>

# Remote
git push
git pull

# Historie
git log --oneline
```

---

## 12.5 Typische Klausuraufgaben

### Aufgabe 1: Rekursion

```scala
// Berechne die Länge einer Liste rekursiv
def laenge[A](liste: List[A]): Int = liste match
  case Nil => 0
  case _ :: rest => 1 + laenge(rest)

// Tail-rekursiv
def laengeTail[A](liste: List[A]): Int =
  @tailrec
  def loop(l: List[A], acc: Int): Int = l match
    case Nil => acc
    case _ :: rest => loop(rest, acc + 1)
  loop(liste, 0)
```

### Aufgabe 2: ADT und Pattern Matching

```scala
sealed trait Baum[+A]
case object Leer extends Baum[Nothing]
case class Knoten[A](wert: A, links: Baum[A], rechts: Baum[A]) extends Baum[A]

def groesse[A](baum: Baum[A]): Int = baum match
  case Leer => 0
  case Knoten(_, l, r) => 1 + groesse(l) + groesse(r)

def summe(baum: Baum[Int]): Int = baum match
  case Leer => 0
  case Knoten(w, l, r) => w + summe(l) + summe(r)
```

### Aufgabe 3: Higher-Order Functions

```scala
// Eigenes map implementieren
def meinMap[A, B](liste: List[A], f: A => B): List[B] = liste match
  case Nil => Nil
  case kopf :: rest => f(kopf) :: meinMap(rest, f)

// Eigenes filter
def meinFilter[A](liste: List[A], p: A => Boolean): List[A] = liste match
  case Nil => Nil
  case kopf :: rest =>
    if p(kopf) then kopf :: meinFilter(rest, p)
    else meinFilter(rest, p)
```

### Aufgabe 4: For-Comprehension

```scala
// Kartesisches Produkt
val paare = for
  x <- List(1, 2, 3)
  y <- List("a", "b")
yield (x, y)
// List((1,a), (1,b), (2,a), (2,b), (3,a), (3,b))

// Mit Filter
val gefiltert = for
  x <- 1 to 10
  y <- 1 to 10
  if x + y == 10
yield (x, y)
```

---

## 12.6 Häufige Fehler vermeiden

### ✗ Falsch: var statt val

```scala
// Schlecht
var ergebnis = 0
for x <- liste do
  ergebnis += x

// Besser
val ergebnis = liste.sum
// oder
val ergebnis = liste.fold(0)(_ + _)
```

### ✗ Falsch: null verwenden

```scala
// Schlecht
def finde(id: Int): String = 
  if gefunden then wert else null

// Besser
def finde(id: Int): Option[String] =
  if gefunden then Some(wert) else None
```

### ✗ Falsch: return verwenden

```scala
// Schlecht
def max(a: Int, b: Int): Int = {
  if (a > b) return a
  else return b
}

// Besser
def max(a: Int, b: Int): Int =
  if a > b then a else b
```

### ✗ Falsch: Nicht exhaustive Patterns

```scala
// Warnung: nicht alle Fälle behandelt!
sealed trait Status
case object Aktiv extends Status
case object Inaktiv extends Status

def pruefen(s: Status) = s match
  case Aktiv => "OK"
  // Inaktiv fehlt!
```

---

## 12.7 Klausur-Tipps

### Zeitmanagement

```
┌─────────────────────────────────────────────────────────────┐
│ 1. Alle Aufgaben durchlesen (5 min)                         │
│ 2. Leichte Aufgaben zuerst                                  │
│ 3. Zeit pro Aufgabe einteilen                               │
│ 4. Nicht an einer Aufgabe festbeißen                        │
│ 5. Am Ende nochmal durchgehen                               │
└─────────────────────────────────────────────────────────────┘
```

### Vor dem Abgeben prüfen

- [ ] Typen angegeben?
- [ ] Pattern Matching exhaustive?
- [ ] Rekursion hat Basisfall?
- [ ] `@tailrec` wo nötig?
- [ ] Keine `var`, `null`, `return`?

---

## 12.8 Probeklausur

Bearbeite die folgenden Aufgaben unter Zeitdruck (90 Minuten):

### Aufgabe 1 (15 Punkte)
Implementiere eine Funktion `verdoppele`, die alle Elemente einer Liste verdoppelt.

### Aufgabe 2 (20 Punkte)
Definiere einen ADT für arithmetische Ausdrücke mit Zahlen, Addition und Multiplikation. Implementiere eine Auswertungsfunktion.

### Aufgabe 3 (15 Punkte)
Schreibe eine tail-rekursive Funktion `umkehren`, die eine Liste umkehrt.

### Aufgabe 4 (20 Punkte)
Implementiere einen einfachen binären Suchbaum mit `einfuegen` und `enthaelt`.

### Aufgabe 5 (15 Punkte)
Schreibe Git-Befehle für: Repo klonen, Branch erstellen, Änderungen committen, pushen.

### Aufgabe 6 (15 Punkte)
Erkläre den Unterschied zwischen `map`, `flatMap` und `filter`. Gib je ein Beispiel.

---

## Zusammenfassung

**Wichtigste Scala-Konzepte:**
- `val` statt `var`
- Pattern Matching
- Case Classes und ADTs
- Higher-Order Functions
- Option statt null

**Wichtigste Git-Konzepte:**
- Feature Branch Workflow
- Commit, Push, Pull
- Merge und Konflikte

**Klausur-Strategie:**
- Leichte Aufgaben zuerst
- Typen immer angeben
- Compiler-Freundlich programmieren

---

## Viel Erfolg bei der Klausur! 🎓
