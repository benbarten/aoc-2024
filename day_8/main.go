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
		row := []rune(scanner.Text())
		grid = append(grid, row)
	}

	type coord struct {
		y, x int
	}

	zero := '.'
	frequencies := make(map[rune][]coord)

	rows := len(grid)
	cols := len(grid[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			val := grid[y][x]
			if val == zero {
				continue
			}

			c := coord{y: y, x: x}
			frequencies[val] = append(frequencies[val], c)
		}
	}

	part1 := make(map[coord]bool)
	part2 := make(map[coord]bool)

	for _, antennas := range frequencies {
		if len(antennas) == 1 {
			continue
		}

		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				first, second := antennas[i], antennas[j]
				part2[first] = true
				part2[second] = true

				deltaY := second.y - first.y
				deltaX := second.x - first.x

				before := coord{
					y: first.y - deltaY,
					x: first.x - deltaX,
				}
				
				// part 1
				if before.y >= 0 && before.y < rows && before.x >= 0 && before.x < cols {
					part1[before] = true
				}

				// part 2
				for before.y >= 0 && before.y < rows && before.x >= 0 && before.x < cols {
					part2[before] = true
					before.y -= deltaY
					before.x -= deltaX
				} 
				
				after := coord{
					y: second.y + deltaY,
					x: second.x + deltaX,
				}

				if after.y >= 0 && after.y < rows && after.x >= 0 && after.x < cols {
					part1[after] = true
				}

				for after.y >= 0 && after.y < rows && after.x >= 0 && after.x < cols {
					part2[after] = true
					after.y += deltaY
					after.x += deltaX
				}
			}
		}
	}

	fmt.Println(len(part1))
	fmt.Println(len(part2))
}
