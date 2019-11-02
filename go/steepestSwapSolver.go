package main

func solveSwapSteepest(distances [][]int) ([]int, int, int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)
	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = findBestSwapNeighbor(permutation, distances)
		reviewedSolutionsNumber += reviewedNeighborSolutions
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount, reviewedSolutionsNumber
}

func solveOptimizedSwapSteepest(distances [][]int) ([]int, int, int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = findBestSwapNeighborOptimized(permutation, distances)
		reviewedSolutionsNumber += reviewedNeighborSolutions
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount, reviewedSolutionsNumber
}
