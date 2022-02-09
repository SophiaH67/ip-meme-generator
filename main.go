package main

import (
	"fmt"
	"memegenerator/meme"
	"os"
)

func main() {
	ip := "127.0.0.1"
	m := meme.New(ip)
	m.Render()
	// Move 127.0.0.1.mkv to 127.0.0.1.mp4
	err := os.Rename(fmt.Sprintf("./output/%s.mkv", ip), fmt.Sprintf("./output/%s.mp4", ip))
	if err != nil {
		panic(err)
	}
}