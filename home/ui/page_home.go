package uihome

import (
	"strings"

	"github.com/torbenschinke/picapsule/channel"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

func PageHome(wnd core.Window, uc channel.UseCases) core.View {
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
					childSecurityDialog(wnd, antiChildDialogPresented),
					RectButton(wnd, strings.ToUpper("Einstellungen"), func() {
						antiChildDialogPresented.Set(true)
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

func childSecurityDialog(wnd core.Window, presented *core.State[bool]) core.View {
	if !presented.Get() {
		return nil
	}

	password := core.AutoState[string](wnd).Observe(func(newValue string) {
		//slog.Info(newValue)
	})
	return alert.Dialog(
		"Sicherheitscode",
		NumPad(password),
		presented,
		alert.Cancel(nil),
		alert.Apply(func() (close bool) {
			// it was not my intention to protect this at all: this is a time capsule and it must be repairable, however children may play with the device and must not harm it
			if password.Get() == "13092025" {
				wnd.Navigation().ForwardTo("settings", nil)
				return true
			}

			return false
		}),
	)
}
