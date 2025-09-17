package main

import (
	"bufio"
	"fmt"
	"io"
	"log" // <-- Import the log package
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/hpcloud/tail"
)

var (
	blueBox   = color.New(color.BgBlue, color.FgWhite).SprintFunc()
	redBox    = color.New(color.BgRed, color.FgWhite).SprintFunc()
	greenBox  = color.New(color.BgGreen, color.FgWhite).SprintFunc()
	yellowBox = color.New(color.BgYellow, color.FgWhite).SprintFunc()
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		processFile(filename)
	} else {
		processStdin()
	}
}

func processFile(filename string) {
	config := tail.Config{
		Follow:   true,
		ReOpen:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd},
		Logger: log.New(io.Discard, "", 0),
	}
	t, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		return
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for line := range t.Lines {
		highlighted := highlightLog(line.Text)
		fmt.Fprintln(writer, highlighted)
		writer.Flush()
	}
}

func processStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		highlighted := highlightLog(line)
		fmt.Fprintln(writer, highlighted)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
	}
}

func highlightLog(line string) string {
	if strings.Contains(line, "ERROR") {
		return strings.ReplaceAll(line, "ERROR", redBox(" ERROR "))
	} else if strings.Contains(line, "WARNING") {
		return strings.ReplaceAll(line, "WARNING", yellowBox(" WARNING "))
	} else if strings.Contains(line, "WARN") {
		return strings.ReplaceAll(line, "WARN", yellowBox(" WARN "))
	} else if strings.Contains(line, "INFO") {
		return strings.ReplaceAll(line, "INFO", greenBox(" INFO "))
	} else if strings.Contains(line, "DEBUG") {
		return strings.ReplaceAll(line, "DEBUG", blueBox(" DEBUG "))
	}
	return line
}
