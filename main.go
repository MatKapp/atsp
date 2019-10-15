package main

import (
	"fmt"
	"time"
)

type fn func([][]int) []int

func main() {
	instanceFilenames := []string{
		"br17",
		"ft53",
		"ft70",
		"ftv33",
		"ftv35",
		"ftv38",
		"ftv44",
		"ftv47",
		"ftv55",
		"ftv70",
	}
	for _, filename := range instanceFilenames {
		fmt.Println(filename)
		path := "data/" + filename + ".atsp"
		distances := readData(path)
		compute(solveHeuristic, distances, "Heuristic")
		compute(solveGreedy, distances, "Greedy")
		steepestElapsed := compute(solveSteepest, distances, "Steepest")
		computeRandom(distances, steepestElapsed)
	}
}

func compute(solve func([][]int) []int, distances [][]int, name string) time.Duration {
	fmt.Println(name)

	start := time.Now()
	for i := 0; i < 10; i++ {
		permutation := solve(distances)
		result := getDistance(permutation, distances)
		fmt.Println(result)
	}

	return time.Since(start)
}

func computeRandom(distances [][]int, time time.Duration){
	fmt.Println("Random")
	permutation := solveRandom(distances, time)
	result := getDistance(permutation, distances)
	fmt.Println(result)
}
