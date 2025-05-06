package database

import(
	"database/sql"
	structions "github.com/Mattia35/badServer/backend/api/structs"
)

func GetProject(db *sql.DB, name string) ([]structions.Project, error) {
	// Ottieni i dati del progetto dal database
	var listProject []structions.Project
	rows, err := db.Query(`SELECT id, name, start_date end_date, department FROM project WHERE name LIKE '`+name+`%'`)
	if err != nil {
		return nil, err
	}
	defer func() { rows.Close() }()

	// per ogni progetto
	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		var project structions.Project
		err = rows.Scan(&project.ID, &project.Name, &project.StartDate, &project.EndDate, &project.DepartmentID)
		if err != nil {
			return nil, err
		}
		listProject = append(listProject, project)
	}

	return listProject, nil
}
