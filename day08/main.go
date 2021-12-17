package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SolveI() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	segments := make(map[int]int)
	segments[0] = 6
	segments[1] = 2
	segments[2] = 5
	segments[3] = 5
	segments[4] = 4
	segments[5] = 5
	segments[6] = 6
	segments[7] = 3
	segments[8] = 7
	segments[9] = 6

	total := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		// patterns := strings.Fields(line[0])
		output := strings.Fields(line[1])
		for _, d := range output {
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				total++
			}
		}
	}
	fmt.Println(total)
}

func main() {
	SolveI()
}
