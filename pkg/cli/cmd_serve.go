// Copyright (c) 2021, The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"github.com/nlpodyssey/langdet/pkg/configuration"
	"github.com/nlpodyssey/langdet/pkg/server"
	"github.com/urfave/cli/v2"
)

var serveFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "host",
		Usage: "server binding address",
		Value: "0.0.0.0",
	},
	&cli.IntFlag{
		Name:  "port",
		Usage: "server listening port",
		Value: 3000,
	},
	&cli.BoolFlag{
		Name:  "tls",
		Usage: "whether to enable TLS",
		Value: false,
	},
	&cli.StringFlag{
		Name:  "tls-cert",
		Usage: `TLS cert file (ignored if tls is disabled)`,
		Value: "server.crt",
	},
	&cli.StringFlag{
		Name:  "tls-key",
		Usage: `TLS key file (ignored if tls is disabled)`,
		Value: "server.key",
	},
}

func (app *App) makeServeCommand() *cli.Command {
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Runs the language detection server",
		Action:  app.serve,
		Flags:   serveFlags,
	}
}

func (app *App) serve(c *cli.Context) error {
	config := makeServerConfig(c)
	srv := server.New(config, app.logger)
	return srv.Run()
}

func makeServerConfig(c *cli.Context) *configuration.ServerConfig {
	return &configuration.ServerConfig{
		Host:       c.String("host"),
		Port:       c.Int("port"),
		TLSEnabled: c.Bool("tls"),
		TLSCert:    c.String("tls-cert"),
		TLSKey:     c.String("tls-key"),
	}
}
