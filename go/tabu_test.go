package main

import (
	"testing"
	"fmt"
	"time"
)

func TestTabuNotWorse(t *testing.T) {
	SIZE := 45
	distances := readData("../data/ftv44.atsp")

	permutation := makePermutation(SIZE)

	for i := 0; i < 2000; i++ {
		permutation = shuffle(permutation)

		start := time.Now()
		tabuPerm, _ := solveTabu(distances, permutation)
		tabuTime := time.Since(start)

		start = time.Now()
		steepestPerm := solveOptimizedSwapSteepest(distances, permutation)
		steepestTime := time.Since(start)

		tabuResult := getDistance(tabuPerm, distances)
		steepestResult := getDistance(steepestPerm, distances)


		if tabuResult > steepestResult {
			fmt.Println("Tabu is worse")
		  t.Error("Tabu can't be worse")
		} else if tabuResult < steepestResult{
			fmt.Println("Tabu is better")
			fmt.Println(tabuResult, steepestResult)
			fmt.Println(tabuTime, steepestTime)
		} else{
			fmt.Println("Equal")
		}
	}
}
