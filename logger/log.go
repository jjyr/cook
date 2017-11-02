package logger

import log "github.com/sirupsen/logrus"

func New() (*log.Logger) {
	return log.New()
}
