package uihome

import (
	"log/slog"
	"os"
	"strings"
	"sync/atomic"

	"github.com/torbenschinke/picapsule/settings"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

var BrowserProcess atomic.Pointer[os.Process]

func PageSettings(wnd core.Window) core.View {
	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper("Einstellungen")),
		BackBar(wnd),
		ui.VStack(
			ui.PrimaryButton(func() {
				wnd.Navigation().ForwardTo("settings/channels", nil)
			}).Title(strings.ToUpper("Video Kanäle")).Frame(ui.Frame{}.FullWidth()),

			ui.PrimaryButton(func() {
				wnd.Navigation().ForwardTo("settings/videos", nil)
			}).Title(strings.ToUpper("Video Dateien")).Frame(ui.Frame{}.FullWidth()),

			ui.PrimaryButton(func() {
				settings.SetWFPanelAutoHide(true)
			}).Title("Taskleiste ausblenden").Frame(ui.Frame{}.FullWidth()),

			ui.PrimaryButton(func() {
				settings.SetWFPanelAutoHide(false)
			}).Title("Taskleiste einblenden").Frame(ui.Frame{}.FullWidth()),

			ui.PrimaryButton(func() {
				settings.Shutdown()
			}).Title("Herunterfahren").Frame(ui.Frame{}.FullWidth()),

			ui.PrimaryButton(func() {
				p := BrowserProcess.Load()
				if p != nil {
					if err := p.Kill(); err != nil {
						slog.Error("failed to kill browser process", "err", err.Error())
					}
				}
				os.Exit(0)
			}).Title("Beenden und zum Desktop").Frame(ui.Frame{}.FullWidth()),

			ui.Text("Networks:\n"+settings.GetNetworks()).Font(ui.Font{Name: ui.MonoFontName}).Frame(ui.Frame{}.FullWidth()),
		).Gap(ui.L16).Frame(ui.Frame{Width: ui.L256}),
	)
}
