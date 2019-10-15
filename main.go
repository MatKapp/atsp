package main

import (
	"fmt"
	"time"
)

type fn func([][]int) ([]int, int)

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

func compute(solve func([][]int) ([]int, int), distances [][]int, name string) time.Duration {
	results := makeArray(10)
	stepCounts := makeArray(10)

	start := time.Now()
	for i := 0; i < 10; i++ {
		permutation, stepCount := solve(distances)
		stepCounts[i] = stepCount
		results[i] = getDistance(permutation, distances)
	}
	elapsed := time.Since(start)
	bestResult := max(results)
	meanResult := mean(results)
	meanSteps := mean(stepCounts)

	fmt.Println(name, elapsed, bestResult, meanResult, meanSteps)
	return elapsed
}

func computeRandom(distances [][]int, availableTime time.Duration) time.Duration{
	start := time.Now()
	permutation := solveRandom(distances, availableTime	)
	result := getDistance(permutation, distances)
	elapsed := time.Since(start)

	fmt.Println("Random", elapsed, result)
	return elapsed
}
