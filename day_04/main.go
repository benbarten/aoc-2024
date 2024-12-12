package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	directions := [][2]int{
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, 1},   // right
		{1, 1},   // down-right
		{1, 0},   // down
		{1, -1},  // down-left
		{0, -1},  // left
	}

	diagonals := map[int]int{
		0: 4,
		2: 6,
		4: 0,
		6: 2,
	}

	pattern1 := "XMAS"
	pattern2 := "MAS"

	countPart1 := 0
	countPart2 := 0
	rows := len(grid)
	cols := len(grid[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			// part 1
			for _, dir := range directions {

				valid := true

				for i := 0; i < len(pattern1); i++ {
					ny := y + dir[0]*i
					nx := x + dir[1]*i

					if ny < 0 || ny >= rows || nx < 0 || nx >= cols || grid[ny][nx] != pattern1[i] {
						valid = false
						break
					}
				}

				if valid {
					countPart1++
				}
			}

			// part 2
			if grid[y][x] == 'A' {
				validCount := 0

				for i, dir := range directions {
					if i % 2 != 0 {
						continue
					}

					ny := y + dir[0]
					nx := x + dir[1]
					if ny < 0 || ny >= rows || nx < 0 || nx >= cols || grid[ny][nx] != pattern2[0] {
						continue
					}

					ny = y + directions[diagonals[i]][0]
					nx = x + directions[diagonals[i]][1]

					if ny < 0 || ny >= rows || nx < 0 || nx >= cols || grid[ny][nx] != pattern2[2] {
						continue
					}	
					validCount++
				}

				if validCount == 2 {
					countPart2++
				}
			}
		}
	}

	// part 1
	fmt.Println(countPart1)
	fmt.Println(countPart2)
}
