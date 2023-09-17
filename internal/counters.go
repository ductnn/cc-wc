package ccwc

import (
	"unicode"
)

func CountWordsInData(data []byte) int {
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

func CountCharsInData(data []byte) int {
	return len(data)
}

func CountLinesInData(data []byte) int {
	lineCount := 0

	for _, b := range data {
		if b == '\n' {
			lineCount++
		}
	}

	return lineCount
}
