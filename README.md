# Deziro UI

从 DezirClip 提取的 React WebView 和 MyGo Native 组件。两套库提供独立的导出，可在同一个仓库里查看源码、主题和渲染对照。

| 导出 | 用途 | 许可证 |
| --- | --- | --- |
| [`@deziro/ui-react`](packages/ui-react) | React 组件、CSS 和 Web 主题 | [GPL‑3.0](packages/ui-react/LICENSE) |
| [`packages/ui-native/uikit`](packages/ui-native/uikit) | Go 原生组件，带标准浅色与深色样式 | [MIT](packages/ui-native/LICENSE) |
| [`packages/ui-native/theme`](packages/ui-native/theme) | Go 主题类型、调色板和样式接口 | [MIT](packages/ui-native/LICENSE) |
| [`packages/ui-presets`](packages/ui-presets) | 由 Web 样式生成的六套原生主题预设 | [GPL‑3.0](packages/ui-presets/LICENSE) |

## 使用

### React WebView

```sh
pnpm add https://github.com/qiyuanhuakai/deziro-ui/releases/download/v0.1.1/deziro-ui-react-0.1.1.tgz
```

```tsx
import ThemedSelect from "@deziro/ui-react/components/ThemedSelect";
import { loadBuiltinTheme } from "@deziro/ui-react/themes";
import "@deziro/ui-react/styles/base.css";
import "@deziro/ui-react/styles/components/index.css";

await loadBuiltinTheme("macos");
document.documentElement.classList.add("theme-macos", "light-mode");
document.body.classList.add("theme-macos", "light-mode");
```

组件接受应用传入的状态和回调。适合 React 19 和能处理 CSS 导入的 Web 构建环境。

### MyGo Native

```sh
go get github.com/qiyuanhuakai/deziro-ui/packages/ui-native/uikit@v0.1.1
```

```go
import "github.com/qiyuanhuakai/deziro-ui/packages/ui-native/uikit"

kit := uikit.Kit{Mode: "dark"}
kit.Apply(c)
```

使用自己的主题时，向 `uikit.NewKit` 传入语义颜色和控件样式。应用需要下方完整的 Web 对照主题时，可以额外导入 GPL 预设模块，并通过 `presets.Select("macos", "light").Kit()` 应用。

```sh
go get github.com/qiyuanhuakai/deziro-ui/packages/ui-presets@v0.1.1
```

## 主题

六套主题均提供浅色与深色模式。**macOS、Liquid Glass、Sci‑Fi** 由 qiyuanhuakai 原创。**Retro、Mica、Acrylic** 延续上游主题样式。由 Web CSS 采集的原生预设也包含共用控件样式，整个预设模块沿用 GPL。单独安装 MIT 原生组件时，可以使用其标准主题或提供自己的样式。

## 组件对照

每张图左侧是 React WebView，右侧是 MyGo Native。四套主题与色彩模式组合各有 41 张图。输入框使用带内容的示例，简洁卡片和右键菜单展示悬浮后的状态。

剪贴板搜索与类型筛选图使用 DezirClip 主界面的同一组件。一个完整的输入框，下方是七种内容类型。可以单独复用 React 的 ClipboardSearch 或 Go 的 Kit.ClipboardSearch。高级搜索栏和搜索历史也提供独立示例。新增组件的源码与导出位于 `main` 分支。

<details open>
<summary>macOS 深色</summary>

| 组件 | 两套库的实测画面 |
| --- | --- |
| 应用顶栏 | ![应用顶栏 WebView 与 Native](docs/comparisons/macos-dark/app-header.png) |
| 剪贴板搜索与类型筛选 | ![剪贴板搜索与类型筛选 WebView 与 Native](docs/comparisons/macos-dark/clipboard-search.png) |
| 虚拟历史列表 | ![虚拟历史列表 WebView 与 Native](docs/comparisons/macos-dark/history-list.png) |
| 文本卡片与字号 | ![小字号、长文本和隐私遮挡卡片](docs/comparisons/macos-dark/history-typography.png) |
| 紧凑卡片与大字号 | ![十八像素字号下的紧凑卡片](docs/comparisons/macos-dark/history-typography-compact.png) |
| 卡片自定义字体 | ![界面字体与等宽字体设置的实际效果](docs/comparisons/macos-dark/history-typography-custom.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/macos-dark/history-card.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/macos-dark/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/macos-dark/history-card-compact-actions.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/macos-dark/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/macos-dark/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/macos-dark/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/macos-dark/settings-section-card.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/macos-dark/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/macos-dark/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/macos-dark/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/macos-dark/search-select-empty.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/macos-dark/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/macos-dark/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/macos-dark/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/macos-dark/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/macos-dark/privacy-rules.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/macos-dark/emoji-favorites.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/macos-dark/settings-data.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/macos-dark/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/macos-dark/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/macos-dark/backup-dialog-import.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/macos-dark/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/macos-dark/settings-update-installing.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/macos-dark/tag-toolbar.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/macos-dark/tag-search.png) |
| 独立搜索历史 | ![独立搜索历史 WebView 与 Native](docs/comparisons/macos-dark/search-history.png) |
| 筛选项 | ![筛选项 WebView 与 Native](docs/comparisons/macos-dark/filters.png) |
| 匹配高亮 | ![匹配高亮 WebView 与 Native](docs/comparisons/macos-dark/highlight.png) |
| 下拉选择 | ![下拉选择 WebView 与 Native](docs/comparisons/macos-dark/select.png) |
| 设置导航 | ![设置导航 WebView 与 Native](docs/comparisons/macos-dark/settings-rail.png) |
| 设置页脚 | ![设置页脚 WebView 与 Native](docs/comparisons/macos-dark/settings-footer.png) |
| 设置折叠区 | ![设置折叠区 WebView 与 Native](docs/comparisons/macos-dark/disclosure.png) |
| 设置帮助 | ![设置帮助 WebView 与 Native](docs/comparisons/macos-dark/help.png) |
| 确认对话框 | ![确认对话框 WebView 与 Native](docs/comparisons/macos-dark/confirm.png) |
| 标签对话框 | ![标签对话框 WebView 与 Native](docs/comparisons/macos-dark/tag-dialog.png) |
| 短提示 | ![短提示 WebView 与 Native](docs/comparisons/macos-dark/toast.png) |
| 进度提示 | ![进度提示 WebView 与 Native](docs/comparisons/macos-dark/progress.png) |
| 附件预览 | ![附件预览 WebView 与 Native](docs/comparisons/macos-dark/attachment-preview.png) |
| 紧凑附件卡片 | ![紧凑模式的附件缩略图和文件名](docs/comparisons/macos-dark/attachment-preview-compact.png) |
| 窄窗附件卡片 | ![窄窗下的附件列数和省略提示](docs/comparisons/macos-dark/attachment-preview-narrow.png) |
| 图片卡片 | ![不同原始尺寸的图片卡片](docs/comparisons/macos-dark/history-image.png) |
| 紧凑图片卡片 | ![紧凑模式的缩略图尺寸](docs/comparisons/macos-dark/history-image-compact.png) |
| 图片预览不可用 | ![图片不可用时的卡片状态](docs/comparisons/macos-dark/history-image-unavailable.png) |
| 预览工具栏 | ![预览工具栏 WebView 与 Native](docs/comparisons/macos-dark/preview-toolbar.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/macos-dark/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/macos-dark/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/macos-dark/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/macos-dark/settings-toggle.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/macos-dark/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/macos-dark/hotkey-binding.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/macos-dark/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/macos-dark/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/macos-dark/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/macos-dark/emoji-grid.png) |
| 独立高级搜索栏 | ![独立高级搜索栏 WebView 与 Native](docs/comparisons/macos-dark/search-bar.png) |

</details>

<details>
<summary>macOS 浅色</summary>

| 组件 | 两套库的实测画面 |
| --- | --- |
| 应用顶栏 | ![应用顶栏 WebView 与 Native](docs/comparisons/app-header.png) |
| 剪贴板搜索与类型筛选 | ![剪贴板搜索与类型筛选 WebView 与 Native](docs/comparisons/clipboard-search.png) |
| 虚拟历史列表 | ![虚拟历史列表 WebView 与 Native](docs/comparisons/history-list.png) |
| 文本卡片与字号 | ![小字号、长文本和隐私遮挡卡片](docs/comparisons/history-typography.png) |
| 紧凑卡片与大字号 | ![十八像素字号下的紧凑卡片](docs/comparisons/history-typography-compact.png) |
| 卡片自定义字体 | ![界面字体与等宽字体设置的实际效果](docs/comparisons/history-typography-custom.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/history-card.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/history-card-compact-actions.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/settings-section-card.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/search-select-empty.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/privacy-rules.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/emoji-favorites.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/settings-data.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/backup-dialog-import.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/settings-update-installing.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/tag-toolbar.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/tag-search.png) |
| 独立搜索历史 | ![独立搜索历史 WebView 与 Native](docs/comparisons/search-history.png) |
| 筛选项 | ![筛选项 WebView 与 Native](docs/comparisons/filters.png) |
| 匹配高亮 | ![匹配高亮 WebView 与 Native](docs/comparisons/highlight.png) |
| 下拉选择 | ![下拉选择 WebView 与 Native](docs/comparisons/select.png) |
| 设置导航 | ![设置导航 WebView 与 Native](docs/comparisons/settings-rail.png) |
| 设置页脚 | ![设置页脚 WebView 与 Native](docs/comparisons/settings-footer.png) |
| 设置折叠区 | ![设置折叠区 WebView 与 Native](docs/comparisons/disclosure.png) |
| 设置帮助 | ![设置帮助 WebView 与 Native](docs/comparisons/help.png) |
| 确认对话框 | ![确认对话框 WebView 与 Native](docs/comparisons/confirm.png) |
| 标签对话框 | ![标签对话框 WebView 与 Native](docs/comparisons/tag-dialog.png) |
| 短提示 | ![短提示 WebView 与 Native](docs/comparisons/toast.png) |
| 进度提示 | ![进度提示 WebView 与 Native](docs/comparisons/progress.png) |
| 附件预览 | ![附件预览 WebView 与 Native](docs/comparisons/attachment-preview.png) |
| 紧凑附件卡片 | ![紧凑模式的附件缩略图和文件名](docs/comparisons/attachment-preview-compact.png) |
| 窄窗附件卡片 | ![窄窗下的附件列数和省略提示](docs/comparisons/attachment-preview-narrow.png) |
| 图片卡片 | ![不同原始尺寸的图片卡片](docs/comparisons/history-image.png) |
| 紧凑图片卡片 | ![紧凑模式的缩略图尺寸](docs/comparisons/history-image-compact.png) |
| 图片预览不可用 | ![图片不可用时的卡片状态](docs/comparisons/history-image-unavailable.png) |
| 预览工具栏 | ![预览工具栏 WebView 与 Native](docs/comparisons/preview-toolbar.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/settings-toggle.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/hotkey-binding.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/emoji-grid.png) |
| 独立高级搜索栏 | ![独立高级搜索栏 WebView 与 Native](docs/comparisons/search-bar.png) |

</details>

<details>
<summary>Liquid Glass 浅色</summary>

| 组件 | 两套库的实测画面 |
| --- | --- |
| 应用顶栏 | ![应用顶栏 WebView 与 Native](docs/comparisons/liquid-glass-light/app-header.png) |
| 剪贴板搜索与类型筛选 | ![剪贴板搜索与类型筛选 WebView 与 Native](docs/comparisons/liquid-glass-light/clipboard-search.png) |
| 虚拟历史列表 | ![虚拟历史列表 WebView 与 Native](docs/comparisons/liquid-glass-light/history-list.png) |
| 文本卡片与字号 | ![小字号、长文本和隐私遮挡卡片](docs/comparisons/liquid-glass-light/history-typography.png) |
| 紧凑卡片与大字号 | ![十八像素字号下的紧凑卡片](docs/comparisons/liquid-glass-light/history-typography-compact.png) |
| 卡片自定义字体 | ![界面字体与等宽字体设置的实际效果](docs/comparisons/liquid-glass-light/history-typography-custom.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/liquid-glass-light/history-card.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/liquid-glass-light/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/liquid-glass-light/history-card-compact-actions.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-light/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-light/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/liquid-glass-light/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-section-card.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/liquid-glass-light/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/liquid-glass-light/search-select-empty.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/liquid-glass-light/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/liquid-glass-light/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/liquid-glass-light/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/liquid-glass-light/privacy-rules.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/liquid-glass-light/emoji-favorites.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-data.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/liquid-glass-light/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/liquid-glass-light/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/liquid-glass-light/backup-dialog-import.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/liquid-glass-light/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/liquid-glass-light/settings-update-installing.png) |
| 透明度 20% | ![透明度 20% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-light/surface-opacity-20.png) |
| 透明度 85% | ![透明度 85% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-light/surface-opacity-85.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/liquid-glass-light/tag-toolbar.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/liquid-glass-light/tag-search.png) |
| 独立搜索历史 | ![独立搜索历史 WebView 与 Native](docs/comparisons/liquid-glass-light/search-history.png) |
| 筛选项 | ![筛选项 WebView 与 Native](docs/comparisons/liquid-glass-light/filters.png) |
| 匹配高亮 | ![匹配高亮 WebView 与 Native](docs/comparisons/liquid-glass-light/highlight.png) |
| 下拉选择 | ![下拉选择 WebView 与 Native](docs/comparisons/liquid-glass-light/select.png) |
| 设置导航 | ![设置导航 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-rail.png) |
| 设置页脚 | ![设置页脚 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-footer.png) |
| 设置折叠区 | ![设置折叠区 WebView 与 Native](docs/comparisons/liquid-glass-light/disclosure.png) |
| 设置帮助 | ![设置帮助 WebView 与 Native](docs/comparisons/liquid-glass-light/help.png) |
| 确认对话框 | ![确认对话框 WebView 与 Native](docs/comparisons/liquid-glass-light/confirm.png) |
| 标签对话框 | ![标签对话框 WebView 与 Native](docs/comparisons/liquid-glass-light/tag-dialog.png) |
| 短提示 | ![短提示 WebView 与 Native](docs/comparisons/liquid-glass-light/toast.png) |
| 进度提示 | ![进度提示 WebView 与 Native](docs/comparisons/liquid-glass-light/progress.png) |
| 附件预览 | ![附件预览 WebView 与 Native](docs/comparisons/liquid-glass-light/attachment-preview.png) |
| 紧凑附件卡片 | ![紧凑模式的附件缩略图和文件名](docs/comparisons/liquid-glass-light/attachment-preview-compact.png) |
| 窄窗附件卡片 | ![窄窗下的附件列数和省略提示](docs/comparisons/liquid-glass-light/attachment-preview-narrow.png) |
| 图片卡片 | ![不同原始尺寸的图片卡片](docs/comparisons/liquid-glass-light/history-image.png) |
| 紧凑图片卡片 | ![紧凑模式的缩略图尺寸](docs/comparisons/liquid-glass-light/history-image-compact.png) |
| 图片预览不可用 | ![图片不可用时的卡片状态](docs/comparisons/liquid-glass-light/history-image-unavailable.png) |
| 预览工具栏 | ![预览工具栏 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-toolbar.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/liquid-glass-light/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-toggle.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/liquid-glass-light/hotkey-binding.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/liquid-glass-light/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/liquid-glass-light/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/liquid-glass-light/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/liquid-glass-light/emoji-grid.png) |
| 独立高级搜索栏 | ![独立高级搜索栏 WebView 与 Native](docs/comparisons/liquid-glass-light/search-bar.png) |

</details>

<details>
<summary>Liquid Glass 深色</summary>

| 组件 | 两套库的实测画面 |
| --- | --- |
| 应用顶栏 | ![应用顶栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/app-header.png) |
| 剪贴板搜索与类型筛选 | ![剪贴板搜索与类型筛选 WebView 与 Native](docs/comparisons/liquid-glass-dark/clipboard-search.png) |
| 虚拟历史列表 | ![虚拟历史列表 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-list.png) |
| 文本卡片与字号 | ![小字号、长文本和隐私遮挡卡片](docs/comparisons/liquid-glass-dark/history-typography.png) |
| 紧凑卡片与大字号 | ![十八像素字号下的紧凑卡片](docs/comparisons/liquid-glass-dark/history-typography-compact.png) |
| 卡片自定义字体 | ![界面字体与等宽字体设置的实际效果](docs/comparisons/liquid-glass-dark/history-typography-custom.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-card.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-card-compact-actions.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-dark/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-dark/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/liquid-glass-dark/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-section-card.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/liquid-glass-dark/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/liquid-glass-dark/search-select-empty.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/liquid-glass-dark/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/liquid-glass-dark/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/liquid-glass-dark/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/liquid-glass-dark/privacy-rules.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/liquid-glass-dark/emoji-favorites.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-data.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/liquid-glass-dark/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/liquid-glass-dark/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/liquid-glass-dark/backup-dialog-import.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/liquid-glass-dark/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/liquid-glass-dark/settings-update-installing.png) |
| 透明度 20% | ![透明度 20% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-dark/surface-opacity-20.png) |
| 透明度 85% | ![透明度 85% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-dark/surface-opacity-85.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/tag-toolbar.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/liquid-glass-dark/tag-search.png) |
| 独立搜索历史 | ![独立搜索历史 WebView 与 Native](docs/comparisons/liquid-glass-dark/search-history.png) |
| 筛选项 | ![筛选项 WebView 与 Native](docs/comparisons/liquid-glass-dark/filters.png) |
| 匹配高亮 | ![匹配高亮 WebView 与 Native](docs/comparisons/liquid-glass-dark/highlight.png) |
| 下拉选择 | ![下拉选择 WebView 与 Native](docs/comparisons/liquid-glass-dark/select.png) |
| 设置导航 | ![设置导航 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-rail.png) |
| 设置页脚 | ![设置页脚 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-footer.png) |
| 设置折叠区 | ![设置折叠区 WebView 与 Native](docs/comparisons/liquid-glass-dark/disclosure.png) |
| 设置帮助 | ![设置帮助 WebView 与 Native](docs/comparisons/liquid-glass-dark/help.png) |
| 确认对话框 | ![确认对话框 WebView 与 Native](docs/comparisons/liquid-glass-dark/confirm.png) |
| 标签对话框 | ![标签对话框 WebView 与 Native](docs/comparisons/liquid-glass-dark/tag-dialog.png) |
| 短提示 | ![短提示 WebView 与 Native](docs/comparisons/liquid-glass-dark/toast.png) |
| 进度提示 | ![进度提示 WebView 与 Native](docs/comparisons/liquid-glass-dark/progress.png) |
| 附件预览 | ![附件预览 WebView 与 Native](docs/comparisons/liquid-glass-dark/attachment-preview.png) |
| 紧凑附件卡片 | ![紧凑模式的附件缩略图和文件名](docs/comparisons/liquid-glass-dark/attachment-preview-compact.png) |
| 窄窗附件卡片 | ![窄窗下的附件列数和省略提示](docs/comparisons/liquid-glass-dark/attachment-preview-narrow.png) |
| 图片卡片 | ![不同原始尺寸的图片卡片](docs/comparisons/liquid-glass-dark/history-image.png) |
| 紧凑图片卡片 | ![紧凑模式的缩略图尺寸](docs/comparisons/liquid-glass-dark/history-image-compact.png) |
| 图片预览不可用 | ![图片不可用时的卡片状态](docs/comparisons/liquid-glass-dark/history-image-unavailable.png) |
| 预览工具栏 | ![预览工具栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-toolbar.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/liquid-glass-dark/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-toggle.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/liquid-glass-dark/hotkey-binding.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/liquid-glass-dark/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/liquid-glass-dark/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/liquid-glass-dark/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/liquid-glass-dark/emoji-grid.png) |
| 独立高级搜索栏 | ![独立高级搜索栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/search-bar.png) |

</details>

## 构建与验证

```sh
pnpm install --frozen-lockfile
pnpm exec playwright install chromium --only-shell
pnpm run build:react
go -C packages/ui-native test ./...
go -C packages/ui-presets test ./...
go -C packages/ui-presets build -o ../../dzc-ui-gallery ./cmd/dzc-ui-gallery
node scripts/compare-ui-components.mjs native=./dzc-ui-gallery examples=all
node scripts/write-ui-comparison-images.mjs 'results=/path/to/result-directory'
node scripts/write-ui-comparison-images.mjs 'results=/path/to/result-directory' theme=macos mode=dark output=docs/comparisons/macos-dark
node scripts/write-ui-comparison-images.mjs 'results=/path/to/result-directory' theme=liquid-glass mode=light output=docs/comparisons/liquid-glass-light
node scripts/write-ui-comparison-images.mjs 'results=/path/to/result-directory' theme=liquid-glass mode=dark output=docs/comparisons/liquid-glass-dark
```

将 `results` 改为对照脚本输出的目录。Windows 下为原生图集程序使用 `dzc-ui-gallery.exe`。

查看 [许可证范围与第三方声明](LICENSE.md)。
