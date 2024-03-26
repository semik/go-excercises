// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for line, files := range counts {
		count := 0
		var fileNames []string
		for file, n := range files {
			count += n
			fileNames = append(fileNames, file)
		}
		if count > 1 {
			fmt.Printf("%d;\t%s;\t%s\n", count, line,
				strings.Join(fileNames, ", "))
		}
	}

}

func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][fileName]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
