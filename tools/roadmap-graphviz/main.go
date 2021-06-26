package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "roadmap-graphviz",
		Usage: "Generate a GraphViz DOT file from a roadmap specification.",
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
				Usage:     "The output DOT file.",
				Required:  false,
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

			dot, err := render(r)
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
