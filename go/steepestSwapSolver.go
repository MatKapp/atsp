package main

func solveSwapSteepest(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	var stepPermutations [][]int

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

		if stepProcessing {
			stepPermutations = append(stepPermutations, permutation)
		}

		if newResult < bestResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount, reviewedSolutionsNumber, stepPermutations
}

func solveOptimizedSwapSteepest(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	reviewedNeighborSolutions := 0
	bestProfit := 0
	var stepPermutations [][]int

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		permutation, reviewedNeighborSolutions, bestProfit = findBestSwapNeighborOptimized(permutation, distances, SIZE)
		reviewedSolutionsNumber += reviewedNeighborSolutions

		if stepProcessing {
			stepPermutations = append(stepPermutations, permutation)
		}

		if bestProfit > 0 {
			resultImproved = true
			stepCount++
		}
	}
	return permutation, stepCount, reviewedSolutionsNumber, stepPermutations
}
