package log

import (
	"github.com/RichardKnop/logging"
	cron2 "github.com/robfig/cron"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger = logging.New(nil, nil, new(logging.ColouredFormatter))

	INFO    = logger[logging.INFO]
	WARNING = logger[logging.WARNING]
	ERROR   = logger[logging.ERROR]
	FATAL   = logger[logging.FATAL]
	DEBUG   = logger[logging.DEBUG]
)

func init() {
	rollingLog := &lumberjack.Logger{
		Filename:   GetLogFilePath("machinery"),
		MaxSize:    0x1000 * 5, // automatic rolling file on it increment than 2GB,
		MaxBackups: 60,         // reserve last 60 day logs
		LocalTime:  true,
		Compress:   true,
	}

	logger = logging.New(rollingLog, rollingLog, new(logging.ColouredFormatter))

	// DEBUG ...
	DEBUG = logger[logging.DEBUG]
	// INFO ...
	INFO = logger[logging.INFO]
	// WARNING ...
	WARNING = logger[logging.WARNING]
	// ERROR ...
	ERROR = logger[logging.ERROR]
	// FATAL ...
	FATAL = logger[logging.FATAL]

	cron := cron2.New()
	err := cron.AddFunc("0 0 23 * * ?", func() {
		_ = rollingLog.Rotate()
	})
	if err != nil {
		panic(err)
	}
	cron.Start()
}

// Set sets a custom logger for all log levels
func Set(l logging.LoggerInterface) {
	DEBUG = l
	INFO = l
	WARNING = l
	ERROR = l
	FATAL = l
}

// SetDebug sets a custom logger for DEBUG level logs
func SetDebug(l logging.LoggerInterface) {
	DEBUG = l
}

// SetInfo sets a custom logger for INFO level logs
func SetInfo(l logging.LoggerInterface) {
	INFO = l
}

// SetWarning sets a custom logger for WARNING level logs
func SetWarning(l logging.LoggerInterface) {
	WARNING = l
}

// SetError sets a custom logger for ERROR level logs
func SetError(l logging.LoggerInterface) {
	ERROR = l
}

// SetFatal sets a custom logger for FATAL level logs
func SetFatal(l logging.LoggerInterface) {
	FATAL = l
}
