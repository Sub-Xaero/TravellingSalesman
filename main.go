package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"subxaero/ga"
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

	distances := []Distance{
		{Bristol, Bath, 14},
		{Bristol, London, 105},
		{Bristol, Leeds, 275},
		{London, Bath, 110},
		{London, Leeds, 227},
		{Bath, Leeds, 291},
		{Glasgow, London, 724},
		{Glasgow, Bath, 850},
		{Glasgow, Bristol, 802},
		{Glasgow, Leeds, 620},

		{Swansea, London, 250},
		{Swansea, Bath, 105},
		{Swansea, Bristol, 98},
		{Swansea, Leeds, 324},
		{Swansea, Glasgow, 821},
	}
	maxDistance := 0
	for _, i := range distances {
		maxDistance += i.Distance
	}
	fmt.Println(distances)

}
