package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func findMax(n []int) int {
	max := math.MinInt
	for _, el := range n {
		if el > max {
			max = el
		}
	}
	return max
}

func abs(x, y int) int {
	if x-y < 0 {
		return -1 * (x - y)
	}
	return x - y
}

func sumTo(x int) int {
	res := 0
	for i := 1; i <= x; i++ {
		res += i
	}
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	nums := []int{}
	for _, c := range line {
		d, _ := strconv.Atoi(c)
		nums = append(nums, d)
	}
	max := findMax(nums)

	distances := make(map[int]int)
	for i := 0; i <= max; i++ {
		for j := range nums {
			// part I
			distances[i] += sumTo(abs(nums[j], i))

			// part II
			// distances[i] += sumTo(abs(nums[j], i))
		}
	}

	min := math.MaxInt
	for _, v := range distances {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
}
