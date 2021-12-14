package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckRows(check [][]int) bool {
	var res bool
	for row := range check {
		res = true
		for col := range check[row] {
			if check[row][col] == 0 {
				res = false
				break
			}
		}
		if res {
			return true
		}
	}
	return res
}

func CheckCols(check [][]int) bool {
	var res bool
	for row := range check {
		res = true
		for col := range check[row] {
			if check[col][row] == 0 {
				res = false
				break
			}
		}
		if res {
			return true
		}
	}
	return res
}

func main() {
	draw := []int{}

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	// read draw
	line := strings.Split(scanner.Text(), ",")
	for _, c := range line {
		d, _ := strconv.Atoi(c)
		draw = append(draw, d)
	}

	// skip first blank line
	scanner.Scan()

	// read boards

	boards := [][][]int{}

	var tmpRow []int
	tmpBoard := [][]int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, tmpBoard)
			tmpBoard = [][]int{}
			continue
		}
		line := strings.Fields(scanner.Text())
		tmpRow = []int{}
		for _, c := range line {
			d, _ := strconv.Atoi(c)
			tmpRow = append(tmpRow, d)
		}
		tmpBoard = append(tmpBoard, tmpRow)

	}
	boards = append(boards, tmpBoard)

	checks := make([][][]int, len(boards))
	for i := range checks {
		checks[i] = make([][]int, len(boards[0]))
		for j := range checks[i] {
			checks[i][j] = make([]int, len(boards[0][0]))
		}
	}

	// CheckLastWinner(draw, boards, checks)
	w, d := CheckLastWinner(draw, boards, checks)
	fmt.Println(GetResult(boards[w], checks[w], d))
}

func GetResult(winner [][]int, winnerChecks [][]int, draw int) int {
	var sumUnmarked int
	for i := range winnerChecks {
		for j := range winnerChecks[i] {
			if winnerChecks[i][j] == 0 {
				sumUnmarked += winner[i][j]
			}
		}
	}
	return sumUnmarked * draw
}

func CheckWinner(draw []int, boards [][][]int, checks [][][]int) (int, int) {
	for _, el := range draw {
		for idx, board := range boards {
			for ri, row := range board {
				for rc, col := range row {
					if col == el {
						checks[idx][ri][rc] = 1
						if CheckRows(checks[idx]) || CheckCols(checks[idx]) {
							return idx, el
						}
					}
				}
			}
		}
	}
	return -1, -1
}

func CheckLastWinner(draw []int, boards [][][]int, checks [][][]int) (int, int) {
	checked := make(map[int]bool)
	for _, el := range draw {
		for idx, board := range boards {
			for ri, row := range board {
				for rc, col := range row {
					if col == el {
						checks[idx][ri][rc] = 1
						if CheckRows(checks[idx]) || CheckCols(checks[idx]) {
							if !checked[idx] {
								checked[idx] = true
								if len(checked) == len(boards) {
									return idx, el
								}
							}
						}
					}
				}
			}
		}
	}
	return -1, -1
}
