package main

import (
	"fmt"
	"time"
)

func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func buildHeap(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func heapSort(arr []int, n int) {
	buildHeap(arr, n)
	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, i, 0)
	}
}

func searchInHeap(arr []int, n int, x int) bool {
	l := 0
	r := n - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == x {
			return true
		} else if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func main() {
	t := time.Now()
	arr := []int{12, 11, 13, 5, 6, 7}
	n := len(arr)
	buildHeap(arr, n)
	fmt.Printf("Sorted array is %v", len(arr))
	fmt.Printf("\nSearching for x = 7 in the sorted array, %v", searchInHeap(arr, n, 7))
	fmt.Printf("\nTime taken: %v", time.Now().Sub(t))

	t = time.Now()
	// random 1000 numbers
	arr = []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, i)
	}
	n = len(arr)
	buildHeap(arr, n)

}
