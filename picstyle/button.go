package picstyle

import (
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Button(views ...core.View) ui.THStack {
	return ui.HStack(views...).BackgroundColor(ui.ColorBackground)
}
