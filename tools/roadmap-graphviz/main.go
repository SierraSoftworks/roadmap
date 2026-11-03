package main

import (
	"log"
	"os"

	"github.com/SierraSoftworks/roadmap"
	"github.com/SierraSoftworks/roadmap/tools/internal/toolconfig"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "roadmap-graphviz",
		Usage: "Generate a GraphViz DOT file from a roadmap specification.",
		Flags: []cli.Flag{
			toolconfig.ProfileFlag("graphviz"),
			&cli.StringFlag{
				Name:      "input",
				Aliases:   []string{"in"},
				Usage:     "The input roadmap.yml file, or a glob pattern (e.g. **/roadmap.yml) matching several of them.",
				Required:  true,
				TakesFile: true,
			},
			&cli.StringFlag{
				Name:      "output",
				Aliases:   []string{"out"},
				Usage:     "The output DOT file.",
				Required:  false,
				TakesFile: true,
			},
		},
		Action: func(c *cli.Context) error {
			return toolconfig.Render(c, func(r *roadmap.Roadmap, opts *toolconfig.Options) (string, error) {
				dot, err := render(r)
				if err != nil {
					return "", err
				}

				return dot, nil
			})
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
