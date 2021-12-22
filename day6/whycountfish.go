package main

import "log"

func main() {
	time := 80
	b := placeIntoArray(fish)
	nb := passTime(time, b)
	tnb := passTime(256, b)
	total := sumBucket(nb)
	total2 := sumBucket(tnb)
	log.Printf("The answer to Part A is is:%v fish after %v days", total, time)
	log.Panicf("The answer to Part B is is:%v fish after %v days", total2, 256)
}

func sumBucket(b [9]int) int {

	t := 0
	for i := 0; i < 9; i++ {
		t += b[i]
	}
	return t
}

func placeIntoArray(f []int) [9]int {
	bucket := [9]int{}
	for i := 0; i < len(f); i++ {
		bucket[f[i]]++
	}
	return bucket
}

func passTime(t int, b [9]int) [9]int {
	monringB := b
	for i := 0; i < t; i++ {
		nBucket := [9]int{}
		for j := 7; j >= 0; j-- {
			nBucket[j] = monringB[j+1]
		}
		nBucket[8] = monringB[0]
		nBucket[6] += monringB[0]
		monringB = nBucket
	}
	return monringB
}
