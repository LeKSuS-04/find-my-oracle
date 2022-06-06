package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ipMaskRow struct {
	mask    string
	checked bool
}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "cache.sqlite")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(ipMasks []string) (*sql.DB, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ip_masks (
			mask STRING NOT NULL PRIMARY KEY,
			checked BOOLEAN NOT NULL DEFAULT FALSE
		);
	`)
	if err != nil {
		return nil, err
	}

	for _, mask := range ipMasks {
		db.Exec(`
			INSERT INTO ip_masks (mask) values ($1);
		`, mask)
	}

	return db, nil
}

func SetIPMaskChecked(db *sql.DB, ipMask string) error {
	_, err := db.Exec(`
		UPDATE ip_masks
		SET checked = TRUE
		WHERE mask = $1;
	`, ipMask)
	return err
}

func IsIPMaskChecked(db *sql.DB, ipMask string) bool {
	row := db.QueryRow(`
		SELECT checked
		FROM ip_masks
		WHERE mask = $1;
	`, ipMask)
	var checked bool
	row.Scan(&checked)
	return checked
}
