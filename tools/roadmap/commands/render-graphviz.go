package commands

import (
	_ "embed"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/SierraSoftworks/roadmap/tools/internal/toolconfig"
	"github.com/urfave/cli/v2"
)

//go:embed templates/roadmap.dot
var graphvizRoadmapTemplate string

var graphvizRenderCommand = cli.Command{
	Name:    "graphviz",
	Aliases: []string{"dot"},
	Usage:   "Generate a GraphViz DOT file from a roadmap specification.",
	Flags: []cli.Flag{
		toolconfig.ProfileFlag("graphviz"),
		&cli.StringFlag{
			Name:      "input",
			Aliases:   []string{"in"},
			Usage:     "The input roadmap.yml `FILE`, or a glob pattern (e.g. **/roadmap.yml) matching several of them.",
			Value:     "roadmap.yml",
			TakesFile: true,
			Required:  true,
		},
		&cli.StringFlag{
			Name:      "output",
			Aliases:   []string{"out"},
			Usage:     "The output `FILE` (defaults to stdout).",
			TakesFile: true,
		},
	},
	Action: func(c *cli.Context) error {
		return toolconfig.Render(c, func(r *roadmap.Roadmap, opts *toolconfig.Options) (string, error) {
			dot, err := renderTextTemplate(r, graphvizRoadmapTemplate, map[string]interface{}{}, template.FuncMap{
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
			})
			if err != nil {
				return "", err
			}

			return dot, nil
		})
	},
}
