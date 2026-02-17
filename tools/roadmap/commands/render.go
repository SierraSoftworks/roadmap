package commands

import (
	"bytes"
	_ "embed"
	"encoding/json"
	htmpl "html/template"
	"log"
	"strings"
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
		"markdown": func(in string) string {
			return strings.TrimSpace(string(markdown.ToHTML([]byte(in), nil, nil)))
		},
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"countByState": func(deliverables []*roadmap.Deliverable, state string) int {
			count := 0
			for _, d := range deliverables {
				if d.State == state {
					count++
				}
			}
			return count
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
		"sub": func(a, b int) int {
			return a - b
		},
		"countByState": func(deliverables []*roadmap.Deliverable, state string) int {
			count := 0
			for _, d := range deliverables {
				if d.State == state {
					count++
				}
			}
			return count
		},
	}
}

func renderHtmlTemplate(r *roadmap.Roadmap, t string, params map[string]interface{}, f htmpl.FuncMap) (string, error) {
	fm := getDefaultHTMLRenderFunctions()

	fm["param"] = func(key string) interface{} {
		return params[key]
	}

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

func renderTextTemplate(r *roadmap.Roadmap, t string, params map[string]interface{}, f template.FuncMap) (string, error) {
	fm := getDefaultTextRenderFunctions()

	fm["param"] = func(key string) interface{} {
		return params[key]
	}

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
