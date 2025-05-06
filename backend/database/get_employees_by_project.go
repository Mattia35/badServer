package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func GetEmplByProj(db *sql.DB, progectId int) (string, error) {
	// Ottieni il nome dell'impiegato
	query := `SELECT name_surname FROM employee WHERE project = '` + strconv.Itoa(progectId) + `'`
	var employee string
	err := db.QueryRow(query).Scan(&employee)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("nessun impiegato trovato con questo progetto")
		}
		return "", err
	}
	return employee, nil
}
