package uikit

import "github.com/egoist/mygo/ui"

// FloatingPanel keeps all overlays single-layer and paints the chosen theme's
// tint and edge independently from the main window.
func (k Kit) FloatingPanel(c *ui.Context, shade ui.Color, padding, gap float32) ui.Element {
	panel := ui.Column(c).Fill().Padding(padding).Gap(gap)
	radius := max(12, c.Theme().Radius)
	border := c.Theme().Border.Alpha(.6)
	var top, bottom string
	var rim ui.Color
	angle, lowerOpacity := float32(180), float32(1)
	dark := k.Mode == "dark"
	switch k.floatingTheme {
	case "mica":
		radius, top, bottom, angle = 7, "#e9f0fa", "#c2d3ea", 145
		border, rim = ui.Hex("#8396b6").Alpha(.53), ui.Hex("#6f8db4").Alpha(.32)
		if dark {
			top, bottom, border = "#30415a", "#151b29", ui.Hex("#748dae").Alpha(.60)
		}
	case "acrylic":
		radius, top, bottom, angle, lowerOpacity = 10, "#f4edf7", "#d6ddec", 224, .82
		border = ui.Hex("#aaa3bf").Alpha(.67)
		if dark {
			top, bottom, border = "#3b3144", "#253040", ui.Hex("#aaa3bf").Alpha(.50)
		}
	case "macos":
		radius, top = 15, "#fdfdfd"
		border = ui.Hex("#babac0").Alpha(.63)
		if dark {
			top, border = "#222224", ui.Hex("#74747c").Alpha(.63)
		}
	case "liquid-glass":
		radius, top, bottom, lowerOpacity = 23, "#e9f7ff", "#d9e8ff", .72
		border, rim = ui.Hex("#ffffff").Alpha(.91), ui.Hex("#ffffff").Alpha(.55)
		if dark {
			top, bottom = "#2b3f57", "#0d1b2f"
			border, rim = ui.Hex("#c8ebff").Alpha(.81), ui.Hex("#b1e2ff").Alpha(.38)
		}
	case "retro":
		color, edge := "#eeece4", "#393733"
		if dark {
			color, edge = "#252524", "#92918b"
		}
		fill := ui.Hex(color)
		fill.A = shade.A
		return panel.Radius(0).Background(fill).BorderWidth(2, 4, 4, 2).BorderColor(ui.Hex(edge))
	}
	panel = panel.Radius(radius).Border(1, border)
	if top == "" {
		return panel.Background(shade)
	}
	first := ui.Hex(top)
	first.A = shade.A
	if bottom == "" {
		return panel.Background(first)
	}
	last := ui.Hex(bottom)
	last.A = uint8(float32(shade.A) * lowerOpacity)
	panel = panel.Gradient(first, last, angle)
	if rim.A != 0 {
		panel.Draw(func(p *ui.Painter, box ui.Rect) {
			box = ui.Rect{X: box.X + 1, Y: box.Y + 1, W: max(0, box.W-2), H: max(0, box.H-2)}
			height := float32(24)
			if k.floatingTheme == "mica" {
				height = 4
			}
			clear := rim
			clear.A = 0
			p.Clip(box, radius-1, func() {
				box.H = min(box.H, height)
				p.FillGradient(box, ui.LinearGradient{From: rim, To: clear, Angle: 180}, 0)
			})
		})
	}
	return panel
}
