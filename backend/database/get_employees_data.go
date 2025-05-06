package database
import (
	"database/sql"
	"errors"
)

func GetEmployeesData(db *sql.DB) (string, error) {
	// Ottieni i dati dell'impiegato
	query := `SELECT id, name, surname, email, phone, department FROM employees`
	var employeeData string
	err := db.QueryRow(query).Scan(&employeeData)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Nessun impiegato trovato con questo ID")
		}
		return "", err
	}
	return employeeData, nil
}