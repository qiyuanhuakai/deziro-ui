package uikit

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// FloatingPanel paints the window tint and outline supplied by the active theme.
func (k Kit) FloatingPanel(c *ui.Context, shade ui.Color, padding, gap float32) ui.Element {
	if c.Vibrancy() && k.backdropPower > 0 {
		shade = shade.Alpha(float32(math.Pow(float64(shade.A)/255, k.backdropPower)))
	}
	palette := c.Theme()
	panel := ui.Column(c).Fill().Padding(padding).Gap(gap).Background(shade)
	if k.floatingTheme == "retro" {
		return panel.Radius(0).BorderWidth(2, 4, 4, 2).BorderColor(palette.Border)
	}
	return panel.Radius(palette.Radius).Border(1, palette.Border)
}
