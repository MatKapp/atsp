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
		permutation = findBestNeighbor(permutation, distances)
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
		permutation = findBestNeighborOptimized(permutation, distances)
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
		}
	}
	return permutation
}

func findBestNeighbor(permutation []int, distances [][]int) []int {
	SIZE := len(permutation)
	result := makeArray(SIZE)
	copy(result, permutation)
	bestResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := createNeighbor(permutation, i, j)
			newResult := getDistance(neighbor, distances)

			if newResult < bestResult {
				bestResult = newResult
				copy(result, neighbor)
			}
		}
	}
	return result
}

func findBestNeighborOptimized(permutation []int, distances [][]int) []int {
	SIZE := len(permutation)

	result := makeArray(SIZE)
	bestNeighbor := []int{0, 0}
	copy(result, permutation)
	bestProfit := 0

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j)

			if neighborProfit > bestProfit {
				bestProfit = neighborProfit
				bestNeighbor[0] = i
				bestNeighbor[1] = j
			}
		}
	}

	copy(result, createNeighbor(permutation, bestNeighbor[0], bestNeighbor[1]))

	return result
}
