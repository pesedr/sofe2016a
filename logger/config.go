package logger

import (
	"io"
	"log"
)

const (
	Dev  = "development"
	Prod = "production"
)

var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

func InitLogger(info, err, debug io.Writer) {
	Info = log.New(info, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(err, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(debug, "DEBUG", log.Ldate|log.Ltime|log.Lshortfile)
}
