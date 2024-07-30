package ls

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Recursive(fileInfos []fs.FileInfo, opts Option, path string) {
	for _, info := range fileInfos {
		// Les dossiers parent et actuel ne sont pas à traiter
		if info.Name() == "." || info.Name() == ".." {
			continue
		}

		if info.IsDir() {
			fmt.Println()
			entryPath := filepath.Join(path, info.Name())
			dir := entryPath[len(opts.Path+"/"):]
			fmt.Printf("./%s:\n", dir)
			ParseArgument(entryPath, opts, false)
		}
	}
}
