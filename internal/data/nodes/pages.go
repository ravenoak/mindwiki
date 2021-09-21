package nodes

import (
	"github.com/ravenoak/mindwiki/internal/data"
)

type Page struct {
	data.Node

	Title string
	Body  string
	Slug  string
	Tags  []*Tag
}
