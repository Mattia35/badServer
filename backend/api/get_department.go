package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	database "github.com/Mattia35/badServer/backend/database"
)

func GetDepartment(db *sql.DB, w http.ResponseWriter, r *http.Request, ctx reqcontext.RequestContext, ps httprouter.Params) {
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
		http.Error(w, "Internal server error: isn't possible to check session", http.StatusInternalServerError)
		return
	}
	if !control {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ottieni il dipartimento dal database
	departments, err := database.GetDepartment(db)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to get department", http.StatusInternalServerError)
		return
	}

	// Imposta l'intestazione della risposta come JSON
	w.Header().Set("Content-Type", "application/json")

	// Scrivi la lista dei dipartimanti nella risposta
	if err := json.NewEncoder(w).Encode(departments); err != nil {
		http.Error(w, "Internal server error: isn't possible to encode list of departments", http.StatusInternalServerError)
		return
	}

}
