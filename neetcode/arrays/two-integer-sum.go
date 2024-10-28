package main

import "fmt"

// Given an array of ints and a target(int), return the two indices i,j so that arr[i]+arr[j]==target; it's guaranteed that there is exacly one pair of indices that satisfy the condition
func TwoIntegerSum(arr []int, target int) (int, int) {
	// solution: we use a map to store the numbers(key) and their index(value); during our array iteration, we search for the number's complement in the map; if it exists, we have found our indices; if not, store our number and move on
	// O(n) speed since we iterate through the array; O(n) space since we store numbers in the map
	numbers := make(map[int]int)
	for i, num := range arr {
		complement := target - num
		if index, exists := numbers[complement]; exists {
			return index, i
		}
		numbers[num] = i
	}
	// we'll never reach this point due to the nature of the problem, but the compiler doesn't know this
	return -1, -1
}

func main() {
	arr := []int{1, 5, 7, 8, 3, 5}
	target := 10
	fmt.Println(TwoIntegerSum(arr, target))
}
