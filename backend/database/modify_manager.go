package database

import (
	"database/sql"
	"strconv"

	structions "github.com/Mattia35/badServer/backend/api/structs"
)

func ModifyManager(db *sql.DB, newManager string, department structions.Department) error {
	// Modifica il manager del dipartimento nel database
	_, err := db.Exec(`UPDATE department SET manager = ` + newManager + `WHERE id = ` + strconv.Itoa(department.ID))
	if err != nil {
		return err
	}
	return nil
}
