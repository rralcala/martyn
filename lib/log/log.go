package log

import (
	"fmt"
	"os"
)

func Info(message string) {
	fmt.Println("INFO:  " + message)
}

func Warning(message string) {
	fmt.Println("WARN:  " + message)
}

func Error(message string) {
	fmt.Println("ERROR: " + message)
}

func Fatal(message string) {
	fmt.Println("FATAL: " + message)
	os.Exit(1)
}
