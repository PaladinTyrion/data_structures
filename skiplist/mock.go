package skiplist

type mockEntry uint64

func (me mockEntry) Compare(other Comparator) int {
	otherU := other.(mockEntry)
	if me == otherU {
		return 0
	}

	if me > otherU {
		return 1
	}

	return -1
}

func newMockEntry(num uint64) mockEntry {
	return mockEntry(num)
}

