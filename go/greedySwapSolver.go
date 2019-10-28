package main

func solveSwapGreedy(distances [][]int) []int {
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
		permutation = findBetterSwapNeighbor(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
	return permutation
}

func solveOptimizedSwapGreedy(distances [][]int) []int {
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
		permutation = findBetterSwapNeighborOptimized(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
	return permutation
}

func findBetterSwapNeighborOptimized(permutation []int, distances [][]int) []int {
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

func findBetterSwapNeighbor(permutation []int, distances [][]int) []int {
	SIZE := len(permutation)
	result := makeArray(SIZE)
	copy(result, permutation)
	oldResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := tmpSwap(permutation, i, j)
			newResult := getDistance(neighbor, distances)

			if newResult < oldResult {
				copy(result, neighbor)
				return result
			}
		}
	}
	return result
}
