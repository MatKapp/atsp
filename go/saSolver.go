package main

func solveSaSolver(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int) {
	SIZE := len(distances)
	permutation := makeArray(SIZE)
	bestPermutation := makeArray(SIZE)
	stepCount := 0
	reviewedSolutionsNumber := 0
	var stepPermutations [][]int

	saTemperature := 0.95
	saCoolingCoefficient := .9
	markovChainLength := newtonSymbol(uint64(SIZE), uint64(2))
	stopCoefficient := uint64(10)
	stepsWithoutImprovementBeforeStop := markovChainLength * stopCoefficient

	//Start processing with swap greedy result
	// permutation, _, _, _ = solveOptimizedSwapGreedy(distances, stepProcessing)

	for i := 0; i < SIZE; i++ {
		permutation[i] = i
	}
	permutation = shuffle(permutation)

	copy(bestPermutation, permutation)
	bestResult := getDistance(permutation, distances)
	stepsWithoutImprovement := uint64(0)
	stepsWithoutTemperatureChange := uint64(0)

	for stepsWithoutImprovement <= stepsWithoutImprovementBeforeStop || saTemperature > .01 {
		// for saTemperature > .01 {
		reviewedNeighborSolutions := 0
		permutation, reviewedNeighborSolutions = saFindBetterSwapNeighborOptimized(permutation, distances, SIZE, saTemperature)
		stepsWithoutTemperatureChange++

		if stepsWithoutTemperatureChange >= markovChainLength && stepsWithoutImprovement >= markovChainLength {
			saTemperature = saCoolingCoefficient * saTemperature
			stepsWithoutTemperatureChange = 0
		}

		if stepProcessing {
			processingStep := makeArray(SIZE)
			copy(processingStep, permutation)
			stepPermutations = append(stepPermutations, processingStep)
		}
		reviewedSolutionsNumber += reviewedNeighborSolutions

		newResult := getDistance(permutation, distances)
		profit := bestResult - newResult

		if profit >= 0 {
			bestResult = newResult
			copy(bestPermutation, permutation)
		}

		stepsWithoutImprovement++
		stepCount++

		if profit > 0 {
			stepsWithoutImprovement = uint64(0)
		}
	}

	return bestPermutation, stepCount, reviewedSolutionsNumber, stepPermutations
}
