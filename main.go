package main

import (
	"fmt"
	"math/rand"
	. "subxaero/ga"
	"time"
)

type Distance struct {
	City1, City2 int
	Distance     int
}

const (
	London        = iota
	Nottingham    = iota
	Birmingham    = iota
	Glasgow       = iota
	Edinborough   = iota
	Cambridge     = iota
	Oxford        = iota
	Bristol       = iota
	Bath          = iota
	Manchester    = iota
	Cardiff       = iota
	Belfast       = iota
	Coventry      = iota
	Plymouth      = iota
	Wolverhampton = iota
)

func main() {
	rand.Seed(time.Now().Unix())

	distances := [][]int{
		{0, 29, 82, 46, 68, 52, 72, 42, 51, 55, 29, 74, 23, 72, 46},
		{29, 0, 55, 46, 42, 43, 43, 23, 23, 31, 41, 51, 11, 52, 21},
		{82, 55, 0, 68, 46, 55, 23, 43, 41, 29, 79, 21, 64, 31, 51},
		{46, 46, 68, 0, 82, 15, 72, 31, 62, 42, 21, 51, 51, 43, 64},
		{68, 42, 46, 82, 0, 74, 23, 52, 21, 46, 82, 58, 46, 65, 23},
		{52, 43, 55, 15, 74, 0, 61, 23, 55, 31, 33, 37, 51, 29, 59},
		{72, 43, 23, 72, 23, 61, 0, 42, 23, 31, 77, 37, 51, 46, 33},
		{42, 23, 43, 31, 52, 23, 42, 0, 33, 15, 37, 33, 33, 31, 37},
		{51, 23, 41, 62, 21, 55, 23, 33, 0, 29, 62, 46, 29, 51, 11},
		{55, 31, 29, 42, 46, 31, 31, 15, 29, 0, 51, 21, 41, 23, 37},
		{29, 41, 79, 21, 82, 33, 77, 37, 62, 51, 0, 65, 42, 59, 61},
		{74, 51, 21, 51, 58, 37, 37, 33, 46, 21, 65, 0, 61, 11, 55},
		{23, 11, 64, 51, 46, 51, 51, 33, 29, 41, 42, 61, 0, 62, 23},
		{72, 52, 31, 43, 65, 29, 46, 31, 51, 23, 59, 11, 62, 0, 59},
		{46, 21, 51, 64, 23, 59, 33, 37, 11, 37, 61, 55, 23, 59, 0},
	}

	maxDistance := 0
	for i := range distances {
		maxDistance += distances[i][0]
	}
	fmt.Println(maxDistance)

	var ga = NewGeneticAlgorithm()
	ga.SetGenerateCandidate(func(length int) ([]int, error) {
		possibilities := []int{
			London,
			Nottingham,
			Birmingham,
			Glasgow,
			Edinborough,
			Cambridge,
			Oxford,
			Bristol,
			Bath,
			Manchester,
			Cardiff,
			Belfast,
			Coventry,
			Plymouth,
			Wolverhampton,
		}

		sequence := make([]int, 0)
		for i := 0; i < length; i++ {
			choice := rand.Int() % len(possibilities)
			sequence = append(sequence, possibilities[choice])
			possibilities = append(possibilities[:choice], possibilities[choice+1:]...)
		}
		sequence = append(sequence, sequence[0])
		return sequence, nil
	})

	ga.SetMutateFunc(func(gene Genome) Genome {
		sequence := gene.Sequence
		newSequence := append(make([]int, 0), gene.Sequence...)
		min := 1
		max := len(sequence) - 1

		choice1 := rand.Int()%(max-min) + min
		choice2 := rand.Int()%(max-min) + min

		newSequence[choice1], newSequence[choice2] = sequence[choice2], sequence[choice1]
		return Genome{Sequence: newSequence}
	})

	sequence, err := ga.GenerateCandidate(len(distances))
	if err != nil {
		panic(err)
	}
	gene := Genome{Sequence: sequence}
	fmt.Println(gene.Sequence)
	fmt.Println(ga.Mutate(gene).Sequence)

	ga.SetFitnessFunc(func(gene Genome) int {
		genomeSequence := gene.Sequence
		totalDistance := 0
		for i := 0; i < len(genomeSequence)-1; i++ {
			city1, city2 := genomeSequence[i], genomeSequence[i+1]
			totalDistance += distances[city1][city2]
		}
		return maxDistance - totalDistance
	})

	var (
		numStrings  = 10
		generations = 50000
		strLength   = len(distances) - 1
	)

	ga.Run(numStrings, strLength, generations, false, true, true)
	fmt.Println(maxDistance)
}
