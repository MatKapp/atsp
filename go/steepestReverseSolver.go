package main

func solveReverseSteepest(distances [][]int) ([]int, int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		permutation = findBestReverseNeighbor(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount
}

func solveOptimizedReverseSteepest(distances [][]int) []int {
	SIZE := len(distances)
	permutation := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		permutation = findBestReverseNeighborOptimized(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
	return permutation
}
