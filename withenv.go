package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)

	lines := 0
	for scan.Scan() {
		txt := scan.Text()
		pair := strings.SplitN(txt, "=", 2)

		if len(pair) != 2 || pair[0] == "" || pair[1] == "" {
			log.Fatal(fmt.Sprintf(".env:%d  - Error parsing \"%s\"", lines, txt))
		}
		os.Setenv(pair[0], pair[1])
		lines++
	}

	if len(os.Args) < 2 {
		log.Fatal("Error: Expecting command to run")
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
