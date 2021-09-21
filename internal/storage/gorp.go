package storage

import (
	"database/sql"

	"github.com/go-gorp/gorp/v3"
	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/ravenoak/mindwiki/internal/data/nodes"
)

type ORP struct {
	config *config.StorageConfig
	db     *sql.DB
	dbMap  *gorp.DbMap
	models map[string]interface{}
	tables map[string]*gorp.TableMap
}

func (o *ORP) Get(id string, t string) (interface{}, error) {
	return nil, nil
}

func (o *ORP) Put(id string, t string, d interface{}) error {
	return nil
}

func (o *ORP) Find(q []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (o *ORP) Open() error {
	var err error
	o.models["pages"] = nodes.Page{}

	o.db, err = sql.Open("sqlite3", o.config.Path)
	if err != nil {
		return err
	}

	// construct a gorp DbMap using SQLite
	o.dbMap = &gorp.DbMap{Db: o.db, Dialect: gorp.SqliteDialect{}}

	for k, v := range o.models {
		o.tables[k] = o.dbMap.AddTableWithName(v, k)
	}

	return nil
}

func (o *ORP) Close() error {
	return o.db.Close()
}

func NewORP(c *config.StorageConfig) *ORP {
	return &ORP{
		config: c,
	}
}
