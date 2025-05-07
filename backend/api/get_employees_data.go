package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	database "github.com/Mattia35/badServer/backend/database"
)

// Handler API per l'ottenimento dei dati degli impiegati
func GetEmployeesData(db *sql.DB, w http.ResponseWriter, r *http.Request, ctx reqcontext.RequestContext, ps httprouter.Params) {
	// Ottieni la sessione dell'utente
	session := ctx.Session

	// Controlla se l'utente è autenticato
	if ctx.Token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ottieni il token dell'utente
	token := ctx.Token

	// Controlla che la sessione sia valida
	control, err := database.CheckSession(db, session, token)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to check session: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if !control {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the search query
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Bad request: isn't possible to get the input", http.StatusInternalServerError)
		return
	}

	// Ottieni i dati degli impiegati dal database
	employees, err := database.GetEmployeesData(db, query)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to get employees data: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Imposta l'intestazione della risposta come JSON
	w.Header().Set("Content-Type", "application/json")

	// Scrivi i dati degli impiegati nella risposta
	if err := json.NewEncoder(w).Encode(employees); err != nil {
		http.Error(w, "Internal server error: isn't possible to encode employees data: " + err.Error(), http.StatusInternalServerError)
		return
	}
}
