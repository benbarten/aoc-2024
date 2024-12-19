package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type robot struct {
	pos      coord
	velocity coord
}

func newGrid() [][]string {
	grid := make([][]string, 103)
	for i := 0; i < len(grid); i++ {
		row := make([]string, 101)
		for j := range row {
			row[j] = "."
		}
		grid[i] = row
	}
	return grid
}

func placeRobots(grid [][]string, robots []robot) [][]string {
	for _, r := range robots {
		grid[r.pos.y][r.pos.x] = "X"
	}
	return grid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(strings.Join(row, " "))
	}
}

var fromRoot = []coord{
	{y: 1, x: 1},
	{y: 1, x: 0},
	{y: 1, x: -1},
	{y: 2, x: -2},
	{y: 2, x: -1},
	{y: 2, x: 0},
	{y: 2, x: 1},
	{y: 2, x: 2},
	{y: 3, x: 0},
}

func isValid(grid [][]string, x, y int) bool {
	for _, c := range fromRoot {
		nextY := y + c.y
		nextX := x + c.x

		if nextY < 0 || nextY >= 103 || nextX < 0 || nextX >= 101 || grid[nextY][nextX] != "X" {
			return false
		}
	}
	return true
}

func hasChristmasTree(grid [][]string) bool {
	for y := 0; y < 103; y++ {
		for x := 0; x < 101; x++ {
			if grid[y][x] != "X" {
				continue
			}

			if isValid(grid, x, y) {
				return true
			}
		}
	}

	return false
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var robots []robot

	for scanner.Scan() {
		text := scanner.Text()
		re := regexp.MustCompile(`-*\d+`)
		matches := re.FindAllString(text, -1)

		var nums []int
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			nums = append(nums, num)
		}

		robots = append(robots, robot{
			pos:      coord{x: nums[0], y: nums[1]},
			velocity: coord{x: nums[2], y: nums[3]},
		})
	}

	// part 1
	var q1, q2, q3, q4 int

	for _, r := range robots {
		for i := 0; i < 100; i++ {
			r.pos.x = (r.pos.x + r.velocity.x + 101) % 101
			r.pos.y = (r.pos.y + r.velocity.y + 103) % 103
		}

		if r.pos.x < 50 {
			if r.pos.y < 51 {
				q1++
			} else if r.pos.y > 51 {
				q3++
			}
		} else if r.pos.x > 50 {
			if r.pos.y < 51 {
				q2++
			} else if r.pos.y > 51 {
				q4++
			}
		}
	}

	// part 2
	i := 0
	for {
		for j := range robots {
			robots[j].pos.x = (robots[j].pos.x + robots[j].velocity.x + 101) % 101
			robots[j].pos.y = (robots[j].pos.y + robots[j].velocity.y + 103) % 103
		}
		grid := newGrid()
		grid = placeRobots(grid, robots)

		if hasChristmasTree(grid) {
			fmt.Println(i+1)
			printGrid(grid)
			break
		}
		i++
	}
}
