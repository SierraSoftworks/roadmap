package commands

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
)

//go:embed templates/roadmap.basic.md
var mdRoadmapBasicTemplate string

//go:embed templates/roadmap.advanced.md
var mdRoadmapAdvancedTemplate string

var mdRenderCommand = cli.Command{
	Name:    "markdown",
	Aliases: []string{"md"},
	Usage:   "Generate a Markdown file which can be rendered to visualize your roadmap specification.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "simple",
			Usage: "Emits simplified Markdown for maximum compatibility.",
		},
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
		f, err := ioutil.ReadFile(c.String("input"))
		if err != nil {
			return err
		}

		r, err := roadmap.Parse(f)
		if err != nil {
			return err
		}

		tmpl := mdRoadmapAdvancedTemplate
		if c.Bool("simple") {
			tmpl = mdRoadmapBasicTemplate
		}

		dot, err := renderTextTemplate(r, tmpl, template.FuncMap{
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
		})
		if err != nil {
			return err
		}

		if c.String("output") != "" {
			return ioutil.WriteFile(c.String("output"), []byte(dot), os.ModePerm)
		} else {
			_, err := fmt.Println(dot)
			return err
		}
	},
}
