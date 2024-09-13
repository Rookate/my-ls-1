package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

// Fonction pour le flag -l pour afficher plus de détail
func Long(fileInfos []fs.FileInfo, opts Option) {
	if len(fileInfos) != 0 {
		totalBlock := 0
		for _, blocksCount := range fileInfos {
			stat, ok := blocksCount.Sys().(*syscall.Stat_t)
			if !ok {
				log.Fatal("Failed to get raw syscall.Stat_t data")
			}
			totalBlock += int(stat.Blocks)
		}

		fmt.Printf("total: %d\n", totalBlock)
	}

	for _, info := range fileInfos {
		var link string
		if (info.Mode() & os.ModeSymlink) != 0 { // Le fichier est-il un lien symbolique ?
			target, err := filepath.EvalSymlinks(filepath.Join(opts.Path, info.Name()))
			if err != nil {
				fmt.Printf("Error resolving symlink: %v\n", err)
				link = " -> <error>"
			} else {
				linkInfo, errLink := os.Lstat(target)
				if errLink != nil {
					fmt.Println("Error getting info: " + errLink.Error())
				}

				normalizedTarget := filepath.Clean(target)
				normalizedPath := filepath.Clean(opts.Path)

				if normalizedPath == normalizedTarget {
					target = "."
				} else if filepath.Dir(normalizedTarget) == normalizedPath {
					normalizedTarget = filepath.Base(normalizedTarget)
				} else if target[:len(filepath.Dir(normalizedPath))] == filepath.Dir(normalizedPath) && normalizedTarget[0] != '/' {
					normalizedTarget = "../" + normalizedTarget[len(filepath.Dir(normalizedPath)+"/"):]
				}
				link = " -> " + Colorize(normalizedTarget, linkInfo)
			}
		}
		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			log.Fatal("Failed to get raw syscall.Stat_t data")
		}

		uid := fmt.Sprint(stat.Uid)
		gid := fmt.Sprint(stat.Gid)

		usr, err := user.LookupId(uid)
		if err != nil {
			fmt.Printf("Error looking up user: %v\n", err)
			usr = &user.User{Name: "unknown"}
		}

		grp, err := user.LookupGroupId(gid)
		if err != nil {
			fmt.Printf("Error looking up group: %v\n", err)
			grp = &user.Group{Name: "unknown"}
		}
		name := Colorize(info.Name(), info)

		fmt.Printf("%s %3d %8v %8v %5d %5v %s%s\n",
			getInfoModes(info),
			stat.Nlink,
			usr.Name,
			grp.Name,
			info.Size(),
			info.ModTime().Format("Jan 02 15:04"),
			name,
			link,
		)
	}
}

func getInfoModes(info fs.FileInfo) string {
	modes := make([]rune, 10, 11)
	for i := 0; i < len(modes); i++ {
		modes[i] = '-'
	}

	if info.IsDir() {
		modes[0] = 'd'
	} else if (info.Mode() & os.ModeSymlink) != 0 {
		modes[0] = 'l'
	} else if (info.Mode() & os.ModeSocket) != 0 {
		modes[0] = 's'
	} else if (info.Mode() & os.ModeNamedPipe) != 0 {
		modes[0] = 'p'
	} else if (info.Mode() & os.ModeDevice) != 0 {
		if (info.Mode() & os.ModeCharDevice) != 0 {
			modes[0] = 'c'
		} else {
			modes[0] = 'b'
		}
	} else {
		modes[0] = '-'
	}

	if (info.Mode() & os.ModeSetuid) != 0 {
		modes[3] = 's'
	}
	if (info.Mode() & os.ModeSetgid) != 0 {
		modes[6] = 's'
	}

	if (info.Mode() & os.ModeSticky) != 0 {
		modes[9] = 't'
	}

	perms := int(info.Mode()) & 0777
	//fmt.Printf("%o -> %o %o %o\n", perms, perms/64, (perms/8)%8, perms&7)

	{
		usr := perms / 64
		if usr >= 4 {
			modes[1] = 'r'
			usr -= 4
		}
		if usr >= 2 {
			modes[2] = 'w'
			usr -= 2
		}
		if usr > 0 && (info.Mode()&os.ModeSetuid) == 0 {
			modes[3] = 'x'
		}
	}
	{
		grp := (perms / 8) % 8
		if grp >= 4 {
			modes[4] = 'r'
			grp -= 4
		}
		if grp >= 2 {
			modes[5] = 'w'
			grp -= 2
		}
		if grp > 0 && (info.Mode()&os.ModeSetgid) == 0 {
			modes[6] = 'x'
		}
	}
	{
		oth := perms & 7
		if oth >= 4 {
			modes[7] = 'r'
			oth -= 4
		}
		if oth >= 2 {
			modes[8] = 'w'
			oth -= 2
		}
		if oth > 0 && (info.Mode()&os.ModeSticky) == 0 {
			modes[9] = 'x'
		}
	}

	return string(modes)
}
