// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

}

// !+rev
// reverse reverses a array of 6 ints in place.
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
