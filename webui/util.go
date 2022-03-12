package webui

import (
	"embed"
	"html/template"
)

//go:embed static/*
var staticFS embed.FS

//go:embed templates/*
var templateFS embed.FS

func mkTmpl(t []string) *template.Template {
	return template.Must(template.ParseFS(templateFS, t...))
}
