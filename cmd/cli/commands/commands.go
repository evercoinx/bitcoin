package commands

import (
	"encoding/hex"
	"fmt"

	"github.com/evercoinx/bitcoin/internal/encoding"
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "address",
			Aliases: []string{"a"},
			Subcommands: []*cli.Command{
				{
					Name:    "encode",
					Aliases: []string{"e"},
					Usage:   "encode hash of public key or script to bitcoin address",
					Action:  encodeHashToAddress,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "address-type",
							Aliases: []string{"t"},
							Value:   "p2pkh",
							Usage:   "address type: p2pkh or p2sh",
						},
					},
				},
				{
					Name:    "decode",
					Aliases: []string{"d"},
					Usage:   "decode bitcoin address to hash of public key or script",
					Action:  decodeAddressToHash,
				},
			},
		},
	}
}

func encodeHashToAddress(ctx *cli.Context) error {
	hash := ctx.Args().First()
	if len(hash) != 40 {
		return fmt.Errorf("invalid hash is specified: %s", hash)
	}

	payload, err := hex.DecodeString(hash)
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
	fmt.Printf("address: %s\n", addr)
	return nil
}

func decodeAddressToHash(ctx *cli.Context) error {
	addr := ctx.Args().First()
	if len(addr) < 14 || len(addr) > 74 {
		return fmt.Errorf("invalid address is specified: %s", addr)
	}

	hash, err := encoding.Base58CheckDecode(addr)
	if err != nil {
		return fmt.Errorf("unable to decode address.\ncause: %w", err)
	}

	fmt.Printf("hash: %x\n", hash)
	return nil
}
