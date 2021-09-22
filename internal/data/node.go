package data

import (
	"fmt"
)

// Node element
type Node struct {
	Id uint64 `db:",primarykey,autoincrement"`
}

// String returns a string representation of the node
func (n Node) String() string {
	return fmt.Sprintf("%+v", n)
}

