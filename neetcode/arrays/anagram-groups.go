package main

import (
	"fmt"
)

// given a list of strings, group all anagrams together in sublists; any output order is fine
// anagram = two strings contain the same characters, but in different order
func AnagramGroups(strs []string) [][]string {
	// solution 1: sort each string/word and then sort these values in the string array; we anagrams are now right next to each other, though we no longer have a correspondence to the initial value, so we need to keep track of it via a struct that holds the initial index so we can go back to the right value
	// speed: O(nlogn) * num strings (this one can be the higher one, calling it m)
	// so speed: O(m * nlogn) and space O(m*n) since we need to store the sorted strings

	//solution 2, a bit more efficient: the same approach from IsAnagram, but more words; we do a frequency vector for the letters of each string; that's speed O(m*n); if we transform that into a string, we can use it as a key in a map which allows us O(1) inserts and searches
	// we add the words as values in a []string slice; their key is the stringified frequency vector
	// space O(m) due to the map having all words as keys / values
	// the real space is like O(2m)=O(m) since I need to copy the map values into a slice; I don't know a why to do it in-place or 1-line at least; I see there's a maps.Values function, but it returns an iterator so we still don't do it in 1 line
	// we also need to transform the map into a slice of slices of strings
	wordMap := make(map[string][]string)
	for _, word := range strs {
		freq := [26]int{0}
		for _, letter := range word {
			freq[letter-'a']++

		}
		// turning frq into string using sprintf
		wordMap[fmt.Sprint(freq)] = append(wordMap[fmt.Sprint(freq)], word)
	}
	res := make([][]string, 0, len(strs))
	for _, val := range wordMap {
		res = append(res, val)
	}
	return res
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(AnagramGroups(strs))
}
