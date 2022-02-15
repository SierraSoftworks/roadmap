package commands

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
)

var serveHtmlCommand = cli.Command{
	Name:  "html",
	Usage: "Serves an HTML version of your roadmap.",
	Flags: []cli.Flag{

		&cli.StringFlag{
			Name:      "input",
			Aliases:   []string{"in"},
			Usage:     "The input roadmap.yml `FILE`.",
			Value:     "roadmap.yml",
			TakesFile: true,
			Required:  true,
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "The `PORT` to listen on.",
			Value:   8080,
		},
	},
	Action: func(c *cli.Context) error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			f, err := ioutil.ReadFile(c.String("input"))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Failed to open roadmap file: %s", err)
				return
			}

			rm, err := roadmap.Parse(f)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error parsing roadmap: %s", err)
				return
			}

			dot, err := renderHtmlTemplate(rm, htmlRoadmapTemplate, template.FuncMap{
				"stylesheet": func() template.CSS {
					return template.CSS(htmlRoadmapCss)
				},
			})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error rendering roadmap: %s", err)
				return
			}

			_, err = bytes.NewBufferString(dot).WriteTo(w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error writing roadmap to output: %s", err)
			}
		})

		fmt.Printf("Listening on http://localhost:%d/...\n", c.Int("port"))
		return http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), nil)
	},
}
