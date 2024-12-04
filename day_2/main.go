package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidDiff(a, b int64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff >= 1 && diff <= 3
}

func isValidSequence(nums []int64) bool {
	if len(nums) < 2 {
		return true
	}
	
	increasing := nums[1] > nums[0]
	
	for i := 1; i < len(nums); i++ {
		if !isValidDiff(nums[i], nums[i-1]) {
			return false
		}
		if increasing && nums[i] <= nums[i-1] {
			return false
		}
		if !increasing && nums[i] >= nums[i-1] {
			return false
		}
	}
	return true
}

func isValidWithRemoval(nums []int64) bool {
	if isValidSequence(nums) {
		return true
	}
	
	// Try removing each number once
	for i := range nums {
		sequence := make([]int64, 0, len(nums)-1)
		sequence = append(sequence, nums[:i]...)
		sequence = append(sequence, nums[i+1:]...)
		
		if isValidSequence(sequence) {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open input file")
	}
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	var levels [][]int64
	
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		level := make([]int64, 0, len(parts))
		for _, part := range parts {
			val, _ := strconv.ParseInt(part, 10, 64)
			level = append(level, val)
		}
		levels = append(levels, level)
	}
	
	// Part 1
	var count1 int64
	for _, level := range levels {
		if isValidSequence(level) {
			count1++
		}
	}
	fmt.Println("Part 1:", count1)
	
	// Part 2
	var count2 int64
	for _, level := range levels {
		if isValidWithRemoval(level) {
			count2++
		}
	}
	fmt.Println("Part 2:", count2)
}