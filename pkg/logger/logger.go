package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Logger is interface of logger
type Logger interface {
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
	Panic(format string, v ...interface{})
	Debug(format string, v ...interface{})
	WithPrefix(prefix string) Logger
	WithRequestID(requestID string) Logger
	WithMethod(method string) Logger
}

// Log is struct for logger
type Log struct {
	zlog      *zerolog.Logger
	requestID string
	method    string
}

// New is constructor for the logger
func New(isDebug ...bool) Log {
	logLevel := zerolog.InfoLevel

	var debug bool
	if isDebug != nil {
		if len(isDebug) != 0 {
			debug = isDebug[0]
		}
	}

	if debug {
		logLevel = zerolog.DebugLevel
	}

	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    true,
		TimeFormat: "2006/01/02 - 15:04:05.000000",
		PartsOrder: []string{
			zerolog.LevelFieldName,
			"request_id",
			zerolog.TimestampFieldName,
			"method",
			zerolog.MessageFieldName,
		},
		FieldsExclude: []string{
			"request_id",
			"method",
		},
		FormatLevel: func(i interface{}) string {
			if l, ok := i.(string); ok {
				switch l {
				case zerolog.LevelTraceValue:
					return "[TRC]"
				case zerolog.LevelDebugValue:
					return "[DBG]"
				case zerolog.LevelInfoValue:
					return "[INF]"
				case zerolog.LevelWarnValue:
					return "[WRN]"
				case zerolog.LevelErrorValue:
					return "[ERR]"
				case zerolog.LevelFatalValue:
					return "[FTL]"
				case zerolog.LevelPanicValue:
					return "[PNC]"
				default:
					return "[???]"
				}
			} else {
				return "[???]"
			}
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s", i)
		},
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(writer).With().Timestamp().Logger()

	return Log{zlog: &logger}
}

// WithPrefix setting a prefix for log message
func (ths Log) WithPrefix(prefix string) Logger {
	ths.requestID = prefix
	return ths
}

// WithRequestID setting a request id for log message
func (ths Log) WithRequestID(reqID string) Logger {
	ths.requestID = reqID
	return ths
}

// WithMethod setting a method for log message
func (ths Log) WithMethod(method string) Logger {
	ths.method = method
	return ths
}

// Info logs a message with info level
func (ths Log) Info(format string, v ...interface{}) {
	ths.prepareMessage(ths.zlog.Info(), format, v...)
}

// Warn logs a message with warn level
func (ths Log) Warn(format string, v ...interface{}) {
	ths.prepareMessage(ths.zlog.Warn(), format, v...)
}

// Warn logs a message with warn level
func (ths Log) Error(format string, v ...interface{}) {
	ths.prepareMessage(ths.zlog.Error(), format, v...)
}

// Fatal logs a message with fatal level
func (ths Log) Fatal(format string, v ...interface{}) {
	ths.prepareMessage(ths.zlog.Fatal(), format, v...)
}

// Panic logs a message with panic level
func (ths Log) Panic(format string, v ...interface{}) {
	ths.prepareMessage(ths.zlog.Panic(), format, v...)
}

// Debug logs a message with debug level
func (ths Log) Debug(format string, v ...interface{}) {
	ths.prepareMessage(ths.zlog.Debug(), format, v...)
}

// prepareMessage prepares the message for log
func (ths Log) prepareMessage(event *zerolog.Event, format string, v ...interface{}) {
	if len(ths.method) != 0 {
		event.Str("method", fmt.Sprintf("| %s", ths.method))
	} else {
		event.Str("method", "")
	}

	if len(ths.requestID) != 0 {
		event.Str("request_id", fmt.Sprintf("[%s]", ths.requestID))
	} else {
		event.Str("request_id", "")
	}

	if v != nil {
		event.Msgf(format, v...)
		return
	}

	event.Msg(format)
}
