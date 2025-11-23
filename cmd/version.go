package cmd

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"

	ver "github.com/library_owner/library_name/internal/version"
)

func version() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "print version",
		Flags: []cli.Flag{},
		Action: func(_ context.Context, _ *cli.Command) error {
			log.Println(ver.Version())

			return nil
		},
	}
}
