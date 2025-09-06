package uihome

import (
	"strings"

	"github.com/torbenschinke/picapsule/channel"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

func PageChannelAdd(wnd core.Window, ucChannel channel.UseCases) core.View {
	name := core.AutoState[string](wnd)
	desc := core.AutoState[string](wnd)
	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper("Kanal hinzufügen")),
		BackBar(wnd),
		Panel(wnd,
			ui.TextField("Name", name.Get()).InputValue(name),
			ui.TextField("Beschreibung", desc.Get()).InputValue(desc).Lines(3),
		).FullWidth(),

		ui.HStack(
			ui.PrimaryButton(func() {
				if _, err := ucChannel.Create(name.Get(), desc.Get()); err != nil {
					alert.ShowBannerError(wnd, err)
					return
				}

				wnd.Navigation().Back()
			}).Title(strings.ToUpper("Hinzufügen")),
		).FullWidth().Alignment(ui.Trailing),
	)
}
