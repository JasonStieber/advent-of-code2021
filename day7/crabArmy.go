package main

import "log"

func main() {
	steps := minDistance(input)
	sumStep := minSumDistance(input)
	log.Printf("The answer to Part A is: %v", steps)
	log.Printf("The answer to part B is: %v", sumStep)
}

func minSumDistance(set []int) int {
	fewestSteps := 0
	for i := 0; i < len(set); i++ {
		steps := totalSumSteps(i, set)
		if fewestSteps == 0 || steps < fewestSteps {
			fewestSteps = steps
		} else {
			return fewestSteps
		}

	}
	return -1
}

func minDistance(set []int) int {
	fewestSteps := 0
	for i := 0; i < len(set); i++ {
		steps := totalSteps(i, set)
		if fewestSteps == 0 || steps < fewestSteps {
			fewestSteps = steps
		} else {
			return fewestSteps
		}

	}
	return -1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func totalSumSteps(target int, set []int) int {
	total := 0
	for i := 0; i < len(set); i++ {
		total += sumToN(abs(set[i] - target))
	}
	return total
}

func totalSteps(target int, set []int) int {
	total := 0
	for i := 0; i < len(set); i++ {
		total += abs(set[i] - target)
	}
	return total
}

func sumToN(n int) int {
	if n < 0 {
		log.Panicf("Expected a non negtive number got got %v", n)
	}
	return (n * (n + 1)) / 2
}
