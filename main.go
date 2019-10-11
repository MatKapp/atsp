package main

import(
	"fmt"
)

func main() {
	distances := readData("data/br17.atsp")

	permutation := solveRandom(distances)
	result := getDistance(permutation, distances)
	fmt.Println("Random")
	fmt.Println(permutation)
	fmt.Println(result)

	permutation = solveHeuristic(distances)
	result = getDistance(permutation, distances)
	fmt.Println("Heuristic")
	fmt.Println(permutation)
	fmt.Println(result)
}
