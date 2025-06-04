package main

import (
	"fmt"
	"reflect"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1     // индекс конца "реальной части" nums1 — O(1)
	j := n - 1     // индекс конца nums2 — O(1)
	k := m + n - 1 // последний индекс итогового массива — O(1)

	// основной цикл — мы сравниваем два массива с конца
	for i >= 0 && j >= 0 { // максимум m + n итераций => O(m + n)
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i] // вставка — O(1)
			i--
		} else {
			nums1[k] = nums2[j] // вставка — O(1)
			j--
		}
		k--
	}

	// если в nums2 остались элементы, мы их просто копируем
	// этот цикл максимум n раз выполнится => O(n)
	for j >= 0 {
		nums1[k] = nums2[j] // вставка — O(1)
		j--
		k--
	}
}

func main() {
	// Тест 1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	expected := []int{1, 2, 2, 3, 5, 6}
	fmt.Println("Test 1:", reflect.DeepEqual(nums1, expected), nums1)

	// Тест 2
	nums1 = []int{4, 5, 6, 0, 0, 0}
	nums2 = []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	expected = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Test 2:", reflect.DeepEqual(nums1, expected), nums1)

	// Тест 3 — nums2 пуст
	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	expected = []int{1}
	fmt.Println("Test 3:", reflect.DeepEqual(nums1, expected), nums1)

	// Тест 4 — nums1 пуст
	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	expected = []int{1}
	fmt.Println("Test 4:", reflect.DeepEqual(nums1, expected), nums1)
}
