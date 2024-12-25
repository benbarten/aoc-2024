package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Element rune

const (
	Wall  Element = '#'
	Free  Element = '.'
	Start Element = 'S'
	End   Element = 'E'
)

type Pos struct {
	X, Y int
	Dir  int
}

var directions = []Pos{
	{Y: -1, X: 0}, // up
	{Y: 0, X: 1},  // right
	{Y: 1, X: 0},  // down
	{Y: 0, X: -1}, // left
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var grid [][]Element
	start, end := Pos{}, Pos{}

	for scanner.Scan() {
		text := scanner.Text()

		var row []Element
		for i, r := range text {
			row = append(row, Element(r))

			if Element(r) == Start {
				start.X = i
				start.Y = len(grid)
			} else if Element(r) == End {
				end.X = i
				end.Y = len(grid)
			}
		}

		grid = append(grid, row)
	}

	// init distance map for each node for dijkstra's shortest path
	visited := make(map[Pos]int, len(grid))
	predecessors := make(map[Pos]map[Pos]struct{}, len(grid))

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			for i := range directions {
				pos := Pos{Y: y, X: x, Dir: i}
				visited[pos] = math.MaxInt
				predecessors[pos] = make(map[Pos]struct{})
			}
		}
	}

	start.Dir = 1

	s := &solution{grid: grid, visited: visited, predecessors: predecessors}
	s.traverse(Pos{}, start, 0, true)

	// part 1
	fmt.Println(s.visited[end])

	endPredecessors := make(map[Pos]bool)

	q := []Pos{end}
	for len(q) > 0 {
		next := []Pos{}
		for _, e := range q {
			endPredecessors[e] = true

			for k := range predecessors[e] {
				next = append(next, k)
			}
		}

		q = next
	}

	deduplicated := make(map[Pos]bool)
	for k := range endPredecessors {
		pos := Pos{Y: k.Y, X: k.X}
		deduplicated[pos] = true
	}

	// part 2
	fmt.Println(len(deduplicated))
}

type solution struct {
	grid         [][]Element
	visited      map[Pos]int
	predecessors map[Pos]map[Pos]struct{}
	end          Pos
}

func (s *solution) traverse(last, current Pos, sum int, isMove bool) {
	if s.grid[current.Y][current.X] == Wall {
		return
	}

	if sum > s.visited[current] {
		return
	} else if sum == s.visited[current] && current != last { // found another shortest path to this node
		s.predecessors[current][last] = struct{}{}
	} else { // found a new new shortest path
		if last.X != 0 && last.Y != 0 && current != last {
			s.predecessors[current] = make(map[Pos]struct{})
			s.predecessors[current][last] = struct{}{}
		}
	}

	s.visited[current] = sum

	if isMove { // we just moved here, we need to turn straight, and 90 degrees to both sides
		for i := range directions {
			if (current.Dir%2) == (i%2) && current.Dir != i { // don't go back where we came from
				continue
			}

			newSum := sum

			if (current.Dir%2) != (i%2) { // 90 degree turn
				newSum += 1000
			}
			next := current
			next.Dir = i
			s.traverse(current, next, newSum, false)

		}
	} else {
		newSum := sum
		newSum++

		nextY := current.Y + directions[current.Dir].Y
		nextX := current.X + directions[current.Dir].X
		next := Pos{Y: nextY, X: nextX, Dir: current.Dir}

		s.traverse(current, next, newSum, true) // move to next field
	}
}
