package main

import (
	"fmt"

	"priority-queue/priorityqueue"
)

func main() {
	minHeap := priorityqueue.NewPriorityQueue(0, func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	minHeap.Push(5)
	minHeap.Push(2)
	minHeap.Push(8)
	minHeap.Push(1)
	minHeap.Push(3)

	fmt.Println("Min-heap pop order:")
	for {
		val := minHeap.Pop()
		if val == nil {
			break
		}
		fmt.Println(val)
	}
}
