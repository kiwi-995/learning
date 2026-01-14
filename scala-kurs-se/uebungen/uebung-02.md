# Übung 2: Datentypen & Variablen

**Thema:** Datentypen, String-Interpolation, Tupel, Typkonvertierung

**Geschätzte Zeit:** 60-90 Minuten

---

## Teil A: Theoretische Fragen

### Frage 2.1
Was ist der Unterschied zwischen `Int` und `Long`? Wann würdest du `Long` verwenden?

### Frage 2.2
Erkläre den Unterschied zwischen `10 / 3` und `10.0 / 3`. Was ist das Ergebnis jeweils?

### Frage 2.3
Was ist String-Interpolation? Erkläre die drei Arten (`s`, `f`, `raw`) mit je einem Beispiel.

### Frage 2.4
Was ist ein Tupel? Wie greifst du auf das zweite Element eines Tupels zu?

### Frage 2.5
Was bedeuten `Unit`, `Null` und `Nothing` in Scala? Wann wird jeder Typ verwendet?

---

## Teil B: REPL-Übungen

### Aufgabe 2.1 – Numerische Typen
Führe folgende Berechnungen in der REPL aus:

```scala
val a = 2147483647  // Größter Int
a + 1               // Was passiert?

val b = 2147483647L  // Als Long
b + 1

17 / 5
17.0 / 5
17 % 5
```

### Aufgabe 2.2 – String-Operationen
Experimentiere mit Strings:

```scala
val text = "  Scala ist großartig!  "

text.length
text.trim
text.toUpperCase
text.contains("großartig")
text.replace("großartig", "super")
text.split(" ")
```

### Aufgabe 2.3 – String-Interpolation
Erstelle folgende Ausgaben mit String-Interpolation:

1. `"Mein Name ist Alice und ich bin 25 Jahre alt."`
2. `"Pi auf 4 Stellen: 3.1416"`
3. `"Der Preis beträgt     9,99 €"` (rechtsbündig mit 10 Zeichen)

### Aufgabe 2.4 – Tupel
Erstelle und manipuliere Tupel:

```scala
val student = ("Max", "Informatik", 3)

// Greife auf alle Elemente zu
// Dekonstruiere das Tupel
// Erstelle ein neues Tupel mit geändertem Semester
```

---

## Teil C: Programmieraufgaben

### Aufgabe 2.5 – Personendaten
Schreibe ein Programm `person.scala`:

```scala
@main def personDaten(): Unit =
  val vorname = "Maria"
  val nachname = "Müller"
  val alter = 22
  val groesse = 1.72  // in Metern
  val gewicht = 65.5  // in kg
  
  // 1. Berechne den BMI: gewicht / (groesse * groesse)
  // 2. Gib alle Daten formatiert aus mit String-Interpolation
  // 3. Formatiere den BMI auf 1 Dezimalstelle
```

Erwartete Ausgabe:
```
=== Personendaten ===
Name: Maria Müller
Alter: 22 Jahre
Größe: 1.72 m
Gewicht: 65.5 kg
BMI: 22.1
```

### Aufgabe 2.6 – Währungsrechner
Schreibe ein Programm `waehrung.scala`, das:
1. Einen Euro-Betrag speichert
2. In Dollar umrechnet (Kurs: 1 EUR = 1.08 USD)
3. In Schweizer Franken umrechnet (Kurs: 1 EUR = 0.96 CHF)
4. Alle Beträge formatiert auf 2 Dezimalstellen ausgibt

### Aufgabe 2.7 – Tupel-Übung
Schreibe ein Programm `koordinaten.scala`:

```scala
@main def koordinaten(): Unit =
  // Definiere zwei Punkte als Tupel (x, y)
  val punkt1 = (3.0, 4.0)
  val punkt2 = (6.0, 8.0)
  
  // 1. Berechne die Distanz zwischen den Punkten
  //    Formel: sqrt((x2-x1)² + (y2-y1)²)
  
  // 2. Berechne den Mittelpunkt
  //    Formel: ((x1+x2)/2, (y1+y2)/2)
  
  // 3. Gib alles aus
```

### Aufgabe 2.8 – Zeitumrechnung ⭐
Schreibe ein Programm `zeit.scala`:

1. Speichere eine Anzahl Sekunden (z.B. 3661)
2. Rechne um in Stunden, Minuten und Sekunden
3. Gib das Ergebnis formatiert aus

Beispiel: `3661 Sekunden = 1 Stunde, 1 Minute, 1 Sekunde`

**Tipp:** Verwende `/` für Division und `%` für den Rest.

---

## Teil D: Verständnisfragen

### Frage 2.6
Was ist das Ergebnis und der Typ von:

```scala
val x = (1 + 2, "Hello" + " World", 3.14 > 2)
```

### Frage 2.7
Erkläre, was bei folgendem Code passiert:

```scala
val preis = 19.99
println(f"Preis: $preis%10.2f EUR")
```

### Frage 2.8
Warum ist dieser Code problematisch?

```scala
val eingabe = "abc"
val zahl = eingabe.toInt
```

Wie würdest du es sicherer machen?

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – Zeugnis-Formatierung
Erstelle ein Programm, das Noten als Tupel speichert und formatiert ausgibt:

```scala
val noten = (("Mathematik", 2.3), ("Deutsch", 1.7), ("Englisch", 2.0))
```

Ausgabe:
```
═══════════════════════════════
       ZEUGNIS
═══════════════════════════════
Mathematik:         2.3
Deutsch:            1.7
Englisch:           2.0
───────────────────────────────
Durchschnitt:       2.0
═══════════════════════════════
```

### Bonus 2 – ASCII-Art Generator
Schreibe ein Programm, das einen Namen als "Banner" ausgibt:

```
Eingabe: "MAX"

Ausgabe:
╔═══════════════╗
║     MAX       ║
╚═══════════════╝
```

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-02.md](../loesungen/loesung-02.md)
