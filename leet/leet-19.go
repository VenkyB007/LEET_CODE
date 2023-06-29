package leet

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := head
	second := dummy

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next

}

func (l *SinglyLinkedList) Add(value int) {
	newNode := &ListNode{
		Val:  value,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		current := l.Tail
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		l.Tail = newNode
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type SinglyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func CreateLinkedList(input []int) *ListNode {
	linkedList := SinglyLinkedList{
		Head: nil,
		Tail: nil,
	}

	for _, i := range input {
		linkedList.Add(i)
	}

	fmt.Println("successfully created linkedlist")
	return linkedList.Head
}
