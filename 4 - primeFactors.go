package main

import "fmt"

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
func divisible(a int, b int) bool {
	if float64(a)/float64(b) == float64(int(a/b)) {
		return true
	}
	return false
}

func factorise(n int) {
	curPrime := 0
	i := 1
	for curPrime <= n {
		curPrime = primeN(i)

		if divisible(n, curPrime) {
			n /= curPrime
			fmt.Println(curPrime)
			i = 0
		}

		i++
	}
}

func main() {
	factorise(6004050)
}
