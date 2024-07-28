package ls

import (
	"io/fs"
	"sort"
)

func SortDisplay(fileInfos []fs.FileInfo, opts Option) {
	if opts.Reverse {
		sort.Slice(fileInfos, func(i int, j int) bool {
			return fileInfos[i].Name() > fileInfos[j].Name()
		})
	} else {
		sort.Slice(fileInfos, func(i int, j int) bool {
			return fileInfos[i].Name() < fileInfos[j].Name()
		})
	}
	if opts.Time {
		sort.Slice(fileInfos, func(i int, j int) bool {
			return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
		})
	}
}
