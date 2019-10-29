package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type fn func([][]int) ([]int, int)

const MIN_TIME = 1000 * 1000
const RUN_COUNT = 10

func main() {
	bestKnownSolutions := map[string]int{
		"br17":   39,
		"ft53":   6905,
		"ftv33":  1286,
		"ft70":   38673,
		"ftv35":  1473,
		"ftv44":  1613,
		"ftv38":  1530,
		"ftv55":  1608,
		"ftv47":  1776,
		"ftv64":  1839,
		"ftv70":  1950,
		"ftv100": 1788,
		"ftv110": 1958,
		"ftv90":  1579,
		"ftv120": 2166,
		"ftv130": 2307,
		"ftv140": 2420,
		"ftv150": 2611,
		"ftv170": 2755,
		"ftv160": 2683,
		"kro124": 36230,
		"p43":    5620,
		"rbg358": 1163,
		"rbg323": 1326,
		"rbg403": 2465,
		"rbg443": 2720,
		"ry48p":  14422,
	}
	instanceFilenames := []string{
		"br17",
		"ftv33",
		"ftv35",
		"ftv38",
		"ftv44",
		"ftv47",
		"ft53",
		"ftv55",
		"ft70",
	}

	gFile, greedyWriter := getWriter("../results/greedy.csv")
	sFile, steepestWriter := getWriter("../results/steepest.csv")
	hFile, heuristicWriter := getWriter("../results/heuristic.csv")
	rFile, randomWriter := getWriter("../results/random.csv")

	defer gFile.Close()
	defer sFile.Close()
	defer hFile.Close()
	defer rFile.Close()
	defer greedyWriter.Flush()
	defer steepestWriter.Flush()
	defer heuristicWriter.Flush()
	defer randomWriter.Flush()

	greedyWriter.Write([]string{"size", "best", "mean", "mean_steps", "std", "time"})
	steepestWriter.Write([]string{"size", "best", "mean", "mean_steps", "std", "time"})
	heuristicWriter.Write([]string{"size", "best"})
	randomWriter.Write([]string{"size", "best", "time"})

	for _, filename := range instanceFilenames {
		fmt.Println()
		fmt.Println(filename)
		path := "../data/" + filename + ".atsp"
		distances := readData(path)

		bestKnown := bestKnownSolutions[filename]
		fmt.Println("Best known: ", bestKnown)

		hOutput := computeHeuristic(distances, bestKnown)
		heuristicWriter.Write(hOutput)

		_, gOutput := computeGS(solveSwapGreedy, distances, bestKnown, "Greedy")
		greedyWriter.Write(gOutput)

		steepestElapsed, sOutput := computeGS(solveReverseSteepest, distances, bestKnown, "Steepest")
		steepestWriter.Write(sOutput)

		rOutput := computeRandom(distances, steepestElapsed, bestKnown)
		randomWriter.Write(rOutput)
	}
}

func computeGS(solve func([][]int) ([]int, int), distances [][]int, bestKnown int, name string) (time.Duration, []string) {
	qualities := make([]float64, RUN_COUNT)
	stepCounts := makeArray(RUN_COUNT)

	start := time.Now()
	for i := 0; i < RUN_COUNT || time.Since(start).Nanoseconds() < MIN_TIME; i++ {
		permutation, stepCount := solve(distances)
		if i < RUN_COUNT {
			stepCounts[i] = stepCount
			result := getDistance(permutation, distances)
			qualities[i] = getQuality(result, bestKnown)
		}
	}
	elapsed := time.Since(start)
	bestResult := minOfArray(qualities)
	fmt.Println(bestResult, qualities)

	meanResult := mean(qualities)

	stdResult := std(qualities)
	meanSteps := meanInt(stepCounts)

	fmt.Println(name, "elapsed: ", elapsed, "best: ", bestResult, "mean: ", meanResult, "steps(mean): ", meanSteps, "std result: ", stdResult)
	output := []string{
		itoa(len(distances)),
		ftoa(bestResult),
		ftoa(meanResult),
		ftoa(meanSteps),
		ftoa(stdResult),
		itoa(int(elapsed.Nanoseconds())),
	}
	return elapsed, output
}

func computeHeuristic(distances [][]int, bestKnown int) []string {
	permutation := solveHeuristic(distances)
	result := getDistance(permutation, distances)
	quality := getQuality(result, bestKnown)
	fmt.Println("Heuristic result: ", quality)
	return []string{
		itoa(len(distances)),
		ftoa(quality),
	}
}

func computeRandom(distances [][]int, availableTime time.Duration, bestKnown int) []string {
	start := time.Now()
	permutation := solveRandom(distances, availableTime)
	result := getDistance(permutation, distances)
	quality := getQuality(result, bestKnown)
	elapsed := time.Since(start)

	fmt.Println("Random, elapsed: ", elapsed, "result: ", quality)
	output := []string{
		itoa(len(distances)),
		ftoa(quality),
		itoa(int(elapsed.Nanoseconds())),
	}
	return output
}

func getWriter(filename string) (*os.File, *csv.Writer) {
	file, _ := os.Create(filename)
	writer := csv.NewWriter(file)
	return file, writer
}
