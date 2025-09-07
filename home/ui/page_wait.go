package uihome

import (
	"strings"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func PageWait(wnd core.Window) core.View {
	antiChildDialogPresented := core.AutoState[bool](wnd)

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
						antiChildDialogPresented.Set(true)
					}).Frame(ui.Frame{MinHeight: "3.5rem"}),
				).FullWidth().Alignment(ui.Trailing),

				ui.Space("13rem"),

				ui.VStack(ui.Text(strings.ToUpper("Bitte warten")).Font(ui.DisplayLarge).Color("#FF6753")).Padding(ui.Padding{}.All(ui.L16)),
			).Alignment(ui.Top).
				Frame(ui.Frame{Height: ui.Full, Width: ui.Full}),
		),
	).Alignment(ui.TopLeading)
}
