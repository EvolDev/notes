// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func sortByFrequency(nums []int) []int {
	freq := make(map[int]int)
	res := []int{}
	for _, n := range nums {
		if _, ok := freq[n]; !ok {
			res = append(res, n)
		}
		freq[n]++
	}

	sort.Slice(res, func(i, j int) bool {
		if freq[res[i]] != freq[res[j]] {
			return freq[res[i]] > freq[res[j]]
		}

		return res[i] < res[j]
	})

	return res
}

func main() {

	nums := []int{4, 3, 1, 6, 3, 4, 3, 6, 1, 4}
	fmt.Println(sortByFrequency(nums))

}
