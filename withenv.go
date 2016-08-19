package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := os.Open(".env")
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

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Expecting command to run")
		os.Exit(0)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
