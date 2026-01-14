# Übung 1: Einführung in Scala

**Thema:** Installation, REPL, erste Programme, val vs var

**Geschätzte Zeit:** 60-90 Minuten

---

## Teil A: Theoretische Fragen

### Frage 1.1
Was bedeutet es, dass Scala eine "Multi-Paradigma-Sprache" ist? Nenne die zwei Hauptparadigmen.

### Frage 1.2
Erkläre den Unterschied zwischen `val` und `var`. Warum sollte man `val` bevorzugen?

### Frage 1.3
Was ist Typinferenz? Gib ein Beispiel, wo der Compiler den Typ automatisch erkennt.

### Frage 1.4
Was ist der Rückgabetyp der folgenden Funktion? Was bedeutet dieser Typ?

```scala
@main def greet(): ??? =
  println("Hello!")
```

### Frage 1.5
Was ist der Unterschied zwischen einem **Ausdruck** und einer **Anweisung** in Scala? Gib je ein Beispiel.

---

## Teil B: REPL-Übungen

Starte die Scala REPL (`scala-cli repl`) und führe folgende Aufgaben aus:

### Aufgabe 1.1 – Einfache Berechnungen
Berechne in der REPL:
- `2 + 3 * 4`
- `(2 + 3) * 4`
- `17 / 5` (Was fällt dir auf?)
- `17.0 / 5`
- `17 % 5` (Modulo-Operator)

### Aufgabe 1.2 – Variablen
Erstelle Variablen in der REPL:
```scala
val vorname = "Max"
val nachname = "Mustermann"
val alter = 20
```

Dann:
1. Versuche, `vorname` einen neuen Wert zuzuweisen. Was passiert?
2. Erstelle eine Variable `vollerName`, die Vor- und Nachname kombiniert
3. Berechne, in welchem Jahr die Person geboren wurde (aktuelles Jahr - Alter)

### Aufgabe 1.3 – Typen erkunden
Nutze `:type` in der REPL, um die Typen folgender Ausdrücke zu ermitteln:
- `42`
- `3.14`
- `"Hallo"`
- `true`
- `List(1, 2, 3)`
- `(1, "text", true)`

### Aufgabe 1.4 – Blöcke als Ausdrücke
Was ist der Wert und Typ von:
```scala
val ergebnis = {
  val x = 10
  val y = 20
  x + y
}
```

---

## Teil C: Programmieraufgaben

### Aufgabe 1.5 – Hello World Variante
Erstelle eine Datei `willkommen.scala`, die:
1. Deinen Namen ausgibt
2. Das aktuelle Semester ausgibt
3. Eine persönliche Nachricht ausgibt

Beispielausgabe:
```
Hallo, ich bin Max!
Aktuelles Semester: Wintersemester 2024/25
Ich freue mich auf Software Engineering!
```

### Aufgabe 1.6 – Variablen-Übung
Erstelle eine Datei `variablen.scala` mit folgendem Code und ergänze die fehlenden Teile:

```scala
@main def variablenDemo(): Unit =
  // 1. Definiere eine unveränderliche Variable für dein Geburtsjahr
  val geburtsjahr = ???
  
  // 2. Berechne dein Alter (als val)
  val alter = ???
  
  // 3. Definiere deinen Namen
  val name = ???
  
  // 4. Gib alles aus
  println(s"Hallo, ich bin $name!")
  println(s"Ich bin $alter Jahre alt.")
  println(s"Geboren im Jahr $geburtsjahr.")
```

### Aufgabe 1.7 – Rechteck-Berechnung
Schreibe ein Programm `rechteck.scala`, das:
1. Zwei `val`-Variablen für Länge und Breite definiert
2. Die Fläche berechnet (Länge × Breite)
3. Den Umfang berechnet (2 × Länge + 2 × Breite)
4. Beides ausgibt

### Aufgabe 1.8 – Temperaturkonverter ⭐
Schreibe ein Programm `temperatur.scala`, das:
1. Eine Temperatur in Celsius als `val` definiert
2. Sie in Fahrenheit umrechnet: `F = C × 9/5 + 32`
3. Sie in Kelvin umrechnet: `K = C + 273.15`
4. Alle drei Werte schön formatiert ausgibt

Beispielausgabe:
```
20.0°C entspricht:
- 68.0°F (Fahrenheit)
- 293.15 K (Kelvin)
```

---

## Teil D: Verständnisfragen

### Frage 1.6
Was passiert bei folgendem Code? Erkläre.

```scala
val x = 10
val y = {
  val x = 20
  x + 5
}
println(x)
println(y)
```

### Frage 1.7
Warum gibt folgender Code einen Fehler? Wie würdest du ihn korrigieren?

```scala
val gruss = "Hallo"
gruss = "Tschüss"
```

### Frage 1.8
Was ist der Unterschied zwischen diesen zwei Varianten?

```scala
// Variante A
val zahl: Int = 42

// Variante B
val zahl = 42
```

---

## Bonusaufgaben ⭐⭐

### Bonus 1 – BMI-Rechner
Schreibe ein Programm, das aus Gewicht (kg) und Größe (m) den BMI berechnet.
Formel: `BMI = Gewicht / (Größe × Größe)`

### Bonus 2 – Projekt-Struktur
Erstelle ein kleines Projekt mit zwei Dateien:
- `Mathematik.scala` – Ein Object mit einer Funktion `quadrat(n: Int): Int`
- `Main.scala` – Ruft die Funktion auf und gibt das Ergebnis aus

---

## Lösungen

Die Lösungen findest du in: [loesungen/loesung-01.md](../loesungen/loesung-01.md)
