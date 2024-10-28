package main

import "fmt"

type LinkedList struct {
	head *Node
}
type Node struct {
	next *Node
	item int
}

func (ll *LinkedList) get(i int) int {
	currentNode := ll.head
	for index := 0; index < i; index++ {
		if currentNode == nil {
			return -1
		}
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return -1
	}
	return currentNode.item
}

func (ll *LinkedList) insertHead(val int) {
	newHead := new(Node)
	newHead.item = val
	if ll.head == nil {
		ll.head = newHead
	} else {
		newHead.next = ll.head
		ll.head = newHead
	}
}

func (ll *LinkedList) insertTail(val int) {
	currentNode := ll.head
	if currentNode == nil {
		currentNode = new(Node)
		currentNode.item = val
		return
	}
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	newNode := new(Node)
	newNode.item = val
	currentNode.next = newNode
}

func (ll *LinkedList) remove(i int) bool {
	// edge cases: i = 0 means removing the head
	// i out of bounds, so we might not actually reach the right element
	// if the head doesn't actually exist, we just change the reference to the next
	// the garbage collector will get rid of the now unreferenced memory
	if ll.head == nil {
		return false
	}
	if i == 0 {
		ll.head = ll.head.next
		return true
	}
	if i == 1 {
		// the element at index 1 might be nil, need to check
		if ll.head.next == nil {
			return false
		}
		ll.head.next = ll.head.next.next
		return true
	}
	currentIndex := 0
	currentNode := ll.head
	for currentIndex < i-1 { // X _ Y
		if currentNode.next == nil {
			return false
		}
		currentIndex++
		currentNode = currentNode.next
	}
	if currentNode.next == nil {
		return false
	}
	currentNode.next = currentNode.next.next
	return true
}
func (ll *LinkedList) getValues() []int {
	values := make([]int, 0, 10)
	currentNode := ll.head
	for currentNode != nil {
		values = append(values, currentNode.item)
		currentNode = currentNode.next
	}
	return values
}

func main() {
	ll := LinkedList{}
	ll.insertHead(20)
	ll.insertHead(30)
	fmt.Println(ll.get(0))
	fmt.Println(ll.get(1))
	fmt.Println(ll.get(2))
	ll.insertTail(40)
	fmt.Println(ll.getValues())
	fmt.Println(ll.remove(1))
	fmt.Println(ll.remove(2))
	fmt.Println(ll.getValues())
	fmt.Println(ll.remove(1))
	fmt.Println(ll.getValues())
}
