package ls

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Création d'une structure fake pour pouvoir attribuer les stats du dossier acutel et parent sous la forme de . et ..
type fakeFileInfo struct {
	name string
	info fs.FileInfo
}

func (f *fakeFileInfo) Name() string       { return f.name }
func (f *fakeFileInfo) Size() int64        { return f.info.Size() }
func (f *fakeFileInfo) Mode() fs.FileMode  { return f.info.Mode() }
func (f *fakeFileInfo) ModTime() time.Time { return f.info.ModTime() }
func (f *fakeFileInfo) IsDir() bool        { return f.info.IsDir() }
func (f *fakeFileInfo) Sys() interface{}   { return f.info.Sys() }

// Fonction pour ne pas afficher les fichiers cachés si opts.All est désactivé comme par exemple .vscode
func HiddenFile(fileInfos []fs.FileInfo, opts Option, path string) []fs.FileInfo {
	if opts.All {
		currentDirInfo, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		parentDirInfo, err := os.Stat(filepath.Dir(path))
		if err != nil {
			log.Fatal(err)
		}

		fileInfos = append([]fs.FileInfo{
			&fakeFileInfo{name: ".", info: currentDirInfo},
			&fakeFileInfo{name: "..", info: parentDirInfo},
		}, fileInfos...)
	} else {
		var filteredFileInfos []fs.FileInfo
		for _, info := range fileInfos {
			if info.Name()[0] != '.' {
				filteredFileInfos = append(filteredFileInfos, info)
			}
		}
		return filteredFileInfos
	}
	return fileInfos
}
