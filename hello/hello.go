package main

import (
	"fmt"
	"rsc.io/quote"
)

func Quote() string {
	return quote.Go()
}

func main() {
	fmt.Println(Quote())
}
