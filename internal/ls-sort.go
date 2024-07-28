package ls

import (
	"io/fs"
	"regexp"
	"sort"
	"strings"
)

// Fonction tri par nom dans l'ordre décroissant si reverse est activé ou bien croissant quand ce n'est pas le cas.
func SortDisplay(fileInfos []fs.FileInfo, opts Option) {
	if opts.Reverse {
		sort.Slice(fileInfos, func(i int, j int) bool {
			return CleanName(fileInfos[i].Name()) > CleanName(fileInfos[j].Name())
		})
	} else {
		sort.Slice(fileInfos, func(i int, j int) bool {
			return CleanName(fileInfos[i].Name()) < CleanName(fileInfos[j].Name())
		})
	}
}

// Fonction pour que tout ce qui n'est pas des caractères alphabétiques ne soit pas pris en compte.
func CleanName(name string) string {
	lowerCaseName := strings.ToLower(name)
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	return re.ReplaceAllString(lowerCaseName, "")
}
