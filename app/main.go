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
		fmt.Printf("%s is a shell builtin \n", text[5:])
	} else {
		path, _ := exec.LookPath(arg)
		fileInfo, _ := os.Stat(path)
		permission := fileInfo.Mode().Perm().String()
		_, _, isExecutable := strings.Cut(permission, "x")
		if isExecutable {
			fmt.Printf("%s is %s \n", arg, path)
			return
		}
		fmt.Printf("%s: not found \n", arg)
	}
}
