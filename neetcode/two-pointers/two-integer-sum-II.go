package main

import "fmt"

// given an array of non-decreasing ints (which means x<=y for any x that comes before y in the array) and a target int; return indices (1-indexed, meaning our array is considered to start at 1) of the 2 numbers that add up to target
// the problem ensures that there's exactly 1 solution every time
// requirement: O(1) space

func TwoIntegerSumII(arr []int, target int) (int, int) {
	// solution: given that we now the array is non-decreasing, we know for sure that checking numbers to the right increases our sum, and to the left decreases it; having the category of this problem also hints to the solution
	// we use 2 pointers, starting at the beginning and end of the array; these starting numbers have a sum; if the sum is < target, we know we have to increase the sum (so we move the left pointer to the right); if the sum is too big, we have to decrease the sum so we move the right pointer to the left; this guarantees we're iterating towards a good solution
	// since we also have the guarantee of there being a solution, we wouldn't need to worry about other constraints, but having left<right should still apply for the for condition
	// another note: if the solution was not guaranteed to be unique, this algorithm would still work by choosing the first solution it found (with the indices as far apart as possible)
	// anyway, O(n) time and O(1) space
	left := 0
	right := len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return left + 1, right + 1
		}
	}
	return -1, -1
}

func main() {
	nums := []int{1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9}
	nums2 := []int{1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	target2 := 20
	fmt.Println(TwoIntegerSumII(nums, target))
	fmt.Println(TwoIntegerSumII(nums2, target2))
}
