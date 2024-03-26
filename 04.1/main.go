// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"fmt"
)

func diffBits(a, b uint8) int {
	//fmt.Printf("   %3d %8b\n", a, a)
	//fmt.Printf("   %3d %8b\n", b, b)

	diff := 0
	s := ""
	for i := 0; i < 8; i++ {
		var mask uint8 = 1 << i
		//fmt.Printf("%d: %3d %8b\n", i, mask, mask)
		if a&mask != b&mask {
			s = "!" + s
			diff++
		} else {
			s = "=" + s
		}
	}
	//fmt.Printf("       %s\ndiff=%d\n\n", s, diff)

	return diff
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	db := 0
	for i := 0; i < len(c1); i++ {
		db += diffBits(c1[i], c2[i])
	}
	fmt.Printf("Pocet ruznych bitu: %d\n", db)
}

//!-
