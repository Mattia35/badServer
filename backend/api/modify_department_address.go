package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	database "github.com/Mattia35/badServer/backend/database"
	"github.com/julienschmidt/httprouter"
)

func ModifyDepAddress(db *sql.DB, w http.ResponseWriter, r *http.Request, ctx reqcontext.RequestContext, ps httprouter.Params) {
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

	// ottieni il dipartimento dall'URL
	nameDepartment := ps.ByName("department")
	if nameDepartment == "" {
		http.Error(w, "Can't get the department from URL", http.StatusBadRequest)
		return
	}

	// Struct per la richiesta di modifica della sede
	type ModifyAddrRequest struct {
		NewAddr string `json:"address"`
	}

	// Ottieni il dipartimento dalla richiesta
	var req ModifyAddrRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request: isn't possible to decode department: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Controlla se il dipartimento è valido
	if nameDepartment == "" {
		http.Error(w, "Bad request: isn't possible to get the input", http.StatusBadRequest)
		return
	}

	// Modifica la sede del dipartimento nel database
	err = database.ModifyDepAddress(db, req.NewAddr, nameDepartment)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to modify address: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Imposta l'intestazione della risposta come JSON
	w.Header().Set("Content-Type", "application/json")
	// Scrivi la risposta di successo
	response := map[string]string{"message": "Address modified successfully"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error: isn't possible to encode response: " + err.Error(), http.StatusInternalServerError)
		return
	}

}
