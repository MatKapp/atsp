package main

import (
  "container/list"
  "math"
  // "fmt"
  )

func solveTabu(distances [][]int, permutation []int) ([]int, int)  {
	SIZE := len(distances)
  // TODO parametrize TABU_SIZE and WAIT_ITERATIONS
  TABU_SIZE := SIZE / 4
  WAIT_ITERATIONS := SIZE
	// permutation := makePermutation(SIZE)
	// permutation = shuffle(permutation)

  bestPermutation := makeArray(SIZE)
  counter := 0
  currentDistance := getDistance(permutation, distances)
  bestDistance := currentDistance
	tabuInProgress := true
  tabuList := list.New()

	for ok := true; ok; ok = tabuInProgress && counter < WAIT_ITERATIONS {
		tabuInProgress = false
		permutation, profit, swap := findTabuNeighbor(permutation, distances, SIZE, tabuList)
    currentDistance -= profit
    // fmt.Println(counter, profit, bestDistance, currentDistance)
    counter += 1

		if swap[0] != swap[1] {
      if currentDistance < bestDistance {
          copy(bestPermutation, permutation)
          bestDistance = currentDistance
          counter = 0
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
	profit := 0

	for i := 0; i < SIZE - 1; i++ {
		for j := i + 1; j < SIZE; j++ {
			profit = countNeighborSwapProfit(permutation, distances, i, j, SIZE)

      // Conditions have been split for performance
			if profit > bestProfit {
        aspirationCriterion := profit > 0 || !listContains(tabuList, []int{i, j})
        if aspirationCriterion {
  				bestProfit = profit
  				neighborI = i
  				neighborJ = j
        }
			}
		}
	}

	tmpSwap(permutation, neighborI, neighborJ)
	return permutation, bestProfit, []int{neighborI, neighborJ}
}

func listContains(lst *list.List, value []int) bool{
  for e := lst.Front(); e != nil; e = e.Next() {
    el := e.Value.([]int)
		if el[0] == value[0] && el[1] == value[1] {
      return true
    }
	}
  return false
}
