package main

import "log"

func main() {
	c := 0
	sc := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			c++
		}
	}
	for j := 3; j < len(input); j++ {
		if input[j] > input[j-3] {
			sc++
		}
	}
	log.Printf("The answer to part A: is %v", c)
	log.Printf("The answer to part B: is %v", sc)
}
