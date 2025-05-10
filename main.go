package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	api "github.com/Mattia35/badServer/backend/api"
	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func main() {
	initDB()
	defer db.Close()
	router := setupRoutes()
	startServer(router)
}

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consenti richieste da qualsiasi origine (oppure specifica: http://localhost:5173)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Session")

		// Rispondi immediatamente a richieste OPTIONS (preflight)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}


// Inizializza la connessione al database e crea tabelle
func initDB() {
	var err error
	dsn := "baduser:badpass@tcp(127.0.0.1:3306)/badserver?parseTime=true&multiStatements=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Errore apertura database:", err)
	}

	// Test connessione
	if err := db.Ping(); err != nil {
		log.Fatal("Errore connessione al database:", err)
	}

	createTables()
}

func createTables() {
	/*
	drops := []string{
		"SET FOREIGN_KEY_CHECKS = 0;",
		"DROP TABLE IF EXISTS token;",
		"DROP TABLE IF EXISTS profile;",
		"DROP TABLE IF EXISTS employee;",
		"DROP TABLE IF EXISTS project;",
		"DROP TABLE IF EXISTS department;",
		"SET FOREIGN_KEY_CHECKS = 1;",
	}
	for _, stmt := range drops {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatal("Errore durante DROP:", err)
		}
	}
		*/

	// Creazione tabelle
	queries := []string{
		`CREATE TABLE IF NOT EXISTS department (
			id INT NOT NULL,
			name VARCHAR(100) NOT NULL,
			manager INT,
			address VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		);`,
		`CREATE TABLE IF NOT EXISTS project (
			id INT NOT NULL,
			name VARCHAR(100) NOT NULL,
			start_date DATE NOT NULL,
			end_date DATE NOT NULL,
			budget FLOAT NOT NULL,
			department INT NOT NULL UNIQUE,
			PRIMARY KEY (id),
			CONSTRAINT fk_project_department FOREIGN KEY (department) REFERENCES department(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS employee (
			id INT NOT NULL,
			name_surname VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			phone VARCHAR(20) NOT NULL UNIQUE,
			address VARCHAR(255) NOT NULL,
			birth_date DATE NOT NULL,
			hire_date DATE NOT NULL,
			salary FLOAT NOT NULL,
			department INT NOT NULL,
			position VARCHAR(100) NOT NULL,
			project INT,
			PRIMARY KEY (id),
			CONSTRAINT fk_employee_project FOREIGN KEY (project) REFERENCES project(id),
			CONSTRAINT fk_employee_department FOREIGN KEY (department) REFERENCES department(id)
		);`,
		`CREATE TABLE IF NOT EXISTS profile (
			username VARCHAR(100) NOT NULL,
			password VARCHAR(255) NOT NULL,
			PRIMARY KEY (username)
		);`,
		`CREATE TABLE IF NOT EXISTS token (
			username VARCHAR(100) NOT NULL,
			token VARCHAR(255) NOT NULL,
			session INT NOT NULL,
			PRIMARY KEY (username, token),
			FOREIGN KEY (username) REFERENCES profile(username) ON DELETE CASCADE
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Errore creazione tabella:", err)
		}
	}

	// Aggiungi foreign key ciclica dopo la creazione di tutte le tabelle
	_, err := db.Exec(`ALTER TABLE department
		ADD CONSTRAINT fk_department_manager FOREIGN KEY (manager) REFERENCES employee(id);`)
	if err != nil {
		if !strings.Contains(err.Error(), "errno: 121") {
			log.Fatal("Errore aggiunta FK manager:", err)
		}
	}
}



func setupRoutes() *httprouter.Router {
	router := httprouter.New()

	// Serve il frontend statico
	router.ServeFiles("/static/*filepath", http.Dir("./frontend"))

	// Login
	router.PUT("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.LoginHandler(db, w, r, ps)
	})

	// Employees data
	router.GET("/profiles/:profile/employees", WithRequestContext(api.GetEmployeesData))

	// Search project
	router.GET("/profiles/:profile/projects", WithRequestContext(api.SearchProject))

	// Get all projects
	router.GET("/profiles/:profile/departments", WithRequestContext(api.GetDepartment))

	// Modify manager
	router.PUT("/profiles/:profile/departments/:department", WithRequestContext(api.ModifyDepAddress))

	return router
}

// Avvia il server HTTP
func startServer(router *httprouter.Router) {
	addr := "localhost:8080"
	fmt.Println("Server avviato su", addr)
	log.Fatal(http.ListenAndServe(addr, enableCORS(router)))
}

func WithRequestContext(
	handler func(*sql.DB, http.ResponseWriter, *http.Request, reqcontext.RequestContext, httprouter.Params),
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqID, _ := uuid.NewV4()

		// Estrai token e sessione
		token := r.Header.Get("Token")
		sessionStr := r.Header.Get("Session")
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
