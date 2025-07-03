// You can edit this code!
// Click here and start typing.
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	size := len(old)
	num := old[size-1]
	*h = old[:size-1]

	return num
}

func GetTopK(list []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range list {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	return *h
}

func main() {
	arr := []int{100, 1, 50, 1, 200, 30, 3, 150, 4, 150}
	k := 3

	fmt.Println(GetTopK(arr, k))
}
