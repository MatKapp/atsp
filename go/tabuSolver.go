package main

import (
  "container/list"
  "math"
  "time"
)

const TABU_SIZE = 100000
const TIME = 100000000

func solveTabu(distances [][]int) ([]int, int)  {
	SIZE := len(distances)
	permutation := makePermutation(SIZE)
  start := time.Now()

	bestResult := getDistance(permutation, distances)
  bestPermutation := permutation
	tabuInProgress := true
  tabuList := list.New()
  tabuList.PushFront(permutation)

	for ok := true; ok; ok = tabuInProgress && time.Since(start) < TIME {
		tabuInProgress = false
		neighbor := findTabuNeighbor(permutation, distances, tabuList)
		newResult := getDistance(neighbor, distances)

		if !arraysEqual(permutation, neighbor) {
      permutation = neighbor
      if newResult < bestResult {
			     bestResult = newResult
           bestPermutation = permutation
      }
			tabuInProgress = true

      tabuList.PushFront(permutation)
      if tabuList.Len() > TABU_SIZE {
          tabuList.Remove(tabuList.Back())
      }
		}
	}
	return bestPermutation, 0
}

func findTabuNeighbor(permutation []int, distances [][]int, tabuList *list.List) []int {
  // If no neighbor is better, returns the same permutation
	SIZE := len(permutation)

	result := makeArray(SIZE)
	bestNeighbor := []int{0, 0}
	copy(result, permutation)
	bestProfit := math.MinInt32

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j)

      neigh := createNeighbor(permutation, i, j)
			if neighborProfit > bestProfit && !listContains(tabuList, neigh) {
				bestProfit = neighborProfit
				bestNeighbor[0] = i
				bestNeighbor[1] = j
			}
		}
	}

	copy(result, createNeighbor(permutation, bestNeighbor[0], bestNeighbor[1]))

	return result
}


func listContains(lst *list.List, value []int) bool{
  for e := lst.Front(); e != nil; e = e.Next() {

		if arraysEqual(e.Value.([]int), value) {
      return true
    }
	}
  return false
}

func arraysEqual(arr1 []int, arr2 []int) bool {
  for i, _ := range arr1 {
    if arr1[i] != arr2[i] {
      return false
    }
  }
  return true
}
