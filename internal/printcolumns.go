package ls

import (
	"fmt"
	"os"
)

// Fonction pui permet de retravailler l'affichage des fichiers pour les mettre en colonnes. Try avec "go run . ls /usr/bin"
func PrintColumns(fileInfos []os.FileInfo) {
	const numberColumns = 6

	numFiles := len(fileInfos)
	filePerColumns := (numFiles + numberColumns - 1) / numberColumns

	columnWidth := 0

	for _, info := range fileInfos {
		nameLen := len(info.Name())
		if nameLen > columnWidth {
			columnWidth = nameLen
		}
	}
	columnWidth += 5

	for i := 0; i < filePerColumns; i++ {
		for j := i; j < numFiles; j += filePerColumns {
			info := fileInfos[j]
			name := Colorize(info.Name(), info)
			fmt.Printf("%-*s", columnWidth, name)
		}
		fmt.Println()
	}
}
