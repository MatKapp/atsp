package main

import "testing"

// test greedy optimization
func TestGreedy(t *testing.T) {
	distances := readData("data/br17.atsp")
	SIZE := len(distances)
	permutation := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	greedyResult := solveGreedy(distances)
	optimizedGreedyResult := solveOptimizedGreedy(distances)

	if areSlicesEqual(greedyResult, optimizedGreedyResult) {
		t.Error("Greedy optimization break the result")
	}
}
