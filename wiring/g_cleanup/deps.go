package main

import (
	"errors"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Filepath string

type FileLogger struct {
	fp *os.File
}

func NewFileLogger(filepath Filepath) (*FileLogger, func(), error) {
	fp, err := os.OpenFile(string(filepath), os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = fp.Close()
		_ = os.Remove(string(filepath))
	}

	return &FileLogger{fp: fp}, cleanup, nil
}

type HalfLiveCat struct {
	log *FileLogger
}

func NewHalfLiveCat(log *FileLogger) (*HalfLiveCat, error) {
	if rand.Intn(100)%2 == 0 {
		return &HalfLiveCat{log: log}, nil
	}
	return nil, errors.New("always error")
}
