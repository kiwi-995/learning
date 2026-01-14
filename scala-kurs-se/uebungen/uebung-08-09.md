# Übung 8-9: Higher-Order Functions & Fortgeschrittene FP

**Thema:** HOFs, Currying, Closures, Option, Either, For-Comprehensions

**Geschätzte Zeit:** 120 Minuten

---

## Teil A: Higher-Order Functions

### Aufgabe 8.1 – Funktionen als Werte
Definiere folgende Funktionen als Werte:

```scala
val verdreifache: Int => Int = ???
val istPositiv: Int => Boolean = ???
val gruss: String => String = ???
```

### Aufgabe 8.2 – Funktionen als Parameter
Implementiere eine flexiblere Filter-Funktion:

```scala
def filtereUndTransformiere[A, B](
  liste: List[A],
  bedingung: A => Boolean,
  transformation: A => B
): List[B] = ???
```

### Aufgabe 8.3 – Currying
Schreibe these Funktionen in curried Form:

```scala
def multipliziere(a: Int)(b: Int): Int = ???
def formatiere(prefix: String)(wert: Any)(suffix: String): String = ???
```

### Aufgabe 8.4 – Closures
Was gibt dieser Code aus? Erkläre warum.

```scala
def createCounters(n: Int): List[() => Int] =
  (1 to n).map(i => () => i * i).toList

val counters = createCounters(3)
counters.foreach(c => println(c()))
```

---

## Teil B: Option und Either

### Aufgabe 8.5 – Option verwenden
```scala
def sichereDivision(a: Int, b: Int): Option[Int] = ???
def sucheInListe[A](liste: List[A], praedikat: A => Boolean): Option[A] = ???
def parseAlter(s: String): Option[Int] = ???  // Nur 0-150 gültig
```

### Aufgabe 8.6 – Either für Validierung
```scala
sealed trait Fehler
case class NameZuKurz(min: Int) extends Fehler
case class EmailUngueltig(grund: String) extends Fehler
case class AlterUngueltig(grund: String) extends Fehler

case class BenutzerDaten(name: String, email: String, alter: Int)

def validiereBenutzer(
  name: String, 
  email: String, 
  alter: String
): Either[Fehler, BenutzerDaten] = ???
```

### Aufgabe 8.7 – Verkettung mit For-Comprehension
```scala
def komplexeBerechnung(
  a: String, 
  b: String, 
  c: String
): Either[String, Int] =
  for
    x <- a.toIntOption.toRight("a ist keine Zahl")
    y <- b.toIntOption.toRight("b ist keine Zahl")
    z <- c.toIntOption.toRight("c ist keine Zahl")
    ergebnis <- ???  // x * y / z, mit Prüfung auf Division durch 0
  yield ergebnis
```

---

## Teil C: Praktische Aufgaben

### Aufgabe 8.8 – Datenverarbeitung Pipeline
```scala
case class Produkt(name: String, preis: Double, kategorie: String)

val produkte = List(
  Produkt("Laptop", 999.99, "Elektronik"),
  Produkt("Buch", 19.99, "Bücher"),
  Produkt("Handy", 599.99, "Elektronik"),
  Produkt("Kaffee", 12.99, "Lebensmittel")
)

// 1. Alle Produkte über 50€
// 2. Nur die Namen
// 3. Alphabetisch sortiert
// 4. Als kommagetrennter String

def teuereProduktNamen(produkte: List[Produkt], minPreis: Double): String = ???
```

### Aufgabe 8.9 – Eigenes flatMap
```scala
def meinFlatMap[A, B](opt: Option[A])(f: A => Option[B]): Option[B] = ???

// Test
meinFlatMap(Some(5))(x => if x > 0 then Some(x * 2) else None)  // Some(10)
meinFlatMap(Some(-5))(x => if x > 0 then Some(x * 2) else None) // None
meinFlatMap(None)(x => Some(x))  // None
```

### Aufgabe 8.10 – traverse und sequence ⭐
```scala
def sequence[A](liste: List[Option[A]]): Option[List[A]] = ???
def traverse[A, B](liste: List[A])(f: A => Option[B]): Option[List[B]] = ???

// Test
sequence(List(Some(1), Some(2), Some(3)))  // Some(List(1, 2, 3))
sequence(List(Some(1), None, Some(3)))     // None

traverse(List("1", "2", "3"))(_.toIntOption)  // Some(List(1, 2, 3))
traverse(List("1", "x", "3"))(_.toIntOption)  // None
```

---

## Lösungen

```scala
// 8.1
val verdreifache: Int => Int = x => x * 3
val istPositiv: Int => Boolean = _ > 0
val gruss: String => String = name => s"Hallo, $name!"

// 8.2
def filtereUndTransformiere[A, B](
  liste: List[A],
  bedingung: A => Boolean,
  transformation: A => B
): List[B] = liste.filter(bedingung).map(transformation)

// 8.5
def sichereDivision(a: Int, b: Int): Option[Int] =
  if b == 0 then None else Some(a / b)

// 8.9
def meinFlatMap[A, B](opt: Option[A])(f: A => Option[B]): Option[B] = opt match
  case Some(a) => f(a)
  case None => None

// 8.10
def sequence[A](liste: List[Option[A]]): Option[List[A]] =
  liste.foldRight(Some(Nil): Option[List[A]]) { (opt, acc) =>
    for a <- opt; as <- acc yield a :: as
  }

def traverse[A, B](liste: List[A])(f: A => Option[B]): Option[List[B]] =
  sequence(liste.map(f))
```
