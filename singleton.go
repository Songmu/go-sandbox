package main

import (
	"fmt"

	"github.com/Songmu/go-sandbox/singleton"
)

func main() {
	deeeet := singleton.GetDeeeter()

	deeeet.Deeeet()
	deeeet.Deeeet()
	singleton.GetDeeeter().Deeeet()
	deeeet.Deeeet()
}

type imitateDeeeet struct {
}

func (ide *imitateDeeeet) Deeeet() {
	fmt.Println("deeeet(偽)です…")
}

func (ide *imitateDeeeet) getAge() {
	fmt.Println("17歳です♥")
}

// var _ singleton.Deeeeter = &imitateDeeeet{}
