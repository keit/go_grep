package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go_grep <pattern> <file-name>")
		os.Exit(1)
  	}

	pattern := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	defer file.Close()

	regex := regexp.MustCompile(pattern)

	line_num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_num++
		line := scanner.Text()
		if regex.Match([]byte(line)) {
			fmt.Printf("#%d: %s\n", line_num, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
