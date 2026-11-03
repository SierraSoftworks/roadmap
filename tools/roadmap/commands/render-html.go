package commands

import (
	_ "embed"
	"html/template"

	"github.com/SierraSoftworks/roadmap"
	"github.com/SierraSoftworks/roadmap/tools/internal/toolconfig"
	"github.com/urfave/cli/v2"
)

//go:embed templates/roadmap.html.css
var htmlRoadmapCss string

//go:embed templates/roadmap.html.watch.js
var htmlRoadmapWatchScript string

//go:embed templates/roadmap.html
var htmlRoadmapTemplate string

var htmlRenderCommand = cli.Command{
	Name:  "html",
	Usage: "Generate an HTML file which can be opened to visualize your roadmap specification.",
	Flags: []cli.Flag{
		toolconfig.ProfileFlag("html"),
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
		&cli.BoolFlag{
			Name:  "collapsed",
			Usage: "Collapse milestone deliverables into expandable sections.",
		},
	},
	Action: func(c *cli.Context) error {
		return toolconfig.Render(c, func(r *roadmap.Roadmap, opts *toolconfig.Options) (string, error) {
			params := map[string]interface{}{
				"collapsed": opts.Bool("collapsed"),
			}
			dot, err := renderHtmlTemplate(r, htmlRoadmapTemplate, params, template.FuncMap{
				"stylesheet": func() template.CSS {
					return template.CSS(htmlRoadmapCss)
				},
				"script": func() template.JS {
					return template.JS("")
				},
			})
			if err != nil {
				return "", err
			}

			return dot, nil
		})
	},
}
