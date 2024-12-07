package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var grid [][]bool
    var startY, startX int

	for scanner.Scan() {
		line := scanner.Text()

		row := make([]bool, 0, len(line))

		for i := range line {
			switch line[i] {
			case '.':
				row = append(row, true)
			case '#':
				row = append(row, false)
			case '^':
				row = append(row, true)
				startY = len(grid)
				startX = i
			}
		}

		grid = append(grid, row)
	}

	directions := [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	cols := len(grid)
	row := len(grid[0])
	y, x := startY, startX

	dir := 0 // start walking up
	visited := make(map[[2]int]bool)

	for {
		visited[[2]int{y, x}] = true

		nextY := y + directions[dir][0]
		nextX := x + directions[dir][1]

		if nextX < 0 || nextX >= row || nextY < 0 || nextY >= cols {
			break
		}

		// turn right until there is a free path
		for !grid[nextY][nextX] {
			dir = (dir + 1) % 4
			nextY = y + directions[dir][0]
			nextX = x + directions[dir][1]

			if nextX < 0 || nextX >= row || nextY < 0 || nextY >= cols {
				break
			}
		}

		y = nextY
		x = nextX
	}

	// part 1
	fmt.Println(len(visited))
	

	validObstacles := 0
    
	// brute force by inserting an obstacle in every possible position apart from the start and existing obstacles and then
	// let the guard run through it

    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[0]); x++ {
            if !grid[y][x] || (y == startY && x == startX) {
                continue
            }

            grid[y][x] = false
            
            if createsLoop(grid, startY, startX, directions) {
                validObstacles++
            }
            
            grid[y][x] = true
        }
    }

	// part 2
    fmt.Println(validObstacles)
}

type Position struct {
    y, x, dir int
}

func createsLoop(grid [][]bool, startY, startX int, directions [][2]int) bool {
    visited := make(map[Position]bool)
    y, x := startY, startX
    dir := 0 

    for {
        pos := Position{y, x, dir}
        if visited[pos] {
            return true // found a loop
        }

        visited[pos] = true

        nextY := y + directions[dir][0]
        nextX := x + directions[dir][1]

        if nextX < 0 || nextX >= len(grid[0]) || nextY < 0 || nextY >= len(grid) {
            return false
        }

        originalDir := dir
        for !grid[nextY][nextX] {
            dir = (dir + 1) % 4
            if dir == originalDir { // stop when all directions were tried
                return false
            }

            nextY = y + directions[dir][0]
            nextX = x + directions[dir][1]
            if nextX < 0 || nextX >= len(grid[0]) || nextY < 0 || nextY >= len(grid) {
                return false
            }
        }

        y, x = nextY, nextX
    }
}
