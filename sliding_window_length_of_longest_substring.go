package main

import "fmt"

// speed O(<= 2n)
// memory O(1), O(min(n,256))
func lengthOfLongestSubstring(s string) int {
	start, maxLen := 0, 0
	lastSeen := map[byte]int{}

	for end := 0; end < len(s); end++ {
		ch := s[end]

		if _, ok := lastSeen[ch]; ok {
			start = max(start, lastSeen[ch]+1)
		}

		lastSeen[ch] = end
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3  ("abc")
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1  ("b")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3  ("wke")
	fmt.Println(lengthOfLongestSubstring(""))         // 0
	fmt.Println(lengthOfLongestSubstring("au"))       // 2
}
