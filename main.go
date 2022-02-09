package main

import (
	"memegenerator/meme"
)

func main() {
	ip := "127.0.0.1"
	m := meme.New(ip)
	m.Render()
}