package main

import (
	"fmt"
	"log"
)

type board [10]solution

type solution [5]square

type square struct {
	num    int
	called bool
}

func main() {
	bds := createsolutionList(boards)
	winner := board{}
	lastNum := 0
	for _, b := range balls {
		for i := 0; i < len(bds); i++ {
			for j := 0; j < 10; j++ {
				bds[i][j] = updateSolution(bds[i][j], b)
				if chkSolution(bds[i][j]) {
					winner = bds[i]
					lastNum = b
					break
				}
			}
		}
		if (winner != board{}) {
			break
		}
	}
	filtered := fillter(winner)
	total := sumUncalledSquares(filtered)
	log.Printf("The winning row is %v", winner)
	log.Printf("The answer to part A is %v \n the winning row %v * the winning number %v", total*lastNum, total, lastNum)
}

func (b board) String() string {
	str := "\n"
	for i := 0; i < len(b); i++ {
		str += fmt.Sprintf("[%v", b[i][0].num)
		for j := 1; j < 5; j++ {
			str += fmt.Sprintf(",%v", b[i][j].num)
		}
		str += "]\n"
	}
	return str
}

func fillter(b board) board {
	for k := 0; k < 10; k++ {
		for l := 0; l < 5; l++ {
			for i := 0; i < 10; i++ {
				for j := 0; j < 5; j++ {
					if b[k][l].num == b[i][j].num && b[k][l].called != b[i][j].called {
						b[k][l].called, b[k][l].called = true, true
					}
				}
			}
		}
	}
	return b
}

func sumUncalledSquares(b board) int {
	sum := 0
	counted := make(map[int]bool)
	for _, s := range b {
		for j := 0; j < 5; j++ {
			if !s[j].called && !counted[s[j].num] {
				sum += s[j].num
				counted[s[j].num] = true
			}
		}
	}
	return sum
}

func sumSolu(s solution) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i].num
	}
	return sum
}

func chkSolution(s solution) bool {
	return s[0].called && s[1].called && s[2].called && s[3].called && s[4].called
}

func updateSolution(s solution, n int) solution {
	for i := 0; i < 5; i++ {
		if s[i].num == n {
			s[i].called = true
		}
	}
	return s
}

// func checkForWinners(boards []Board) (bool []int) {
// 	for _, board := range boards {
// 		for _, sol, := range board.solutions {

// 		}
// 	}
// }

func createsolutionList(boards [][5][5]int) []board {
	allBrds := []board{}
	// redo here

	for _, b := range boards { // run though each board
		solbrd := board{}
		for r, row := range b {
			solbrd[r] = mapRowToSquare(row[:]...)
			solbrd[r+5] = mapRowToSquare(b[0][r], b[1][r], b[2][r], b[3][r], b[4][r])
		}
		allBrds = append(allBrds, solbrd)
	}
	return allBrds
}

func mapRowToSquare(nums ...int) [5]square {
	var row = [5]square{}
	for i, n := range nums {
		var s = square{n, false}
		row[i] = s
	}
	return row
}

// func setUpBoards(intbrds [][5][5]int, key map[int]square) []Board {
// 	var boards = []Board{}
// 	for _, b := range intbrds { // run though each board
// 		var solution = [10][5]*square{}
// 		for r, row := range b {
// 			solution[r] = mapRowToSquare(key, row[:]...)
// 			solution[5+r] = mapRowToSquare(key, b[0][r], b[1][r], b[2][r], b[3][r], b[4][r])
// 		}
// 		boards = append(boards, Board{solution})
// 	}
// 	return boards
// }

/*
79 67 21 84 71
25 22 19 80 13
10 63 90 78 33
93 50 89 58 87
91  7 45  6 41
*/
