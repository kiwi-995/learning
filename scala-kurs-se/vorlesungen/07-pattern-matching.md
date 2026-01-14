# Vorlesung 7: Pattern Matching & ADTs

## Lernziele

Nach dieser Vorlesung kannst du:
- Pattern Matching vollständig nutzen
- Algebraische Datentypen modellieren
- Guards und Extraktoren verwenden
- Exhaustive Matching verstehen
- Typische Muster erkennen

---

## 7.1 Grundlagen Pattern Matching

Pattern Matching ist eines der **mächtigsten Features** von Scala:

```scala
def beschreibe(x: Any): String = x match
  case 0 => "Null"
  case 1 => "Eins"
  case "Hallo" => "Eine Begrüßung"
  case true => "Wahr"
  case _ => "Etwas anderes"
```

### Match als Ausdruck

```scala
val ergebnis = zahl match
  case 1 => "eins"
  case 2 => "zwei"
  case _ => "viele"
```

---

## 7.2 Pattern-Arten

### Literal Patterns

```scala
x match
  case 0 => "Null"
  case "test" => "Test-String"
  case true => "Wahr"
```

### Variable Patterns

```scala
x match
  case n => s"Der Wert ist $n"  // Bindet n an x
```

### Wildcard Pattern

```scala
x match
  case _ => "Egal was"  // Fängt alles, bindet nicht
```

### Typmuster (Type Patterns)

```scala
def typBeschreibung(x: Any): String = x match
  case s: String => s"String der Länge ${s.length}"
  case i: Int => s"Integer: $i"
  case l: List[_] => s"Liste mit ${l.length} Elementen"
  case _ => "Unbekannter Typ"
```

### Tuple Patterns

```scala
val punkt = (3, 4)

punkt match
  case (0, 0) => "Ursprung"
  case (x, 0) => s"Auf X-Achse bei $x"
  case (0, y) => s"Auf Y-Achse bei $y"
  case (x, y) => s"Punkt ($x, $y)"
```

### List Patterns

```scala
liste match
  case Nil => "Leer"
  case x :: Nil => s"Ein Element: $x"
  case x :: y :: rest => s"Mindestens zwei: $x, $y, ..."
  case List(a, b, c) => s"Genau drei: $a, $b, $c"
```

---

## 7.3 Case Class Patterns (Dekonstruktion)

Case Classes können **dekonstruiert** werden:

```scala
case class Person(name: String, alter: Int)

val alice = Person("Alice", 25)

alice match
  case Person("Alice", age) => s"Alice ist $age Jahre alt"
  case Person(n, a) if a >= 18 => s"$n ist erwachsen"
  case Person(n, a) => s"$n ist minderjährig ($a Jahre)"
```

### Verschachtelte Patterns

```scala
case class Adresse(stadt: String, plz: String)
case class Mitarbeiter(name: String, adresse: Adresse)

val m = Mitarbeiter("Max", Adresse("Tübingen", "72074"))

m match
  case Mitarbeiter(n, Adresse("Tübingen", _)) => 
    s"$n arbeitet in Tübingen"
  case Mitarbeiter(n, Adresse(s, _)) => 
    s"$n arbeitet in $s"
```

---

## 7.4 Guards (Wächter)

Guards fügen Bedingungen zu Patterns hinzu:

```scala
zahl match
  case n if n < 0 => "Negativ"
  case n if n == 0 => "Null"
  case n if n > 0 && n < 10 => "Kleine positive Zahl"
  case n => "Große Zahl"
```

```scala
person match
  case Person(name, alter) if alter >= 18 && alter < 65 =>
    s"$name ist im arbeitsfähigen Alter"
  case Person(name, alter) if alter >= 65 =>
    s"$name ist in Rente"
  case Person(name, _) =>
    s"$name ist noch jung"
```

---

## 7.5 Algebraische Datentypen (ADTs) ⭐

ADTs sind **das Herzstück** funktionaler Programmierung!

### Sum Types (Oder-Typen)

```scala
sealed trait Ampelfarbe
case object Rot extends Ampelfarbe
case object Gelb extends Ampelfarbe
case object Grün extends Ampelfarbe
```

### Product Types (Und-Typen)

```scala
case class Position(x: Double, y: Double)
// Position HAT x UND y
```

### Kombiniert

```scala
sealed trait Form
case class Kreis(radius: Double) extends Form
case class Rechteck(breite: Double, hoehe: Double) extends Form
case class Dreieck(a: Double, b: Double, c: Double) extends Form
```

### Rekursive ADTs

```scala
sealed trait Liste[+A]
case object Leer extends Liste[Nothing]
case class Cons[A](kopf: A, rest: Liste[A]) extends Liste[A]

// Beispiel: Cons(1, Cons(2, Cons(3, Leer)))
```

---

## 7.6 Exhaustive Matching

Bei `sealed` Traits warnt der Compiler bei unvollständigem Matching:

```scala
sealed trait Ergebnis
case class Erfolg(wert: Int) extends Ergebnis
case class Fehler(nachricht: String) extends Ergebnis

def verarbeite(e: Ergebnis): String = e match
  case Erfolg(w) => s"Wert: $w"
  // Warnung: Fehler nicht behandelt!
```

### Mit Enum (Scala 3)

```scala
enum Status:
  case Aktiv, Pausiert, Beendet

def beschreibe(s: Status): String = s match
  case Status.Aktiv => "Läuft"
  case Status.Pausiert => "Wartet"
  case Status.Beendet => "Fertig"
```

---

## 7.7 Option, Either, Try

### Option – Optionale Werte

```scala
sealed trait Option[+A]
case class Some[A](wert: A) extends Option[A]
case object None extends Option[Nothing]
```

```scala
def teilen(a: Int, b: Int): Option[Int] =
  if b == 0 then None
  else Some(a / b)

teilen(10, 2) match
  case Some(ergebnis) => println(s"Ergebnis: $ergebnis")
  case None => println("Division durch Null!")
```

### Either – Zwei Möglichkeiten

```scala
sealed trait Either[+L, +R]
case class Left[L](wert: L) extends Either[L, Nothing]
case class Right[R](wert: R) extends Either[Nothing, R]
```

```scala
def parseZahl(s: String): Either[String, Int] =
  s.toIntOption match
    case Some(n) => Right(n)
    case None => Left(s"'$s' ist keine Zahl")

parseZahl("42") match
  case Right(n) => println(s"Zahl: $n")
  case Left(err) => println(s"Fehler: $err")
```

---

## 7.8 Pattern Matching auf Collections

```scala
def summe(liste: List[Int]): Int = liste match
  case Nil => 0
  case kopf :: rest => kopf + summe(rest)

def letztes[A](liste: List[A]): Option[A] = liste match
  case Nil => None
  case x :: Nil => Some(x)
  case _ :: rest => letztes(rest)
```

### Mit @ Operator (Binding)

```scala
liste match
  case alle @ (_ :: _ :: _) => 
    s"Liste mit mindestens 2 Elementen: $alle"
  case _ => "Zu kurz"
```

---

## 7.9 Extraktoren (unapply)

Eigene Pattern-Matcher definieren:

```scala
object Gerade:
  def unapply(n: Int): Boolean = n % 2 == 0

object Ungerade:
  def unapply(n: Int): Boolean = n % 2 != 0

42 match
  case Gerade() => "Gerade!"
  case Ungerade() => "Ungerade!"
```

### Mit Extraktion

```scala
object Email:
  def unapply(s: String): Option[(String, String)] =
    val teile = s.split("@")
    if teile.length == 2 then Some((teile(0), teile(1)))
    else None

"alice@example.com" match
  case Email(benutzer, domain) => 
    s"Benutzer: $benutzer, Domain: $domain"
  case _ => "Keine gültige Email"
```

---

## 7.10 Praktische Beispiele

### Auswertung arithmetischer Ausdrücke

```scala
sealed trait Expr
case class Num(n: Double) extends Expr
case class Add(l: Expr, r: Expr) extends Expr
case class Mul(l: Expr, r: Expr) extends Expr
case class Neg(e: Expr) extends Expr

def eval(e: Expr): Double = e match
  case Num(n) => n
  case Add(l, r) => eval(l) + eval(r)
  case Mul(l, r) => eval(l) * eval(r)
  case Neg(e) => -eval(e)

// (3 + 4) * 2
val expr = Mul(Add(Num(3), Num(4)), Num(2))
eval(expr)  // 14.0
```

### JSON-Verarbeitung

```scala
sealed trait Json
case class JObject(fields: Map[String, Json]) extends Json
case class JArray(items: List[Json]) extends Json
case class JString(s: String) extends Json
case class JNumber(n: Double) extends Json
case class JBool(b: Boolean) extends Json
case object JNull extends Json

def findString(json: Json, key: String): Option[String] = json match
  case JObject(fields) => fields.get(key).flatMap {
    case JString(s) => Some(s)
    case _ => None
  }
  case _ => None
```

---

## Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Literal Pattern | `case 42 =>` |
| Variable Pattern | `case x =>` bindet x |
| Wildcard | `case _ =>` fängt alles |
| Type Pattern | `case s: String =>` |
| Case Class | `case Person(n, a) =>` |
| Guard | `case x if x > 0 =>` |
| Sealed Trait | Exhaustive Matching |
| ADT | Sum + Product Types |
| @ Operator | `case all @ List(x, _*) =>` |
| Extractor | `unapply` für eigene Patterns |

---

## Nächste Schritte

In **Vorlesung 8** lernen wir:
- Higher-Order Functions
- Currying und Partial Application
- Closures

➡️ [Weiter zu Vorlesung 8: Higher-Order Functions](./08-higher-order-functions.md)
