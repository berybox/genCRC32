package main

import (
	"context"
	"fmt"

	"github.com/berybox/genCRC32/pkg/gen32"
	"github.com/urfave/cli/v3"
)

var (
	commandList = &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List appropriate CRC32 polynomials",
		Action: func(_ context.Context, _ *cli.Command) error {

			for i, poly := range gen32.CRC32Polynomials {
				fmt.Printf("%2d: 0x%x\n", i, poly)
			}

			return nil
		},
	}
)
