package main

import (
	"datastructures"
)

func main() {
	queue := new(datastructures.Queue)

	queue.Enqueue(3)
	queue.Enqueue(15)
	queue.Enqueue(7)
	queue.Enqueue(9)
	queue.Enqueue(0)

	queue.Print()

}
