package main

import (
    "os"
)

type NaiveLogger struct {
    file *os.File
}

func NewNaiveLogger(filename string) *NaiveLogger {
    f, _ := os.Create(filename)
    return &NaiveLogger{file: f}
}

func (l *NaiveLogger) Log(entry LogEntry) {
    l.file.WriteString(entry.String())
    l.file.Sync()
}

func (l *NaiveLogger) Close() {
    l.file.Close()
}