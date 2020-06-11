// +build ignore

package main

import (
	"github.com/baojiweicn/Surge/util/log"
)

var logger = log.Get("ExampleName")

func main() {
	logger.SetLevel(log.INFO)
	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Warnf("This is a number %v", 1)
}
