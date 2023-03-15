package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text := "Hello world! This is a test. I'm a text value."
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Println("Word count: ", count)
}
