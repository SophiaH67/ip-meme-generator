package meme

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type Meme struct {
	opPath string
	edPath string
	ip string
	outputPath string
}

func New(ip string) Meme {
	m := Meme{
		opPath: "./assets/op.mp4",
		edPath: "./assets/ed.mp4",
		ip: ip,
		outputPath: fmt.Sprintf("./output/%s.mkv", ip),
	}
	return m
}

func (m *Meme) renderEd() *ffmpeg.Stream {
	ed := ffmpeg.Input(m.edPath).
	// Draw the IP on the middle of the image
		Filter("drawtext", ffmpeg.Args{
			fmt.Sprintf("text='%s'", m.ip),
			"fontsize=32",
			"fontcolor=black",
			"x=w/2",
			"y=h/2",
			"box=1:1",
		}).
		// Scale to 512x512
		Filter("scale", ffmpeg.Args{"512:512"}).
		// setdar to 1:1
		Filter("setdar", ffmpeg.Args{"1/1"})
	return ed
}

func (m *Meme) Render() {
	// Concatenate the op and ed
	ed := m.renderEd()
	base := ffmpeg.Concat([]*ffmpeg.Stream{
		ffmpeg.Input(m.opPath).
			Filter("scale", ffmpeg.Args{"512:512"}).
			Filter("setdar", ffmpeg.Args{"1/1"}),
		ed,
	})
	err := ffmpeg.Output([]*ffmpeg.Stream{
			base,
			ffmpeg.Input("./assets/sound.mp3"),
		}, m.outputPath, ffmpeg.KwArgs{"y": ""}).
		Run()

	if err != nil {
		panic(err)
	}
}