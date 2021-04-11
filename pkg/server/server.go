// Copyright (c) 2021, The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"github.com/nlpodyssey/langdet/pkg/api"
	"github.com/nlpodyssey/langdet/pkg/configuration"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is the main implementation of api.ApiServer.
type Server struct {
	api.UnimplementedApiServer
	config *configuration.ServerConfig
	logger zerolog.Logger
}

// New creates a new Server.
func New(config *configuration.ServerConfig, logger zerolog.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

// DetectLanguage detects the language of a text.
func (s *Server) DetectLanguage(context.Context, *api.DetectLanguageRequest) (*api.DetectLanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectLanguage not implemented")
}
