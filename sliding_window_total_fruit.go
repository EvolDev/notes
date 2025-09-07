package main

import "fmt"

func totalFruit(fruits []int) int {
	left, maxLen := 0, 0
	basket := map[int]int{}

	for right := 0; right < len(fruits); right++ {
		basket[fruits[right]]++

		for len(basket) > 2 {
			basket[fruits[left]]--

			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}

			left++
		}

		maxLen = max(maxLen, right-left+1)

	}

	return maxLen

}

func main() {
	fmt.Println(totalFruit([]int{1, 2, 1}))                         // 3
	fmt.Println(totalFruit([]int{3, 3, 2, 1, 2}))                   // 3
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))                      // 3
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))                   // 4
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4})) // 5
}
