package main

import (
	"fmt"
)

// General Scratchpad for quick ideas.

func isPrime(x int) bool {
	for i := 2; i <= x/2; i++ {
		if float64(x)/float64(i) == float64(int64(x/i)) {
			return false
		}
	}
	return true
}

func primeN(n int) int {
	pc := 0
	curPrime := 1
	for pc < n {
		curPrime++
		if isPrime(curPrime) {
			pc++
		}
	}
	return curPrime
}

func main() {
	n := 0
	for {
		fmt.Print("New prime? (Y/n): ")
		var text string
		fmt.Scanln(&text)
		if text == "n" {
			break
		} else {
			fmt.Println(primeN(n))
		}
		n++
	}
}
