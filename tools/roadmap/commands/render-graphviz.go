package commands

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
)

//go:embed templates/roadmap.dot
var graphvizRoadmapTemplate string

var graphvizRenderCommand = cli.Command{
	Name:    "graphviz",
	Aliases: []string{"dot"},
	Usage:   "Generate a GraphViz DOT file from a roadmap specification.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:      "input",
			Aliases:   []string{"in"},
			Usage:     "The input roadmap.yml `FILE`.",
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
		f, err := os.ReadFile(c.String("input"))
		if err != nil {
			return err
		}

		r, err := roadmap.Parse(f)
		if err != nil {
			return err
		}

		dot, err := renderTextTemplate(r, graphvizRoadmapTemplate, template.FuncMap{
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
			return err
		}

		if c.String("output") != "" {
			return os.WriteFile(c.String("output"), []byte(dot), os.ModePerm)
		} else {
			_, err := fmt.Println(dot)
			return err
		}
	},
}
