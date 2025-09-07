package settings

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func SetWFPanelAutoHide(enabled bool) {
	if err := setAutohide("/home/tschinke/.config/wf-panel-pi.ini", enabled); err != nil {
		slog.Error(err.Error())
	}
}

// setAutohide updates the wf-panel-pi.ini file and enables or disables autohide.
// - configPath: path to the wf-panel-pi.ini file (usually ~/.config/wf-panel-pi.ini)
// - enable: true to set autohide=true, false to set autohide=false
func setAutohide(configPath string, enable bool) error {
	// Open existing config file
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	var lines []string
	autohideFound := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "autohide=") {
			autohideFound = true
			if enable {
				lines = append(lines, "autohide=true")
			} else {
				lines = append(lines, "autohide=false")
			}
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading config: %w", err)
	}

	// Add autohide entry if not found
	if !autohideFound {
		if enable {
			lines = append(lines, "autohide=true")
		} else {
			lines = append(lines, "autohide=false")
		}
	}

	// Overwrite the config file
	err = os.WriteFile(configPath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	// Restart the panel so changes take effect
	//exec.Command("killall", "wf-panel-pi").Run()
	//exec.Command("wf-panel-pi").Start()

	if enable {
		slog.Info("Autohide has been enabled.")
	} else {
		slog.Info("Autohide has been disabled.")
	}
	return nil
}
