# Lösungen zu Übung 5

---

## Teil A: Theoretische Fragen

### Lösung 5.1
**Trait:**
- Ermöglicht Mehrfachvererbung (Mixins)
- Kann nicht mit `new` instanziiert werden
- Seit Scala 3: Kann Konstruktor-Parameter haben

**Abstrakte Klasse:**
- Nur Einfachvererbung möglich
- Kann Konstruktor-Parameter haben
- Bessere Java-Interoperabilität

**Wann was verwenden:**
- Trait: Wenn mehrere Klassen das Verhalten mischen sollen
- Abstrakte Klasse: Wenn Java-Kompatibilität wichtig oder klare "ist-ein" Beziehung

### Lösung 5.2
`sealed` bedeutet, dass der Trait nur in der **gleichen Datei** erweitert werden kann.

**Vorteile:**
1. Compiler kennt alle Subtypen
2. Exhaustive Pattern Matching mit Warnungen
3. Ideal für ADTs (Algebraische Datentypen)

### Lösung 5.3
Bei `class D extends A, B, C` ist die Linearisierung:
```
D → C → B → A → AnyRef → Any
```

Von rechts nach links, ohne Duplikate. Bei `super`-Aufrufen wird diese Kette durchlaufen.

### Lösung 5.4
Ein **ADT** besteht aus:
1. **Sum Types**: "ist entweder A oder B" (sealed trait + case classes)
2. **Product Types**: "hat A und B" (case class mit Feldern)

Beispiel: `Option[A]` ist entweder `Some(a)` oder `None`.

### Lösung 5.5
Ein **Self-Type** zwingt einen Trait, nur mit bestimmten anderen Typen gemischt zu werden:

```scala
trait Logger:
  def log(msg: String): Unit

trait Service:
  self: Logger =>  // Self-Type
  def doWork(): Unit =
    log("Working...")  // Logger-Methode verfügbar
```

---

## Teil B: REPL-Übungen

### Lösung 5.1
```scala
trait Sprechend:
  def sprechen(): String

class Mensch(val name: String) extends Sprechend:
  def sprechen() = s"Hallo, ich bin $name"

class Roboter(val modell: String) extends Sprechend:
  def sprechen() = s"Beep boop, Modell $modell"

val sprecher: List[Sprechend] = List(
  Mensch("Alice"),
  Roboter("R2D2")
)
sprecher.foreach(s => println(s.sprechen()))
```

### Lösung 5.2
```scala
trait Fliegend:
  def fliegen(): Unit = println("Ich fliege!")

trait Sprechend:
  def sprechen(): Unit

class Papagei extends Fliegend, Sprechend:
  def sprechen() = println("Polly will einen Cracker!")

val polly = Papagei()
polly.fliegen()   // "Ich fliege!"
polly.sprechen()  // "Polly will einen Cracker!"
```

### Lösung 5.3
```scala
sealed trait Verkehrsmittel
case class Auto(marke: String) extends Verkehrsmittel
case class Fahrrad(typ: String) extends Verkehrsmittel
case object Roller extends Verkehrsmittel

def beschreibe(v: Verkehrsmittel): String = v match
  case Auto(marke) => s"Ein $marke Auto"
  case Fahrrad(typ) => s"Ein $typ Fahrrad"
  case Roller => "Ein Roller"

println(beschreibe(Auto("BMW")))   // "Ein BMW Auto"
println(beschreibe(Roller))        // "Ein Roller"
```

---

## Teil C: Programmieraufgaben

### Lösung 5.4 – Tierhierarchie
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
  override def beschreibung: String =
    s"${super.beschreibung} (Lebensraum: $lebensraum)"

case class Hund(name: String, besitzer: String) extends Haustier:
  def laut = "Wuff"

case class Katze(name: String, besitzer: String) extends Haustier:
  def laut = "Miau"

case class Wolf(name: String, lebensraum: String) extends Wildtier:
  def laut = "Auuu"

val tiere: List[Tier] = List(
  Hund("Bello", "Max"),
  Katze("Minka", "Anna"),
  Wolf("Grey", "Wald")
)

tiere.foreach(t => println(t.beschreibung))
```

### Lösung 5.5 – Ausdrucks-Parser
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
  def auswerten(expr: Ausdruck, env: Map[String, Double]): Double = expr match
    case Zahl(w) => w
    case Variable(n) => env.getOrElse(n, 
      throw new RuntimeException(s"Variable $n nicht definiert"))
    case Plus(l, r) => auswerten(l, env) + auswerten(r, env)
    case Minus(l, r) => auswerten(l, env) - auswerten(r, env)
    case Mal(l, r) => auswerten(l, env) * auswerten(r, env)
    case Geteilt(l, r) => auswerten(l, env) / auswerten(r, env)
    case Negation(e) => -auswerten(e, env)

  def alsString(expr: Ausdruck): String = expr match
    case Zahl(w) => w.toString
    case Variable(n) => n
    case Plus(l, r) => s"(${alsString(l)} + ${alsString(r)})"
    case Minus(l, r) => s"(${alsString(l)} - ${alsString(r)})"
    case Mal(l, r) => s"(${alsString(l)} * ${alsString(r)})"
    case Geteilt(l, r) => s"(${alsString(l)} / ${alsString(r)})"
    case Negation(e) => s"(-${alsString(e)})"

  def vereinfachen(expr: Ausdruck): Ausdruck = expr match
    // 0 + x = x
    case Plus(Zahl(0), x) => vereinfachen(x)
    case Plus(x, Zahl(0)) => vereinfachen(x)
    // x - 0 = x
    case Minus(x, Zahl(0)) => vereinfachen(x)
    // 1 * x = x
    case Mal(Zahl(1), x) => vereinfachen(x)
    case Mal(x, Zahl(1)) => vereinfachen(x)
    // 0 * x = 0
    case Mal(Zahl(0), _) => Zahl(0)
    case Mal(_, Zahl(0)) => Zahl(0)
    // x / 1 = x
    case Geteilt(x, Zahl(1)) => vereinfachen(x)
    // Rekursiv vereinfachen
    case Plus(l, r) => Plus(vereinfachen(l), vereinfachen(r))
    case Minus(l, r) => Minus(vereinfachen(l), vereinfachen(r))
    case Mal(l, r) => Mal(vereinfachen(l), vereinfachen(r))
    case Geteilt(l, r) => Geteilt(vereinfachen(l), vereinfachen(r))
    case Negation(e) => Negation(vereinfachen(e))
    case other => other

// Test
val expr = Plus(Zahl(0), Mal(Variable("x"), Zahl(1)))
println(Ausdruck.alsString(expr))        // "(0.0 + (x * 1.0))"
println(Ausdruck.alsString(Ausdruck.vereinfachen(expr)))  // "x"
```

### Lösung 5.6 – JSON-Datentyp
```scala
sealed trait JsonWert
case class JsonObjekt(felder: Map[String, JsonWert]) extends JsonWert
case class JsonArray(elemente: List[JsonWert]) extends JsonWert
case class JsonString(wert: String) extends JsonWert
case class JsonZahl(wert: Double) extends JsonWert
case class JsonBool(wert: Boolean) extends JsonWert
case object JsonNull extends JsonWert

object Json:
  def format(json: JsonWert, einrueckung: Int = 0): String =
    val indent = "  " * einrueckung
    val nextIndent = "  " * (einrueckung + 1)
    
    json match
      case JsonString(s) => s"\"$s\""
      case JsonZahl(n) => if n == n.toLong then n.toLong.toString else n.toString
      case JsonBool(b) => b.toString
      case JsonNull => "null"
      case JsonArray(elemente) =>
        if elemente.isEmpty then "[]"
        else elemente.map(e => nextIndent + format(e, einrueckung + 1))
          .mkString("[\n", ",\n", s"\n$indent]")
      case JsonObjekt(felder) =>
        if felder.isEmpty then "{}"
        else felder.map { case (k, v) => 
          s"$nextIndent\"$k\": ${format(v, einrueckung + 1)}"
        }.mkString("{\n", ",\n", s"\n$indent}")

// Test
val person = JsonObjekt(Map(
  "name" -> JsonString("Max"),
  "alter" -> JsonZahl(25),
  "aktiv" -> JsonBool(true),
  "hobbies" -> JsonArray(List(
    JsonString("Scala"),
    JsonString("Git")
  ))
))

println(Json.format(person))
```

### Lösung 5.7 – Plugin-System
```scala
trait Plugin:
  def name: String
  def version: String
  def execute(): String

trait Loggable:
  self: Plugin =>
  def log(msg: String): Unit = 
    println(s"[${name} v${version}] $msg")

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

class CalculatorPlugin extends Plugin, Loggable:
  val name = "Calculator"
  val version = "2.0"
  
  def execute(): String =
    log("Calculating...")
    s"2 + 2 = ${2 + 2}"

// Test
val plugins: List[Plugin] = List(GreeterPlugin(), CalculatorPlugin())
plugins.foreach(p => println(p.execute()))
```

---

## Teil D: Verständnisfragen

### Lösung 5.6
```scala
trait A { def x = "A" }
trait B extends A { override def x = "B" + super.x }
trait C extends A { override def x = "C" + super.x }
class D extends B, C

println(D().x)  // "CBA"
```

Linearisierung: `D → C → B → A`
- `D.x` ruft `C.x` auf
- `C.x` = "C" + `super.x` = "C" + `B.x`
- `B.x` = "B" + `super.x` = "B" + `A.x`
- `A.x` = "A"
- Ergebnis: "C" + "B" + "A" = "CBA"

### Lösung 5.7
```scala
class MeineApp extends Logger:  // Fehler!
```

**Problem:** `Logger` hat Self-Type `self: Datenbank =>`, aber `MeineApp` ist keine `Datenbank`.

**Korrektur:**
```scala
class MeineApp extends Logger, Datenbank:
  def log(msg: String) = println(msg)
  def verbinden(): Unit = println("Verbunden")
```

### Lösung 5.8
**Version A (kovariant mit +A):**
- `None` ist ein Singleton (`case object`)
- `None` hat Typ `Option[Nothing]`
- Wegen Kovarianz: `Option[Nothing]` ist Subtyp von `Option[A]` für alle A
- Speichereffizient: nur ein `None`-Objekt

**Version B (invariant):**
- `None[A]()` muss mit Typparameter instanziiert werden
- Für jeden Typ A ein eigenes None-Objekt
- Weniger elegant, mehr Speicher

---

## Bonusaufgaben

### Bonus 1 – Zustandsautomat
```scala
sealed trait Zustand
case object Gesperrt extends Zustand
case object Entsperrt extends Zustand

sealed trait Ereignis  
case object Muenze extends Ereignis
case object Drehen extends Ereignis

def transition(zustand: Zustand, ereignis: Ereignis): Zustand =
  (zustand, ereignis) match
    case (Gesperrt, Muenze) => Entsperrt
    case (Entsperrt, Drehen) => Gesperrt
    case (z, _) => z  // Sonst: Zustand bleibt

// Simulation
val ereignisse = List(Drehen, Muenze, Drehen, Muenze, Muenze, Drehen)
val endzustand = ereignisse.foldLeft(Gesperrt: Zustand)(transition)
println(s"Endzustand: $endzustand")  // Gesperrt
```

### Bonus 2 – Typenklasse Pattern
```scala
trait Show[A]:
  def show(a: A): String

object Show:
  given Show[Int] = (n: Int) => n.toString
  given Show[String] = (s: String) => s"\"$s\""
  given Show[Boolean] = (b: Boolean) => b.toString
  
  given [A](using sa: Show[A]): Show[List[A]] = (list: List[A]) =>
    list.map(sa.show).mkString("[", ", ", "]")
  
extension [A](a: A)
  def show(using s: Show[A]): String = s.show(a)

// Test
println(42.show)              // "42"
println("Hallo".show)         // "\"Hallo\""
println(true.show)            // "true"
println(List(1, 2, 3).show)   // "[1, 2, 3]"
```
