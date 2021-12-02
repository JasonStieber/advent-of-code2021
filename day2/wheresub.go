package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	cord := parseInput(input)
	depth, dist := cord["down"]-cord["up"], cord["forward"]
	log.Printf("The answer to part A is %v, with a depth of: %v and a distance of %v", depth*dist, depth, dist)

}

func parseInput(s string) map[string]int {
	m := make(map[string]int)
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		verbNum := strings.Split(l, " ")
		num, err := strconv.Atoi(verbNum[1])
		if err != nil {
			fmt.Errorf("%v was expected to be anumber %v ", num, err)
		}
		m[verbNum[0]] += num
	}
	return m
}
