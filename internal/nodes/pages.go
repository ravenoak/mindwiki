package nodes

import (
	"github.com/ravenoak/mindwiki/internal/primitives"
)

type Page struct {
	primitives.Node

	Title string
	Body  string
	Slug  string
	Tags  []*Tag
}
