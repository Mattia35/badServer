package database

import (
	"database/sql"
)

func ModifyDepAddress(db *sql.DB, newAddr string, nameDepartment string) error {
	// Modifica la sede del dipartimento nel database
	_, err := db.Exec(`UPDATE department SET address = '` + newAddr + `' WHERE name = '` + nameDepartment + `'`)
	if err != nil {
		return err
	}
	return nil
}
