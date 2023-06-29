package main

import (
	"fmt"
	"leetcode/leet"
)

func main() {

	// leet-19
	result := leet.RemoveNthFromEnd(leet.CreateLinkedList([]int{1, 2, 3, 4, 5, 6}), 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
