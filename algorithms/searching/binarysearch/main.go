package main

import "fmt"

var a = []int{
	1, 4, 5, 6, 8, 10, 11, 12, 12, 13, 14, 15, 19, 20, 20,
}

func main() {
	x := 0
	fmt.Println("Enter number to search :")
	fmt.Scanf("%d", &x)
	l := 0
	r := len(a) - 1
	i := binarySearch(a, l, r, x)
	if i == -1 {
		fmt.Printf("Could not finc number, Sorry !!!!")
	} else {
		fmt.Printf("Number found @ position : %d", i+1)
	}
}

func binarySearch(a []int, l int, r int, x int) int {
	if l <= r {
		mid := (l + r) / 2
		if a[mid] == x {
			return mid
		} else if a[mid] < x {
			return binarySearch(a, mid+1, r, x)
		} else {
			return binarySearch(a, l, mid-1, x)
		}
	}
	return -1
}
