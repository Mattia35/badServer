package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	structions "github.com/Mattia35/badServer/backend/api/structs"
	database "github.com/Mattia35/badServer/backend/database"
)

func SearchProject(db *sql.DB, w http.ResponseWriter, r *http.Request, ctx reqcontext.RequestContext, ps httprouter.Params) {

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
		http.Error(w, "Internal server error: isn't possible to check session: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !control {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var project structions.Project
	// ottiene il nome del progetto dalla richiesta
	project.Name = r.URL.Query().Get("name")
	if project.Name == "" {
		http.Error(w, "Bad request: isn't possible to get the input", http.StatusInternalServerError)
		return
	}

	var listProject []structions.Project
	// ottieni i dati del progetto dal database
	listProject, err = database.GetProject(db, project.Name)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to get project data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	type TotalProj struct {
		Project  []structions.Project `json:"project"`
		Employee []string             `json:"employees"`
	}
	var response TotalProj
	response.Project = listProject
	// per ogni progetto ottiene l'elenco di utenti che ne fanno parte
	for i := 0; i < len(listProject); i++ {
		// ottieni i dati degli utenti che fanno parte del progetto
		employee, err := database.GetEmplByProj(db, listProject[i].ID)
		if err != nil {
			http.Error(w, "Internal server error: isn't possible to get employee from project: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response.Employee = append(response.Employee, employee)
	}

	// Imposta l'intestazione della risposta come JSON
	w.Header().Set("Content-Type", "application/json")

	// Scrivi la lista dei progetti nella risposta
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error: isn't possible to encode list of project: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
