package main

import "log"

func main() {
	steps := testEachSlot(input)
	log.Printf("The answer to example is: %v", steps)
}

func testEachSlot(set []int) int {
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

func totalSteps(target int, set []int) int {
	total := 0
	for i := 0; i < len(set); i++ {
		total += abs(set[i] - target)
	}
	return total
}
