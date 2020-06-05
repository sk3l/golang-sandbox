package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var numstr = os.Args[1]

	var numint = 0
	var numpow = 0

	for i := len(numstr) - 1; i >= 0; i-- {

		//fmt.Println("i = %d", i)
		var digit = numstr[i] - 0x30

		if i == 0 && numint > 0 && numstr[i] == '-' {
			numint = 0 - numint
		} else if numstr[i] < 0x30 || numstr[i] > 0x39 {
			fmt.Printf("Encountered non-digit char %c\n", numstr[i])
			os.Exit(1)
		} else {
			numint += int(digit) * int(math.Pow10(numpow))
			numpow++
		}
	}

	fmt.Println(numint)
	os.Exit(0)
}
