package Queue

import (
	"fmt"
	"strconv"
)

// Linked List to get O(1) complexity

type Node struct {
	Data float64
	next *Node
}

var rear, front *Node

func isEmpty() bool {
	return rear == nil && front == nil
}

func Enqueue(data float64) {
	elem := Node{Data: data, next: nil}

	if isEmpty() {
		front = &elem
		rear = &elem
		return
	}

	rear.next = &elem
	rear = &elem
}

func Dequeue() (elem Node) {
	elem = *front

	if isEmpty() {
		return
	}
	if front == rear {
		front = nil
		rear = nil
	} else {
		front = front.next
	}

	return
}

func Length() int {
	temp := front
	count := 1
	for {
		if temp == rear {
			break
		}
		temp = temp.next
		count++
	}
	return count
}

func PrintQueue() {
	if isEmpty() {
		fmt.Println("Nothing to print!")
		return
	}

	temp := front
	count := 1
	for {
		fmt.Println("Number "+strconv.Itoa(count)+":", temp.Data)

		if temp == rear {
			break
		}
		temp = temp.next
		count++
	}
}
