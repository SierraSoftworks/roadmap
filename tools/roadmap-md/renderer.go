package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"log"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/pkg/errors"
)

//go:embed roadmap.md
var roadmapTemplate string

func render(r *roadmap.Roadmap) (string, error) {
	tmpl := template.Must(template.New("roadmap").Funcs(template.FuncMap{
		"json": func(in string) string {
			out, err := json.Marshal(in)
			if err != nil {
				log.Fatal(err)
			}

			return string(out)
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
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(roadmapTemplate))

	buf := bytes.NewBufferString("")

	if err := tmpl.Execute(buf, r); err != nil {
		return "", errors.Wrap(err, "graphviz: could not generate roadmap from template")
	}

	return buf.String(), nil
}
