// sort_by_frequency_unique.go
package main

import (
	"fmt"
	"reflect"
	"sort"
)

// SortByFrequencyUnique:
// Возвращает список УНИКАЛЬНЫХ значений, отсортированных:
// 1) по частоте по возрастанию (freq ASC),
// 2) при равной частоте — по значению по убыванию (value DESC).
//
// Пример: [1,1,2,2,2,3] -> [3,1,2]
func SortByFrequencyUnique(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	freqList := map[int]int{}
	sl := []int{}

	for _, num := range nums {
		if _, ok := freqList[num]; !ok {
			sl = append(sl, num)
		}
		freqList[num]++
	}

	sort.Slice(sl, func(i, j int) bool {
		if freqList[sl[i]] != freqList[sl[j]] {
			return freqList[sl[i]] < freqList[sl[j]]
		}

		return sl[i] > sl[j]
	})

	return sl
}

func mustEqual(got, want interface{}, msg string) {
	if !reflect.DeepEqual(got, want) {
		panic(fmt.Sprintf("%s: got=%v want=%v", msg, got, want))
	}
}

func main() {
	type TC struct {
		in   []int
		want []int
	}
	cases := []TC{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		// базовые
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 2}}, // 1×3, 2×3, 1×2 -> 3,1,2
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 2}},    // 1×1, 2×2, 3×2 -> 1,3,2 (tie: 3>2)
		// с отрицательными и нулём
		{[]int{-1, -1, -2, -2, -2, 3}, []int{3, -1, -2}}, // 3×1, -1×2, -2×3
		{[]int{0, 0, 1, 1}, []int{1, 0}},                 // freq 2 и 2 -> value DESC: 1, 0
		// равные частоты у нескольких значений
		{[]int{4, 5, 6, 5, 4, 3}, []int{6, 3, 5, 4}}, // freq(6)=1,freq(3)=1 -> 6,3; freq(5)=2,freq(4)=2 -> 5,4
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},       // все по 1 -> по value DESC
	}

	for i, tc := range cases {
		got := SortByFrequencyUnique(tc.in)
		mustEqual(got, tc.want, fmt.Sprintf("FREQ-UNIQUE[%d]", i))
	}

	fmt.Println("Всё ок")
}
