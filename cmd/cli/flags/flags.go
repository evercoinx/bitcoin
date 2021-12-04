package flags

import "github.com/urfave/cli/v2"

func New() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "address-type",
			Aliases: []string{"at"},
			Value:   "p2pkh",
			Usage:   "address type: p2pkh or p2sh",
		},
	}
}
