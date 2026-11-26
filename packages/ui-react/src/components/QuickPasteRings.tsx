import { useLayoutEffect, useRef, useState } from "react";
import "../styles/components/quick-paste-rings.css";

export type QuickPasteDock = "left" | "right" | "";

interface RingItem {
    id: string | number;
    label: string;
    compactLabel?: string;
    detail?: string;
    active?: boolean;
    onSelect: () => void;
    onHighlight?: () => void;
}

interface QuickPasteRingsProps {
    dock: QuickPasteDock;
    categories: RingItem[];
    entries: RingItem[];
    categoryLabel: string;
    entryLabel: string;
    emptyLabel?: string;
}

function ringSector(index: number, count: number, outer: boolean, dock: QuickPasteDock, width: number, height: number) {
    const radius = Math.max(0, Math.min(height / 2 - 4, dock ? width - 8 : width / 2 - 4));
    const cx = dock === "left" ? 0 : dock === "right" ? width : width / 2;
    const cy = height / 2;
    const direction = dock === "right" ? -1 : 1;
    const inner = radius * (outer ? .69 : .28);
    const edge = radius * (outer ? 1 : .66);
    const step = (dock ? Math.PI : 2 * Math.PI) / count;
    const middle = -Math.PI / 2 + step * (index + (dock ? .5 : 0));
    const halfAngle = step / 2 - .018;
    const point = (r: number, angle: number) => ({
        x: cx + direction * r * Math.cos(angle),
        y: cy + r * Math.sin(angle),
    });
    const label = point((inner + edge) / 2, middle);
    const samples = 18;
    const arc = Array.from({ length: samples + 1 }, (_, i) => middle - halfAngle + 2 * halfAngle * i / samples);
    const points = [...arc.map(a => point(edge, a)), ...arc.reverse().map(a => point(inner, a))];
    const left = Math.min(...points.map(p => p.x));
    const top = Math.min(...points.map(p => p.y));
    const sectorWidth = Math.max(...points.map(p => p.x)) - left;
    const sectorHeight = Math.max(...points.map(p => p.y)) - top;
    let halfWidth = outer ? 54 : 36;
    const halfHeight = outer ? 12 : 7;
    const contains = (x: number, y: number) => {
        const dx = direction * (x - cx);
        const dy = y - cy;
        const distance = Math.hypot(dx, dy);
        const angle = Math.atan2(dy, dx) - middle;
        return distance >= inner + 2 && distance <= edge - 2 &&
            Math.abs(Math.atan2(Math.sin(angle), Math.cos(angle))) <= halfAngle;
    };
    while (halfWidth > 2 && ![-halfWidth, halfWidth].every(dx =>
        [-halfHeight, halfHeight].every(dy => contains(label.x + dx, label.y + dy)))) {
        halfWidth--;
    }
    return {
        bounds: {
            left, top, width: sectorWidth, height: sectorHeight,
            clipPath: "polygon(" + points.map(p => (p.x - left) + "px " + (p.y - top) + "px").join(",") + ")",
        },
        label: { left: label.x - left, top: label.y - top, width: halfWidth * 2 },
    };
}

export function QuickPasteRings({ dock, categories, entries, categoryLabel, entryLabel, emptyLabel }: QuickPasteRingsProps) {
    const view = useRef<HTMLDivElement>(null);
    const [size, setSize] = useState({ width: 0, height: 0 });
    useLayoutEffect(() => {
        const node = view.current;
        if (!node) return;
        const resize = () => setSize({ width: node.clientWidth, height: node.clientHeight });
        resize();
        const observer = new ResizeObserver(resize);
        observer.observe(node);
        return () => observer.disconnect();
    }, []);

    const renderRing = (items: RingItem[], count: number, outer: boolean) =>
        Array.from({ length: count }, (_, index) => {
            const item = items[index];
            const geometry = ringSector(index, count, outer, dock, size.width, size.height);
            if (!item) return <span key={index} aria-hidden="true" className="quick-paste-ring__empty-sector" style={geometry.bounds} />;
            const label = dock ? item.compactLabel ?? item.label : item.label;
            return <button type="button" key={item.id} style={geometry.bounds}
                className={outer ? "quick-paste-ring__entry" : "quick-paste-ring__category"}
                aria-label={item.label} aria-pressed={item.active ?? false}
                onClick={item.onSelect} onPointerEnter={item.onHighlight}
                title={item.detail ? item.label + " · " + item.detail : item.label}>
                <span className="quick-paste-ring__label" style={geometry.label}>
                    <span>{label}</span>
                    {item.detail && <small>{item.detail}</small>}
                </span>
            </button>;
        });
    return <div ref={view} className={"quick-paste-ring" + (dock ? " quick-paste-ring--" + dock : "")}>
        {size.width > 0 && size.height > 0 && <>
            <div role="group" aria-label={categoryLabel}>{renderRing(categories, categories.length, false)}</div>
            <div role="group" aria-label={entryLabel}>{renderRing(entries, 6, true)}</div>
        </>}
        {entries.length === 0 && emptyLabel && <span className="quick-paste-ring__empty">{emptyLabel}</span>}
    </div>;
}
