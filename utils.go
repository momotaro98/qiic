package qiic

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenBrowser opens browser with taken URL.
func OpenBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		return fmt.Errorf("%s is not supported", runtime.GOOS)
	}

	return nil
}
