package main

import (
	"fmt"
	"slices"
)

// Binary searching the given element; return index on found, -1 on not found
func BinarySearch(arr []int, left int, right int, val int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == val {
		return mid
	}
	if val < arr[mid] {
		// need to search to the left
		return BinarySearch(arr, left, mid-1, val)
	}
	// else we search to the right
	return BinarySearch(arr, mid+1, right, val)
}
func BinarySearchIterative(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}
	left, right := 0, len(arr)-1
	// need <= here; if we put <, then if the sought element is right in the middle at the beginning we miss entrance in the for and return -1
	for left <= right {
		mid := (left + right) / 2
		if val < arr[mid] {
			// need to search to the left
			right = mid - 1
		} else if val > arr[mid] {
			// need to search to the right
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := make([]int, 0, 10)
	arr = append(arr, 2, 6, 9, 3, 4, 1, 8, 5)
	// Binary seach on a vector/array/slice works by ensuring we have the array sorted
	slices.Sort(arr)
	fmt.Println(arr)

	const val = 6
	foundIndex := BinarySearch(arr, 0, len(arr)-1, val)
	foundIndex2 := BinarySearchIterative(arr, val)
	fmt.Printf("Found recursively %v at index %v\n", val, foundIndex)
	fmt.Printf("Found iteratively %v at index %v\n", val, foundIndex2)
	//! There's also another binary search using trees
	// This one requires that we store our data in a binary search tree, where each node has at most 2 children and the left child <= parent <= right child
	// knowing these 2 conditions, we can search our element within O(height), where height = the distance between the root of the tree and the deepest leaf
	// an imbalanced tree would have some nodes being found quickly and others not so much; that's why balanced trees are mostly used;
	// if I ever implement a BST here, it will definitely be in a separate file

}
