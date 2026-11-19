# Deziro UI Web-derived theme presets

This GPL-3.0 Go module contains the six theme catalogs generated from the React
WebView CSS. It imports the separate MIT native kit. From an application view,
choose a preset and apply its kit to the current MyGo UI context.

```go
import presets "github.com/qiyuanhuakai/deziro-ui/packages/ui-presets"

kit := presets.Select("macos", "light").Kit()
kit.Apply(c)
```

`retro`, `mica`, `acrylic`, `liquid-glass`, `macos`, and `scifi` each accept
`light` and `dark`. The macOS, Liquid Glass, and Sci-Fi themes are original
designs by qiyuanhuakai. The other themes build on upstream styles. Each
generated catalog also includes shared Web control styles.

The native comparison gallery is under `cmd/dzc-ui-gallery`. See the top-level
README for screenshots and reproduction commands.
