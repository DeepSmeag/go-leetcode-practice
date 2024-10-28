package main

import "fmt"

// Insertion sort has O(n^2) speed and O(n) space due to the array requirements
func main() {
	arr := make([]int, 0, 10)
	arr = append(arr, 2, 6, 9, 3, 4, 1, 8, 5)
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		fmt.Println(i)
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(arr)
			}
		}
	}
	fmt.Println(arr)
}
