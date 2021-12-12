package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetData() [][]int {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	nums := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		tmp := []int{}
		for _, c := range chars {
			d, _ := strconv.Atoi(c)
			tmp = append(tmp, d)
		}
		nums = append(nums, tmp)
	}
	return nums
}

func GetDec(b []int) int {
	res := 0
	for i := 0; i < len(b); i++ {
		res += b[i] * int(math.Pow(2, float64(len(b)-i-1)))
	}
	return res
}

func SolveI() {
	nums := GetData()
	gammaRate := make([]int, len(nums[0]))
	epsilonRate := make([]int, len(nums[0]))

	for i := range nums[0] {
		tmp := 0
		for _, n := range nums {
			if n[i] == 1 {
				tmp += 1
			} else {
				tmp -= 1
			}
		}
		if tmp >= 0 {
			gammaRate[i] = 1
			epsilonRate[i] = 0
		} else {
			gammaRate[i] = 0
			epsilonRate[i] = 1
		}
	}
	fmt.Println(GetDec(gammaRate) * GetDec(epsilonRate))
}

func main() {
	SolveI()
}
