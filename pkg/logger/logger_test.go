package logger

import "testing"

func Test_Logger(t *testing.T) {
	if Logger != nil {
		t.Error("could not hanppen this ~")
	}

	if err := InitLogger("./testdata"); err != nil {
		t.Errorf("could not InitLogger: %v", err)
	}

	Logger.Info("info")
	Logger.Errorf("error: %v", nil)
	Logger.WithField("key", "value").Info("info")
}
