package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	set := make(map[byte]bool)
	left := 0

	for right := 0; right < len(s); right++ {

		for set[s[right]] {
			delete(set, s[left])
			left++
		}

		set[s[right]] = true

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func main() {
	inputs := []string{
		"abcabcbb", // ожидаем 3
		"bbbbb",    // ожидаем 1
		"pwwkew",   // ожидаем 3
		"",         // ожидаем 0
		"a",        // ожидаем 1
		"abcddefg", // ожидаем 4 ("defg")
	}

	for _, s := range inputs {
		fmt.Printf("Input: %q → Output: %d\n", s, lengthOfLongestSubstring(s))
	}
}
