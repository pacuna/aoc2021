package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func SolveI() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var (
		x1 int
		x2 int
		y1 int
		y2 int
	)

	covered := make(map[Point]int)

	for scanner.Scan() {
		line := scanner.Text()
		r := strings.NewReader(line)
		fmt.Fscanf(r, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		p1 := Point{x: x1, y: y1}
		p2 := Point{x: x2, y: y2}

		var start, end int
		if p1.x == p2.x {
			if p1.y >= p2.y {
				start = p2.y
				end = p1.y
			} else {
				start = p1.y
				end = p2.y
			}
			for i := start; i <= end; i++ {
				covered[Point{x: p1.x, y: i}]++
			}
		} else if p1.y == p2.y {
			if p1.x >= p2.x {
				start = p2.x
				end = p1.x
			} else {
				start = p1.x
				end = p2.x
			}
			for i := start; i <= end; i++ {
				covered[Point{x: i, y: p1.y}]++
			}
		}
	}

	res := 0
	for _, v := range covered {
		if v >= 2 {
			res++
		}
	}
	fmt.Println(res)
}

func main() {
	SolveI()
}
