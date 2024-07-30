package ls

import (
	"os"
	"strings"
)

const (
	ColorReset     = "\033[0m"
	ColorBold      = "\033[1m"
	ColorFGBlack   = "\033[30m"
	ColorFGRed     = "\033[31m"
	ColorFGGreen   = "\033[32m"
	ColorFGYellow  = "\033[33m"
	ColorFGBlue    = "\033[34m"
	ColorFGMagenta = "\033[35m"
	ColorFGCyan    = "\033[36m"
	ColorFGWhite   = "\033[37m"
	ColorBGBlack   = "\033[40m"
	ColorBGRed     = "\033[41m"
	ColorBGGreen   = "\033[42m"
	ColorBGYellow  = "\033[43m"
	ColorBGBlue    = "\033[44m"
	ColorBGMagenta = "\033[45m"
	ColorBGCyan    = "\033[46m"
	ColorBGWhite   = "\033[47m"
)

func Colorize(name string, info os.FileInfo) string {
	var color string
	if info.IsDir() {
		color = ColorBold + ColorFGBlue
	} else if strings.HasSuffix(name, ".tar") || strings.HasSuffix(name, ".gz") || strings.HasSuffix(name, ".zip") {
		color = ColorBold + ColorFGRed
	} else if (info.Mode() & os.ModeSetuid) != 0 {
		color = ColorFGWhite + ColorBGRed
	} else if (info.Mode() & os.ModeSetgid) != 0 {
		color = ColorFGBlack + ColorBGYellow
	} else if (info.Mode() & os.ModeSymlink) != 0 {
		color = ColorBold + ColorFGCyan
	} else if (info.Mode() & 0100) != 0 {
		color = ColorBold + ColorFGGreen
	}
	return color + name + ColorReset
}
