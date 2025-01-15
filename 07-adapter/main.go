package main

import "adapter/weirdlogger"

type Logger interface {
	Log(message string)
}
type LoggerAdapter struct {
	ThirdParty *weirdlogger.ThirdPartyLogger
}

func (a *LoggerAdapter) Log(message string) {
	a.ThirdParty.LogMessage(message)
}

func main() {
	thirdPartyLogger := &weirdlogger.ThirdPartyLogger{}
	adapter := &LoggerAdapter{ThirdParty: thirdPartyLogger}

	var logger Logger = adapter
	logger.Log("This is a log message.") // No direct dependency on ThirdPartyLogger
}
