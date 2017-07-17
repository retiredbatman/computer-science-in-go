package main

import "fmt"

var a = []int{
	10, 20, 11, 5, 8, 4, 19, 12, 12, 13, 6, 1, 20, 15, 14,
}

func main() {
	high := len(a) - 1
	low := 0
	quickSort(a, low, high)
	fmt.Printf("quick sorted array : %v", a)
}

func quickSort(a []int, low int, high int) {
	if low < high {
		pi := partition(a, low, high)

		quickSort(a, low, pi-1)
		quickSort(a, pi+1, high)
	}
}

func partition(a []int, low int, high int) int {
	pivot := a[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if a[j] < pivot {
			i++
			swap(a, i, j)
		}
	}
	swap(a, i+1, high)
	return i + 1
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
