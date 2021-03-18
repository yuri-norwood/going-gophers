package main

import (
	"os"

	"github.com/yuri-norwood/bfk"
)

// HelloWorld is a brainf*** program that prints "hello, world!"
const HelloWorld = "+[>>>->-[>->----<<<]>>]>.---.>+..+++.>>.<.>>---.<<<.+++.------.<-.>>+."

func main() {
	program := bfk.Parse(HelloWorld)
	program.Execute(os.Stdout)
}
