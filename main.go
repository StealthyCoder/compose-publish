package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	commandLine "github.com/urfave/cli/v2"

	"github.com/foundriesio/compose-publish/pkg"
)

const banner = `
	   |\/|
	\__|__|__/

`

func main() {
	var file string
	var digestFile string
	var dryRun bool

	fmt.Print(banner)
	app := &commandLine.App{
		Name:  "compose-ref",
		Usage: "Reference Compose Specification implementation",
		Flags: []commandLine.Flag{
			&commandLine.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Value:       "docker-compose.yml",
				Usage:       "Load Compose file `FILE`",
				Destination: &file,
			},
			&commandLine.StringFlag{
				Name:        "digest-file",
				Aliases:     []string{"d"},
				Required:    false,
				Usage:       "Save sha256 digest of bundle to a file",
				Destination: &digestFile,
			},
			&commandLine.BoolFlag{
				Name:        "dryrun",
				Required:    false,
				Usage:       "Show what would be done, but don't actually publish",
				Destination: &dryRun,
			},
		},
		Action: func(c *commandLine.Context) error {
			target := c.Args().Get(0)
			if len(target) == 0 {
				return errors.New("Missing required argument: TARGET:[TAG]")
			}
			return pkg.DoPublish(file, target, digestFile, dryRun)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
