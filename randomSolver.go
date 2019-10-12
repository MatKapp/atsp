package main

import (
  "math"
)

func solveRandom(distances [][]int) []int{
  minDistance := math.MaxInt32
  SIZE := len(distances)
  bestPermutation := makeArray(SIZE)
  array := makeArray(SIZE)
  for i := 0; i < SIZE; i++ {
    array[i] = i
  }

  for i := 0; i < 1000; i++{
    array = shuffle(array)
    distance := getDistance(array, distances)
    if distance < minDistance{
      minDistance = distance
      copy(bestPermutation, array)
    }
  }
  return bestPermutation
}
