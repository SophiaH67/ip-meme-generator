package main

import (
	"log"
	"memegenerator/meme"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	log.Printf("%s requested %s", ip, r.URL.Path)

	m := meme.New(ip)
	m.Render()

	http.ServeFile(w, r, m.OutputPath)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}