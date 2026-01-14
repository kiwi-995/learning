# Vorlesung 9: Fortgeschrittene FP-Konzepte

## Lernziele

Nach dieser Vorlesung kannst du:
- Option, Either, Try richtig einsetzen
- For-Comprehensions für Monad-Komposition nutzen
- Fehler funktional behandeln
- Die Monad-Gesetze verstehen

---

## 9.1 Option – Optionale Werte

`Option[A]` repräsentiert einen Wert, der vorhanden (`Some(a)`) oder abwesend (`None`) sein kann.

```scala
val vielleicht: Option[Int] = Some(42)
val nichts: Option[Int] = None
```

### Warum Option?

```scala
// Schlecht: null
def findePerson(id: Int): Person = 
  if gefunden then person else null  // 💀 NullPointerException!

// Gut: Option
def findePerson(id: Int): Option[Person] =
  if gefunden then Some(person) else None
```

### Mit Option arbeiten

```scala
val opt: Option[Int] = Some(42)

// Pattern Matching
opt match
  case Some(n) => println(s"Wert: $n")
  case None => println("Kein Wert")

// Methoden
opt.getOrElse(0)     // 42 (oder 0 bei None)
opt.map(_ * 2)       // Some(84)
opt.filter(_ > 50)   // None
opt.flatMap(n => Some(n + 1))  // Some(43)
opt.isDefined        // true
opt.isEmpty          // false
opt.foreach(println) // Gibt 42 aus
```

---

## 9.2 Either – Zwei Möglichkeiten

`Either[L, R]` ist entweder `Left(l)` (oft Fehler) oder `Right(r)` (oft Erfolg).

```scala
def teilen(a: Int, b: Int): Either[String, Int] =
  if b == 0 then Left("Division durch Null")
  else Right(a / b)

teilen(10, 2)  // Right(5)
teilen(10, 0)  // Left("Division durch Null")
```

### Mit Either arbeiten

```scala
val result: Either[String, Int] = Right(42)

result match
  case Right(n) => println(s"Erfolg: $n")
  case Left(err) => println(s"Fehler: $err")

// Either ist rechts-biased
result.map(_ * 2)       // Right(84)
result.flatMap(n => Right(n + 1))  // Right(43)
result.getOrElse(0)     // 42

// Links-Seite transformieren
result.left.map(_.toUpperCase)
```

### Validierung mit Either

```scala
case class User(name: String, email: String, alter: Int)

def validiereNamen(name: String): Either[String, String] =
  if name.length >= 2 then Right(name)
  else Left("Name zu kurz")

def validiereEmail(email: String): Either[String, String] =
  if email.contains("@") then Right(email)
  else Left("Ungültige Email")

def validiereAlter(alter: Int): Either[String, Int] =
  if alter >= 0 && alter <= 150 then Right(alter)
  else Left("Ungültiges Alter")

def validiereUser(name: String, email: String, alter: Int): Either[String, User] =
  for
    n <- validiereNamen(name)
    e <- validiereEmail(email)
    a <- validiereAlter(alter)
  yield User(n, e, a)
```

---

## 9.3 Try – Ausnahmen einfangen

`Try[A]` fängt Exceptions und wandelt sie in Werte um:

```scala
import scala.util.{Try, Success, Failure}

def parseZahl(s: String): Try[Int] = Try(s.toInt)

parseZahl("42")   // Success(42)
parseZahl("abc")  // Failure(NumberFormatException)
```

### Mit Try arbeiten

```scala
val result: Try[Int] = Try("42".toInt)

result match
  case Success(n) => println(s"Zahl: $n")
  case Failure(e) => println(s"Fehler: ${e.getMessage}")

result.getOrElse(0)
result.map(_ * 2)
result.flatMap(n => Try(100 / n))
result.toOption     // Some(42)
result.toEither     // Right(42)
```

### Recover

```scala
val unsicher = Try(10 / 0)

unsicher.recover {
  case _: ArithmeticException => 0
}  // Success(0)

unsicher.recoverWith {
  case _: ArithmeticException => Success(0)
}  // Success(0)
```

---

## 9.4 For-Comprehensions ⭐

For-Comprehensions sind **syntaktischer Zucker** für `flatMap` und `map`:

```scala
// Mit For-Comprehension
for
  a <- Some(3)
  b <- Some(4)
yield a + b  // Some(7)

// Entspricht:
Some(3).flatMap(a => Some(4).map(b => a + b))
```

### Mehrere Werte verketten

```scala
def berechne(x: String, y: String): Option[Int] =
  for
    a <- x.toIntOption
    b <- y.toIntOption
    if b != 0
    ergebnis = a / b
  yield ergebnis

berechne("10", "2")  // Some(5)
berechne("10", "0")  // None (Guard schlägt fehl)
berechne("ab", "2")  // None (Parsing schlägt fehl)
```

### Mit Either

```scala
def workflow(input: String): Either[String, Int] =
  for
    parsed <- Try(input.toInt).toEither.left.map(_.getMessage)
    validated <- if parsed > 0 then Right(parsed) else Left("Muss positiv sein")
    result = validated * 2
  yield result

workflow("5")   // Right(10)
workflow("-5")  // Left("Muss positiv sein")
workflow("abc") // Left("For input string: \"abc\"")
```

---

## 9.5 Monad-Konzept

Ein **Monad** ist ein Typ `M[A]` mit:
- `pure(a)` / `apply(a)`: Wert einpacken
- `flatMap(f: A => M[B])`: Verketten

### Monad-Gesetze

```scala
// 1. Linksidentität: pure(a).flatMap(f) == f(a)
Some(5).flatMap(x => Some(x * 2)) == Some(10)  // ✓

// 2. Rechtsidentität: m.flatMap(pure) == m
Some(5).flatMap(Some(_)) == Some(5)  // ✓

// 3. Assoziativität: m.flatMap(f).flatMap(g) == m.flatMap(a => f(a).flatMap(g))
```

### Option ist ein Monad

```scala
val opt: Option[Int] = Some(5)

// pure
def pure[A](a: A): Option[A] = Some(a)

// flatMap
opt.flatMap(x => Some(x * 2))  // Some(10)
```

---

## 9.6 Kombinatoren

### sequence – Liste von Options zu Option von Liste

```scala
def sequence[A](opts: List[Option[A]]): Option[List[A]] =
  opts.foldRight(Some(Nil): Option[List[A]]) { (opt, acc) =>
    for
      a <- opt
      list <- acc
    yield a :: list
  }

sequence(List(Some(1), Some(2), Some(3)))  // Some(List(1, 2, 3))
sequence(List(Some(1), None, Some(3)))     // None
```

### traverse – map + sequence

```scala
def traverse[A, B](list: List[A])(f: A => Option[B]): Option[List[B]] =
  sequence(list.map(f))

traverse(List("1", "2", "3"))(_.toIntOption)  // Some(List(1, 2, 3))
traverse(List("1", "x", "3"))(_.toIntOption)  // None
```

---

## 9.7 Praktische Muster

### Railway-Oriented Programming

```scala
def schritt1(input: String): Either[String, Int] =
  input.toIntOption.toRight("Keine Zahl")

def schritt2(n: Int): Either[String, Int] =
  if n > 0 then Right(n) else Left("Muss positiv sein")

def schritt3(n: Int): Either[String, String] =
  Right(s"Ergebnis: ${n * 2}")

def pipeline(input: String): Either[String, String] =
  for
    a <- schritt1(input)
    b <- schritt2(a)
    c <- schritt3(b)
  yield c
```

### Fehlerakkumulation (kurzer Einblick)

```scala
// Mit Cats/ZIO kann man alle Fehler sammeln statt beim ersten abzubrechen
// Dies geht über den Scope dieses Kurses hinaus
```

---

## Zusammenfassung

| Typ | Bedeutung | Erfolg | Fehler |
|-----|-----------|--------|--------|
| `Option[A]` | Optionaler Wert | `Some(a)` | `None` |
| `Either[L, R]` | Zwei Möglichkeiten | `Right(r)` | `Left(l)` |
| `Try[A]` | Ausnahmefrei | `Success(a)` | `Failure(e)` |

**For-Comprehension:**
```scala
for
  a <- optionA
  b <- optionB if bedingung
  c = berechnung
yield ergebnis
```

---

## Nächste Schritte

In **Vorlesung 10** lernen wir:
- Git Grundlagen
- Repositories, Commits, Branches
- Merge und Rebase

➡️ [Weiter zu Vorlesung 10: Git Grundlagen](./10-git-grundlagen.md)
