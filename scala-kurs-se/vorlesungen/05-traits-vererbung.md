# Vorlesung 5: Traits & Vererbung

## Lernziele

Nach dieser Vorlesung kannst du:
- Traits definieren und als Mixins verwenden
- Den Unterschied zwischen Traits und abstrakten Klassen verstehen
- Vererbungshierarchien in Scala aufbauen
- Sealed Traits für ADTs nutzen
- Die Linearisierung verstehen

---

## 5.1 Traits – Grundlagen

**Traits** sind wie Interfaces in Java, aber mächtiger: Sie können auch **implementierten Code** enthalten.

### Einfacher Trait

```scala
trait Begruesbar:
  def begruessung: String
  
  def hallo(): Unit = println(s"Hallo! $begruessung")
```

### Trait implementieren

```scala
class Person(val name: String) extends Begruesbar:
  def begruessung: String = s"Ich bin $name."

val p = Person("Alice")
p.hallo()  // "Hallo! Ich bin Alice."
```

---

## 5.2 Mehrfachvererbung mit Mixins

In Scala kann eine Klasse **mehrere Traits** mischen:

```scala
trait Schwimmer:
  def schwimmen(): Unit = println("Ich schwimme!")

trait Laeufer:
  def laufen(): Unit = println("Ich laufe!")

trait Radfahrer:
  def radfahren(): Unit = println("Ich fahre Rad!")

class Triathlet extends Schwimmer, Laeufer, Radfahrer:
  def wettkampf(): Unit =
    schwimmen()
    radfahren()
    laufen()

val sportler = Triathlet()
sportler.wettkampf()
```

### Syntax: extends und with

```scala
// Scala 3 (bevorzugt)
class C extends Trait1, Trait2, Trait3

// Scala 2 Stil (noch gültig)
class C extends Trait1 with Trait2 with Trait3
```

---

## 5.3 Vererbung

### Klassen erweitern

```scala
class Tier(val name: String):
  def laut(): String = "..."

class Hund(name: String) extends Tier(name):
  override def laut(): String = "Wuff!"

class Katze(name: String) extends Tier(name):
  override def laut(): String = "Miau!"

val bello = Hund("Bello")
println(bello.laut())  // "Wuff!"
```

### override-Keyword

Das `override`-Keyword ist **verpflichtend**, wenn du eine Methode überschreibst:

```scala
class Basis:
  def methode(): String = "Basis"

class Abgeleitet extends Basis:
  override def methode(): String = "Abgeleitet"
  
  // Fehler ohne override:
  // def methode(): String = "..."  // Kompiliert nicht!
```

### super aufrufen

```scala
class Basis:
  def info(): String = "Basis-Info"

class Abgeleitet extends Basis:
  override def info(): String = 
    super.info() + " + Erweiterung"
```

---

## 5.4 Abstract Classes vs Traits

### Wann Abstract Class?

```scala
abstract class Tier(val name: String):
  def laut(): String  // Abstrakt
  
  def vorstellen(): Unit =
    println(s"Ich bin $name und sage: ${laut()}")
```

### Unterschiede

| Feature | Abstract Class | Trait |
|---------|---------------|-------|
| Konstruktor-Parameter | ✓ Ja | ✓ Ja (Scala 3) |
| Mehrfachvererbung | ✗ Nein | ✓ Ja |
| Mit Java kompatibel | ✓ Ja | Eingeschränkt |
| Zustand | ✓ Ja | ✓ Ja |

### Trait-Parameter (Scala 3)

```scala
trait Begruesbar(name: String):
  def hallo(): Unit = println(s"Hallo, $name!")

class Person(name: String) extends Begruesbar(name)
```

---

## 5.5 Sealed Traits ⭐

Ein **sealed trait** kann nur in der gleichen Datei erweitert werden. Ideal für ADTs!

```scala
sealed trait Form
case class Kreis(radius: Double) extends Form
case class Rechteck(breite: Double, hoehe: Double) extends Form
case class Dreieck(a: Double, b: Double, c: Double) extends Form
```

### Exhaustive Pattern Matching

Der Compiler warnt bei fehlenden Fällen:

```scala
def flaeche(form: Form): Double = form match
  case Kreis(r) => math.Pi * r * r
  case Rechteck(b, h) => b * h
  // WARNUNG: Dreieck nicht behandelt!
```

```scala
// Vollständig:
def flaeche(form: Form): Double = form match
  case Kreis(r) => math.Pi * r * r
  case Rechteck(b, h) => b * h
  case Dreieck(a, b, c) =>
    val s = (a + b + c) / 2
    math.sqrt(s * (s - a) * (s - b) * (s - c))
```

---

## 5.6 Algebraische Datentypen (ADTs)

ADTs bestehen aus:
- **Sum Types**: "ist entweder A oder B" (sealed trait + case classes)
- **Product Types**: "hat A und B" (case class mit mehreren Feldern)

### Beispiel: Option nachbauen

```scala
sealed trait MeinOption[+A]
case class Etwas[A](wert: A) extends MeinOption[A]
case object Nichts extends MeinOption[Nothing]
```

### Beispiel: Ausdruck-Baum

```scala
sealed trait Ausdruck
case class Zahl(wert: Double) extends Ausdruck
case class Variable(name: String) extends Ausdruck
case class Plus(links: Ausdruck, rechts: Ausdruck) extends Ausdruck
case class Mal(links: Ausdruck, rechts: Ausdruck) extends Ausdruck

def auswerten(expr: Ausdruck, env: Map[String, Double]): Double =
  expr match
    case Zahl(w) => w
    case Variable(n) => env(n)
    case Plus(l, r) => auswerten(l, env) + auswerten(r, env)
    case Mal(l, r) => auswerten(l, env) * auswerten(r, env)

// 3 + x * 2
val formel = Plus(Zahl(3), Mal(Variable("x"), Zahl(2)))
auswerten(formel, Map("x" -> 5))  // 13.0
```

---

## 5.7 Trait Linearisierung

Bei Mehrfachvererbung: Welche Methode wird aufgerufen?

### Das Diamond Problem

```scala
trait A:
  def greet(): String = "A"

trait B extends A:
  override def greet(): String = "B -> " + super.greet()

trait C extends A:
  override def greet(): String = "C -> " + super.greet()

class D extends B, C

val d = D()
println(d.greet())  // Was kommt raus?
```

### Linearisierung

Scala löst dies durch **Linearisierung**. Die Reihenfolge für `D extends B, C`:

```
D → C → B → A → AnyRef → Any
```

Also: `d.greet()` = `"C -> B -> A"`

### Regel

Die Linearisierung geht von **rechts nach links** durch die Traits, ohne Duplikate.

---

## 5.8 Self Types

Traits können verlangen, dass sie nur mit bestimmten anderen Typen gemischt werden:

```scala
trait Logger:
  def log(msg: String): Unit

trait DatabaseAccess:
  self: Logger =>  // Muss mit Logger gemischt werden!
  
  def query(sql: String): Unit =
    log(s"Executing: $sql")
    // Datenbankabfrage...

// OK:
class MyDB extends DatabaseAccess, Logger:
  def log(msg: String) = println(msg)

// Fehler:
// class BadDB extends DatabaseAccess  // Logger fehlt!
```

---

## 5.9 Polymorphismus

### Subtyp-Polymorphismus

```scala
val tiere: List[Tier] = List(
  Hund("Bello"),
  Katze("Minka"),
  Hund("Rex")
)

for tier <- tiere do
  println(s"${tier.name} sagt: ${tier.laut()}")
```

### Kovariant, Kontravariant, Invariant

```scala
class Container[+A]     // Kovariant: Container[Hund] ist Container[Tier]
class Consumer[-A]      // Kontravariant
class Invariant[A]      // Invariant (default)
```

---

## Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Trait | Interface mit Implementation, Mehrfachvererbung |
| Mixin | Trait zu Klasse hinzumischen |
| `extends` | Erbt von Klasse oder Trait |
| `override` | Markiert überschriebene Methoden |
| `sealed` | Trait nur in gleicher Datei erweiterbar |
| ADT | Sum Types + Product Types |
| Linearisierung | Auflösung bei Mehrfachvererbung |
| Self Type | `self: T =>` erfordert Mixin mit T |

---

## Nächste Schritte

In **Vorlesung 6** lernen wir:
- Collections im Detail
- map, filter, flatMap, fold
- For-Comprehensions auf Collections

➡️ [Weiter zu Vorlesung 6: Collections](./06-collections.md)
