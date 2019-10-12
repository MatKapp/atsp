package main

import(
	"fmt"
)

type fn func([][]int) []int

func main() {
	distances := readData("data/br17.atsp")
	compute(solveRandom, distances, "Random")
	compute(solveHeuristic, distances, "Heuristic")
}

func compute(solve func([][]int) []int, distances [][]int, name string) {
	permutation := solve(distances)
	result := getDistance(permutation, distances)
	fmt.Println(name)
	fmt.Println(permutation)
	fmt.Println(result)
}
