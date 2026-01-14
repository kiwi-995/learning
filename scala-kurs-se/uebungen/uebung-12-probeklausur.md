# Übung 12: Klausurvorbereitung - Probeklausur

**Zeit:** 90 Minuten  
**Erlaubte Hilfsmittel:** Keine

---

## Aufgabe 1: Grundlagen (15 Punkte)

### a) Typen und Werte (5 Punkte)
Was sind die Typen und Werte der folgenden Ausdrücke?

```scala
val a = List(1, 2, 3).map(_ * 2)
val b = if true then "ja" else 5
val c = (1, "hello", List(1,2))
val d = Some(42).filter(_ > 50)
val e = List("a", "b").flatMap(s => List(s, s.toUpperCase))
```

### b) val vs var (5 Punkte)
Erkläre den Unterschied. Warum ist `val` in funktionaler Programmierung zu bevorzugen?

### c) Pattern Matching (5 Punkte)
Schreibe Patterns für:
1. Eine leere Liste
2. Eine Liste mit genau einem Element
3. Eine Liste mit mindestens zwei Elementen

---

## Aufgabe 2: Rekursion (20 Punkte)

### a) Normale Rekursion (10 Punkte)
Implementiere eine Funktion `summe(n: Int): Int`, die die Summe aller Zahlen von 1 bis n berechnet.

```scala
def summe(n: Int): Int = ???

// summe(5) = 1+2+3+4+5 = 15
```

### b) Tail-Rekursion (10 Punkte)
Schreibe die Funktion aus a) als tail-rekursive Version mit `@tailrec`.

```scala
import scala.annotation.tailrec

def summeTail(n: Int): Int = ???
```

---

## Aufgabe 3: ADT und Pattern Matching (20 Punkte)

Definiere einen ADT für eine einfache Programmiersprache mit:
- Zahlen (Int)
- Variablen (String)
- Addition zweier Ausdrücke
- Multiplikation zweier Ausdrücke

```scala
sealed trait Ausdruck
// ???
```

Implementiere dann:

```scala
def auswerten(expr: Ausdruck, env: Map[String, Int]): Int = ???
```

Bei unbekannten Variablen soll eine Exception geworfen werden.

---

## Aufgabe 4: Collections (15 Punkte)

Gegeben:
```scala
case class Student(name: String, punkte: Int, fach: String)

val studenten = List(
  Student("Alice", 85, "Info"),
  Student("Bob", 72, "Mathe"),
  Student("Charlie", 91, "Info"),
  Student("Diana", 68, "Mathe"),
  Student("Eve", 79, "Info")
)
```

Schreibe **einen Ausdruck** (mit Methodenverkettung oder For-Comprehension), der:

a) Alle Informatik-Studenten mit über 80 Punkten findet und deren Namen zurückgibt

b) Den Durchschnitt aller Punkte berechnet

c) Studenten nach Fach gruppiert

---

## Aufgabe 5: Higher-Order Functions (15 Punkte)

### a) Eigenes filter (10 Punkte)
Implementiere `meinFilter` selbst (ohne die eingebaute Methode):

```scala
def meinFilter[A](liste: List[A])(bedingung: A => Boolean): List[A] = ???
```

### b) Currying (5 Punkte)
Was ist der Unterschied zwischen:
```scala
def f1(a: Int, b: Int): Int = a + b
def f2(a: Int)(b: Int): Int = a + b
```

Gib ein Beispiel, wo Version 2 praktischer ist.

---

## Aufgabe 6: Git (15 Punkte)

### a) Befehle (10 Punkte)
Schreibe die Git-Befehle für:
1. Repository klonen von `https://github.com/user/repo.git`
2. Neuen Branch `feature/login` erstellen und wechseln
3. Datei `Login.scala` zum Commit hinzufügen
4. Commit mit Nachricht "Implement login"
5. Branch zum Remote pushen

### b) Workflow (5 Punkte)
Erkläre kurz den Feature Branch Workflow.

---

## Musterlösungen

### Aufgabe 1a
```scala
val a = List(1, 2, 3).map(_ * 2)  // List[Int] = List(2, 4, 6)
val b = if true then "ja" else 5  // "ja": String (Int würde auch passen → Any)
val c = (1, "hello", List(1,2))   // (Int, String, List[Int])
val d = Some(42).filter(_ > 50)   // None: Option[Int]
val e = List("a", "b").flatMap(s => List(s, s.toUpperCase))  
// List[String] = List("a", "A", "b", "B")
```

### Aufgabe 2
```scala
def summe(n: Int): Int =
  if n <= 0 then 0
  else n + summe(n - 1)

@tailrec
def summeTail(n: Int, acc: Int = 0): Int =
  if n <= 0 then acc
  else summeTail(n - 1, acc + n)
```

### Aufgabe 3
```scala
sealed trait Ausdruck
case class Zahl(n: Int) extends Ausdruck
case class Variable(name: String) extends Ausdruck
case class Plus(l: Ausdruck, r: Ausdruck) extends Ausdruck
case class Mal(l: Ausdruck, r: Ausdruck) extends Ausdruck

def auswerten(expr: Ausdruck, env: Map[String, Int]): Int = expr match
  case Zahl(n) => n
  case Variable(name) => env.getOrElse(name, throw new Exception(s"Unbekannt: $name"))
  case Plus(l, r) => auswerten(l, env) + auswerten(r, env)
  case Mal(l, r) => auswerten(l, env) * auswerten(r, env)
```

### Aufgabe 4
```scala
// a)
studenten.filter(s => s.fach == "Info" && s.punkte > 80).map(_.name)

// b)
studenten.map(_.punkte).sum.toDouble / studenten.length

// c)
studenten.groupBy(_.fach)
```

### Aufgabe 5a
```scala
def meinFilter[A](liste: List[A])(bedingung: A => Boolean): List[A] = liste match
  case Nil => Nil
  case kopf :: rest =>
    if bedingung(kopf) then kopf :: meinFilter(rest)(bedingung)
    else meinFilter(rest)(bedingung)
```

### Aufgabe 6a
```bash
git clone https://github.com/user/repo.git
git checkout -b feature/login
git add Login.scala
git commit -m "Implement login"
git push -u origin feature/login
```
