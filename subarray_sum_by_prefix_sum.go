package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixSum := map[int]int{0: 1}

	for _, num := range nums {
		sum += num

		if prefixSum[sum-k] > 0 {
			count += prefixSum[sum-k]
		}

		prefixSum[sum]++ //0:1, 1:1, 2:1
	}

	return count
}

func main() {
	tests := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1}, 2},                 // 2
		{[]int{1, 2, 3}, 3},                 // 2
		{[]int{3, 4, 7, 2, -3, 1, 4, 2}, 7}, // 4
		{[]int{1}, 0},                       // 0
	}

	for _, test := range tests {
		fmt.Printf("Input: %v, k=%d → Output: %d\n", test.nums, test.k, subarraySum(test.nums, test.k))
	}
}
