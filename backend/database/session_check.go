package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func CheckSession(db *sql.DB, session int, token string) (bool, error) {
	var session_check int
	var query = `SELECT session FROM token WHERE session = '` + strconv.Itoa(session) + `' AND token = '` + token + `'`

	err := db.QueryRow(query).Scan(&session_check)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, err
}
