package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capLetter := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	fmt.Printf("Before: %c\n", capitalized)
	for i := 0; i < len(data); i++ {
		go capLetter(data[i])
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("After: %c\n", capitalized)
}
