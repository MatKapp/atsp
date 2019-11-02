package main

func findBestSwapNeighbor(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)
	copy(result, permutation)
	bestResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := createNeighbor(permutation, i, j)
			newResult := getDistance(neighbor, distances)
			reviewedSolutionsNumber++

			if newResult < bestResult {
				bestResult = newResult
				copy(result, neighbor)
			}
		}
	}
	return result, reviewedSolutionsNumber
}

func findBestSwapNeighborOptimized(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)
	bestNeighbor := []int{0, 0}
	copy(result, permutation)
	bestProfit := 0

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j)
			reviewedSolutionsNumber++

			if neighborProfit > bestProfit {
				bestProfit = neighborProfit
				bestNeighbor[0] = i
				bestNeighbor[1] = j
			}
		}
	}

	copy(result, createNeighbor(permutation, bestNeighbor[0], bestNeighbor[1]))

	return result, reviewedSolutionsNumber
}

func findBestReverseNeighbor(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutions := 0
	result := makeArray(SIZE)
	copy(result, permutation)
	bestResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := createNeighbor(permutation, i, j)
			newResult := getDistance(neighbor, distances)
			reviewedSolutions++

			if newResult < bestResult {
				bestResult = newResult
				copy(result, neighbor)
			}
		}
	}
	return result, reviewedSolutions
}

func findBestReverseNeighborOptimized(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0

	result := makeArray(SIZE)
	bestNeighbor := []int{0, 0}
	copy(result, permutation)
	bestProfit := 0

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j)
			reviewedSolutionsNumber++

			if neighborProfit > bestProfit {
				bestProfit = neighborProfit
				bestNeighbor[0] = i
				bestNeighbor[1] = j
			}
		}
	}

	copy(result, createNeighbor(permutation, bestNeighbor[0], bestNeighbor[1]))

	return result, reviewedSolutionsNumber
}

func findBetterSwapNeighborOptimized(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborSwapProfit(permutation, distances, i, j)
			reviewedSolutionsNumber++

			if neighborProfit > 0 {
				neighbor := createNeighbor(permutation, i, j)
				copy(result, neighbor)
				return result, reviewedSolutionsNumber
			}
		}
	}
	return permutation, reviewedSolutionsNumber
}

func findBetterSwapNeighbor(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)
	copy(result, permutation)
	oldResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := tmpSwap(permutation, i, j)
			newResult := getDistance(neighbor, distances)
			reviewedSolutionsNumber++

			if newResult < oldResult {
				copy(result, neighbor)
				return result, reviewedSolutionsNumber
			}
		}
	}
	return result, reviewedSolutionsNumber
}

func findBetterReverseNeighborOptimized(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j)
			reviewedSolutionsNumber++

			if neighborProfit > 0 {
				neighbor := createNeighbor(permutation, i, j)
				copy(result, neighbor)
				return result, reviewedSolutionsNumber
			}
		}
	}
	return permutation, reviewedSolutionsNumber
}

func findBetterReverseNeighbor(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)
	oldResult := getDistance(permutation, distances)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighbor := createNeighbor(permutation, i, j)
			newResult := getDistance(neighbor, distances)
			reviewedSolutionsNumber++

			if newResult < oldResult {
				copy(result, neighbor)
				return result, reviewedSolutionsNumber
			}
		}
	}
	return permutation, reviewedSolutionsNumber
}

func createNeighbor(permutation []int, start int, end int) []int {
	distance := end - start
	SIZE := len(permutation)
	result := makeArray(SIZE)
	copy(result, permutation)

	for i := 0; i < distance/2; i++ {
		temp := result[start+i]
		result[start+i] = result[end-i]
		result[end-i] = temp
	}
	return result
}

func countNeighborDistanceDifference(permutation []int, distances [][]int, start int, end int) int {
	actualPartialDistance := getPartialDistance(permutation, distances, start, end)
	newPartialDistance := getPartialDistanceReversed(permutation, distances, start, end)
	return actualPartialDistance - newPartialDistance
}

func countNeighborSwapProfit(perm []int, distances [][]int, start int, end int) int {
	SIZE := len(perm)

	if start == end {
		return 0
	}

	actualPartialDistance := 0
	newPartialDistance := 0

	actualPartialDistance += distances[perm[(start-1+SIZE)%SIZE]][perm[start]]
	actualPartialDistance += distances[perm[start]][perm[start+1]]
	actualPartialDistance += distances[perm[end-1]][perm[end]]
	actualPartialDistance += distances[perm[end]][perm[(end+1)%SIZE]]

	newPartialDistance += distances[perm[(start-1+SIZE)%SIZE]][perm[end]]
	newPartialDistance += distances[perm[end]][perm[start+1]]
	newPartialDistance += distances[perm[end-1]][perm[start]]
	newPartialDistance += distances[perm[start]][perm[(end+1)%SIZE]]

	return actualPartialDistance - newPartialDistance
}

func tmpSwap(permutation []int, first int, second int) []int {
	tmp := permutation[first]
	permutation[first] = permutation[second]
	permutation[second] = tmp
	return permutation
}

func bitSwap(permutation []int, first int, second int) []int {
	if first != second {
		permutation[first] ^= permutation[second]
		permutation[second] ^= permutation[first]
		permutation[first] ^= permutation[second]
	}
	return permutation
}
