package handler

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplateRenderer creates a new TemplateRenderer instance
func NewTemplateRenderer(glob string) *TemplateRenderer {
	allFiles, err := filepath.Glob(glob)
	if err != nil {
		panic(err)
	}

	templates := template.Must(template.ParseFiles(allFiles...))
	return &TemplateRenderer{templates: templates}
}
