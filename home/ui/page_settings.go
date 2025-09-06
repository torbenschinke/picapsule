package uihome

import (
	"os"
	"strings"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

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
			os.Exit(0)
		}).Title("Kill Process"),
	)
}
