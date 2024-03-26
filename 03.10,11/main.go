// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buff bytes.Buffer

	// 400001.002
	// calculate position of . from end
	decimalDot := strings.IndexByte(s, '.')
	if decimalDot > 0 {
		decimalDot = len(s) - decimalDot
	}

	// process string from end and every 3 characters insert comma
	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		buff.WriteByte(s[i])

		if decimalDot > 0 {
			// we are still procesing part after dot
			decimalDot--
		} else {
			if s[i] >= '0' && s[i] <= '9' {
				c++
				if (c%3 == 0) && (i != 0) && !((s[i-1] == '+') || (s[i-1] == '-')) {
					buff.WriteByte(',')
				}
			}
		}
	}

	/* reverse order of string
	   I failed to find library function to reverse order of string or slice */
	reversed := buff.Bytes()
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}

//!-
