package commands

import (
	"encoding/hex"
	"fmt"

	"github.com/evercoinx/bitcoin/internal/encoding"
	"github.com/urfave/cli/v2"
)

func New() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "converts hash of public key or script to bitcoin address",
			Action: func(ctx *cli.Context) error {
				publicKeyHash := ctx.Args().First()
				if len(publicKeyHash) != 40 {
					return fmt.Errorf("invalid hash is specified: %s", publicKeyHash)
				}

				payload, err := hex.DecodeString(publicKeyHash)
				if err != nil {
					return fmt.Errorf("unable to decode hash.\ncause: %w", err)
				}

				var ver encoding.AddressVersion
				switch ctx.String("address-type") {
				case "p2pkh":
					ver = encoding.AddressVersionPublicKeyHash
				case "p2sh":
					ver = encoding.AddressVersionScriptHash
				default:
					return fmt.Errorf("invalid address type is specified: %s", ver)
				}

				addr := encoding.Base58CheckEncode(payload, ver)
				fmt.Printf("bitcoin address: %s\n", addr)
				return nil
			},
		},
	}
}
