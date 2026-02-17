package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"
	"log"

	"github.com/SierraSoftworks/roadmap"
	"github.com/gomarkdown/markdown"
	"github.com/pkg/errors"
)

//go:embed page.css
var roadmapCss string

//go:embed roadmap.html
var roadmapTemplate string

func render(r *roadmap.Roadmap, params map[string]interface{}) (string, error) {
	tmpl := template.Must(template.New("roadmap").Funcs(template.FuncMap{
		"param": func(key string) interface{} {
			return params[key]
		},
		"json": func(in string) string {
			out, err := json.Marshal(in)
			if err != nil {
				log.Fatal(err)
			}

			return string(out)
		},
		"markdown": func(in string) template.HTML {
			return template.HTML(markdown.ToHTML([]byte(in), nil, nil))
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
		"stylesheet": func() template.CSS {
			return template.CSS(roadmapCss)
		},
	}).Parse(roadmapTemplate))

	buf := bytes.NewBufferString("")

	if err := tmpl.Execute(buf, r); err != nil {
		return "", errors.Wrap(err, "roadmap-html: could not generate roadmap from template")
	}

	return buf.String(), nil
}
