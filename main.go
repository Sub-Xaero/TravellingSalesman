package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"subxaero/ga"
	"time"
)

var (
	globalChance = 6
	mutateChance = globalChance
	numStrings   = globalChance
	strLength    = globalChance
	generations  = 1000
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

	ga.SetGenerateBitString(func(length int) string {
		possibilities := []int{
			London,
			Bath,
			Bristol,
			Leeds,
			Glasgow,
			Swansea,
		}

		bitstring := ""
		for i := 0; i < length; i++ {
			choice := rand.Int() % len(possibilities)
			bitstring += strconv.Itoa(possibilities[choice])
			possibilities = append(possibilities[:choice], possibilities[choice+1:]...)
		}
		bitstring += string(bitstring[0])
		return bitstring
	})

	ga.SetCrossoverFunc(func(gene, spouse ga.Genome) []ga.Genome {

		return []ga.Genome{gene, spouse}
	})

	ga.SetMutateFunc(func(gene ga.Genome, chance int) ga.Genome {
		sequence := []rune(gene.Sequence)
		choice1 := (rand.Int() % (len(sequence) - 2)) + 1 // Start and end city must not change, so swap between
		choice2 := (rand.Int() % (len(sequence) - 2)) + 1 //

		sequence[choice1], sequence[choice2] = sequence[choice2], sequence[choice1]
		//fmt.Println(string(sequence))
		return ga.Genome{string(sequence)}
	})

	ga.SetFitnessFunc(func(gene ga.Genome) int {
		genomeSequence := gene.Sequence
		totalDistance := 0
		for index := range genomeSequence {
			if index == len(genomeSequence)-1 {
				break
			}
			city1, _ := strconv.Atoi(string(genomeSequence[index]))
			city2, _ := strconv.Atoi(string(genomeSequence[index+1]))

			for _, j := range distances {
				if (j.City1 == city1 && j.City2 == city2) || (j.City2 == city1 && j.City1 == city2) {
					//fmt.Println("Distance between", city1, "and", city2, "is", j.Distance)
					totalDistance += j.Distance
					break
				}
			}
		}
		return maxDistance - totalDistance
	})

	ga.GeneticAlgorithm(numStrings, strLength, generations, mutateChance)
}
