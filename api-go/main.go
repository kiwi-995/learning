/*
Minimal Todo API — Go / net/http + database/sql
=================================================
Run:  go run .
Test: curl http://localhost:8080/todos

Key concepts demonstrated:
  • Routing             — http.HandleFunc with manual method checking
  • JSON encoding       — json.NewEncoder / json.NewDecoder
  • Path parsing        — strings.TrimPrefix to extract IDs from URLs
  • Database            — database/sql with prepared statements
  • Error handling      — Go's explicit if err != nil pattern
  • Struct tags         — `json:"..."` controls JSON field names
  • No framework!       — Everything is standard library (+ SQLite driver)

This is MORE code than Python/Ruby — and that's the point.
Go makes you handle every detail explicitly, which means fewer surprises
in production. You'll appreciate this clarity in larger projects.
*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "modernc.org/sqlite" // Pure-Go SQLite driver (no CGo needed)
)

// ─── Data Model ─────────────────────────────────────────────────
// Struct tags (`json:"..."`) tell the JSON encoder what field names to use.
// `omitempty` means the field is omitted from JSON if it's the zero value.

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// For create/update requests, we use pointers so we can tell the
// difference between "not provided" (nil) and "set to zero value".
type TodoInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

// ─── Global DB handle ───────────────────────────────────────────
var db *sql.DB

func main() {
	var err error

	// Open SQLite database (creates file if it doesn't exist)
	db, err = sql.Open("sqlite", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id    INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT    NOT NULL,
			done  BOOLEAN NOT NULL DEFAULT 0
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// ─── Routing ────────────────────────────────────────────────
	// Go's standard library doesn't have a router with method matching,
	// so we handle it manually. This teaches you what routers actually do!

	http.HandleFunc("/todos", todosHandler)   // matches /todos
	http.HandleFunc("/todos/", todoHandler)   // matches /todos/123

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ─── Collection handler: GET /todos, POST /todos ────────────────
func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

// ─── Single-item handler: GET/PUT/DELETE /todos/:id ─────────────
func todoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL: "/todos/42" → "42"
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid ID"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodo(w, r, id)
	case http.MethodPut:
		updateTodo(w, r, id)
	case http.MethodDelete:
		deleteTodo(w, r, id)
	default:
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

// ─── Handlers ───────────────────────────────────────────────────

func listTodos(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, title, done FROM todos")
	if err != nil {
		http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []Todo{} // empty slice, not nil (so JSON returns [] not null)
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			http.Error(w, `{"error":"Scan error"}`, http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	writeJSON(w, http.StatusOK, todos)
}

func getTodo(w http.ResponseWriter, _ *http.Request, id int) {
	var t Todo
	err := db.QueryRow("SELECT id, title, done FROM todos WHERE id = ?", id).
		Scan(&t.ID, &t.Title, &t.Done)
	if err == sql.ErrNoRows {
		http.Error(w, `{"error":"Todo not found"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var input TodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}
	if input.Title == nil || *input.Title == "" {
		http.Error(w, `{"error":"title is required"}`, http.StatusUnprocessableEntity)
		return
	}

	done := false
	if input.Done != nil {
		done = *input.Done
	}

	result, err := db.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", *input.Title, done)
	if err != nil {
		http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	todo := Todo{ID: int(id), Title: *input.Title, Done: done}
	writeJSON(w, http.StatusCreated, todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request, id int) {
	// First check if the todo exists
	var t Todo
	err := db.QueryRow("SELECT id, title, done FROM todos WHERE id = ?", id).
		Scan(&t.ID, &t.Title, &t.Done)
	if err == sql.ErrNoRows {
		http.Error(w, `{"error":"Todo not found"}`, http.StatusNotFound)
		return
	}

	var input TodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Apply partial updates
	if input.Title != nil {
		t.Title = *input.Title
	}
	if input.Done != nil {
		t.Done = *input.Done
	}

	_, err = db.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?", t.Title, t.Done, id)
	if err != nil {
		http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func deleteTodo(w http.ResponseWriter, _ *http.Request, id int) {
	result, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, `{"error":"Todo not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 — no body
}

// ─── Helper ─────────────────────────────────────────────────────
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
