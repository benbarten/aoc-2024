package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	input := []int{0, 7, 6618216, 26481, 885, 42, 202642, 8791}

	// part 1
	blinks := 25

	ll := list.New()
	for _, item := range input {
		ll.PushBack(item)
	}

	for blink := 0; blink < blinks; blink++ {
		el := ll.Front()
		for el != nil {
			val := el.Value.(int)
			if val == 0 {
				el.Value = 1
				el = el.Next()
				continue
			}

			str := strconv.Itoa(val)
			if len(str)%2 == 0 {
				leftStr := str[:len(str)/2]
				rightStr := str[len(str)/2:]

				left, _ := strconv.Atoi(leftStr)
				right, _ := strconv.Atoi(rightStr)

				el.Value = left
				ll.InsertAfter(right, el)
				el = el.Next().Next()
				continue
			}

			el.Value = el.Value.(int) * 2024
			el = el.Next()
		}
	}

	fmt.Println(ll.Len())

	// part 2: the linked list does not scale and we need to leverage caching
	var sum int
	for _, num := range input {
		sum += sumCached(num, 75)
	}
	fmt.Println(sum)
}

type blinkNum struct {
	num, blinks int
}

var cache = map[blinkNum]int{}

func sumCached(num int, blinks int) int {
	if val, ok := cache[blinkNum{num, blinks}]; ok {
		return val
	}

	if blinks == 0 {
		return 1
	}

	if num == 0 {
		return sumCached(1, blinks-1)
	} else if str := strconv.Itoa(num); len(str)%2 == 0 {
		left, _ := strconv.Atoi(str[:len(str)/2])
		right, _ := strconv.Atoi(str[len(str)/2:])

		sum := sumCached(left, blinks-1) + sumCached(right, blinks-1)
		cache[blinkNum{num, blinks}] = sum
		return sum
	} else {
		sum := sumCached(num*2024, blinks-1)
		cache[blinkNum{num, blinks}] = sum
		return sum
	}
}
