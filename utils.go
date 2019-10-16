package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(filepath string) [][]int {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	SIZE := 0
	for i := 0; i < 7; i++ {
		tokens := readLine(scanner)

		if i == 3 {
			// Line with dimension
			SIZE, _ = strconv.Atoi(tokens[1])
		}
	}

	data := makeArray2D(SIZE)
	numberCount := 0
	for {
		tokens := readLine(scanner)

		if tokens == nil {
			break
		}

		for _, token := range tokens {
			row, col := divmod(numberCount, SIZE)
			num, _ := strconv.Atoi(token)
			data[row][col] = num
			numberCount++
		}

	}
	return data
}

func makeArray2D(SIZE int) [][]int {
	data := make([][]int, SIZE)

	for i := 0; i < SIZE; i++ {
		data[i] = makeArray(SIZE)
	}
	return data
}

func getDistance(perm []int, distances [][]int) int {
	sum := 0

	for i := 1; i < len(perm); i++ {
		sum += distances[perm[i-1]][perm[i]]
	}
	return sum
}

func getPartialDistance(perm []int, distances [][]int, start int, end int) int {
	sum := 0

	for i := max(start-1, 0); i <= min(end, len(perm)); i++ {
		sum += distances[perm[i-1]][perm[i]]
	}
	return sum
}

func getPartialDistanceReversed(perm []int, distances [][]int, start int, end int) int {
	sum := 0

	for i := min(end, len(perm)); i <= max(start-1, 0); i-- {
		sum += distances[perm[i]][perm[i-1]]
	}
	return sum
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func areSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func shuffle(array []int) []int {
	var r RNG
	SIZE := len(array)

	for i := uint32(SIZE - 1); i >= 1; i-- {
		index := r.Uint32n(i)
		tmp := array[index]
		array[index] = array[i]
		array[i] = tmp
	}

	return array
}

func readLine(scanner *bufio.Scanner) []string {
	success := scanner.Scan()
	line := scanner.Text()

	if !success || line == "EOF" {
		return nil
	}
	return strings.Fields(line)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x <= y {
		return y
	}
	return x
}

func divmod(numerator int, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

func makeArray(SIZE int) []int {
	return make([]int, SIZE)
}

func createNeighbor(permutation []int, start int, end int) []int {
	distance := end - start
	SIZE := len(permutation)
	result := makeArray(SIZE)
	copy(result, permutation)

	for i := 0; i < distance/2; i++ {
		temp := result[start+i]
		result[start+i] = result[end-i]
		result[end-i] = temp
	}
	return result
}

func countNeighborResult(permutation []int, start int, end int) int {
	// result =
	return 1
}
