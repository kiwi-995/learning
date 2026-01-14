# Vorlesung 8: Higher-Order Functions

## Lernziele

Nach dieser Vorlesung kannst du:
- Funktionen als Parameter übergeben
- Funktionen als Rückgabewerte verwenden
- Currying und Partial Application anwenden
- Closures verstehen
- Funktionstypen lesen und schreiben

---

## 8.1 Funktionen als Werte

In Scala sind Funktionen **First-Class Citizens**:

```scala
// Funktion als Wert speichern
val verdopple: Int => Int = (x: Int) => x * 2

// Funktion aufrufen
verdopple(5)  // 10

// Kurzschreibweise
val verdopple2 = (x: Int) => x * 2
```

### Funktionstypen

```scala
val f1: Int => Int = x => x + 1              // Ein Parameter
val f2: (Int, Int) => Int = (a, b) => a + b  // Zwei Parameter
val f3: () => String = () => "Hallo"         // Keine Parameter
val f4: Int => Int => Int = a => b => a + b  // Curried
```

---

## 8.2 Funktionen als Parameter

```scala
def anwenden(f: Int => Int, wert: Int): Int = f(wert)

val quadrat = (x: Int) => x * x
val inkrement = (x: Int) => x + 1

anwenden(quadrat, 5)     // 25
anwenden(inkrement, 5)   // 6
anwenden(_ * 3, 5)       // 15 (anonyme Funktion)
```

### Praktisches Beispiel

```scala
def filterUndTransformiere(
  zahlen: List[Int],
  filter: Int => Boolean,
  transform: Int => Int
): List[Int] =
  zahlen.filter(filter).map(transform)

val ergebnis = filterUndTransformiere(
  List(1, 2, 3, 4, 5, 6),
  _ % 2 == 0,    // Gerade Zahlen
  _ * 10         // Mal 10
)
// List(20, 40, 60)
```

---

## 8.3 Funktionen als Rückgabewerte

```scala
def multiplizierer(faktor: Int): Int => Int =
  (x: Int) => x * faktor

val mal2 = multiplizierer(2)
val mal5 = multiplizierer(5)

mal2(10)  // 20
mal5(10)  // 50
```

### Factory für Funktionen

```scala
def createValidator(min: Int, max: Int): String => Boolean =
  (s: String) => s.length >= min && s.length <= max

val passwortValidator = createValidator(8, 20)
val usernameValidator = createValidator(3, 15)

passwortValidator("geheim123")  // true
usernameValidator("ab")         // false
```

---

## 8.4 Closures

Eine **Closure** ist eine Funktion, die auf Variablen aus ihrem Umgebungskontext zugreift:

```scala
var faktor = 2
val multipliziere = (x: Int) => x * faktor  // Closure!

multipliziere(5)  // 10

faktor = 3
multipliziere(5)  // 15 (faktor hat sich geändert!)
```

### Closure mit val (sicherer)

```scala
def createCounter(): () => Int =
  var count = 0
  () => { count += 1; count }

val counter1 = createCounter()
val counter2 = createCounter()

counter1()  // 1
counter1()  // 2
counter2()  // 1 (eigener Zähler)
```

---

## 8.5 Currying

**Currying** wandelt eine Funktion mit mehreren Parametern in eine Kette von Funktionen um:

```scala
// Normale Funktion
def addiere(a: Int, b: Int): Int = a + b
addiere(2, 3)  // 5

// Curried Version
def addiereCurried(a: Int)(b: Int): Int = a + b
addiereCurried(2)(3)  // 5

// Oder
val add2 = addiereCurried(2)  // Int => Int
add2(3)  // 5
```

### Automatisches Currying

```scala
val normale = (a: Int, b: Int, c: Int) => a + b + c
val curried = normale.curried  // Int => Int => Int => Int

curried(1)(2)(3)  // 6

val umgekehrt = Function.uncurried(curried)  // (Int, Int, Int) => Int
```

---

## 8.6 Partial Application

**Partial Application** fixiert einige Parameter einer Funktion:

```scala
def gruss(gruß: String, name: String): String =
  s"$gruß, $name!"

// Teilweise angewendet
val halloSagen = gruss("Hallo", _)  // String => String
val tschueSagen = gruss("Tschüss", _)

halloSagen("Max")    // "Hallo, Max!"
tschueSagen("Max")   // "Tschüss, Max!"
```

### Mit Currying

```scala
def log(level: String)(message: String): Unit =
  println(s"[$level] $message")

val info = log("INFO")
val error = log("ERROR")

info("Programm gestartet")
error("Etwas ging schief")
```

---

## 8.7 Wichtige Higher-Order Functions

### map

```scala
List(1, 2, 3).map(_ * 2)           // List(2, 4, 6)
List("a", "b").map(_.toUpperCase)  // List("A", "B")
```

### filter

```scala
List(1, 2, 3, 4, 5).filter(_ > 2)  // List(3, 4, 5)
```

### fold / reduce

```scala
List(1, 2, 3).fold(0)(_ + _)     // 6
List(1, 2, 3).reduce(_ * _)     // 6
```

### flatMap

```scala
List(1, 2, 3).flatMap(x => List(x, x * 10))
// List(1, 10, 2, 20, 3, 30)
```

---

## 8.8 Eigene Higher-Order Functions

### Retry-Funktion

```scala
def retry[A](times: Int)(operation: => A): Option[A] =
  if times <= 0 then None
  else
    try Some(operation)
    catch case _: Exception => retry(times - 1)(operation)

// Verwendung
val ergebnis = retry(3) {
  // Risikoreiche Operation
  if math.random() < 0.5 then throw new Exception("Fehlgeschlagen")
  else "Erfolg!"
}
```

### Timing-Funktion

```scala
def zeit[A](operation: => A): (A, Long) =
  val start = System.currentTimeMillis()
  val result = operation
  val ende = System.currentTimeMillis()
  (result, ende - start)

val (ergebnis, ms) = zeit {
  Thread.sleep(100)
  42
}
println(s"Ergebnis: $ergebnis in $ms ms")
```

### Using-Pattern (Resource Management)

```scala
def using[R <: AutoCloseable, A](resource: R)(f: R => A): A =
  try f(resource)
  finally resource.close()

// Verwendung
val inhalt = using(new java.io.FileReader("test.txt")) { reader =>
  // Lesen...
  "Dateiinhalt"
}
```

---

## 8.9 Funktionskomposition

```scala
val f: Int => Int = _ + 1
val g: Int => Int = _ * 2

// andThen: erst f, dann g
val fDannG = f.andThen(g)  // (x + 1) * 2
fDannG(3)  // 8

// compose: erst g, dann f
val gDannF = f.compose(g)  // (x * 2) + 1
gDannF(3)  // 7
```

---

## 8.10 By-Name Parameter

```scala
def logIfDebug(message: => String): Unit =
  if debugMode then println(message)

// message wird nur ausgewertet, wenn debugMode = true!
logIfDebug(expensiveOperation())  // Lazy!
```

---

## Zusammenfassung

| Konzept | Syntax | Beispiel |
|---------|--------|----------|
| Funktionsliteral | `(x) => expr` | `(x: Int) => x * 2` |
| Funktionstyp | `A => B` | `Int => String` |
| HOF (Parameter) | `def f(g: A => B)` | `def map(f: Int => Int)` |
| HOF (Rückgabe) | `def f(): A => B` | `def multiplier(n: Int): Int => Int` |
| Currying | `def f(a)(b)` | `def add(a: Int)(b: Int)` |
| Closure | Freie Variablen | `val f = x => x + y` |
| By-Name | `param: => A` | `def log(msg: => String)` |

---

## Nächste Schritte

In **Vorlesung 9** lernen wir:
- Option, Either, Try im Detail
- For-Comprehensions als Monad-Komposition
- Fehlerbehandlung funktional

➡️ [Weiter zu Vorlesung 9: Fortgeschrittene FP-Konzepte](./09-fortgeschrittene-fp.md)
