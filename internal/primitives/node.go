package primitives

import (
	"fmt"
)

// Node element to keep element and next node together
type Node struct {
	ID uint64
}

// String returns a string representation of the node
func (n Node) String() string {
	return fmt.Sprintf("%+v", n)
}

