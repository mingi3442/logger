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

func (l logLevel) colorLogLevel() string {
  switch l {
  case trace:
    return "\033[37m[TRACE]\033[0m" // 흰색
  case debug:
    return "\033[34m[DEBUG]\033[0m" // 파란색
  case info:
    return "\033[32m[INFO]\033[0m" // 녹색
  case notice:
    return "\033[36m[NOTICE]\033[0m" // 청록색
  case warn:
    return "\033[33m[WARN]\033[0m" // 노란색
  case error:
    return "\033[31m[ERROR]\033[0m" // 빨간색
  case critical:
    return "\033[35m[CRITICAL]\033[0m" // 자주색
  case fatal:
    return "\033[41m[FATAL]\033[0m" // 빨간 배경
  default:
    return "\033[37m[UNKNOWN]\033[0m" // 흰색
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
  // colorCode := logMsg.Level.Color()
  return fmt.Sprintf("%s %s: %s%s", logMsg.Level.colorLogLevel(), time.Now().Format(time.RFC3339), logMsg.Message, "\033[0m")
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
