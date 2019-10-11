package main

import(
	"fmt"
)

func main() {
	distances := readData("data/br17.atsp")
	fmt.Println(distances)
	permutation := solveRandom(distances)
	result := getDistance(permutation, distances)
	fmt.Println(permutation)
	fmt.Println(result)
}
