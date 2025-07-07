package main

import "fmt"

func twoSumSorted(arr []int, target int) (int, int, bool) {
	i := 0
	j := len(arr) - 1

	for i < j {
		sum := arr[i] + arr[j]
		if sum == target {
			return arr[i], arr[j], true
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return -1, -1, false
}
func main() {
	arr := []int{1, 2, 3, 4, 6}
	target := 6

	fmt.Println(twoSumSorted(arr, target))
}
