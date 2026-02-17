package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "roadmap-md",
		Usage: "Generate a Markdown file which can be rendered to visualize your roadmap specification.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "input",
				Aliases:   []string{"in"},
				Usage:     "The input roadmap.yml file.",
				Required:  true,
				TakesFile: true,
			},
			&cli.StringFlag{
				Name:      "output",
				Aliases:   []string{"out"},
				Usage:     "The output Markdown file.",
				Required:  false,
				TakesFile: true,
			},
			&cli.BoolFlag{
				Name:  "simple",
				Usage: "Emits simplified Markdown for maximum compatibility.",
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
			f, err := os.ReadFile(c.String("input"))
			if err != nil {
				return err
			}

			r, err := roadmap.Parse(f)
			if err != nil {
				return err
			}

			template := roadmapTemplateAdvanced

			if c.String("template") != "" {
				t, err := os.ReadFile(c.String("template"))
				if err != nil {
					return err
				}
				template = string(t)
			} else if c.Bool("simple") {
				template = roadmapTemplateBasic
			}

			dot, err := render(r, template, c.Bool("collapsed"))
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
