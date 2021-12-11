package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func PartI() {
	result := 0
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	cur := math.MaxInt
	for scanner.Scan() {
		line := scanner.Text()
		d, _ := strconv.Atoi(line)
		if d > cur {
			result++
		}
		cur = d
	}
	fmt.Println(result)
}

func PartII() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	nums := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		d, _ := strconv.Atoi(line)
		nums = append(nums, d)
	}

	prevSum := math.MaxInt
	var windowStart, curSum, result int

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		curSum += nums[windowEnd]
		if windowEnd > 1 {
			if curSum > prevSum {
				result++
			}
			prevSum = curSum
			curSum -= nums[windowStart]
			windowStart++
		}
	}
	fmt.Println(result)
}

func main() {
	PartII()
}
