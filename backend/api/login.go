package api

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	database "github.com/Mattia35/badServer/backend/database"
)

// Handler API per il login
func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// struct per la richiesta di login
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var loginRequest LoginRequest

	// Decodifica la richiesta JSON
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if loginRequest.Username == "" || loginRequest.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Confronta le credenziali con quelle nel database
	err := database.CheckCredentials(db, loginRequest.Username, loginRequest.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error: isn't possibile to check credentials", http.StatusInternalServerError)
		}
		return
	}

	// Struct risposta
	type AuthUser struct {
		Username string `json:"username"`
		Token    string `json:"token"`
		Session  int    `json:"session"`
	}

	// Crea un token di autorizzazione
	token, err := GenerateSecureToken(32)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to generate the token", http.StatusInternalServerError)
		return
	}

	// Salva nel db il token generato
	session, err := database.SaveToken(db, loginRequest.Username, token)
	if err != nil {
		http.Error(w, "Internal server error: isn't possible to save the token", http.StatusInternalServerError)
		return
	}

	var authUser AuthUser
	authUser.Username = loginRequest.Username
	authUser.Token = token
	authUser.Session = session

	// Encode the AuthUser object in JSON and send it to the client.
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(authUser); err != nil {
		http.Error(w, "Internal server error: ins't possible to encode the token", http.StatusInternalServerError)
		return
	}
}

// Genera un token sicuro di una lunghezza specificata
func GenerateSecureToken(length int) (string, error) {

	// Crea un array di byte vuoto della lunghezza desiderata
	bytes := make([]byte, length)

	// Riempi i byte con dati casuali sicuri (da crypto/rand)
	_, err := rand.Read(bytes)

	// Se c'è un errore, lo restituisce
	if err != nil {
		return "", err
	}

	// Converte i byte casuali in una stringa in formato Base64 URL-safe
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
