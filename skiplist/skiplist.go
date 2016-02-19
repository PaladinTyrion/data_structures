package skiplist

type SkipList struct {
	maxLevel	uint8
	level 		uint8
	head        *node
	length      uint64
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

func (sl *SkipList) search(val Comparator, cache nodes) (resNode *node) {
	if sl.length == 0 {
		return nil
	}

	if val == nil {
		return nil
	}

	n := sl.head
	var alreadyChecked *node
	for i := int(sl.level)-1; i >= 0 ; i-- {
		for ;n.forward[i] != nil && n.forward[i] != alreadyChecked &&
			n.forward[i].entry != nil && n.forward[i].Compare(val) < 0; {
			n = n.forward[i]
		}

		alreadyChecked = n
		if cache != nil {
			cache[i] = n
		}
	}
	return n.forward[0]
}

func (sl *SkipList) Seek(vals ...Comparator) (results Comparators) {
	results = make(Comparators, 0, len(vals))

	for _, val := range vals {
		n := sl.search(val, nil)
		if n != nil && n.Compare(val) == 0 {
			results = append(results, n.entry)
		}
	}
	return results
}

func (sl *SkipList) insert(n *node, val Comparator, cache nodes, allowOverwrite bool) (overWrite Comparator) {
	if !allowOverwrite && n != nil && n.entry.Compare(val) == 0 {
		return val
	}

	if val == nil {
		return nil
	}

	nLevel := generatorLevel(sl.maxLevel)
	if nLevel > sl.level {
		for i := sl.level; i < nLevel; i++ {
			cache[i] = sl.head
		}
		sl.level = nLevel
	}

	nn := newNode(val, sl.maxLevel)
	for i := 0; i < int(nLevel); i++ {
		nn.forward[i] = cache[i].forward[i]
		cache[i].forward[i] = nn
	}

	sl.length++

	return nil
}


func (sl *SkipList) Insert(val Comparator) (overwrite Comparator) {
	//cache saves pre-Nodes of searching val in the skiplist at every level
	cache := make(nodes, sl.maxLevel)
	n := sl.search(val, cache)
	return sl.insert(n, val, cache, false)
}

func (sl *SkipList) delete(n *node, val Comparator, cache nodes) (del Comparator) {
	if n == nil || n.Compare(val) != 0 {
		return nil
	}

	sl.length--
	for i := 0; i < int(sl.level); i++ {
		cache[i].forward[i] = n.forward[i]
	}

	for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
		sl.level--
	}

	return n.entry
}

func (sl *SkipList) Delete(val Comparator) (del Comparator) {
	cache := make(nodes, sl.maxLevel)
	n := sl.search(val, cache)
	return sl.delete(n, val, cache)
}

func (sl *SkipList) InsertMul(vals ...Comparator) (overwrites Comparators) {
	for _, val := range vals {
		cache := make(nodes, sl.maxLevel)
		n := sl.search(val, cache)
		overwrite := sl.insert(n, val, cache, false)
		if overwrite != nil {
			overwrites = append(overwrites, overwrite)
		}
	}
	return overwrites
}

func (sl *SkipList) DeleteMul(vals ...Comparator) (dels Comparators) {
	for _, val := range vals {
		cache := make(nodes, sl.maxLevel)
		n := sl.search(val, cache)
		del := sl.delete(n, val, cache)
		if del != nil {
			dels = append(dels, del)
		}
	}
	return dels
}