import { cp, copyFile } from "node:fs/promises";

await cp(new URL("../src/styles", import.meta.url), new URL("../dist/styles", import.meta.url), { recursive: true });
await copyFile(new URL("../src/base.css", import.meta.url), new URL("../dist/base.css", import.meta.url));
