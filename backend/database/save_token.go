package database

import (
	"database/sql"
	"errors"
	"strconv"
)

func SaveToken(db *sql.DB, username string, token string) (int, error) {

	// Prova a ottenere la sessione massima
	query := `SELECT MAX(session) FROM token WHERE username = '` + username + `'`
	var _maxSession = sql.NullInt64{Int64: 0, Valid: false}
	row, err := db.Query(query)
	if err != nil {
		return 0, err
	}

	var maxSession int
	for row.Next() {
		if row.Err() != nil {
			return 0, err
		}

		err = row.Scan(&_maxSession)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}

		if !_maxSession.Valid {
			maxSession = 1
		} else {
			maxSession = int(_maxSession.Int64)
		}
	}

	// Inserisce il token nel database, associandolo all'username e alla sessione
	query2 := `INSERT INTO token (username, token, session) VALUES ('` + username + `', '` + token + `', ` + strconv.Itoa(maxSession) + `)`
	_, err = db.Exec(query2)
	return maxSession, err
}
