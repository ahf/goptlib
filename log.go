package pt

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type LogSeverity int

const (
	Debug   LogSeverity = 0
	Info    LogSeverity = 1
	Notice  LogSeverity = 2
	Warning LogSeverity = 3
	Error   LogSeverity = 4
)

func (log_severity LogSeverity) String() string {
	labels := [...]string{
		"debug",
		"info",
		"notice",
		"warning",
		"error",
	}

	value := log_severity

	// Smaller values than Debug? Reset to Debug.
	if value < Debug {
		value = Debug
	}

	// Larger values than Error? Reset to Error.
	if value > Error {
		value = Error
	}

	return labels[value]
}

type logger struct {
	writer   io.Writer
	severity LogSeverity
}

func (l logger) Write(p []byte) (n int, err error) {
	data := map[string]string{
		"SEVERITY": l.severity.String(),
		// Remove the trailing new line that the `log` package appends to all strings.
		"MESSAGE": strings.TrimRight(string(p), "\n"),
	}
	log := fmt.Sprintf("LOG %s\n", kvline_encode(data))

	return io.WriteString(l.writer, log)
}

func NewPTLogger(severity LogSeverity, prefix string, flags int) *log.Logger {
	result := log.New(logger{writer: Stdout, severity: severity}, prefix, flags)
	return result
}
