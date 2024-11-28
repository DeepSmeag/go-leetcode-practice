package main

import "fmt"

/*
*
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.
*/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// this is a problem that fits with compiled languages; Go/C/C++/Java/Rust should be the targets
	// as potential solutions, we could do multiple passes in order to avoid having a bigger space complexity
	// I'd take the bigger space - smaller time complexity approach

	// then we can iterate through the list normally; the issue is with assigning the random pointers when they point to a node we haven't yet created
	// space - O(n); time O(n); more exactly I'd say 3n space and 2n time (2 iterations)
	copyHead := Node{}
	copyNode := &copyHead
	ogNode := head
	currentIndex := 0
	// we'll have 2 maps:
	// indexToAddress - bind index of node to the address of that node; goal is to find index of the Random which we must point to
	// valToIndex - bind value of a node to its index
	// on the forward pass
	valToIndex := make(map[int]int)
	indexToAddress := make(map[int]*Node)
	for ogNode != nil {
		copyNode.Val = ogNode.Val
		if ogNode.Next != nil {
			newCopy := Node{}
			copyNode.Next = &newCopy
		} else {
			copyNode.Next = nil
		}
		valToIndex[copyNode.Val] = currentIndex
		indexToAddress[currentIndex] = copyNode

		ogNode = ogNode.Next
		copyNode = copyNode.Next
		currentIndex++
	}
	// at this point we have the list, now we need to build the randoms
	// we iterate with the original at the same time so we can use the Val of the random the nodes point to in order to find the address of our copy list
	copyNode = &copyHead
	ogNode = head
	for copyNode != nil {
		if ogNode.Random != nil {
			copyNode.Random = indexToAddress[valToIndex[ogNode.Random.Val]]
		} else {
			copyNode.Random = nil
		}
		copyNode = copyNode.Next
		ogNode = ogNode.Next
	}
	return &copyHead
	// as an extra note; I've seen a more in-place option were we double elements in the list and then clean it up; in the end that uses extra memory as well; I'm not sure if it essentially doubles the list or not; with this approach, we technically triple it, since we have 2 maps which contain information for each node;
	// that solution is fit for situations where we don't have much memory; most modern problems seek to optimize time, not space
}

func main() {
	// we get the list input as an array technically, that should be converted into a proper list with pointers; that function is not the object of this problem so I'll do the lists by hand
	node1 := Node{7, nil, nil}
	node2 := Node{13, nil, nil}
	node3 := Node{11, nil, nil}
	node4 := Node{10, nil, nil}
	node5 := Node{1, nil, nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = nil
	node1.Random = nil
	node2.Random = &node1
	node3.Random = &node5
	node4.Random = &node3
	node5.Random = &node1

	copyHead := copyRandomList(&node1)
	printList(&node1)
	fmt.Println("")
	printList(copyHead) // should show the same thing as the initial list

}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("Node adr: %p | val: %d | next: %p | rand: %p\n", head, head.Val, head.Next, head.Random)
		head = head.Next
	}
}
