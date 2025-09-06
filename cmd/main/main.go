package main

import (
	_ "embed"
	"iter"
	"log/slog"
	"os"
	"os/exec"
	"time"

	"github.com/torbenschinke/picapsule/channel"
	uihome "github.com/torbenschinke/picapsule/home/ui"
	"github.com/torbenschinke/picapsule/video"
	"github.com/worldiety/option"
	"go.wdy.de/nago/application"
	"go.wdy.de/nago/application/theme"
	"go.wdy.de/nago/application/user"
	"go.wdy.de/nago/auth"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui/form"
	"go.wdy.de/nago/web/vuejs"
)

//go:embed palm.mp4
var palm application.StaticBytes

func main() {
	os.Setenv("NO_SSL", "true") // without ssl normal browsers other than localhost will get nil pointer in session management
	application.Configure(func(cfg *application.Configurator) {
		cfg.SetApplicationID("de.worldiety.tutorial_72")
		cfg.Serve(vuejs.Dist())

		cfg.SetHost("0.0.0.0")

		myTheme := option.Must(cfg.ThemeManagement())
		myBaseColors := theme.BaseColors{
			Main:        "#2f3749",
			Interactive: "#e5432a",
			Accent:      "#36a3cd",
		}

		myDark := myTheme.UseCases.Calculations.DarkMode(myBaseColors)

		option.MustZero(myTheme.UseCases.UpdateColors(user.SU(), theme.Colors{
			Dark:  myDark,
			Light: myDark,
		}))

		repoChannel := option.Must(application.JSONRepository[channel.Channel](cfg, "pic.channel"))
		ucChannel := channel.NewUseCases(repoChannel)

		videoFiles := option.Must(cfg.FileStore("videos"))
		repoVideo := option.Must(application.JSONRepository[video.Video](cfg, "pic.video"))
		ucVideo := video.NewUseCases(repoVideo, videoFiles)
		cfg.AddContextValue(core.ContextValue("pic.videos", form.AnyUseCaseList[video.Video, video.ID](func(subject auth.Subject) iter.Seq2[video.Video, error] {
			return ucVideo.FindAll()
		})))

		cfg.RootView("video/details", func(wnd core.Window) core.View {
			return uihome.PageVideoDetails(wnd, ucVideo, ucChannel)
		})

		cfg.RootView("video/play", func(wnd core.Window) core.View {
			return uihome.PageVideoPlay(wnd, ucVideo)
		})

		cfg.RootView("settings", func(wnd core.Window) core.View {
			return uihome.PageSettings(wnd)
		})

		cfg.RootView("settings/videos", func(wnd core.Window) core.View {
			return uihome.PageVideos(wnd, ucVideo)
		})

		cfg.RootView("settings/videos/add", func(wnd core.Window) core.View {
			return uihome.PageVideoAdd(wnd, ucVideo)
		})

		cfg.RootView("settings/videos/edit", func(wnd core.Window) core.View {
			return uihome.PageVideoEdit(wnd, ucVideo)
		})

		cfg.RootView("settings/channels", func(wnd core.Window) core.View {
			return uihome.PageChannels(wnd, ucChannel)
		})

		cfg.RootView("settings/channels/add", func(wnd core.Window) core.View {
			return uihome.PageChannelAdd(wnd, ucChannel)
		})

		cfg.RootView("settings/channels/edit", func(wnd core.Window) core.View {
			return uihome.PageChannelEdit(wnd, ucChannel)
		})

		cfg.RootView(".", func(wnd core.Window) core.View {
			return uihome.PageHome(wnd, ucChannel)
		})

		cfg.HandleFunc("/api/video/download", video.NewHandler(ucVideo, videoFiles))

		go launchBrowser()
	}).Run()
}

func launchBrowser() {
	slog.Info("waiting before launch browser")
	time.Sleep(500 * time.Millisecond)
	//cmd := exec.Command("chromium-browser", "--kiosk", "http://localhost:3000", "---no-user-gesture-required", "--ignore-autoplay-restrictions")
	cmd := exec.Command("firefox", "--kiosk", "http://localhost:3000")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "DISPLAY=:0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		slog.Error("failed to launch browser:", "err", err)
	}

	uihome.BrowserProcess.Store(cmd.Process)

	if err := cmd.Wait(); err != nil {
		slog.Error("failed to await browser:", "err", err)
	}
}
