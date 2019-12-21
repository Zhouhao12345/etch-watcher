package logger

import (
	"log"
	"os"
	"zhouhao.com/elevator/application/services/config"
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
	var filePath string
	var err error

	filePath = config.LoggerConfig.GetFilePath()
	if Logger, err = NewLogger(filePath); err != nil {
		panic(err)
	} else {
		log.Print("Logger Service Init Success")
	}
}
