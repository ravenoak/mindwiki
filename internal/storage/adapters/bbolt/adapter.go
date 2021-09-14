package bbolt

import (
	"time"

	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

type BBoltAdapter struct {
	config *config.StorageConfig
	db     *bbolt.DB
}

func (a BBoltAdapter) Delete(set, id []byte) (interface{}, error) {
	return nil, nil
}

func (a BBoltAdapter) Insert(set, id []byte, object interface{}) error {
	return nil
}

func (a BBoltAdapter) Query(parameters []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (k *BBoltAdapter) Open() error {
	var err error
	k.db, err = bbolt.Open(k.config.Path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	return err
}

func (a BBoltAdapter) Close() {
	if err := a.db.Close(); err != nil {
		log.Fatal().Err(err)
	}
}

func (k *BBoltAdapter) update(fn func(*bbolt.Tx) error) error {
	return k.db.Update(fn)
}

func (k *BBoltAdapter) view(fn func(*bbolt.Tx) error) error {
	return k.db.View(fn)
}

func (k *BBoltAdapter) bucketGet(bucket, key []byte) func(tx *bbolt.Tx) error {
	return func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return InvalidBucketName{bucketName: string(bucket)}
		}
		b.Get(key)
		return nil
	}
}

func (k *BBoltAdapter) set(key []byte, value []byte) func(tx *bbolt.Tx) error {
	return func(tx *bbolt.Tx) error {
		return nil
	}
}

func NewBBoltAdapter(config *config.StorageConfig) *BBoltAdapter {
	if config == nil {
		log.Fatal().Msg("config cannot be nil")
	}
	s := BBoltAdapter{
		config: config,
	}
	return &s
}
