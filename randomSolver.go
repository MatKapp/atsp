package main

import (
  "fmt"
)

func solveRandom(distances [][]int) []int{
  minDistance := 99999
  SIZE := len(distances)
  bestPermutation := make([]int, SIZE)
  array := make([]int, SIZE)

  for i := 0; i < 1000; i++{
    array = shuffle(array)
    distance := getDistance(array, distances)
    if distance < minDistance{
      minDistance = distance
      fmt.Println(minDistance)
      copy(bestPermutation, array)
    }
  }
  return bestPermutation
}
