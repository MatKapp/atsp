package main

import (
  "container/list"
  "math"
  // "fmt"
  )

func solveTabu(distances [][]int, stepProcessing bool) ([]int, int, int, [][]int)  {
	SIZE := len(distances)
  // TODO parametrize TABU_SIZE and WAIT_ITERATIONS
  TABU_SIZE := SIZE / 4
  WAIT_ITERATIONS := SIZE
	stepCount := 0
	reviewedSolutionsNumber := 0
	permutation := makePermutation(SIZE)
	// permutation = shuffle(permutation)

  bestPermutation := makeArray(SIZE)
  currentDistance := getDistance(permutation, distances)
  bestDistance := currentDistance
	tabuInProgress := true
  tabuList := list.New()

	for ok := true; ok; ok = tabuInProgress && stepCount < WAIT_ITERATIONS {
		tabuInProgress = false
		permutation, profit, swap, revSolutions := findTabuNeighbor(permutation, distances, SIZE, tabuList)

    stepCount++
		reviewedSolutionsNumber += revSolutions
    currentDistance -= profit

		if swap[0] != swap[1] {
      if currentDistance < bestDistance {
          copy(bestPermutation, permutation)
          bestDistance = currentDistance
          stepCount = 0
      }
			tabuInProgress = true

      tabuList.PushFront(swap)
      if tabuList.Len() > TABU_SIZE {
          tabuList.Remove(tabuList.Back())
      }
		}
	}
	return bestPermutation, stepCount, reviewedSolutionsNumber, [][]int{}
}

func findTabuNeighbor(permutation []int, distances [][]int, SIZE int, tabuList *list.List) ([]int, int, []int, int) {
	swap := []int{0, 0}
	bestProfit := math.MinInt32
  revSolutions := 0
	profit := 0

	for i := 0; i < SIZE - 1; i++ {
		for j := i + 1; j < SIZE; j++ {
			profit = countNeighborSwapProfit(permutation, distances, i, j, SIZE)
      revSolutions++

      // Conditions have been split for performance
			if profit > bestProfit {
        aspirationCriterion := profit > 0 || !listContains(tabuList, i, j)
        if aspirationCriterion {
  				bestProfit = profit
  				swap = []int{i, j}
        }
			}
		}
	}

	tmpSwap(permutation, swap[0], swap[1])
	return permutation, bestProfit, swap, revSolutions
}

func listContains(lst *list.List, v1 int, v2 int) bool{
  for e := lst.Front(); e != nil; e = e.Next() {
    el := e.Value.([]int)
		if el[0] == v1 && el[1] == v2 {
      return true
    }
	}
  return false
}
