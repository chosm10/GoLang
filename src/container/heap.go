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
	result := h[i] < h[j]
	fmt.Println("LESS: ", h[i], h[j], result)
	return result
}

func (h MinHeap) Swap(i, j int) {
	fmt.Println("SWAP: ", h[i], h[j])
	h[i], h[j] = h[j], h[i]
	fmt.Println("Now: ", h)
}

func (h *MinHeap) Push(a interface{}) {
	fmt.Println("PUSH: ", a)
	*h = append(*h, a.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	length := len(old)
	result := old[length-1]
	*h = old[0 : length-1]
	return result
}

func main() {
	arr := new(MinHeap)
	heap.Init(arr)
	heap.Push(arr, 3)
	heap.Push(arr, 2)
	heap.Push(arr, 5)
	heap.Push(arr, 7)
	heap.Pop(arr)

	fmt.Println("MIN: ", (*arr)[0])
	fmt.Println("MAX: ", (*arr)[len(*arr)-1])
}
