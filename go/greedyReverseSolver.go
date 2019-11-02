package main

func solveReverseGreedy(distances [][]int) []int {
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
		permutation = findBetterReverseNeighbor(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
	return permutation
}

func solveOptimizedReverseGreedy(distances [][]int) []int {
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
		permutation = findBetterReverseNeighborOptimized(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
	return permutation
}
