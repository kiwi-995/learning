# Lösungen zu Übung 3

---

## Teil A: Theoretische Fragen

### Lösung 3.1
In Scala ist `if/else` ein **Ausdruck**, d.h. er hat immer einen Rückgabewert.

**Vorteile:**
- Direktes Zuweisen möglich: `val x = if (bed) a else b`
- Kein Deklarieren von Variablen nötig
- Fördert funktionalen Stil (keine `var`)
- Kompakterer Code

```scala
// Java-Stil (Statement)
String result;
if (x > 0) { result = "positiv"; }
else { result = "negativ"; }

// Scala-Stil (Ausdruck)
val result = if x > 0 then "positiv" else "negativ"
```

### Lösung 3.2
- **`for ... do`**: Führt Seiteneffekte aus (println, etc.), gibt `Unit` zurück
- **`for ... yield`**: Transformiert Elemente und gibt neue Collection zurück

```scala
// for ... do → Seiteneffekte
for i <- 1 to 3 do println(i)  // Typ: Unit

// for ... yield → neue Collection
val xs = for i <- 1 to 3 yield i * 2  // Typ: IndexedSeq[Int]
```

### Lösung 3.3
**Tail-Rekursion:** Der rekursive Aufruf ist die **letzte** Operation der Funktion.

**Wichtigkeit:** Der Compiler kann die Rekursion in eine Schleife umwandeln → konstanter Stack-Verbrauch, kein StackOverflowError.

**`@tailrec`:** Annotation, die den Compiler prüfen lässt, ob die Funktion wirklich tail-rekursiv ist. Gibt Compile-Fehler, wenn nicht.

### Lösung 3.4
```scala
def f1(x: Int) = x * 2       // Rückgabetyp inferiert (Int)
def f2(x: Int): Int = x * 2  // Rückgabetyp explizit
```

Funktional identisch! Für öffentliche APIs ist explizite Typisierung empfohlen.

### Lösung 3.5
`while`-Schleifen sollten vermieden werden weil:
1. Sie erfordern `var` (veränderliche Variable)
2. Sie sind nicht funktional (Seiteneffekt-basiert)
3. Alternativen sind ausdrucksstärker:
   - Rekursion
   - `for`-Comprehensions
   - Collection-Methoden wie `map`, `filter`, `fold`

---

## Teil B: REPL-Übungen

### Lösung 3.1
```scala
scala> val zahl = 42
val zahl: Int = 42

scala> val vorzeichen = if zahl > 0 then "positiv" else if zahl < 0 then "negativ" else "null"
val vorzeichen: String = positiv
```

### Lösung 3.2
```scala
scala> for i <- 1 to 10 yield i * i
val res0: IndexedSeq[Int] = Vector(1, 4, 9, 16, 25, 36, 49, 64, 81, 100)

scala> for i <- 1 to 10 if i % 2 == 0 yield i * i
val res1: IndexedSeq[Int] = Vector(4, 16, 36, 64, 100)

scala> for a <- 1 to 3; b <- 1 to 3 yield (a, b)
val res2: IndexedSeq[(Int, Int)] = Vector((1,1), (1,2), (1,3), (2,1), (2,2), (2,3), (3,1), (3,2), (3,3))
```

### Lösung 3.3
```scala
def greet(name: String): String = s"Hello, $name!"
def add(a: Int, b: Int): Int = a + b
def isEven(n: Int): Boolean = n % 2 == 0
def absoluteValue(n: Int): Int = if n >= 0 then n else -n
```

---

## Teil C: Programmieraufgaben

### Lösung 3.4 – Notenberechnung
```scala
def berechneNote(punkte: Int): String =
  if punkte >= 90 then "Sehr gut (1)"
  else if punkte >= 80 then "Gut (2)"
  else if punkte >= 70 then "Befriedigend (3)"
  else if punkte >= 60 then "Ausreichend (4)"
  else "Nicht bestanden (5)"

@main def notenTest(): Unit =
  println(berechneNote(95))  // "Sehr gut (1)"
  println(berechneNote(75))  // "Befriedigend (3)"
  println(berechneNote(45))  // "Nicht bestanden (5)"
```

### Lösung 3.5 – FizzBuzz
```scala
def fizzBuzz(n: Int): List[String] =
  (for i <- 1 to n yield
    if i % 15 == 0 then "FizzBuzz"
    else if i % 3 == 0 then "Fizz"
    else if i % 5 == 0 then "Buzz"
    else i.toString
  ).toList

@main def fizzBuzzTest(): Unit =
  println(fizzBuzz(15).mkString(", "))
```

### Lösung 3.6 – Rekursive Summe
```scala
def summe(n: Int): Int =
  if n <= 0 then 0
  else n + summe(n - 1)

@main def summeTest(): Unit =
  println(summe(5))   // 15
  println(summe(10))  // 55
```

### Lösung 3.7 – Tail-rekursive Summe
```scala
import scala.annotation.tailrec

def summeTail(n: Int): Int =
  @tailrec
  def loop(current: Int, akkumulator: Int): Int =
    if current <= 0 then akkumulator
    else loop(current - 1, akkumulator + current)
  
  loop(n, 0)

@main def summeTailTest(): Unit =
  println(summeTail(5))      // 15
  println(summeTail(10))     // 55
  println(summeTail(100000)) // 5000050000 (ohne StackOverflow!)
```

### Lösung 3.8 – Primzahlen
```scala
def istPrim(n: Int): Boolean =
  if n <= 1 then false
  else if n == 2 then true
  else !(2 to math.sqrt(n).toInt).exists(t => n % t == 0)

def primzahlenBis(n: Int): List[Int] =
  (2 to n).filter(istPrim).toList

@main def primTest(): Unit =
  println(istPrim(7))        // true
  println(istPrim(10))       // false
  println(primzahlenBis(20)) // List(2, 3, 5, 7, 11, 13, 17, 19)
```

### Lösung 3.9 – Collatz-Folge
```scala
def collatz(n: Int): List[Int] =
  @tailrec
  def loop(current: Int, acc: List[Int]): List[Int] =
    if current == 1 then (acc :+ 1)
    else if current % 2 == 0 then loop(current / 2, acc :+ current)
    else loop(3 * current + 1, acc :+ current)
  
  loop(n, List.empty)

@main def collatzTest(): Unit =
  println(collatz(6))   // List(6, 3, 10, 5, 16, 8, 4, 2, 1)
  println(collatz(27))  // Lange Folge!
```

---

## Teil D: Verständnisfragen

### Lösung 3.6
```scala
val xs = for i <- 1 to 5 if i != 3 yield i * 10
// xs = Vector(10, 20, 40, 50)
```

Die 3 wird durch den Filter `if i != 3` ausgeschlossen.

### Lösung 3.7
```scala
def length(list: List[Int]): Int =
  if list.isEmpty then 0
  else 1 + length(list.tail)
```

**NEIN**, diese Funktion ist **nicht tail-rekursiv!**

Nach dem rekursiven Aufruf `length(list.tail)` muss noch `1 + ...` berechnet werden. Die letzte Operation ist die Addition, nicht der rekursive Aufruf.

**Tail-rekursive Version:**
```scala
@tailrec
def lengthTail(list: List[Int], acc: Int = 0): Int =
  if list.isEmpty then acc
  else lengthTail(list.tail, acc + 1)
```

### Lösung 3.8
```scala
def mystery(n: Int): Int =
  if n == 0 then 0
  else n + mystery(n - 1)

mystery(4) // = ?
```

Das ist die Summe von 1 bis n!

```
mystery(4) = 4 + mystery(3)
           = 4 + (3 + mystery(2))
           = 4 + (3 + (2 + mystery(1)))
           = 4 + (3 + (2 + (1 + mystery(0))))
           = 4 + (3 + (2 + (1 + 0)))
           = 4 + 3 + 2 + 1 + 0
           = 10
```

---

## Bonusaufgaben

### Bonus 1 – Pascal'sches Dreieck
```scala
def pascal(spalte: Int, zeile: Int): Int =
  if spalte == 0 || spalte == zeile then 1
  else pascal(spalte - 1, zeile - 1) + pascal(spalte, zeile - 1)

@main def pascalTest(): Unit =
  println(pascal(0, 0))  // 1
  println(pascal(1, 2))  // 2
  println(pascal(2, 4))  // 6
  
  // Dreieck ausgeben
  for zeile <- 0 to 5 do
    print(" " * (5 - zeile))
    for spalte <- 0 to zeile do
      print(s"${pascal(spalte, zeile)} ")
    println()
```

### Bonus 2 – GGT
```scala
import scala.annotation.tailrec

@tailrec
def ggt(a: Int, b: Int): Int =
  if b == 0 then a
  else ggt(b, a % b)

@main def ggtTest(): Unit =
  println(ggt(48, 18))   // 6
  println(ggt(100, 35))  // 5
```

### Bonus 3 – Permutationen
```scala
def permutationen[A](liste: List[A]): List[List[A]] =
  liste match
    case Nil => List(Nil)
    case _ =>
      for
        elem <- liste
        rest <- permutationen(liste.filterNot(_ == elem))
      yield elem :: rest

@main def permTest(): Unit =
  println(permutationen(List(1, 2, 3)))
  // List(List(1, 2, 3), List(1, 3, 2), List(2, 1, 3), 
  //      List(2, 3, 1), List(3, 1, 2), List(3, 2, 1))
```
