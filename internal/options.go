package ccwc

import (
	"flag"
	"fmt"
)

const Version = "v0.1.0"

const Usage = `
ccwc ` + Version + `
A Go word count (wc) clone - print newline, word, char, and byte counts for each file 

USAGE:
	gowc [FLAGS] [OPTIONS] [input]...

FLAGS:
	-c, --bytes              Print the byte counts.
	-m, --chars              Print the character counts.
	-l, --lines              Print the newline counts.
	-w, --words              Print the word counts.
	-h, --help               Display help and exit.
	-V, --version            Output version information and exit.

OPTIONS:
	-f <file>     Read input from the newline-terminated list of filenames in the given file.
ARGS:
	<input>...    Input file names
`

type Options struct {
	Words    bool
	Lines    bool
	Chars    bool
	Bytes    bool
	Version  bool
	FilePath string
}

func ParseOptions() *Options {
	opts := new(Options)
	flag.BoolVar(&opts.Words, "w", false, "Count words")
	flag.BoolVar(&opts.Lines, "l", false, "Count lines")
	flag.BoolVar(&opts.Chars, "m", false, "Count chars")
	flag.BoolVar(&opts.Bytes, "c", false, "Count bytes")
	flag.BoolVar(&opts.Version, "v", false, "print version")
	flag.StringVar(&opts.FilePath, "f", "", "Path to the input file")

	flag.Usage = func() { fmt.Print(Usage) }

	flag.Parse()

	if !opts.Lines && !opts.Words && !opts.Bytes && !opts.Chars {
		opts.Lines = true
		opts.Words = true
		opts.Bytes = true
	}

	return opts
}
