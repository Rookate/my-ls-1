package ls

import (
	"io/fs"
	"sort"
)

func SortTime(fileInfos []fs.FileInfo, opts Option) {
	var sorter func(i int, j int) bool

	if opts.Time {
		if opts.Reverse {
			sorter = func(i int, j int) bool {
				return fileInfos[i].ModTime().Before((fileInfos[j].ModTime()))
			}
		} else {
			sorter = func(i int, j int) bool {
				return fileInfos[i].ModTime().After((fileInfos[j].ModTime()))
			}
		}
	}

	if sorter != nil {
		sort.Slice(fileInfos, sorter)

	}
}
