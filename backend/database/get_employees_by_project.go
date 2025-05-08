package database

import (
	"database/sql"
	"strconv"
)

func GetEmplByProj(db *sql.DB, progectId int) ([]string, error) {
	// Ottieni il nome dell'impiegato
	query := `SELECT name_surname FROM employee WHERE project = '` + strconv.Itoa(progectId) + `'`
	var employees []string
	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return employees, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee string
		err := rows.Scan(&employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
