package main

import (
	"fmt"
)

func makeWork(work []rune) chan rune {
	out := make(chan rune)

	go func() {
		defer close(out)

		for _, w := range work {
			out <- w
		}
	}()

	return out
}

func fanOut(in chan rune) chan rune {
	out := make(chan rune)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func main() {
	work := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	in := makeWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)

	for {
		v1, o1 := <-out1
		v2, o2 := <-out2
		v3, o3 := <-out3

		if !o1 && !o2 && !o3 {
			break
		}
		if o1 {
			fmt.Printf("Output 1 got: %c\n", v1)
		}
		if o2 {
			fmt.Printf("Output 2 got: %c\n", v2)
		}
		if o3 {
			fmt.Printf("Output 3 got: %c\n", v3)
		}
	}
}
