# Lösungen zu Übung 7

---

## Teil A: Theoretische Fragen

### Lösung 7.1
Ein **ADT** (Algebraischer Datentyp) modelliert Daten als typsichere Strukturen:

**Sum Type** ("oder"): Ein Wert ist EINER von mehreren Typen
```scala
sealed trait Ampel
case object Rot extends Ampel
case object Grün extends Ampel
```

**Product Type** ("und"): Ein Wert HAT mehrere Komponenten
```scala
case class Person(name: String, alter: Int)
```

### Lösung 7.2
**Exhaustive Matching:** Der Compiler stellt sicher, dass alle möglichen Fälle behandelt werden.

`sealed` ist wichtig, weil:
- Alle Subtypen müssen in der gleichen Datei definiert sein
- Compiler kennt alle Möglichkeiten
- Warnung bei fehlenden Cases

### Lösung 7.3
- **Variable** (`case x =>`) bindet den Wert an `x`, kann im Body verwendet werden
- **Wildcard** (`case _ =>`) fängt alles, bindet aber nicht

```scala
x match
  case n => println(n)  // n ist verfügbar
  case _ => println("?")  // _ kann nicht verwendet werden
```

### Lösung 7.4
Ein **Guard** ist eine zusätzliche `if`-Bedingung im Pattern:

```scala
person match
  case Person(name, alter) if alter >= 18 => "Erwachsen"
  case Person(name, _) => "Minderjährig"
```

---

## Teil B: REPL-Übungen

### Lösung 7.2
```scala
def beschreibePunkt(p: (Int, Int)): String = p match
  case (0, 0) => "Ursprung"
  case (x, 0) => s"Auf X-Achse bei x=$x"
  case (0, y) => s"Auf Y-Achse bei y=$y"
  case (x, y) if x == y => s"Auf Diagonale bei $x"
  case (x, y) => s"Punkt ($x, $y)"
```

### Lösung 7.3
```scala
def listenInfo[A](l: List[A]): String = l match
  case Nil => "Leere Liste"
  case x :: Nil => s"Ein Element: $x"
  case x :: y :: Nil => s"Zwei Elemente: $x und $y"
  case x :: y :: _ => s"Mindestens drei, beginnt mit $x, $y"
```

---

## Teil C: Programmieraufgaben

### Lösung 7.4 – Optionale Werte
```scala
def sicheresDivision(a: Int, b: Int): Option[Int] =
  if b == 0 then None else Some(a / b)

def wurzel(n: Double): Option[Double] =
  if n < 0 then None else Some(math.sqrt(n))

def ersteZahl(text: String): Option[Int] =
  text.split("\\D+").filter(_.nonEmpty).headOption.flatMap(_.toIntOption)

def berechne(a: String, b: String): Option[Int] =
  for
    x <- a.toIntOption
    y <- b.toIntOption
    result <- sicheresDivision(x, y)
  yield result

// Test
println(berechne("10", "2"))  // Some(5)
println(berechne("10", "0"))  // None
println(berechne("abc", "2")) // None
```

### Lösung 7.5 – Either für Fehlerbehandlung
```scala
sealed trait ValidierungsFehler
case class ZuKurz(min: Int) extends ValidierungsFehler
case class ZuLang(max: Int) extends ValidierungsFehler
case class UngueltigesZeichen(zeichen: Char) extends ValidierungsFehler

def validierePasswort(pw: String): Either[ValidierungsFehler, String] =
  if pw.length < 8 then
    Left(ZuKurz(8))
  else if pw.length > 20 then
    Left(ZuLang(20))
  else
    pw.find(c => !c.isLetterOrDigit) match
      case Some(c) => Left(UngueltigesZeichen(c))
      case None => Right(pw)

// Test
println(validierePasswort("abc"))          // Left(ZuKurz(8))
println(validierePasswort("password123"))  // Right(password123)
println(validierePasswort("pass@word"))    // Left(UngueltigesZeichen(@))
```

### Lösung 7.6 – Binärbaum
```scala
sealed trait Baum[+A]
case object Leer extends Baum[Nothing]
case class Knoten[A](wert: A, links: Baum[A], rechts: Baum[A]) extends Baum[A]

object Baum:
  def einfuegen[A: Ordering](baum: Baum[A], wert: A): Baum[A] =
    val ord = summon[Ordering[A]]
    baum match
      case Leer => Knoten(wert, Leer, Leer)
      case Knoten(w, l, r) =>
        if ord.lt(wert, w) then Knoten(w, einfuegen(l, wert), r)
        else Knoten(w, l, einfuegen(r, wert))
  
  def suchen[A: Ordering](baum: Baum[A], wert: A): Boolean =
    val ord = summon[Ordering[A]]
    baum match
      case Leer => false
      case Knoten(w, l, r) =>
        if ord.equiv(wert, w) then true
        else if ord.lt(wert, w) then suchen(l, wert)
        else suchen(r, wert)
  
  def inOrder[A](baum: Baum[A]): List[A] = baum match
    case Leer => Nil
    case Knoten(w, l, r) => inOrder(l) ++ List(w) ++ inOrder(r)
  
  def hoehe[A](baum: Baum[A]): Int = baum match
    case Leer => 0
    case Knoten(_, l, r) => 1 + math.max(hoehe(l), hoehe(r))
  
  def anzahl[A](baum: Baum[A]): Int = baum match
    case Leer => 0
    case Knoten(_, l, r) => 1 + anzahl(l) + anzahl(r)

// Test
var b: Baum[Int] = Leer
for n <- List(5, 3, 7, 1, 4, 6, 8) do
  b = Baum.einfuegen(b, n)
println(Baum.inOrder(b))   // List(1, 3, 4, 5, 6, 7, 8)
println(Baum.suchen(b, 4)) // true
println(Baum.hoehe(b))     // 3
```

### Lösung 7.7 – Taschenrechner
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
  
  def auswerten(expr: Ausdruck, env: Umgebung): Either[String, Double] = expr match
    case Zahl(n) => Right(n)
    case Variable(name) => 
      env.get(name).toRight(s"Variable '$name' nicht definiert")
    case Plus(l, r) => 
      for x <- auswerten(l, env); y <- auswerten(r, env) yield x + y
    case Minus(l, r) => 
      for x <- auswerten(l, env); y <- auswerten(r, env) yield x - y
    case Mal(l, r) => 
      for x <- auswerten(l, env); y <- auswerten(r, env) yield x * y
    case Geteilt(l, r) =>
      for
        x <- auswerten(l, env)
        y <- auswerten(r, env)
        result <- if y == 0 then Left("Division durch Null") else Right(x / y)
      yield result
    case Potenz(b, e) =>
      for x <- auswerten(b, env); y <- auswerten(e, env) yield math.pow(x, y)
  
  def formatieren(expr: Ausdruck): String = expr match
    case Zahl(n) => n.toString
    case Variable(name) => name
    case Plus(l, r) => s"(${formatieren(l)} + ${formatieren(r)})"
    case Minus(l, r) => s"(${formatieren(l)} - ${formatieren(r)})"
    case Mal(l, r) => s"(${formatieren(l)} * ${formatieren(r)})"
    case Geteilt(l, r) => s"(${formatieren(l)} / ${formatieren(r)})"
    case Potenz(b, e) => s"(${formatieren(b)} ^ ${formatieren(e)})"

// Test: (x + 2) * 3
val expr = Mal(Plus(Variable("x"), Zahl(2)), Zahl(3))
println(Rechner.formatieren(expr))                    // ((x + 2.0) * 3.0)
println(Rechner.auswerten(expr, Map("x" -> 5)))       // Right(21.0)
```

---

## Teil D: Extraktoren

### Lösung 7.9
```scala
object Email:
  def unapply(s: String): Option[(String, String)] =
    val teile = s.split("@")
    if teile.length == 2 && teile(0).nonEmpty && teile(1).contains(".") then
      Some((teile(0), teile(1)))
    else
      None

// Test
"test@example.com" match
  case Email(user, domain) => println(s"User: $user, Domain: $domain")
  case _ => println("Keine Email")
```

### Lösung 7.10
```scala
object Datum:
  def unapply(s: String): Option[(Int, Int, Int)] =
    // Format: DD.MM.YYYY
    val deutschFormat = """(\d{1,2})\.(\d{1,2})\.(\d{4})""".r
    // Format: YYYY-MM-DD
    val isoFormat = """(\d{4})-(\d{1,2})-(\d{1,2})""".r
    
    s match
      case deutschFormat(tag, monat, jahr) =>
        Some((tag.toInt, monat.toInt, jahr.toInt))
      case isoFormat(jahr, monat, tag) =>
        Some((tag.toInt, monat.toInt, jahr.toInt))
      case _ => None

// Test
println("25.12.2024" match
  case Datum(t, m, j) => s"$t. Monat $m im Jahr $j"
  case _ => "Kein Datum")

println("2024-12-25" match
  case Datum(t, m, j) => s"$t. Monat $m im Jahr $j"
  case _ => "Kein Datum")
```

---

## Bonusaufgaben

### Bonus 1 – Lambda-Kalkül
```scala
sealed trait Term
case class Var(name: String) extends Term
case class App(f: Term, arg: Term) extends Term
case class Lam(param: String, body: Term) extends Term

def freeVars(t: Term): Set[String] = t match
  case Var(name) => Set(name)
  case App(f, arg) => freeVars(f) ++ freeVars(arg)
  case Lam(param, body) => freeVars(body) - param

def substitute(t: Term, x: String, s: Term): Term = t match
  case Var(name) if name == x => s
  case Var(name) => Var(name)
  case App(f, arg) => App(substitute(f, x, s), substitute(arg, x, s))
  case Lam(param, body) if param == x => Lam(param, body)
  case Lam(param, body) => Lam(param, substitute(body, x, s))

// Test
val id = Lam("x", Var("x"))  // λx.x
println(freeVars(id))  // Set()
println(freeVars(Var("y")))  // Set(y)
```

### Bonus 2 – Regex-Matcher
```scala
sealed trait Regex
case object Empty extends Regex
case class Lit(c: Char) extends Regex
case class Seq(r1: Regex, r2: Regex) extends Regex
case class Alt(r1: Regex, r2: Regex) extends Regex
case class Star(r: Regex) extends Regex

def matches(regex: Regex, input: String): Boolean =
  def matchesPrefix(r: Regex, s: String): List[String] = r match
    case Empty => List(s)
    case Lit(c) => 
      if s.nonEmpty && s.head == c then List(s.tail)
      else Nil
    case Seq(r1, r2) =>
      matchesPrefix(r1, s).flatMap(rest => matchesPrefix(r2, rest))
    case Alt(r1, r2) =>
      matchesPrefix(r1, s) ++ matchesPrefix(r2, s)
    case Star(r) =>
      s :: matchesPrefix(r, s).flatMap(rest => 
        if rest != s then matchesPrefix(Star(r), rest) else Nil)
  
  matchesPrefix(regex, input).contains("")

// Test: (a|b)*c
val pattern = Seq(Star(Alt(Lit('a'), Lit('b'))), Lit('c'))
println(matches(pattern, "aabc"))  // true
println(matches(pattern, "bc"))    // true
println(matches(pattern, "c"))     // true
println(matches(pattern, "ab"))    // false
```
