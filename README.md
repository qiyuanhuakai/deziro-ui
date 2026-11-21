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

每张图左侧是 React WebView，右侧是 MyGo Native。macOS 与 Liquid Glass 的浅色和深色共 288 张对照图。输入框使用带内容的示例，简洁卡片操作栏和右键菜单展示悬浮后的状态。

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
| 普通卡片中的富文本 | ![带粗体、斜体和链接的普通卡片](docs/comparisons/macos-dark/history-rich-text.png) |
| 紧凑卡片中的富文本 | ![紧凑模式下的富文本截断](docs/comparisons/macos-dark/history-rich-text-compact.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/macos-dark/history-card.png) |
| 图片卡片的 OCR 状态和识别结果 | ![普通卡片展开 OCR 结果](docs/comparisons/macos-dark/history-ocr.png) |
| 简洁卡片的 OCR 结果 | ![简洁卡片展开 OCR 结果](docs/comparisons/macos-dark/history-ocr-compact.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/macos-dark/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/macos-dark/history-card-compact-actions.png) |
| 剪贴板卡片标签 | ![标签的颜色、文字与圆角](docs/comparisons/macos-dark/history-tags.png) |
| 简洁卡片标签 | ![覆盖在卡片右下角的标签](docs/comparisons/macos-dark/history-tags-compact.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/macos-dark/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/macos-dark/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/macos-dark/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/macos-dark/settings-section-card.png) |
| 设置项、说明和数字输入 | ![字体说明和数字输入框](docs/comparisons/macos-dark/settings-field.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/macos-dark/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/macos-dark/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/macos-dark/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/macos-dark/search-select-empty.png) |
| 默认应用搜索容错 | ![程序名称拼写错误时仍返回正确选项](docs/comparisons/macos-dark/search-select-typo.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/macos-dark/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/macos-dark/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/macos-dark/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/macos-dark/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/macos-dark/privacy-rules.png) |
| 应用排除输入框 | ![输入框获得焦点时的主题效果](docs/comparisons/macos-dark/privacy-source-exclusion-focus.png) |
| 隐私检测开关与自定义规则 | ![窄窗下的检测开关换行与拖动调高后的规则输入框](docs/comparisons/macos-dark/privacy-detection.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/macos-dark/emoji-favorites.png) |
| 空表情收藏 | ![空收藏提示和居中的添加按钮](docs/comparisons/macos-dark/emoji-favorites-empty.png) |
| 添加表情悬停 | ![添加按钮悬停后的轮廓和位移](docs/comparisons/macos-dark/emoji-favorites-hover.png) |
| 拖入表情图片 | ![拖入图片时的收藏区域提示](docs/comparisons/macos-dark/emoji-favorites-drop.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/macos-dark/settings-data.png) |
| 数据目录悬浮 | ![数据路径的主题悬浮样式](docs/comparisons/macos-dark/settings-data-hover.png) |
| 备份入口悬浮 | ![导出按钮的主题悬浮样式](docs/comparisons/macos-dark/settings-data-action-hover.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/macos-dark/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/macos-dark/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/macos-dark/backup-dialog-import.png) |
| 确认覆盖导入 | ![勾选后启用覆盖导入](docs/comparisons/macos-dark/backup-dialog-replace.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/macos-dark/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/macos-dark/settings-update-installing.png) |
| 更新操作 | ![检查、安装和发布页](docs/comparisons/macos-dark/settings-update-actions.png) |
| 下载中的更新操作 | ![下载时的按钮状态](docs/comparisons/macos-dark/settings-update-actions-busy.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/macos-dark/tag-toolbar.png) |
| 标签批量选择栏 | ![清除选择、已选数量和删除](docs/comparisons/macos-dark/tag-batch-toolbar.png) |
| 标签卡片批量选择 | ![独立复选框与选中项](docs/comparisons/macos-dark/tag-history-selection.png) |
| 简洁标签卡片批量选择 | ![简洁卡片的选中状态](docs/comparisons/macos-dark/tag-history-selection-compact.png) |
| 快速粘贴面板 | ![65% 透明度设置下的面板、搜索框与选中结果](docs/comparisons/macos-dark/quick-paste-panel.png) |
| 主题管理 | ![内置主题、已导入主题及导出和移除操作](docs/comparisons/macos-dark/theme-manager.png) |
| 自定义背景 | ![更换和清除背景、文件名与透明度](docs/comparisons/macos-dark/background-settings.png) |
| 截图操作 | ![区域截图和全屏截图](docs/comparisons/macos-dark/capture-actions.png) |
| OCR 引擎状态 | ![OCR 可用状态与重试按钮](docs/comparisons/macos-dark/capture-engine.png) |
| 快速粘贴界面设置 | ![固定的 MyGo Native 界面](docs/comparisons/macos-dark/fixed-capture-renderer.png) |
| CLI 设置 | ![已安装 CLI、PATH 副本冲突与随包指南](docs/comparisons/macos-dark/cli-settings.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/macos-dark/tag-search.png) |
| 搜索框标签建议 | ![标签建议、键盘选中态和主题色](docs/comparisons/macos-dark/tag-suggestions.png) |
| 标签建议空状态 | ![标签建议为空时的提示与重试](docs/comparisons/macos-dark/tag-suggestions-empty.png) |
| 设置栏长标题 | ![设置侧栏的长标题换行](docs/comparisons/macos-dark/settings-rail-long-label.png) |
| 字体选择 | ![下拉选项使用各自的字体](docs/comparisons/macos-dark/font-select.png) |
| 文本预览字号 | ![十八像素字号下的双行文本预览](docs/comparisons/macos-dark/preview-text.png) |
| 默认应用长名称 | ![默认应用菜单内的长名称换行](docs/comparisons/macos-dark/search-select-long-options.png) |
| 代码沿用文本程序 | ![代码选择器显示沿用的文本程序](docs/comparisons/macos-dark/search-select-app-inherited.png) |
| 同名程序 | ![默认应用菜单用路径区分同名程序](docs/comparisons/macos-dark/search-select-app-same-name.png) |
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
| 固定预览工具栏 | ![固定预览工具栏 WebView 与 Native](docs/comparisons/macos-dark/preview-toolbar.png) |
| 悬浮预览工具栏 | ![悬浮预览工具栏 WebView 与 Native](docs/comparisons/macos-dark/preview-toolbar-hover.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/macos-dark/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/macos-dark/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/macos-dark/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/macos-dark/settings-toggle.png) |
| 带说明的设置开关 | ![粘贴后置顶与 Win+V 开关](docs/comparisons/macos-dark/settings-toggle-help.png) |
| 方向键预览从属设置 | ![方向键选择和预览开关](docs/comparisons/macos-dark/settings-keyboard-preview.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/macos-dark/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/macos-dark/hotkey-binding.png) |
| 在说明中清除可选快捷键 | ![单快捷键的清除操作](docs/comparisons/macos-dark/hotkey-optional.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/macos-dark/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/macos-dark/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/macos-dark/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/macos-dark/emoji-grid.png) |
| 滚动时自动显示后续符号 | ![滚动到底部后的符号面板](docs/comparisons/macos-dark/character-panel-scroll.png) |
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
| 普通卡片中的富文本 | ![带粗体、斜体和链接的普通卡片](docs/comparisons/history-rich-text.png) |
| 紧凑卡片中的富文本 | ![紧凑模式下的富文本截断](docs/comparisons/history-rich-text-compact.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/history-card.png) |
| 图片卡片的 OCR 状态和识别结果 | ![普通卡片展开 OCR 结果](docs/comparisons/history-ocr.png) |
| 简洁卡片的 OCR 结果 | ![简洁卡片展开 OCR 结果](docs/comparisons/history-ocr-compact.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/history-card-compact-actions.png) |
| 剪贴板卡片标签 | ![标签的颜色、文字与圆角](docs/comparisons/history-tags.png) |
| 简洁卡片标签 | ![覆盖在卡片右下角的标签](docs/comparisons/history-tags-compact.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/settings-section-card.png) |
| 设置项、说明和数字输入 | ![字体说明和数字输入框](docs/comparisons/settings-field.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/search-select-empty.png) |
| 默认应用搜索容错 | ![程序名称拼写错误时仍返回正确选项](docs/comparisons/search-select-typo.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/privacy-rules.png) |
| 应用排除输入框 | ![输入框获得焦点时的主题效果](docs/comparisons/privacy-source-exclusion-focus.png) |
| 隐私检测开关与自定义规则 | ![窄窗下的检测开关换行与拖动调高后的规则输入框](docs/comparisons/privacy-detection.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/emoji-favorites.png) |
| 空表情收藏 | ![空收藏提示和居中的添加按钮](docs/comparisons/emoji-favorites-empty.png) |
| 添加表情悬停 | ![添加按钮悬停后的轮廓和位移](docs/comparisons/emoji-favorites-hover.png) |
| 拖入表情图片 | ![拖入图片时的收藏区域提示](docs/comparisons/emoji-favorites-drop.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/settings-data.png) |
| 数据目录悬浮 | ![数据路径的主题悬浮样式](docs/comparisons/settings-data-hover.png) |
| 备份入口悬浮 | ![导出按钮的主题悬浮样式](docs/comparisons/settings-data-action-hover.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/backup-dialog-import.png) |
| 确认覆盖导入 | ![勾选后启用覆盖导入](docs/comparisons/backup-dialog-replace.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/settings-update-installing.png) |
| 更新操作 | ![检查、安装和发布页](docs/comparisons/settings-update-actions.png) |
| 下载中的更新操作 | ![下载时的按钮状态](docs/comparisons/settings-update-actions-busy.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/tag-toolbar.png) |
| 标签批量选择栏 | ![清除选择、已选数量和删除](docs/comparisons/tag-batch-toolbar.png) |
| 标签卡片批量选择 | ![独立复选框与选中项](docs/comparisons/tag-history-selection.png) |
| 简洁标签卡片批量选择 | ![简洁卡片的选中状态](docs/comparisons/tag-history-selection-compact.png) |
| 快速粘贴面板 | ![65% 透明度设置下的面板、搜索框与选中结果](docs/comparisons/quick-paste-panel.png) |
| 主题管理 | ![内置主题、已导入主题及导出和移除操作](docs/comparisons/theme-manager.png) |
| 自定义背景 | ![更换和清除背景、文件名与透明度](docs/comparisons/background-settings.png) |
| 截图操作 | ![区域截图和全屏截图](docs/comparisons/capture-actions.png) |
| OCR 引擎状态 | ![OCR 可用状态与重试按钮](docs/comparisons/capture-engine.png) |
| 快速粘贴界面设置 | ![固定的 MyGo Native 界面](docs/comparisons/fixed-capture-renderer.png) |
| CLI 设置 | ![已安装 CLI、PATH 副本冲突与随包指南](docs/comparisons/cli-settings.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/tag-search.png) |
| 搜索框标签建议 | ![标签建议、键盘选中态和主题色](docs/comparisons/tag-suggestions.png) |
| 标签建议空状态 | ![标签建议为空时的提示与重试](docs/comparisons/tag-suggestions-empty.png) |
| 设置栏长标题 | ![设置侧栏的长标题换行](docs/comparisons/settings-rail-long-label.png) |
| 字体选择 | ![下拉选项使用各自的字体](docs/comparisons/font-select.png) |
| 文本预览字号 | ![十八像素字号下的双行文本预览](docs/comparisons/preview-text.png) |
| 默认应用长名称 | ![默认应用菜单内的长名称换行](docs/comparisons/search-select-long-options.png) |
| 代码沿用文本程序 | ![代码选择器显示沿用的文本程序](docs/comparisons/search-select-app-inherited.png) |
| 同名程序 | ![默认应用菜单用路径区分同名程序](docs/comparisons/search-select-app-same-name.png) |
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
| 固定预览工具栏 | ![固定预览工具栏 WebView 与 Native](docs/comparisons/preview-toolbar.png) |
| 悬浮预览工具栏 | ![悬浮预览工具栏 WebView 与 Native](docs/comparisons/preview-toolbar-hover.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/settings-toggle.png) |
| 带说明的设置开关 | ![粘贴后置顶与 Win+V 开关](docs/comparisons/settings-toggle-help.png) |
| 方向键预览从属设置 | ![方向键选择和预览开关](docs/comparisons/settings-keyboard-preview.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/hotkey-binding.png) |
| 在说明中清除可选快捷键 | ![单快捷键的清除操作](docs/comparisons/hotkey-optional.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/emoji-grid.png) |
| 滚动时自动显示后续符号 | ![滚动到底部后的符号面板](docs/comparisons/character-panel-scroll.png) |
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
| 普通卡片中的富文本 | ![带粗体、斜体和链接的普通卡片](docs/comparisons/liquid-glass-light/history-rich-text.png) |
| 紧凑卡片中的富文本 | ![紧凑模式下的富文本截断](docs/comparisons/liquid-glass-light/history-rich-text-compact.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/liquid-glass-light/history-card.png) |
| 图片卡片的 OCR 状态和识别结果 | ![普通卡片展开 OCR 结果](docs/comparisons/liquid-glass-light/history-ocr.png) |
| 简洁卡片的 OCR 结果 | ![简洁卡片展开 OCR 结果](docs/comparisons/liquid-glass-light/history-ocr-compact.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/liquid-glass-light/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/liquid-glass-light/history-card-compact-actions.png) |
| 剪贴板卡片标签 | ![标签的颜色、文字与圆角](docs/comparisons/liquid-glass-light/history-tags.png) |
| 简洁卡片标签 | ![覆盖在卡片右下角的标签](docs/comparisons/liquid-glass-light/history-tags-compact.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-light/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-light/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/liquid-glass-light/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-section-card.png) |
| 设置项、说明和数字输入 | ![字体说明和数字输入框](docs/comparisons/liquid-glass-light/settings-field.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/liquid-glass-light/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/liquid-glass-light/search-select-empty.png) |
| 默认应用搜索容错 | ![程序名称拼写错误时仍返回正确选项](docs/comparisons/liquid-glass-light/search-select-typo.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/liquid-glass-light/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/liquid-glass-light/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/liquid-glass-light/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/liquid-glass-light/privacy-rules.png) |
| 应用排除输入框 | ![输入框获得焦点时的主题效果](docs/comparisons/liquid-glass-light/privacy-source-exclusion-focus.png) |
| 隐私检测开关与自定义规则 | ![窄窗下的检测开关换行与拖动调高后的规则输入框](docs/comparisons/liquid-glass-light/privacy-detection.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/liquid-glass-light/emoji-favorites.png) |
| 空表情收藏 | ![空收藏提示和居中的添加按钮](docs/comparisons/liquid-glass-light/emoji-favorites-empty.png) |
| 添加表情悬停 | ![添加按钮悬停后的轮廓和位移](docs/comparisons/liquid-glass-light/emoji-favorites-hover.png) |
| 拖入表情图片 | ![拖入图片时的收藏区域提示](docs/comparisons/liquid-glass-light/emoji-favorites-drop.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-data.png) |
| 数据目录悬浮 | ![数据路径的主题悬浮样式](docs/comparisons/liquid-glass-light/settings-data-hover.png) |
| 备份入口悬浮 | ![导出按钮的主题悬浮样式](docs/comparisons/liquid-glass-light/settings-data-action-hover.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/liquid-glass-light/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/liquid-glass-light/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/liquid-glass-light/backup-dialog-import.png) |
| 确认覆盖导入 | ![勾选后启用覆盖导入](docs/comparisons/liquid-glass-light/backup-dialog-replace.png) |
| 玻璃弹窗背景 | ![浅色玻璃下仍可辨认的设置页内容](docs/comparisons/liquid-glass-light/backup-dialog-glass.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/liquid-glass-light/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/liquid-glass-light/settings-update-installing.png) |
| 更新操作 | ![检查、安装和发布页](docs/comparisons/liquid-glass-light/settings-update-actions.png) |
| 下载中的更新操作 | ![下载时的按钮状态](docs/comparisons/liquid-glass-light/settings-update-actions-busy.png) |
| 透明度 20% | ![透明度 20% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-light/surface-opacity-20.png) |
| 透明度 85% | ![透明度 85% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-light/surface-opacity-85.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/liquid-glass-light/tag-toolbar.png) |
| 标签批量选择栏 | ![清除选择、已选数量和删除](docs/comparisons/liquid-glass-light/tag-batch-toolbar.png) |
| 标签卡片批量选择 | ![独立复选框与选中项](docs/comparisons/liquid-glass-light/tag-history-selection.png) |
| 简洁标签卡片批量选择 | ![简洁卡片的选中状态](docs/comparisons/liquid-glass-light/tag-history-selection-compact.png) |
| 快速粘贴面板 | ![65% 透明度设置下的面板、搜索框与选中结果](docs/comparisons/liquid-glass-light/quick-paste-panel.png) |
| 主题管理 | ![内置主题、已导入主题及导出和移除操作](docs/comparisons/liquid-glass-light/theme-manager.png) |
| 自定义背景 | ![更换和清除背景、文件名与透明度](docs/comparisons/liquid-glass-light/background-settings.png) |
| 截图操作 | ![区域截图和全屏截图](docs/comparisons/liquid-glass-light/capture-actions.png) |
| OCR 引擎状态 | ![OCR 可用状态与重试按钮](docs/comparisons/liquid-glass-light/capture-engine.png) |
| 快速粘贴界面设置 | ![固定的 MyGo Native 界面](docs/comparisons/liquid-glass-light/fixed-capture-renderer.png) |
| CLI 设置 | ![已安装 CLI、PATH 副本冲突与随包指南](docs/comparisons/liquid-glass-light/cli-settings.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/liquid-glass-light/tag-search.png) |
| 搜索框标签建议 | ![标签建议、键盘选中态和主题色](docs/comparisons/liquid-glass-light/tag-suggestions.png) |
| 标签建议空状态 | ![标签建议为空时的提示与重试](docs/comparisons/liquid-glass-light/tag-suggestions-empty.png) |
| 设置栏长标题 | ![设置侧栏的长标题换行](docs/comparisons/liquid-glass-light/settings-rail-long-label.png) |
| 字体选择 | ![下拉选项使用各自的字体](docs/comparisons/liquid-glass-light/font-select.png) |
| 文本预览字号 | ![十八像素字号下的双行文本预览](docs/comparisons/liquid-glass-light/preview-text.png) |
| 默认应用长名称 | ![默认应用菜单内的长名称换行](docs/comparisons/liquid-glass-light/search-select-long-options.png) |
| 代码沿用文本程序 | ![代码选择器显示沿用的文本程序](docs/comparisons/liquid-glass-light/search-select-app-inherited.png) |
| 同名程序 | ![默认应用菜单用路径区分同名程序](docs/comparisons/liquid-glass-light/search-select-app-same-name.png) |
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
| 固定预览工具栏 | ![固定预览工具栏 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-toolbar.png) |
| 悬浮预览工具栏 | ![悬浮预览工具栏 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-toolbar-hover.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/liquid-glass-light/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/liquid-glass-light/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-toggle.png) |
| 带说明的设置开关 | ![粘贴后置顶与 Win+V 开关](docs/comparisons/liquid-glass-light/settings-toggle-help.png) |
| 方向键预览从属设置 | ![方向键选择和预览开关](docs/comparisons/liquid-glass-light/settings-keyboard-preview.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/liquid-glass-light/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/liquid-glass-light/hotkey-binding.png) |
| 在说明中清除可选快捷键 | ![单快捷键的清除操作](docs/comparisons/liquid-glass-light/hotkey-optional.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/liquid-glass-light/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/liquid-glass-light/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/liquid-glass-light/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/liquid-glass-light/emoji-grid.png) |
| 滚动时自动显示后续符号 | ![滚动到底部后的符号面板](docs/comparisons/liquid-glass-light/character-panel-scroll.png) |
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
| 普通卡片中的富文本 | ![带粗体、斜体和链接的普通卡片](docs/comparisons/liquid-glass-dark/history-rich-text.png) |
| 紧凑卡片中的富文本 | ![紧凑模式下的富文本截断](docs/comparisons/liquid-glass-dark/history-rich-text-compact.png) |
| 历史卡片 | ![历史卡片 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-card.png) |
| 图片卡片的 OCR 状态和识别结果 | ![普通卡片展开 OCR 结果](docs/comparisons/liquid-glass-dark/history-ocr.png) |
| 简洁卡片的 OCR 结果 | ![简洁卡片展开 OCR 结果](docs/comparisons/liquid-glass-dark/history-ocr-compact.png) |
| 历史卡片信息栏 | ![历史卡片信息栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-card-meta.png) |
| 简洁卡片悬浮操作栏 | ![简洁卡片悬浮操作栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/history-card-compact-actions.png) |
| 剪贴板卡片标签 | ![标签的颜色、文字与圆角](docs/comparisons/liquid-glass-dark/history-tags.png) |
| 简洁卡片标签 | ![覆盖在卡片右下角的标签](docs/comparisons/liquid-glass-dark/history-tags-compact.png) |
| 文字处理侧边子菜单 | ![文字处理侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-dark/context-menu.png) |
| 隐私侧边子菜单 | ![隐私侧边子菜单 WebView 与 Native](docs/comparisons/liquid-glass-dark/context-menu-privacy.png) |
| 文本、数字和多行输入框 | ![文本、数字和多行输入框 WebView 与 Native](docs/comparisons/liquid-glass-dark/text-fields.png) |
| 设置分组卡片 | ![设置分组卡片 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-section-card.png) |
| 设置项、说明和数字输入 | ![字体说明和数字输入框](docs/comparisons/liquid-glass-dark/settings-field.png) |
| 横屏设置分栏与通栏输入框 | ![横屏设置分栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-section-grid.png) |
| 设置选项卡片和更新通道 | ![设置选项卡片和更新通道 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-choices.png) |
| 字体与默认应用选择器 | ![字体选中值和应用搜索结果](docs/comparisons/liquid-glass-dark/search-select.png) |
| 默认应用搜索无匹配项 | ![应用选择器的无匹配项提示](docs/comparisons/liquid-glass-dark/search-select-empty.png) |
| 默认应用搜索容错 | ![程序名称拼写错误时仍返回正确选项](docs/comparisons/liquid-glass-dark/search-select-typo.png) |
| 多快捷键与移除按钮 | ![多快捷键与移除按钮 WebView 与 Native](docs/comparisons/liquid-glass-dark/hotkey-list.png) |
| Windows 接管 Win+V 时的快捷键列表 | ![Win+V 键帽和独立的快捷键移除按钮](docs/comparisons/liquid-glass-dark/hotkey-list-win-v.png) |
| 数字小键盘与标点快捷键 | ![主键区 1、小键盘 1 与分号的独立键帽](docs/comparisons/liquid-glass-dark/hotkey-list-extended.png) |
| 设置内删除按钮 | ![设置内删除按钮 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-remove.png) |
| 隐私规则卡片与应用排除输入框 | ![隐私规则行内确认和多行应用排除输入框](docs/comparisons/liquid-glass-dark/privacy-rules.png) |
| 应用排除输入框 | ![输入框获得焦点时的主题效果](docs/comparisons/liquid-glass-dark/privacy-source-exclusion-focus.png) |
| 隐私检测开关与自定义规则 | ![窄窗下的检测开关换行与拖动调高后的规则输入框](docs/comparisons/liquid-glass-dark/privacy-detection.png) |
| 表情收藏与添加 | ![表情收藏与添加 WebView 与 Native](docs/comparisons/liquid-glass-dark/emoji-favorites.png) |
| 空表情收藏 | ![空收藏提示和居中的添加按钮](docs/comparisons/liquid-glass-dark/emoji-favorites-empty.png) |
| 添加表情悬停 | ![添加按钮悬停后的轮廓和位移](docs/comparisons/liquid-glass-dark/emoji-favorites-hover.png) |
| 拖入表情图片 | ![拖入图片时的收藏区域提示](docs/comparisons/liquid-glass-dark/emoji-favorites-drop.png) |
| 数据路径和导入导出 | ![数据路径和导入导出 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-data.png) |
| 数据目录悬浮 | ![数据路径的主题悬浮样式](docs/comparisons/liquid-glass-dark/settings-data-hover.png) |
| 备份入口悬浮 | ![导出按钮的主题悬浮样式](docs/comparisons/liquid-glass-dark/settings-data-action-hover.png) |
| 备份导出 | ![导出格式、文件路径和操作按钮](docs/comparisons/liquid-glass-dark/backup-dialog.png) |
| 加密备份导出 | ![加密备份的密码和确认输入框](docs/comparisons/liquid-glass-dark/backup-dialog-encrypted.png) |
| 备份导入 | ![导入方式、文件路径和内容预览](docs/comparisons/liquid-glass-dark/backup-dialog-import.png) |
| 确认覆盖导入 | ![勾选后启用覆盖导入](docs/comparisons/liquid-glass-dark/backup-dialog-replace.png) |
| 玻璃弹窗背景 | ![深色玻璃下仍可辨认的设置页内容](docs/comparisons/liquid-glass-dark/backup-dialog-glass.png) |
| 更新下载进度 | ![当前版本、通道版本和下载百分比](docs/comparisons/liquid-glass-dark/settings-update.png) |
| 更新签名与安装状态 | ![验证通过后的安装和重启状态](docs/comparisons/liquid-glass-dark/settings-update-installing.png) |
| 更新操作 | ![检查、安装和发布页](docs/comparisons/liquid-glass-dark/settings-update-actions.png) |
| 下载中的更新操作 | ![下载时的按钮状态](docs/comparisons/liquid-glass-dark/settings-update-actions-busy.png) |
| 透明度 20% | ![透明度 20% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-dark/surface-opacity-20.png) |
| 透明度 85% | ![透明度 85% 下的设置卡片、下拉菜单和历史卡片](docs/comparisons/liquid-glass-dark/surface-opacity-85.png) |
| 标签管理操作栏 | ![标签管理操作栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/tag-toolbar.png) |
| 标签批量选择栏 | ![清除选择、已选数量和删除](docs/comparisons/liquid-glass-dark/tag-batch-toolbar.png) |
| 标签卡片批量选择 | ![独立复选框与选中项](docs/comparisons/liquid-glass-dark/tag-history-selection.png) |
| 简洁标签卡片批量选择 | ![简洁卡片的选中状态](docs/comparisons/liquid-glass-dark/tag-history-selection-compact.png) |
| 快速粘贴面板 | ![65% 透明度设置下的面板、搜索框与选中结果](docs/comparisons/liquid-glass-dark/quick-paste-panel.png) |
| 主题管理 | ![内置主题、已导入主题及导出和移除操作](docs/comparisons/liquid-glass-dark/theme-manager.png) |
| 自定义背景 | ![更换和清除背景、文件名与透明度](docs/comparisons/liquid-glass-dark/background-settings.png) |
| 截图操作 | ![区域截图和全屏截图](docs/comparisons/liquid-glass-dark/capture-actions.png) |
| OCR 引擎状态 | ![OCR 可用状态与重试按钮](docs/comparisons/liquid-glass-dark/capture-engine.png) |
| 快速粘贴界面设置 | ![固定的 MyGo Native 界面](docs/comparisons/liquid-glass-dark/fixed-capture-renderer.png) |
| CLI 设置 | ![已安装 CLI、PATH 副本冲突与随包指南](docs/comparisons/liquid-glass-dark/cli-settings.png) |
| 标签选择器与搜索框 | ![窄屏标签选择、侧栏搜索与条目搜索](docs/comparisons/liquid-glass-dark/tag-search.png) |
| 搜索框标签建议 | ![标签建议、键盘选中态和主题色](docs/comparisons/liquid-glass-dark/tag-suggestions.png) |
| 标签建议空状态 | ![标签建议为空时的提示与重试](docs/comparisons/liquid-glass-dark/tag-suggestions-empty.png) |
| 设置栏长标题 | ![设置侧栏的长标题换行](docs/comparisons/liquid-glass-dark/settings-rail-long-label.png) |
| 字体选择 | ![下拉选项使用各自的字体](docs/comparisons/liquid-glass-dark/font-select.png) |
| 文本预览字号 | ![十八像素字号下的双行文本预览](docs/comparisons/liquid-glass-dark/preview-text.png) |
| 默认应用长名称 | ![默认应用菜单内的长名称换行](docs/comparisons/liquid-glass-dark/search-select-long-options.png) |
| 代码沿用文本程序 | ![代码选择器显示沿用的文本程序](docs/comparisons/liquid-glass-dark/search-select-app-inherited.png) |
| 同名程序 | ![默认应用菜单用路径区分同名程序](docs/comparisons/liquid-glass-dark/search-select-app-same-name.png) |
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
| 固定预览工具栏 | ![固定预览工具栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-toolbar.png) |
| 悬浮预览工具栏 | ![悬浮预览工具栏 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-toolbar-hover.png) |
| 图片预览 | ![图片预览 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-image.png) |
| 预览窗口缩放手柄 | ![缩放手柄 WebView 与 Native](docs/comparisons/liquid-glass-dark/preview-resize-handles.png) |
| 滚轮感应区域 | ![滚轮感应区域 WebView 与 Native](docs/comparisons/liquid-glass-dark/wheel-hot-zone.png) |
| 设置开关 | ![设置开关 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-toggle.png) |
| 带说明的设置开关 | ![粘贴后置顶与 Win+V 开关](docs/comparisons/liquid-glass-dark/settings-toggle-help.png) |
| 方向键预览从属设置 | ![方向键选择和预览开关](docs/comparisons/liquid-glass-dark/settings-keyboard-preview.png) |
| 设置滑块 | ![设置滑块 WebView 与 Native](docs/comparisons/liquid-glass-dark/settings-slider.png) |
| 快捷键绑定 | ![快捷键绑定 WebView 与 Native](docs/comparisons/liquid-glass-dark/hotkey-binding.png) |
| 在说明中清除可选快捷键 | ![单快捷键的清除操作](docs/comparisons/liquid-glass-dark/hotkey-optional.png) |
| 回到顶部 | ![回到顶部 WebView 与 Native](docs/comparisons/liquid-glass-dark/scroll-to-top.png) |
| 表情与符号搜索框 | ![表情与符号搜索框 WebView 与 Native](docs/comparisons/liquid-glass-dark/character-search.png) |
| 表情标签页 | ![表情标签页 WebView 与 Native](docs/comparisons/liquid-glass-dark/emoji-tabs.png) |
| 表情网格与连续滚动 | ![表情网格 WebView 与 Native](docs/comparisons/liquid-glass-dark/emoji-grid.png) |
| 滚动时自动显示后续符号 | ![滚动到底部后的符号面板](docs/comparisons/liquid-glass-dark/character-panel-scroll.png) |
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
