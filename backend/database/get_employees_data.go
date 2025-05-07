package database

import (
	"database/sql"
	"errors"

	structions "github.com/Mattia35/badServer/backend/api/structs"
)

func GetEmployeesData(db *sql.DB, name_and_surname string) ([]structions.Employee, error) {
	Employees := []structions.Employee{}
	// Ottieni i dati dell'impiegato
	query := `SELECT name_surname, email, phone, department, position, COALESCE(project, 0) FROM employee WHERE name_surname LIKE '` + name_and_surname + `%'`
	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("nessun impiegato trovato con questo nome")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee structions.Employee
		err := rows.Scan(&employee.NameSurname, &employee.Email, &employee.Phone, &employee.Department, &employee.Position, &employee.Project)
		if err != nil {
			return nil, err
		}
		Employees = append(Employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Employees, nil
}
