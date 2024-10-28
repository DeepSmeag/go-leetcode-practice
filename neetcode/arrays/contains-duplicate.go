package main

import "fmt"

// Given an array (slice) , return true if any value appears more than once
func ContainsDuplicates(arr []int) bool {
	// we use a map to store newly-met values; if at any point we try to store a value that's already stored, we know we have a duplicate
	// O(n) time with O(n) space
	// the tradeoff here is space; we could also do O(1) space by quicksorting, but then we have O(nlogn) time
	// edge case - empty, in which case there's nothing to iterate over so we return false; it's covered already
	numMap := make(map[int]bool)
	for _, num := range arr {
		if _, exists := numMap[num]; exists {
			return true
		}
		numMap[num] = true
	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicates([]int{1, 4, 6, 4, 3, 6, 7, 8, 9}))
	fmt.Println(ContainsDuplicates([]int{1, 4, 6, 3, 7, 8, 9}))
}
