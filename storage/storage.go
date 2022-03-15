package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

const (
	timeout      = 1 * time.Second
	errPathEmpty = "you suck, path is empty"
)

type Storage struct {
	dbPath string
	store  *bbolt.DB
}

type Storable interface {
	Save(tx *bbolt.Tx) error
}

func NewStorage(dbPath string) (*Storage, error) {
	if dbPath == "" {
		return nil, errors.New(errPathEmpty)
	}
	s := &Storage{
		dbPath: dbPath,
	}
	return s, nil
}

func (s *Storage) Open() error {
	var err error
	s.store, err = bbolt.Open(s.dbPath, 0600, &bbolt.Options{Timeout: timeout})
	return err
}

func (s *Storage) Close() error {
	return s.store.Close()
}

func (s *Storage) NewTx(writeable bool) (*bbolt.Tx, error) {
	tx, err := s.store.Begin(writeable)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *Storage) DBStats() string {
	t := s.store.Stats()
	log.Debug().Interface("s.store.Stats()", t).Msg("")
	return fmt.Sprintf("%+v", t)
}
