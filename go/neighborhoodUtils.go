package main

import (
	"math/rand"
)

func findBestSwapNeighbor(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)
	copy(result, permutation)
	bestResult := getDistance(permutation, distances)

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			tmpSwap(permutation, i, j)
			newResult := getDistance(permutation, distances)
			reviewedSolutionsNumber++

			if newResult < bestResult {
				bestResult = newResult
				copy(result, permutation)
			}
			tmpSwap(permutation, i, j)
		}
	}
	return result, reviewedSolutionsNumber
}

func findBestSwapNeighborOptimized(permutation []int, distances [][]int, SIZE int) ([]int, int, int) {
	reviewedSolutionsNumber := 0
	neighborI := 0
	neighborJ := 0
	bestProfit := 0
	neighborProfit := 0

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			neighborProfit = countNeighborSwapProfit(permutation, distances, i, j, SIZE)
			reviewedSolutionsNumber++

			if neighborProfit > bestProfit {
				bestProfit = neighborProfit
				neighborI = i
				neighborJ = j
			}
		}
	}

	tmpSwap(permutation, neighborI, neighborJ)

	return permutation, reviewedSolutionsNumber, bestProfit
}

func findBestReverseNeighbor(permutation []int, distances [][]int) ([]int, int) {
	SIZE := len(permutation)
	reviewedSolutions := 0
	result := makeArray(SIZE)
	copy(result, permutation)
	bestResult := getDistance(permutation, distances)

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
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

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j, SIZE)
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

func findBetterSwapNeighborOptimized(permutation []int, distances [][]int, SIZE int) ([]int, int) {
	reviewedSolutionsNumber := 0

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			neighborProfit := countNeighborSwapProfit(permutation, distances, i, j, SIZE)
			reviewedSolutionsNumber++

			if neighborProfit > 0 {
				tmpSwap(permutation, i, j)
				return permutation, reviewedSolutionsNumber
			}
		}
	}
	return permutation, reviewedSolutionsNumber
}

func saFindBetterSwapNeighborOptimized(permutation []int, distances [][]int, SIZE int, temperature float64) ([]int, int) {
	reviewedSolutionsNumber := 0
	offset := rand.Intn(SIZE)
	iWithOffset := 0
	jWithOffset := 0

	for i := 0; i < SIZE-1; i++ {
		iWithOffset = (i + offset) % (SIZE - 1)
		for j := i + 1; j < SIZE; j++ {

			jWithOffset = (j + offset) % SIZE

			neighborProfit := countNeighborSwapProfit(permutation, distances, iWithOffset, jWithOffset, SIZE)
			reviewedSolutionsNumber++

			if neighborProfit > 0 || randBool(temperature*float64((1/min(4, max(1, abs(neighborProfit)))))) {
				tmpSwap(permutation, iWithOffset, jWithOffset)
				return permutation, reviewedSolutionsNumber
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

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			tmpSwap(permutation, i, j)
			newResult := getDistance(permutation, distances)
			reviewedSolutionsNumber++

			if newResult < oldResult {
				copy(result, permutation)
				return result, reviewedSolutionsNumber
			}
			tmpSwap(permutation, i, j)
		}
	}
	return result, reviewedSolutionsNumber
}

func findBetterReverseNeighborOptimized(permutation []int, distances [][]int, SIZE int) ([]int, int) {
	reviewedSolutionsNumber := 0

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j, SIZE)
			reviewedSolutionsNumber++

			if neighborProfit > 0 {
				permutation = reversePart(permutation, i, j, SIZE)
				return permutation, reviewedSolutionsNumber
			}
		}
	}
	return permutation, reviewedSolutionsNumber
}

func findBetterReverseNeighbor(permutation []int, distances [][]int, SIZE int) ([]int, int) {
	reviewedSolutionsNumber := 0
	result := makeArray(SIZE)
	oldResult := getDistance(permutation, distances)

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
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
	distance := end - start + 1
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

func reversePart(permutation []int, start int, end int, SIZE int) []int {
	distance := end - start + 1
	temp := 0

	for i := 0; i < distance/2; i++ {
		temp = permutation[start+i]
		permutation[start+i] = permutation[end-i]
		permutation[end-i] = temp
	}
	return permutation
}

func countNeighborDistanceDifference(permutation []int, distances [][]int, start int, end int, SIZE int) int {
	actualPartialDistance := getPartialDistance(permutation, distances, start, end, SIZE)
	newPartialDistance := getPartialDistanceReversedOptimized(permutation, distances, start, end, SIZE)

	return actualPartialDistance - newPartialDistance
}

func countNeighborSwapProfit(perm []int, distances [][]int, start int, end int, SIZE int) int {

	if start == end {
		return 0
	}

	// correct in case of swap neighborhood
	if start > end {
		temp := start
		start = end
		end = temp
	}

	actualPartialDistance := 0
	newPartialDistance := 0

	areNeighbors := end == start+1
	areOverfllowNeighbors := end-start == SIZE-1

	if !areOverfllowNeighbors {
		actualPartialDistance += distances[perm[(start-1+SIZE)%SIZE]][perm[start]]
	}

	if !areNeighbors {
		actualPartialDistance += distances[perm[start]][perm[start+1]]
	}

	actualPartialDistance += distances[perm[end-1]][perm[end]]
	actualPartialDistance += distances[perm[end]][perm[(end+1)%SIZE]]

	perm = tmpSwap(perm, start, end)

	if !areOverfllowNeighbors {
		newPartialDistance += distances[perm[(start-1+SIZE)%SIZE]][perm[start]]
	}

	if !areNeighbors {
		newPartialDistance += distances[perm[start]][perm[start+1]]
	}

	newPartialDistance += distances[perm[end-1]][perm[end]]
	newPartialDistance += distances[perm[end]][perm[(end+1)%SIZE]]

	perm = tmpSwap(perm, start, end)

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
