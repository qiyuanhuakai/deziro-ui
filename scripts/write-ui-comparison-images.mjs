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
  ["history-card", "compact-hover", 90, 0, "history-card-compact-actions"],
  ["context-menu", "text", 215],
  ["context-menu", "privacy", 235, 0, "context-menu-privacy"],
  ["text-fields", "default", 275],
  ["app-header", "pinned", 185],
  ["history-list", "selected", 245],
  ["history-card", "selected", 110],
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
  ["preview-toolbar", "pinned", 100],
  ["preview-image", "zoomed", 275],
  ["preview-resize-handles", "southeast", 185],
  ["wheel-hot-zone", "scrolled", 185],
];
const additional = [
  ["settings-section-card", "expanded", 125],
  ["settings-section-card", "wide", 195, 0, "settings-section-grid"],
  ["settings-choices", "default", 185],
  ["hotkey-list", "multiple", 120],
  ["emoji-favorites", "default", 125],
  ["settings-data", "default", 185],
  ["settings-remove", "default", 115],
  ["tag-toolbar", "default", 85],
  ["clipboard-search", "default", 140],
  ["settings-toggle", "enabled", 80],
  ["settings-slider", "default", 95],
  ["hotkey-binding", "default", 85],
  ["scroll-to-top", "default", 165],
  ["character-search", "default", 82],
  ["emoji-tabs", "symbols", 82],
  ["emoji-grid", "default", 105],
];
const components = [...figures, ...additional];
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
