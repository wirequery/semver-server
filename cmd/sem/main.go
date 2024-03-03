package main

import (
	"github.com/urfave/cli/v2"
	"github.com/wirequery/semver-server/internal/server"
	"github.com/wirequery/semver-server/pkg/store"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "sem",
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "port",
					},
				},
				Action: func(c *cli.Context) error {
					port := 12345
					if c.Int("port") != 0 {
						port = c.Int("port")
					}
					instance := server.Server{
						Port:         port,
						VersionStore: store.NewInMemoryStore(),
					}
					instance.Serve()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
