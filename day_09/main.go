package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	b, _ := os.ReadFile("input.txt")

	var diskmap []int
	for _, r := range string(b) {
		if !unicode.IsNumber(r) {
			panic("invalid input")
		}

		diskmap = append(diskmap, int(r-'0'))
	}

	var part1 []rune
	var part2 []rune
	fileID := 0

	for i, d := range diskmap {
		var token rune
		if i%2 == 0 { // file
			token = rune('0' + fileID)
			fileID++
		} else { // free space
			token = '.'
		}

		for i := 0; i < d; i++ {
			part1 = append(part1, token)
			part2 = append(part2, token)
		}
	}

	start := 0
	end := len(part1) - 1

	for start < end {
		if part1[start] != '.' {
			start++
			continue
		}

		if part1[end] == '.' {
			end--
			continue
		}

		part1[start] = part1[end]
		part1[end] = '.'
	}

	// part 1
	fmt.Println(checksum(part1))

	start = 0
	windowStart := len(part2) - 1
	windowEnd := len(part2) - 1

	for windowStart > 0 && start < len(part2) {
		// iterate until we find a window
		if part2[windowEnd] == '.' {
			windowEnd--
			windowStart = windowEnd
			continue
		}

		// determine the window to move
		if part2[windowStart-1] == part2[windowEnd] {
			windowStart--
			continue
		}

		// find the first .
		if part2[start] != '.' {
			start++
			continue
		}

		// we did not find a fitting window
		if start >= windowStart {
			windowEnd = windowStart - 1
			windowStart = windowEnd
			start = 0
			continue
		}
	
		// find the first fitting free space
		freeWindowStart := start
		freeWindowEnd := start
		
		// length of free space to find
		windowLen := windowEnd - windowStart + 1

		for freeWindowEnd < windowStart {
			freeWindowLen := (freeWindowEnd - freeWindowStart + 1)

			if freeWindowLen < windowLen && part2[freeWindowEnd+1] == '.' {
				freeWindowEnd++
				continue
			}
			
			// current free window is too small, find start of next free window
			if freeWindowLen < windowLen {
				start = freeWindowEnd+1
			} else {
				for i := 0; i < windowLen; i++ {
					part2[freeWindowStart+i] = part2[windowStart+i]
					part2[windowStart+i] = '.'
				}
				windowEnd = windowStart-1
				windowStart = windowEnd
				start = 0
			}
			break
		}
	}

	fmt.Println(checksum(part2))
}

func checksum(arr []rune) int {
	checksum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '.' {
			continue
		}
		checksum += i * int(arr[i]-'0')
	}

	return checksum
}
