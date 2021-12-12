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

	var (
		dir      string
		mag      int
		hor, dep int
	)
	for scanner.Scan() {
		line := scanner.Text()
		r := strings.NewReader(line)
		fmt.Fscanf(r, "%s %d", &dir, &mag)
		switch dir {
		case "forward":
			hor += mag
		case "up":
			dep -= mag
		case "down":
			dep += mag
		}
	}
	fmt.Println(hor * dep)
}

func SolveII() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var (
		dir      string
		mag      int
		hor, dep int
		aim      int
	)
	for scanner.Scan() {
		line := scanner.Text()
		r := strings.NewReader(line)
		fmt.Fscanf(r, "%s %d", &dir, &mag)
		switch dir {
		case "forward":
			hor += mag
			dep += aim * mag
		case "up":
			aim -= mag
		case "down":
			aim += mag
		}
	}
	fmt.Println(hor * dep)
}

func main() {
	SolveI()
	SolveII()
}
