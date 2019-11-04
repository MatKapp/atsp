package main

func solveSwapGreedy(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	var stepPermutations [][]int

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	// permutation = shuffle(permutation)
	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = findBetterSwapNeighbor(permutation, distances)
		reviewedSolutionsNumber += reviewedNeighborSolutions
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount, reviewedSolutionsNumber, stepPermutations
}

func solveOptimizedSwapGreedy(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	var stepPermutations [][]int

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	// permutation = shuffle(permutation)
	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = findBetterSwapNeighborOptimized(permutation, distances)
		reviewedSolutionsNumber += reviewedNeighborSolutions
		newResult := getDistance(permutation, distances)

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount, reviewedSolutionsNumber, stepPermutations
}
