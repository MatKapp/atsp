package main

import (
	"log"
	"testing"
	"time"
)

// test greedy optimization (COMMENT PERMUTATAION SHUFFLE BEFORE THE TEST)
func TestGreedy(t *testing.T) {
	distances := readData("data/br17.atsp")
	greedyResult := []int{}
	optimizedGreedyResult := []int{}

	start := time.Now()
	for i := 0; i <= 1000; i++ {
		greedyResult = solveGreedy(distances)
	}
	elapsed := time.Since(start)
	log.Printf("greedy took %s", elapsed)

	start = time.Now()
	for i := 0; i <= 1000; i++ {
		optimizedGreedyResult = solveOptimizedGreedy(distances)
	}
	elapsed = time.Since(start)
	log.Printf("optimizedGreedy took %s", elapsed)

	if !areSlicesEqual(greedyResult, optimizedGreedyResult) {
		t.Error("Greedy optimization broke the result")
	}
}
