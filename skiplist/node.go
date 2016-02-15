package skiplist

// nodes is a slice of type node
type nodes []*node

type node struct {
	// forward denotes the forward pointing pointers in this node
	forward 	nodes

	entry 		Comparator
}

// node implements Comparator interface
func (n *node) Compare(cmp Comparator) int {
	return n.entry.Compare(cmp)
}

func newNode(val Comparator, maxlevel uint8) *node {
	return &node{
		entry: val,
		forward: make(nodes, maxlevel),
	}
}