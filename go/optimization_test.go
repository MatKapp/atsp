package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// test greedy optimization (COMMENT PERMUTATAION SHUFFLE BEFORE THE TEST)
func TestGreedy(t *testing.T) {
	distances := readData("../data/br17.atsp")
	greedyResult := []int{}
	optimizedGreedyResult := []int{}

	start := time.Now()

	for i := 0; i <= 10000; i++ {
		greedyResult, _, _, _ = solveSwapGreedy(distances, false)
	}
	elapsed := time.Since(start)
	fmt.Println(greedyResult)
	log.Printf("greedy took %s", elapsed)

	start = time.Now()
	for i := 0; i <= 10000; i++ {
		optimizedGreedyResult, _, _, _ = solveOptimizedSwapGreedy(distances, false)
	}
	fmt.Println(optimizedGreedyResult)
	elapsed = time.Since(start)
	log.Printf("optimizedGreedy took %s", elapsed)

	if !areSlicesEqual(greedyResult, optimizedGreedyResult) {
		t.Error("Greedy optimization broke the result")
	}
}

func TestBestSwapNeighborOptimized(t *testing.T) {
	distances := readData("../data/br17.atsp")
	perm := []int{10, 7, 8, 16, 0, 11, 14, 3, 4, 5, 6, 2, 13, 12, 1, 9, 15}
	result, _ := findBestSwapNeighbor(perm, distances)
	resultOptimized, _ := findBestSwapNeighborOptimized(perm, distances)

	fmt.Println(result)
	fmt.Println(resultOptimized)

	if !areSlicesEqual(result, resultOptimized) {
		t.Error("Best Swap optimization broke the result")
	}
}

// test swap optimization (COMMENT PERMUTATAION SHUFFLE BEFORE THE TEST)
func TestSteepest(t *testing.T) {
	distances := readData("../data/br17.atsp")
	steepestResult := []int{}
	optimizedSteepestResult := []int{}

	start := time.Now()

	for i := 0; i <= 10000; i++ {
		steepestResult, _, _, _ = solveSwapSteepest(distances, false)
	}
	elapsed := time.Since(start)
	fmt.Println(steepestResult)
	log.Printf("greedy took %s", elapsed)

	start = time.Now()
	for i := 0; i <= 10000; i++ {
		optimizedSteepestResult, _, _ = solveOptimizedSwapSteepest(distances, false)
	}
	fmt.Println(optimizedSteepestResult)
	elapsed = time.Since(start)
	log.Printf("optimizedSteepest took %s", elapsed)

	if !areSlicesEqual(steepestResult, optimizedSteepestResult) {
		t.Error("Steepest optimization broke the result")
	}
}
