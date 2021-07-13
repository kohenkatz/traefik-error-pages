package main

import (
	"embed"
	"io/fs"
	"github.com/gorilla/handlers"
	"log"
	"os"
	"net/http"
)

//go:embed pages/*
var pages embed.FS

func main() {
	subFS, _ := fs.Sub(pages, "pages")
	pagesFS := http.FS(subFS)
	fs := http.FileServer(pagesFS)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, handlers.ProxyHeaders(fs)))

	log.Println("Listening on :3000...")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
