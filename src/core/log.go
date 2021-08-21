package core

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func (level LogLevel) String() string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	}
	panic("wrong log level")
}

type Logger struct {
	manager EventManagerI
	w       io.Writer
	level   LogLevel
}

func NewLogger(manager EventManagerI, w io.Writer, level LogLevel) *Logger {
	return &Logger{
		manager: manager,
		w:       w,
		level:   level,
	}
}

func (logger *Logger) write(level LogLevel, format string, args ...interface{}) {
	time := logger.manager.GetCurrentTimeOffset()
	_, file, line, _ := runtime.Caller(2)

	idx := strings.LastIndexByte(file, '/')
	idx = strings.LastIndexByte(file[:idx], '/')
	file = file[idx+1:]

	str := fmt.Sprintf(format, args...)

	str = fmt.Sprintf("%s %s %s:%d %s\n", time, level, file, line, str)

	logger.w.Write([]byte(str))
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.level <= LogLevelDebug {
		logger.write(LogLevelDebug, format, args...)
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.level <= LogLevelInfo {
		logger.write(LogLevelInfo, format, args...)
	}
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.level <= LogLevelWarn {
		logger.write(LogLevelWarn, format, args...)
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.level <= LogLevelError {
		logger.write(LogLevelError, format, args...)
	}
}
