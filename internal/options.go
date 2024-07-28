package ls

import (
	"log"
	"os"
	"strings"
)

type Option struct {
	List      bool
	All       bool
	Long      bool
	Reverse   bool
	Recursive bool
	Time      bool
	Path      string
	Filenames []string
	Root      bool
	FileCount int
}

func ParseOptions(args []string) Option {
	var opts Option
	var isArg bool
	var key string

	for _, arg := range args {
		if arg == "ls" {
			opts.List = true
			continue
		}

		isArg, key, _ = IsArgument(arg)
		if isArg {
			if len(key) == 1 {
				switch key[0] {
				case 'a':
					opts.All = true
				case 'l':
					opts.Long = true
				case 'r':
					opts.Reverse = true
				case 'R':
					opts.Recursive = true
				case 't':
					opts.Time = true
				}
			} else {
				switch key {
				case "all":
					opts.All = true
				case "reverse":
					opts.Reverse = true
				case "recursive":
					opts.Recursive = true
				default:
					for _, letter := range strings.Split(key, "") {
						switch letter[0] {
						case 'a':
							opts.All = true
						case 'l':
							opts.Long = true
						case 'r':
							opts.Reverse = true
						case 'R':
							opts.Recursive = true
						case 't':
							opts.Time = true
						}
					}
				}
			}
		} else {
			if opts.Path == "" {
				if strings.HasPrefix(arg, "/") || strings.Contains(arg, ":") {
					opts.Path = arg
				} else {
					opts.Filenames = append(opts.Filenames, arg)
				}
			} else {
				opts.Filenames = append(opts.Filenames, arg)
			}
		}
	}
	if opts.Path == "" {
		cmd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		opts.Path = cmd
	}

	return opts
}

/*
Checks whether the function's parameter is an argument, to use as terminal commands' option.

The argument may be a single "-" with a single character, or "--" with a single word, either with a value ("=" and the value after it) or not.
*/
func IsArgument(s string) (bool, string, string) {
	if len(s) < 2 || s[0] != '-' {
		return false, "", ""
	}
	startArg, startVal := 1, -1
	arg, value := "", ""

	if s[startArg] == '-' || s == "ls" {
		startArg++
	}

	for i := startArg; i < len(s); i++ {
		if s[i] == '=' {
			if i != startArg {
				arg = s[startArg:i]
				startVal = i + 1
				break
			} else {
				return false, "", ""
			}
		}
	}

	if startVal != -1 {
		value = s[startVal:]
	} else {
		arg = s[startArg:]
	}

	return true, arg, value
}
