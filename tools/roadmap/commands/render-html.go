package commands

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"

	"github.com/SierraSoftworks/roadmap"
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

		dot, err := renderHtmlTemplate(r, htmlRoadmapTemplate, template.FuncMap{
			"stylesheet": func() template.CSS {
				return template.CSS(htmlRoadmapCss)
			},
			"script": func() template.JS {
				return template.JS("")
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
