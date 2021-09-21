package data

import (
	"fmt"
)

// Node element to keep element and next node together
type Node struct {
	Id uint64 `db:",primarykey,autoincrement"`
}

// String returns a string representation of the node
func (n Node) String() string {
	return fmt.Sprintf("%+v", n)
}

