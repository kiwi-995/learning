# Übung 7: Pattern Matching & ADTs

**Thema:** Pattern Matching, Algebraische Datentypen, Extraktoren

**Geschätzte Zeit:** 90-120 Minuten

---

## Teil A: Theoretische Fragen

### Frage 7.1
Was ist ein Algebraischer Datentyp (ADT)? Erkläre Sum Types und Product Types.

### Frage 7.2
Was bedeutet "Exhaustive Matching"? Warum ist `sealed` wichtig dafür?

### Frage 7.3
Was ist der Unterschied zwischen einem Pattern mit Variable und dem Wildcard `_`?

### Frage 7.4
Was macht ein Guard in einem Pattern Match? Gib ein Beispiel.

---

## Teil B: REPL-Übungen

### Aufgabe 7.1 – Grundlegendes Matching
```scala
def klassifiziere(x: Any): String = x match
  case 0 => "Null"
  case s: String => s"String: $s"
  case i: Int if i > 0 => "Positive Zahl"
  case i: Int => "Negative Zahl"
  case _ => "Unbekannt"

// Teste mit verschiedenen Werten
```

### Aufgabe 7.2 – Tuple Matching
```scala
def beschreibePunkt(p: (Int, Int)): String = p match
  case (0, 0) => ???
  case (x, 0) => ???
  case (0, y) => ???
  case (x, y) if x == y => ???
  case (x, y) => ???
```

### Aufgabe 7.3 – List Matching
```scala
def listenInfo[A](l: List[A]): String = l match
  case Nil => ???
  case x :: Nil => ???
  case x :: y :: Nil => ???
  case x :: y :: _ => ???
```

---

## Teil C: Programmieraufgaben

### Aufgabe 7.4 – Optionale Werte
Implementiere sichere Operationen:

```scala
def sicheresDivision(a: Int, b: Int): Option[Int] = ???
def wurzel(n: Double): Option[Double] = ???  // None für negative Zahlen
def ersteZahl(text: String): Option[Int] = ???  // Erste Zahl im Text

// Kombiniere die Ergebnisse
def berechne(a: String, b: String): Option[Int] =
  ???  // Parse beide Strings und dividiere
```

### Aufgabe 7.5 – Either für Fehlerbehandlung
```scala
sealed trait ValidierungsFehler
case class ZuKurz(min: Int) extends ValidierungsFehler
case class ZuLang(max: Int) extends ValidierungsFehler
case class UngueltigesZeichen(zeichen: Char) extends ValidierungsFehler

def validierePasswort(pw: String): Either[ValidierungsFehler, String] = ???
// - Mindestens 8 Zeichen
// - Maximal 20 Zeichen
// - Nur Buchstaben und Zahlen

// Teste mit verschiedenen Passwörtern
```

### Aufgabe 7.6 – Binärbaum ⭐
Implementiere einen binären Suchbaum:

```scala
sealed trait Baum[+A]
case object Leer extends Baum[Nothing]
case class Knoten[A](wert: A, links: Baum[A], rechts: Baum[A]) extends Baum[A]

object Baum:
  def einfuegen[A: Ordering](baum: Baum[A], wert: A): Baum[A] = ???
  def suchen[A: Ordering](baum: Baum[A], wert: A): Boolean = ???
  def inOrder[A](baum: Baum[A]): List[A] = ???
  def hoehe[A](baum: Baum[A]): Int = ???
  def anzahl[A](baum: Baum[A]): Int = ???
```

### Aufgabe 7.7 – Einfacher Taschenrechner ⭐
```scala
sealed trait Ausdruck
case class Zahl(n: Double) extends Ausdruck
case class Variable(name: String) extends Ausdruck
case class Plus(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Minus(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Mal(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Geteilt(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Potenz(basis: Ausdruck, exponent: Ausdruck) extends Ausdruck

object Rechner:
  type Umgebung = Map[String, Double]
  
  def auswerten(expr: Ausdruck, env: Umgebung): Either[String, Double] = ???
  def ableiten(expr: Ausdruck, variable: String): Ausdruck = ???
  def vereinfachen(expr: Ausdruck): Ausdruck = ???
  def formatieren(expr: Ausdruck): String = ???
```

### Aufgabe 7.8 – Mini-Interpreter ⭐⭐
```scala
sealed trait Befehl
case class Setze(variable: String, wert: Ausdruck) extends Befehl
case class Drucke(ausdruck: Ausdruck) extends Befehl
case class Wenn(bedingung: Ausdruck, dann: List[Befehl], sonst: List[Befehl]) extends Befehl
case class Wiederhole(anzahl: Int, befehle: List[Befehl]) extends Befehl

object Interpreter:
  type Zustand = Map[String, Double]
  
  def ausfuehren(programme: List[Befehl], zustand: Zustand): Zustand = ???
```

---

## Teil D: Extraktoren

### Aufgabe 7.9 – Email-Extraktor
```scala
object Email:
  def unapply(s: String): Option[(String, String)] = ???

// Test:
"test@example.com" match
  case Email(user, domain) => s"User: $user, Domain: $domain"
  case _ => "Keine Email"
```

### Aufgabe 7.10 – Datum-Extraktor
```scala
object Datum:
  def unapply(s: String): Option[(Int, Int, Int)] = ???
  // Erkennt "DD.MM.YYYY" oder "YYYY-MM-DD"

"25.12.2024" match
  case Datum(tag, monat, jahr) => s"$tag. $monat $jahr"
  case _ => "Kein Datum"
```

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – Lambda-Kalkül
```scala
sealed trait Term
case class Var(name: String) extends Term
case class App(f: Term, arg: Term) extends Term
case class Lam(param: String, body: Term) extends Term

// Beispiel: λx. λy. x y  (K-Kombinator)
val K = Lam("x", Lam("y", App(Var("x"), Var("y"))))

def freeVars(t: Term): Set[String] = ???
def substitute(t: Term, x: String, s: Term): Term = ???
def betaReduce(t: Term): Term = ???
```

### Bonus 2 – Regex-Matcher
```scala
sealed trait Regex
case object Empty extends Regex        // ""
case class Lit(c: Char) extends Regex  // einzelnes Zeichen
case class Seq(r1: Regex, r2: Regex) extends Regex  // r1 gefolgt von r2
case class Alt(r1: Regex, r2: Regex) extends Regex  // r1 oder r2
case class Star(r: Regex) extends Regex  // r beliebig oft

def matches(regex: Regex, input: String): Boolean = ???

// Test: (a|b)*c matched "aabc", "bc", "c", aber nicht "ab"
val pattern = Seq(Star(Alt(Lit('a'), Lit('b'))), Lit('c'))
```

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-07.md](../loesungen/loesung-07.md)
