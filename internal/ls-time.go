package ls

import (
	"io/fs"
	"sort"
)

func SortTime(fileInfos []fs.FileInfo, opts Option) {
	if opts.Time {
		sort.Slice(fileInfos, func(i int, j int) bool {
			return fileInfos[i].ModTime().After((fileInfos[j].ModTime()))
		})
	}
}
