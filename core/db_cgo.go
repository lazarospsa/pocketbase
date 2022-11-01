//go:build cgo

package core

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pocketbase/dbx"
)

func connectDB(dbPath string) (*dbx.DB, error) {
	// note: the busy_timeout pragma must be first because
	// the connection needs to be set to block on busy before WAL mode
	// is set in case it hasn't been already set by another connection
	pragmas := "_busy_timeout=10000&_journal_mode=WAL&_foreign_keys=1&_synchronous=NORMAL"

	db, openErr := dbx.MustOpen("sqlite3", fmt.Sprintf("%s?%s", dbPath, pragmas))
	if openErr != nil {
		return nil, openErr
	}

	// additional pragmas not supported through the dsn string
	_, err := db.NewQuery(`
		pragma journal_size_limit = 100000000;
	`).Execute()

	return db, err
}
