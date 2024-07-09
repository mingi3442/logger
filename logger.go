package logger

import (
  "fmt"
  "time"
)

// logLevel type
type logLevel int

// Log levels
const (
  trace logLevel = iota
  debug
  info
  notice
  warn
  error
  critical
  fatal
)

const (
  colorReset   = "\033[0m"
  colorWhite   = "\033[37m"
  colorBlue    = "\033[34m"
  colorGreen   = "\033[32m"
  colorCyan    = "\033[36m"
  colorYellow  = "\033[33m"
  colorRed     = "\033[31m"
  colorMagenta = "\033[35m"
  colorBgRed   = "\033[41m"
)

func (l logLevel) colorLogLevel() string {
  switch l {
  case trace:
    return colorWhite + "[TRACE]" + colorReset
  case debug:
    return colorBlue + "[DEBUG]" + colorReset
  case info:
    return colorGreen + "[INFO]" + colorReset
  case notice:
    return colorCyan + "[NOTICE]" + colorReset
  case warn:
    return colorYellow + "[WARN]" + colorReset
  case error:
    return colorRed + "[ERROR]" + colorReset
  case critical:
    return colorMagenta + "[CRITICAL]" + colorReset
  case fatal:
    return colorBgRed + "[FATAL]" + colorReset
  default:
    return colorWhite + "[UNKNOWN]" + colorReset
  }
}

func (l logLevel) colorTimestamp() string {
  switch l {
  case trace:
    return colorWhite
  case debug:
    return colorBlue
  case info:
    return colorGreen
  case notice:
    return colorCyan
  case warn:
    return colorYellow
  case error:
    return colorRed
  case critical:
    return colorMagenta
  case fatal:
    return colorBgRed
  default:
    return colorWhite
  }
}

// LogMessage struct
type LogMessage struct {
  Level   logLevel
  Message string
}

// logChan is a channel for log messages
var logChan chan LogMessage

// package가 읽혀질 때 최초 시작
func init() {
  logChan = make(chan LogMessage)
  go logPrint(logChan)
}

func logPrint(logChan <-chan LogMessage) {
  for logMsg := range logChan {
    formattedMsg := formatLogMessage(logMsg)
    fmt.Println(formattedMsg)
  }
}

func formatLogMessage(logMsg LogMessage) string {
  timestamp := logMsg.Level.colorTimestamp() + time.Now().Format(time.RFC3339) + colorReset
  return fmt.Sprintf("%s %s: %s", timestamp, logMsg.Level.colorLogLevel(), logMsg.Message)
}

func log(level logLevel, msg string) {
  logChan <- LogMessage{Level: level, Message: msg}
}

func Trace(format string, args ...interface{}) {
  log(trace, fmt.Sprintf(format, args...))
}

func Debug(format string, args ...interface{}) {
  log(debug, fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
  log(info, fmt.Sprintf(format, args...))
}

func Notice(format string, args ...interface{}) {
  log(notice, fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
  log(warn, fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
  log(error, fmt.Sprintf(format, args...))
}

func Critical(format string, args ...interface{}) {
  log(critical, fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...interface{}) {
  log(fatal, fmt.Sprintf(format, args...))
}
