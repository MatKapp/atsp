package main

import (
	"fmt"
	"time"
	"os"
	"encoding/csv"
)

type fn func([][]int) ([]int, int)

const MIN_TIME = 1000 * 1000
const RUN_COUNT = 10

func main() {
	bestKnownSolutions := map[string]int{
		"br17": 39,
		"ft53": 6905,
		"ftv33": 1286,
		"ft70": 38673,
		"ftv35": 1473,
		"ftv44": 1613,
		"ftv38": 1530,
		"ftv55": 1608,
		"ftv47": 1776,
		"ftv64": 1839,
		"ftv70": 1950,
		"ftv100": 1788,
		"ftv110": 1958,
		"ftv90": 1579,
		"ftv120": 2166,
		"ftv130": 2307,
		"ftv140": 2420,
		"ftv150": 2611,
		"ftv170": 2755,
		"ftv160": 2683,
		"kro124": 36230,
		"p43": 5620,
		"rbg358": 1163,
		"rbg323": 1326,
		"rbg403": 2465,
		"rbg443": 2720,
		"ry48p": 14422,
	}
	instanceFilenames := []string{
		"br17",
		"ft53",
		"ft70",
		"ftv33",
		"ftv35",
		"ftv38",
		"ftv44",
		"ftv47",
		"ftv55",
		"ftv70",
	}

	gFile, greedyWriter := getWriter("results/greedy.csv")
	sFile, steepestWriter := getWriter("results/steepest.csv")
	hFile, heuristicWriter := getWriter("results/heuristic.csv")
	rFile, randomWriter := getWriter("results/random.csv")

	defer gFile.Close()
	defer sFile.Close()
	defer hFile.Close()
	defer rFile.Close()
	defer greedyWriter.Flush()
	defer steepestWriter.Flush()
	defer heuristicWriter.Flush()
	defer randomWriter.Flush()

	for _, filename := range instanceFilenames {
		fmt.Println()
		fmt.Println(filename)
		path := "data/" + filename + ".atsp"
		distances := readData(path)

		bestKnown := bestKnownSolutions[filename]
		fmt.Println("Best known: ", bestKnown)

		hOutput := computeHeuristic(distances)
		heuristicWriter.Write(hOutput)

		_, gOutput := compute(solveGreedy, distances, bestKnown, "Greedy")
		greedyWriter.Write(gOutput)

		steepestElapsed, sOutput := compute(solveSteepest, distances, bestKnown, "Steepest")
		steepestWriter.Write(sOutput)

		rOutput := computeRandom(distances, steepestElapsed)
		randomWriter.Write(rOutput)
	}
}

func compute(solve func([][]int) ([]int, int), distances [][]int, bestKnown int, name string) (time.Duration, []string) {
	qualities := makeArray(RUN_COUNT)
	stepCounts := makeArray(RUN_COUNT)

	start := time.Now()
	for i := 0; i < RUN_COUNT || time.Since(start).Nanoseconds() < MIN_TIME; i++ {
		permutation, stepCount := solve(distances)
		if i < RUN_COUNT {
			stepCounts[i] = stepCount
			distance := getDistance(permutation, distances)
			qualities[i] = distance - bestKnown
		}
	}
	elapsed := time.Since(start)
	bestResult := min(qualities)
	meanResult := mean(qualities)

	stdResult := std(qualities)
	meanSteps := mean(stepCounts)

	fmt.Println(name, "elapsed: ", elapsed, "best: ", bestResult, "mean: ", meanResult, "steps(mean): ", meanSteps, "std result: ", stdResult)
	output := []string{
		itoa(bestResult),
		ftoa(meanResult),
		ftoa(meanSteps),
		ftoa(stdResult),
		itoa(int(elapsed.Nanoseconds())),
	}
	return elapsed, output
}

func computeHeuristic(distances [][]int) []string{
	permutation := solveHeuristic(distances)
	result := getDistance(permutation, distances)
	fmt.Println("Heuristic result: ", result)
	return []string{itoa(result)}
}

func computeRandom(distances [][]int, availableTime time.Duration) []string{
	start := time.Now()
	permutation := solveRandom(distances, availableTime)
	result := getDistance(permutation, distances)
	elapsed := time.Since(start)

	fmt.Println("Random, elapsed: ", elapsed, "result: ", result)
	output := []string{
		itoa(result),
		itoa(int(elapsed.Nanoseconds())),
	}
	return output
}

func getWriter(filename string) (*os.File, *csv.Writer){
	file, _ := os.Create(filename)
	writer := csv.NewWriter(file)
	return file, writer
}
