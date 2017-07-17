package main

import "log"

var a = []int{
	10, 20, 11, 5, 8, 4, 19, 12, 12, 13, 6, 1, 20, 15, 14,
}

func main() {
	selectionSort(a)
	log.Println(a)
}

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := a[i]
		s1 := i
		s2 := 0
		for j := i + 1; j < len(a); j++ {
			if min > a[j] {
				min = a[j]
				s2 = j
			}
		}
		if s2 != 0 {
			swap(a, s1, s2)
		}
	}
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
