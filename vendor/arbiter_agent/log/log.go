package log

import "github.com/gookit/color"

type Level int

const (
	All Level = 0
	InfoLevel Level = 1
	WarnLevel Level = 2
	ErrorLevel Level = 3
)

var level = All

func Error (printStr string) {
	color.Error.Printf("[Error] %s\n", printStr)
}

func Warn (printStr interface{}) {
	if level <= WarnLevel {
		color.Warn.Printf("[Warn] %s\n", printStr)
	}
}

func Info (printStr interface{}) {
	if level <= InfoLevel {
		color.Info.Printf("[Info] %s\n", printStr)
	}
}

func Debug (printStr interface{}) {
	if level == All {
		color.Cyan.Printf("[Debug] %s\n", printStr)
	}
}

func SetLogLevel (logLevel Level) {
	level = logLevel
}