# Übung 3: Kontrollstrukturen & Funktionen

**Thema:** If/Else, For-Comprehensions, Funktionen, Rekursion

**Geschätzte Zeit:** 90-120 Minuten

---

## Teil A: Theoretische Fragen

### Frage 3.1
Warum ist `if/else` in Scala ein "Ausdruck" und nicht ein "Statement"? Was ist der Vorteil davon?

### Frage 3.2
Erkläre den Unterschied zwischen `for ... do` und `for ... yield`. Wann verwendest du welches?

### Frage 3.3
Was ist Tail-Rekursion? Warum ist sie wichtig? Was macht die `@tailrec`-Annotation?

### Frage 3.4
Was ist der Unterschied zwischen diesen zwei Schreibweisen?

```scala
def f1(x: Int) = x * 2
def f2(x: Int): Int = x * 2
```

### Frage 3.5
Warum sollte man `while`-Schleifen in Scala vermeiden?

---

## Teil B: REPL-Übungen

### Aufgabe 3.1 – If/Else als Ausdruck
Schreibe in der REPL:

```scala
val zahl = 42
val vorzeichen = if zahl > 0 then "positiv" else if zahl < 0 then "negativ" else "null"

// Teste mit verschiedenen Werten
val x = -5
val y = 0
```

### Aufgabe 3.2 – For-Comprehensions
Experimentiere mit For-Ausdrücken:

```scala
// Quadratzahlen von 1 bis 10
for i <- 1 to 10 yield i * i

// Nur gerade Quadrate
for i <- 1 to 10 if i % 2 == 0 yield i * i

// Paare
for 
  a <- 1 to 3
  b <- 1 to 3
yield (a, b)
```

### Aufgabe 3.3 – Funktionen definieren
Definiere und teste diese Funktionen in der REPL:

```scala
def greet(name: String): String = s"Hello, $name!"
def add(a: Int, b: Int): Int = a + b
def isEven(n: Int): Boolean = ???
def absoluteValue(n: Int): Int = ???
```

---

## Teil C: Programmieraufgaben

### Aufgabe 3.4 – Notenberechnung
Schreibe eine Funktion `berechneNote(punkte: Int): String`, die:
- Bei 90+ Punkten "Sehr gut (1)" zurückgibt
- Bei 80-89 Punkten "Gut (2)"
- Bei 70-79 Punkten "Befriedigend (3)"
- Bei 60-69 Punkten "Ausreichend (4)"
- Unter 60 Punkten "Nicht bestanden (5)"

```scala
@main def notenTest(): Unit =
  println(berechneNote(95))  // "Sehr gut (1)"
  println(berechneNote(75))  // "Befriedigend (3)"
  println(berechneNote(45))  // "Nicht bestanden (5)"
```

### Aufgabe 3.5 – FizzBuzz
Implementiere das klassische FizzBuzz-Problem:
- Für Zahlen von 1 bis n:
  - Wenn durch 3 teilbar: "Fizz"
  - Wenn durch 5 teilbar: "Buzz"
  - Wenn durch 3 und 5 teilbar: "FizzBuzz"
  - Sonst: die Zahl selbst

```scala
def fizzBuzz(n: Int): List[String] = ???

// fizzBuzz(15) sollte enthalten:
// "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", 
// "11", "Fizz", "13", "14", "FizzBuzz"
```

### Aufgabe 3.6 – Rekursive Summe
Schreibe eine rekursive Funktion, die die Summe der Zahlen von 1 bis n berechnet:

```scala
def summe(n: Int): Int = ???

// summe(5) = 1 + 2 + 3 + 4 + 5 = 15
// summe(10) = 55
```

### Aufgabe 3.7 – Tail-rekursive Summe ⭐
Schreibe die Summen-Funktion als **tail-rekursive** Version:

```scala
import scala.annotation.tailrec

def summeTail(n: Int): Int =
  @tailrec
  def loop(current: Int, akkumulator: Int): Int = ???
  
  loop(n, 0)
```

### Aufgabe 3.8 – Primzahlen
Schreibe Funktionen zum Arbeiten mit Primzahlen:

```scala
// Prüft, ob eine Zahl prim ist
def istPrim(n: Int): Boolean = ???

// Gibt alle Primzahlen bis n zurück
def primzahlenBis(n: Int): List[Int] = ???

// Test
println(istPrim(7))   // true
println(istPrim(10))  // false
println(primzahlenBis(20)) // List(2, 3, 5, 7, 11, 13, 17, 19)
```

### Aufgabe 3.9 – Collatz-Folge ⭐
Die Collatz-Vermutung: Starte mit einer Zahl n.
- Wenn n gerade: n / 2
- Wenn n ungerade: 3n + 1
- Wiederhole, bis n = 1

Schreibe eine Funktion, die die Collatz-Folge berechnet:

```scala
def collatz(n: Int): List[Int] = ???

// collatz(6) = List(6, 3, 10, 5, 16, 8, 4, 2, 1)
```

---

## Teil D: Verständnisfragen

### Frage 3.6
Was ist das Ergebnis von:

```scala
val xs = for i <- 1 to 5 if i != 3 yield i * 10
```

### Frage 3.7
Ist diese Funktion tail-rekursiv? Begründe!

```scala
def length(list: List[Int]): Int =
  if list.isEmpty then 0
  else 1 + length(list.tail)
```

### Frage 3.8
Was gibt diese Funktion zurück?

```scala
def mystery(n: Int): Int =
  if n == 0 then 0
  else n + mystery(n - 1)

mystery(4) // = ?
```

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – Pascal'sches Dreieck
Berechne den Wert an Position (Zeile, Spalte) im Pascal'schen Dreieck:

```
        1           Zeile 0
       1 1          Zeile 1
      1 2 1         Zeile 2
     1 3 3 1        Zeile 3
    1 4 6 4 1       Zeile 4
```

```scala
def pascal(spalte: Int, zeile: Int): Int = ???

// pascal(0, 0) = 1
// pascal(1, 2) = 2
// pascal(2, 4) = 6
```

### Bonus 2 – Größter gemeinsamer Teiler (GGT)
Implementiere den Euklidischen Algorithmus:

```scala
@tailrec
def ggt(a: Int, b: Int): Int = ???

// ggt(48, 18) = 6
// ggt(100, 35) = 5
```

### Bonus 3 – Alle Permutationen
Schreibe eine Funktion, die alle Permutationen einer Liste zurückgibt:

```scala
def permutationen[A](liste: List[A]): List[List[A]] = ???

// permutationen(List(1, 2, 3)) = 
// List(
//   List(1, 2, 3), List(1, 3, 2),
//   List(2, 1, 3), List(2, 3, 1),
//   List(3, 1, 2), List(3, 2, 1)
// )
```

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-03.md](../loesungen/loesung-03.md)
