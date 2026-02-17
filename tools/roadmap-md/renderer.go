package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"log"
	"strings"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/gomarkdown/markdown"
	"github.com/pkg/errors"
)

//go:embed roadmap.advanced.md
var roadmapTemplateAdvanced string

//go:embed roadmap.basic.md
var roadmapTemplateBasic string

func render(r *roadmap.Roadmap, t string) (string, error) {
	tmpl := template.Must(template.New("roadmap").Funcs(template.FuncMap{
		"json": func(in string) string {
			out, err := json.Marshal(in)
			if err != nil {
				log.Fatal(err)
			}

			return string(out)
		},
		"markdown": func(in string) string {
			return strings.TrimSpace(string(markdown.ToHTML([]byte(in), nil, nil)))
		},
		"stateColor": func(state string) string {
			switch state {
			case "TODO":
				return "#aaa"
			case "DOING":
				return "#63B2EB"
			case "DONE":
				return "#3EAF7C"
			case "SKIP":
				return "#F65BD2"
			default:
				return "#aaa"
			}
		},
		"requirementColor": func(requirement string) string {
			switch requirement {
			case "MUST":
				return "#E06446"
			case "SHOULD":
				return "#E0AF2F"
			case "MAY":
				return "#3ABDE0"
			default:
				return "#888"
			}
		},
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(t))

	buf := bytes.NewBufferString("")

	if err := tmpl.Execute(buf, r); err != nil {
		return "", errors.Wrap(err, "roadmap-md: could not generate roadmap from template")
	}

	return buf.String(), nil
}
