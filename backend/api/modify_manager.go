package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	structions "github.com/Mattia35/badServer/backend/api/structs"
	database "github.com/Mattia35/badServer/backend/database"
	"github.com/julienschmidt/httprouter"
)

func ModifyManager(db *sql.DB, w http.ResponseWriter, r *http.Request, ctx reqcontext.RequestContext, ps httprouter.Params) {
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

	// ottieni il dipartimento dall'URL
	var department structions.Department
	department.ID, err = strconv.Atoi(ps.ByName("department"))
	if err != nil {
		http.Error(w, "Can't get the department from URL", http.StatusBadRequest)
		return
	}

	// Ottieni il dipartimento dalla richiesta
	var newManager string
	err = json.NewDecoder(r.Body).Decode(&newManager)
	if err != nil {
		http.Error(w, "Bad request: isn't possible to decode department", http.StatusBadRequest)
		return
	}

	// Controlla se il dipartimento è valido
	if department.Name == "" {
		http.Error(w, "Bad request: isn't possible to get the input", http.StatusBadRequest)
		return
	}

	// Modifica il manager del dipartimento nel database
	err = database.ModifyManager(db, newManager, department)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to modify manager", http.StatusInternalServerError)
		return
	}

	// Imposta l'intestazione della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	// Scrivi la risposta di successo
	response := map[string]string{"message": "Manager modified successfully"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error: isn't possible to encode response", http.StatusInternalServerError)
		return
	}

}
