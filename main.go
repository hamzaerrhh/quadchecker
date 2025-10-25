package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// Read all data coming from stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Print what was received
	fmt.Print(string(data))

	// take all and get the line and cols then store in x and y
	// after u execute on all after u compare them one by one
	// and also add a condition at top to check if it is a quad function or not (not all line or not column)

	cmd := exec.Command("./quadC", "1", "1")

	// Run it and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output as a string
	fmt.Print(string(output))
}
