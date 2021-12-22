package main

import (
	"fmt"
	"log"
)

func main() {
	grid := make(map[cord]int)
	moreThan1 := 0
	for i := 0; i < len(lines); i++ {
		grid = travelLine(lines[i], grid)
	}
	for _, t := range grid {
		if t > 1 {
			moreThan1++
		}
	}

	log.Printf("The answer to part B is: %v", moreThan1)

}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func travelLine(l line, m map[cord]int) map[cord]int {
	dx, dy, err := travelDir(l)
	if err != nil {
		log.Printf(error.Error(err))
		return m
	}
	if dx != 0 && dy != 0 {
		for x, y := l.start.x, l.start.y; x != l.end.x+dx; {
			var c = cord{x, y}
			m[c]++
			x, y = x+dx, y+dy
		}
		return m
	} else if dx != 0 {
		for x := l.start.x; x != l.end.x+dx; x += dx {
			var c = cord{x, l.end.y}
			m[c]++
		}
		return m
	} else if dy != 0 {
		for y := l.start.y; y != l.end.y+dy; y += dy {
			var c = cord{l.end.x, y}
			m[c]++
		}
		return m
	}
	panic("We should have moved here")
	return m
}

func travelDir(l line) (int, int, error) {
	if abs(l.start.x-l.end.x) == abs(l.start.y-l.end.y) {
		if l.start.x > l.end.x && l.start.y > l.end.y {
			return -1, -1, nil
		} else if l.start.x > l.end.x && l.start.y < l.end.y {
			return -1, 1, nil
		} else if l.start.x < l.end.x && l.start.y > l.end.y {
			return 1, -1, nil
		} else if l.start.x < l.end.x && l.start.y < l.end.y {
			return 1, 1, nil
		}
	} else if l.start.x == l.end.x && l.start.y < l.end.y {
		return 0, 1, nil
	} else if l.start.x == l.end.x && l.start.y > l.end.y {
		return 0, -1, nil
	} else if l.start.y == l.end.y && l.start.x < l.end.x {
		return 1, 0, nil
	} else if l.start.y == l.end.y && l.start.x > l.end.x {
		return -1, 0, nil
	}
	return 0, 0, fmt.Errorf("Not in a strait line")
}
