package commands

import "github.com/urfave/cli/v2"

var app = &cli.App{
	Name:      "roadmap",
	Usage:     "The command line tool used to manage your roadmap.yml files.",
	Copyright: "Copyright © Sierra Softworks 2021",
	Authors: []*cli.Author{
		{Name: "Benjamin Pannell", Email: "contact@sierrasoftworks.com"},
	},
	Commands: cli.Commands{
		&cli.Command{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "Serve your roadmap.yml file over HTTP.",
			Subcommands: cli.Commands{
				&serveHtmlCommand,
			},
		},
		&cli.Command{
			Name:    "render",
			Aliases: []string{"to"},
			Usage:   "Renders a roadmap.yml file into the provided format.",
			Subcommands: cli.Commands{
				&graphvizRenderCommand,
				&htmlRenderCommand,
				&mdRenderCommand,
				&frontmatterMdRenderCommand,
			},
		},
		&validateCommand,
	},
}

func App() *cli.App {
	return app
}
