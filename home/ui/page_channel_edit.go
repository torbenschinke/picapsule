package uihome

import (
	"os"
	"strings"

	"github.com/torbenschinke/picapsule/channel"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"go.wdy.de/nago/presentation/ui/form"
)

func PageChannelEdit(wnd core.Window, ucChannel channel.UseCases) core.View {
	id := channel.ID(wnd.Values()["channel"])
	optChan, err := ucChannel.FindByID(id)
	if err != nil {
		return MainErr(wnd, err)
	}

	if optChan.IsNone() {
		return MainErr(wnd, os.ErrNotExist)
	}

	model := core.AutoState[channel.Channel](wnd).Init(func() channel.Channel {
		return optChan.Unwrap()
	})

	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper("Kanal bearbeiten")),
		BackBar(wnd),
		ui.HStack(
			ui.SecondaryButton(func() {
				if err := ucChannel.Delete(model.Get().ID); err != nil {
					alert.ShowBannerError(wnd, err)
					return
				}

				wnd.Navigation().Back()
			}).Title("Jetzt löschen"),
		).FullWidth().Alignment(ui.Trailing),
		Panel(wnd,
			form.Auto(form.AutoOptions{Window: wnd}, model),
		).FullWidth(),

		ui.HStack(
			ui.PrimaryButton(func() {
				if err := ucChannel.Update(model.Get()); err != nil {
					alert.ShowBannerError(wnd, err)
					return
				}

				wnd.Navigation().Back()
			}).Title(strings.ToUpper("Speichern")),
		).FullWidth().Alignment(ui.Trailing),
	)
}
