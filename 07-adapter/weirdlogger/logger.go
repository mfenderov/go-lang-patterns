package weirdlogger

import "fmt"

type ThirdPartyLogger struct{}

func (l *ThirdPartyLogger) LogMessage(message string) {
	fmt.Println("[ThirdPartyLogger]:", message)
}
