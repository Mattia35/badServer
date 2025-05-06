package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	api "github.com/Mattia35/badServer/backend/api"
	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func main() {
	initDB()
	defer db.Close()
	router := setupRoutes()
	startServer(router)
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
            name_surname TEXT NOT NULL,
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
			PRIMARY KEY (username)
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

func setupRoutes() *httprouter.Router {
	router := httprouter.New()

	// Serve il frontend statico
	router.ServeFiles("/*filepath", http.Dir("./frontend"))

	// Login
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.LoginHandler(db, w, r, ps)
	})

	// Employees data
	router.GET("/:profile/employees", WithRequestContext(api.GetEmployeesData))

	// Search project
	router.GET("/:profile/projects", WithRequestContext(api.SearchProject))

	// Get all projects
	router.GET("/:profile/departments", WithRequestContext(api.GetDepartment))

	// Modify manager
	router.PUT("/:profile/departments/:department", WithRequestContext(api.ModifyManager))

	return router
}

// Avvia il server HTTP
func startServer(router *httprouter.Router) {
	addr := "localhost:8080"
	fmt.Println("Server avviato su", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func WithRequestContext(
	handler func(*sql.DB, http.ResponseWriter, *http.Request, reqcontext.RequestContext, httprouter.Params),
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqID, _ := uuid.NewV4()

		// Estrai token e sessione
		token := r.Header.Get("Authorization")
		sessionStr := r.Header.Get("X-Session")
		session := 0
		if parsedSession, err := strconv.Atoi(sessionStr); err == nil {
			session = parsedSession
		}

		logger := logrus.WithField("req_id", reqID.String())
		if token != "" {
			logger = logger.WithField("token", token)
		}
		if sessionStr != "" {
			logger = logger.WithField("session", session)
		}

		ctx := reqcontext.RequestContext{
			ReqUUID: reqID,
			Token:   token,
			Session: session,
			Logger:  logger,
		}

		handler(db, w, r, ctx, ps)
	}
}
