package uihome

import (
	"os"
	"strings"

	"github.com/torbenschinke/picapsule/video"
	"go.wdy.de/nago/pkg/xstrings"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	video2 "go.wdy.de/nago/presentation/ui/video"
)

func PageVideoPlay(wnd core.Window, ucVideo video.UseCases) core.View {
	id := video.ID(wnd.Values()["video"])
	optVid, err := ucVideo.FindByID(id)
	if err != nil {
		return MainErr(wnd, err)
	}

	if optVid.IsNone() {
		return MainErr(wnd, os.ErrNotExist)
	}

	vid := optVid.Unwrap()

	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper(xstrings.EllipsisEnd(vid.Title, 20))),
		BackBar(wnd),
		video2.Video("/api/video/download?id="+core.URI(vid.ID)).
			AutoPlay(true).
			Loop(true).
			Controls(true).
			Frame(ui.Frame{MaxHeight: "calc(100vh - 8.5rem)"}),
	)
}
