package main

import (
	"fmt"
	"math"
)

func makeWork(input []int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, in := range input {
			out <- in
		}
	}()

	return out
}

func filter(in chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()

	return out
}

func square(in chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2)
			out <- int(value)
		}
	}()

	return out
}

func multiply(in chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := i * 2
			out <- value
		}
	}()

	return out
}

func main() {
	in := makeWork([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	filtered := filter(in)
	squared := square(filtered)
	multiplied := multiply(squared)

	for v := range multiplied {
		fmt.Println(v)
	}
}
