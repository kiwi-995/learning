# Vorlesung 1: Einführung in Scala

## Lernziele

Nach dieser Vorlesung kannst du:
- Erklären, was Scala ist und warum es verwendet wird
- Scala und SBT auf deinem Computer installieren
- Ein erstes Scala-Programm schreiben und ausführen
- Die Scala REPL für Experimente nutzen
- Den Unterschied zwischen `val` und `var` verstehen

---

## 1.1 Was ist Scala?

**Scala** (Scalable Language) ist eine moderne Programmiersprache, die zwei Paradigmen vereint:

- **Objektorientierte Programmierung (OOP)**: Alles ist ein Objekt
- **Funktionale Programmierung (FP)**: Funktionen sind First-Class Citizens

### Warum Scala lernen?

```
┌─────────────────────────────────────────────────────────────┐
│                    Vorteile von Scala                       │
├─────────────────────────────────────────────────────────────┤
│ ✓ Läuft auf der JVM (Java Virtual Machine)                  │
│ ✓ Volle Interoperabilität mit Java                          │
│ ✓ Prägnante, ausdrucksstarke Syntax                         │
│ ✓ Starkes Typsystem mit Typinferenz                         │
│ ✓ Funktionale Programmierung nativ unterstützt              │
│ ✓ Wird von Firmen wie Twitter, LinkedIn, Netflix genutzt    │
└─────────────────────────────────────────────────────────────┘
```

### Scala 3 vs Scala 2

Wir verwenden in diesem Kurs **Scala 3** (auch "Dotty" genannt). Scala 3 wurde 2021 veröffentlicht und bringt:

- Vereinfachte Syntax (weniger geschweifte Klammern nötig)
- Verbesserte Enums
- Union Types und Intersection Types
- Kontextuelle Abstraktionen (`given`/`using` statt `implicit`)

> **Hinweis**: Prof. Brachthäuser ist Co-Autor des offiziellen Scala 3 Buchs!

---

## 1.2 Installation

### Schritt 1: Scala CLI installieren

Die einfachste Methode ist **Scala CLI**, ein modernes Tool für Scala 3:

**macOS (mit Homebrew):**
```bash
brew install Virtuslab/scala-cli/scala-cli
```

**Linux/macOS (ohne Homebrew):**
```bash
curl -sSLf https://scala-cli.virtuslab.org/get | sh
```

**Windows (PowerShell):**
```powershell
irm https://scala-cli.virtuslab.org/get | iex
```

### Schritt 2: Installation testen

```bash
scala-cli version
```

Du solltest etwas wie `Scala CLI version 1.x.x` sehen.

### Schritt 3: REPL starten

```bash
scala-cli repl
```

Du siehst nun den interaktiven Scala-Prompt:
```
scala>
```

### Alternative: SBT (Scala Build Tool)

Für größere Projekte wird oft **SBT** verwendet:

```bash
# macOS
brew install sbt

# Oder via Coursier (empfohlen)
cs install sbt
```

---

## 1.3 Dein erstes Scala-Programm

### Hello World

Erstelle eine Datei `hello.scala`:

```scala
@main def hello(): Unit =
  println("Hallo, Scala!")
```

**Ausführen mit Scala CLI:**
```bash
scala-cli run hello.scala
```

**Ausgabe:**
```
Hallo, Scala!
```

### Erklärung des Codes

```scala
@main def hello(): Unit =
  println("Hallo, Scala!")
```

| Element | Bedeutung |
|---------|-----------|
| `@main` | Markiert die Funktion als Einstiegspunkt |
| `def` | Definiert eine Funktion |
| `hello` | Name der Funktion |
| `()` | Keine Parameter |
| `Unit` | Rückgabetyp (entspricht `void` in Java) |
| `=` | Der Funktionskörper folgt |
| `println(...)` | Gibt Text auf der Konsole aus |

### Alternative Schreibweise (mit Klammern)

```scala
@main def hello(): Unit = {
  println("Hallo, Scala!")
  println("Willkommen zum Kurs!")
}
```

In Scala 3 sind geschweifte Klammern bei Einrückung optional!

---

## 1.4 Die Scala REPL

Die **REPL** (Read-Eval-Print-Loop) ist perfekt zum Experimentieren:

```bash
scala-cli repl
```

### Beispiel-Session

```scala
scala> 1 + 1
val res0: Int = 2

scala> "Hallo" + " Welt"
val res1: String = Hallo Welt

scala> (1 to 10).sum
val res2: Int = 55

scala> List(1, 2, 3).map(_ * 2)
val res3: List[Int] = List(2, 4, 6)
```

**Nützliche REPL-Befehle:**
- `:help` – Zeigt alle Befehle
- `:quit` oder `:q` – Beendet die REPL
- `:reset` – Setzt die Session zurück
- `:type <expr>` – Zeigt den Typ eines Ausdrucks

---

## 1.5 Variablen: val vs var

Scala unterscheidet zwischen **unveränderlichen** und **veränderlichen** Variablen:

### val – Unveränderlich (Empfohlen!)

```scala
val name = "Alice"
val alter = 21

// Das geht NICHT:
// name = "Bob"  // Fehler: Reassignment to val
```

### var – Veränderlich (Sparsam verwenden!)

```scala
var zaehler = 0
zaehler = zaehler + 1  // OK
zaehler += 1           // OK (Kurzschreibweise)
```

### Warum val bevorzugen?

```
┌─────────────────────────────────────────────────────────────┐
│               Vorteile von Unveränderlichkeit               │
├─────────────────────────────────────────────────────────────┤
│ ✓ Einfacher zu verstehen (Wert ändert sich nie)             │
│ ✓ Thread-sicher (keine Race Conditions)                     │
│ ✓ Keine versteckten Seiteneffekte                           │
│ ✓ Compiler kann besser optimieren                           │
│ ✓ Funktionaler Stil fördert sauberen Code                   │
└─────────────────────────────────────────────────────────────┘
```

> **Regel**: Verwende immer `val`, es sei denn, du hast einen guten Grund für `var`.

---

## 1.6 Typinferenz

Scala kann Typen oft **automatisch erkennen**:

```scala
val zahl = 42           // Typ: Int (erkannt)
val text = "Hallo"      // Typ: String (erkannt)
val pi = 3.14159        // Typ: Double (erkannt)
val flag = true         // Typ: Boolean (erkannt)
```

Du kannst Typen auch **explizit angeben**:

```scala
val zahl: Int = 42
val text: String = "Hallo"
val pi: Double = 3.14159
val flag: Boolean = true
```

### Wann Typen explizit angeben?

- Bei öffentlichen Methoden (für Dokumentation)
- Wenn der Compiler den Typ nicht erkennen kann
- Wenn der abgeleitete Typ zu allgemein ist

```scala
// Explizit für Klarheit
def berechneGehalt(stunden: Int, satz: Double): Double =
  stunden * satz
```

---

## 1.7 Ausdrücke vs Anweisungen

**Wichtiges Konzept**: In Scala ist (fast) alles ein **Ausdruck** mit einem Wert!

### Beispiele

```scala
// if/else ist ein Ausdruck
val maximum = if (5 > 3) 5 else 3

// Block ist ein Ausdruck (letzte Zeile = Wert)
val ergebnis = {
  val a = 10
  val b = 20
  a + b  // Dieser Wert wird zurückgegeben: 30
}

// Auch Zuweisungen haben einen Wert (Unit)
val x = { val y = 5 }  // x hat Typ Unit
```

---

## 1.8 Projekt-Setup mit Scala CLI

Für mehrere Dateien erstelle ein Projektverzeichnis:

```
mein-projekt/
├── src/
│   ├── Main.scala
│   └── Hilfe.scala
└── project.scala
```

**project.scala** (Projektkonfiguration):
```scala
//> using scala 3.3
//> using option -deprecation
```

**src/Main.scala**:
```scala
@main def start(): Unit =
  val nachricht = Hilfe.begruessung("Scala-Lerner")
  println(nachricht)
```

**src/Hilfe.scala**:
```scala
object Hilfe:
  def begruessung(name: String): String =
    s"Willkommen, $name!"
```

**Ausführen:**
```bash
scala-cli run mein-projekt/
```

---

## 1.9 IDE-Setup

### IntelliJ IDEA (Empfohlen)

1. Lade [IntelliJ IDEA](https://www.jetbrains.com/idea/) herunter (Community Edition reicht)
2. Installiere das **Scala**-Plugin (Einstellungen → Plugins)
3. Öffne dein Projekt

### VS Code

1. Installiere [VS Code](https://code.visualstudio.com/)
2. Installiere die **Metals**-Extension
3. Öffne deinen Projektordner

---

## Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Scala | Multi-Paradigma-Sprache (OOP + FP) auf der JVM |
| `val` | Unveränderliche Variable (bevorzugt) |
| `var` | Veränderliche Variable (sparsam nutzen) |
| `@main` | Kennzeichnet den Programm-Einstiegspunkt |
| REPL | Interaktive Shell zum Experimentieren |
| Typinferenz | Compiler erkennt Typen automatisch |
| Ausdrücke | Fast alles hat einen Wert |

---

## Nächste Schritte

In **Vorlesung 2** lernen wir:
- Alle Datentypen im Detail
- String-Operationen und -Interpolation
- Tupel und erste Datenstrukturen

➡️ [Weiter zu Vorlesung 2: Datentypen & Variablen](./02-datentypen.md)

---

## Ressourcen

- [Offizielle Scala-Dokumentation](https://docs.scala-lang.org/)
- [Scala 3 Book](https://docs.scala-lang.org/scala3/book/introduction.html)
- [Scala CLI Dokumentation](https://scala-cli.virtuslab.org/)
