package main

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
      copy(bestPermutation, array)
    }
  }
  return bestPermutation
}
