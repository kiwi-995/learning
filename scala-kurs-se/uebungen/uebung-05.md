# Übung 5: Traits & Vererbung

**Thema:** Traits, Mixins, Vererbung, Sealed Traits, ADTs

**Geschätzte Zeit:** 90-120 Minuten

---

## Teil A: Theoretische Fragen

### Frage 5.1
Was ist der Unterschied zwischen einem Trait und einer abstrakten Klasse? Wann würdest du welches verwenden?

### Frage 5.2
Was bedeutet `sealed` bei einem Trait? Welchen Vorteil bietet es?

### Frage 5.3
Erkläre das Konzept der Linearisierung. Was passiert bei `class D extends A, B, C`?

### Frage 5.4
Was ist ein ADT (Algebraischer Datentyp)? Aus welchen zwei Komponenten besteht er?

### Frage 5.5
Was ist ein Self-Type? Gib ein Beispiel.

---

## Teil B: REPL-Übungen

### Aufgabe 5.1 – Einfache Traits
Definiere und teste in der REPL:

```scala
trait Sprechend:
  def sprechen(): String

class Mensch(val name: String) extends Sprechend:
  def sprechen() = s"Hallo, ich bin $name"

class Roboter(val modell: String) extends Sprechend:
  def sprechen() = s"Beep boop, Modell $modell"
```

### Aufgabe 5.2 – Mixins
Erstelle Traits und mische sie:

```scala
trait Fliegend:
  def fliegen(): Unit = println("Ich fliege!")

trait Sprechend:
  def sprechen(): Unit

class Papagei extends Fliegend, Sprechend:
  def sprechen() = println("Polly will einen Cracker!")

// Teste beide Methoden
```

### Aufgabe 5.3 – Sealed Trait
Erstelle einen sealed trait und teste Pattern Matching:

```scala
sealed trait Verkehrsmittel
case class Auto(marke: String) extends Verkehrsmittel
case class Fahrrad(typ: String) extends Verkehrsmittel
case object Roller extends Verkehrsmittel

def beschreibe(v: Verkehrsmittel): String = ???
```

---

## Teil C: Programmieraufgaben

### Aufgabe 5.4 – Tierhierarchie
Erstelle eine Tierhierarchie mit Traits:

```scala
trait Tier:
  def name: String
  def laut: String
  def beschreibung: String = s"$name sagt: $laut"

trait Haustier extends Tier:
  def besitzer: String
  override def beschreibung: String = 
    s"${super.beschreibung} (Besitzer: $besitzer)"

trait Wildtier extends Tier:
  def lebensraum: String

case class Hund(name: String, besitzer: String) extends Haustier:
  def laut = "Wuff"

case class Katze(name: String, besitzer: String) extends Haustier:
  def laut = "Miau"

case class Wolf(name: String, lebensraum: String) extends Wildtier:
  def laut = "Auuu"

// Teste mit einer Liste von Tieren
```

### Aufgabe 5.5 – Ausdrucks-Parser
Erweitere den Ausdrucks-Baum aus der Vorlesung:

```scala
sealed trait Ausdruck

case class Zahl(wert: Double) extends Ausdruck
case class Variable(name: String) extends Ausdruck
case class Plus(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Minus(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Mal(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Geteilt(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Negation(expr: Ausdruck) extends Ausdruck

object Ausdruck:
  def auswerten(expr: Ausdruck, env: Map[String, Double]): Double = ???
  def alsString(expr: Ausdruck): String = ???
  def vereinfachen(expr: Ausdruck): Ausdruck = ???
```

Implementiere:
1. `auswerten`: Berechnet den Wert
2. `alsString`: Erzeugt lesbaren String wie `"(3 + x) * 2"`
3. `vereinfachen`: Vereinfacht z.B. `Plus(Zahl(0), x)` zu `x`

### Aufgabe 5.6 – JSON-Datentyp ⭐
Modelliere einen JSON-Wert als ADT:

```scala
sealed trait JsonWert

case class JsonObjekt(felder: Map[String, JsonWert]) extends JsonWert
case class JsonArray(elemente: List[JsonWert]) extends JsonWert
case class JsonString(wert: String) extends JsonWert
case class JsonZahl(wert: Double) extends JsonWert
case class JsonBool(wert: Boolean) extends JsonWert
case object JsonNull extends JsonWert

object Json:
  def format(json: JsonWert, einrueckung: Int = 0): String = ???
  def parse(s: String): JsonWert = ???  // Bonus!
```

Beispiel:
```scala
val person = JsonObjekt(Map(
  "name" -> JsonString("Max"),
  "alter" -> JsonZahl(25),
  "hobbies" -> JsonArray(List(
    JsonString("Scala"),
    JsonString("Git")
  ))
))

println(Json.format(person))
// {
//   "name": "Max",
//   "alter": 25,
//   "hobbies": ["Scala", "Git"]
// }
```

### Aufgabe 5.7 – Plugin-System ⭐
Erstelle ein erweiterbares Plugin-System:

```scala
trait Plugin:
  def name: String
  def version: String
  def execute(): String

trait Loggable:
  self: Plugin =>
  def log(msg: String): Unit = 
    println(s"[${name}] $msg")

trait Configurable:
  self: Plugin =>
  def config: Map[String, String]
  def getConfig(key: String): Option[String] = config.get(key)

class GreeterPlugin extends Plugin, Loggable, Configurable:
  val name = "Greeter"
  val version = "1.0"
  val config = Map("greeting" -> "Hello")
  
  def execute(): String =
    log("Executing...")
    getConfig("greeting").getOrElse("Hi") + ", World!"

// Erstelle weitere Plugins
```

---

## Teil D: Verständnisfragen

### Frage 5.6
Was ist die Ausgabe? Erkläre die Linearisierung.

```scala
trait A { def x = "A" }
trait B extends A { override def x = "B" + super.x }
trait C extends A { override def x = "C" + super.x }
class D extends B, C

println(D().x)  // ?
```

### Frage 5.7
Warum kompiliert dieser Code nicht?

```scala
trait Datenbank:
  def verbinden(): Unit

trait Logger:
  self: Datenbank =>
  def log(msg: String): Unit

class MeineApp extends Logger:  // Fehler!
  def log(msg: String) = println(msg)
```

### Frage 5.8
Was ist der Unterschied zwischen diesen zwei Definitionen?

```scala
// Version A
sealed trait Option[+A]
case class Some[A](value: A) extends Option[A]
case object None extends Option[Nothing]

// Version B
sealed trait Option[A]
case class Some[A](value: A) extends Option[A]
case class None[A]() extends Option[A]
```

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – Zustandsautomat
Modelliere einen einfachen Zustandsautomaten:

```scala
sealed trait Zustand
case object Gesperrt extends Zustand
case object Entsperrt extends Zustand

sealed trait Ereignis  
case object Muenze extends Ereignis
case object Drehen extends Ereignis

def transition(zustand: Zustand, ereignis: Ereignis): Zustand = ???

// Simuliere: Drehkreuz (Muenze entsperrt, Drehen sperrt wieder)
```

### Bonus 2 – Typenklasse Pattern
Implementiere eine einfache Show-Typklasse:

```scala
trait Show[A]:
  def show(a: A): String

object Show:
  given Show[Int] = _.toString
  given Show[String] = s => s"\"$s\""
  given [A](using sa: Show[A]): Show[List[A]] = ???
  
extension [A](a: A)
  def show(using s: Show[A]): String = s.show(a)

// Teste
println(42.show)
println("Hallo".show)
println(List(1, 2, 3).show)
```

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-05.md](../loesungen/loesung-05.md)
