# Lösungen zu Übung 4

---

## Teil A: Theoretische Fragen

### Lösung 4.1
- **class**: Blaupause für Objekte. Es können mehrere Instanzen erstellt werden mit `new`.
- **object**: Singleton – es gibt genau eine Instanz. Kein `new` nötig. Wird für Utility-Funktionen, Konstanten und Factory-Methoden verwendet.

### Lösung 4.2
```scala
class A(x: Int)      // x nur intern verfügbar (private)
class B(val x: Int)  // x öffentlich lesbar
class C(var x: Int)  // x öffentlich les- und schreibbar
```

### Lösung 4.3
Ein **Companion Object** ist ein Object mit dem gleichen Namen wie eine Klasse in derselben Datei.

**Vorteile:**
1. Zugriff auf `private` Members der Klasse
2. Factory-Methoden via `apply()`
3. Ersetzt `static` aus Java
4. Platz für Konstanten und Utility-Methoden

### Lösung 4.4
Eine **Case Class** ist eine spezialisierte Klasse für unveränderliche Daten.

**Automatisch generiert:**
- `apply()` – Factory-Methode
- `unapply()` – Für Pattern Matching
- `toString` – Formatierte Ausgabe
- `equals` / `hashCode` – Strukturelle Gleichheit
- `copy()` – Kopie mit optionalen Änderungen
- Alle Parameter als `val` (public)

### Lösung 4.5
**Case Class verwenden für:**
- Unveränderliche Daten (DTOs, Value Objects)
- Daten, die verglichen werden sollen
- Pattern Matching
- Kurze Datencontainer

**Normale Klasse verwenden für:**
- Veränderlicher Zustand nötig
- Komplexe Initialisierungslogik
- Vererbungshierarchien (außer sealed)
- Wenn `equals` Referenzgleichheit sein soll

---

## Teil B: REPL-Übungen

### Lösung 4.1
```scala
class Auto(val marke: String, val baujahr: Int):
  def alter: Int = 2024 - baujahr
  override def toString = s"$marke ($baujahr)"

val bmw = Auto("BMW", 2020)
println(bmw)        // "BMW (2020)"
println(bmw.alter)  // 4
```

### Lösung 4.2
```scala
object StringUtils:
  def istPalindrom(s: String): Boolean =
    val bereinigt = s.toLowerCase.filter(_.isLetter)
    bereinigt == bereinigt.reverse
  
  def anzahlVokale(s: String): Int =
    s.toLowerCase.count("aeiouäöü".contains(_))

StringUtils.istPalindrom("Anna")      // true
StringUtils.istPalindrom("Hallo")     // false
StringUtils.anzahlVokale("Hallo")     // 2
```

### Lösung 4.3
```scala
case class Student(name: String, matrikelNr: Int, fach: String)

val s1 = Student("Max", 12345, "Informatik")
val s2 = Student("Max", 12345, "Informatik")

s1 == s2                          // true (strukturelle Gleichheit)
s1.toString                       // "Student(Max,12345,Informatik)"
s1.copy(fach = "Mathematik")      // Student(Max,12345,Mathematik)
```

---

## Teil C: Programmieraufgaben

### Lösung 4.4 – Bruchrechnung
```scala
case class Bruch(zaehler: Int, nenner: Int):
  require(nenner != 0, "Nenner darf nicht 0 sein")
  
  def plus(other: Bruch): Bruch =
    Bruch(
      zaehler * other.nenner + other.zaehler * nenner,
      nenner * other.nenner
    ).kuerzen
  
  def minus(other: Bruch): Bruch =
    Bruch(
      zaehler * other.nenner - other.zaehler * nenner,
      nenner * other.nenner
    ).kuerzen
  
  def mal(other: Bruch): Bruch =
    Bruch(zaehler * other.zaehler, nenner * other.nenner).kuerzen
  
  def durch(other: Bruch): Bruch =
    mal(Bruch(other.nenner, other.zaehler))
  
  def kuerzen: Bruch =
    val teiler = Bruch.ggt(zaehler, nenner)
    val vorzeichen = if nenner < 0 then -1 else 1
    Bruch(vorzeichen * zaehler / teiler, vorzeichen * nenner / teiler)
  
  override def toString = s"$zaehler/$nenner"

object Bruch:
  def ggt(a: Int, b: Int): Int = if b == 0 then a.abs else ggt(b, a % b)

// Test
val b1 = Bruch(1, 2)
val b2 = Bruch(1, 4)
println(b1.plus(b2))   // 3/4
println(b1.mal(b2))    // 1/8
println(b1.durch(b2))  // 2/1
```

### Lösung 4.5 – Geometrie
```scala
case class Punkt(x: Double, y: Double):
  def distanzZu(other: Punkt): Double =
    val dx = x - other.x
    val dy = y - other.y
    math.sqrt(dx * dx + dy * dy)

case class Kreis(mittelpunkt: Punkt, radius: Double):
  require(radius > 0, "Radius muss positiv sein")
  
  def flaeche: Double = math.Pi * radius * radius
  def umfang: Double = 2 * math.Pi * radius
  def enthaelt(p: Punkt): Boolean = mittelpunkt.distanzZu(p) <= radius

case class Rechteck(eckeUntenLinks: Punkt, breite: Double, hoehe: Double):
  require(breite > 0 && hoehe > 0, "Dimensionen müssen positiv sein")
  
  def flaeche: Double = breite * hoehe
  def umfang: Double = 2 * (breite + hoehe)
  def eckeObenRechts: Punkt = Punkt(eckeUntenLinks.x + breite, eckeUntenLinks.y + hoehe)
  
  def enthaelt(p: Punkt): Boolean =
    p.x >= eckeUntenLinks.x && p.x <= eckeObenRechts.x &&
    p.y >= eckeUntenLinks.y && p.y <= eckeObenRechts.y

// Test
val kreis = Kreis(Punkt(0, 0), 5)
println(kreis.enthaelt(Punkt(3, 4)))  // true (Distanz = 5)
println(kreis.enthaelt(Punkt(4, 4)))  // false
```

### Lösung 4.6 – Bankkonto mit History
```scala
case class Transaktion(art: String, betrag: Double, zeitstempel: Long)

class Bankkonto(val inhaber: String, anfangsbetrag: Double):
  require(anfangsbetrag >= 0, "Anfangsbetrag muss positiv sein")
  
  private var _kontostand: Double = anfangsbetrag
  private var _historie: List[Transaktion] = List(
    Transaktion("Eröffnung", anfangsbetrag, System.currentTimeMillis())
  )
  
  def kontostand: Double = _kontostand
  def historie: List[Transaktion] = _historie
  
  def einzahlen(betrag: Double): Unit =
    require(betrag > 0, "Betrag muss positiv sein")
    _kontostand += betrag
    _historie = _historie :+ Transaktion("Einzahlung", betrag, System.currentTimeMillis())
  
  def abheben(betrag: Double): Boolean =
    if betrag > 0 && betrag <= _kontostand then
      _kontostand -= betrag
      _historie = _historie :+ Transaktion("Abhebung", -betrag, System.currentTimeMillis())
      true
    else
      false
  
  def kontoauszug(): String =
    val header = s"=== Konto von $inhaber ===\n"
    val transaktionen = _historie.map { t =>
      f"${t.art}%-12s: ${t.betrag}%+10.2f EUR"
    }.mkString("\n")
    val footer = f"\n${"="*30}\nKontostand: $_kontostand%10.2f EUR"
    header + transaktionen + footer

// Test
val konto = Bankkonto("Max", 100.0)
konto.einzahlen(50.0)
konto.abheben(30.0)
println(konto.kontoauszug())
```

### Lösung 4.7 – Spielkarten
```scala
enum Farbe:
  case Kreuz, Pik, Herz, Karo
  
  def symbol: Char = this match
    case Kreuz => '♣'
    case Pik => '♠'
    case Herz => '♥'
    case Karo => '♦'

enum Wert(val punkte: Int):
  case Zwei extends Wert(2)
  case Drei extends Wert(3)
  case Vier extends Wert(4)
  case Fuenf extends Wert(5)
  case Sechs extends Wert(6)
  case Sieben extends Wert(7)
  case Acht extends Wert(8)
  case Neun extends Wert(9)
  case Zehn extends Wert(10)
  case Bube extends Wert(10)
  case Dame extends Wert(10)
  case Koenig extends Wert(10)
  case Ass extends Wert(11)

case class Karte(farbe: Farbe, wert: Wert):
  override def toString = s"${wert} ${farbe.symbol}"

object Kartendeck:
  def vollstaendig: List[Karte] =
    for
      farbe <- Farbe.values.toList
      wert <- Wert.values.toList
    yield Karte(farbe, wert)
  
  def mischen(karten: List[Karte]): List[Karte] =
    scala.util.Random.shuffle(karten)

// Test
val deck = Kartendeck.vollstaendig
println(s"Deck hat ${deck.length} Karten")  // 52
val gemischt = Kartendeck.mischen(deck)
println(gemischt.take(5))
```

### Lösung 4.8 – Todo-Liste
```scala
enum Prioritaet:
  case Hoch, Mittel, Niedrig

case class Todo(
  id: Int,
  titel: String,
  beschreibung: String,
  prioritaet: Prioritaet,
  erledigt: Boolean = false
)

class TodoListe:
  private var todos: List[Todo] = List.empty
  private var naechsteId: Int = 1
  
  def hinzufuegen(titel: String, beschreibung: String, prio: Prioritaet): Todo =
    val todo = Todo(naechsteId, titel, beschreibung, prio)
    todos = todos :+ todo
    naechsteId += 1
    todo
  
  def erledigen(id: Int): Boolean =
    todos.find(_.id == id) match
      case Some(todo) =>
        todos = todos.map(t => if t.id == id then t.copy(erledigt = true) else t)
        true
      case None => false
  
  def loeschen(id: Int): Boolean =
    val vorher = todos.length
    todos = todos.filterNot(_.id == id)
    todos.length < vorher
  
  def alleAnzeigen(): Unit =
    todos.foreach { t =>
      val status = if t.erledigt then "✓" else " "
      println(f"[$status] ${t.id}%3d | ${t.prioritaet}%-7s | ${t.titel}")
    }
  
  def offene(): List[Todo] = todos.filterNot(_.erledigt)
  
  def nachPrioritaet(): List[Todo] =
    todos.sortBy(t => t.prioritaet.ordinal)

// Test
val liste = TodoListe()
liste.hinzufuegen("Scala lernen", "Vorlesung 4", Prioritaet.Hoch)
liste.hinzufuegen("Einkaufen", "Milch und Brot", Prioritaet.Niedrig)
liste.alleAnzeigen()
```

---

## Teil D: Verständnisfragen

### Lösung 4.6
```scala
case class Box(var inhalt: Int)

val b1 = Box(10)
val b2 = b1.copy()
b1.inhalt = 20

println(b1.inhalt)  // 20
println(b2.inhalt)  // 10 (copy erstellt neue Instanz!)
println(b1 == b2)   // false (unterschiedlicher inhalt)
```

`copy()` erstellt eine neue Instanz mit den gleichen Werten zum Zeitpunkt des Kopierens.

### Lösung 4.7
```scala
class Geheimnis(geheim: String):
  def vergleiche(other: Geheimnis): Boolean =
    this.geheim == other.geheim  // FEHLER!
```

**Problem:** `geheim` ist in `other` private und nicht zugänglich.

**Korrektur 1:** Parameter als `val`:
```scala
class Geheimnis(val geheim: String):
  def vergleiche(other: Geheimnis): Boolean =
    this.geheim == other.geheim  // OK
```

**Korrektur 2:** Case Class verwenden:
```scala
case class Geheimnis(geheim: String)
// == funktioniert automatisch!
```

### Lösung 4.8
```scala
object Singleton       // Normales Singleton
case object CaseSingleton  // Case Object
```

**Unterschied:**
- `case object` hat `toString`, `hashCode`, funktioniert mit Pattern Matching
- Wird oft für Enum-ähnliche Konstrukte verwendet

---

## Bonusaufgaben

### Bonus 1 – Komplexe Zahlen
```scala
case class Complex(real: Double, imag: Double):
  def +(other: Complex): Complex = 
    Complex(real + other.real, imag + other.imag)
  
  def -(other: Complex): Complex = 
    Complex(real - other.real, imag - other.imag)
  
  def *(other: Complex): Complex =
    Complex(
      real * other.real - imag * other.imag,
      real * other.imag + imag * other.real
    )
  
  def abs: Double = math.sqrt(real * real + imag * imag)
  
  def konjugiert: Complex = Complex(real, -imag)
  
  override def toString = 
    if imag >= 0 then f"$real%.2f + ${imag}%.2fi"
    else f"$real%.2f - ${-imag}%.2fi"

object Complex:
  val i: Complex = Complex(0, 1)
  val zero: Complex = Complex(0, 0)
  val one: Complex = Complex(1, 0)

// Test
val z1 = Complex(3, 4)
val z2 = Complex(1, -2)
println(z1 + z2)     // 4.00 + 2.00i
println(z1 * z2)     // 11.00 - 2.00i
println(z1.abs)      // 5.0
```

### Bonus 2 – Binärbaum
```scala
sealed trait Baum[+A]
case object Leer extends Baum[Nothing]
case class Knoten[A](wert: A, links: Baum[A], rechts: Baum[A]) extends Baum[A]

object Baum:
  def einfuegen[A](baum: Baum[A], wert: A)(using ord: Ordering[A]): Baum[A] =
    baum match
      case Leer => Knoten(wert, Leer, Leer)
      case Knoten(v, l, r) =>
        if ord.lt(wert, v) then Knoten(v, einfuegen(l, wert), r)
        else Knoten(v, l, einfuegen(r, wert))
  
  def enthaelt[A](baum: Baum[A], wert: A)(using ord: Ordering[A]): Boolean =
    baum match
      case Leer => false
      case Knoten(v, l, r) =>
        if ord.equiv(wert, v) then true
        else if ord.lt(wert, v) then enthaelt(l, wert)
        else enthaelt(r, wert)
  
  def inorder[A](baum: Baum[A]): List[A] =
    baum match
      case Leer => Nil
      case Knoten(v, l, r) => inorder(l) ++ List(v) ++ inorder(r)

// Test
var baum: Baum[Int] = Leer
for n <- List(5, 3, 7, 1, 4, 6, 8) do
  baum = Baum.einfuegen(baum, n)
println(Baum.inorder(baum))  // List(1, 3, 4, 5, 6, 7, 8)
```
