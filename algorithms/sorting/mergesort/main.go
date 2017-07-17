package main

import "log"

var a = []int{
	10, 20, 11, 5, 8, 4, 19, 12, 12, 13, 6, 1, 20, 15, 14,
}

var b []int

func main() {
	b := mergeSort(a)
	log.Printf("final : %v", b)
}

func mergeSort(a []int) []int {
	return mSort(a)
}

func mSort(a []int) []int {
	log.Printf("len a : %d", len(a))
	if len(a) <= 1 {
		return a
	}
	low := 0
	high := len(a)
	mid := int(low + high/2)
	log.Printf("mid : %d", mid)
	a1 := mSort(a[:mid])
	a2 := mSort(a[mid:])
	log.Printf("a1 after m sort : %v", a1)
	log.Printf("a2 after m sort : %v", a2)
	return merge(a1, a2)
	//return a
}

func merge(a1 []int, a2 []int) []int {
	log.Printf("a1 merge : %v", a1)
	log.Printf("a2 merge : %v", a2)
	tl := len(a1) + len(a2)
	b := make([]int, tl)
	log.Println(tl)
	i, j := 0, 0
	t := i
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			b[t] = a1[i]
			i++
			t++
		} else {
			b[t] = a2[j]
			j++
			t++
		}
	}
	for i < len(a1) {
		b[t] = a1[i]
		t++
		i++
	}
	for j < len(a2) {
		b[t] = a2[j]
		t++
		j++
	}
	return b
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
