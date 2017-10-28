package qloger

import (
	"testing"
	"time"
)

func TestDefaultLogOutput(t *testing.T) {
	logger, err := NewQLogger("/tmp/test.log", "info")
	if err != nil {
		t.Fatal(err.Error())
	}

	logger.Infoln("world")
}

func TestRotatorLogOutput(t *testing.T) {
	logger, err := NewRotatorQLogger("/tmp/test.log", "debug", true, true, true, 1024*5)
	if err != nil {
		t.Fatal(err.Error())
	}

	for i := 0; i < 1000; i++ {
		logger.Infoln("hello")
		time.Sleep(80 * time.Millisecond)
	}
}
