package main

import (
	"math"
	"time"
)

func solveRandom(distances [][]int, availableTime time.Duration) []int {
	start := time.Now()
	minDistance := math.MaxInt32
	SIZE := len(distances)
	bestPermutation := makeArray(SIZE)
	array := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		array[i] = i
	}

	for {
		array = shuffle(array)
		distance := getDistance(array, distances)

		if distance < minDistance {
			minDistance = distance
			copy(bestPermutation, array)
		}
		if time.Since(start) > availableTime {
			break
		}
	}
	return bestPermutation
}
