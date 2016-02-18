package skiplist
import (
	"math/rand"
	"time"
	"sync"
)

const (
	p = 0.5
)

// generator is seeded with unix nanosecond and only executed once
// ensuring all random numbers come from the same randomly seeded generator.
var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

var glLock sync.Mutex

// generatorLevel generator a threadsafe level with a range from 0 to MaxLevel.
func generatorLevel(maxlevel uint8) (level uint8) {
	glLock.Lock()
	defer glLock.Unlock()

	for level = uint8(1); level <= maxlevel; level++ {
		if generator.Float64() >= p {
			return level
		}
	}
	return level
}
