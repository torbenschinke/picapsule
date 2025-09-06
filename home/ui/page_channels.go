package uihome

import (
	"github.com/torbenschinke/picapsule/channel"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func PageChannels(wnd core.Window, uc channel.UseCases) core.View {
	return ui.VStack(
		TopBar(wnd, "VIDEO KANÄLE"),
		BackBar(wnd),
		ui.HStack(
			ui.PrimaryButton(func() {
				wnd.Navigation().ForwardTo("settings/channels/add", nil)
			}).Title("HINZUFÜGEN"),
		).FullWidth().Alignment(ui.Trailing),

		List[channel.Channel, channel.ID](
			wnd,
			5,
			uc.FindByID,
			uc.FindAll(),
			func(t channel.Channel) core.View {
				return PillButton(wnd, t.Title, func() {
					wnd.Navigation().ForwardTo("settings/channels/edit", core.Values{"channel": string(t.ID)})
				})
			},
		),
	).Gap(ui.L8).
		Alignment(ui.Top).
		BackgroundColor(ui.ColorBlack).
		Frame(ui.Frame{}.MatchScreen())
}
