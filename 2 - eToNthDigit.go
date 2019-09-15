package main

// General Scratchpad for quick ideas.

import "fmt"
import "math"

func eToNthDigit(n int) string {
	fmtString := "%." + fmt.Sprintf("%v", n+1) + "v"
	return fmt.Sprintf(fmtString, math.E)
}

func main() {
	fmt.Println(eToNthDigit(100))
}
