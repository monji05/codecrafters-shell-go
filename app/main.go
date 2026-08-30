package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var builtInCommands = []string{"echo", "exit", "type"}

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
			break
		} else if strings.HasPrefix(text, "echo") {
			fmt.Println(text[5:])
		} else if strings.HasPrefix(text, "type") {
			callType(text)
		} else {
			fmt.Printf("%s: command not found \n", text)
		}
		fmt.Print("$ ")
	}
}

func callType(text string) {
	arg := text[4:]
	arg = strings.TrimSpace(arg)
	if slices.Index(builtInCommands, arg) != -1 {
		fmt.Printf("%s is a shell built in \n", text[4:])
	} else {
		fmt.Printf("%s: not found \n", arg)
	}
}
