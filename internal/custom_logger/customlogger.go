package customlogger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	filePath string
}

func NewLogger(filePath string) (*Logger, error) {
	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		// If the file doesn't exist, create a new one
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to create log file: %s", err)
			}
			file.Close()
		} else {
			return nil, fmt.Errorf("failed to check log file status: %s", err)
		}
	}

	return &Logger{filePath: filePath}, nil
}

func (l *Logger) Log(message string) {
	// Open the log file in append mode
	file, err := os.OpenFile(l.filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %s\n", err)
		return
	}
	defer file.Close()

	// Get the current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Construct the log entry
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)

	// Write the log entry to the file
	_, err = file.WriteString(logEntry)
	if err != nil {
		fmt.Printf("Failed to write log entry: %s\n", err)
		return
	}
}