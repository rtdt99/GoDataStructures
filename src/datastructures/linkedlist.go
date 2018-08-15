package datastructures

import "fmt"

type LinkedList struct {
	head *LinkedListNode
}

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

func (ll *LinkedList) AddLast(val int) {
	if ll.head == nil {
		ll.head = &LinkedListNode{val, nil}
	} else {
		temp := ll.head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &LinkedListNode{val, nil}
	}
}

func (ll *LinkedList) Delete(val int) {
	if ll.head == nil {
		return
	}

	temp := ll.head
	for temp != nil && temp.Next != nil {
		if temp.Val == val {
			nextNode := temp.Next
			temp.Val = nextNode.Val
			temp.Next = nextNode.Next
		}
		temp = temp.Next
	}

	if temp.Val == val {
		h := ll.head
		for h.Next != temp {
			h = h.Next
		}
		h.Next = nil
	}

}

func (ll *LinkedList) Print() {
	if ll.head == nil {
		return
	}
	temp := ll.head

	for temp != nil {
		fmt.Printf("%d\t", temp.Val)
		temp = temp.Next
	}
}
