package app

import (
	"context"
)

type (
	HTTPServinator interface {
		Shutdown(context.Context) error
		Start() error
	}

	Storinator interface {
		Get(objectId string, objectType string) (interface{}, error)
		Put(objectId string, objectType string, objectData interface{}) error
		Find(queryParameters []interface{}) ([]interface{}, error)

		Open() error
		Close() error
	}
)
