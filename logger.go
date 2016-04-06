package hippo

import (
	"log"
	"time"
)

type Logger struct{}

func (l *Logger) Handle(ct *Context, next HandlerInterface) *Error {
	start := time.Now()
	err := next.Handle(ct)
	log.Printf("%s %s %s", ct.Request.Method, ct.Request.URL.Path, time.Since(start))
	return err
}

func NewLogger() *Logger {
	return &Logger{}
}
