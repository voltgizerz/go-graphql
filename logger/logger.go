package logger

import (
	log "github.com/sirupsen/logrus"
)

// Log - logrus logging.
//
//revive:disable:import-shadowing
var (
	Log *log.Logger
)

// SetupLog - return logrus.
func init() {
	Log = log.New()
}
