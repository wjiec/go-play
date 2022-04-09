// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"bytes"
)

// Injectors from wire.go:

func NewAlerter() Alerter {
	writer := _wireBufferValue
	message := _wireMessageValue
	simpleAlerter := NewSimpleAlerter(writer, message)
	return simpleAlerter
}

var (
	_wireBufferValue  = &bytes.Buffer{}
	_wireMessageValue = Message("alert message")
)
