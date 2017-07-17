package main

import (
	"log"
)

var a = []int{
	10, 20, 5, 11, 8, 4, 19, 12, 12, 13, 6, 1, 20, 15, 14,
}

func main() {
	bubbleSort(a)
	log.Println(a)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			}
		}
	}
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
