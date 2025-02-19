package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Template struct {
	tmpl *template.Template
}

func NewTemplate() (*Template, error) {
	templatesDir := "./src/templates"
	pattern := filepath.Join(templatesDir, "*.html")
	
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	return &Template{tmpl: tmpl}, nil
}

func (t *Template) Render(w http.ResponseWriter, name string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return t.tmpl.ExecuteTemplate(w, name, data)
}
