package main

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
)

//go:embed all:web/build
var webBuildFS embed.FS

type webBuildServer struct {
	fs http.FileSystem
}

func NewBuildServer() http.FileSystem {
	build, _ := fs.Sub(webBuildFS, "web/build")
	return &webBuildServer{
		fs: http.FS(build),
	}
}

func (f *webBuildServer) Open(name string) (http.File, error) {
	file, err := f.fs.Open(name)
	if errors.Is(err, fs.ErrNotExist) {
		return f.fs.Open("index.html")
	}
	return file, err
}
