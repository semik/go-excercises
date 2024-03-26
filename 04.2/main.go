package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	// Command line arguments
	sha256Flag := flag.Bool("sha256", false, "Calculate SHA256 hash")
	sha384Flag := flag.Bool("sha384", false, "Calculate SHA384 hash")
	sha512Flag := flag.Bool("sha512", false, "Calculate SHA512 hash")
	flag.Parse()

	// Create a new hasher based on cmdline
	var hasher hash.Hash = sha256.New()
	if *sha256Flag {
		// to make compiler happy
	} else if *sha384Flag {
		hasher = sha512.New384()
	} else if *sha512Flag {
		hasher = sha512.New()
	}

	// Copy the standard input to the hasher
	_, err := io.Copy(hasher, os.Stdin)
	if err != nil {
		fmt.Println("Error reading standard input:", err)
		os.Exit(1)
	}

	fmt.Printf("Hasher is type: %T\n", hasher)

	// Sum the hash and print it
	hashSum := hasher.Sum(nil)
	fmt.Printf("%x\n", hashSum)
}
