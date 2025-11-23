package cmd

import (
	"github.com/urfave/cli/v3"

	ver "github.com/library_owner/library_name/internal/version"
)

func Lefthook() *cli.Command {
	return &cli.Command{
		Name:                  "library_name",
		Usage:                 "Conventional Commits prompt generator",
		Version:               ver.Version(),
		Commands:              commands,
		EnableShellCompletion: true,
	}
}
