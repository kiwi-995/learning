# Lösungen zu Übung 2

---

## Teil A: Theoretische Fragen

### Lösung 2.1
- `Int`: 32 Bit, Wertebereich ca. -2 bis +2 Milliarden
- `Long`: 64 Bit, viel größerer Wertebereich

`Long` verwenden bei:
- Zeitstempeln (Millisekunden seit 1970)
- Dateigrößen
- Großen IDs
- Wenn Int-Bereich nicht ausreicht

### Lösung 2.2
```scala
10 / 3    // = 3 (Integer-Division, Nachkommastellen abgeschnitten)
10.0 / 3  // = 3.333... (Double-Division, korrekt)
```

Bei Integer-Division wird immer abgerundet (Richtung 0).

### Lösung 2.3

**s-Interpolation** – Variablen einsetzen:
```scala
val name = "Max"
s"Hallo $name"  // "Hallo Max"
```

**f-Interpolation** – Formatierung:
```scala
val pi = 3.14159
f"Pi: $pi%.2f"  // "Pi: 3.14"
```

**raw-Interpolation** – Keine Escape-Sequenzen:
```scala
raw"Pfad: C:\neu"  // "Pfad: C:\neu" (Backslash bleibt)
```

### Lösung 2.4
Ein **Tupel** gruppiert mehrere Werte verschiedener Typen:
```scala
val person = ("Alice", 25, true)
person._2  // 25 (zweites Element, 1-basiert!)
```

Oder mit Dekonstruktion:
```scala
val (name, alter, aktiv) = person
alter  // 25
```

### Lösung 2.5
- **Unit**: Kein sinnvoller Rückgabewert (wie `void`). Verwendet für Seiteneffekte.
  ```scala
  def sagHallo(): Unit = println("Hallo")
  ```

- **Null**: Fehlender Referenzwert. Sollte vermieden werden, stattdessen `Option`.
  ```scala
  val x: String = null  // Schlecht!
  ```

- **Nothing**: Subtyp aller Typen. Für Funktionen, die nie normal zurückkehren.
  ```scala
  def fehler(): Nothing = throw new Exception()
  ```

---

## Teil B: REPL-Übungen

### Lösung 2.1
```scala
scala> val a = 2147483647
val a: Int = 2147483647

scala> a + 1
val res0: Int = -2147483648  // Überlauf! Wird negativ

scala> val b = 2147483647L
val b: Long = 2147483647

scala> b + 1
val res1: Long = 2147483648  // Korrekt!

scala> 17 / 5
val res2: Int = 3  // Integer-Division

scala> 17.0 / 5
val res3: Double = 3.4  // Double-Division

scala> 17 % 5
val res4: Int = 2  // Rest (Modulo)
```

### Lösung 2.2
```scala
scala> val text = "  Scala ist großartig!  "
scala> text.length       // 25
scala> text.trim         // "Scala ist großartig!"
scala> text.toUpperCase  // "  SCALA IST GROSSARTIG!  "
scala> text.contains("großartig")  // true
scala> text.replace("großartig", "super")  // "  Scala ist super!  "
scala> text.split(" ")   // Array("", "", "Scala", "ist", "großartig!", "", "")
```

### Lösung 2.3
```scala
// 1.
val name = "Alice"
val alter = 25
s"Mein Name ist $name und ich bin $alter Jahre alt."

// 2.
val pi = 3.14159265
f"Pi auf 4 Stellen: $pi%.4f"

// 3.
val preis = 9.99
f"Der Preis beträgt $preis%10.2f €"
```

### Lösung 2.4
```scala
val student = ("Max", "Informatik", 3)

// Zugriff
student._1  // "Max"
student._2  // "Informatik"
student._3  // 3

// Dekonstruktion
val (name, fach, semester) = student

// Neues Tupel (Tupel sind unveränderlich!)
val naechstesSemester = (student._1, student._2, student._3 + 1)
// oder
val naechstesSemester = (name, fach, semester + 1)
```

---

## Teil C: Programmieraufgaben

### Lösung 2.5 – Personendaten
```scala
@main def personDaten(): Unit =
  val vorname = "Maria"
  val nachname = "Müller"
  val alter = 22
  val groesse = 1.72
  val gewicht = 65.5
  
  val bmi = gewicht / (groesse * groesse)
  
  println("=== Personendaten ===")
  println(s"Name: $vorname $nachname")
  println(s"Alter: $alter Jahre")
  println(s"Größe: $groesse m")
  println(s"Gewicht: $gewicht kg")
  println(f"BMI: $bmi%.1f")
```

### Lösung 2.6 – Währungsrechner
```scala
@main def waehrung(): Unit =
  val euroBetrag = 100.0
  val kursUSD = 1.08
  val kursCHF = 0.96
  
  val dollar = euroBetrag * kursUSD
  val franken = euroBetrag * kursCHF
  
  println(f"$euroBetrag%.2f EUR entspricht:")
  println(f"  $dollar%.2f USD (Kurs: $kursUSD)")
  println(f"  $franken%.2f CHF (Kurs: $kursCHF)")
```

### Lösung 2.7 – Koordinaten
```scala
@main def koordinaten(): Unit =
  val punkt1 = (3.0, 4.0)
  val punkt2 = (6.0, 8.0)
  
  val (x1, y1) = punkt1
  val (x2, y2) = punkt2
  
  // Distanz
  val dx = x2 - x1
  val dy = y2 - y1
  val distanz = math.sqrt(dx * dx + dy * dy)
  
  // Mittelpunkt
  val mittelpunkt = ((x1 + x2) / 2, (y1 + y2) / 2)
  
  println(s"Punkt 1: $punkt1")
  println(s"Punkt 2: $punkt2")
  println(f"Distanz: $distanz%.2f")
  println(s"Mittelpunkt: $mittelpunkt")
```

### Lösung 2.8 – Zeitumrechnung
```scala
@main def zeit(): Unit =
  val gesamtSekunden = 3661
  
  val stunden = gesamtSekunden / 3600
  val restNachStunden = gesamtSekunden % 3600
  val minuten = restNachStunden / 60
  val sekunden = restNachStunden % 60
  
  println(s"$gesamtSekunden Sekunden = $stunden Stunde(n), $minuten Minute(n), $sekunden Sekunde(n)")
```

---

## Teil D: Verständnisfragen

### Lösung 2.6
```scala
val x = (1 + 2, "Hello" + " World", 3.14 > 2)
```

- **Ergebnis:** `(3, "Hello World", true)`
- **Typ:** `(Int, String, Boolean)`

### Lösung 2.7
```scala
val preis = 19.99
println(f"Preis: $preis%10.2f EUR")
```

Ausgabe: `Preis:      19.99 EUR`

- `%10.2f` bedeutet: Mindestbreite 10 Zeichen, 2 Dezimalstellen
- Der Wert wird rechtsbündig mit Leerzeichen aufgefüllt

### Lösung 2.8
```scala
val eingabe = "abc"
val zahl = eingabe.toInt  // FEHLER: NumberFormatException!
```

**Problem:** "abc" kann nicht in eine Zahl konvertiert werden.

**Sichere Lösung:**
```scala
val eingabe = "abc"
val zahl: Option[Int] = eingabe.toIntOption

zahl match
  case Some(n) => println(s"Die Zahl ist $n")
  case None => println("Ungültige Eingabe!")
```

---

## Bonusaufgaben

### Bonus 1 – Zeugnis-Formatierung
```scala
@main def zeugnis(): Unit =
  val noten = (("Mathematik", 2.3), ("Deutsch", 1.7), ("Englisch", 2.0))
  
  val (mathe, deutsch, englisch) = noten
  val durchschnitt = (mathe._2 + deutsch._2 + englisch._2) / 3
  
  println("═══════════════════════════════")
  println("       ZEUGNIS                 ")
  println("═══════════════════════════════")
  println(f"${mathe._1}%-15s ${mathe._2}%5.1f")
  println(f"${deutsch._1}%-15s ${deutsch._2}%5.1f")
  println(f"${englisch._1}%-15s ${englisch._2}%5.1f")
  println("───────────────────────────────")
  println(f"Durchschnitt:   $durchschnitt%5.1f")
  println("═══════════════════════════════")
```

### Bonus 2 – ASCII-Art Generator
```scala
@main def banner(): Unit =
  val name = "MAX"
  val breite = name.length + 8
  
  val rahmenOben = "╔" + "═" * breite + "╗"
  val rahmenUnten = "╚" + "═" * breite + "╝"
  val inhalt = f"║${name}%${(breite + name.length) / 2}s${" " * ((breite - name.length) / 2)}║"
  
  println(rahmenOben)
  println(inhalt)
  println(rahmenUnten)
```
