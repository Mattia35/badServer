package myserver

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
)


var db *sql.DB
func main() {
    // 1. Inizializza database e crea tabelle
    initDB()
    // 2. Configura rotte
    setupRoutes()
    // 3. Avvia server
    startServer()
}


// Inizializza la connessione al database e crea tabelle
func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./app.db")
    if err != nil {
        log.Fatal("Errore apertura database:", err)
    }

    // Crea le tabelle se non esistono
    createTables()
}

// Crea le tabelle necessarie
func createTables() {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS employee (
            id INTEGER NOT NULL,
            name TEXT NOT NULL,
            surname TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			phone TEXT NOT NULL UNIQUE,
			address TEXT NOT NULL,
			birth_date DATE NOT NULL,
			hire_date DATE NOT NULL,
			salary REAL NOT NULL,
			department TEXT NOT NULL,
			position TEXT NOT NULL,
			PRIMARY KEY (id)
        );`,

        `CREATE TABLE IF NOT EXISTS department (
            id INTEGER NOT NULL,
			name TEXT NOT NULL,
			manager_id INTEGER,
			PRIMARY KEY (id),
			FOREIGN KEY (manager_id) REFERENCES employee(id)
        );`,

		`CREATE TABLE IF NOT EXISTS project (
			id INTEGER NOT NULL,
			name TEXT NOT NULL,
			start_date DATE NOT NULL,
			end_date DATE NOT NULL,
			budget REAL NOT NULL,
			department_id INTEGER NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (department_id) REFERENCES department(id)
		);`,
    }


    for _, query := range queries {
        _, err := db.Exec(query)
        if err != nil {
            log.Fatal("Errore creazione tabella:", err)
        }
    }
}


// Configura tutte le rotte
func setupRoutes() {
    // Serve frontend statico
    fs := http.FileServer(http.Dir("./frontend"))
    http.Handle("/", fs)
    // API di esempio
    http.HandleFunc("/api/ping", pingHandler)
}


// Avvia il server HTTP
func startServer() {
    addr := "localhost:8080"
    fmt.Println("Server avviato su", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}


// Handler API di esempio
func pingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // Crea la risposta da inviare
    response := map[string]string{"message": "pong"}
    // Serializza la risposta in JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Scrivi la risposta JSON
    w.Write(jsonResponse)
}