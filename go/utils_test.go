package main

import (
	"fmt"
	"math"
	"testing"
)

func TestGetPartialDistance(t *testing.T) {
	distances := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 8, 7},
		{6, 5, 6, 7},
	}

	permutation := []int{1, 2, 3, 0}
	reversedPermutation := []int{0, 3, 2, 1} //3 + 6 + 4
	SIZE := len(permutation)

	distance := getDistance(permutation[0:len(permutation)], distances)
	partialDistance := getPartialDistance(permutation, distances, 0, len(permutation)-1, SIZE)

	if distance != partialDistance {
		t.Error("GetPartialDistance error")
	}

	distance = 20
	partialDistance = getPartialDistance(permutation, distances, 0, 2, SIZE) //6 + 7 + 6 + 1

	if distance != partialDistance {
		t.Error("GetPartialDistance error")
	}

	distance = 20
	partialDistance = getPartialDistance(permutation, distances, 1, 3, SIZE) //6 + 7 + 6 + 1

	if distance != partialDistance {
		t.Error("GetPartialDistance error")
	}

	reversedDistance := getDistance(reversedPermutation[0:len(reversedPermutation)], distances)
	reversedPartialDistance := getPartialDistanceReversed(permutation, distances, 0, SIZE, SIZE)

	if reversedDistance != reversedPartialDistance {
		t.Error("GetPartialDistanceReversed error")
	}
}

func TestCreateNeighbor(t *testing.T) {
	perm := []int{0, 1, 2, 3, 4, 5, 6}
	neighbor := []int{0, 1, 2, 4, 3, 5, 6}

	result := createNeighbor(perm, 3, 4)

	if !areSlicesEqual(neighbor, result) {
		t.Error("create neighbor error")
	}

	neighbor = []int{0, 1, 4, 3, 2, 5, 6}

	result = createNeighbor(perm, 2, 4)

	if !areSlicesEqual(neighbor, result) {
		t.Error("create neighbor error")
	}

	neighbor = []int{6, 5, 4, 3, 2, 1, 0}

	result = createNeighbor(perm, 0, 6)

	if !areSlicesEqual(neighbor, result) {
		t.Error("create neighbor error")
	}

	neighbor = []int{1, 0, 2, 3, 4, 5, 6}

	result = createNeighbor(perm, 0, 1)

	if !areSlicesEqual(neighbor, result) {
		t.Error("create neighbor error")
	}

	neighbor = []int{0, 1, 2, 3, 4, 6, 5}

	result = createNeighbor(perm, 5, 6)

	if !areSlicesEqual(neighbor, result) {
		t.Error("create neighbor error")
	}
}

func TestGetPartialDistanceReversedOptimized(t *testing.T) {
	distances := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 8, 7},
		{6, 5, 6, 7},
	}

	permutation := []int{1, 2, 3, 0}
	SIZE := len(permutation)

	reversedPartialDistance := getPartialDistanceReversed(permutation, distances, 1, 2, SIZE)
	reversedPartialDistanceOptimized := getPartialDistanceReversedOptimized(permutation, distances, 1, 2, SIZE)

	if reversedPartialDistance != reversedPartialDistanceOptimized {
		t.Error("ReversedPartialDistance  error")
	}

	reversedPartialDistance = getPartialDistanceReversed(permutation, distances, 0, 3, SIZE)
	reversedPartialDistanceOptimized = getPartialDistanceReversedOptimized(permutation, distances, 0, 3, SIZE)

	if reversedPartialDistance != reversedPartialDistanceOptimized {
		t.Error("ReversedPartialDistance  error")
	}
}
func TestReversePermutationPart(t *testing.T) {
	perm := []int{1, 2, 3, 4}
	SIZE := len(perm)
	reversed := makeArray(SIZE)
	copy(reversed, perm)

	reversed = reversePermutationPart(reversed, 0, 2)

	if !areSlicesEqual([]int{3, 2, 1, 4}, reversed) {
		t.Error("ReversePermutationPart error")
	}
}

func TestSwap(t *testing.T) {
	perm := []int{1, 2, 3, 4, 5, 6, 7}
	SIZE := len(perm)
	swapped1 := makeArray(SIZE)
	swapped2 := makeArray(SIZE)
	copy(swapped1, perm)
	copy(swapped2, perm)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			tmpSwap(swapped1, i, j)
			swapped2 = bitSwap(swapped2, i, j)

			if !areSlicesEqual(swapped1, swapped2) {
				t.Error("swapped values are different")
			}
		}
	}
}

func TestCountNeighborSwapProfit(t *testing.T) {
	distances := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 8, 7},
		{6, 5, 6, 7},
	}
	perm := []int{0, 2, 1, 3}
	SIZE := len(perm)

	profit := countNeighborSwapProfit(perm, distances, 0, 2, SIZE) // (9 + 7 + 6 + 2) - (5 + 6 + 8 + 3) = 24 - 22 = 2

	fmt.Println(profit)

	if profit != 2 {
		t.Error("Profit error")
	}

	perm2 := []int{1, 2, 44, 39, 40, 41, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
		18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 3, 4, 5, 42, 43, 0}
	distances2 := readData("../data/ftv44.atsp")

	value := getDistance(perm2, distances2)
	tmpSwap(perm2, 8, 9)
	newValue := getDistance(perm2, distances2)
	tmpSwap(perm2, 8, 9)

	if countNeighborSwapProfit(perm2, distances2, 8, 9, len(perm2)) != value-newValue {
		t.Error("Neighbor swap error")
	}
}

func TestSimilarity(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{7, 6, 5, 4, 3, 2, 1}

	result := countSimilarity(a, b)

	if result != 0 {
		t.Error("Similarity failed")
	}

	a = []int{1, 2, 3, 4, 5, 7, 6}
	b = []int{7, 6, 5, 4, 3, 2, 1}

	result = countSimilarity(a, b)

	if math.Round(result*100)/100 != math.Round(1/7.0*100)/100 {
		t.Error("Similarity failed")
	}
}
