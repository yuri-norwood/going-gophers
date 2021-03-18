package main

import (
	"os"

	"github.com/yuri-norwood/bfk"
)

const HelloWorld = "+[>>>->-[>->----<<<]>>]>.---.>+..+++.>>.<.>>---.<<<.+++.------.<-.>>+."

func main() {
	program := bfk.Parse(HelloWorld)
	program.Execute(os.Stdout) // Prints "hello, world!"
}
