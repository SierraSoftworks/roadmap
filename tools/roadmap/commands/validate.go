package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/SierraSoftworks/roadmap"
	"github.com/urfave/cli/v2"
)

var validateCommand = cli.Command{
	Name:      "validate",
	Aliases:   []string{"lint"},
	Usage:     "Validates that a roadmap.yml file conforms to the standard schema.",
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

		fmt.Println("Validation succeeded!")

		return nil
	},
}
