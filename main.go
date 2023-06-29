package main

import (
	"fmt"
	"leetcode/leet"
)

func main() {

	//leet-74
	fmt.Println(leet.SearchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3))
	fmt.Println(leet.SearchMatrix([][]int{
		{1, 3},
	}, 3))

	// leet-19
	result := leet.RemoveNthFromEnd(leet.CreateLinkedList([]int{1, 2, 3, 4, 5, 6}), 2)
	for result != nil {
		fmt.Print(result.Val, "  ")
		result = result.Next
	}
}
