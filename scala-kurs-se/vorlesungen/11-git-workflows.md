# Vorlesung 11: Git Workflows & Zusammenarbeit

## Lernziele

Nach dieser Vorlesung kannst du:
- Mit Remote Repositories arbeiten
- Push, Pull und Fetch verstehen
- Den Feature Branch Workflow anwenden
- Pull Requests verstellen und reviewen
- Merge-Konflikte in Teams lösen

---

## 11.1 Remote Repositories

### Remote hinzufügen

```bash
git remote add origin https://github.com/user/repo.git
git remote -v  # Remotes anzeigen
```

### Push (Hochladen)

```bash
# Erster Push
git push -u origin main

# Danach
git push

# Bestimmter Branch
git push origin feature-xyz
```

### Pull (Herunterladen + Merge)

```bash
git pull
git pull origin main
```

### Fetch (Nur herunterladen)

```bash
git fetch
git fetch origin

# Dann manuell mergen
git merge origin/main
```

---

## 11.2 Feature Branch Workflow ⭐

Der **Standard-Workflow** für Teamprojekte:

```
1. Vom main-Branch starten
2. Feature-Branch erstellen  
3. Änderungen machen und committen
4. Push zum Remote
5. Pull Request erstellen
6. Code Review
7. Merge in main
```

### Schritt für Schritt

```bash
# 1. Aktuellen main holen
git checkout main
git pull

# 2. Feature-Branch erstellen
git checkout -b feature/benutzer-login

# 3. Entwickeln und committen
# ... Änderungen ...
git add .
git commit -m "Implementiere Login-Formular"

# ... weitere Änderungen ...
git add .
git commit -m "Füge Validierung hinzu"

# 4. Push zum Remote
git push -u origin feature/benutzer-login

# 5. Pull Request auf GitHub/GitLab erstellen
# 6. Code Review durch Teammitglieder
# 7. Merge (auf GitHub/GitLab oder lokal)
```

---

## 11.3 Pull Requests

### Was ist ein Pull Request?

```
┌─────────────────────────────────────────────────────────────┐
│                      Pull Request                           │
├─────────────────────────────────────────────────────────────┤
│ Titel: Feature: Benutzer-Login implementieren               │
│                                                             │
│ Beschreibung:                                               │
│ - Login-Formular hinzugefügt                                │
│ - Validierung implementiert                                 │
│ - Session-Verwaltung ergänzt                                │
│                                                             │
│ Reviewer: @alice, @bob                                      │
│ Status: Bereit zum Review                                   │
└─────────────────────────────────────────────────────────────┘
```

### Gute PR-Beschreibung

```markdown
## Zusammenfassung
Kurze Beschreibung der Änderungen.

## Änderungen
- Punkt 1
- Punkt 2

## Tests
- [ ] Unit-Tests hinzugefügt
- [ ] Manuell getestet

## Screenshots (falls relevant)
...
```

---

## 11.4 Code Reviews

### Als Reviewer

```
┌─────────────────────────────────────────────────────────────┐
│                    Code Review Checkliste                   │
├─────────────────────────────────────────────────────────────┤
│ ✓ Verständlicher Code?                                      │
│ ✓ Keine offensichtlichen Bugs?                              │
│ ✓ Tests vorhanden?                                          │
│ ✓ Edge Cases berücksichtigt?                                │
│ ✓ Konsistenter Stil?                                        │
│ ✓ Keine doppelter Code?                                     │
│ ✓ Commit-Messages sinnvoll?                                 │
└─────────────────────────────────────────────────────────────┘
```

### Feedback geben

- **Konstruktiv** formulieren
- **Konkret** sein ("In Zeile 42..." statt "Irgendwo...")
- Zwischen **Muss** und **Vorschlag** unterscheiden
- **Fragen** stellen statt befehlen

---

## 11.5 Rebase vs Merge

### Merge

```
          A───B───C  feature
         /         \
────●───●───●───●───●  main (Merge-Commit)
```

### Rebase

```
                    A'──B'──C'  feature
                   /
────●───●───●───●  main
```

```bash
# Feature-Branch auf main rebasen
git checkout feature
git rebase main

# Dann mergen (Fast-Forward)
git checkout main
git merge feature
```

### Wann was?

| Situation | Empfehlung |
|-----------|------------|
| Öffentliche Branches | Merge |
| Lokale Feature-Branches | Rebase |
| Linearen Verlauf gewünscht | Rebase |
| Merge-Commits explizit zeigen | Merge |

---

## 11.6 Konflikte bei Zusammenarbeit

### Szenario

```
Alice und Bob bearbeiten beide README.md

Alice:
- Pushed zuerst

Bob:
- Will pushen, aber:
  "Updates were rejected because the remote contains work..."
```

### Lösung

```bash
# Bobs Workflow
git pull --rebase  # Holt Alice' Änderungen, rebaset darauf

# Bei Konflikt:
# 1. Konflikt lösen
# 2. git add README.md
# 3. git rebase --continue

git push
```

---

## 11.7 Git Workflow-Modelle

### GitHub Flow (einfach)

```
main ────●────●────●────●────●────●
              │              │
              └───feature────┘
```

- Immer von `main` branchen
- Kurze Feature-Branches
- Häufige Pulls und Merges

### Gitflow (komplex)

```
main      ────────────●────────────●
                     /            /
release         ●───●            /
               /                /
develop  ●────●────●────●────●─●
              │         │
              └─feature─┘
```

Für Projekte mit geplanten Releases.

---

## 11.8 Best Practices

### Commits

```
✓ Kleine, fokussierte Commits
✓ Aussagekräftige Messages
✓ Ein logischer Schritt pro Commit
✗ "WIP" oder "Misc" als Message
✗ Riesige Commits mit vielen Änderungen
```

### Branches

```
✓ Beschreibende Namen: feature/user-login
✓ Kurzlebig (2-3 Tage ideal)
✗ Lange lebende Feature-Branches
✗ Vage Namen: "fix", "test", "stuff"
```

### Zusammenarbeit

```
✓ Regelmäßig pullen
✓ Code Reviews ernst nehmen
✓ Kommunizieren bei Konflikten
✗ Force-Push auf shared Branches
✗ main direkt bearbeiten
```

---

## 11.9 Nützliche Befehle

```bash
# Letzten Commit ändern
git commit --amend

# Branch-Tracking setzen
git branch --set-upstream-to=origin/main main

# Remote-Branches löschen
git push origin --delete feature-xyz

# Lokale Referenzen zu Remote aufräumen
git fetch --prune

# Cherry-Pick (einzelnen Commit übernehmen)
git cherry-pick abc1234

# Blame (wer hat was wann geändert)
git blame datei.scala
```

---

## Zusammenfassung

| Konzept | Beschreibung |
|---------|--------------|
| Remote | Entferntes Repository (z.B. GitHub) |
| Push | Lokale Commits hochladen |
| Pull | Remote-Änderungen herunterladen + mergen |
| Fetch | Remote-Änderungen nur herunterladen |
| Pull Request | Anfrage zum Mergen mit Review |
| Code Review | Begutachtung von Änderungen |
| Rebase | Commits auf neue Basis setzen |

---

## Nächste Schritte

In **Vorlesung 12** wiederholen wir alles und bereiten uns auf die Klausur vor!

➡️ [Weiter zu Vorlesung 12: Klausurvorbereitung](./12-klausurvorbereitung.md)
