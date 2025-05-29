// main executable
package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "genCRC32",
		Usage: "Collision-free CRC32 based hashing of DNA sequences",
		Commands: []*cli.Command{
			commandList,
			commandHash,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
