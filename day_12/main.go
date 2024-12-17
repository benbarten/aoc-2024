package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	directions := [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	rows := len(grid)
	cols := len(grid[0])

	var price, price2 int

	visited := make(map[[2]int][][2]int) // stores which nodes were visited and which directions had fences

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			var area, perimeter, perimeter2 int

			q := [][2]int{{y, x}}

			for len(q) > 0 {
				current := q[0]
				q = q[1:]

				if _, ok := visited[current]; ok {
					continue
				}

				area++
				fences := [][2]int{}

				for i, dir := range directions {
					nextY := current[0] + dir[0]
					nextX := current[1] + dir[1]
					next := [2]int{nextY, nextX}

					// if we hit an edge of the map or another plant, we need a fence
					if nextY < 0 || nextY >= rows || nextX < 0 || nextX >= cols || grid[nextY][nextX] != grid[y][x] {
						fences = append(fences, dir)
						perimeter++
						perimeter2++

						// part 2: if we have a fence to the left, we want to check one elem higher and lower, whether we
						// we have visited them already and they have a fence on the left. Same for the right. Same if
						// we hit a fence on the top of the bottom, we check the elements to the right and left.

						var compareDirections [][2]int
						if i%2 == 0 { // top / bottom
							compareDirections = append(compareDirections, directions[1], directions[3])
						} else { // left / right
							compareDirections = append(compareDirections, directions[0], directions[2])
						}

						for _, compareDir := range compareDirections {
							compareY := current[0] + compareDir[0]
							compareX := current[1] + compareDir[1]
							compareNext := [2]int{compareY, compareX}

							if compareY < 0 || compareY >= rows || compareX < 0 || compareX >= cols || grid[compareY][compareX] != grid[y][x] {
								continue
							}

							compareFences, ok := visited[compareNext]
							if !ok {
								continue
							}

							var hasOverlap bool
							for _, fence := range compareFences {
								if fence == dir {
									hasOverlap = true
									break
								}
							}

							if hasOverlap {
								perimeter2--
							}
						}

						continue
					}

					q = append(q, next)
				}

				visited[current] = fences
			}

			price += area * perimeter
			price2 += area * perimeter2
		}
	}

	fmt.Println(price)
	fmt.Println(price2)
}
