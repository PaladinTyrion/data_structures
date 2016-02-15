package skiplist

type SkipList struct {
	maxLevel	uint8
	level 		uint8
	head        *node
	length      uint64

	// In the insert/delete case, cache record the pre-Node
	// before operation at every level.
	cache		nodes
}

func (sl *SkipList) init(para interface{}) {
	switch para.(type) {
		case uint8:
		sl.maxLevel = 8
		case uint16:
		sl.maxLevel = 16
		case uint32:
		sl.maxLevel = 32
		case uint64, uint:
		sl.maxLevel = 64
		default:
		sl.maxLevel = 32
	}
	sl.cache = make(nodes, sl.maxLevel)
	sl.head = newNode(nil, sl.maxLevel)
}

func NewSkipList(para interface{}) *SkipList {
	sl := &SkipList{}
	sl.init(para)
	return sl
}

// Len returns the number of items in this skiplist.
func (sl *SkipList) Len() uint64 {
	return sl.length
}

func (s *SkipList) seach(val Comparator) (cache *node) {
	return
}

func (s *SkipList) Insert(val Comparator) (overwrite Comparator) {
	return
}

func (sl *SkipList) Delete(val Comparator) (del Comparator) {
	return
}

func (s *SkipList) InsertMul(vals ...Comparators) (overwrite Comparators) {
	return
}

func (sl *SkipList) DeleteMul(vals ...Comparators) (del Comparators) {
	return
}