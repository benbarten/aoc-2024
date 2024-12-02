package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"container/heap"
)

type minHeap []int64

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int64))
}

func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m *minHeap) Peek() any {
	return (*m)[0]
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open input file")
	}

	scanner := bufio.NewScanner(f)

	left1, right1 := &minHeap{}, &minHeap{}
	left2, right2 := &minHeap{}, &minHeap{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		l, _ := strconv.ParseInt(parts[0], 10, 64)
		r, _ := strconv.ParseInt(parts[3], 10, 64)

		// create heaps for both parts
		heap.Push(left1, l)
		heap.Push(left2, l)
		heap.Push(right1, r)
		heap.Push(right2, r)
	}

	var sum int64

	for left1.Len() > 0 && right1.Len() > 0 {
		lv, rv := heap.Pop(left1), heap.Pop(right1)
		sum += abs(lv.(int64) - rv.(int64))
	}

	// Part 1
	fmt.Println(sum)


	var score int64

	for left2.Len() > 0 {
		lv := heap.Pop(left2).(int64)

		var multiplier int64
		for right2.Len() > 0 && right2.Peek().(int64) <= lv {
			rv := heap.Pop(right2).(int64)

			if rv == lv {
				multiplier++
			}
		}

		score += lv * multiplier
	}

	fmt.Println(score)
}
