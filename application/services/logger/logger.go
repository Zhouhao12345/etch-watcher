package logger

import (
	"log"
	"os"
)

const (
	FilePath = "/tmp/watch.log"
)

type logging struct {
	filename string
	*log.Logger
}

var Logger *logging

func NewLogger(filename string) (*logging, error) {
	var (
		logFile *os.File
		err error
	)
	if logFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		return nil, err
	} else {
		return &logging{
			filename: filename,
			Logger:   log.New(logFile, "Etcd Watcher Applicant: ", log.Lshortfile),
		}, nil
	}
}


func init()  {
	var err error
	if Logger, err = NewLogger(FilePath); err != nil {
		panic(err)
	}
}
