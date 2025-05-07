package database

import (
	"database/sql"
	"strconv"
)

func ModifyManager(db *sql.DB, newManager int, nameDepartment string) error {
	// Modifica il manager del dipartimento nel database
	_, err := db.Exec(`UPDATE department SET manager = '` + strconv.Itoa(newManager) + `' WHERE name = '` + nameDepartment + `'`)
	if err != nil {
		return err
	}
	return nil
}
