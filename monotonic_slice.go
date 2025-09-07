// monotonic_slice.go
package main

import (
	"fmt"
	"reflect"
)

//
// A) Проверить, что слайс монотонен неубывающе (non-decreasing).
//    func IsMonotonicNonDecreasing(nums []int) bool
//
// B) Проверить, что слайс монотонен (либо неубывающий, либо невозрастающий).
//    func IsMonotonicEither(nums []int) bool
//
// C) Найти длину максимального монотонного (нестрого) подотрезка в любом направлении.
//    func LongestMonotonicRunEither(nums []int) int
//

func IsMonotonicNonDecreasing(nums []int) bool {
	for i := 0; i+1 < len(nums); i++ {
		n1 := nums[i]
		n2 := nums[i+1]

		if n1 > n2 {
			return false
		}
	}

	return true
}

func IsMonotonicEither(nums []int) bool {
	isDecr, isIncr := false, false

	for i := 0; i+1 < len(nums); i++ {
		n1 := nums[i]
		n2 := nums[i+1]

		if n1 < n2 {
			isIncr = true
		}
		if n1 > n2 {
			isDecr = true
		}
		if isIncr && isDecr {
			return false
		}
	}

	return true
}

func LongestMonotonicRunEither(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	maxLen := 1
	incrCounter := 1
	decrCounter := 1
	for i := 1; i < len(nums); i++ {
		n1 := nums[i-1]
		n2 := nums[i]

		if n1 < n2 {
			incrCounter++
			decrCounter = 1
		} else if n1 > n2 {
			decrCounter++
			incrCounter = 1
		} else {
			incrCounter++
			decrCounter++
		}

		if incrCounter > maxLen {
			maxLen = incrCounter
		}
		if decrCounter > maxLen {
			maxLen = decrCounter
		}
	}

	return maxLen
}

func mustEqual(got, want interface{}, msg string) {
	if !reflect.DeepEqual(got, want) {
		panic(fmt.Sprintf("%s: got=%v want=%v", msg, got, want))
	}
}

func main() {
	// Тесты для A (non-decreasing)
	func() {
		type TC struct {
			in   []int
			want bool
		}
		cases := []TC{
			{[]int{}, true},
			{[]int{1}, true},
			{[]int{1, 2, 2, 3}, true},
			{[]int{1, 1, 1}, true},
			{[]int{3, 2, 2, 1}, false},
			{[]int{1, 3, 2}, false},
		}
		for i, tc := range cases {
			got := IsMonotonicNonDecreasing(tc.in)
			mustEqual(got, tc.want, fmt.Sprintf("A[%d]", i))
			_ = i
			_ = tc
		}
	}()

	// Тесты для B (either direction)
	func() {
		type TC struct {
			in   []int
			want bool
		}
		cases := []TC{
			{[]int{}, true},
			{[]int{1}, true},
			{[]int{1, 2, 2, 3}, true}, // неубывающий
			{[]int{5, 4, 4, 1}, true}, // невозрастающий
			{[]int{1, 3, 2}, false},   // меняет направление
			{[]int{3, 3, 3}, true},    // плоский — ок
		}
		for i, tc := range cases {
			got := IsMonotonicEither(tc.in)
			mustEqual(got, tc.want, fmt.Sprintf("B[%d]", i))
			_ = i
			_ = tc
		}
	}()

	// Тесты для C (longest run either direction)
	func() {
		type TC struct {
			in   []int
			want int
		}
		cases := []TC{
			{[]int{}, 0},
			{[]int{7}, 1},
			{[]int{1, 2, 2, 1, 0}, 4},       // [1,2,2] длина 3; [2,1,0] не монотонно; лучший: [2,2,1] нет; подберёшь свой пример
			{[]int{5, 4, 4, 4, 3, 2}, 6},    // весь массив невозрастающий
			{[]int{1, 3, 5, 4, 6, 8, 7}, 3}, // несколько участков
			{[]int{2, 2, 1, 2, 2}, 3},
		}
		for i, tc := range cases {
			got := LongestMonotonicRunEither(tc.in)
			mustEqual(got, tc.want, fmt.Sprintf("C[%d]", i))
			_ = i
			_ = tc
		}
	}()

	fmt.Println("Всё ок")
}
