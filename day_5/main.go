package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	rules := make(map[int]map[int]bool)
	var updates [][]int

	scanningRules := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanningRules = false
			continue
		}

		if scanningRules {
			left, right, _ := strings.Cut(line, "|")
			before, _ := strconv.Atoi(left)
			after, _ := strconv.Atoi(right)
			if _, ok := rules[before]; !ok {
				rules[before] = make(map[int]bool)
			}

			rules[before][after] = true
		} else {
			parts := strings.Split(line, ",")
			pages := make([]int, 0, len(parts))
			for _, p := range parts {
				page, _ := strconv.Atoi(p)
				pages = append(pages, page)
			}
			updates = append(updates, pages)
		}
	}

	type node struct {
		val  int
		next *node
	}

	var validSum, invalidSum int

	for _, update := range updates {
		head := &node{
			val: update[0],
		}

		isValid := true
		
		for i := 1; i < len(update); i++ {
			n := &node{
				val: update[i],
			}

			var last *node
			current := head
			
			for current != nil {
				// if n.val should be before, place it before
				if rules[n.val][current.val] {
					if last == nil { // we have a new head
						n.next = current
						head = n
					} else { // insert before current
						last.next = n
						n.next = current
					}
					isValid = false
					break
				}

				last = current
				current = current.next
			}

			if current == nil {
				last.next = n
			}
		}

		validate := head
		for validate.next != nil {
			if !rules[validate.val][validate.next.val] {
				panic("invalid rule")
			}
			validate = validate.next
		}

		// find middle element of linked list with slow & fast pointer method
		slow := head
		fast := head

		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
		}

		if isValid{
			validSum += slow.val
		} else {
			invalidSum += slow.val
		}
	}

	// part 1
	fmt.Println(validSum)

	// part 2
	fmt.Println(invalidSum)
}
