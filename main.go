package main

import (
	"fmt"
	"time"
)

type fn func([][]int) ([]int, int)

const MIN_TIME = 1000 * 1000 * 1000

func main() {
	bestKnownSolutions := map[string]int{
		"br17": 39,
		"ft53": 6905,
		"ftv33": 1286,
		"ft70": 38673,
		"ftv35": 1473,
		"ftv44": 1613,
		"ftv38": 1530,
		"ftv55": 1608,
		"ftv47": 1776,
		"ftv64": 1839,
		"ftv70": 1950,
		"ftv100": 1788,
		"ftv110": 1958,
		"ftv90": 1579,
		"ftv120": 2166,
		"ftv130": 2307,
		"ftv140": 2420,
		"ftv150": 2611,
		"ftv170": 2755,
		"ftv160": 2683,
		"kro124": 36230,
		"p43": 5620,
		"rbg358": 1163,
		"rbg323": 1326,
		"rbg403": 2465,
		"rbg443": 2720,
		"ry48p": 14422,
	}
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
		fmt.Println()
		fmt.Println(filename)
		path := "data/" + filename + ".atsp"
		distances := readData(path)

		bestKnown := bestKnownSolutions[filename]
		fmt.Println("Best known: ", bestKnown)

		computeHeuristic(distances)
		compute(solveGreedy, distances, "Greedy")
		steepestElapsed := compute(solveSteepest, distances, "Steepest")
		computeRandom(distances, steepestElapsed)
	}
}

func compute(solve func([][]int) ([]int, int), distances [][]int, name string) time.Duration {
	results := makeArray(10)
	stepCounts := makeArray(10)

	start := time.Now()
	for i := 0; i < 10 || time.Since(start).Nanoseconds() < MIN_TIME; i++ {
		permutation, stepCount := solve(distances)
		if i < 10 {
			stepCounts[i] = stepCount
			results[i] = getDistance(permutation, distances)
		}
	}
	elapsed := time.Since(start)
	bestResult := max(results)
	meanResult := mean(results)

	stdResult := std(results)
	meanSteps := mean(stepCounts)

	fmt.Println(name, "elapsed: ", elapsed, "best: ", bestResult, "mean: ", meanResult, "steps(mean): ", meanSteps, "std result: ", stdResult)
	return elapsed
}

func computeHeuristic(distances [][]int){
	permutation := solveHeuristic(distances)
	result := getDistance(permutation, distances)
	fmt.Println("Heuristic result: ", result)
}

func computeRandom(distances [][]int, availableTime time.Duration) time.Duration{
	start := time.Now()
	permutation := solveRandom(distances, availableTime	)
	result := getDistance(permutation, distances)
	elapsed := time.Since(start)

	fmt.Println("Random, elapsed: ", elapsed, "result: ", result)
	return elapsed
}
