package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Element rune

const (
	Free  Element = '.'
	Wall  Element = '#'
	Start Element = 'S'
	End   Element = 'E'
)

type coord struct {
	x, y int
}

var directions = []coord{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)

	var grid [][]Element
	var start, end coord

	for s.Scan() {
		text := s.Text()

		var row []Element
		for i, r := range text {
			el := Element(r)
			row = append(row, el)

			if el == Start {
				start = coord{x: i, y: len(grid)}
			} else if el == End {
				end = coord{x: i, y: len(grid)}
			}
		}

		grid = append(grid, row)
	}

	dists := make(map[coord]int)
	dist := 0

	// build the regular path
	path := list.New()
	path.PushBack(start)

	for {
		current := path.Back()
		currentVal := current.Value.(coord)
		dists[currentVal] = dist

		if currentVal == end {
			break
		}

		for _, dir := range directions {
			nextX := currentVal.x + dir.x
			nextY := currentVal.y + dir.y
			next := coord{x: nextX, y: nextY}

			if current.Prev() != nil && current.Prev().Value.(coord) == next { // don't go back
				continue
			}

			if grid[nextY][nextX] == Free || grid[nextY][nextX] == End {
				path.PushBack(next)
				break
			}
		}
		dist++
	}

	var cheats []int

	cur := path.Front()
	for cur != nil {
		curVal := cur.Value.(coord)
		curDist := dists[curVal]

		for _, dir := range directions {

			aX := curVal.x + dir.x
			aY := curVal.y + dir.y
			if grid[aY][aX] != Wall {
				continue // a cheat is only if we go through a wall, not jump ahead
			}

			bX := curVal.x + 2*dir.x
			bY := curVal.y + 2*dir.y
			b := coord{x: bX, y: bY}

			if bX >= 0 && bX < len(grid[0]) && bY >= 0 && bY < len(grid) && (grid[bY][bX] == Free || grid[bY][bX] == End) && dists[b] > curDist { // there's a free spot behind the wall and we're not going back
				diff := dists[b] - curDist - 2
				if diff >= 100 {
					cheats = append(cheats, diff)
				}
			}
		}

		cur = cur.Next()
	}

	// part 1
	fmt.Println(len(cheats))

	cheats = []int{}

	type state struct {
		current *coord
		last    *coord
		dist    int
		maxDist int
	}

	type startEnd struct {
		start, end coord
	}

	cur = path.Front()
	for cur != nil {
		curVal := cur.Value.(coord)
		curDist := dists[curVal]

		start := state{
			current: &curVal,
			last:    nil,
			dist:    0,
			maxDist: curDist,
		}

		visited := make(map[coord]bool)
		uniqueCheats := make(map[startEnd]int)

		q := list.New()
		q.PushBack(start)
		for q.Len() > 0 {
			currentState := q.Front().Value.(state)
			q.Remove(q.Front())

			//fmt.Printf("X: %d Y: %d Dist: %d Wall: %v\n", currentState.current.x, currentState.current.y, currentState.dist, currentState.crossedWall)

			if visited[*currentState.current] {
				continue
			}

			if currentState.dist > 20 {
				continue
			}

			visited[*currentState.current] = true

			qX := currentState.current.x
			qY := currentState.current.y
			qCell := grid[qY][qX]

			if (qCell == Free || qCell == End) {
				diff := dists[*currentState.current] - curDist - currentState.dist
				p := startEnd{start: curVal, end: *currentState.current}
				if diff >= 100 { // TODO change this for final solution
					uniqueCheats[p] = diff
				}
			}

			for _, dir := range directions {
				aX := qX + dir.x
				aY := qY + dir.y
				a := coord{x: aX, y: aY}

				if currentState.last != nil && a == *currentState.last {
					continue // avoid going back
				}

				if aX < 0 || aX >= len(grid[0]) || aY < 0 || aY >= len(grid) {
					continue
				}

				next := state{
					last:    currentState.current,
					current: &a,
					dist:    currentState.dist + 1,
				}

				q.PushBack(next)
			}
		}

		for _, uc := range uniqueCheats {
			cheats = append(cheats, uc)
		}

		cur = cur.Next()
	}

	fmt.Println(len(cheats))
}
