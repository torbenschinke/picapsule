package uihome

import (
	"github.com/torbenschinke/picapsule/video"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func PageVideos(wnd core.Window, uc video.UseCases) core.View {
	return Main(wnd,
		TopBar(wnd, "VIDEO KANÄLE"),
		BackBar(wnd),
		ui.HStack(
			ui.PrimaryButton(func() {
				wnd.Navigation().ForwardTo("settings/videos/add", nil)
			}).Title("HINZUFÜGEN"),
		).FullWidth().Alignment(ui.Trailing),

		List[video.Video, video.ID](
			wnd,
			5,
			uc.FindByID,
			uc.FindAllIdent(),
			func(t video.Video) core.View {
				return PillButton(wnd, t.Title, func() {
					wnd.Navigation().ForwardTo("settings/videos/edit", core.Values{"video": string(t.ID)})
				})
			},
		),
	)
}
