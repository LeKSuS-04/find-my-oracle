package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

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
			check_result BOOLEAN
		);
	`)
	if err != nil {
		return nil, err
	}

	for _, mask := range ipMasks {
		db.Exec(`
			INSERT INTO ip_masks (mask) values ($1)
		`, mask)
	}

	return db, nil
}

func SetIPMask(db *sql.DB, ipMask string, checkResult bool) error {
	_, err := db.Exec(`
		UPDATE ip_masks
		SET check_result = $1
		WHERE mask = $2	
	`, checkResult, ipMask)
	return err
}
