package handler

import (
	"html/template"
	"io"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/guregu/null/zero"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type TemplateRenderer struct {
	templates *template.Template
}

func toString(data interface{}) string {
	if reflect.TypeOf(data).String() == reflect.TypeOf(zero.String{}).String() {
		return data.(zero.String).ValueOrZero()
	}
	return cast.ToString(data)
}

func strTitle(data interface{}) string {
	return strings.Title(strings.ToLower(data.(string)))
}

func templateHelperFunc() template.FuncMap {
	return map[string]any{
		"helper_str_title": strTitle,
		"helper_to_string": toString,
	}
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

	templates := template.Must(template.New("").Funcs(templateHelperFunc()).ParseFiles(allFiles...))
	ptr := &TemplateRenderer{templates: templates}
	return ptr
}
