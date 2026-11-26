package uikit

import (
	"math"
	"time"

	"github.com/egoist/mygo/ui"
)

type QuickPasteRingItem struct {
	Key, Label, CompactLabel, Detail string
	Active                           bool
	Select, Hover                    func()
}

type quickPasteSector struct {
	bounds ui.Rect
	label  ui.Point
	width  float32
	path   ui.Path
}

// QuickPasteRings draws an inner category ring and six outer entry sectors.
// At either edge, the same controls occupy an inward-facing half circle.
func (k Kit) QuickPasteRings(c *ui.Context, dock string, slots int, tint ui.Color, categories, entries []QuickPasteRingItem) {
	ring := ui.Box(c).FillWidth().Grow(1).MinHeight(0).Clip()
	size := ring.Bounds()
	ring.Draw(func(_ *ui.Painter, bounds ui.Rect) {
		if bounds.W != size.W || bounds.H != size.H {
			c.Invalidate()
		}
	})
	if size.W <= 0 || size.H <= 0 {
		return
	}
	ring.Children(func() {
		k.quickPasteRing(c, dock, size, tint, categories, len(categories), false)
		k.quickPasteRing(c, dock, size, tint, entries, slots, true)
	})
}

func (k Kit) quickPasteRing(c *ui.Context, dock string, size ui.Rect, tint ui.Color, items []QuickPasteRingItem, count int, outer bool) {
	palette := c.Theme()
	for index := range count {
		sector := makeQuickPasteSector(index, count, outer, dock, size.W, size.H)
		bounds := sector.bounds
		if index >= len(items) {
			ui.Box(c.Key(index)).Absolute().Left(bounds.X).Top(bounds.Y).Size(bounds.W, bounds.H).
				PassThrough().Draw(func(p *ui.Painter, rect ui.Rect) {
				drawQuickPasteSector(p, rect, sector, tint.Alpha(float32(tint.A)*.42/255))
			})
			continue
		}

		item := items[index]
		button := ui.ButtonBase(c.Key(struct {
			Outer bool
			Key   string
		}{outer, item.Key})).Absolute().Left(bounds.X).Top(bounds.Y).Size(bounds.W, bounds.H).
			Label(item.Label).Tooltip(item.Label).Role(ui.RoleButton).AccessSelected(item.Active).KeepFocus().
			Cursor(ui.CursorPointer).FocusRing(false).
			HitTest(func(x, y float32, rect ui.Rect) bool {
				return sector.path.Contains(x-rect.X, y-rect.Y)
			})
		fill, text := tint, palette.Text
		if item.Active {
			fill, text = palette.Accent, palette.AccentText
		} else if button.Hovered() || button.Focused() {
			fill = palette.SurfaceHover
		}
		fill = button.AnimateColor("fill", fill, ui.Animation{Duration: 140 * time.Millisecond})
		button.Draw(func(p *ui.Painter, rect ui.Rect) {
			drawQuickPasteSector(p, rect, sector, fill)
		}).Children(func() {
			label, font, height := item.Label, float32(12), float32(15)
			if outer {
				font, height = 13, 26
			} else if dock != "" {
				font = 10
				if item.CompactLabel != "" {
					label = item.CompactLabel
				}
			}
			ui.Column(c).Absolute().Left(sector.label.X-sector.width/2).Top(sector.label.Y-height/2).
				Size(sector.width, height).Center().PassThrough().Children(func() {
				ui.Text(c, label).FillWidth().SingleLine().TextAlign(ui.Center).TextColor(text).FontSize(font).PassThrough()
				if item.Detail != "" {
					ui.Text(c, item.Detail).FillWidth().SingleLine().TextAlign(ui.Center).
						TextColor(text.Alpha(.82)).FontSize(10).PassThrough()
				}
			})
		})
		if button.Hovered() && item.Hover != nil {
			item.Hover()
		}
		if button.Clicked() && button.ClickModifiers() == 0 && item.Select != nil {
			item.Select()
		}
	}
}

func drawQuickPasteSector(p *ui.Painter, rect ui.Rect, sector quickPasteSector, fill ui.Color) {
	p.Transform(ui.Translation(rect.X, rect.Y), 0, 0, func() {
		p.FillPath(&sector.path, fill)
	})
}

func makeQuickPasteSector(index, count int, outer bool, dock string, width, height float32) quickPasteSector {
	radius := max(0, min(float64(height)/2-4, float64(width)/2-4))
	cx, cy := float64(width)/2, float64(height)/2
	direction, span, offset, start := 1.0, 2*math.Pi, 0.0, -math.Pi/2
	switch dock {
	case "top":
		radius = max(0, min(float64(height)-4, float64(width)/2-4))
		cy, direction, span, offset, start = 0, -1, math.Pi, .5, 0
	case "left", "right":
		radius = max(0, min(float64(height)/2-4, float64(width)-8))
		cx, span, offset = 0, math.Pi, .5
		if dock == "right" {
			cx, direction = float64(width), -1
		}
	}
	inner, edge := .28*radius, .66*radius
	halfWidth, halfHeight := 36.0, 7.0
	if outer {
		inner, edge = .69*radius, radius
		halfWidth, halfHeight = 54, 12
	}
	step := span / float64(count)
	middle := start + step*(float64(index)+offset)
	halfAngle := step/2 - .018
	point := func(r, angle float64) ui.Point {
		return ui.Point{X: float32(cx + direction*r*math.Cos(angle)), Y: float32(cy + r*math.Sin(angle))}
	}
	label := point((inner+edge)/2, middle)
	const samples = 18
	points := [2 * (samples + 1)]ui.Point{}
	for i := range samples + 1 {
		angle := middle - halfAngle + 2*halfAngle*float64(i)/samples
		points[i] = point(edge, angle)
		points[len(points)-1-i] = point(inner, angle)
	}
	left, top, right, bottom := points[0].X, points[0].Y, points[0].X, points[0].Y
	for _, p := range points[1:] {
		left, top, right, bottom = min(left, p.X), min(top, p.Y), max(right, p.X), max(bottom, p.Y)
	}

	// Keep the horizontal label within the curved sides of the sector.
	contains := func(dx, dy float64) bool {
		x, y := direction*(float64(label.X)+dx-cx), float64(label.Y)+dy-cy
		angle := math.Atan2(y, x) - middle
		distance := math.Hypot(x, y)
		return distance >= inner+2 && distance <= edge-2 &&
			math.Abs(math.Atan2(math.Sin(angle), math.Cos(angle))) <= halfAngle
	}
	for halfWidth > 2 && !(contains(-halfWidth, -halfHeight) && contains(halfWidth, -halfHeight) &&
		contains(-halfWidth, halfHeight) && contains(halfWidth, halfHeight)) {
		halfWidth--
	}

	sector := quickPasteSector{
		bounds: ui.Rect{X: left, Y: top, W: right - left, H: bottom - top},
		label:  ui.Point{X: label.X - left, Y: label.Y - top},
		width:  float32(2 * halfWidth),
	}
	sector.path.MoveTo(points[0].X-left, points[0].Y-top)
	for _, p := range points[1:] {
		sector.path.LineTo(p.X-left, p.Y-top)
	}
	sector.path.Close()
	return sector
}
