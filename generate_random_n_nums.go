package main

import (
	"fmt"
	"math/rand"
	"time"
)

// speed O(N)
// mem O(N)
func uniqRandn(n int) []int {
	N := n * 10
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i
	}

	rand.Shuffle(N, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	return nums[:n]
}

// best speed O(n)
// worst speed O(n) collisions
// mem O(n)
func uniqRandnWithSeed(n int) []int {
	rand.Seed(time.Now().UnixNano())

	m := make(map[int]struct{})
	res := make([]int, 0, n)

	for len(res) < n {
		num := rand.Intn(n * 10)
		if _, exists := m[num]; !exists {
			m[num] = struct{}{}
			res = append(res, num)
		}
	}

	//или fmt.Println(rand.Perm(n * 10)[:n])
	return res
}

func main() {
	fmt.Println(uniqRandn(10))
	fmt.Println(uniqRandnWithSeed(10))
}
