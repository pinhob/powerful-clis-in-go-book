package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	wordCount := count(os.Stdin)
	fmt.Println(wordCount)
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
