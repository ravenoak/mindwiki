package echo

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type HTMLTemplate struct {
	templates *template.Template
}

func (t *HTMLTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
