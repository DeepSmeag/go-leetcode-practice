package main

import (
	"fmt"
	"slices"
)

// given array, return all triplets where the sum of the 3 numbers = 0; the indices of the triplets should be distinct
// also no duplicate triplets (this gives a hint to solution); output can be in any order
func ThreeIntegerSum(arr []int) [][]int {
	// solution: two sum I and II help here; sort so we have a non-decreasing order, that's O(nlogn);
	// we're searching for an a,b,c triplet such that a+b+c=0; we iterate through the elements-2 for a and have it fixed; then we apply two pointer searching to find b and c
	// we could literally call TwoIntegerSumII to help us
	sortedArr := append(make([]int, 0, len(arr)), arr...)
	slices.Sort(sortedArr)
	res := make([][]int, 0, len(arr)/5) // just a random heuristic for the capacity
	for i := 0; i < len(sortedArr)-2; i++ {
		// this is our fixed a
		val := sortedArr[i]
		if val > 0 {
			// if our smallest value >0, we're not hitting 0-sum for sure
			break
		}
		if i > 0 && val == sortedArr[i-1] {
			continue // same thing as with increasing left below, we don't want duplicate solutions
		}
		left := i + 1
		right := len(sortedArr) - 1
		for left < right {
			sum := val + sortedArr[left] + sortedArr[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{val, sortedArr[left], sortedArr[right]})
				left++
				right--
				// we don't break cause there might be other solutions we could miss
				// and to make sure we don't get duplicate solutions, at least 1 variable needs to change; we could go for left or right here, but left is the intuitive choice
				for sortedArr[left] == sortedArr[left-1] && left < right {
					left++
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	numsEmpty := []int{0, 1, 1}
	numsZero := []int{0, 0, 0}
	fmt.Println(ThreeIntegerSum(nums))
	fmt.Println(ThreeIntegerSum(numsEmpty))
	fmt.Println(ThreeIntegerSum(numsZero))
}

//! ARCHIVE (bad thinking but keeping it anyway for reference)
//solution thinking: so there's this hint that triplets need to be distinct; so example: -1,1,0,2,-1,1; technically, we have 3 different triplet types
// -1,1,0; -1,-1,2 but we also have 0,-1,1 (which is the same triplet as the first one); we can't put the same one, even though they do have different indices
// an option would be to first get rid of duplicates, but we introduce unnecessary steps; we don't need to make another list, we just need to have unique numbers; this reminds me of sets / maps in this case
// and given this number is in the two pointer category, we know we need 2 pointers; if we have 2 numbers, the 3rd is given by target - first 2;
// so I need a way to iterate linearly through the array and get distinct ways of reaching the sum 0
// twopointer left-right with searching for the complement in the map gives me whether it's possible; I still need to check whether the element in the map has different indices compared to the left and right
// if the array is sorted, then it's easier cause I have the guarantee I'm searching through all the options; that would guarantee me an O(nlogn) solution
// in this case, I'd have left-right with min and max; and then I search in the map for something in the middle; and the map has key the number & values=[]int indices where it can be found
//_ ... _ and I search in the map; if not found, I know ...do I?
// ... I realize there's so me wrong thinking above; I'm keeping this as a history of my thinking, even though it's wrong; two-pointer on the array by itself, sorted or not, still forces me to do O(n^2) if I want all the options

// 2nd try
// solution thinking: we have the hint of no duplicates, which means sets or maps; we also know we've got to do something related to two pointers because of the category; if we didn't know the category, the guess would've come from the idea of searching 2 things; why two things? a+b = target -c; so c is the complement of a+b, as long as we have 2 numbers we know the third; when we search for 2 things at once we know it has to do with two pointers
// in this case searching in the array itself isn't useful, sorted or not we'd still have to check every pair possible
// so I'm thinking we use a map to extract unique numbers (<1000) and their frequency ( so we know if we can use the same number more times)
// then we can iterate through that and extract an array with them so we can use two pointers on those unique numbers; if we sort those, we know lowest is on the left, highest is on the right;
// we also know that each triplet changes if one number changes; so if we only work with unique numbers in our search..nvm a [0 0 0] solution would not be findable this way

//try 3
// solution thinking: first solving two sum and two sum II cause they help
// the problem is essentially "find a,b,c such that a+b+c=0"
// which we can transform into "find a,b such that their sum = -c"
// which simplifies the problem greatly now
// if we sort our array, making time O(nlogn) overall, we know left=lowest, right=highest
// if we have repeating numbers, as long as one of the variables changes the overall triplet changes; exception is only 0, where it can fit with multiple other pairs;
// if we have a sorted array, we have the guarantee that we can cover pairs of two pointers left-right
// let's say we have the first scenario _ ... _ (lowest, highest and something in between); we can do some checks here; 1st: does c which completes the sum exist? if yes, then we have a trio and we put it there; if not, then we still need to move on; but which direction? consider this: a-lowest, b-highest, -c=target (call it d to make the logic easier); so a+b=d; if d < a (meaning it has to be to the left), then we know for sure we need higher sums to have a chance to find an existing d (meaning we have to move a to the left); if d>b (meaning it's to the right of b), then we know we have to decrease our sum to be able to find d; if a<d<b and we still don't find it, then how do we move the pointers?

// UPDATE: I looked into the proposed efficient solution; here I was trying to get something better than O(n^2) and apparently that's the optimum; I was looking for an O(nlogn), which means a linear parsing to generate every option after the sort;
