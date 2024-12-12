package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func next(input string, end int) int {
	return min(len(input)-1, end+1)
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := io.ReadAll(f)

	input := string(b)

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var sum int

	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])

		sum += left * right
	}

	// Part 1
	fmt.Println(sum)

	sum = 0

	commandRegex := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
    matches = commandRegex.FindAllStringSubmatch(input, -1)

    enabled := true  // enable from the start

    for _, match := range matches {
        command := match[1]
        
        switch {
        case command == "do()":
            enabled = true
        case command == "don't()":
            enabled = false
        case enabled && command[:3] == "mul":
            left, _ := strconv.Atoi(match[2])
            right, _ := strconv.Atoi(match[3])
            sum += left * right
        }
    }

	fmt.Println(sum)
}
