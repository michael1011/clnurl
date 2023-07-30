package router

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
)

const outDir = "out"

//go:embed out/* out/_next/static/**/* out/_next/static/chunks/pages/*
var Assets embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func assetHandler() http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(outDir, name)
		return Assets.Open(assetPath)
	})

	return http.FileServer(http.FS(handler))
}
