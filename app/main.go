package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	if scanner.Err() != nil {
		return
	}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			os.Exit(1)
		} else {
			fmt.Printf("%s: command not found \n", text)
		}
		fmt.Print("$ ")
	}
}
