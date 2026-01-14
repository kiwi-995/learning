# Vorlesung 4: Klassen & Objekte

## Lernziele

Nach dieser Vorlesung kannst du:
- Klassen mit Konstruktor-Parametern definieren
- Companion Objects und Factory-Methoden verstehen
- Case Classes für Datenmodellierung nutzen
- Sichtbarkeit mit private/protected steuern
- Den Unterschied zwischen Klassen und Objects erklären

---

## 4.1 Klassen in Scala

### Grundlegende Klassendefinition

```scala
class Person(name: String, alter: Int):
  def vorstellen(): String =
    s"Hallo, ich bin $name und $alter Jahre alt."
```

**Instanziierung:**
```scala
val alice = Person("Alice", 25)
println(alice.vorstellen())  // "Hallo, ich bin Alice und 25 Jahre alt."
```

### Primärer Konstruktor

In Scala ist der **Konstruktor Teil der Klassendeklaration**:

```scala
class Rechteck(breite: Double, hoehe: Double):
  // Konstruktor-Parameter sind hier direkt verfügbar
  val flaeche: Double = breite * hoehe
  val umfang: Double = 2 * (breite + hoehe)
  
  override def toString: String =
    s"Rechteck($breite x $hoehe)"
```

```scala
val r = Rechteck(5.0, 3.0)
println(r.flaeche)   // 15.0
println(r.umfang)    // 16.0
println(r)           // "Rechteck(5.0 x 3.0)"
```

---

## 4.2 Sichtbarkeit von Parametern

### Standardmäßig: Privat

```scala
class Geheimnis(wert: Int):  // wert ist NICHT von außen zugänglich
  def getWert: Int = wert

val g = Geheimnis(42)
// g.wert        // FEHLER: wert is not accessible
g.getWert        // 42
```

### Mit val: Öffentlich lesbar

```scala
class Person(val name: String, val alter: Int)

val p = Person("Bob", 30)
p.name           // "Bob" ✓
p.alter          // 30 ✓
// p.name = "X"  // FEHLER: Zuweisung nicht erlaubt (val)
```

### Mit var: Öffentlich les- und schreibbar

```scala
class Zaehler(var wert: Int):
  def inkrement(): Unit = wert += 1

val z = Zaehler(0)
z.wert           // 0
z.inkrement()
z.wert           // 1
z.wert = 100     // OK (var!)
```

### Übersicht

| Definition | Lesbar von außen? | Schreibbar? |
|------------|-------------------|-------------|
| `class C(x: Int)` | Nein | Nein |
| `class C(val x: Int)` | Ja | Nein |
| `class C(var x: Int)` | Ja | Ja |
| `class C(private val x: Int)` | Nein | Nein |

---

## 4.3 Methoden und Felder

### Methoden definieren

```scala
class Bankkonto(inhaber: String, private var kontostand: Double):
  
  def einzahlen(betrag: Double): Unit =
    require(betrag > 0, "Betrag muss positiv sein")
    kontostand += betrag
  
  def abheben(betrag: Double): Boolean =
    if betrag <= kontostand then
      kontostand -= betrag
      true
    else
      false
  
  def gibKontostand: Double = kontostand
  
  override def toString: String =
    f"Konto von $inhaber: $kontostand%.2f EUR"
```

```scala
val konto = Bankkonto("Max", 100.0)
konto.einzahlen(50.0)
println(konto.gibKontostand)  // 150.0
konto.abheben(30.0)           // true
println(konto)                // "Konto von Max: 120.00 EUR"
```

### require und assert

```scala
class Alter(wert: Int):
  require(wert >= 0, "Alter kann nicht negativ sein")
  require(wert <= 150, "Alter unrealistisch hoch")

Alter(25)   // OK
Alter(-5)   // IllegalArgumentException
```

---

## 4.4 Auxiliary Constructors

Zusätzliche Konstruktoren mit `this`:

```scala
class Punkt(val x: Double, val y: Double):
  
  // Sekundär-Konstruktor: Nur ein Wert (Punkt auf Diagonale)
  def this(wert: Double) = this(wert, wert)
  
  // Sekundär-Konstruktor: Ursprung
  def this() = this(0.0, 0.0)
  
  override def toString = s"Punkt($x, $y)"
```

```scala
val p1 = Punkt(3.0, 4.0)  // Punkt(3.0, 4.0)
val p2 = Punkt(5.0)       // Punkt(5.0, 5.0)
val p3 = Punkt()          // Punkt(0.0, 0.0)
```

---

## 4.5 Objects (Singleton)

Ein **Object** ist eine einzelne Instanz – kein `new` nötig!

```scala
object MathHelfer:
  val pi: Double = 3.14159265
  
  def quadrat(x: Double): Double = x * x
  
  def kreisflaeche(radius: Double): Double =
    pi * quadrat(radius)
```

**Verwendung:**
```scala
println(MathHelfer.pi)           // 3.14159265
println(MathHelfer.quadrat(5))   // 25.0
println(MathHelfer.kreisflaeche(2.0))  // 12.566...
```

### Wann Objects verwenden?

```
┌─────────────────────────────────────────────────────────────┐
│                   Objects verwenden für:                    │
├─────────────────────────────────────────────────────────────┤
│ ✓ Utility-Funktionen (wie static in Java)                   │
│ ✓ Konstanten                                                │
│ ✓ Einstiegspunkte (@main)                                   │
│ ✓ Singleton-Pattern                                         │
│ ✓ Factory-Methoden (Companion Objects)                      │
└─────────────────────────────────────────────────────────────┘
```

---

## 4.6 Companion Objects

Ein **Companion Object** hat den gleichen Namen wie eine Klasse und steht in der gleichen Datei.

```scala
class Kreis private (val radius: Double):
  def flaeche: Double = Kreis.pi * radius * radius
  def umfang: Double = 2 * Kreis.pi * radius
  
object Kreis:
  val pi: Double = 3.14159265
  
  // Factory-Methode
  def apply(radius: Double): Kreis =
    require(radius > 0, "Radius muss positiv sein")
    new Kreis(radius)
  
  // Einheitskreis
  val einheitskreis: Kreis = new Kreis(1.0)
```

**Verwendung:**
```scala
val k1 = Kreis(5.0)        // Ruft apply() auf
val k2 = Kreis.einheitskreis

println(k1.flaeche)        // 78.5398...
println(Kreis.pi)          // 3.14159265
```

### Was macht Companion Objects besonders?

1. **Zugriff auf private Members** der Klasse
2. **apply-Methode** ermöglicht `Klasse(args)` statt `new Klasse(args)`
3. Ersetzt **static** aus Java

---

## 4.7 Case Classes ⭐

**Case Classes** sind die bevorzugte Art, Daten in Scala zu modellieren!

```scala
case class Person(name: String, alter: Int)
```

### Was bekommst du automatisch?

```scala
case class Punkt(x: Double, y: Double)

val p1 = Punkt(3, 4)       // Kein 'new' nötig (apply generiert)
val p2 = Punkt(3, 4)

// 1. Automatische toString
println(p1)                // "Punkt(3.0,4.0)"

// 2. Automatische equals (strukturell, nicht Referenz)
p1 == p2                   // true

// 3. Automatische hashCode
p1.hashCode == p2.hashCode // true

// 4. copy-Methode
val p3 = p1.copy(x = 10)   // Punkt(10.0, 4.0)

// 5. Pattern Matching Unterstützung
p1 match
  case Punkt(x, y) => println(s"x=$x, y=$y")
```

### Case Classes sind unveränderlich!

```scala
case class User(name: String, email: String)

val user = User("Alice", "alice@example.com")
// user.name = "Bob"  // FEHLER: val kann nicht geändert werden

// Stattdessen: Kopie mit Änderung
val neuerUser = user.copy(name = "Bob")
```

### Case Classes vs normale Klassen

| Feature | Normale Klasse | Case Class |
|---------|----------------|------------|
| `new` nötig? | Ja | Nein |
| `toString` | "@hashcode" | "Name(felder)" |
| `equals` | Referenz | Struktur |
| `copy` | Nein | Ja |
| Pattern Matching | Manuell | Automatisch |
| Felder | Private (default) | Public val (default) |

---

## 4.8 Case Objects

Für Singletons als Teil einer Typhierarchie:

```scala
sealed trait Ampelfarbe
case object Rot extends Ampelfarbe
case object Gelb extends Ampelfarbe
case object Gruen extends Ampelfarbe

def darfFahren(farbe: Ampelfarbe): Boolean = farbe match
  case Gruen => true
  case _ => false
```

---

## 4.9 Enums in Scala 3 ⭐

Scala 3 hat native Enum-Unterstützung:

```scala
enum Wochentag:
  case Montag, Dienstag, Mittwoch, Donnerstag, Freitag, Samstag, Sonntag

enum Planet(val masse: Double, val radius: Double):
  case Merkur extends Planet(3.3e23, 2.4e6)
  case Venus extends Planet(4.9e24, 6.1e6)
  case Erde extends Planet(6.0e24, 6.4e6)
  // ...
  
  def schwerkraft: Double = 6.67e-11 * masse / (radius * radius)
```

**Verwendung:**
```scala
val heute = Wochentag.Mittwoch

heute match
  case Wochentag.Samstag | Wochentag.Sonntag => "Wochenende!"
  case _ => "Arbeitstag"

// Alle Werte
Wochentag.values  // Array(Montag, Dienstag, ...)

// Ordinal
Wochentag.Mittwoch.ordinal  // 2
```

---

## 4.10 Private und Protected

### private

Nur innerhalb der Klasse/Object zugänglich:

```scala
class Geheimvoll:
  private val geheim = "sssh"
  private def interneMethode() = geheim.toUpperCase
  
  def oeffentlich(): String = interneMethode()
```

### protected

Zugänglich für Klasse und Subklassen:

```scala
class Basis:
  protected val wert = 42

class Abgeleitet extends Basis:
  def leseWert = wert  // OK
```

### private[scope]

Beschränkt auf bestimmten Scope:

```scala
package meinpaket

class MeineKlasse:
  private[meinpaket] val paketPrivat = "sichtbar im Paket"
  private[this] val nurHier = "nur in dieser Instanz"
```

---

## Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| `class` | Blaupause für Objekte |
| `object` | Singleton-Instanz |
| Companion Object | Object + Class mit gleichem Namen |
| `case class` | Unveränderliche Datenklasse mit extras |
| `enum` | Aufzählungstyp (Scala 3) |
| `val` in Konstruktor | Öffentlich lesbar |
| `var` in Konstruktor | Öffentlich les-/schreibbar |
| `private` | Nur innerhalb der Klasse |
| `apply()` | Ermöglicht `Klasse(args)` |

---

## Nächste Schritte

In **Vorlesung 5** lernen wir:
- Traits und Mixins
- Vererbung in Scala
- Sealed Traits für ADTs

➡️ [Weiter zu Vorlesung 5: Traits & Vererbung](./05-traits-vererbung.md)
