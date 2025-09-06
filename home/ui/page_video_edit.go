package uihome

import (
	"os"
	"strings"

	"github.com/torbenschinke/picapsule/video"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"go.wdy.de/nago/presentation/ui/form"
)

func PageVideoEdit(wnd core.Window, ucVideo video.UseCases) core.View {
	id := video.ID(wnd.Values()["video"])
	optVid, err := ucVideo.FindByID(id)
	if err != nil {
		return MainErr(wnd, err)
	}

	if optVid.IsNone() {
		return MainErr(wnd, os.ErrNotExist)
	}

	model := core.AutoState[video.Video](wnd).Init(func() video.Video {
		return optVid.Unwrap()
	})

	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper("Video bearbeiten")),
		BackBar(wnd),
		ui.HStack(
			ui.SecondaryButton(func() {
				if err := ucVideo.Delete(model.Get().ID); err != nil {
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
				if err := ucVideo.Update(model.Get()); err != nil {
					alert.ShowBannerError(wnd, err)
					return
				}

				wnd.Navigation().Back()
			}).Title(strings.ToUpper("Speichern")),
		).FullWidth().Alignment(ui.Trailing),
	)
}
