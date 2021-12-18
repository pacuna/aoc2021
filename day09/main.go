package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SolveI() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	M := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			d, _ := strconv.Atoi(string(c))
			row = append(row, d)
		}
		M = append(M, row)
	}

	lowPoints := []int{}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if isLowPoint(i, j, M) {
				lowPoints = append(lowPoints, M[i][j])
			}
		}
	}
	fmt.Println(riskLevel(lowPoints))

}

func riskLevel(lowPoints []int) int {
	rl := 0
	for _, lp := range lowPoints {
		rl += lp + 1
	}
	return rl
}

func isLowPoint(i, j int, M [][]int) bool {
	var (
		prevX int
		nextX int
		prevY int
		nextY int
	)
	if i-1 < 0 {
		prevX = M[i][j] + 1
	} else {
		prevX = M[i-1][j]
	}

	if i+1 >= len(M) {
		nextX = M[i][j] + 1
	} else {
		nextX = M[i+1][j]
	}

	if j-1 < 0 {
		prevY = M[i][j] + 1
	} else {
		prevY = M[i][j-1]
	}

	if j+1 >= len(M[i]) {
		nextY = M[i][j] + 1
	} else {
		nextY = M[i][j+1]
	}
	point := M[i][j]
	// fmt.Println("testing ", point)
	// fmt.Println("comparing with ", prevX, nextX, prevY, nextY)
	// fmt.Println(prevX > point && nextX > point && prevY > point && nextY > point)
	return prevX > point && nextX > point && prevY > point && nextY > point
}

func main() {
	SolveI()
}
