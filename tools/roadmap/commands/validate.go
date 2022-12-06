package commands

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/SierraSoftworks/roadmap"
	"github.com/SierraSoftworks/roadmap/tools/roadmap/commands/linting"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var validateCommand = cli.Command{
	Name:      "validate",
	Aliases:   []string{"lint"},
	Usage:     "Validates that a roadmap.yml file conforms to the standard schema and reports any linter issues.",
	ArgsUsage: "./roadmap.yml",
	Action: func(c *cli.Context) error {
		f, err := ioutil.ReadFile(c.Args().First())
		if err != nil {
			return err
		}

		r, err := roadmap.Parse(f)
		if err != nil {
			return err
		}

		err = r.Validate()
		if err != nil {
			return err
		}

		issues := linting.Validate(r)
		fatal := false
		for _, issue := range issues {
			fmt.Print("[")
			switch issue.Level {
			case linting.LEVEL_ERROR:
				fmt.Print(color.RedString("%s", issue.Level))
			case linting.LEVEL_WARNING:
				fmt.Print(color.YellowString("%s", issue.Level))
			case linting.LEVEL_INFO:
				fmt.Print(color.BlueString("%s", issue.Level))
			}

			fatal = fatal || issue.IsError()

			fmt.Printf("]: %s\n", issue.Message)
		}

		if fatal {
			return errors.New("linter validation failed")
		}

		if len(issues) == 0 {
			fmt.Println("Validation succeeded!")
		}

		return nil
	},
}
