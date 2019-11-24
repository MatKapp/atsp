package main

func solveSaSolver(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	var stepPermutations [][]int

	saTemperature := 1.
	saCoolingCoefficient := 0.98

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	//Start processing with swap greedy result
	permutation, _, _, _ = solveOptimizedSwapGreedy(distances, stepProcessing)
	bestResult := getDistance(permutation, distances)
	resultImproved := true

	for ok := true; ok; ok = resultImproved {
		resultImproved = false
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = findBetterSwapNeighborOptimized(permutation, distances, SIZE)

		if stepProcessing {
			processingStep := makeArray(SIZE)
			copy(processingStep, permutation)
			stepPermutations = append(stepPermutations, processingStep)
		}
		reviewedSolutionsNumber += reviewedNeighborSolutions

		newResult := getDistance(permutation, distances)
		profit := newResult - bestResult

		saProbability := saTemperature * float64((1 / max(1, abs(profit))))
		updateResult := false

		if profit > 0 {
			updateResult = true
		} else {
			updateResult = randBool(saProbability)
			saTemperature *= saCoolingCoefficient
		}

		if updateResult {
			bestResult = newResult
			resultImproved = true
			stepCount++
		}
	}

	return permutation, stepCount, reviewedSolutionsNumber, stepPermutations
}
