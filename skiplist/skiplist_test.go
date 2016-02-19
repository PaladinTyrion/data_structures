package skiplist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func generateSimpleMockEntries(init, num int) Comparators {
	entries := make(Comparators, 0, num)
	for i := 0; i < num; i++ {
		entries = append(entries, newMockEntry(uint64(init+i)))
	}

	return entries
}

func generateRandomMockEntries(num int) Comparators {
	entries := make(Comparators, 0, num)
	for i := 0; i < num; i++ {
		entries = append(entries, newMockEntry(uint64(rand.Int())))
	}

	return entries
}

func TestSeekAndInsert(t *testing.T) {
	m1 := newMockEntry(15)
	m2 := newMockEntry(61)
	m3 := newMockEntry(33)
	m := make(Comparators, 0)
	m = append(m, m1, m2, m3)

	sl := NewSkipList(uint8(0))
	assert.Equal(t, uint64(0), sl.Len())
	assert.Equal(t, Comparators{}, sl.Seek(m1))

	overwritten := sl.Insert(m1)
	assert.Equal(t, Comparators{m1}, sl.Seek(m1))
	assert.Equal(t, uint64(1), sl.Len())
	assert.Equal(t, nil, overwritten)
	assert.Equal(t, Comparators{}, sl.Seek(mockEntry(1)))

	overwrittens := sl.InsertMul(m...)
	assert.Equal(t, Comparators{m2}, sl.Seek(m2))
	assert.Equal(t, Comparators{}, sl.Seek(mockEntry(77)))
	assert.Equal(t, uint64(3), sl.Len())
	assert.Equal(t, Comparators{m1}, overwrittens)

}

func TestDelete(t *testing.T) {
	m := generateSimpleMockEntries(33, 15)
	m1 := newMockEntry(35)

	m2 := newMockEntry(45)
	m3 := newMockEntry(46)
	m4 := newMockEntry(47)
	mdels := make(Comparators, 3)
	mdels = append(mdels, m2, m3, m4)

	sl := NewSkipList(uint8(0))
	overwrittens := sl.InsertMul(m...)
	assert.Equal(t, uint64(15), sl.Len())
	assert.Equal(t, Comparators(nil), overwrittens)
	del := sl.Delete(m1)
	assert.Equal(t, uint64(14), sl.Len())
	assert.Equal(t, m1, del)
	dels := sl.DeleteMul(mdels...)
	assert.Equal(t, uint64(11), sl.Len())
	assert.Equal(t, Comparators{m2, m3, m4}, dels)
}
