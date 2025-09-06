package uihome

import (
	"strings"

	"github.com/torbenschinke/picapsule/channel"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func PageHome(wnd core.Window, uc channel.UseCases) core.View {
	return Main(
		wnd,
		TopBar(wnd, "Zeitkapsel"),
		Arc(
			wnd,
			strings.ToUpper("Erinnerungen"),
			strings.ToUpper("Videos"),
			VStack(wnd,
				ui.HStack(
					RectButton(wnd, strings.ToUpper("Einstellungen"), func() {
						wnd.Navigation().ForwardTo("settings", nil)
					}).Frame(ui.Frame{MinHeight: "3.5rem"}),
				).FullWidth().Alignment(ui.Trailing),

				List[channel.Channel, channel.ID](
					wnd,
					6,
					uc.FindByID,
					uc.FindAll(),
					func(t channel.Channel) core.View {
						return PillButton(wnd, t.Title, func() {
							wnd.Navigation().ForwardTo("video/details", core.Values{"channel": string(t.ID)})
						})
					},
				),
			).Alignment(ui.Top).
				Frame(ui.Frame{Height: ui.Full, Width: ui.Full}),
		),
	).Alignment(ui.TopLeading)
}
