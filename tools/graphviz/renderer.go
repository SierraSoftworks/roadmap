package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"regexp"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/pkg/errors"
)

var idLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const roadmapTemplate = `
{{- $root := . -}}

digraph Roadmap {
  rankdir=LR;
  label={{ .Title | json }};
  fontname="Arial";
  compound=true;
  labelloc="t";

  node[style="filled",shape="rect",color="orange",fontname="Arial",fontsize=8];
  edge[weight=1,group="milestones"];

  {
	rank=same;
	
	start[label="Start",shape="house",color=""];

  {{- range $index, $milestone := .Milestones -}}
  {{- with $milestone -}}
  {{- $id := .Title | id }}
  
    {{ $id }}[label={{ .Title | json }},tooltip={{ .Description | json }}];
  {{- if gt $index 0 }}
    {{ (index $root.Milestones (add $index -1)).Title | id }} -> {{ $id }};
  {{- else }}
    start -> {{ $id }};
  {{- end }}
  {{- end }}
  {{- end }}
  }

  node[color="grey"];
  edge[weight=5,color="grey",penwidth=0.4,arrowsize=0.4,group="dependencies"];

  {{ range .Milestones }}
  {{- $id := .Title | id }}

  subgraph cluster_{{ $id }} {
	label={{ .Title | json }};
	fontsize=8;
	penwidth=0.6;

	# Deliverables for {{ $id }}
	{{ range .Deliverables -}}
	{{ .Title | id }}[label={{ .Title | json }},tooltip={{ .Description | json }},color="{{ .State | stateColor }}",labelhref={{ .Reference | json }}];
	{{ .Title | id }} -> {{ $id }};
	{{- end }}
  }
  {{ end }}
}
`

func render(r *roadmap.Roadmap) (string, error) {
	tmpl := template.Must(template.New("roadmap").Funcs(template.FuncMap{
		"newid": func() string {
			b := make([]rune, 10)
			for i := range b {
				b[i] = idLetters[rand.Intn(len(idLetters))]
			}
			return string(b)
		},
		"id": func(title string) string {
			return string(regexp.MustCompile("[^a-zA-Z0-9_]").ReplaceAll([]byte(title), []byte("_")))
		},
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
				return "grey"
			case "DOING":
				return "lightblue"
			case "DONE":
				return "palegreen"
			case "SKIP":
				return "pink"
			default:
				return "grey"
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
