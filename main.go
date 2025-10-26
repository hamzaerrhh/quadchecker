package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}
	input := string(data)
	cols, line := CalcLinesAndCols(input)
	Checker(input, line, cols)
}

func CalcLinesAndCols(s string) (cols, lines int) {
	if len(s) == 0 {
		printNotQuad()
	}

	lines = 0
	cols = 0
	currentCol := 0
	firstLine := true

	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			// End of line
			if firstLine {
				cols = currentCol
				if cols == 0 {
					printNotQuad()
				}
				firstLine = false
			} else {
				if currentCol != cols {
					printNotQuad()
				}
			}
			currentCol = 0
			lines++
		} else {
			currentCol++
		}
	}

	// Handle case if last line doesn't end with '\n'
	if currentCol > 0 {
		if firstLine {
			cols = currentCol
		} else if currentCol != cols {
			printNotQuad()
		}
		lines++
	}

	return cols, lines
}

// helper to print error and exit
func printNotQuad() {
	os.Stdout.WriteString("Not a quad function\n")
	os.Exit(0)
}

func Checker(str string, line, cols int) {
	quads := []string{"A", "B", "C", "D", "E"}
	var matches []string

	for _, q := range quads {
		cmd := exec.Command("./quad"+q, fmt.Sprint(cols), fmt.Sprint(line))
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			continue // skip if program missing or error
		}

		output := out.String()

		// compare ignoring possible trailing newline differences
		if strings.TrimSpace(output) == strings.TrimSpace(str) {
			matches = append(matches, fmt.Sprintf("[quad%s] [%d] [%d]", q, cols, line))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(matches, " || "))
}
