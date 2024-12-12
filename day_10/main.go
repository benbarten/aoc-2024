package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		text := scanner.Text()
		line := make([]int, 0, len(text))
		for _, r := range text {
			line = append(line, int(r-'0'))
		}
		grid = append(grid, line)
	}

	rows := len(grid)
	cols := len(grid[0])

	type point struct {
		y, x int
		val  int
	}

	directions := [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	var count1 int
	var count2 int

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] != 0 {
				continue
			}

			// multiple paths can lead to the same 9, so we need to deduplicate for each trailhead
			visited := make(map[point]bool)

			q := []point{{y: y, x: x, val: 0}}
			for len(q) > 0 {
				current := q[len(q)-1]
				q = q[:len(q)-1]

				if current.val == 9 {
					visited[current] = true
					count2++
					continue
				}

				for _, dir := range directions {
					nextY := current.y + dir[0]
					nextX := current.x + dir[1]

					if nextY >= 0 && nextY < rows && nextX >= 0 && nextX < cols && grid[nextY][nextX] == current.val+1 {
						q = append(q, point{y: nextY, x: nextX, val: grid[nextY][nextX]})
					}
				}
			}

			count1 += len(visited)
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}
