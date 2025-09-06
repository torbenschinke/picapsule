package uihome

import (
	"log/slog"
	"net"
	"os"
	"strings"
	"sync/atomic"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

var BrowserProcess atomic.Pointer[os.Process]

func PageSettings(wnd core.Window) core.View {
	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper("Einstellungen")),
		BackBar(wnd),
		ui.PrimaryButton(func() {
			wnd.Navigation().ForwardTo("settings/channels", nil)
		}).Title(strings.ToUpper("Video Kanäle")),
		ui.PrimaryButton(func() {
			wnd.Navigation().ForwardTo("settings/videos", nil)
		}).Title(strings.ToUpper("Video Dateien")),
		ui.PrimaryButton(func() {
			p := BrowserProcess.Load()
			if p != nil {
				if err := p.Kill(); err != nil {
					slog.Error("failed to kill browser process", "err", err.Error())
				}
			}
			os.Exit(0)
		}).Title("Kill Process"),

		ui.Text("Networks:\n"+getNetworks()).Font(ui.Font{Name: ui.MonoFontName}),
	)
}

func getNetworks() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err.Error()
	}

	var networks []string
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			networks = append(networks, err.Error())
			continue
		}

		for _, addr := range addrs {
			networks = append(networks, addr.String())
		}
	}

	return strings.Join(networks, "\n")
}
