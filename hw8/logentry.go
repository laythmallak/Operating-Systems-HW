package main

import (
    "fmt"
    "time"
)

type LogEntry struct {
    Level   string
    Context string
    Message string
}

func (le LogEntry) String() string {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    return fmt.Sprintf("[%s] [%s] [%s] %s\n", timestamp, le.Level, le.Context, le.Message)
}