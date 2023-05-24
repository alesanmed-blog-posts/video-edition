package main

import (
	"fmt"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	overlayFile := ffmpeg_go.Input("/path/to/static/overlays/black.png")

	textStartTime := 5
	textEndTime := 20
	fadeInDuration := 1.5
	fadeOutDuration := 1.5

	ffmpeg_go.Input("/path/to/static/videos/base_video.mp4").
		Overlay(overlayFile, "repeat", ffmpeg_go.KwArgs{
			"x": 0,
			"y": 0,
		}).
		Drawtext("Aleware,\nTremenda Sabrosura", 0, 0, false, ffmpeg_go.KwArgs{
			"x":              "(w - text_w) / 2",
			"y":              "(h - text_h) / 2",
			"fontsize":       "h/10",
			"line_spacing":   30,
			"fontfile":       "/path/to/static/fonts/Roboto-Regular.ttf",
			"fontcolor_expr": fmt.Sprint("ffffff%{ eif\\: clip(255*(between(t, ", float64(textStartTime)+fadeInDuration, ", ", float64(textEndTime)-fadeOutDuration, ") + ((t - ", textStartTime, ") / ", fadeInDuration, ") * between(t, ", textStartTime, ", ", float64(textStartTime)+fadeInDuration, ") + (-(t - ", textEndTime, ")/", fadeOutDuration, ") * between(t, ", float64(textEndTime)-fadeOutDuration, ", ", textEndTime, ")), 0, 255) \\: x\\: 2}"),
		}).
		Output("/path/to/static/videos/aleware_video.mp4", ffmpeg_go.KwArgs{
			"c:a": "copy",
		}).
		OverWriteOutput().ErrorToStdOut().Run()
}
