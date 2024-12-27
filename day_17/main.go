package main

import (
	"container/list"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type State struct {
	A, B, C int
	iptr    int
	result  []int
}

type Instruction interface {
	Exec(s *State, op int)
}

func (s *State) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return s.A
	case 5:
		return s.B
	case 6:
		return s.C
	}

	// 7 is invalid, so we don't check for it here
	return -1
}

func (s *State) print() {
	fmt.Printf("A: %b %d\n", s.A, s.A)
	fmt.Printf("B: %b %d\n", s.B, s.B)
	fmt.Printf("C: %b %d\n", s.C, s.C)
}

type adv struct{}

func (a adv) Exec(s *State, op int) {
	s.A >>= s.combo(op)
}

type bxl struct{}

func (b bxl) Exec(s *State, op int) {
	s.B ^= op
}

type bst struct{}

func (b bst) Exec(s *State, op int) {
	s.B = s.combo(op) % 8
}

type jnz struct{}

func (j jnz) Exec(s *State, op int) {
	if s.A == 0 {
		return
	}

	s.iptr = op
}

type bxc struct{}

func (b bxc) Exec(s *State, op int) {
	s.B ^= s.C
}

type out struct{}

func (o out) Exec(s *State, op int) {
	s.result = append(s.result, s.combo(op) % 8)
}

type bdv struct{}

func (b bdv) Exec(s *State, op int) {
	s.B = s.A >> s.combo(op)
}

type cdv struct{}

func (c cdv) Exec(s *State, op int) {
	s.C = s.A >> s.combo(op)
}

var instructions = map[int]Instruction{
	0: adv{},
	1: bxl{},
	2: bst{},
	3: jnz{},
	4: bxc{},
	5: out{},
	6: bdv{},
	7: cdv{},
}

var insNames = map[int]string{
	0: "adv",
	1: "bxl",
	2: "bst",
	3: "jnz",
	4: "bxc",
	5: "out",
	6: "bdv",
	7: "cdv",
}

func main() {

	seq := []int{0, 3, 5, 4, 3, 0}
	s := &State{A: 117440}

	fmt.Printf("A: %b\n", s.A)

	for s.iptr < len(seq) {
		opcode := seq[s.iptr]
		operand := seq[s.iptr+1]
		s.iptr += 2
		instructions[opcode].Exec(s, operand)
	}

	var result []string
	for _, r := range s.result {
		result = append(result, strconv.Itoa(r))
	}

	// part 1
	fmt.Println("Out: ", strings.Join(result, ","))

	seq = []int{2, 4, 1, 5, 7, 5, 1, 6, 0, 3, 4, 2, 5, 5, 3, 0}

	end := &State{
		A:      0,
		result: seq,
		iptr:   len(seq),
	}

	part2 := reverse(end, seq)
	slices.Sort(part2)

	// part 2
	fmt.Println(part2[0])

	// validate
	s = &State{A: part2[0]}
	for s.iptr < len(seq) {
		opcode := seq[s.iptr]
		operand := seq[s.iptr+1]
		s.iptr += 2
		instructions[opcode].Exec(s, operand)
	}

	var validateResult []string
	for _, r := range s.result {
		validateResult = append(validateResult, strconv.Itoa(r))
	}

	fmt.Println("Validate out: ", strings.Join(validateResult, ","))
}

func reverse(end *State, seq []int) []int {
    type QueueItem struct {
        a    int
        n    int
    }
    
    queue := list.New()
    queue.PushBack(QueueItem{a: 0, n: 1})
    
    for queue.Len() > 0 {
        item := queue.Remove(queue.Front()).(QueueItem)
        a, n := item.a, item.n
        
        if n > len(seq) {
            return []int{a}
        }
        
        for i := 0; i < 8; i++ {
            a2 := (a << 3) | i
            
            s := &State{A: a2}
            target := seq[len(seq)-n:]
            
            for s.iptr < len(seq) {
                opcode := seq[s.iptr]
                operand := seq[s.iptr+1]
                s.iptr += 2
                instructions[opcode].Exec(s, operand)
            }
            
            if matchesOutput(s.result, target) {
                queue.PushBack(QueueItem{a: a2, n: n + 1})
            }
        }
    }
    
    return nil
}

func matchesOutput(output []int, target []int) bool {
    if len(output) != len(target) {
        return false
    }
    for i := range output {
        if output[i] != target[i] {
            return false
        }
    }
    return true
}
