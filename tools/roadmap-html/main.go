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
		Name:  "roadmap-html",
		Usage: "Generate an HTML file which can be opened to visualize your roadmap specification.",
		Flags: []cli.Flag{
			toolconfig.ProfileFlag("html"),
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
				Usage:     "The output HTML file.",
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
				params := map[string]interface{}{
					"collapsed": opts.Bool("collapsed"),
				}
				dot, err := render(r, params)
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
