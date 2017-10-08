package main

import (
	"math/rand"
	"time"
)


type Distance struct {
	City1, City2 int
	Distance     int
}

const (
	_       = iota
	London  = iota
	Bath    = iota
	Bristol = iota
	Leeds   = iota
	Glasgow = iota
	Swansea = iota
)

func main() {
	rand.Seed(time.Now().Unix())
}
