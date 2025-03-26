package main

import (
	"fmt"
	"os"
)

type ILoxLogger interface {
	Info(message string)
	debug(message string)
	_Error(message string)
}

type LoxLogger struct{}

func (logger *LoxLogger) Info(message string) {
	fmt.Fprintln(os.Stdin, message)
}

func (logger *LoxLogger) _Error(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func NewLogger() *LoxLogger {
	return &LoxLogger{}
}
