package main

import (
    "os"
)

type ChannelLogger struct {
    file     *os.File
    channel  chan LogEntry
    done     chan bool
    buffer   []string
    batchMax int
}

func NewChannelLogger(filename string) *ChannelLogger {
    f, _ := os.Create(filename)
    cl := &ChannelLogger{
        file:     f,
        channel:  make(chan LogEntry, 100),
        done:     make(chan bool),
        batchMax: 10,
    }
    go cl.listen()
    return cl
}

func (l *ChannelLogger) listen() {
    for {
        select {
        case entry := <-l.channel:
            l.buffer = append(l.buffer, entry.String())
            if len(l.buffer) == l.batchMax {
                for _, line := range l.buffer {
                    l.file.WriteString(line)
                }
                l.file.Sync()
                l.buffer = l.buffer[:0]
            }
        case <-l.done:
            for _, line := range l.buffer {
                l.file.WriteString(line)
            }
            l.file.Sync()
            l.file.Close()
            return
        }
    }
}

func (l *ChannelLogger) Log(entry LogEntry) {
    l.channel <- entry
}

func (l *ChannelLogger) Close() {
    l.done <- true
}