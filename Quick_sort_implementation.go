
// 1.quicksort

package main

import (
	"fmt"
)

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}
	
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

$go run main.go
[1 2 3 4 10 14 16 18 20 23 24 40 41 43 46 51 53 55 56 57 59 62 68 73 74 78 79 84 87 88 89 90 93 95 99]

