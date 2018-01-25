package logUtils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

// Logger is our extension of the log15 Logger type
type Logger struct {
	log15.Logger
	logErrorsWithStackTrace bool
}

const (
	// DebugLevel - duh
	DebugLevel = log15.LvlDebug
	// InfoLevel - duh
	InfoLevel = log15.LvlInfo
	// WarnLevel - duh
	WarnLevel = log15.LvlWarn
	// ErrorLevel - duh
	ErrorLevel = log15.LvlError
	// CritLevel - duh
	CritLevel = log15.LvlCrit
)

var (
	// Always set the default to Info
	globalLogLevel = InfoLevel
)

// SetLevel sets the global log level for log15-based loggers
func SetLevel(level log15.Lvl) {
	globalLogLevel = level
}

// NewLogger creates a new logger with the default handlers
func NewLogger(args ...interface{}) Logger {
	logger := log15.New(args...)
	handler := logger.GetHandler()
	handler = log15.CallerFuncHandler(handler)
	handler = log15.CallerFileHandler(handler)
	// Get the log level from the env
	logLvl := InfoLevel // Default
	env := os.Getenv("CHATBOT_DEBUG")
	if len(env) > 0 {
		lvl, err := strconv.ParseBool(env)
		if lvl == true && err == nil {
			logLvl = DebugLevel
		}
	}
	handler = log15.LvlFilterHandler(logLvl, handler)
	logger.SetHandler(handler)
	logErrorsWithStackTrace, err := strconv.ParseBool(os.Getenv("CHATBOT_LOG_ERRORS_WITH_STACK_TRACE"))
	if err != nil {
		logErrorsWithStackTrace = false
	}
	return Logger{
		Logger:                  logger,
		logErrorsWithStackTrace: logErrorsWithStackTrace,
	}
}

// NewEnvLogger gives a logger using the env var supplied for specifying the loglevel
// If the envVar contains bogus info, defaults to Info level
// The env var should be one of debug | info | warn | error | crit
func NewEnvLogger(lvlEnvVar string, args ...interface{}) Logger {
	logger := log15.New(args...)
	handler := logger.GetHandler()
	handler = log15.CallerFuncHandler(handler)
	handler = log15.CallerFileHandler(handler)
	// Get the log level from the env
	strLvl := os.Getenv(lvlEnvVar)
	if strLvl != "debug" && strLvl != "info" && strLvl != "warn" &&
		strLvl != "error" && strLvl != "crit" {
		strLvl = "info"
	}
	loglvl, _ := log15.LvlFromString(strLvl)
	handler = log15.LvlFilterHandler(loglvl, handler)
	logger.SetHandler(handler)
	logErrorsWithStackTrace, err := strconv.ParseBool(os.Getenv("FD_LOG_ERRORS_WITH_STACK_TRACE"))
	if err != nil {
		logErrorsWithStackTrace = false
	}
	return Logger{
		Logger:                  logger,
		logErrorsWithStackTrace: logErrorsWithStackTrace,
	}
}

// ChildLogger is a method that lets us create child loggers of the same type
func (logger Logger) ChildLogger(ctx ...interface{}) Logger {
	// delegate to log15 to create a new logger
	childLogger := logger.Logger.New(ctx)
	// return the new wrapped Logger
	return Logger{
		Logger:                  childLogger,
		logErrorsWithStackTrace: logger.logErrorsWithStackTrace,
	}
}

// Printf is implemented so that our logger conforms to the GorpLogger
// interface and can be used for logging queries in
// that require gorp.DbMap
// SOURCE: https://godoc.org/github.com/coopernurse/gorp#GorpLogger
// NOTE: https://github.com/go-gorp/gorp
func (logger Logger) Printf(arg1 string, args ...interface{}) {
	output := fmt.Sprintf(arg1, args...)
	logger.Logger.Debug(output)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrorStackTrace logs a stack trace for the given error
func (logger Logger) Error(arg1 string, args ...interface{}) {
	logger.Logger.Error(arg1, args...)

	if logger.logErrorsWithStackTrace {
		for _, arg := range args {
			if err, ok := arg.(stackTracer); ok {
				logger.Logger.Error(fmt.Sprintf("*** Stack trace ***\n%+v", err))
			}
		}
	}
}
