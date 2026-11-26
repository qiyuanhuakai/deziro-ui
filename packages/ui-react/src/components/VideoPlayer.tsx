import { LoaderCircle, Pause, Play, Repeat, Volume2, VolumeX } from "lucide-react";
import { useEffect, useRef, useState, type ReactNode } from "react";
import "../styles/components/video-player.css";

export interface VideoPlayerLabels {
    play: string;
    pause: string;
    progress: string;
    mute: string;
    unmute: string;
    volume: string;
    loop: string;
}

interface VideoPlayerProps {
    source?: string;
    poster?: string;
    loading?: boolean;
    status?: string;
    labels: VideoPlayerLabels;
    children?: ReactNode;
    onOpen: () => void;
    onError: () => void;
}

function timestamp(value: number) {
    const seconds = Math.floor(value);
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor(seconds / 60) % 60;
    const clock = `${minutes.toString().padStart(2, "0")}:${(seconds % 60).toString().padStart(2, "0")}`;
    return hours ? `${hours}:${clock}` : clock;
}

export function VideoPlayer({ source, poster, loading, status, labels, children, onOpen, onError }: VideoPlayerProps) {
    const video = useRef<HTMLVideoElement>(null);
    const [position, setPosition] = useState(0);
    const [duration, setDuration] = useState(0);
    const [playing, setPlaying] = useState(false);
    const [muted, setMuted] = useState(false);
    const [volume, setVolume] = useState(1);
    const [loop, setLoop] = useState(false);

    useEffect(() => {
        const element = video.current;
        return () => {
            if (!element) return;
            element.pause();
            element.removeAttribute("src");
            element.load();
        };
    }, []);

    const togglePlayback = () => {
        const element = video.current;
        if (!element) return;
        if (element.paused) void element.play().catch(() => undefined);
        else element.pause();
    };
    const playLabel = playing ? labels.pause : labels.play;
    const muteLabel = muted ? labels.unmute : labels.mute;
    return <div className="dz-video-player">
        <div className="dz-video-stage">
            {source ? <video ref={video} src={source} poster={poster}
                disablePictureInPicture disableRemotePlayback autoPlay playsInline preload="metadata"
                aria-label={labels.play} onError={onError}
                onPlay={() => setPlaying(true)} onPause={() => setPlaying(false)} onEnded={() => setPlaying(false)}
                onTimeUpdate={event => setPosition(event.currentTarget.currentTime)}
                onDurationChange={event => {
                    const length = event.currentTarget.duration;
                    setDuration(Number.isFinite(length) ? length : 0);
                }} /> : <>
                {poster && <img src={poster} alt="" />}
                <button type="button" className="dz-video-start" aria-label={labels.play}
                    title={labels.play} disabled={loading} onClick={onOpen}>
                    {loading ? <LoaderCircle aria-hidden="true" /> : <Play aria-hidden="true" />}
                </button>
            </>}
        </div>
        {source ? <div className="dz-video-controls">
            <input className="dz-video-progress" type="range" aria-label={labels.progress}
                aria-valuetext={timestamp(position)} min={0} max={duration || 1}
                step={0.1} value={Math.min(position, duration)} disabled={duration <= 0}
                onChange={event => {
                    if (video.current) video.current.currentTime = event.currentTarget.valueAsNumber;
                    setPosition(event.currentTarget.valueAsNumber);
                }} />
            <div className="dz-video-actions">
                <button type="button" aria-label={playLabel} title={playLabel} onClick={togglePlayback}>
                    {playing ? <Pause aria-hidden="true" /> : <Play aria-hidden="true" />}
                </button>
                <span className="dz-video-clock">{timestamp(position)} / {timestamp(duration)}</span>
                <button type="button" aria-label={muteLabel} title={muteLabel} aria-pressed={muted}
                    onClick={() => {
                        if (video.current) video.current.muted = !muted;
                        setMuted(!muted);
                    }}>
                    {muted || volume === 0 ? <VolumeX aria-hidden="true" /> : <Volume2 aria-hidden="true" />}
                </button>
                <input className="dz-video-volume" type="range" aria-label={labels.volume}
                    min={0} max={1} step={0.05} value={volume}
                    onChange={event => {
                        const next = event.currentTarget.valueAsNumber;
                        if (video.current) {
                            video.current.volume = next;
                            video.current.muted = false;
                        }
                        setVolume(next);
                        setMuted(false);
                    }} />
                <button type="button" aria-label={labels.loop} title={labels.loop} aria-pressed={loop}
                    onClick={() => {
                        if (video.current) video.current.loop = !loop;
                        setLoop(!loop);
                    }}>
                    <Repeat aria-hidden="true" />
                </button>
            </div>
        </div> : <>
            {status && <div role="status" className="dz-video-status">{status}</div>}
            {children}
        </>}
    </div>;
}
