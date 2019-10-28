package main

import (
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

	distance := getDistance(permutation[0:len(permutation)], distances)
	partialDistance := getPartialDistance(permutation, distances, 0, len(permutation)-1)

	if distance != partialDistance {
		t.Error("GetPartialDistance error")
	}

	distance = 20
	partialDistance = getPartialDistance(permutation, distances, 0, 2) //6 + 7 + 6 + 1

	if distance != partialDistance {
		t.Error("GetPartialDistance error")
	}

	distance = 20
	partialDistance = getPartialDistance(permutation, distances, 1, 3) //6 + 7 + 6 + 1

	if distance != partialDistance {
		t.Error("GetPartialDistance error")
	}

	reversedDistance := getDistance(reversedPermutation[0:len(reversedPermutation)], distances)
	reversedPartialDistance := getPartialDistanceReversed(permutation, distances, 0, len(permutation))

	if reversedDistance != reversedPartialDistance {
		t.Error("GetPartialDistanceReversed error")
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
			swapped1 = tmpSwap(swapped1, i, j)
			swapped2 = bitSwap(swapped2, i, j)

			if !areSlicesEqual(swapped1, swapped2) {
				t.Error("swapped values are different")
			}
		}
	}
}