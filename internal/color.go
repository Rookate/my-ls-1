package ls

import (
	"os"
	"strings"
)

const (
	ColorReset   = "\033[0m"
	ColorBold    = "\033[1m"
	ColorBlue    = "\033[34m"
	ColorGreen   = "\033[32m"
	ColorCyan    = "\033[36m"
	ColorRed     = "\033[31m"
	ColorYellow  = "\033[33m"
	ColorMagenta = "\033[35m"
)

func Colorize(name string, info os.FileInfo) string {
	var color string
	if info.IsDir() {
		color = ColorBold + ColorBlue
	} else if strings.HasSuffix(name, ".tar") || strings.HasSuffix(name, ".gz") || strings.HasSuffix(name, ".zip") {
		color = ColorBold + ColorRed
	} else if (info.Mode() & os.ModeSymlink) != 0 {
		color = ColorBold + ColorCyan
	} else if (info.Mode() & 0100) != 0 {
		color = ColorBold + ColorGreen
	}
	return color + name + ColorReset
}
