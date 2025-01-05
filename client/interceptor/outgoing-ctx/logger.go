package outgoing_ctx

import (
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the logger for the outgoing context debugger
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger is the logger for the outgoing context debugger
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Initialize the mode named logger
	namedLogger, err := gologgermodenamed.NewDefaultLogger(header, modeLogger)
	if err != nil {
		return nil, err
	}

	return &Logger{logger: namedLogger}, nil
}

// LogKeyValue logs the key value
func (l *Logger) LogKeyValue(key string, value string) {
	l.logger.Debug(
		"outgoing context key '"+key+"' value",
		value,
	)
}
