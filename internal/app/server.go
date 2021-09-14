package app

import (
	"context"
)

type HTTPServer interface {
	Shutdown(context.Context) error
	Start() error
}

