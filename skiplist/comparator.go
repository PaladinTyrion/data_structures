package skiplist

// Comparator is an interface that represents items that can
// be compared.
type Comparator interface {
	// Compare compares this interface with another.
	// Returns a positive number if this interface is greater,
	//      0 if equal,
	//      a negative number if less.
	Compare(Comparator) int
}

// Comparators is a list of type Comparator.
type Comparators []Comparator


