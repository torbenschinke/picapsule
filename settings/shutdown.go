package settings

import (
	"log/slog"
	"os/exec"
)

func Shutdown() {
	if err := exec.Command("sudo", "shutdown", "now").Run(); err != nil {
		slog.Error(err.Error())
	}
}
