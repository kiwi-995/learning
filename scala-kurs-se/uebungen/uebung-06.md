# Übung 6: Collections

**Thema:** List, Vector, Set, Map, map/filter/fold, For-Comprehensions

**Geschätzte Zeit:** 90-120 Minuten

---

## Teil A: Theoretische Fragen

### Frage 6.1
Was ist der Unterschied zwischen `foldLeft` und `foldRight`? Wann macht der Unterschied etwas aus?

### Frage 6.2
Erkläre den Unterschied zwischen `map` und `flatMap`. Wann verwendest du welches?

### Frage 6.3
Warum ist `0 :: liste` bei einer List effizient, aber `liste :+ 0` langsam?

### Frage 6.4
Was ist der Unterschied zwischen immutable und mutable Collections? Warum bevorzugt man in Scala immutable?

---

## Teil B: REPL-Übungen

### Aufgabe 6.1 – Grundoperationen
```scala
val zahlen = List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

// Teste folgende Operationen:
zahlen.head
zahlen.tail
zahlen.take(3)
zahlen.drop(3)
zahlen.filter(_ % 2 == 0)
zahlen.map(_ * 2)
zahlen.sum
zahlen.product
```

### Aufgabe 6.2 – Verkettung
```scala
val worte = List("Scala", "ist", "toll")

// Kombiniere zu:
// 1. "Scala ist toll" (mit Leerzeichen)
// 2. Alle Großbuchstaben
// 3. Nur Worte mit mehr als 3 Buchstaben
```

### Aufgabe 6.3 – Map und Set
```scala
val studentennoten = Map(
  "Alice" -> 1.3,
  "Bob" -> 2.7,
  "Charlie" -> 1.0,
  "Diana" -> 2.0
)

// 1. Finde Alices Note
// 2. Alle Studenten mit Note besser als 2.0
// 3. Durchschnittsnote berechnen
// 4. Füge "Eve" mit Note 1.7 hinzu
```

---

## Teil C: Programmieraufgaben

### Aufgabe 6.4 – Statistik
Implementiere Statistik-Funktionen:

```scala
object Statistik:
  def mittelwert(zahlen: List[Double]): Double = ???
  def varianz(zahlen: List[Double]): Double = ???
  def standardabweichung(zahlen: List[Double]): Double = ???
  def median(zahlen: List[Double]): Double = ???
  def modus(zahlen: List[Double]): Double = ???  // Häufigster Wert
```

### Aufgabe 6.5 – Wortzähler
Schreibe einen Wortzähler:

```scala
def zähleWorte(text: String): Map[String, Int] = ???

val text = "Scala ist toll Scala ist funktional Scala ist objektorientiert"
val gezählt = zähleWorte(text)
// Map("Scala" -> 3, "ist" -> 3, "toll" -> 1, "funktional" -> 1, "objektorientiert" -> 1)
```

### Aufgabe 6.6 – Matrixoperationen
Implementiere Matrizen als Listen von Listen:

```scala
type Matrix = List[List[Double]]

object MatrixOps:
  def transponieren(m: Matrix): Matrix = ???
  def addieren(a: Matrix, b: Matrix): Matrix = ???
  def multiplizieren(a: Matrix, b: Matrix): Matrix = ???
  def skalar(m: Matrix, s: Double): Matrix = ???
  def alsString(m: Matrix): String = ???
```

### Aufgabe 6.7 – Primfaktorzerlegung ⭐
```scala
def primfaktoren(n: Int): List[Int] = ???

primfaktoren(12)   // List(2, 2, 3)
primfaktoren(100)  // List(2, 2, 5, 5)
primfaktoren(17)   // List(17)
```

### Aufgabe 6.8 – Gruppierung ⭐
Implementiere eine Funktion, die Elemente nach einem Kriterium gruppiert:

```scala
case class Student(name: String, fach: String, semester: Int)

val studenten = List(
  Student("Alice", "Informatik", 3),
  Student("Bob", "Mathematik", 2),
  Student("Charlie", "Informatik", 2),
  Student("Diana", "Physik", 1),
  Student("Eve", "Informatik", 3)
)

// 1. Gruppiere nach Fach
// 2. Gruppiere nach Semester
// 3. Finde alle Informatik-Studenten im 3. Semester
// 4. Berechne Durchschnittssemester pro Fach
```

---

## Teil D: For-Comprehensions

### Aufgabe 6.9
Schreibe folgende Ausdrücke als For-Comprehension und umgekehrt:

```scala
// 1. Als For-Comprehension schreiben:
List(1, 2, 3).flatMap(x => List(4, 5, 6).map(y => x * y))

// 2. Als map/flatMap/filter schreiben:
for
  x <- List(1, 2, 3)
  y <- List(1, 2, 3)
  if x != y
yield (x, y)
```

### Aufgabe 6.10 – Schachbrett
Generiere alle Positionen auf einem Schachbrett:

```scala
val positionen = ??? // Alle 64 Positionen als List[(Char, Int)]
// List(('a', 1), ('a', 2), ..., ('h', 8))
```

### Aufgabe 6.11 – Pythagoräische Tripel ⭐
Finde alle pythagoräischen Tripel (a, b, c) mit a² + b² = c² und a, b, c ≤ 100:

```scala
val tripel = for
  ???
yield ???

// Sollte enthalten: (3, 4, 5), (5, 12, 13), (8, 15, 17), ...
```

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – Eigene Collection-Methoden
Implementiere diese Methoden selbst (ohne die eingebauten zu verwenden):

```scala
def meinMap[A, B](liste: List[A])(f: A => B): List[B] = ???
def meinFilter[A](liste: List[A])(p: A => Boolean): List[A] = ???
def meinFold[A, B](liste: List[A])(start: B)(f: (B, A) => B): B = ???
def meinFlatMap[A, B](liste: List[A])(f: A => List[B]): List[B] = ???
```

### Bonus 2 – Sliding Window Maximum
Finde für jedes Fenster der Größe k das Maximum:

```scala
def slidingMax(zahlen: List[Int], k: Int): List[Int] = ???

slidingMax(List(1, 3, -1, -3, 5, 3, 6, 7), 3)
// List(3, 3, 5, 5, 6, 7)
// Fenster: [1,3,-1] -> 3, [3,-1,-3] -> 3, [-1,-3,5] -> 5, ...
```

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-06.md](../loesungen/loesung-06.md)
