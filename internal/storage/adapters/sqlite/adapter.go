package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/rs/zerolog/log"
)

type SQLiteAdapter struct {
	config *config.StorageConfig
	db     *sql.DB
}

func (a SQLiteAdapter) Delete(set, id []byte) (interface{}, error) {
	return nil, nil
}

func (a SQLiteAdapter) Insert(set, id []byte, object interface{}) error {
	return nil
}

func (a SQLiteAdapter) Query(parameters []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (a *SQLiteAdapter) Open() error {
	var err error
	dsn := "file://" + a.config.Path
	a.db, err = sql.Open("sqlite3", dsn)
	return err
}

func (a SQLiteAdapter) Close() {
	if err := a.db.Close(); err != nil {
		log.Fatal().Err(err)
	}
}
