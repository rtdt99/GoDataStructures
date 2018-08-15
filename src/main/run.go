package main

import (
	"datastructures"
	"fmt"
)

func main() {
	myLinkedList := &datastructures.LinkedList{}

	myLinkedList.AddLast(12)
	myLinkedList.AddLast(9)
	myLinkedList.AddLast(2)
	myLinkedList.AddLast(11)
	myLinkedList.AddLast(9)
	myLinkedList.Print()

	fmt.Println()
	myLinkedList.Delete(7)
	myLinkedList.Delete(9)
	myLinkedList.Print()
}
