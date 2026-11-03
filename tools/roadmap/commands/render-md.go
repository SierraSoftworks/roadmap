package commands

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/SierraSoftworks/roadmap/tools/internal/toolconfig"
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
		toolconfig.ProfileFlag("markdown"),
		&cli.BoolFlag{
			Name:  "simple",
			Usage: "Emits simplified Markdown for maximum compatibility.",
		},
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
		&cli.StringFlag{
			Name:      "template",
			Aliases:   []string{"temp"},
			Usage:     "The input template file.",
			Required:  false,
			TakesFile: true,
		},
		&cli.BoolFlag{
			Name:  "collapsed",
			Usage: "Collapse milestone deliverables into expandable sections.",
		},
	},
	Action: func(c *cli.Context) error {
		return toolconfig.Render(c, func(r *roadmap.Roadmap, opts *toolconfig.Options) (string, error) {
			tmpl := mdRoadmapAdvancedTemplate
			if opts.String("template") != "" {
				t, err := os.ReadFile(opts.Path("template"))
				if err != nil {
					return "", err
				}
				tmpl = string(t)
			} else if opts.Bool("simple") {
				tmpl = mdRoadmapBasicTemplate
			}

			params := map[string]interface{}{
				"collapsed": opts.Bool("collapsed"),
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
			})
			if err != nil {
				return "", err
			}

			return dot, nil
		})
	},
}
