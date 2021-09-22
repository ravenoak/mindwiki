package data

import (
	"fmt"
)

type Predicate interface {
}

// Edge element
type Edge struct {
	Id        uint64 `db:",primarykey,autoincrement"`
	Subject   *Node
	Predicate Predicate
	Object    *Node
}

// String returns a string representation of the edge
func (e Edge) String() string {
	return fmt.Sprintf("%+v", e)
}
