package main

import (
	"log"
	"strings"
)

type scramClock struct {
	obsrv [10]string
	dsply [4]string
}

func main() {
	parced := parseInput(input)
	totalUnique := cntUnique(parced)
	log.Printf("The answer to part A is %v", totalUnique)
}

func cntUnique(d []scramClock) int {
	cnt := 0
	for _, inp := range d {
		for i := 0; i < 4; i++ {
			l := len(inp.dsply[i])
			if l == 2 || l == 3 || l == 4 || l == 7 {
				cnt++
			}
		}
	}
	return cnt
}

func parseInput(s string) []scramClock {
	output := []scramClock{}
	lines := strings.Split(s, "\n")
	for _, data := range lines {
		d := scramClock{}
		twoPrts := strings.Split(data, "|")
		obsrv := strings.Fields(twoPrts[0])
		dpsly := strings.Fields(twoPrts[1])
		copy(d.dsply[:], dpsly)
		copy(d.obsrv[:], obsrv)
		output = append(output, d)
	}
	return output
}

//..strings.Fields
