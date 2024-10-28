package main

import "fmt"

// given an array of numbers and a number ok, find top k most frequent numbers in array
// we have the guarantee that there is only 1 answer (so no situation in which top k but the i-th frequency has many numbers so we would go past k); so in a sense we have a guaranteed uniqueness of number frequency
func TopKfrequent(arr []int, k int) []int {
	// solution: we use a frequency map on first iteration to know how many times each number appears; why map? so we don't waste memory if numbers can be in range [-5000,5000] then we'd need a 10k vector; with maps, it's the bare minimum
	// now we could sort  the values in this map based on frequency, but that would land us in the realm of O(nlogn) speed; we can do O(n)
	// we take an array of len(arr) elements; each index signifies the number of appearances the values have; for index 2, the value is a list of the numbers in arr which have appeared 2 times
	// this takes linear time to build and linear space as well; if we then iterate through this array from end to beginning, we can get our top k numbers
	// edge case, len(arr) < k case in which the algorithm still covers our situation; other edge cases when len>=k but there's no k distinct numbers, so we cannot output them
	// This solution is O(n) speed and O(n) space; we do 3 iterations through our arrays and take some memory, but speed is often prioritized cause there's enough space generally
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
	countSlice := make([][]int, len(arr), len(arr))
	for key, value := range freq {
		countSlice[value] = append(countSlice[value], key)
	}
	res := make([]int, 0, k)
	for i := len(countSlice) - 1; i > -1; i-- {
		if len(countSlice[i]) > 0 {
			res = append(res, countSlice[i]...)
			if len(res) >= k {
				break
			}
		}
	}
	return res

}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4} // this input isn't valid per problem specification, but it highlights the situation in which we have multiple freqeuencies for the same number
	nums2 := []int{1, 2, 2}
	k := 3
	fmt.Println(TopKfrequent(nums, k))
	fmt.Println(TopKfrequent(nums2, k))

}
