package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

var directions = []xy{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

const gridSize = 70
const fallenBytes = 1024

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var coords []xy
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		coord := xy{
			x: x,
			y: y,
		}
		coords = append(coords, coord)
	}

	grid := make([][]int, gridSize+1)
	for i := 0; i <= gridSize; i++ {
		row := make([]int, gridSize+1)
		grid[i] = row
	}

	for i := 0; i < fallenBytes; i++ {
		grid[coords[i].y][coords[i].x] = 1
	}

	result := traverse(grid)

	fmt.Println(result)

	i := fallenBytes-1
	for traverse(grid) != -1 {
		grid[coords[i].y][coords[i].x] = 1
		if traverse(grid) == -1 {
			break
		}
		i++
	}
	fmt.Printf("%d,%d\n", coords[i].x, coords[i].y)
}

func traverse(grid [][]int) int {
	type node struct {
		xy   xy
		dist int
	}

	start := node{xy: xy{0, 0}, dist: 0}
	end := xy{gridSize, gridSize}

	visited := make(map[xy]bool)

	q := list.New()
	q.PushBack(start)

	for q.Len() > 0 {
		front := q.Front()
		q.Remove(front)

		val := front.Value.(node)

		if visited[val.xy] {
			continue
		}

		visited[val.xy] = true

		if val.xy == end {
			return val.dist
		}

		for _, dir := range directions {
			nextX := val.xy.x + dir.x
			nextY := val.xy.y + dir.y
			nextXY := xy{x: nextX, y: nextY}

			if nextX >= 0 && nextX <= gridSize && nextY >= 0 && nextY <= gridSize && grid[nextY][nextX] == 0 {
				next := node{xy: nextXY, dist: val.dist + 1}
				q.PushBack(next)
			}
		}
	}

	return -1
}
