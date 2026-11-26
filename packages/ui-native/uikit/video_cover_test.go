package uikit

import (
	"github.com/egoist/mygo/ui"
	"testing"
)

func TestVideoCoverPlayReceivesPointerAbovePoster(t *testing.T) {
	opened := false
	view := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Children(func() {
			if (Kit{}).VideoCover(c, "Play video", false, func(c *ui.Context) {
				ui.Box(c).Fill().Cursor(ui.CursorPointer).Background(ui.Hex("#346bb8"))
			}) {
				opened = true
			}
		})
	}, 480, 360)
	if err := view.Click("Play video"); err != nil {
		t.Fatal(err)
	}
	if !opened {
		t.Fatal("the poster intercepted the Play button pointer event")
	}
}
