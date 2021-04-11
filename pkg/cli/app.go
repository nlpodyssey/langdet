// Copyright (c) 2021, The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"github.com/nlpodyssey/langdet/pkg/configuration"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

// App is the main structure of the CLI application.
type App struct {
	*cli.App
	config *configuration.AppConfig
	logger zerolog.Logger
}

var appFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "loglevel",
		Usage:    `Minimum severity level for log messages: "error", "warn", "info", "debug", or "trace"`,
		Required: false,
		Value:    "info",
	},
}

// NewApp creates a new App.
func NewApp() *App {
	app := &App{
		App: &cli.App{
			HelpName:  "langdet",
			Usage:     "A simple language detection service",
			Flags:     appFlags,
			Reader:    os.Stdin,
			Writer:    os.Stdout,
			ErrWriter: os.Stderr,
		},
		config: new(configuration.AppConfig),
	}
	app.Before = app.makeBeforeFunc()
	app.Commands = app.makeAppCommands()
	return app
}

func (app *App) makeBeforeFunc() cli.BeforeFunc {
	return func(c *cli.Context) error {
		logLevel := c.String("loglevel")
		var err error
		app.config.LogLevel, err = zerolog.ParseLevel(logLevel)
		if err != nil {
			return fmt.Errorf("invalid loglevel %#v", logLevel)
		}
		app.logger = newLogger(app.config.LogLevel)
		return nil
	}
}

func (app *App) makeAppCommands() []*cli.Command {
	return []*cli.Command{
		app.makeServeCommand(),
		app.makeTrainCommand(),
	}
}

func newLogger(level zerolog.Level) zerolog.Logger {
	w := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}
	return zerolog.New(w).With().Timestamp().Logger().Level(level)
}
