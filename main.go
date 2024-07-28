package main

import (
	ls "ls/internal"
	"os"
)

func main() {
	args := os.Args[1:]
	opts := ls.ParseOptions(args)
	ls.ParseArgument(opts.Path, opts, true)
}
