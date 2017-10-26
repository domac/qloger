package qloger

import (
	"testing"
)

func TestDefaultLogOutput(t *testing.T) {
	logger, err := NewQLogger("/tmp/test.log", "info")
	if err != nil {
		t.Fatal(err.Error())
	}

	logger.Infoln("world")
}

func TestRotatorLogOutput(t *testing.T) {
	logger, err := NewRotatorQLogger("/tmp/test.log", "debug", true, false, true, 1024*5)
	if err != nil {
		t.Fatal(err.Error())
	}

	logger.Infoln("hello")
}
