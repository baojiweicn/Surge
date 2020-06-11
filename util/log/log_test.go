package log_test

import (
	"testing"

	"github.com/baojiweicn/Surge/util/log"
)

func TestLog(t *testing.T) {
	logger := log.Get("ExampleName")
	// No assertions.
	logger.SetLevel(log.DEBUG)
	logger.Debug(nil)
	logger.Info(nil)
	logger.Warn(nil)
	logger.Error(nil)
	logger.Debugf("hello %s", "world")
	logger.Infof("hello %s", "world")
	logger.Warnf("hello %s", "world")
	logger.Errorf("hello %s", "world")
}
