package main

// General Scratchpad for quick ideas.

import "fmt"
import "math"

func piToNthDigit(n int) string {
	fmtString := "%." + fmt.Sprintf("%v", n+1) + "v"
	return fmt.Sprintf(fmtString, math.Pi)
}

func main() {
	fmt.Println(piToNthDigit(2))
}
