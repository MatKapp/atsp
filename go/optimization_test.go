package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// test greedy optimization (COMMENT PERMUTATAION SHUFFLE BEFORE THE TEST)
func TestGreedy(t *testing.T) {

	// if !AreGreedyResultsEqual("br17.atsp") {
	// 	t.Error("Greedy optimization broke the result")
	// }

	if !AreGreedyResultsEqual("ftv44.atsp") {
		t.Error("Greedy optimization broke the result")
	}
}

// test swap optimization (COMMENT PERMUTATAION SHUFFLE BEFORE THE TEST)
func TestSteepest(t *testing.T) {
	// if !AreSteepestResultsEqual("br17.atsp") {
	// 	t.Error("Steepest optimization broke the result")
	// }

	if !AreSteepestResultsEqual("ftv44.atsp") {
		t.Error("Steepest optimization broke the result")
	}
}

func AreGreedyResultsEqual(instanceName string) bool {
	distances := readData("../data/" + instanceName)
	greedyResult := []int{}
	optimizedGreedyResult := []int{}
	stepCount := 0

	start := time.Now()

	for i := 0; i <= 10000; i++ {
		greedyResult, stepCount, _, _ = solveSwapGreedy(distances, false)
	}
	elapsed := time.Since(start)
	fmt.Println(greedyResult)
	fmt.Println("greedyStepCount: " + itoa(stepCount))
	println(itoa(getDistance(greedyResult, distances)))
	log.Printf("greedy took %s", elapsed)

	start = time.Now()
	for i := 0; i <= 10000; i++ {
		optimizedGreedyResult, stepCount, _, _ = solveOptimizedSwapGreedy(distances, false)
	}
	elapsed = time.Since(start)
	fmt.Println(optimizedGreedyResult)
	fmt.Println("optimizedGreedyStepCount: " + itoa(stepCount))
	println(itoa(getDistance(optimizedGreedyResult, distances)))
	log.Printf("optimizedGreedy took %s", elapsed)

	return areSlicesEqual(greedyResult, optimizedGreedyResult)
}

func AreSteepestResultsEqual(instanceName string) bool {
	distances := readData("../data/" + instanceName)
	steepestResult := []int{}
	optimizedSteepestResult := []int{}

	start := time.Now()

	for i := 0; i <= 10000; i++ {
		steepestResult, _, _, _ = solveSwapSteepest(distances, false)
	}
	elapsed := time.Since(start)
	fmt.Println(steepestResult)
	log.Printf("steepest took %s", elapsed)

	start = time.Now()
	for i := 0; i <= 10000; i++ {
		optimizedSteepestResult, _, _ = solveOptimizedSwapSteepest(distances, false)
	}
	fmt.Println(optimizedSteepestResult)
	elapsed = time.Since(start)
	log.Printf("optimizedSteepest took %s", elapsed)

	return areSlicesEqual(steepestResult, optimizedSteepestResult)
}

func TestBestSwapNeighborOptimized(t *testing.T) {
	distances := readData("../data/br17.atsp")
	perm := []int{10, 7, 8, 16, 0, 11, 14, 3, 4, 5, 6, 2, 13, 12, 1, 9, 15}
	result, _ := findBestSwapNeighbor(perm, distances)
	resultOptimized, _, _ := findBestSwapNeighborOptimized(perm, distances, len(perm))

	fmt.Println(result)
	fmt.Println(resultOptimized)

	if !areSlicesEqual(result, resultOptimized) {
		t.Error("Best Swap optimization broke the result")
	}
}

func TestGetBetterNeighbor(t *testing.T) {
	distances := readData("../data/ftv44.atsp")
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	permutationOptimized := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	i := 0

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
		permutationOptimized[i] = i
	}
	// permutation = shuffle(permutation)
	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = findBetterSwapNeighbor(permutation, distances)
		permutationOptimized, _ = findBetterSwapNeighborOptimized(permutationOptimized, distances, SIZE)

		if !areSlicesEqual(permutation, permutationOptimized) {
			t.Error("Error optimization")
		}

		reviewedSolutionsNumber += reviewedNeighborSolutions
		newResult := getDistance(permutation, distances)
		i++

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
}

func TestGetBestNeighbor(t *testing.T) {
	distances := readData("../data/ftv44.atsp")
	SIZE := len(distances)
	result := makeArray(SIZE)
	optimizedResult := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		result[i] = i
		optimizedResult[i] = i
	}
	bestResult := getDistance(result, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		result, _ = findBestSwapNeighbor(result, distances)
		optimizedResult, _, _ = findBestSwapNeighborOptimized(optimizedResult, distances, SIZE)

		if !areSlicesEqual(result, optimizedResult) {
			fmt.Println("____________________________________________")
			t.Error("Results are not equal")
		}

		fmt.Println(result)
		fmt.Println(optimizedResult)
		newResult := getDistance(result, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
}
