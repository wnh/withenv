package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	filePrefix := ""
	subArgs := os.Args[:]
	expectedCount := 2

	if len(os.Args) >= 2 && strings.HasPrefix(os.Args[1], "--") {
		filePrefix = os.Args[1][2:] + "."
		subArgs = os.Args[1:]
		expectedCount += 1
	}

	if len(os.Args) < expectedCount {
		fmt.Fprintln(os.Stderr, "Error: Expecting command to run")
		os.Exit(0)
	}

	fname := fmt.Sprintf(".%senv", filePrefix)
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	scan := bufio.NewScanner(file)

	lines := 0
	for scan.Scan() {
		txt := strings.TrimSpace(scan.Text())
		//skip on empty lines
		if len(txt) == 0 || strings.HasPrefix(txt, "#") {
			continue
		}

		pair := strings.SplitN(txt, "=", 2)

		if len(pair) != 2 || pair[0] == "" || pair[1] == "" {
			fmt.Fprintf(os.Stderr, ".env:%d  - Error parsing \"%s\"", lines, txt)
			os.Exit(1)
		}
		os.Setenv(pair[0], pair[1])
		lines++
	}

	cmd := exec.Command(subArgs[1], subArgs[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
