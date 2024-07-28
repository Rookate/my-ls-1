package ls

import "io/fs"

// Fonction pour ne pas afficher les fichiers cachés si opts.All est désactivé comme par exemple .vscode
func HiddenFile(fileInfos []fs.FileInfo, opts Option) []fs.FileInfo {
	if !opts.All {
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
