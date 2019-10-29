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

func findBetterReverseNeighborOptimized(permutation []int, distances [][]int) []int {
	SIZE := len(permutation)
	result := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j)

			if neighborProfit > 0 {
				neighbor := createNeighbor(permutation, i, j)
				copy(result, neighbor)
				return result
			}
		}
	}
	return permutation
}

func findBetterReverseNeighbor(permutation []int, distances [][]int) []int {
	SIZE := len(permutation)
	result := makeArray(SIZE)
	oldResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := createNeighbor(permutation, i, j)
			newResult := getDistance(neighbor, distances)

			if newResult < oldResult {
				copy(result, neighbor)
				return result
			}
		}
	}
	return permutation
}
