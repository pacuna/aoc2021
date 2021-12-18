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

func SolveII() {
	lookUpTable := make(map[byte][]int)
	lookUpTable['a'] = []int{0, 2, 3, 5, 6, 7, 8, 9}
	lookUpTable['b'] = []int{0, 1, 2, 3, 4, 7, 8, 9}
	lookUpTable['c'] = []int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	lookUpTable['d'] = []int{0, 2, 3, 5, 6, 8, 9}
	lookUpTable['e'] = []int{0, 2, 6, 8}
	lookUpTable['f'] = []int{0, 4, 5, 6, 8, 9}
	lookUpTable['g'] = []int{2, 3, 4, 5, 6, 8, 9}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	// acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		patterns := strings.Fields(line[0])
		output := strings.Fields(line[1])
		m := make(map[byte][]byte)
		for _, p := range patterns {
			if len(p) == 2 {
				fmt.Println(p)
				m[p[0]] = []byte{'b', 'c'}
				m[p[1]] = []byte{'b', 'c'}
				// ab => a could be b|c, b could be b|c
				// dab => d has to be a, a could be b|c, b could b|c
				// eafd => e could be f|g, a could be b|c, f could be f|g, d has to be a
				// acedgfb => a could be b|c, c cannnot be (b|c|a|f|g) h can be d|e
			}else if len(p) == 3 {
				m[p[0]] = []byte{'a', 'b', 'c'}
				m[p[1]] = []byte{'a', 'b', 'c'}
				m[p[2]] = []byte{'a', 'b', 'c'}
			}
		}
		fmt.Println(patterns, output)
	}

}

func main() {
	SolveII()
}
