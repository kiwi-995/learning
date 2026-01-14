# Lösungen zu Übung 1

---

## Teil A: Theoretische Fragen

### Lösung 1.1
Scala ist eine Multi-Paradigma-Sprache, weil sie zwei Programmierparadigmen unterstützt:
1. **Objektorientierte Programmierung (OOP)**: Alles ist ein Objekt, Klassen, Vererbung
2. **Funktionale Programmierung (FP)**: Funktionen als First-Class Values, Unveränderlichkeit

### Lösung 1.2
- `val` definiert eine **unveränderliche** Variable (wie `final` in Java)
- `var` definiert eine **veränderliche** Variable

**Warum `val` bevorzugen:**
- Einfacher zu verstehen (Wert ändert sich nie)
- Thread-sicher (keine Race Conditions)
- Ermöglicht funktionalen Programmierstil
- Compiler kann besser optimieren

### Lösung 1.3
Typinferenz bedeutet, dass der Compiler den Typ einer Variable automatisch erkennt:

```scala
val zahl = 42        // Compiler erkennt: Int
val text = "Hallo"   // Compiler erkennt: String
val pi = 3.14        // Compiler erkennt: Double
```

### Lösung 1.4
```scala
@main def greet(): Unit =
  println("Hello!")
```

Der Rückgabetyp ist `Unit`. `Unit` entspricht `void` in Java/C – es bedeutet, dass die Funktion keinen sinnvollen Wert zurückgibt, sondern nur einen Seiteneffekt hat (hier: Konsolenausgabe).

### Lösung 1.5
**Ausdruck (Expression):** Hat immer einen Wert
```scala
val x = 1 + 1        // 1 + 1 ist ein Ausdruck mit Wert 2
val max = if (a > b) a else b  // if/else ist ein Ausdruck
```

**Anweisung (Statement):** Führt eine Aktion aus, kein Rückgabewert (in Scala selten)
```scala
println("Hallo")     // Hat Rückgabewert Unit (quasi kein Wert)
var x = 0; x = 5     // Zuweisung hat Rückgabewert Unit
```

In Scala sind die meisten Konstrukte Ausdrücke – das ist ein funktionales Merkmal!

---

## Teil B: REPL-Übungen

### Lösung 1.1 – Einfache Berechnungen
```scala
scala> 2 + 3 * 4
val res0: Int = 14   // Punkt vor Strich

scala> (2 + 3) * 4
val res1: Int = 20   // Klammern ändern Reihenfolge

scala> 17 / 5
val res2: Int = 3    // Integer-Division! Nachkommastellen werden abgeschnitten

scala> 17.0 / 5
val res3: Double = 3.4  // Mit Double korrekte Division

scala> 17 % 5
val res4: Int = 2    // Modulo: Rest nach Division
```

### Lösung 1.2 – Variablen
```scala
scala> val vorname = "Max"
val vorname: String = Max

scala> val nachname = "Mustermann"
val nachname: String = Mustermann

scala> val alter = 20
val alter: Int = 20

scala> vorname = "Moritz"
-- Error: Reassignment to val vorname

scala> val vollerName = vorname + " " + nachname
val vollerName: String = Max Mustermann

scala> val geburtsjahr = 2024 - alter
val geburtsjahr: Int = 2004
```

### Lösung 1.3 – Typen erkunden
```scala
:type 42                  // Int
:type 3.14                // Double
:type "Hallo"             // String
:type true                // Boolean
:type List(1, 2, 3)       // List[Int]
:type (1, "text", true)   // (Int, String, Boolean) - ein Tupel
```

### Lösung 1.4 – Blöcke als Ausdrücke
```scala
val ergebnis = {
  val x = 10
  val y = 20
  x + y
}
```
- **Wert:** `30`
- **Typ:** `Int`

Der Block gibt den Wert der letzten Zeile zurück (`x + y = 30`).

---

## Teil C: Programmieraufgaben

### Lösung 1.5 – Hello World Variante
```scala
@main def willkommen(): Unit =
  println("Hallo, ich bin Max!")
  println("Aktuelles Semester: Wintersemester 2024/25")
  println("Ich freue mich auf Software Engineering!")
```

### Lösung 1.6 – Variablen-Übung
```scala
@main def variablenDemo(): Unit =
  // 1. Definiere eine unveränderliche Variable für dein Geburtsjahr
  val geburtsjahr = 2003
  
  // 2. Berechne dein Alter (als val)
  val alter = 2024 - geburtsjahr
  
  // 3. Definiere deinen Namen
  val name = "Max"
  
  // 4. Gib alles aus
  println(s"Hallo, ich bin $name!")
  println(s"Ich bin $alter Jahre alt.")
  println(s"Geboren im Jahr $geburtsjahr.")
```

### Lösung 1.7 – Rechteck-Berechnung
```scala
@main def rechteck(): Unit =
  val laenge = 5.0
  val breite = 3.0
  
  val flaeche = laenge * breite
  val umfang = 2 * laenge + 2 * breite
  
  println(s"Rechteck: $laenge x $breite")
  println(s"Fläche: $flaeche")
  println(s"Umfang: $umfang")
```

### Lösung 1.8 – Temperaturkonverter
```scala
@main def temperatur(): Unit =
  val celsius = 20.0
  
  val fahrenheit = celsius * 9.0 / 5.0 + 32
  val kelvin = celsius + 273.15
  
  println(s"$celsius°C entspricht:")
  println(s"- $fahrenheit°F (Fahrenheit)")
  println(s"- $kelvin K (Kelvin)")
```

---

## Teil D: Verständnisfragen

### Lösung 1.6
```scala
val x = 10
val y = {
  val x = 20  // Neues x im inneren Scope, überdeckt äußeres x
  x + 5       // Verwendet inneres x = 20
}
println(x)    // Ausgabe: 10 (äußeres x)
println(y)    // Ausgabe: 25 (20 + 5)
```

Das innere `val x = 20` **überdeckt (shadows)** das äußere `x`, aber nur innerhalb des Blocks.

### Lösung 1.7
```scala
val gruss = "Hallo"
gruss = "Tschüss"  // FEHLER: Reassignment to val
```

**Problem:** `val` kann nicht neu zugewiesen werden.

**Korrektur mit `var`:**
```scala
var gruss = "Hallo"
gruss = "Tschüss"  // OK
```

**Bessere Korrektur (funktionaler Stil):**
```scala
val gruss1 = "Hallo"
val gruss2 = "Tschüss"  // Neue Variable statt Neuzuweisung
```

### Lösung 1.8
```scala
// Variante A - Expliziter Typ
val zahl: Int = 42

// Variante B - Typinferenz
val zahl = 42
```

**Unterschied:** Funktional identisch! In beiden Fällen ist `zahl` vom Typ `Int`.

**Wann explizit:**
- Bei öffentlichen APIs (Dokumentation)
- Wenn Compiler falschen Typ ableitet
- Für bessere Lesbarkeit in komplexem Code

---

## Bonusaufgaben

### Bonus 1 – BMI-Rechner
```scala
@main def bmiRechner(): Unit =
  val gewicht = 75.0    // kg
  val groesse = 1.80    // m
  
  val bmi = gewicht / (groesse * groesse)
  
  println(s"Gewicht: $gewicht kg")
  println(s"Größe: $groesse m")
  println(f"BMI: $bmi%.1f")
```

### Bonus 2 – Projekt-Struktur

**Mathematik.scala:**
```scala
object Mathematik:
  def quadrat(n: Int): Int = n * n
```

**Main.scala:**
```scala
@main def start(): Unit =
  val eingabe = 5
  val ergebnis = Mathematik.quadrat(eingabe)
  println(s"Das Quadrat von $eingabe ist $ergebnis")
```

**Ausführen:**
```bash
scala-cli run .
```

**Ausgabe:**
```
Das Quadrat von 5 ist 25
```
