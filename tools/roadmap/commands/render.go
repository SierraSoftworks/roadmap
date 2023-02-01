package commands

import (
	"bytes"
	_ "embed"
	"encoding/json"
	htmpl "html/template"
	"log"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/gomarkdown/markdown"
	"github.com/pkg/errors"
)

func getDefaultTextRenderFunctions() template.FuncMap {
	return template.FuncMap{
		"json": func(in interface{}) string {
			out, err := json.Marshal(in)
			if err != nil {
				log.Fatal(err)
			}

			return string(out)
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
}

func getDefaultHTMLRenderFunctions() htmpl.FuncMap {
	return htmpl.FuncMap{
		"json": func(in string) string {
			out, err := json.Marshal(in)
			if err != nil {
				log.Fatal(err)
			}

			return string(out)
		},
		"markdown": func(in string) htmpl.HTML {
			return htmpl.HTML(markdown.ToHTML([]byte(in), nil, nil))
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
}

func renderHtmlTemplate(r *roadmap.Roadmap, t string, f htmpl.FuncMap) (string, error) {
	fm := getDefaultHTMLRenderFunctions()

	for k, v := range f {
		fm[k] = v
	}

	tmpl := htmpl.Must(htmpl.New("roadmap").Funcs(fm).Parse(t))

	buf := bytes.NewBufferString("")

	if err := tmpl.Execute(buf, r); err != nil {
		return "", errors.Wrap(err, "roadmap: could not generate roadmap from template")
	}

	return buf.String(), nil
}

func renderTextTemplate(r *roadmap.Roadmap, t string, f template.FuncMap) (string, error) {
	fm := getDefaultTextRenderFunctions()

	for k, v := range f {
		fm[k] = v
	}

	tmpl := template.Must(template.New("roadmap").Funcs(fm).Parse(t))

	buf := bytes.NewBufferString("")

	if err := tmpl.Execute(buf, r); err != nil {
		return "", errors.Wrap(err, "roadmap: could not generate roadmap from template")
	}

	return buf.String(), nil
}
