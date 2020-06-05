package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	var mySeed, err = strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		fmt.Printf("OhNo ParseInt Error!")
		os.Exit(1)
	}
	rand.Seed(mySeed)
	fmt.Println("My favorite number is", rand.Intn(10))
}
