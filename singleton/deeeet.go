package singleton

import (
	"fmt"
	"strings"
	"sync"
)

// Deeeeter implements Deeeet() method
type Deeeeter interface {
	Deeeet()
	getAge()
}

type deeeet struct {
	age int
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
	de.mu.RLock()
	de.age++
	de.mu.RUnlock()
	fmt.Printf("d%stです…\n", strings.Repeat("e", de.age))
}

func (de *deeeet) getAge() {
	fmt.Println(de.age)
}
