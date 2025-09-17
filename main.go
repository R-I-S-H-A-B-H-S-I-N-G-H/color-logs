package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// all time best time : 0.30s user 0.02s system 103% cpu 0.304 total
// cur best time      : 0.34s user 0.01s system 103% cpu 0.341 total

var blueBox = color.New(color.BgBlue, color.FgWhite).SprintFunc()
var redBox = color.New(color.BgRed, color.FgWhite).SprintFunc()
var greenBox = color.New(color.BgGreen, color.FgWhite).SprintFunc()
var yellowBox = color.New(color.BgYellow, color.FgWhite).SprintFunc()

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Create a buffered writer for stdout
	writer := bufio.NewWriter(os.Stdout)
	// CRITICAL: Ensure the buffer is flushed before the program exits.
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		highlighted := highlightLog(line)
		fmt.Fprintln(writer, highlighted)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func highlightLog(line string) string {
	// Use an if/else if chain to short-circuit
	if strings.Contains(line, "ERROR") {
		return strings.ReplaceAll(line, "ERROR", redBox(" ERROR "))
	}

	if strings.Contains(line, "WARNING") { // Important: Check for "WARNING" before "WARN"
		return strings.ReplaceAll(line, "WARNING", yellowBox(" WARNING "))
	}

	if strings.Contains(line, "WARN") {
		return strings.ReplaceAll(line, "WARN", yellowBox(" WARN "))
	}

	if strings.Contains(line, "INFO") {
		return strings.ReplaceAll(line, "INFO", greenBox(" INFO "))
	}

	if strings.Contains(line, "DEBUG") {
		return strings.ReplaceAll(line, "DEBUG", blueBox(" DEBUG "))
	}

	// If no keywords are found, return the original line
	return line
}
