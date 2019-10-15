package main

import (
	"math"
)

func solveHeuristic(distances [][]int) ([]int, int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	set := make(map[int]bool) // New empty set

	current := 0

	for i := 0; i < SIZE; i++ {
		set[current] = true // Add
		permutation[i] = current
		next := findNext(current, distances, set, SIZE)
		current = next
	}
	return permutation, 1
}

func findNext(index int, distances [][]int, set map[int]bool, SIZE int) int {
	minDistance := math.MaxInt32
	next := 0

	for i := 0; i < SIZE; i++ {
		distance := distances[index][i]

		if !set[i] && distance < minDistance {
			minDistance = distance
			next = i
		}
	}
	return next
}
