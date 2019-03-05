package main

import (
	flag "github.com/ogier/pflag"
)

var fizzBuzzNumber int

func main() {
	flag.Parse()
}

func init() {
	flag.IntVarP(&fizzBuzzNumber, "number", "n", 0, "Is it Fizz or Buzz !!")
}

// $ go get github.com/ogier/pflag

// $ go run fizzBuzz.go --help
