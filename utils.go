package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
	"fmt"
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
	sum += distances[SIZE - 1][0]
	return sum
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

func divmod(numerator int, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

func makeArray(SIZE int) []int {
	return make([]int, SIZE)
}

func max(array []int) int{
	maxValue := math.MinInt32
	for _, val := range array{
		if val > maxValue{
			maxValue = val
		}
	}
	return maxValue
}

func min(array []int) int{
	minValue := math.MaxInt32
	for _, val := range array{
		if val < minValue{
			minValue = val
		}
	}
	return minValue
}

func mean(array []int) float64 {
	sum := 0.0
	for _, val := range array{
		sum += float64(val)
	}
	return sum / float64(len(array))
}

func std(array []int) float64 {
	meanValue := mean(array)

	sum := 0.0
	for _, val := range array{
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

func itoa(value int) string{
	return strconv.Itoa(value)
}

func ftoa(value float64) string{
	return fmt.Sprintf("%f", value)
}
