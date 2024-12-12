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

	cases := make(map[int][]int)

	for scanner.Scan() {
		text := scanner.Text()
		
		before, after, _ := strings.Cut(text, ":")
		result, _ := strconv.Atoi(before)
		parts := strings.Split(after, " ")
		
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			cases[result] = append(cases[result], num)
		}
	}

	var sum int
	for result, nums := range cases {
		if eval(nums[0], nums, 1, result) {
			sum += result
		}
	}

	fmt.Println(sum)
}

func eval(current int, nums []int, i int, result int) bool {
	if i == len(nums) {
		return current == result
	}

	added := current + nums[i]
	multiplied := current * nums[i]
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, nums[i]))

	return eval(added, nums, i+1, result) || eval(multiplied, nums, i+1, result) || eval(concatenated, nums, i+1, result)
}