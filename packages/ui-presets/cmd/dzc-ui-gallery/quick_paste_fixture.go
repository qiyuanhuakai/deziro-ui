package main

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/qiyuanhuakai/deziro-ui/packages/ui-native/uikit"
	"github.com/qiyuanhuakai/deziro-ui/packages/ui-presets"
)

func (g *gallery) quickPasteFixture(c *ui.Context, preset presets.Preset) {
	opacity, _ := strconv.Atoi(g.state)
	if g.fixture == "quick-paste-rings" {
		opacity = 65
	}
	kit := preset.Kit()
	kit.Apply(c)
	palette := *c.Theme()
	palette.Background = ui.Hex("#6c6c6c")
	c.SetTheme(&palette)
	if g.fixture == "quick-paste-rings" {
		g.quickPasteRingFixture(c, kit, preset.WindowBackground(94))
		return
	}
	kit.QuickPastePanel(c, preset.WindowBackground(100).Alpha(float32(opacity)/100)).Children(func() {
		kit.QuickPasteSearch(c, &g.query, "Search clipboard…")
		ui.Column(c).FillWidth().Children(func() {
			kit.QuickPasteRow(c, "Release checklist", "Editor", true)
			kit.QuickPasteRow(c, "Design notes", "Browser", false)
			kit.QuickPasteRow(c, "Preview package ready for testing", "Terminal", false)
		})
	})
}

func (g *gallery) quickPasteRingFixture(c *ui.Context, kit uikit.Kit, tint ui.Color) {
	dock := g.state
	if dock == "full" {
		dock = ""
	}
	categories := []uikit.QuickPasteRingItem{
		{Key: "text", Label: "Text", Active: true},
		{Key: "image", Label: "Image", CompactLabel: "Img"},
		{Key: "file", Label: "File"},
		{Key: "link", Label: "Link"},
		{Key: "code", Label: "Code"},
		{Key: "video", Label: "Video", CompactLabel: "Vid"},
		{Key: "rich", Label: "Rich text", CompactLabel: "Rich"},
	}
	entries := []uikit.QuickPasteRingItem{
		{Key: "release", Label: "Release notes", Detail: "Editor", Active: true},
		{Key: "design", Label: "Design notes", Detail: "Browser"},
		{Key: "report", Label: "Report.pdf", Detail: "Explorer"},
		{Key: "test", Label: "Run tests", Detail: "Terminal"},
		{Key: "update", Label: "Update guide", Detail: "Editor"},
		{Key: "colors", Label: "Colors.png", Detail: "Paint"},
	}
	slots := 8
	if dock != "" {
		slots = 5
	}
	kit.QuickPasteRings(c, dock, slots, tint, categories, entries)
}
