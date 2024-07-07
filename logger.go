package logger

import (
  "fmt"
  "time"
)

// logLevel type
type logLevel int

// Log levels
const (
  TRACE logLevel = iota
  DEBUG
  INFO
  NOTICE
  WARN
  ERROR
  CRITICAL
  FATAL
)

func (l logLevel) colorLogLevel() string {
  switch l {
  case TRACE:
    return "\033[37m[TRACE]\033[0m" // 흰색
  case DEBUG:
    return "\033[34m[DEBUG]\033[0m" // 파란색
  case INFO:
    return "\033[32m[INFO]\033[0m" // 녹색
  case NOTICE:
    return "\033[36m[NOTICE]\033[0m" // 청록색
  case WARN:
    return "\033[33m[WARN]\033[0m" // 노란색
  case ERROR:
    return "\033[31m[ERROR]\033[0m" // 빨간색
  case CRITICAL:
    return "\033[35m[CRITICAL]\033[0m" // 자주색
  case FATAL:
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
  log(TRACE, fmt.Sprintf(format, args...))
}

func Debug(format string, args ...interface{}) {
  log(DEBUG, fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
  log(INFO, fmt.Sprintf(format, args...))
}

func Notice(format string, args ...interface{}) {
  log(NOTICE, fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
  log(WARN, fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
  log(ERROR, fmt.Sprintf(format, args...))
}

func Critical(format string, args ...interface{}) {
  log(CRITICAL, fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...interface{}) {
  log(FATAL, fmt.Sprintf(format, args...))
}
