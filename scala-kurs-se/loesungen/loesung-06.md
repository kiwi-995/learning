# Lösungen zu Übung 6

---

## Teil A: Theoretische Fragen

### Lösung 6.1
**foldLeft:** Verarbeitet von links nach rechts: `((start op a1) op a2) op a3`
**foldRight:** Verarbeitet von rechts nach links: `a1 op (a2 op (a3 op start))`

Der Unterschied ist wichtig bei:
- Nicht-assoziativen Operationen (Subtraktion, Division)
- Wenn die Reihenfolge relevant ist (z.B. String-Verkettung)
- Bei manchen rekursiven Strukturen

```scala
List(1, 2, 3).foldLeft(10)(_ - _)   // ((10-1)-2)-3 = 4
List(1, 2, 3).foldRight(10)(_ - _)  // 1-(2-(3-10)) = -8
```

### Lösung 6.2
**map:** Transformiert jedes Element 1:1 → Ergebnis hat gleiche Länge
**flatMap:** Transformiert jedes Element zu einer Collection und "flacht" das Ergebnis

```scala
List(1, 2, 3).map(x => List(x, x))     // List(List(1,1), List(2,2), List(3,3))
List(1, 2, 3).flatMap(x => List(x, x)) // List(1, 1, 2, 2, 3, 3)
```

Verwende `flatMap` wenn die Funktion eine Collection zurückgibt und du ein flaches Ergebnis willst.

### Lösung 6.3
List ist eine **verkettete Liste**:
- `0 :: liste`: Neuen Knoten vorne anhängen, zeigt auf alte Liste → O(1)
- `liste :+ 0`: Muss gesamte Liste durchlaufen um am Ende anzuhängen → O(n)

### Lösung 6.4
**Immutable:** Kann nicht geändert werden, Operationen erzeugen neue Collections
**Mutable:** Kann in-place geändert werden

Vorteile von immutable:
- Thread-sicher (keine Race Conditions)
- Einfacher zu verstehen (keine versteckten Änderungen)
- Funktionaler Stil
- Structural Sharing ermöglicht Effizienz

---

## Teil B: REPL-Übungen

### Lösung 6.2
```scala
val worte = List("Scala", "ist", "toll")

// 1. Mit Leerzeichen verbinden
worte.mkString(" ")  // "Scala ist toll"

// 2. Alle Großbuchstaben
worte.map(_.toUpperCase)  // List("SCALA", "IST", "TOLL")

// 3. Nur lange Worte
worte.filter(_.length > 3)  // List("Scala", "toll")
```

### Lösung 6.3
```scala
val studentennoten = Map(
  "Alice" -> 1.3,
  "Bob" -> 2.7,
  "Charlie" -> 1.0,
  "Diana" -> 2.0
)

// 1. Alices Note
studentennoten("Alice")  // 1.3

// 2. Studenten mit Note besser als 2.0
studentennoten.filter((_, note) => note < 2.0)
// Map("Alice" -> 1.3, "Charlie" -> 1.0)

// 3. Durchschnittsnote
studentennoten.values.sum / studentennoten.size  // 1.75

// 4. Eve hinzufügen
studentennoten + ("Eve" -> 1.7)
```

---

## Teil C: Programmieraufgaben

### Lösung 6.4 – Statistik
```scala
object Statistik:
  def mittelwert(zahlen: List[Double]): Double =
    if zahlen.isEmpty then 0.0
    else zahlen.sum / zahlen.length
  
  def varianz(zahlen: List[Double]): Double =
    if zahlen.isEmpty then 0.0
    else
      val m = mittelwert(zahlen)
      zahlen.map(x => (x - m) * (x - m)).sum / zahlen.length
  
  def standardabweichung(zahlen: List[Double]): Double =
    math.sqrt(varianz(zahlen))
  
  def median(zahlen: List[Double]): Double =
    if zahlen.isEmpty then 0.0
    else
      val sorted = zahlen.sorted
      val n = sorted.length
      if n % 2 == 1 then sorted(n / 2)
      else (sorted(n / 2 - 1) + sorted(n / 2)) / 2
  
  def modus(zahlen: List[Double]): Double =
    if zahlen.isEmpty then 0.0
    else zahlen.groupBy(identity).maxBy(_._2.length)._1

// Test
val daten = List(1.0, 2.0, 2.0, 3.0, 4.0, 4.0, 4.0, 5.0)
println(s"Mittelwert: ${Statistik.mittelwert(daten)}")  // 3.125
println(s"Median: ${Statistik.median(daten)}")          // 3.5
println(s"Modus: ${Statistik.modus(daten)}")            // 4.0
```

### Lösung 6.5 – Wortzähler
```scala
def zähleWorte(text: String): Map[String, Int] =
  text
    .toLowerCase
    .split("\\s+")
    .groupBy(identity)
    .map((wort, vorkommen) => wort -> vorkommen.length)

val text = "Scala ist toll Scala ist funktional Scala ist objektorientiert"
println(zähleWorte(text))
// Map(toll -> 1, funktional -> 1, scala -> 3, ist -> 3, objektorientiert -> 1)
```

### Lösung 6.6 – Matrixoperationen
```scala
type Matrix = List[List[Double]]

object MatrixOps:
  def transponieren(m: Matrix): Matrix =
    if m.isEmpty || m.head.isEmpty then Nil
    else m.map(_.head) :: transponieren(m.map(_.tail))
  
  def addieren(a: Matrix, b: Matrix): Matrix =
    a.zip(b).map((zeileA, zeileB) => 
      zeileA.zip(zeileB).map(_ + _))
  
  def skalar(m: Matrix, s: Double): Matrix =
    m.map(_.map(_ * s))
  
  def multiplizieren(a: Matrix, b: Matrix): Matrix =
    val bt = transponieren(b)
    a.map(zeileA => 
      bt.map(spalteB => 
        zeileA.zip(spalteB).map(_ * _).sum))
  
  def alsString(m: Matrix): String =
    m.map(zeile => zeile.map(x => f"$x%6.2f").mkString(" "))
      .mkString("\n")

// Test
val m1 = List(List(1.0, 2.0), List(3.0, 4.0))
val m2 = List(List(5.0, 6.0), List(7.0, 8.0))
println(MatrixOps.alsString(MatrixOps.multiplizieren(m1, m2)))
```

### Lösung 6.7 – Primfaktorzerlegung
```scala
def primfaktoren(n: Int): List[Int] =
  def loop(zahl: Int, faktor: Int, acc: List[Int]): List[Int] =
    if zahl <= 1 then acc.reverse
    else if faktor * faktor > zahl then (zahl :: acc).reverse
    else if zahl % faktor == 0 then loop(zahl / faktor, faktor, faktor :: acc)
    else loop(zahl, faktor + 1, acc)
  
  loop(n, 2, Nil)

println(primfaktoren(12))   // List(2, 2, 3)
println(primfaktoren(100))  // List(2, 2, 5, 5)
println(primfaktoren(17))   // List(17)
```

### Lösung 6.8 – Gruppierung
```scala
case class Student(name: String, fach: String, semester: Int)

val studenten = List(
  Student("Alice", "Informatik", 3),
  Student("Bob", "Mathematik", 2),
  Student("Charlie", "Informatik", 2),
  Student("Diana", "Physik", 1),
  Student("Eve", "Informatik", 3)
)

// 1. Nach Fach gruppieren
val nachFach = studenten.groupBy(_.fach)

// 2. Nach Semester gruppieren
val nachSemester = studenten.groupBy(_.semester)

// 3. Informatik im 3. Semester
val info3 = studenten.filter(s => s.fach == "Informatik" && s.semester == 3)
// List(Student(Alice,Informatik,3), Student(Eve,Informatik,3))

// 4. Durchschnittssemester pro Fach
val durchschnitt = nachFach.map { (fach, studis) =>
  fach -> studis.map(_.semester).sum.toDouble / studis.length
}
// Map(Informatik -> 2.67, Mathematik -> 2.0, Physik -> 1.0)
```

---

## Teil D: For-Comprehensions

### Lösung 6.9
```scala
// 1. Als For-Comprehension:
for
  x <- List(1, 2, 3)
  y <- List(4, 5, 6)
yield x * y

// 2. Als map/flatMap/filter:
List(1, 2, 3).flatMap(x => 
  List(1, 2, 3).filter(y => x != y).map(y => (x, y)))
```

### Lösung 6.10 – Schachbrett
```scala
val positionen = for
  spalte <- 'a' to 'h'
  zeile <- 1 to 8
yield (spalte, zeile)

positionen.toList
// List((a,1), (a,2), ..., (h,8))
```

### Lösung 6.11 – Pythagoräische Tripel
```scala
val tripel = for
  a <- 1 to 100
  b <- a to 100  // b >= a vermeidet Duplikate
  c <- b to 100  // c >= b
  if a * a + b * b == c * c
yield (a, b, c)

tripel.toList
// List((3,4,5), (5,12,13), (6,8,10), (7,24,25), (8,15,17), ...)
```

---

## Bonusaufgaben

### Bonus 1 – Eigene Collection-Methoden
```scala
def meinMap[A, B](liste: List[A])(f: A => B): List[B] = liste match
  case Nil => Nil
  case kopf :: rest => f(kopf) :: meinMap(rest)(f)

def meinFilter[A](liste: List[A])(p: A => Boolean): List[A] = liste match
  case Nil => Nil
  case kopf :: rest =>
    if p(kopf) then kopf :: meinFilter(rest)(p)
    else meinFilter(rest)(p)

def meinFold[A, B](liste: List[A])(start: B)(f: (B, A) => B): B = liste match
  case Nil => start
  case kopf :: rest => meinFold(rest)(f(start, kopf))(f)

def meinFlatMap[A, B](liste: List[A])(f: A => List[B]): List[B] = liste match
  case Nil => Nil
  case kopf :: rest => f(kopf) ++ meinFlatMap(rest)(f)

// Test
println(meinMap(List(1, 2, 3))(_ * 2))            // List(2, 4, 6)
println(meinFilter(List(1, 2, 3, 4))(_ % 2 == 0)) // List(2, 4)
println(meinFold(List(1, 2, 3))(0)(_ + _))        // 6
println(meinFlatMap(List(1, 2))(x => List(x, x))) // List(1, 1, 2, 2)
```

### Bonus 2 – Sliding Window Maximum
```scala
def slidingMax(zahlen: List[Int], k: Int): List[Int] =
  zahlen.sliding(k).map(_.max).toList

// Test
println(slidingMax(List(1, 3, -1, -3, 5, 3, 6, 7), 3))
// List(3, 3, 5, 5, 6, 7)
```
