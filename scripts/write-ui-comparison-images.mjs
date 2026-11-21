import { mkdir, readFile } from "node:fs/promises";
import { join, resolve } from "node:path";
import { chromium } from "playwright";

const flags = new URLSearchParams(process.argv.slice(2).join("&"));
const artifacts = flags.get("results");
const destination = resolve(flags.get("output") ?? "docs/comparisons");
const theme = flags.get("theme") ?? "macos";
const mode = flags.get("mode") ?? "light";
if (!artifacts) throw Error("Pass results=<directory written by compare-ui-components.mjs>");

// Each published figure shows one component using the same tested theme and
// state, trimmed to keep its visible controls legible in the README.
const figures = [
  ["settings-toggle-help", "expanded", 165],
  ["settings-keyboard-preview", "default", 155],
  ["character-panel", "after-scroll", 310, 0, "character-panel-scroll"],
  ["settings-field", "expanded", 200],
  ["history-card", "compact-hover", 90, 0, "history-card-compact-actions"],
  ["context-menu", "text", 215],
  ["context-menu", "privacy", 235, 0, "context-menu-privacy"],
  ["text-fields", "default", 275],
  ["app-header", "pinned", 185],
  ["history-list", "selected", 245],
  ["history-card", "selected", 110],
  ["history-typography", "small", 276, 0, "history-typography"],
  ["history-typography", "compact-large", 205, 0, "history-typography-compact"],
  ["history-typography", "custom", 292, 0, "history-typography-custom"],
  ["history-rich-text", "default", 150, 0, "history-rich-text"],
  ["history-rich-text", "compact", 112, 0, "history-rich-text-compact"],
  ["history-card-meta", "default", 90],
  ["search-bar", "filled", 110],
  ["search-history", "default", 135],
  ["filters", "default", 95],
  ["highlight", "default", 85],
  ["select", "popup", 250],
  ["settings-rail", "selected", 180],
  ["settings-footer", "default", 165],
  ["disclosure", "expanded", 110],
  ["help", "expanded", 105],
  ["confirm", "default", 285, 65],
  ["tag-dialog", "default", 285, 65],
  ["toast", "default", 130],
  ["progress", "default", 140],
  ["attachment-preview", "media", 180],
  ["attachment-preview", "compact-media", 132, 0, "attachment-preview-compact"],
  ["attachment-preview", "narrow-media", 200, 0, "attachment-preview-narrow"],
  ["history-image", "default", 405],
  ["history-image", "compact", 196, 0, "history-image-compact"],
  ["history-image", "unavailable", 145, 0, "history-image-unavailable"],
  ["history-ocr", "expanded", 235],
  ["history-ocr", "compact-expanded", 155, 0, "history-ocr-compact"],
  ["preview-toolbar", "pinned", 100],
  ["preview-toolbar", "default", 100, 0, "preview-toolbar-hover"],
  ["preview-image", "zoomed", 275],
  ["preview-resize-handles", "southeast", 185],
  ["wheel-hot-zone", "scrolled", 185],
];
const additional = [
  ["settings-section-card", "expanded", 125],
  ["settings-section-card", "wide", 195, 0, "settings-section-grid"],
  ["settings-choices", "default", 185],
  ["hotkey-list", "multiple", 120],
  ["hotkey-list", "win-v", 120, 0, "hotkey-list-win-v"],
  ["hotkey-list", "extended", 120, 0, "hotkey-list-extended"],
  ["emoji-favorites", "default", 125],
  ["settings-data", "default", 185],
  ["settings-data", "hover-path", 185, 0, "settings-data-hover"],
  ["backup-dialog", "default", 310, 55, "backup-dialog"],
  ["backup-dialog", "encrypted", 396, 12, "backup-dialog-encrypted"],
  ["backup-dialog", "import", 376, 22, "backup-dialog-import"],
  ["backup-dialog", "replace-confirmed", 407, 6, "backup-dialog-replace"],
  ["settings-update", "default", 155, 0, "settings-update"],
  ["settings-update", "installing", 175, 0, "settings-update-installing"],
  ["settings-remove", "default", 115],
  ["privacy-rules", "default", 390],
  ["tag-toolbar", "default", 85],
  ["tag-search", "default", 180],
  ["search-select", "apps-filtered", 218],
  ["search-select", "empty", 218, 0, "search-select-empty"],
  ["search-select", "typo", 218, 0, "search-select-typo"],
  ["clipboard-search", "default", 140],
  ["settings-toggle", "enabled", 80],
  ["settings-slider", "default", 95],
  ["hotkey-binding", "default", 85],
  ["scroll-to-top", "default", 165],
  ["character-search", "default", 82],
  ["emoji-tabs", "symbols", 82],
  ["emoji-grid", "default", 105],
];
const opacityFigures = [
  ["surface-opacity", "20", 312, 0, "surface-opacity-20"],
  ["surface-opacity", "85", 312, 0, "surface-opacity-85"],
];
const components = theme === "liquid-glass" ? [...figures, ...additional, ...opacityFigures] : [...figures, ...additional];
const fixture = flags.get("fixture");
const requested = fixture?.split(",");
const gallery = requested ? components.filter(([name, , , , basename]) => requested.includes(basename ?? name)) : components;
if (gallery.length === 0) throw Error("Unknown comparison component: " + fixture);
const additionalNames = new Set(additional.map(([name]) => name));
const results = JSON.parse(await readFile(join(artifacts, "results.json"), "utf8"));
const reports = new Map(results.map(row => [row.id, row]));
const capture = async (id, side) => "data:image/png;base64," + (await readFile(join(artifacts, id + "_" + side + ".png"))).toString("base64");
for (const [component, state] of gallery) {
  const id = [theme, mode, component, state].join("_");
  const report = reports.get(id);
  if (report?.passed) continue;
  // The supplemental figures also show any measured differences between the two renderers.
  if (additionalNames.has(component) && report?.sameText && report.expectedState && report.errors.length === 0 &&
      report.inkWeb > 0 && report.inkNative > 0) continue;
  throw Error("Missing a verified render for " + id);
}

await mkdir(destination, { recursive: true });
const browser = await chromium.launch({ headless: true });
try {
  const page = await browser.newPage({ viewport: { width: 1156, height: 485 }, deviceScaleFactor: 1 });
  for (const [component, state, height, offset = 0, basename = component] of gallery) {
    const id = [theme, mode, component, state].join("_");
    const [web, native] = await Promise.all([capture(id, "web"), capture(id, "native")]);
    await page.setContent(`<!doctype html><html lang="en"><meta charset="utf-8"><style>
      * { box-sizing: border-box }
      body { margin: 0; font: 600 15px system-ui, sans-serif; color: #162238; background: #e7ebf2 }
      main { display: grid; grid-template-columns: repeat(2, 560px); gap: 12px; width: 1156px; padding: 12px }
      figure { margin: 0; overflow: hidden; border-radius: 8px; background: #fff }
      figcaption { padding: 10px 14px; border-bottom: 1px solid #d8dfea }
      .view { height: ${height}px; overflow: hidden }
      img { display: block; width: 560px; height: 420px; margin-top: -${offset}px }
    </style><main>
      <figure><figcaption>React WebView</figcaption><div class="view"><img src="${web}" alt=""></div></figure>
      <figure><figcaption>MyGo Native</figcaption><div class="view"><img src="${native}" alt=""></div></figure>
    </main></html>`);
    await page.locator("main").screenshot({ path: join(destination, basename + ".png") });
  }
} finally {
  await browser.close();
}
console.log(`Wrote ${gallery.length} comparisons to ${destination}`);
