package database

import (
	"database/sql"
	structions "github.com/Mattia35/badServer/backend/api/structs"
)

func GetDepartment(db *sql.DB) ([]structions.Department, error) {
	var departments []structions.Department
	rows, err := db.Query(`SELECT id, name, COALESCE(manager, 0) FROM department`)
	if err != nil {
		return nil, err
	}
	defer func() { rows.Close() }()
	
	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		var department structions.Department
		err = rows.Scan(&department.ID, &department.Name, &department.Manager)
		if err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	return departments, nil
}
