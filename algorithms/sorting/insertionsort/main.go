package main

import "log"

var a = []int{
	10, 20, 5, 11, 8, 4, 19, 12, 12, 13, 6, 1, 20, 15, 14,
}

func main() {
	insertionSort(a)
	log.Println(a)
}

func insertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		h := i
		toBeInserted := a[i]
		for h > 0 && a[h-1] > toBeInserted {
			a[h] = a[h-1]
			h--
		}
		a[h] = toBeInserted
	}
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
