package logging

import "testing"

func TestLogger(t *testing.T) {
	log := GetLogger("module")
	log.Trace("trace")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	log.Fatal("fatal")
}
