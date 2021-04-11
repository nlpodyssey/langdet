// Copyright (c) 2021, The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var trainFlags []cli.Flag

func (app *App) makeTrainCommand() *cli.Command {
	return &cli.Command{
		Name:    "train",
		Aliases: []string{"t"},
		Usage:   "Trains a model for language detection",
		Action:  app.train,
		Flags:   trainFlags,
	}
}

func (app *App) train(*cli.Context) error {
	return fmt.Errorf("train command not implemented")
}
