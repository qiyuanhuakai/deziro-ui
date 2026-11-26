package uikit

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"
)

var videoIcons = map[string]*ui.SVG{
	"play": icon("Play"), "pause": icon("Pause"), "repeat": icon("Repeat"),
	"volume": icon("Volume2"), "muted": icon("VolumeX"),
}

type VideoLabels struct {
	Play, Pause, Progress, Mute, Unmute, Volume, Loop string
}

// PinnedVideo draws a playback viewport and controls within the preview surface.
func (k Kit) PinnedVideo(c *ui.Context, player *ui.VideoPlayer, labels VideoLabels) {
	ui.Column(c).FillWidth().Grow(1).MinHeight(0).Radius(14).Clip().Background(ui.Hex("#161d2a")).Children(func() {
		ui.Box(c).FillWidth().Grow(1).MinHeight(76).Background(ui.Hex("#080b10")).Center().Children(func() {
			ui.Video(c, player).Fill().Fit(ui.Contain)
		})
		ui.Column(c).FillWidth().Padding(3, 8, 6, 8).Children(func() {
			duration, position := player.Duration(), player.Position()
			progress := position.Seconds()
			if videoSlider(c.Key("position"), &progress, duration.Seconds(), .1, labels.Progress).FillWidth().Changed() {
				player.Seek(time.Duration(progress * float64(time.Second)))
			}
			ui.Row(c).FillWidth().Gap(4).AlignItems(ui.Center).Children(func() {
				playIcon, playLabel := "play", labels.Play
				if player.Playing() {
					playIcon, playLabel = "pause", labels.Pause
				}
				if videoButton(c.Key("playback"), playIcon, playLabel, false, false, 31).Clicked() {
					if player.Playing() {
						player.Pause()
					} else {
						player.Play()
					}
				}
				ui.Text(c, videoClock(position)+" / "+videoClock(duration)).FontSize(11).
					TextColor(ui.Hex("#f2f4fa")).Grow(1)
				muteIcon, muteLabel := "volume", labels.Mute
				if player.Muted() {
					muteIcon, muteLabel = "muted", labels.Unmute
				}
				if videoButton(c.Key("mute"), muteIcon, muteLabel, player.Muted(), false, 31).Clicked() {
					player.SetMuted(!player.Muted())
				}
				volume := player.Volume()
				if videoSlider(c.Key("volume"), &volume, 1, .05, labels.Volume).WidthPercent(23).MinWidth(28).MaxWidth(86).Changed() {
					player.SetVolume(volume)
					player.SetMuted(false)
				}
				if videoButton(c.Key("loop"), "repeat", labels.Loop, player.Loop(), false, 31).Clicked() {
					player.SetLoop(!player.Loop())
				}
			})
		})
	})
}

// VideoCover shows the player before opening the underlying file.
func (k Kit) VideoCover(c *ui.Context, label string, loading bool) bool {
	clicked := false
	ui.Box(c).FillWidth().Grow(1).MinHeight(96).Radius(14).Center().
		Gradient(ui.Hex("#303f58"), ui.Hex("#080b10"), 180).Children(func() {
		clicked = videoButton(c, "play", label, false, loading, 50).Clicked()
	})
	return clicked
}

func videoButton(c *ui.Context, glyph, label string, active, disabled bool, size float32) ui.Element {
	button := ui.ButtonBase(c).Label(label).Tooltip(label).Size(size, size).Center().
		FocusRing(true).Cursor(actionCursor(disabled)).Selected(active).Disabled(disabled)
	ink := ui.Hex("#f2f4fa")
	if size > 32 {
		button.Radius(size/2).Border(1, ui.Hex("#637698")).Background(ui.Hex("#202c42"))
	} else {
		button.Radius(8)
	}
	if button.Hovered() || button.FocusVisible() || active {
		button.Background(ui.Hex("#33547e"))
	}
	if active || disabled {
		ink = ui.Hex("#8cbdff")
	}
	button.Children(func() {
		ui.Icon(c, videoIcons[glyph]).Size(size/2, size/2).TextColor(ink).PassThrough()
	})
	return button
}

func videoClock(value time.Duration) string {
	seconds := int64(value.Seconds())
	if seconds >= 3600 {
		return fmt.Sprintf("%d:%02d:%02d", seconds/3600, seconds/60%60, seconds%60)
	}
	return fmt.Sprintf("%02d:%02d", seconds/60, seconds%60)
}

func videoSlider(c *ui.Context, value *float64, maximum, step float64, label string) ui.Element {
	slider := ui.SliderBase(c, value, 0, max(step, maximum)).Height(19).PaddingX(5).
		Step(step).Disabled(maximum <= 0).FocusRing(true).Label(label).Tooltip(label)
	return slider.Draw(func(p *ui.Painter, r ui.Rect) {
		fraction := min(1, max(0, float32(*value/max(step, maximum))))
		line := ui.Rect{X: r.X + 5, Y: r.Y + (r.H-3)/2, W: max(0, r.W-10), H: 3}
		p.Fill(line, ui.Hex("#414f66"), 1.5)
		line.W *= fraction
		p.Fill(line, ui.Hex("#8cbdff"), 1.5)
		knob := ui.Rect{X: r.X + line.W, Y: r.Y + (r.H-10)/2, W: 10, H: 10}
		p.Fill(knob, ui.Hex("#e6efff"), 5)
	})
}
