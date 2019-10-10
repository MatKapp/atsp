package main

import(
	"fmt"
)

func main() {
	data := readData("data/br17.atsp")
	fmt.Println(data)
	result := solveRandom(data)
	fmt.Println(result)
}
