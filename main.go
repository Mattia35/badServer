package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
	_ "github.com/mattn/go-sqlite3"
	api "github.com/Mattia35/badServer/backend/api"
	"github.com/Mattia35/badServer/backend/api/requestContext"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"strconv"
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
			department INTEGER NOT NULL,
			position TEXT NOT NULL,
			project INTEGER,
			PRIMARY KEY (id),
			FOREIGN KEY (project) REFERENCES project(id),
			FOREIGN KEY (department) REFERENCES department(id)
        );`,

        `CREATE TABLE IF NOT EXISTS department (
            id INTEGER NOT NULL,
			name TEXT NOT NULL,
			manager INTEGER,
			PRIMARY KEY (id),
			FOREIGN KEY (manager) REFERENCES employee(id)
        );`,

		`CREATE TABLE IF NOT EXISTS project (
			id INTEGER NOT NULL,
			name TEXT NOT NULL,
			start_date DATE NOT NULL,
			end_date DATE NOT NULL,
			budget REAL NOT NULL,
			department INTEGER NOT NULL UNIQUE,
			PRIMARY KEY (id)
			FOREIGN KEY (department) REFERENCES department(id)
				ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS profile (
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			PRIMARY KEY (username, password)
		);`,

		`CREATE TABLE IF NOT EXISTS token (
			username TEXT NOT NULL,
			token TEXT NOT NULL,
			session INTEGER NOT NULL,
			PRIMARY KEY (username, token)
			FOREIGN KEY (username) REFERENCES profile(username)
				ON DELETE CASCADE
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


    // API endpoint
	// login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		api.LoginHandler(db, w, r)
	})
	// get employees data
	http.HandleFunc("/:profile/employees", WithRequestContext(api.GetEmployeesData))
}


// Avvia il server HTTP
func startServer() {
    addr := "localhost:8080"
    fmt.Println("Server avviato su", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}

func WithRequestContext(handler func(*sql.DB, http.ResponseWriter, *http.Request, reqcontext.RequestContext)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID, _ := uuid.NewV4()

		// Estrai il token dall'header Authorization
		token := r.Header.Get("Authorization")

		// Estrai la sessione da un header (es: "X-Session")
		sessionStr := r.Header.Get("X-Session")
		session := 0
		if sessionStr != "" {
			if parsedSession, err := strconv.Atoi(sessionStr); err == nil {
				session = parsedSession
			}
		}

		// Costruisci il logger
		logger := logrus.WithField("req_id", reqID.String())
		if token != "" {
			logger = logger.WithField("token", token)
		}
		if sessionStr != "" {
			logger = logger.WithField("session", session)
		}

		// Crea il contesto
		ctx := reqcontext.RequestContext{
			ReqUUID: reqID,
			Token:   token,
			Session: session,
			Logger:  logger,
		}

		// Esegui l'handler
		handler(db, w, r, ctx)
	}
}


