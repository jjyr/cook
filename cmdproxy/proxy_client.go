package cmdproxy

import (
	"io"
)

type ProxyClient interface {
	Start(name string, args ... string) error
	Wait() error
	Run(name string, args ... string) error
	StdinPipe() io.WriteCloser
	StderrPipe() io.Reader
	StdoutPipe() io.Reader
	Close() error
}
