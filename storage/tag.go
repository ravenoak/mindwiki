package storage

import (
	"go.etcd.io/bbolt"
)

const tagBucketName = "tag"

type Tag struct {
	Name        string
	Description string
	Slug        string
}

func (t *Tag) Save(tx *bbolt.Tx) error {
	b := tx.Bucket([]byte(tagBucketName))
	if err := b.Put([]byte(t.Slug+"-Name"), []byte(t.Name)); err != nil {
		return err
	}
	return nil
}
