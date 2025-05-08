package main

import (
    "os"
    "sync"
)

type MutexLogger struct {
    file     *os.File
    mutex    sync.Mutex
    buffer   []string
    count    int
    batchMax int
}

func NewMutexLogger(filename string) *MutexLogger {
    f, _ := os.Create(filename)
    return &MutexLogger{
        file:     f,
        buffer:   make([]string, 0, 10),
        batchMax: 10,
    }
}

func (l *MutexLogger) Log(entry LogEntry) {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    l.buffer = append(l.buffer, entry.String())
    l.count++

    if l.count == l.batchMax {
        for _, line := range l.buffer {
            l.file.WriteString(line)
        }
        l.file.Sync()
        l.buffer = l.buffer[:0]
        l.count = 0
    }
}

func (l *MutexLogger) Close() {
    l.mutex.Lock()
    for _, line := range l.buffer {
        l.file.WriteString(line)
    }
    l.file.Sync()
    l.file.Close()
    l.mutex.Unlock()
}