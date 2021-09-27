package storage

import (
	"time"

	"go.etcd.io/bbolt"
)

const (
	timeout = 1 * time.Second
)

type Storage struct {
	dbPath string
	store  *bbolt.DB
}

func (s Storage) Open() error {
	var err error
	s.store, err = bbolt.Open(s.dbPath, 0600, &bbolt.Options{Timeout: timeout})
	return err
}
