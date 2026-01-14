# Übung 4: Klassen & Objekte

**Thema:** Klassen, Objects, Companion Objects, Case Classes, Enums

**Geschätzte Zeit:** 90-120 Minuten

---

## Teil A: Theoretische Fragen

### Frage 4.1
Was ist der Unterschied zwischen einer `class` und einem `object` in Scala?

### Frage 4.2
Erkläre den Unterschied zwischen diesen drei Klassendefinitionen:

```scala
class A(x: Int)
class B(val x: Int)
class C(var x: Int)
```

### Frage 4.3
Was ist ein Companion Object? Welche Vorteile bietet es?

### Frage 4.4
Was ist eine Case Class? Welche Methoden werden automatisch generiert?

### Frage 4.5
Wann würdest du eine Case Class vs. eine normale Klasse verwenden?

---

## Teil B: REPL-Übungen

### Aufgabe 4.1 – Einfache Klasse
Definiere in der REPL:

```scala
class Auto(val marke: String, val baujahr: Int):
  def alter: Int = 2024 - baujahr
  override def toString = s"$marke ($baujahr)"

// Erstelle Instanzen und teste die Methoden
```

### Aufgabe 4.2 – Object
Erstelle ein Utility-Object:

```scala
object StringUtils:
  def istPalindrom(s: String): Boolean = ???
  def anzahlVokale(s: String): Int = ???

// Teste mit verschiedenen Strings
```

### Aufgabe 4.3 – Case Class
Experimentiere mit Case Classes:

```scala
case class Student(name: String, matrikelNr: Int, fach: String)

val s1 = Student("Max", 12345, "Informatik")
val s2 = Student("Max", 12345, "Informatik")

// Teste: s1 == s2, s1.toString, s1.copy(fach = "Mathematik")
```

---

## Teil C: Programmieraufgaben

### Aufgabe 4.4 – Bruchrechnung
Erstelle eine Klasse für Brüche:

```scala
case class Bruch(zaehler: Int, nenner: Int):
  require(nenner != 0, "Nenner darf nicht 0 sein")
  
  def plus(other: Bruch): Bruch = ???
  def minus(other: Bruch): Bruch = ???
  def mal(other: Bruch): Bruch = ???
  def durch(other: Bruch): Bruch = ???
  def kuerzen: Bruch = ???
  
  override def toString = s"$zaehler/$nenner"

object Bruch:
  def ggt(a: Int, b: Int): Int = if b == 0 then a.abs else ggt(b, a % b)
```

**Teste mit:**
```scala
val b1 = Bruch(1, 2)
val b2 = Bruch(1, 4)
println(b1.plus(b2))  // 3/4
println(b1.mal(b2))   // 1/8
```

### Aufgabe 4.5 – Geometrie
Modelliere geometrische Formen mit Case Classes:

```scala
case class Punkt(x: Double, y: Double):
  def distanzZu(other: Punkt): Double = ???

case class Kreis(mittelpunkt: Punkt, radius: Double):
  def flaeche: Double = ???
  def umfang: Double = ???
  def enthaelt(p: Punkt): Boolean = ???

case class Rechteck(eckeUntenLinks: Punkt, breite: Double, hoehe: Double):
  def flaeche: Double = ???
  def umfang: Double = ???
  def eckeObenRechts: Punkt = ???
  def enthaelt(p: Punkt): Boolean = ???
```

### Aufgabe 4.6 – Bankkonto mit History
Erweitere das Bankkonto mit einer Transaktionshistorie:

```scala
case class Transaktion(art: String, betrag: Double, zeitstempel: Long)

class Bankkonto(val inhaber: String, anfangsbetrag: Double):
  private var _kontostand: Double = anfangsbetrag
  private var _historie: List[Transaktion] = List(
    Transaktion("Eröffnung", anfangsbetrag, System.currentTimeMillis())
  )
  
  def kontostand: Double = _kontostand
  def historie: List[Transaktion] = _historie
  
  def einzahlen(betrag: Double): Unit = ???
  def abheben(betrag: Double): Boolean = ???
  def kontoauszug(): String = ???
```

### Aufgabe 4.7 – Spielkarten ⭐
Modelliere ein Kartenspiel mit Enums:

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
  // ... weitere Werte
  case Ass extends Wert(11)

case class Karte(farbe: Farbe, wert: Wert):
  override def toString = s"${wert} ${farbe.symbol}"

object Kartendeck:
  def vollstaendig: List[Karte] = ???  // Alle 52 Karten
  def mischen(karten: List[Karte]): List[Karte] = ???
```

### Aufgabe 4.8 – Todo-Liste ⭐
Erstelle ein Todo-System:

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
  
  def hinzufuegen(titel: String, beschreibung: String, prio: Prioritaet): Todo = ???
  def erledigen(id: Int): Boolean = ???
  def loeschen(id: Int): Boolean = ???
  def alleAnzeigen(): Unit = ???
  def offene(): List[Todo] = ???
  def nachPrioritaet(): List[Todo] = ???
```

---

## Teil D: Verständnisfragen

### Frage 4.6
Was passiert bei diesem Code? Erkläre die Ausgabe.

```scala
case class Box(var inhalt: Int)

val b1 = Box(10)
val b2 = b1.copy()
b1.inhalt = 20

println(b1.inhalt)  // ?
println(b2.inhalt)  // ?
println(b1 == b2)   // ?
```

### Frage 4.7
Warum funktioniert dieser Code nicht?

```scala
class Geheimnis(geheim: String):
  def vergleiche(other: Geheimnis): Boolean =
    this.geheim == other.geheim  // FEHLER!
```

Wie würdest du es korrigieren?

### Frage 4.8
Was ist der Unterschied zwischen:

```scala
object Singleton

case object CaseSingleton
```

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – Komplexe Zahlen
Implementiere komplexe Zahlen:

```scala
case class Complex(real: Double, imag: Double):
  def +(other: Complex): Complex = ???
  def -(other: Complex): Complex = ???
  def *(other: Complex): Complex = ???
  def abs: Double = ???  // Betrag
  def konjugiert: Complex = ???
  
  override def toString = 
    if imag >= 0 then s"$real + ${imag}i"
    else s"$real - ${-imag}i"

object Complex:
  val i: Complex = Complex(0, 1)
  val zero: Complex = Complex(0, 0)
  val one: Complex = Complex(1, 0)
```

### Bonus 2 – Binärbaum
Modelliere einen binären Suchbaum:

```scala
sealed trait Baum[+A]
case object Leer extends Baum[Nothing]
case class Knoten[A](wert: A, links: Baum[A], rechts: Baum[A]) extends Baum[A]

object Baum:
  def einfuegen[A](baum: Baum[A], wert: A)(using ord: Ordering[A]): Baum[A] = ???
  def enthaelt[A](baum: Baum[A], wert: A)(using ord: Ordering[A]): Boolean = ???
  def inorder[A](baum: Baum[A]): List[A] = ???
```

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-04.md](../loesungen/loesung-04.md)
