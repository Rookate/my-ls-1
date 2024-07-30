package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

/* Initialization de la variable root à true au lancement du programme pour
que si la fonction recursive est active alors qu'il ne passe pas dans "if len(opts.Filnames)"*/

func ParseArgument(path string, opts Option, root bool) {
	if opts.List {
		entries, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Error reading directory path")
			log.Fatal(err)
			return
		}

		var fileInfos []fs.FileInfo
		if len(opts.Filenames) > 0 && root {
			for _, filenames := range opts.Filenames {
				entrypath := filepath.Join(path, filenames)
				info, err := os.Lstat(entrypath)
				// os.Lstat pour ne pas suivre les fichiers cibles ce qui nous permet de savoir si un fichier est un lien symbolique ou non
				if err != nil {
					fmt.Printf("Error %s does not exist or cannot be accessed\n", filenames)
					continue
				}

				if info.IsDir() {
					if opts.Recursive && opts.Root {
						fmt.Println(".:")
					} else {
						fmt.Printf("./%s:\n", filenames)
					}
					ParseArgument(entrypath, opts, false)
				} else {
					fileInfos = append(fileInfos, info)
					fileInfos = HiddenFile(fileInfos, opts, path)
				}
				DisplayContent(fileInfos, opts, path)
			}
		} else {
			if opts.Recursive && root {
				fmt.Println(".:")
			}

			for _, entry := range entries {
				info, err := entry.Info()
				if err != nil {
					fmt.Printf("Error getting info: %s %s", info, info.Name())
				}
				fileInfos = append(fileInfos, info)
			}
			fileInfos = HiddenFile(fileInfos, opts, path)
			DisplayContent(fileInfos, opts, path)
		}
	}
}
