package main

import (
	"testing"
	"fmt"
	"time"
)

func TestTabuNotWorse(t *testing.T) {
	SIZE := 34
	distances := readData("../data/ftv33.atsp")

	permutation := makePermutation(SIZE)
	perm2 := makePermutation(SIZE)

	tabuSum := 0
	tabuBetter := 0
	tabuWorse := 0

	for i := 0; i < 1000; i++ {
		permutation = shuffle(permutation)
		copy(perm2, permutation)

		start := time.Now()
		tabuPerm, _ := solveTabu(distances, permutation)
		tabuTime := time.Since(start)

		start = time.Now()
		steepestPerm := solveOptimizedSwapSteepest(distances, perm2)
		steepestTime := time.Since(start)

		tabuResult := getDistance(tabuPerm, distances)
		steepestResult := getDistance(steepestPerm, distances)

		fmt.Println(tabuResult, steepestResult)
		fmt.Println(tabuTime, steepestTime)

		tabuSum += tabuResult
		if tabuResult > steepestResult {
			tabuWorse++
		} else if tabuResult < steepestResult{
			tabuBetter++
		}
	}
	// 1694748 868 0
	fmt.Println(tabuSum, tabuBetter, tabuWorse)
}
