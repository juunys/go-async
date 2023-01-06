package main

import "fmt"

func generateRandInt() chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	ch := generateRandInt()

	for {
		n, open := <-ch
		if !open {
			break
		}
		fmt.Println(n)
	}
}
