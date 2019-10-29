package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	SIZE := len(perm)

	for i := 1; i < SIZE; i++ {
		sum += distances[perm[i-1]][perm[i]]
	}

	sum += distances[perm[SIZE-1]][perm[0]]

	return sum
}

func getPartialDistance(perm []int, distances [][]int, start int, end int) int {
	SIZE := len(perm)
	numberOfSteps := end - start + 2
	sum := 0
	startIndex := (start - 1 + SIZE) % SIZE

	for i := 0; i < numberOfSteps; i++ {
		sum += distances[perm[(startIndex+i)%SIZE]][perm[(startIndex+i+1)%SIZE]]
	}

	if start == 0 && end == SIZE-1 {
		sum -= distances[perm[SIZE-1]][perm[0]]
	}

	return sum
}

func getPartialDistanceReversed(perm []int, distances [][]int, start int, end int) int {
	sum := 0
	perm = reversePermutationPart(perm, start, end)
	sum = getPartialDistance(perm, distances, start, end)
	perm = reversePermutationPart(perm, start, end)
	return sum
}

func reversePermutationPart(perm []int, start int, end int) []int {
	swapsNumber := (end - start) / 2

	for i := 0; i < swapsNumber; i++ {
		bitSwap(perm, start+i, end-i)
	}
	return perm
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

func maxOfArray(array []int) int {
	maxValue := math.MinInt32
	for _, val := range array {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

func minOfArray(array []float64) float64 {
	minValue := math.MaxFloat64
	for _, val := range array {
		if val < minValue {
			minValue = val
		}
	}
	return minValue
}

func mean(array []float64) float64 {
	sum := 0.0
	for _, val := range array {
		sum += val
	}
	return sum / float64(len(array))
}

func meanInt(array []int) float64 {
	sum := 0.0
	for _, val := range array {
		sum += float64(val)
	}
	return sum / float64(len(array))
}

func std(array []float64) float64 {
	meanValue := mean(array)

	sum := 0.0
	for _, val := range array {
		diff := float64(val) - meanValue
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(len(array)))
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

func countNeighborDistanceDifference(permutation []int, distances [][]int, start int, end int) int {
	actualPartialDistance := getPartialDistance(permutation, distances, start, end)
	newPartialDistance := getPartialDistanceReversed(permutation, distances, start, end)
	return actualPartialDistance - newPartialDistance
}

func countNeighborResult(permutation []int, start int, end int) int {
	// result =
	return 1
}

func tmpSwap(permutation []int, first int, second int) []int {
	tmp := permutation[first]
	permutation[first] = permutation[second]
	permutation[second] = tmp
	return permutation
}

func bitSwap(permutation []int, first int, second int) []int {
	if first != second {
		permutation[first] ^= permutation[second]
		permutation[second] ^= permutation[first]
		permutation[first] ^= permutation[second]
	}
	return permutation
}

func itoa(value int) string {
	return strconv.Itoa(value)
}

func ftoa(value float64) string {
	return fmt.Sprintf("%f", value)
}

func getQuality(result int, bestKnown int) float64 {
	return float64(result) / float64(bestKnown)
}