package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"

	reqcontext "github.com/Mattia35/badServer/backend/api/requestContext"
	structions "github.com/Mattia35/badServer/backend/api/structs"
	database "github.com/Mattia35/badServer/backend/database"
)

func searchProject(db *sql.DB, w http.ResponseWriter, r *http.Request, ctx reqcontext.RequestContext) {

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

	var project structions.Project
	// ottiene il nome del progetto dalla richiesta
	project.Name = r.URL.Query().Get("Name")
	validQuerySearch := regexp.MustCompile(`^[a-z0-9]{1,13}$`)
	if !validQuerySearch.MatchString(project.Name) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var listProject []structions.Project
	// ottieni i dati del progetto dal database
	listProject, err = database.GetProject(db, project.Name)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to get project data", http.StatusInternalServerError)
		return
	}

	type TotalProj struct {
		Project []structions.Project `json:"project"`
		Users string `json:"users"`
	}
	var response TotalProj
	response.Project = listProject
	// per ogni progetto ottiene l'elenco di utenti che ne fanno parte
	for i := 0; i < len(listProject); i++ {
		// ottieni i dati degli utenti che fanno parte del progetto
		response.Users, err = database.GetEmplByProj(db, listProject[i].ID)
		if err != nil {
			http.Error(w, "Internal server error: isn't possible to get users from project", http.StatusInternalServerError)
			return
		}
	}

	// Imposta l'intestazione della risposta come JSON
	w.Header().Set("Content-Type", "application/json")

	// Scrivi i dati degli impiegati nella risposta
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error: isn't possible to encode list of project", http.StatusInternalServerError)
		return
	}

}
