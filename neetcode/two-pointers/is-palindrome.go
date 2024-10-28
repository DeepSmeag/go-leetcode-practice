package main

import (
	"fmt"
)

// check if given string is palindrome; palindrome = reads the same front-to-back and back-to-front
func IsPalindrome(str string) bool {
	// solution: it's a classic 2 pointer problem, we start with 1 pointer at the start of the string and 1 at the end, and converge towards the middle while checking for same letters
	// O(n) speed, we iterate through half the word but still check the entire word -1 letter possibly, depending on how many characters it has
	// O(1) space, only variables
	// also...need to take into account Go's implementation details; if the string has unicode characters, we can't access them directly by indexing; so we first need to convert to rune slice; this might create 2 iterations, still O(n); don't know underlying mechanism Go uses to convert to rune array
	letters := []rune(str)
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	strNo := "abacada"
	strYes := "abacaba"
	fmt.Println(IsPalindrome(strNo))
	fmt.Println(IsPalindrome(strYes))
}
