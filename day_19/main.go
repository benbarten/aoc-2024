package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) Insert(pattern string) {
	node := t.root
	for _, r := range pattern {
		next, ok := node.children[r]
		if !ok {
			next = NewNode()
			node.children[r] = next
		}

		node = next
	}

	node.isEnd = true
}

func (t *Trie) HasPattern(pattern string) bool {
	node := t.root
	for _, r := range pattern {
		child, ok := node.children[r]
		if !ok {
			return false
		}

		node = child
	}

	return node.isEnd
}

func (t *Trie) HasPrefix(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		child, ok := node.children[r]
		if !ok {
			return false
		}

		node = child
	}

	return true
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var patterns, designs []string

	scanner.Scan()
	text := scanner.Text()
	text = strings.ReplaceAll(text, " ", "")
	patterns = strings.Split(text, ",")
	scanner.Scan()

	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}

	trie := NewTrie()

	for _, p := range patterns {
		trie.Insert(p)
	}

	var count int
	for _, design := range designs {
		if isValid(trie, design) {
			count++
		}
	}

	// part 1
	fmt.Println(count)

	count = 0
	for _, design := range designs {
		count += validOptions(trie, design)
	}

	// part 2
	fmt.Println(count)
}

func isValid(trie *Trie, design string) bool {
    dp := make([]bool, len(design)+1)
    dp[0] = true
    
    for i := 1; i <= len(design); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && trie.HasPattern(design[j:i]) {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[len(design)]
}

func validOptions(trie *Trie, design string) int {
    dp := make([]int, len(design)+1)
    dp[0] = 1
    
    for i := 1; i <= len(design); i++ {
        for j := 0; j < i; j++ {
            current := dp[j]
			if current == 0 || !trie.HasPattern(design[j:i]) {
				continue
			}
            
			dp[i] += current
        }
    }

    return dp[len(design)]
}