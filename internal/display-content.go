package ls

import (
	"fmt"
	"io/fs"
)

func DisplayContent(fileInfos []fs.FileInfo, opts Option, path string) {

	SortDisplay(fileInfos, opts)

	if opts.Long {
		Long(fileInfos, opts)
	} else if len(fileInfos) < 20 {
		for _, info := range fileInfos {
			name := info.Name()
			fmt.Printf(Colorize(name, info) + "   ")
		}
		if !opts.Long {
			fmt.Println()
		}
	} else {
		PrintColumns(fileInfos)
	}

	if opts.Recursive {
		Recursive(fileInfos, opts, path)
	}
}
