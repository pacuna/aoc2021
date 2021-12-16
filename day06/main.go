package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	pop := []int{}
	for _, c := range line {
		d, _ := strconv.Atoi(c)
		pop = append(pop, d)
	}

	hsh := make(map[int]int)

	for _, el := range pop {
		hsh[el]++
	}

	for i := 0; i < 256; i++ {
		tmp := make(map[int]int)
		for k, v := range hsh {
			if k == 0 {
				tmp[6] += v
				tmp[8] += v
			} else {
				tmp[k-1] += v
			}
		}
		hsh = tmp
	}

	res := 0
	for _, v := range hsh {
		res += v
	}
	fmt.Println(res)
}
