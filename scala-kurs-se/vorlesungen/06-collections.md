# Vorlesung 6: Collections

## Lernziele

Nach dieser Vorlesung kannst du:
- Die wichtigsten Collection-Typen unterscheiden
- map, filter, flatMap, fold anwenden
- For-Comprehensions auf Collections nutzen
- Immutable vs Mutable Collections verstehen
- Performance-Charakteristiken kennen

---

## 6.1 Übersicht: Scala Collections

```
┌─────────────────────────────────────────────────────────────┐
│                     Iterable                                │
├────────────────────┬────────────────────┬───────────────────┤
│        Seq         │        Set         │       Map         │
├────────────────────┼────────────────────┼───────────────────┤
│ List, Vector,      │ HashSet, TreeSet,  │ HashMap, TreeMap, │
│ Array, ArrayBuffer │ SortedSet          │ SortedMap         │
└────────────────────┴────────────────────┴───────────────────┘
```

### Immutable vs Mutable

```scala
import scala.collection.mutable

// Immutable (Standard, bevorzugt!)
val liste = List(1, 2, 3)
val set = Set("a", "b", "c")
val map = Map("x" -> 1, "y" -> 2)

// Mutable (für Performance-kritische Fälle)
val buffer = mutable.ArrayBuffer(1, 2, 3)
val mutableSet = mutable.Set("a", "b")
val mutableMap = mutable.Map("x" -> 1)
```

---

## 6.2 List

**List** ist die wichtigste Collection in funktionalem Scala:

```scala
val zahlen = List(1, 2, 3, 4, 5)
val leer = List.empty[Int]
val range = (1 to 10).toList
```

### Grundoperationen

```scala
val xs = List(1, 2, 3, 4, 5)

xs.head          // 1 (erstes Element)
xs.tail          // List(2, 3, 4, 5) (Rest)
xs.last          // 5 (letztes)
xs.init          // List(1, 2, 3, 4) (alle außer letztes)

xs.length        // 5
xs.isEmpty       // false
xs.nonEmpty      // true

xs(2)            // 3 (Index-Zugriff, langsam bei List!)
xs.take(3)       // List(1, 2, 3)
xs.drop(2)       // List(3, 4, 5)
xs.slice(1, 4)   // List(2, 3, 4)
```

### Elemente hinzufügen

```scala
val xs = List(2, 3, 4)

0 :: xs          // List(0, 2, 3, 4) - Prepend (effizient!)
xs :+ 5          // List(2, 3, 4, 5) - Append (langsam!)
xs ++ List(5, 6) // List(2, 3, 4, 5, 6) - Verketten
```

### Pattern Matching auf Listen

```scala
def summe(liste: List[Int]): Int = liste match
  case Nil => 0
  case kopf :: rest => kopf + summe(rest)

def beschreibe[A](liste: List[A]): String = liste match
  case Nil => "Leer"
  case x :: Nil => s"Ein Element: $x"
  case x :: y :: rest => s"Mehrere, startet mit $x, $y"
```

---

## 6.3 Vector

**Vector** bietet schnellen Random Access und ist gut für große Datenmengen:

```scala
val vec = Vector(1, 2, 3, 4, 5)

vec(2)           // O(1) - effektiv konstant!
vec :+ 6         // O(1) - Append effizient
0 +: vec         // O(1) - Prepend auch effizient
```

### List vs Vector

| Operation | List | Vector |
|-----------|------|--------|
| Prepend (::) | O(1) | O(log n) |
| Append (:+) | O(n) | O(log n) |
| Index-Zugriff | O(n) | O(log n) |
| Head/Tail | O(1) | O(log n) |

**Faustregel:** 
- List für rekursive Algorithmen
- Vector für Random Access

---

## 6.4 Set und Map

### Set

```scala
val set = Set(1, 2, 3, 2, 1)  // Set(1, 2, 3) - keine Duplikate!

set.contains(2)   // true
set + 4           // Set(1, 2, 3, 4)
set - 2           // Set(1, 3)
set ++ Set(4, 5)  // Set(1, 2, 3, 4, 5)
```

### Map

```scala
val alter = Map("Alice" -> 25, "Bob" -> 30)

alter("Alice")              // 25
alter.get("Charlie")        // None
alter.getOrElse("Eve", 0)   // 0

alter + ("Charlie" -> 35)   // Map mit Charlie
alter - "Bob"               // Map ohne Bob

alter.keys                  // Set(Alice, Bob)
alter.values                // Iterable(25, 30)
```

---

## 6.5 Transformationen: map, filter, flatMap

### map – Elemente transformieren

```scala
val zahlen = List(1, 2, 3, 4, 5)

zahlen.map(x => x * 2)        // List(2, 4, 6, 8, 10)
zahlen.map(_ * 2)             // Kurzschreibweise
zahlen.map(x => x.toString)   // List("1", "2", "3", "4", "5")
```

### filter – Elemente filtern

```scala
zahlen.filter(x => x % 2 == 0)   // List(2, 4)
zahlen.filter(_ > 3)             // List(4, 5)
zahlen.filterNot(_ > 3)          // List(1, 2, 3)
```

### flatMap – map + flatten

```scala
val worte = List("Hallo", "Welt")

worte.map(_.toList)      // List(List('H','a','l','l','o'), List('W','e','l','t'))
worte.flatMap(_.toList)  // List('H','a','l','l','o','W','e','l','t')

// Praktisches Beispiel
val optionen = List(Some(1), None, Some(2), None, Some(3))
optionen.flatMap(x => x)  // List(1, 2, 3)
```

### Verkettung

```scala
val ergebnis = (1 to 100)
  .filter(_ % 2 == 0)      // Gerade Zahlen
  .map(_ * 2)              // Verdoppeln
  .filter(_ < 50)          // Unter 50
  .toList

// List(4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48)
```

---

## 6.6 Reduce und Fold

### reduce – Zu einem Wert zusammenführen

```scala
val zahlen = List(1, 2, 3, 4, 5)

zahlen.reduce((a, b) => a + b)    // 15
zahlen.reduce(_ + _)              // 15 (Kurzschreibweise)
zahlen.reduce(_ * _)              // 120 (Produkt)
zahlen.reduce(_ max _)            // 5 (Maximum)
```

### fold – Mit Startwert

```scala
zahlen.fold(0)(_ + _)       // 15, startet mit 0
zahlen.fold(1)(_ * _)       // 120, startet mit 1
zahlen.fold(100)(_ - _)     // 100-1-2-3-4-5 = 85

List.empty[Int].fold(0)(_ + _)  // 0 (funktioniert bei leerer Liste!)
```

### foldLeft und foldRight

```scala
// foldLeft: von links nach rechts
List(1, 2, 3).foldLeft(0)((acc, x) => acc + x)  // ((0+1)+2)+3 = 6

// foldRight: von rechts nach links
List(1, 2, 3).foldRight(0)((x, acc) => x + acc)  // 1+(2+(3+0)) = 6

// Unterschied bei nicht-assoziativen Operationen:
List("a", "b", "c").foldLeft("")(_ + _)   // "abc"
List("a", "b", "c").foldRight("")(_ + _)  // "abc" (hier gleich)

List(1, 2, 3).foldLeft(List.empty[Int])((acc, x) => x :: acc)  
// List(3, 2, 1) - Umkehren!
```

---

## 6.7 Weitere nützliche Methoden

### Prüfungen

```scala
val xs = List(1, 2, 3, 4, 5)

xs.exists(_ > 3)       // true (mindestens einer > 3)
xs.forall(_ > 0)       // true (alle > 0)
xs.contains(3)         // true
xs.find(_ > 3)         // Some(4) (erstes passende)
xs.count(_ % 2 == 0)   // 2 (Anzahl gerader)
```

### Gruppieren und Partitionieren

```scala
val xs = List(1, 2, 3, 4, 5, 6)

xs.partition(_ % 2 == 0)        // (List(2,4,6), List(1,3,5))
xs.groupBy(_ % 3)               // Map(0 -> List(3,6), 1 -> List(1,4), 2 -> List(2,5))
xs.grouped(2).toList            // List(List(1,2), List(3,4), List(5,6))
xs.sliding(3).toList            // List(List(1,2,3), List(2,3,4), List(3,4,5), ...)
```

### Sortieren

```scala
val xs = List(3, 1, 4, 1, 5, 9, 2, 6)

xs.sorted                       // List(1, 1, 2, 3, 4, 5, 6, 9)
xs.sortBy(x => -x)              // Absteigend
xs.sortWith((a, b) => a > b)    // Absteigend
```

### Kombinieren

```scala
val a = List(1, 2, 3)
val b = List("a", "b", "c")

a.zip(b)                        // List((1,"a"), (2,"b"), (3,"c"))
a.zipWithIndex                  // List((1,0), (2,1), (3,2))
List((1,"a"), (2,"b")).unzip    // (List(1, 2), List("a", "b"))
```

---

## 6.8 For-Comprehensions auf Collections

For-Comprehensions sind syntaktischer Zucker für map/flatMap/filter:

```scala
// Einfaches map
for x <- List(1, 2, 3) yield x * 2
// = List(1, 2, 3).map(x => x * 2)

// Mit Filter
for x <- List(1, 2, 3, 4, 5) if x % 2 == 0 yield x * 2
// = List(...).filter(x => x % 2 == 0).map(x => x * 2)

// Verschachtelt (flatMap)
for
  x <- List(1, 2, 3)
  y <- List("a", "b")
yield (x, y)
// = List((1,"a"), (1,"b"), (2,"a"), (2,"b"), (3,"a"), (3,"b"))
```

### Äquivalenz

```scala
// For-Comprehension
for
  person <- personen
  if person.alter >= 18
  freund <- person.freunde
yield freund.name

// Entspricht:
personen
  .filter(_.alter >= 18)
  .flatMap(_.freunde)
  .map(_.name)
```

---

## 6.9 Lazy Collections

### LazyList (früher: Stream)

```scala
// Elemente werden erst bei Bedarf berechnet
val naturlicheZahlen = LazyList.from(1)
naturlicheZahlen.take(5).toList  // List(1, 2, 3, 4, 5)

// Fibonacci-Folge (unendlich!)
lazy val fibs: LazyList[BigInt] = 
  BigInt(0) #:: BigInt(1) #:: fibs.zip(fibs.tail).map(_ + _)

fibs.take(10).toList  // List(0, 1, 1, 2, 3, 5, 8, 13, 21, 34)
```

### View

```scala
// Lazy Transformationen
val ergebnis = (1 to 1000000)
  .view
  .map(_ * 2)
  .filter(_ % 3 == 0)
  .take(10)
  .toList  // Nur die benötigten werden berechnet
```

---

## Zusammenfassung

| Methode | Beschreibung | Beispiel |
|---------|--------------|----------|
| `map` | Transformiert jedes Element | `xs.map(_ * 2)` |
| `filter` | Behält passende Elemente | `xs.filter(_ > 0)` |
| `flatMap` | map + flatten | `xs.flatMap(f)` |
| `fold` | Reduziert mit Startwert | `xs.fold(0)(_ + _)` |
| `reduce` | Reduziert ohne Startwert | `xs.reduce(_ + _)` |
| `exists` | Existiert mindestens einer? | `xs.exists(_ > 5)` |
| `forall` | Gilt für alle? | `xs.forall(_ > 0)` |
| `find` | Erstes passendes Element | `xs.find(_ > 5)` |
| `groupBy` | Gruppiert nach Kriterium | `xs.groupBy(_ % 2)` |
| `zip` | Paart zwei Collections | `a.zip(b)` |

---

## Nächste Schritte

In **Vorlesung 7** lernen wir:
- Pattern Matching im Detail
- Algebraische Datentypen
- Exhaustive Matching

➡️ [Weiter zu Vorlesung 7: Pattern Matching & ADTs](./07-pattern-matching.md)
