package database

import (
	"database/sql"

	structions "github.com/Mattia35/badServer/backend/api/structs"
)

func CheckCredentials(db *sql.DB, username string, password string) (string, error) {
	// Controlla se le credenziali sono valide
	query := `SELECT username, password FROM profile WHERE username = '` + username + `' AND password = '` + password + `'`
	var profile structions.Profile
	err := db.QueryRow(query).Scan(&profile.Username, &profile.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "",err
		}
		return "",err
	}
	return profile.Username, nil
}
