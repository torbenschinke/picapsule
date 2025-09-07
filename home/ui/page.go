package uihome

import (
	"iter"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/worldiety/option"
	"go.wdy.de/nago/pkg/data"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"go.wdy.de/nago/presentation/ui/pager"
)

func Arc(wnd core.Window, text0, text1 string, body core.View) core.View {
	var pos *core.State[ui.Transformation]
	pos = core.AutoState[ui.Transformation](wnd).Init(func() ui.Transformation {
		go func() {
			time.Sleep(500 * time.Millisecond)
			pos.Set(ui.Transformation{})
		}()

		return ui.Transformation{
			TranslateX: "-10.5rem",
		}
	})

	return ui.HStack(
		ui.VStack(
			ui.VStack(
				ui.Text(text0).Color(ui.ColorBlack).Padding(ui.Padding{Right: ui.L8}),
				ui.VStack(
					ui.Text(text1).Color("#959BB0"),
				).BackgroundColor("#50566B").
					Alignment(ui.TopTrailing).
					Padding(ui.Padding{Right: ui.L8}).
					Frame(ui.Frame{}.FullHeight().FullWidth()).
					Border(ui.Border{TopLeftRadius: ui.L32, TopWidth: ui.L2, LeftWidth: ui.L2}.Color(ui.ColorBlack)),
			).BackgroundColor("#686E85").
				Alignment(ui.TopTrailing).
				Frame(ui.Frame{}.FullHeight().FullWidth()).
				Padding(ui.Padding{Right: ui.L8}).
				Border(ui.Border{TopLeftRadius: ui.L48}).
				Padding(ui.Padding{Left: ui.L48, Top: ui.L32}),
		).Alignment(ui.TopTrailing).
			ID("asdf").
			Animation(ui.AnimateTransition).
			Transformation(pos.Get()).
			Frame(ui.Frame{Height: "calc(100vh - 4.5rem)", MinWidth: "10rem"}),

		body,
	).FullWidth().Alignment(ui.TopLeading)
}

func TopBar(wnd core.Window, title string) core.View {
	pos := core.AutoState[ui.Transformation](wnd).Init(func() ui.Transformation {
		return ui.Transformation{
			TranslateY: "-5rem",
		}
	})

	go func() {
		time.Sleep(500 * time.Millisecond)
		pos.Set(ui.Transformation{})
	}()

	return ui.VStack(
		ui.HStack().
			BackgroundColor("#252C3A").
			Frame(ui.Frame{Height: ui.L48, Width: ui.Full}),
		ui.HStack().
			BackgroundColor("#454B5D").
			Frame(ui.Frame{Height: ui.L8, Width: ui.Full}),
		ui.HStack(
			ui.Text(title).Color("#B36B56").Font(ui.Font{Size: "4rem"}),
		).Position(ui.Position{
			Type:  ui.PositionAbsolute,
			Right: "8rem",
		}).BackgroundColor(ui.ColorBlack).
			Frame(ui.Frame{Height: "4rem"}).
			Padding(ui.Padding{}.All(ui.L8)),
	).Gap(ui.L4).
		Animation(ui.AnimateTransition).
		Transformation(pos.Get()).
		FullWidth().
		Border(ui.Border{TopLeftRadius: ui.L8, BottomLeftRadius: ui.L8})

}

func BackBar(wnd core.Window) core.View {
	var pos *core.State[ui.Transformation]
	pos = core.AutoState[ui.Transformation](wnd).Init(func() ui.Transformation {
		go func() {
			time.Sleep(500 * time.Millisecond)
			pos.Set(ui.Transformation{})
		}()

		return ui.Transformation{
			TranslateX: "-8rem",
		}
	})

	return ui.VStack(
		ui.SecondaryButton(func() {
			wnd.Navigation().Back()
		}).Title(strings.ToUpper("zurück")),
	).Animation(ui.AnimateTransition).
		Transformation(pos.Get()).FullWidth().
		Alignment(ui.Leading).
		Padding(ui.Padding{}.Vertical(ui.L8))
}

func Panel(wnd core.Window, views ...core.View) ui.TVStack {

	var pos *core.State[ui.Transformation]
	pos = core.AutoState[ui.Transformation](wnd).Init(func() ui.Transformation {
		go func() {
			time.Sleep(500 * time.Millisecond)
			pos.Set(ui.Transformation{})

		}()

		return ui.Transformation{
			TranslateX: "100rem",
		}
	})

	return ui.VStack(
		views...,
	).Gap(ui.L32).
		Animation(ui.AnimateTransition).
		Transformation(pos.Get()).FullWidth().
		Font(ui.DisplayMedium).
		BackgroundColor("#252C3A").
		Padding(ui.Padding{}.All(ui.L16)).
		Border(ui.Border{}.Radius(ui.L16)).(ui.TVStack)
}

func Main(wnd core.Window, views ...core.View) ui.TVStack {
	return ui.VStack(views...).
		Append(alert.BannerMessages(wnd)).
		Gap(ui.L8).
		Alignment(ui.Top).
		BackgroundColor(ui.ColorBlack).
		Frame(ui.Frame{}.MatchScreen()).(ui.TVStack)
}

func MainErr(wnd core.Window, err error) ui.TVStack {
	return Main(wnd,
		BackBar(wnd),
		alert.BannerError(err),
		Panel(wnd, ui.PrimaryButton(func() {
			wnd.Navigation().ForwardTo(".", nil)
		}).Title("Home")),
	)
}

func List[T data.Aggregate[ID], ID data.IDType](wnd core.Window, pageSize int, findByID func(id ID) (option.Opt[T], error), all iter.Seq2[ID, error], toView func(T) core.View) core.View {
	var pos *core.State[ui.Transformation]
	pageIdx := core.AutoState[int](wnd).Observe(func(newValue int) {
		pos.Set(ui.Transformation{TranslateX: "100vw"})
		go func() {
			time.Sleep(500 * time.Millisecond)
			pos.Set(ui.Transformation{})

		}()
	})

	pos = core.AutoState[ui.Transformation](wnd).Init(func() ui.Transformation {
		go func() {
			time.Sleep(500 * time.Millisecond)
			pos.Set(ui.Transformation{})

		}()

		return ui.Transformation{
			TranslateX: "100vw",
		}
	})

	page, err := data.Paginate(findByID, all, data.PaginateOptions{
		PageIdx:  pageIdx.Get(),
		PageSize: pageSize,
	})

	if err != nil {
		return alert.BannerError(err)
	}

	return ui.VStack(
		ui.ForEach(page.Items, toView)...,
	).Append(
		ui.Spacer(),
		ui.HStack(
			pager.Pager(pageIdx).Count(page.PageCount).Frame(ui.Frame{}.FullWidth()).Visible(page.PageCount > 1),
		).Font(ui.DisplaySmall),
	).
		Gap(ui.L16).
		FullWidth().
		Transformation(pos.Get()).
		Animation(ui.AnimateTransition).
		Border(ui.Border{}.Radius(ui.L16)).
		Padding(ui.Padding{}.All(ui.L16))
}

func PillButton(wnd core.Window, text string, action func()) core.View {
	return PillView(wnd, ui.Text(strings.ToUpper(text)).Font(ui.DisplayMedium), ui.Padding{}.Horizontal(ui.L8).Vertical(ui.L4), action)
}

func PillView(wnd core.Window, text core.View, pad ui.Padding, action func()) core.View {
	color := ui.Color("#52596E")
	return ui.HStack(
		ui.HStack().
			BackgroundColor(color).
			Frame(ui.Frame{MinWidth: "3rem", MinHeight: "4rem"}).
			Border(ui.Border{TopLeftRadius: ui.L96, BottomLeftRadius: ui.L96}),
		ui.VStack(text).
			Animation(ui.AnimateTransition).
			BackgroundColor(color).
			Frame(ui.Frame{}.FullWidth()).
			Padding(pad),
		ui.HStack().
			BackgroundColor(color).
			Frame(ui.Frame{MinWidth: "4rem", MinHeight: "4rem"}),
	).Alignment(ui.Stretch).
		Action(action).
		Gap(ui.L4).
		FullWidth()
}

func RectButton(wnd core.Window, text string, action func()) ui.THStack {
	color := ui.Color("#676E85")
	return ui.HStack(
		ui.Text(text).Color(ui.ColorBlack),
	).BackgroundColor(color).
		Action(action).
		Padding(ui.Padding{}.All(ui.L8)).(ui.THStack)
}

func VStack(wnd core.Window, views ...core.View) ui.TVStack {
	var pos *core.State[ui.Transformation]
	pos = core.AutoState[ui.Transformation](wnd).Init(func() ui.Transformation {
		go func() {
			time.Sleep(500 * time.Millisecond)
			pos.Set(ui.Transformation{})
		}()

		return ui.Transformation{
			TranslateX: "100vw",
		}
	})

	return ui.VStack(views...).Animation(ui.AnimateTransition).Transformation(pos.Get())
}

func NumPad(text *core.State[string]) core.View {
	num := 1
	return ui.VStack(
		ui.VStack(
			slices.Collect(func(yield func(view core.View) bool) {
				for range 3 {
					yield(ui.HStack(
						slices.Collect(func(yield func(view core.View) bool) {
							for range 3 {
								char := strconv.Itoa(num)
								yield(
									ui.PrimaryButton(func() {
										text.Set(text.Get() + char)
										text.Notify()
									}).Title(char),
								)
								num++
							}
						})...,
					).Gap(ui.L16))
				}
			})...,
		).Gap(ui.L16).
			Append(ui.PrimaryButton(func() {
				text.Set(text.Get() + strconv.Itoa(0))
				text.Notify()
			}).Title("0")),
	).FullWidth()
}
