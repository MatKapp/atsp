package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type fn func([][]int) ([]int, int)

const MIN_TIME = 1000 * 1000
const DEFAULT_RUN_COUNT = 10
const STEP_RUN_COUNT = 400
const GS_COUNT_STEPS = 10
const STEP = STEP_RUN_COUNT / GS_COUNT_STEPS

func main() {
	bestKnownSolutions := initializaBestKnownSolutions()
	instanceFilenames := initializeFileNames()

	stepProcessingInstanceFilename := "kro124p"

	swapGreedyFile, swapGreedyWriter := getWriter("../results/swapGreedy.csv")
	reverseGreedyFile, reverseGreedyWriter := getWriter("../results/reverseGreedy.csv")
	swapSteepestFile, swapSteepestWriter := getWriter("../results/swapSteepest.csv")
	reverseSteepestFile, reverseSteepestWriter := getWriter("../results/reverseSteepest.csv")
	hFile, heuristicWriter := getWriter("../results/heuristic.csv")
	rFile, randomWriter := getWriter("../results/random.csv")
	stepMeanProcessingFile, stepMeanProcessingWriter := getWriter("../results/stepMeanProcessing-" + stepProcessingInstanceFilename + ".csv")
	stepBestProcessingFile, stepBestProcessingWriter := getWriter("../results/stepBestProcessing-" + stepProcessingInstanceFilename + ".csv")
	stepSimilarityFile, stepSimilarityWriter := getWriter("../results/similarity-" + stepProcessingInstanceFilename + ".csv")

	defer swapGreedyFile.Close()
	defer reverseGreedyFile.Close()
	defer swapSteepestFile.Close()
	defer reverseSteepestFile.Close()
	defer hFile.Close()
	defer rFile.Close()
	defer stepMeanProcessingFile.Close()
	defer stepBestProcessingFile.Close()
	defer stepSimilarityFile.Close()
	defer swapGreedyWriter.Flush()
	defer reverseGreedyWriter.Flush()
	defer swapSteepestWriter.Flush()
	defer reverseSteepestWriter.Flush()
	defer heuristicWriter.Flush()
	defer randomWriter.Flush()
	defer stepMeanProcessingWriter.Flush()
	defer stepBestProcessingWriter.Flush()
	defer stepSimilarityWriter.Flush()

	swapGreedyWriter.Write([]string{"size", "best", "mean", "mean_steps", "std", "time", "reviewed_solutions", "quality_time"})
	reverseGreedyWriter.Write([]string{"size", "best", "mean", "mean_steps", "std", "time", "reviewed_solutions", "quality_time"})
	swapSteepestWriter.Write([]string{"size", "best", "mean", "mean_steps", "std", "time", "reviewed_solutions", "quality_time"})
	reverseSteepestWriter.Write([]string{"size", "best", "mean", "mean_steps", "std", "time", "reviewed_solutions", "quality_time"})
	heuristicWriter.Write([]string{"size", "best"})
	randomWriter.Write([]string{"size", "best", "time"})
	stepMeanProcessingWriter.Write([]string{"step", "iteration_num", "quality"})
	stepBestProcessingWriter.Write([]string{"step", "iteration_num", "quality"})
	stepSimilarityWriter.Write([]string{"step", "quality", "similarity"})

	for _, filename := range instanceFilenames {
		fmt.Println()
		fmt.Println(filename)
		path := "../data/" + filename + ".atsp"
		distances := readData(path)

		stepProcessing := false

		if filename == stepProcessingInstanceFilename {
			stepProcessing = true
		}

		bestKnown := bestKnownSolutions[filename]
		fmt.Println("Best known: ", bestKnown)

		hOutput := computeHeuristic(distances, bestKnown)
		heuristicWriter.Write(hOutput)

		_, swapGreedyOutput, meanResult, bestResult, stepPermutations, qualities := computeGS(solveOptimizedSwapGreedy, distances, bestKnown, "SwapGreedy", stepProcessing)

		if stepProcessing {
			for index, element := range meanResult {
				stepMeanProcessingWriter.Write([]string{itoa(index), itoa(STEP * (index + 1)), ftoa(element)})
			}

			for index, element := range bestResult {
				stepBestProcessingWriter.Write([]string{itoa(index), itoa(STEP * (index + 1)), ftoa(element)})
			}

			_, indexOfMin := minFloatOfArray(qualities)
			bestPermutation := stepPermutations[indexOfMin]

			for index, element := range stepPermutations {
				stepSimilarityWriter.Write([]string{itoa(index), ftoa(qualities[index]), ftoa(countSimilarity(element, bestPermutation))})
			}
		}

		swapGreedyWriter.Write(swapGreedyOutput)
		_, reverseGreedyOutput, _, _, _, _ := computeGS(solveReverseGreedy, distances, bestKnown, "ReverseGreedy", false)
		reverseGreedyWriter.Write(reverseGreedyOutput)

		swapSteepestElapsed, swapSteepestOutput, _, _, _, _ := computeGS(solveOptimizedSwapSteepest, distances, bestKnown, "SwapSteepest", false)
		swapSteepestWriter.Write(swapSteepestOutput)
		_, reverseSteepestOutput, _, _, _, _ := computeGS(solveReverseSteepest, distances, bestKnown, "ReverseSteepest", false)
		reverseSteepestWriter.Write(reverseSteepestOutput)

		rOutput := computeRandom(distances, swapSteepestElapsed, bestKnown)
		randomWriter.Write(rOutput)
	}
}

func computeGS(solve func([][]int, bool) ([]int, int, int, [][]int), distances [][]int, bestKnown int, name string, stepProcessing bool) (time.Duration, []string, []float64, []float64, [][]int, []float64) {
	runCount := DEFAULT_RUN_COUNT

	if stepProcessing {
		runCount = STEP_RUN_COUNT
	}

	qualities := make([]float64, runCount)
	stepCounts := makeArray(runCount)
	reviewedSolutionsNumbers := makeArray(runCount)
	start := time.Now()
	meanResultsAfterStep := make([]float64, GS_COUNT_STEPS-1)
	bestResultsAfterStep := make([]float64, GS_COUNT_STEPS-1)
	stepPermutationsResult := make([][]int, runCount)

	if stepProcessing {

		for i := 0; i < runCount || time.Since(start).Nanoseconds() < MIN_TIME; i++ {
			permutation, stepCount, reviewedSolutions, _ := solve(distances, stepProcessing)
			if i < runCount {
				if i%STEP == 0 {
					if i > 0 {
						meanResultsAfterStep[(i/STEP)-1] = mean(qualities[0:i])
						bestResultsAfterStep[(i/STEP)-1] = minOfArray(qualities[0:i])
					}
				}
				stepPermutationsResult[i] = permutation
				stepCounts[i] = stepCount
				reviewedSolutionsNumbers[i] = reviewedSolutions
				result := getDistance(permutation, distances)
				qualities[i] = getQuality(result, bestKnown)
			}
		}
	} else {
		for i := 0; i < runCount || time.Since(start).Nanoseconds() < MIN_TIME; i++ {
			permutation, stepCount, reviewedSolutions, _ := solve(distances, stepProcessing)

			if i < runCount {
				stepCounts[i] = stepCount
				reviewedSolutionsNumbers[i] = reviewedSolutions
				result := getDistance(permutation, distances)
				qualities[i] = getQuality(result, bestKnown)
			}
		}
	}

	elapsed := time.Since(start)
	bestResult := minOfArray(qualities)
	fmt.Println(bestResult, qualities)

	meanResult := mean(qualities)

	stdResult := std(qualities)
	meanSteps := meanInt(stepCounts)
	meanReviewedSolutions := meanInt(reviewedSolutionsNumbers)

	fmt.Println(name, "elapsed: ", elapsed, "best: ", bestResult, "mean: ", meanResult, "steps(mean): ", meanSteps, "std result: ", stdResult, "reviewed solutions(mean): ", meanReviewedSolutions)
	output := []string{
		itoa(len(distances)),
		ftoa(bestResult),
		ftoa(meanResult),
		ftoa(meanSteps),
		ftoa(stdResult),
		itoa(int(elapsed.Nanoseconds())),
		ftoa(meanReviewedSolutions),
		ftoa(meanResult / float64(elapsed.Milliseconds())),
	}
	return elapsed, output, meanResultsAfterStep, bestResultsAfterStep, stepPermutationsResult, qualities
}

func initializaBestKnownSolutions() map[string]int {
	return map[string]int{
		"br17":    39,
		"ft53":    6905,
		"ftv33":   1286,
		"ft70":    38673,
		"ftv35":   1473,
		"ftv44":   1613,
		"ftv38":   1530,
		"ftv55":   1608,
		"ftv47":   1776,
		"ftv64":   1839,
		"ftv70":   1950,
		"ftv100":  1788,
		"ftv110":  1958,
		"ftv90":   1579,
		"ftv120":  2166,
		"ftv130":  2307,
		"ftv140":  2420,
		"ftv150":  2611,
		"ftv170":  2755,
		"ftv160":  2683,
		"kro124p": 36230,
		"p43":     5620,
		"rbg358":  1163,
		"rbg323":  1326,
		"rbg403":  2465,
		"rbg443":  2720,
		"ry48p":   14422,
	}
}

func initializeFileNames() []string {
	return []string{
		"br17",
		"ftv33",
		"ftv35",
		"ftv38",
		"ftv44",
		"ftv47",
		"ft53",
		"ftv55",
		"ft70",
		"kro124p",
	}
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
