package main

import "fmt"

// function that returns true if 2 given words are anagrams of one another; false otherwise
// anagram = contains same letters, but different order
func IsAnagram(s1 string, s2 string) bool {
	// solution: use a vector with length of english alphabet ('z'-'a') where we store occurences of letters in s1; then we go through s2 and decrease the occurences; if at any point we reach <0, return false; at the end, another iteration through the letter vector; if anything != 0, return false
	// O(n) speed (n = max(len(s1),len(s2))), O(1) space (since array[26] is fixed)
	letters := ['z' - 'a']int{0}
	for _, l := range s1 {
		letters[l-'a']++
	}
	for _, l := range s2 {
		if letters[l-'a'] == 0 {
			return false
		}
		letters[l-'a']--
	}
	for _, l := range letters {
		if l != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1, s2 := "racecar", "carrace"
	fmt.Println(IsAnagram(s1, s2))
	s1, s2 = "car", "jar"
	fmt.Println(IsAnagram(s1, s2))
}
