package Common

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	StartTime   = time.Now()
	logFilePath string
)

// SetLogFile sets the file path to log to
func SetLogFile(path string) {
	logFilePath = path
	_ = os.MkdirAll(filepath.Dir(path), 0755)
	_ = os.WriteFile(path, []byte{}, 0644)
}

func logLine(level, service, msg string) {
	line := fmt.Sprintf("(%s || %s) - %s [%s] - %s\n",
		time.Now().Format(time.RFC3339),
		time.Since(StartTime).String(),
		level,
		service,
		msg,
	)
	_ = os.WriteFile(logFilePath, []byte(line), os.ModeAppend|0644)
}

func Error(msg string, err error, service string) {
	logLine("ERR", service, fmt.Sprintf("%s - %v", msg, err))
}

func Warn(msg string, err error, service string) {
	logLine("WARN", service, fmt.Sprintf("%s - %v", msg, err))
}
