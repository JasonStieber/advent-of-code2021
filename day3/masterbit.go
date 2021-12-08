package main

import (
	"log"
	"strconv"
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
		} else {
			epsilon = epsilon | 1
		}

	}
	var o2Key, c02Key [12]int
	for j := 0; j < 12; j++ {
		if btTotal[j] > 0 {
			o2Key[j] = 1
		} else {
			c02Key[j] = 1
		}
	}
	binaries := convertToInt(bits)

	for k := 0; k < 12; k++ {
		n := countKthBit(binaries, k)
		log.Printf("The remaining nums is %v", len(binaries))
		if n > 0 || n == 0 {
			binaries = filtered(binaries, k, true)
		}

		if len(binaries) == 1 {
			log.Printf("the o2 number is: %12b", binaries[0])
		}
	}
	// co2 := filterco2(binaries, 0)
	//log.Printf("Our bit is %v ", btTotal)
	log.Printf("in binary \n%12b \n%12b", gamma, epsilon)
	log.Printf("the answer to part A: is %v", gamma*epsilon)
	// log.Printf("The answer to part B: is %v\n o2 and c02 numbers are \n%12b \n%12b  ", o2*co2, o2, co2)
}

func convertToInt(s []string) []int {
	var b []int
	for _, n := range s {
		i, err := strconv.ParseInt(n, 2, 64)
		if err != nil {
			log.Fatalf("cannot process string %v into a number error: %v", n, err)
		}
		b = append(b, int(i))
	}
	return b
}

func filtered(bits []int, n int, on bool) []int {
	var filt = []int{}
	for _, b := range bits {
		if on && b&(1<<(11-n)) != 0 {
			filt = append(filt, b)
		} else if !on && b|^(1<<(11-n)) != 0b111111111111 {
			filt = append(filt, b)
		}
	}
	return filt
}

func countKthBit(nums []int, loc int) int {
	total := 0
	for _, n := range nums {
		if n&(1<<(11-loc)) != 0 {
			total++
		} else {
			total--
		}
	}
	return total
}

func parseIntoSlice(s string) []string {
	arr := strings.Split(s, "\n")
	return arr
}
