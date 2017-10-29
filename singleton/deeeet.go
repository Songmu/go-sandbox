package singleton

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

// Deeeeter implements Deeeet() method
type Deeeeter interface {
	Deeeet()
	getAge()
}

type deeeet struct {
	age int64
	mu  sync.RWMutex
}

var (
	d Deeeeter
	o sync.Once
)

// GetDeeeter gets the Deeeeter
func GetDeeeter() Deeeeter {
	o.Do(func() {
		d = &deeeet{}
	})
	return d
}

// Deeeet desu...
func (de *deeeet) Deeeet() {
	age := int(atomic.AddInt64(&de.age, 1))
	fmt.Printf("d%stです…\n", strings.Repeat("e", age))
}

func (de *deeeet) getAge() {
	fmt.Println(de.age)
}
