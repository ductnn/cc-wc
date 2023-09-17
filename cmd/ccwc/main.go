package main

import (
	"fmt"
	"io/ioutil"

	ccwc "github.com/ductnn/cc-wc/internal"
)

func main() {
	// wc – word, line, character, and byte count.
	opts := ccwc.ParseOptions()

	if opts.Version {
		fmt.Printf("ccwc %s", ccwc.Version)
		return
	}

	if opts.FilePath == "" {
		fmt.Println("Please provide a file path using the -f flag.")
		fmt.Printf(ccwc.Usage)
		return
	}

	fileData, err := ioutil.ReadFile(opts.FilePath)
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	c := &ccwc.Counter{}

	if opts.Bytes {
		byteCount := c.CountBytesInData(fileData)
		fmt.Printf("Byte count: %d\n", byteCount)
	}

	if opts.Chars {
		characterCount := c.CountCharsInData(fileData)
		fmt.Printf("Character count: %d\n", characterCount)
	}

	if opts.Words {
		wordCount := c.CountWordsInData(fileData)
		fmt.Printf("Word count: %d\n", wordCount)
	}

	if opts.Lines {
		lineCount := c.CountLinesInData(fileData)
		fmt.Printf("Line count: %d\n", lineCount)
	}
}
