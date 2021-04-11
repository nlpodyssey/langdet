// Copyright (c) 2021, The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package configuration

import (
	"github.com/rs/zerolog"
)

// AppConfig provides a minimum set of configuration parameters
// which are shared among different commands.
type AppConfig struct {
	// LogLevel is the minimum severity level for log messages.
	LogLevel zerolog.Level
}

// ServerConfig is the server configuration.
type ServerConfig struct {
	// Host is the server binding address.
	Host string
	// Port is the server listening port.
	Port int
	// TLSEnabled reports whether to enable TLS.
	TLSEnabled bool
	// TLSCert is the TLS cert file. It is ignored if TLSEnabled is false.
	TLSCert string
	// TLSKey is the TLS key file. It is ignored if TLSEnabled is false.
	TLSKey string
}
