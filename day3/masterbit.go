package main

import (
	"log"
	"strings"
)

func main() {
	var btTotal [12]int
	bits := parseIntoSlice(input)
	for _, b := range bits {
		for i := 0; i < 12; i++ {
			if string(b[i]) == "1" {
				btTotal[i]++
			} else {
				btTotal[i]--
			}
		}
	}
	gamma := 0
	epsilon := 0
	for _, n := range btTotal {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if n > 0 {
			gamma = gamma | 1
			log.Printf("We are here")
		} else {
			epsilon = epsilon | 1
		}

	}
	log.Printf("Our bit is %v ", btTotal)
	log.Printf("in binary \n%12b \n%12b", gamma, epsilon)
	log.Printf("the answer to part A: is %v", gamma*epsilon)
}

func parseIntoSlice(s string) []string {
	arr := strings.Split(s, "\n")
	return arr
}
