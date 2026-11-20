package uikit

import (
	"strings"
	"time"

	"github.com/egoist/mygo/ui"
)

// SettingsToggle preserves checkbox accessibility while drawing a switch.
func (k Kit) SettingsToggle(c *ui.Context, on *bool, label string) ui.Element {
	style, palette := k.style("select"), c.Theme()
	row := ui.CheckboxBase(c, on).Label(label).FillWidth().MinHeight(34).Gap(10).AlignItems(ui.Center).FocusRing(true)
	row.Children(func() {
		text(c, label, style).Grow(1).Wrap().PassThrough()
		track := ui.Box(c).Size(44, 22).Shrink(0).PassThrough()
		fill, offset := palette.TextMuted.Alpha(.34), float32(3)
		if *on {
			fill, offset = palette.Accent, 25
		} else if row.Hovered() {
			fill = palette.TextMuted.Alpha(.5)
		}
		duration := time.Duration(k.palette().TransitionMs) * time.Millisecond
		if duration > 0 {
			fill = track.AnimateColor("switch-track", fill, ui.Animation{Duration: duration})
			offset = track.Animate("switch-position", offset, duration)
		}
		track.Radius(11).Background(fill).Children(func() {
			ui.Box(c).Absolute().Left(offset).Top(3).Size(16, 16).Radius(8).
				Background(palette.AccentText).PassThrough()
		})
	})
	return row
}

func (k Kit) HotkeyBinding(c *ui.Context, label, value, placeholder string, recording, disabled bool) ui.Element {
	style := k.style("select")
	if recording {
		style = k.style("selectFocus")
	}
	field := ui.ButtonBase(c).Label(label).Height(style.Height).Gap(6).AlignItems(ui.Center).FocusRing(true).Disabled(disabled)
	if !recording {
		style.Fill, style.Border, style.Outline, style.Shadows = [4]uint8{}, [4]uint8{}, nil, nil
		style.BorderWidth, style.BorderWidths, style.Padding = 0, nil, [4]float32{}
	}
	paint(field, k.interactive(c, field, style, recording))
	field.Children(func() {
		if value == "" {
			text(c, placeholder, style).PassThrough()
			return
		}
		cap := k.style("select")
		cap.Padding = [4]float32{0, 6, 0, 6}
		cap.Radius, cap.Size, cap.Weight = 4, 11, 600
		for _, key := range strings.Split(value, "+") {
			paint(ui.Box(c).Height(22).MinWidth(22).Center().PassThrough(), cap).
				Children(func() { text(c, strings.ToUpper(key), cap).SingleLine().PassThrough() })
		}
	})
	return field
}

// Checkbox draws the themed control on MyGo's accessible toggle primitive.
func (k Kit) Checkbox(c *ui.Context, checked *bool, label string) ui.Element {
	style := k.style("select")
	row := ui.CheckboxBase(c, checked).Label(label).FillWidth().MinHeight(32).
		AlignItems(ui.Center).Gap(8).FocusRing(true)
	row.Children(func() {
		fill, border := c.Theme().Surface, c.Theme().Border
		if *checked {
			fill, border = c.Theme().Accent, c.Theme().Accent
		} else if row.Hovered() || row.FocusVisible() {
			border = c.Theme().Accent
		}
		mark := ui.Box(c).Size(17, 17).Shrink(0)
		box := style
		box.Fill, box.Border = rgba(fill), rgba(border)
		box.Padding, box.Outline, box.Shadows = [4]float32{}, nil, nil
		box.BorderWidths = nil
		box.BorderWidth, box.Radius = 1, min(5, style.Radius)
		paint(mark, k.animatePaint(mark, box))
		if *checked {
			mark.Children(func() {
				ui.Text(c, "✓").Center().Fill().TextColor(c.Theme().AccentText).FontSize(12).PassThrough()
			})
		}
		text(c, label, style).Grow(1).Wrap().PassThrough()
	})
	return row
}

func (k Kit) TextField(c *ui.Context, value *string, label string) ui.Element {
	return k.styleField(c, ui.TextInputBase(c, value).Label(label))
}

func (k Kit) TextArea(c *ui.Context, value *string, label string) ui.Element {
	return k.styleField(c, ui.TextAreaBase(c, value).Label(label))
}

func (k Kit) styleField(c *ui.Context, field ui.Element) ui.Element {
	style := k.style("select")
	if field.Focused() {
		style = k.style("selectFocus")
	} else if field.Hovered() {
		style = k.style("selectHover")
	}
	style = k.animatePaint(field, style)
	return styledText(paint(field.Height(style.Height), style), style).FocusRing(true)
}

func (k Kit) Combo(c *ui.Context, value *string, label, placeholder string, choices []string, disabled bool) (ui.Element, bool) {
	type draft struct{ text, saved string }
	box := ui.Box(c).Height(k.style("select").Height)
	edit := ui.Local(box, "deziro-combobox", func() draft { return draft{text: *value, saved: *value} })
	if edit.saved != *value {
		edit.text, edit.saved = *value, *value
	}
	changed := false
	box.Children(func() {
		parts := ui.ComboboxBase(c, &edit.text)
		field := k.styleField(c, parts.Input.Fill().Label(label).Placeholder(placeholder).Disabled(disabled))
		parts.Popup(func(panel ui.Element) {
			paint(panel, k.style("selectMenu"))
			query := strings.ToLower(strings.TrimSpace(edit.text))
			ui.Scroll(c).FillWidth().MaxHeight(215).Children(func() {
				for _, choice := range choices {
					if parts.Filtering() && !strings.Contains(strings.ToLower(choice), query) {
						continue
					}
					item := parts.Item(choice).Label(choice)
					style := k.style("selectOption")
					if item.Highlighted() || choice == *value {
						style = k.style("selectSelected")
					}
					style = k.interactive(c, item, style, item.Highlighted() || choice == *value)
					paint(item.FillWidth().Height(style.Height), style)
					item.Children(func() { text(c, choice, style).SingleLine().PassThrough() })
				}
			})
		})
		if choice, ok := parts.Chosen(); ok && !disabled {
			changed = *value != choice
			*value, edit.text, edit.saved = choice, choice, choice
		} else if !field.Focused() && !parts.Open() {
			edit.text = *value
		}
	})
	return box, changed
}

// SettingsLabel paints the visible label of a settings row.
func (k Kit) SettingsLabel(c *ui.Context, label string) ui.Element {
	return text(c, label, k.style("settingsLabel"))
}

func (k Kit) Slider(c *ui.Context, value *float64, minimum, maximum float64, label string) ui.Element {
	track, fill, thumb := k.style("settingsRangeTrack"), k.style("settingsRangeFill"), k.style("settingsRangeThumb")
	inset := thumb.Width / 2
	slider := ui.SliderBase(c, value, minimum, maximum).Height(max(20, thumb.Height)).
		PaddingX(inset).Label(label)
	focused := slider.FocusVisible()
	focusStyle := k.style("settingsRangeFocus")
	slider.Draw(func(p *ui.Painter, bounds ui.Rect) {
		fraction := float32(0)
		if maximum > minimum {
			fraction = float32(min(1, max(0, (*value-minimum)/(maximum-minimum))))
		}
		rail := ui.Rect{X: bounds.X, Y: bounds.Y + (bounds.H-track.Height)/2, W: bounds.W, H: track.Height}
		p.Fill(rail, styleColor(track.Fill), track.Radius)
		p.Clip(rail, track.Radius, func() {
			progress := ui.Rect{X: rail.X, Y: rail.Y, W: rail.W * fraction, H: rail.H}
			p.Fill(progress, styleColor(fill.Fill), 0)
		})
		p.Stroke(rail, styleColor(track.Border), track.Radius, track.BorderWidth)
		knob := ui.Rect{X: bounds.X + max(0, bounds.W-thumb.Width)*fraction,
			Y: bounds.Y + (bounds.H-thumb.Height)/2, W: thumb.Width, H: thumb.Height}
		for _, shadow := range thumb.Shadows {
			if !shadow.Inset {
				v := shadow.Values
				p.Shadow(knob, thumb.Radius, v[0], v[1], v[2], v[3], styleColor(shadow.Color))
			}
		}
		p.Fill(knob, styleColor(thumb.Fill), thumb.Radius)
		p.Stroke(knob, styleColor(thumb.Border), thumb.Radius, thumb.BorderWidth)
		if focused {
			if ring := focusStyle.Outline; ring != nil {
				offset := ring.Offset + ring.Width
				r := ui.Rect{X: rail.X - offset, Y: rail.Y - offset, W: rail.W + 2*offset, H: rail.H + 2*offset}
				p.Stroke(r, styleColor(ring.Color), min(track.Radius, rail.H/2)+offset, ring.Width)
			} else {
				p.Stroke(knob, p.Theme().Accent, thumb.Radius, 2)
			}
		}
	})
	return slider
}
