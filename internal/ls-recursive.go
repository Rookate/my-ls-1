package ls

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Recursive(fileInfos []fs.FileInfo, opts Option, path string) {
	for _, info := range fileInfos {
		// Condition pour éviter que la récursive s'applique sur les repos . et .. car sinon ça fait une boucle à l'infinie

		if info.IsDir() {
			fmt.Println()
			entryPath := filepath.Join(path, info.Name())
			dir := entryPath[len(opts.Path+"/"):]
			fmt.Printf("./%s:\n", dir)
			ParseArgument(entryPath, opts, false)
		}
	}
}
