package logger

import (
	"errors"
	"log"
	"os"
)

type FileLogger struct {
	file   *os.File
	logger *log.Logger
	isOpen bool
}

func NewLogger(filename string, prefix string) (FileLogger, error) {
	fl := FileLogger{isOpen: false}
	f, err := os.OpenFile(filename+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fl, err
	}
	fl.file = f
	fl.logger = log.New(f, prefix, log.LstdFlags)
	fl.isOpen = true
	return fl, nil
}

func (fl *FileLogger) Println(message string) error {
	log.Println(message)
	if fl.isOpen {
		fl.logger.Println(message)
	} else {
		return errors.New("file connection is closed")
	}
	return nil
}

func (fl *FileLogger) EndLogger() {
	if fl.isOpen {
		_ = fl.file.Close()
		fl.isOpen = false
	}
}

func SingleLog(filename string, prefix string, message string) error {
	log.Println(message)
	f, err := os.OpenFile(filename+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logger := log.New(f, prefix, log.LstdFlags)
	logger.Println(message)
	return nil
}
