package uihome

import (
	"os"
	"strings"

	"github.com/torbenschinke/picapsule/channel"
	"github.com/torbenschinke/picapsule/video"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func PageVideoDetails(wnd core.Window, uc video.UseCases, ucChan channel.UseCases) core.View {
	chanId := channel.ID(wnd.Values()["channel"])
	optChan, err := ucChan.FindByID(chanId)
	if err != nil {
		return MainErr(wnd, err)
	}

	if optChan.IsNone() {
		return MainErr(wnd, os.ErrNotExist)
	}

	channel := optChan.Unwrap()

	return Main(wnd,
		TopBar(wnd, "VIDEO DETAILS"),
		BackBar(wnd),

		List[video.Video, video.ID](
			wnd,
			1,
			uc.FindByID,
			func(yield func(video.ID, error) bool) {
				for _, id := range channel.Videos {
					if !yield(id, nil) {
						return
					}
				}
			},
			func(t video.Video) core.View {
				return PillView(
					wnd,
					ui.VStack(
						ui.VStack(
							ui.Text(t.Title).Font(ui.TitleLarge).Color("#67C9EF"),
							ui.Text(t.Description),
						).FullWidth().Alignment(ui.TopLeading).Padding(ui.Padding{}.All(ui.L8)),
						ui.Spacer(),
						ui.VStack(ui.Text(strings.ToUpper("Aufzeichnung abspielen")).Color(ui.ColorBlack)).
							FullWidth().
							BackgroundColor("#DE4228").
							Padding(ui.Padding{}.All(ui.L8)),
					).Alignment(ui.TopLeading).
						Gap(ui.L16).
						Frame(ui.Frame{Height: "calc(100vh - 15.5rem)", Width: ui.Full}),
					ui.Padding{},
					func() {
						wnd.Navigation().ForwardTo("video/play", core.Values{"video": string(t.ID)})
					},
				)
			},
		),
	)
}
