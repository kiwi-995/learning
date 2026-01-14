# Übung 10-11: Git Grundlagen & Workflows

**Thema:** Git-Befehle, Branches, Merges, Remote Workflows

**Geschätzte Zeit:** 90 Minuten

---

## Teil A: Git Grundlagen

### Aufgabe 10.1 – Repository erstellen
Schreibe die Git-Befehle für folgende Schritte:

1. Erstelle ein neues Verzeichnis `mein-projekt`
2. Initialisiere ein Git-Repository
3. Erstelle eine Datei `README.md`
4. Stage die Datei
5. Erstelle den ersten Commit

### Aufgabe 10.2 – Status verstehen
Was zeigt `git status` in folgenden Situationen?

1. Direkt nach `git init`
2. Nach `git add datei.txt`
3. Nach `git commit`
4. Nach Änderung einer bereits committeten Datei

### Aufgabe 10.3 – Log lesen
Erkläre die Ausgabe:

```
$ git log --oneline
abc1234 (HEAD -> main) Füge Validierung hinzu
def5678 Implementiere Login
ghi9012 Initiales Setup
```

---

## Teil B: Branches und Merges

### Aufgabe 10.4 – Branch-Workflow
Schreibe die Befehle für:

1. Neuen Branch `feature/user-profile` erstellen
2. Zum Branch wechseln
3. Einige Commits machen
4. Zurück zu `main` wechseln
5. Feature-Branch mergen
6. Feature-Branch löschen

### Aufgabe 10.5 – Merge-Konflikt
Gegeben diese Situation:

```
Datei: config.scala

<<<<<<< HEAD
val maxUsers = 100
=======
val maxUsers = 50
>>>>>>> feature/limits
```

1. Was ist passiert?
2. Wie löst du den Konflikt?
3. Welche Befehle brauchst du danach?

---

## Teil C: Remote Workflows

### Aufgabe 10.6 – Remote einrichten
```bash
# 1. Remote hinzufügen
???

# 2. Branch pushen
???

# 3. Änderungen holen und mergen
???
```

### Aufgabe 10.7 – Feature Branch Workflow
Beschreibe den kompletten Ablauf von "neue Feature-Idee" bis "Merge in main":

1. ...
2. ...
3. ...

---

## Teil D: Praktische Übung

### Aufgabe 10.8 – Mini-Projekt ⭐
Führe folgende Schritte praktisch aus:

```bash
# 1. Neues Repository erstellen
mkdir git-uebung && cd git-uebung
git init

# 2. README erstellen und committen
echo "# Git Übung" > README.md
git add README.md
git commit -m "Initial commit"

# 3. Feature-Branch erstellen
git checkout -b feature/scala-code

# 4. Scala-Datei hinzufügen
cat > Main.scala << 'EOF'
@main def hello(): Unit =
  println("Hello from Git!")
EOF

git add Main.scala
git commit -m "Add main function"

# 5. Feature erweitern
cat >> Main.scala << 'EOF'

def add(a: Int, b: Int): Int = a + b
EOF

git add Main.scala
git commit -m "Add helper function"

# 6. Zurück zu main und mergen
git checkout main
git merge feature/scala-code

# 7. Log anzeigen
git log --oneline --graph
```

---

## Lösungen

### Lösung 10.1
```bash
mkdir mein-projekt
cd mein-projekt
git init
echo "# Mein Projekt" > README.md
git add README.md
git commit -m "Initialer Commit"
```

### Lösung 10.4
```bash
git branch feature/user-profile
git checkout feature/user-profile
# oder: git checkout -b feature/user-profile

# Änderungen machen...
git add .
git commit -m "Implementiere User Profile"

git checkout main
git merge feature/user-profile
git branch -d feature/user-profile
```

### Lösung 10.5
1. Zwei Branches haben dieselbe Zeile unterschiedlich geändert
2. Marker entfernen, gewünschten Wert behalten: `val maxUsers = 100`
3. 
```bash
git add config.scala
git commit  # Merge-Commit abschließen
```

### Lösung 10.6
```bash
# 1. Remote hinzufügen
git remote add origin https://github.com/user/repo.git

# 2. Branch pushen
git push -u origin main

# 3. Änderungen holen und mergen
git pull origin main
```

### Lösung 10.7
1. `git checkout main && git pull` - Aktuellen Stand holen
2. `git checkout -b feature/new-feature` - Branch erstellen
3. Code schreiben, regelmäßig committen
4. `git push -u origin feature/new-feature` - Push zum Remote
5. Pull Request auf GitHub erstellen
6. Code Review durch Teammitglieder
7. Nach Approval: Merge auf GitHub oder lokal
8. `git checkout main && git pull` - Aktualisieren
9. `git branch -d feature/new-feature` - Lokalen Branch löschen
