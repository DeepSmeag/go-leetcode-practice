package main

import "fmt"

/*
QuickSort is intended to sort in-place by continuously choosing a pivot
and we ensure that entries to the left of the pivot are less than it, those to right are greater
and we recursively call the same algorithm for the left and right sides
*/
func QuickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	// choosing pivot as the middle of our available slice
	mid := (left + right) / 2
	fmt.Printf("Pivot is %d\n", arr[mid])
	// moving the pivot to the left
	arr[left], arr[mid] = arr[mid], arr[left]
	var i, j, d = left, right, 0
	for i < j {
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			d = 1 - d
		}
		i += d
		j -= 1 - d
	}

	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

func main() {
	arr := make([]int, 0, 10)
	arr = append(arr, 2, 6, 9, 3, 4, 1, 8, 5)
	arr2 := make([]int, len(arr), 10)
	copy(arr2, arr)
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	QuickSortRe(arr2, 0, len(arr2)-1)
	fmt.Println(arr)
	fmt.Println(arr2)
}

/*
Step by step:
2 6 9 3 4 1 8 5
Choosing pivot 3, moving to the left:
3 6 9 2 4 1 8 5; now sorting with i=left, j=right
3 < 5 so we let it be, j--
3 < 8 so we let it be, j--
3 > 1 so we swap them & d=1 (d is a marker for where to increase); i++
1 6 9 2 4 3 8 5
6 > 3 so we swap them, d=0 so we'll decrease to the right
1 3 9 2 4 6 8 5
3 < 4 so we let it be
3 > 2 so we swap them & d = 1 (decreasing from the right), i++
1 2 9 3 4 6 8 5
9 > 3 so we swap them & d = 0, j--
1 2 3 9 4 6 8 5 and now i points to 3, j points to 3 (i==j)
now that i = j we know that all that's to the left of i is < pivot, all that's to the right is >pivot
So we can recursively call the same function

The tough part with quick sort imo is to find the efficient way of swapping values around
*/

// implementing this again without looking at any resource to practice and also explain along the way
func QuickSortRe(arr []int, left int, right int) {
	// the main thing of QuickSort is that is does O(nlogn) sorting in-place
	// otherwise we might use trees or maps to sort our stuff, or heaps or whatever
	if left >= right {
		return
	}
	// we choose a pivot; there are a few options: left-most, right-most, middle, median, random; I don't like the idea of calculating the median since it's extra work; I like middle / random
	mid := (left + right) / 2 // this is our pivot's index
	// pivot = an element chosen to divide the array into lower elements and higher elements; everything to the left of the pivot must be lower than it, everything to the right must be higher
	// first we move the pivot to the left
	arr[left], arr[mid] = arr[mid], arr[left]
	var leftIndex, rightIndex, increaseLeft int = left, right, 0
	// we use these indexes to go through the array and swap elements so that we reach our desired state;
	// the algorithm is ping-pong like; we start with our pivot on the left and compare leftIndex(the pivot's position for now) with the element at rightIndex
	// as long as arr[leftIndex] <= arr[rightIndex], we know the pivot has elements higher than it to the right; as long as this happens, rightIndex-- cause we now they're good
	// when we reach an arr[leftIndex] > arr[rightIndex], we've got a smaller element to the right; by swapping them, now our pivot is at the rightmost place where we still know for sure that everything to the right of it is higher; now we also change increaseLeft to 1 so we know that moving forward we're checking for elements at arr[leftIndex] who are higher than our pivot, and our pivot is at arr[rightIndex];
	// the whole idea is to iterate only once through the array, from both directions
	for leftIndex < rightIndex {
		if arr[leftIndex] > arr[rightIndex] {
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex] // swap
			increaseLeft = 1 - increaseLeft                                   // toggle "direction"
		}
		// the idea of having this toggle of "increaseLeft" is to avoid doing extra if checks for the direction; it's more efficient than having if-else where we increase/decrease our indexes in both the if and else clauses, but differently; though it's harder to get it at first
		leftIndex += increaseLeft
		rightIndex -= 1 - increaseLeft
	}
	// when we're here, leftIndex==rightIndex and they're pivot's position
	// now we also sort in the same way the regions to the left and to the right by recursively calling, Divide et Impera-style
	QuickSortRe(arr, left, leftIndex-1)
	QuickSortRe(arr, leftIndex+1, right)
}
