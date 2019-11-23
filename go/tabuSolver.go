package main

import (
  "container/list"
  "time"
  "math"
  )

const TABU_SIZE = 50
const TIME = 100000000

func solveTabu(distances [][]int, permutation []int) ([]int, int)  {
	SIZE := len(distances)
	// permutation := makePermutation(SIZE)
  start := time.Now()

  oldPermutation := makeArray(SIZE)
  bestPermutation := permutation
	tabuInProgress := true
  tabuList := list.New()

	for ok := true; ok; ok = tabuInProgress && time.Since(start) < TIME {
		tabuInProgress = false
    copy(oldPermutation, permutation)
		permutation, profit, swap := findTabuNeighbor(permutation, distances, SIZE, tabuList)

		if !arraysEqual(oldPermutation, permutation) {
      if profit > 0 {
           bestPermutation = permutation
      }
			tabuInProgress = true

      tabuList.PushFront(swap)
      if tabuList.Len() > TABU_SIZE {
          tabuList.Remove(tabuList.Back())
      }
		}
	}
	return bestPermutation, 0
}

func findTabuNeighbor(permutation []int, distances [][]int, SIZE int, tabuList *list.List) ([]int, int, []int) {
	neighborI := 0
	neighborJ := 0
	bestProfit := math.MinInt32
	neighborProfit := 0

	for i := 0; i < SIZE-1; i++ {
		for j := i + 1; j < SIZE; j++ {
			neighborProfit = countNeighborSwapProfit(permutation, distances, i, j, SIZE)

			if neighborProfit > bestProfit && !listContains(tabuList, []int{i, j}){
				bestProfit = neighborProfit
				neighborI = i
				neighborJ = j
			}
		}
	}

	tmpSwap(permutation, neighborI, neighborJ)
	return permutation, bestProfit, []int{neighborI, neighborJ}
}

// func findTabuNeighbor(permutation []int, distances [][]int, tabuList *list.List) ([]int, []int) {
//   // If there is no valid neighbor, returns the same permutation
// 	SIZE := len(permutation)
//
// 	result := makeArray(SIZE)
// 	bestNeighbor := []int{0, 0}
// 	copy(result, permutation)
// 	bestProfit := math.MinInt32
//
// 	for i := 0; i < SIZE; i++ {
// 		for j := 0; j < SIZE; j++ {
// 			neighborProfit := countNeighborDistanceDifference(permutation, distances, i, j, SIZE)
//
// 			if neighborProfit > bestProfit {//&& !listContains(tabuList, []int{i, j}) {
// 				bestProfit = neighborProfit
// 				bestNeighbor[0] = i
// 				bestNeighbor[1] = j
// 			}
// 		}
// 	}
//
// 	copy(result, createNeighbor(permutation, bestNeighbor[0], bestNeighbor[1]))
//
// 	return result, bestNeighbor
// }
//
//
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
