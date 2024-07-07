package main

import "github.com/mingi3442/logger"

func main() {
  logger.Trace("This is a trace message")
  logger.Debug("This is a debug message")
  logger.Info("This is an info message")
  logger.Notice("This is a notice message")
  logger.Warn("This is a warning message")
  logger.Error("This is an error message")
  logger.Critical("This is a critical message")
  logger.Fatal("This is a fatal message")
}
