package main

import "fmt"

func rotateRight(a []int, n int) []int {
	if n == 0 {
		// does Go have exception for errors
	} else if n >= len(a) {
		// again error
	} else {
		aa := make([]int, len(a))
		copy(aa[n:], a)
		copy(aa[:], a[len(a)-n:])
		return aa
	}

	return a
}

func main() {
	aOrig := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for i := 1; i <= 4; i++ {
		a := aOrig

		fmt.Println(a)
		a = rotateRight(a, i)
		fmt.Println(a)
		fmt.Println()
	}

}
