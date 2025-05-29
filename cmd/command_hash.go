package main

import (
	"context"
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/berybox/genCRC32/pkg/gen32"
	"github.com/urfave/cli/v3"
)

var (
	flagPolynomial = &cli.StringFlag{Name: "polynomial", Aliases: []string{"p"}, Required: false, Usage: "CRC32 polynomial. Can use ID from list, hex (as 0x...) or integer", Value: "0"}
	flagSequence   = &cli.StringFlag{Name: "sequence", Aliases: []string{"s"}, Required: true, Usage: "DNA sequence to hash"}

	commandHash = &cli.Command{
		Name:    "hash",
		Aliases: []string{"h"},
		Usage:   "Hash DNA sequence string using genCRC32",
		Flags: []cli.Flag{
			flagPolynomial,
			flagSequence,
		},

		Action: func(_ context.Context, c *cli.Command) error {
			input := []byte(c.String(flagSequence.Name))

			polynomial, err := parsePolynomial(c.String(flagPolynomial.Name))
			if err != nil {
				return err
			}

			table := crc32.MakeTable(polynomial)

			preproc := gen32.PreprocessBytes(input)

			checksum := crc32.Checksum(preproc, table)

			fmt.Printf("%x\n", checksum)

			return nil
		},
	}
)

func parsePolynomial(poly string) (uint32, error) {
	i, err := strconv.ParseInt(poly, 0, 32)
	if err != nil {
		return 0, err
	}

	if i >= 0 && i < int64(len(gen32.CRC32Polynomials)) {
		return gen32.CRC32Polynomials[i], nil
	}

	return uint32(i), nil
}
