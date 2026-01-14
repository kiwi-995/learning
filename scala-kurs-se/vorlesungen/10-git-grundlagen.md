# Vorlesung 10: Git Grundlagen

## Lernziele

Nach dieser Vorlesung kannst du:
- Ein Git-Repository erstellen und klonen
- Änderungen stagen und committen
- Den Verlauf mit `git log` ansehen
- Branches erstellen und wechseln
- Einfache Merges durchführen

---

## 10.1 Was ist Git?

**Git** ist ein **verteiltes Versionskontrollsystem**:

```
┌─────────────────────────────────────────────────────────────┐
│                     Vorteile von Git                        │
├─────────────────────────────────────────────────────────────┤
│ ✓ Vollständige Historie aller Änderungen                    │
│ ✓ Arbeiten ohne Netzwerk möglich (verteilt)                 │
│ ✓ Branching und Merging einfach                             │
│ ✓ Jeder hat vollständige Kopie                              │
│ ✓ Industriestandard                                         │
└─────────────────────────────────────────────────────────────┘
```

---

## 10.2 Repository erstellen

### Neues Repository

```bash
# In bestehendem Ordner
cd mein-projekt
git init

# Neuen Ordner erstellen
git init neues-projekt
cd neues-projekt
```

### Repository klonen

```bash
git clone https://github.com/user/repo.git
git clone https://github.com/user/repo.git mein-ordner
```

---

## 10.3 Die drei Bereiche

```
┌──────────────────┐    git add     ┌──────────────────┐   git commit   ┌──────────────────┐
│                  │ ─────────────> │                  │ ─────────────> │                  │
│  Working Dir     │                │  Staging Area    │                │   Repository     │
│  (Arbeitskopie)  │ <───────────── │  (Index)         │                │   (.git)         │
│                  │    checkout    │                  │                │                  │
└──────────────────┘                └──────────────────┘                └──────────────────┘
```

- **Working Directory**: Deine Dateien, die du bearbeitest
- **Staging Area**: Änderungen, die im nächsten Commit sein werden
- **Repository**: Die gespeicherte Historie

---

## 10.4 Grundlegende Befehle

### Status prüfen

```bash
git status
```

Zeigt:
- Geänderte Dateien
- Gestage Änderungen
- Untracked Dateien

### Änderungen stagen

```bash
# Eine Datei
git add datei.scala

# Mehrere Dateien
git add src/Main.scala src/Utils.scala

# Alle Änderungen
git add .

# Interaktiv auswählen
git add -p
```

### Commit erstellen

```bash
# Mit Message
git commit -m "Beschreibung der Änderung"

# Editor öffnen für längere Message
git commit

# Stage + Commit (nur bestehende Dateien)
git commit -am "Beschreibung"
```

### Gute Commit-Messages

```
┌─────────────────────────────────────────────────────────────┐
│                Commit-Message Format                        │
├─────────────────────────────────────────────────────────────┤
│ Kurze Zusammenfassung (max 50 Zeichen)                      │
│                                                             │
│ Optionaler längerer Text nach einer Leerzeile.              │
│ Erklärt WARUM, nicht WAS (das sieht man im Code).          │
│                                                             │
│ - Aufzählung möglich                                        │
│ - Imperativ verwenden ("Add feature" nicht "Added")         │
└─────────────────────────────────────────────────────────────┘
```

---

## 10.5 Historie ansehen

### Log

```bash
# Standard
git log

# Kompakt
git log --oneline

# Mit Graph
git log --oneline --graph --all

# Bestimmte Datei
git log -- datei.scala

# Letzten 5 Commits
git log -5
```

### Änderungen anzeigen

```bash
# Unstaged Änderungen
git diff

# Staged Änderungen
git diff --staged

# Zwischen Commits
git diff abc123 def456

# Bestimmte Datei
git diff datei.scala
```

---

## 10.6 Branches

### Branch-Konzept

```
         feature
            │
    ┌───┬───┼───┬───┐
    │   │   │   │   │
────●───●───●───●───●──────  main
    │       │
    │       └─ Commits auf feature
    └───────── Commits auf main
```

### Branch-Befehle

```bash
# Alle Branches anzeigen
git branch

# Neuen Branch erstellen
git branch feature-xyz

# Branch wechseln
git checkout feature-xyz
# oder (moderner)
git switch feature-xyz

# Erstellen und wechseln
git checkout -b feature-xyz
# oder
git switch -c feature-xyz

# Branch löschen
git branch -d feature-xyz
```

---

## 10.7 Merge

### Fast-Forward Merge

Wenn der Zielbranch keine neuen Commits hat:

```
Vorher:                  Nachher:
    feature                  feature
        │                        │
────●───●                ────●───●
    │                        │
   main                     main
```

```bash
git checkout main
git merge feature
```

### 3-Way Merge

Wenn beide Branches neue Commits haben:

```
Vorher:                  Nachher:
        feature                  feature
            │                        │
────●───●───●            ────●───●───●───●  (Merge-Commit)
    │                        │       │
    ●───●                    ●───●───┘
        │                        │
       main                     main
```

```bash
git checkout main
git merge feature
```

---

## 10.8 Merge-Konflikte

Wenn Git nicht automatisch mergen kann:

```
<<<<<<< HEAD
// Deine Änderung
val x = 10
=======
// Ihre Änderung
val x = 20
>>>>>>> feature
```

### Konflikte lösen

1. Datei öffnen
2. Konfliktmarker finden (`<<<<<<<`, `=======`, `>>>>>>>`)
3. Entscheiden, was bleiben soll
4. Marker entfernen
5. `git add datei.scala`
6. `git commit`

---

## 10.9 Nützliche Befehle

### Änderungen rückgängig

```bash
# Unstaged Änderungen verwerfen
git checkout -- datei.scala
git restore datei.scala  # moderner

# Staged Änderungen unstagen
git reset HEAD datei.scala
git restore --staged datei.scala  # moderner

# Letzten Commit rückgängig (behält Änderungen)
git reset --soft HEAD~1

# Letzten Commit rückgängig (verwirft Änderungen)
git reset --hard HEAD~1
```

### Temporär speichern (Stash)

```bash
# Änderungen stashen
git stash

# Mit Beschreibung
git stash save "Experimentelles Feature"

# Liste anzeigen
git stash list

# Wiederherstellen
git stash pop       # Entfernt aus Stash
git stash apply     # Behält in Stash
```

---

## 10.10 .gitignore

Dateien vom Tracking ausschließen:

```gitignore
# Kompilierte Dateien
*.class
target/

# IDE-Dateien
.idea/
.vscode/
*.iml

# OS-Dateien
.DS_Store
Thumbs.db

# Lokale Konfiguration
.env
secrets.conf

# Abhängigkeiten
node_modules/
```

---

## Zusammenfassung

| Befehl | Beschreibung |
|--------|--------------|
| `git init` | Repository erstellen |
| `git clone <url>` | Repository klonen |
| `git status` | Status anzeigen |
| `git add <datei>` | Änderungen stagen |
| `git commit -m "..."` | Commit erstellen |
| `git log --oneline` | Historie anzeigen |
| `git branch <name>` | Branch erstellen |
| `git checkout <name>` | Branch wechseln |
| `git merge <name>` | Branch mergen |
| `git stash` | Änderungen temporär speichern |

---

## Nächste Schritte

In **Vorlesung 11** lernen wir:
- Remote Repositories
- Push, Pull, Fetch
- Git Workflows für Teams

➡️ [Weiter zu Vorlesung 11: Git Workflows](./11-git-workflows.md)
