package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		repl()
	case 1:
		runFile(args[0])
	default:
		fmt.Println("Usage: glox [file]")
		os.Exit(1)
	}

}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		if ok := scanner.Scan(); !ok {
			break
		}
		line := scanner.Text()
		run(line)
	}
}

func runFile(inputFile string) {}

func run(string) {
	fmt.Println("Doing thing")
}
