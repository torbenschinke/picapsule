package uihome

import (
	"fmt"
	"strings"

	"github.com/torbenschinke/picapsule/video"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

func PageVideoAdd(wnd core.Window, ucVideo video.UseCases) core.View {
	return Main(
		wnd,
		TopBar(wnd, strings.ToUpper("Videos hinzufügen")),
		BackBar(wnd),
		Panel(wnd,
			ui.PrimaryButton(func() {
				wnd.ImportFiles(core.ImportFilesOptions{
					Multiple:         true,
					AllowedMimeTypes: []string{"video/mp4", "video/m4v", "video/x-m4v"},
					OnCompletion: func(files []core.File) {
						for _, file := range files {
							if _, err := ucVideo.Create(file); err != nil {
								alert.ShowBannerError(wnd, err)
							} else {
								alert.ShowBannerMessage(wnd, alert.Message{
									Title:   "Video hinzugefügt",
									Message: fmt.Sprintf("Erfolgreich hochgeladen: %s", file.Name()),
									Intent:  alert.IntentOk,
								})
							}
						}
					},
				})
			}).Title(strings.ToUpper("Hochladen")),
		).FullWidth(),
	)
}
