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
			} else if _, err := exec.LookPath(args[0]); err == nil {
				execCommand := exec.Command(cmd, args...)
				execCommand.Stdout = os.Stdout
				execCommand.Stderr = os.Stderr
				execCommand.Run()
			} else {
				fmt.Printf("%s: not found \n", args[0])
			}
		} else {
			fmt.Printf("%s: command not found", cmd)
		}
	}
}
