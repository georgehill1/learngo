package main

// General Scratchpad for quick ideas.

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		tmp := a
		a += b
		b = tmp
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
