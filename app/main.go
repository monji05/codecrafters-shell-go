package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")

	fmt.Print(text)
	fmt.Print(": command not found")
}
