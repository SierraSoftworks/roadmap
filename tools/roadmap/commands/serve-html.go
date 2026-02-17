package commands

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

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
		&cli.BoolFlag{
			Name:  "collapsed",
			Usage: "Collapse milestone deliverables into expandable sections.",
		},
	},
	Action: func(c *cli.Context) error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			f, err := os.ReadFile(c.String("input"))
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

			params := map[string]interface{}{
				"collapsed": c.Bool("collapsed"),
			}
			dot, err := renderHtmlTemplate(rm, htmlRoadmapTemplate, params, template.FuncMap{
				"stylesheet": func() template.CSS {
					return template.CSS(htmlRoadmapCss)
				},
				"script": func() template.JS {
					return template.JS(htmlRoadmapWatchScript)
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

		http.HandleFunc("/watch", func(w http.ResponseWriter, r *http.Request) {
			stat, err := os.Stat(c.String("input"))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Failed to open roadmap file")
				return
			}

			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(": ready\n\n\n"))

			lastUpdated := stat.ModTime()
			for {
				time.Sleep(1 * time.Second)
				if _, err = w.Write([]byte(": ping\n\n\n")); err != nil {
					break
				}

				if stat, err := os.Stat(c.String("input")); err == nil {
					if stat.ModTime().After(lastUpdated) {
						w.Write([]byte("event: reload\ndata: reload\n\n\n"))
						lastUpdated = stat.ModTime()
					}
				} else {
					w.Write([]byte("event: error\ndata: error\n\n\n"))
					fmt.Printf("Failed to watch roadmap file for changes: %s", err)
					return
				}

				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}
		})

		fmt.Printf("Listening on http://localhost:%d/\n", c.Int("port"))
		return http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), nil)
	},
}
