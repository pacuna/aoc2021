package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getData(filepath string) [][]int {
	f, _ := os.Open(filepath)

	scanner := bufio.NewScanner(f)

	data := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		num := []int{}
		for _, c := range line {
			d, _ := strconv.Atoi(string(c))
			num = append(num, d)
		}
		data = append(data, num)
	}
	return data
}

func findMostCommon(data [][]int, pos int) int {
	var zeroes, ones int
	for _, n := range data {
		if n[pos] == 1 {
			ones++
		} else {
			zeroes++
		}
	}
	if ones >= zeroes {
		return 1
	} else {
		return 0
	}
}

func DtoB(n []int) int {
	res := 0
	for idx, d := range n {
		res += d * int(math.Pow(2, float64(len(n)-idx-1)))
	}
	return res
}

func main() {
	data_oxy := getData("input.txt")
	data_co2 := getData("input.txt")

	pos := 0
	for len(data_oxy) != 1 {
		tmp := [][]int{}
		mcv := findMostCommon(data_oxy, pos)
		for _, el := range data_oxy {
			if el[pos] == mcv {
				tmp = append(tmp, el)
			}
		}
		data_oxy = tmp
		pos++
	}

	pos = 0
	for len(data_co2) != 1 {
		tmp := [][]int{}
		mcv := findMostCommon(data_co2, pos)
		for _, el := range data_co2 {
			if el[pos] != mcv {
				tmp = append(tmp, el)
			}
		}
		data_co2 = tmp
		pos++
	}

	fmt.Println(data_oxy, data_co2, DtoB(data_oxy[0]), DtoB(data_co2[0]))
	fmt.Println(DtoB(data_oxy[0]) * DtoB(data_co2[0]))
}
