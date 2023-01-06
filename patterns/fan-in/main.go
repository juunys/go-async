package main

import (
	"fmt"
	"sync"
)

func makeWork(work []rune) chan rune {
	ch := make(chan rune)
	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}

func fanIn(inputs ...chan rune) chan rune {
	out := make(chan rune)

	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan rune) {
			for {
				value, open := <-ch
				if !open {
					defer wg.Done()
					break
				}
				out <- value
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	w1 := makeWork([]rune{'a', 'c', 'e', 'g'})
	w2 := makeWork([]rune{'b', 'd', 'f', 'h'})

	out := fanIn(w1, w2)

	for value := range out {
		fmt.Printf("Value --> %c\n", value)
	}
}
