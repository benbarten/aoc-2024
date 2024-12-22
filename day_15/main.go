package main

import (
	"bufio"
	"fmt"
	"os"

)

type Element rune

const (
	Wall          Element = '#'
	Box           Element = 'O'
	WideboxOpen   Element = '['
	WideboxClosed Element = ']'
	Free          Element = '.'
	Robot         Element = '@'
)

type Direction rune

const (
	Left  Direction = '<'
	Right Direction = '>'
	Up    Direction = '^'
	Down  Direction = 'v'
)

type Position struct {
	X, Y int
}

type Widebox struct {
	X0, X1 int
	Y      int
}

func move(p Position, d Direction) Position {
	switch d {
	case Up:
		return Position{X: p.X, Y: p.Y - 1}
	case Right:
		return Position{X: p.X + 1, Y: p.Y}
	case Down:
		return Position{X: p.X, Y: p.Y + 1}
	case Left:
		return Position{X: p.X - 1, Y: p.Y}
	}
	return Position{}
}

func isElement(grid [][]Element, p Position, e Element) bool {
	return grid[p.Y][p.X] == e
}

func printGrid(grid [][]Element) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	// grid for part 1
	/*
		##########
		#..O..O.O#
		#......O.#
		#.OO..O.O#
		#..O@..O.#
		#O#..O...#
		#O..O..O.#
		#.OO.O.OO#
		#....O...#
		##########
	*/
	var grid [][]Element
	pos := Position{}

	// grid for part 2
	/*
		####################
		##[].......[].[][]##
		##[]...........[].##
		##[]........[][][]##
		##[]......[]....[]##
		##..##......[]....##
		##..[]............##
		##..@......[].[][]##
		##......[][]..[]..##
		####################
	*/
	var grid2 [][]Element
	pos2 := Position{}

	var directions []Direction

	var scanDirections bool
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanDirections = true
			continue
		}

		var row []Element
		var row2 []Element

		for i, r := range text {
			if !scanDirections {
				row = append(row, Element(r))

				if Element(r) == Robot {
					pos.X = i
					pos.Y = len(grid)
				}

				switch Element(r) {
				case Box:
					row2 = append(row2, WideboxOpen, WideboxClosed)
				case Robot:
					pos2.X = len(row2)
					pos2.Y = len(grid2)
					row2 = append(row2, Robot, Free)
				default:
					row2 = append(row2, Element(r), Element(r))
				}
			} else {
				directions = append(directions, Direction(r))
			}
		}

		grid = append(grid, row)
		grid2 = append(grid2, row2)
	}

	// walk part 1
	for _, dir := range directions {
		next := move(pos, dir)

		if isElement(grid, next, Wall) {
			continue
		}

		if isElement(grid, next, Free) {
			grid[pos.Y][pos.X] = Free
			pos = next
			continue
		}

		// we hit a box at next: @O

		nextNext := move(next, dir)
		for isElement(grid, nextNext, Box) { // @OO - continue
			nextNext = move(nextNext, dir)
		}

		if isElement(grid, nextNext, Wall) { // @OO# - do nothing
			continue
		}

		// @OO.

		// free space at nextNext, move everything one further by moving the first box to the free space,
		// moving the robot one ahead and free up the current position: .@OO
		grid[nextNext.Y][nextNext.X] = Box
		grid[next.Y][next.X] = Robot
		grid[pos.Y][pos.X] = Free
		pos = next
	}

	var sum int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == Box {
				sum += 100*y + x
			}
		}
	}

	// part 1
	fmt.Println(sum)

	grid = grid2
	pos = pos2
	printGrid(grid)

	// walk part 2
	for _, dir := range directions {
		fmt.Println("Direction: ", string(dir))

		next := move(pos, dir)

		if isElement(grid, next, Wall) {
			continue
		}

		if isElement(grid, next, Free) {
			grid[next.Y][next.X] = Robot
			grid[pos.Y][pos.X] = Free
			pos = next
			continue
		}

		// we hit a box at next: @[ / @]

		if dir == Left || dir == Right {
			nextNext := move(next, dir)

			for isElement(grid, nextNext, WideboxOpen) || isElement(grid, nextNext, WideboxClosed) { // @[][] - continue
				nextNext = move(nextNext, dir)
			}

			if isElement(grid, nextNext, Wall) { // @OO# - do nothing
				continue
			}

			if dir == Right {
				// @[][].
				for i := nextNext.X; i > pos.X; i-- {
					grid[pos.Y][i], grid[pos.Y][i-1] = grid[pos.Y][i-1], grid[pos.Y][i]
				}
			} else {
				// .[][]@
				for i := nextNext.X; i < pos.X; i++ {
					grid[pos.Y][i], grid[pos.Y][i+1] = grid[pos.Y][i+1], grid[pos.Y][i]
				}
			}
			grid[pos.Y][pos.X] = Free
			grid[next.Y][next.X] = Robot
			pos = next
			continue
		}

		// for up/down movements, we cannot just determine the width of the widest level, we need to handle each level individually
		// .#.......
		// ....[]....
		// .[]..[]...
		// ..[][]....
		// ...[].....
		// ....@.....

		// determine starting widebox
		var wb Widebox
		wb.Y = next.Y
		if grid[next.Y][next.X] == WideboxOpen {
			wb.X0 = next.X
			wb.X1 = next.X + 1
		} else {
			wb.X0 = next.X - 1
			wb.X1 = next.X
		}

		t, ok := tree(grid, wb, dir)

		if !ok { // some box is in the way
			continue
		}

		// pop off each widebox, move into direction and paint old one free
		for i := len(t) - 1; i >= 0; i-- {
			nwb := t[i]

			var nextY int
			switch dir {
			case Up:
				nextY = nwb.Y - 1
			case Down:
				nextY = nwb.Y + 1
			}

			grid[nextY][nwb.X0] = WideboxOpen
			grid[nextY][nwb.X1] = WideboxClosed

			grid[nwb.Y][nwb.X0] = Free
			grid[nwb.Y][nwb.X1] = Free
		}

		grid[pos.Y][pos.X] = Free
		grid[next.Y][next.X] = Robot
		pos = next
	}

	sum = 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == WideboxOpen {
				sum += 100*y + x
			}
		}
	}

	fmt.Println(sum)
}

func tree(grid [][]Element, wb Widebox, dir Direction) ([]Widebox, bool) {
	all := []Widebox{wb}
	q := map[Widebox]struct{}{
		wb: {},
	}

	for len(q) > 0 {
		current := q
		q = make(map[Widebox]struct{})
		
		for qwb := range current {
			var nextY int
			switch dir {
			case Up:
				nextY = qwb.Y - 1
			case Down:
				nextY = qwb.Y + 1
			}

			// once one box is blocked, we can't move the entire tree
			if grid[nextY][qwb.X0] == Wall || grid[nextY][qwb.X1] == Wall {
				return nil, false
			}

			if grid[nextY][qwb.X0] == WideboxOpen {
				// ..[]..
				// ..[]..
				nwb := Widebox{Y: nextY, X0: qwb.X0, X1: qwb.X1}
				q[nwb] = struct{}{}
			}

			if grid[nextY][qwb.X0] == WideboxClosed {
				// .[]...
				// ..[]..
				nwb := Widebox{Y: nextY, X0: qwb.X0 - 1, X1: qwb.X0}
				q[nwb] = struct{}{}
			}

			if grid[nextY][qwb.X1] == WideboxOpen {
				// ...[].
				// ..[]..
				nwb := Widebox{Y: nextY, X0: qwb.X1, X1: qwb.X1 + 1}
				q[nwb] = struct{}{}
			}
		}

		for el := range q {
			all = append(all, el)
		}
	}

	return all, true
}
