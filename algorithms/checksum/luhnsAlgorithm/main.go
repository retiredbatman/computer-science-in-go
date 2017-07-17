package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 0
	fmt.Println("Please enter the number for luhn check")
	fmt.Scanf("%d", &x)
	fmt.Printf("Is the number luhn check valid %v", luhnCheck(x))
}

func luhnCheck(x int) bool {
	sx := strconv.Itoa(x)
	l := len(sx)
	ax := make([]int, l)
	for i := 0; i < l; i++ {
		s := fmt.Sprintf("%s", string(sx[i]))
		ax[i], _ = strconv.Atoi(s)
	}
	fmt.Println(ax)
	for i := l - 2; i >= 0; i = i - 2 {
		n := ax[i]
		n = n * 2
		if n > 9 {
			n = n - 9
		}
		ax[i] = n
	}
	fmt.Println(ax)
	s := 0
	for i := 0; i < l; i++ {
		s = s + ax[i]
	}
	if s%10 == 0 {
		return true
	}
	return false
}
