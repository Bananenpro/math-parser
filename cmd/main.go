package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	input := readInput("Enter an expression: ")
	tokens, err := Scan(input)
	if err != nil {
		fmt.Println("Fatal error while scanning:", err)
		os.Exit(1)
	}

	fmt.Println("Tokens:", tokens)

	ast, err := Parse(tokens)
	if err != nil {
		fmt.Println("Fatal error while parsing:", err)
		os.Exit(1)
	}

	result, err := Evaluate(ast)
	if err != nil {
		fmt.Println("Fatal error while evaluating:", err)
		os.Exit(1)
	}

	fmt.Printf("The result is %v.\n", result)
}
