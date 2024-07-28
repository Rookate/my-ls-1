package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"syscall"
)

// Fonction pour le flag -l pour afficher plus de détail
func Long(fileInfos []fs.FileInfo, opts Option) {
	for _, info := range fileInfos {
		var link string
		if info.Mode()&os.ModeSymlink != 0 { //condition pour savoir si le ficher est un lien symbolique
			target, err := os.Readlink(info.Name())
			if err != nil {
				fmt.Printf("Error resolving symlink: %v\n", err)
				link = " -> <error>"
			} else {
				link = " -> " + target
			}
		} else {
			link = ""
		}
		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			log.Fatal("Failed to get raw syscall.Stat_t data")
		}

		uid := fmt.Sprint(stat.Uid)
		gid := fmt.Sprint(stat.Gid)

		usr, err := user.LookupGroupId(uid)
		if err != nil {
			log.Fatal(err)
		}

		grp, err := user.LookupGroupId(gid)
		if err != nil {
			log.Fatal(err)
		}
		name := Colorize(info.Name(), info)
		fmt.Printf("%10v %3d %8v %8v %5d %5v %s %s\n",
			info.Mode(),
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
