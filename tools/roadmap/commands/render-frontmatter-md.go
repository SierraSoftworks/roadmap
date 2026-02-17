package commands

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v3"
)

//go:embed templates/roadmap.basic.frontmatter.md
var frontmatterMdRoadmapBasicTemplate string

//go:embed templates/roadmap.advanced.frontmatter.md
var frontmatterMdRoadmapAdvancedTemplate string

var frontmatterMdRenderCommand = cli.Command{
	Name:    "frontmatter-markdown",
	Aliases: []string{"fmd"},
	Usage:   "Generate a Markdown file which can be rendered to visualize your roadmap specification, including a frontmatter preamble.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "simple",
			Usage: "Emits simplified Markdown for maximum compatibility.",
		},
		&cli.StringFlag{
			Name:    "frontmatter",
			Aliases: []string{"m"},
			Usage:   "Additional frontmatter fields to include in the frontmatter, should be provided in YAML format.",
			Value:   "{}",
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
		&cli.BoolFlag{
			Name:  "collapsed",
			Usage: "Collapse milestone deliverables into expandable sections.",
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

		tmpl := frontmatterMdRoadmapAdvancedTemplate
		if c.Bool("simple") {
			tmpl = frontmatterMdRoadmapBasicTemplate
		}

		params := map[string]interface{}{
			"collapsed": c.Bool("collapsed"),
		}
		dot, err := renderTextTemplate(r, tmpl, params, template.FuncMap{
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
			"metadata": func() map[string]interface{} {
				m := map[string]interface{}{}
				if err := yaml.Unmarshal([]byte(c.String("frontmatter")), &m); err != nil {
					return map[string]interface{}{
						"__parse_error": err.Error(),
					}
				}

				if m["title"] == nil {
					m["title"] = r.Title
				}

				if m["description"] == nil {
					m["description"] = r.Description
				}

				return m
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
