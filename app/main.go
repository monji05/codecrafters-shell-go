package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var builtInCommands = []string{"echo", "exit", "type"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		}

		command = strings.TrimSpace(command)
		tokens := strings.Split(command, " ")

		cmd, args := tokens[0], tokens[1:]

		if cmd == "exit" {
			break
		} else if cmd == "echo" {
			fmt.Printf("%s \n", command[5:])
		} else if cmd == "type" {
			if slices.Contains(builtInCommands, args[0]) {
				fmt.Printf("%s is a shell builtin \n", args[0])
			} else if path, err := exec.LookPath(args[0]); err == nil {
				fmt.Printf("%s is %s \n", args[0], path)
			} else {
				fmt.Printf("%s: not found \n", args[0])
			}
		} else {
			fmt.Printf("Program was passed %d args (including program name).\n", len(tokens))
			fmt.Printf("Arg #0 (program name): %s\n", cmd)
			for index, arg := range args {
				fmt.Printf("Arg #%d: %s\n", index+1, arg)
			}
		}
	}
}
