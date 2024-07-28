package ls

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Recursive(fileInfos []fs.FileInfo, opts Option, path string) {
	for _, info := range fileInfos {
		if info.IsDir() {
			fmt.Println()
			enntryPath := filepath.Join(path, info.Name())
			dir := filepath.Dir(enntryPath)
			lastpart := filepath.Base(dir)
			fullpath := filepath.Join(lastpart, info.Name())
			fmt.Printf("./%s:\n", fullpath)
			ParseArgument(enntryPath, opts, false)
		}
	}
}
