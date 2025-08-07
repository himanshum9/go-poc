package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// User represents a user object
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// App wraps the DB connection
type App struct {
	DB *sql.DB
}

// Initialize DB connection
func (a *App) InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=yourpass dbname=yourdb sslmode=disable"
	var err error
	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err = a.DB.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}
	log.Println("DB connected ✅")
}

// GET /users
func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := a.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

// POST /users
func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := a.DB.QueryRow(
		"INSERT INTO users(name, email) VALUES($1, $2) RETURNING id",
		u.Name, u.Email,
	).Scan(&u.ID)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}

// Main function
func main() {
	app := &App{}
	app.InitDB()
	defer app.DB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/users", app.getUsers).Methods("GET")
	r.HandleFunc("/users", app.createUser).Methods("POST")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
