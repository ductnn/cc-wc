package ccwc

import (
	"unicode"
)

type Counter struct {
	countWords bool
	countChars bool
	countLines bool
}

func (c *Counter) CountWordsInData(data []byte) int {
	inWord := false
	wordCount := 0

	for _, b := range data {
		if unicode.IsSpace(rune(b)) {
			inWord = false
		} else {
			if !inWord {
				wordCount++
			}
			inWord = true
		}
	}

	return wordCount
}

func (c *Counter) CountCharsInData(data []byte) int {
	return len(data)
}

func (c *Counter) CountLinesInData(data []byte) int {
	lineCount := 0

	for _, b := range data {
		if b == '\n' {
			lineCount++
		}
	}

	return lineCount
}
