package main

import (
    "fmt"
    "sync"
    "time"
)

func runBenchmark(loggerType string, logger interface {
    Log(LogEntry)
    Close()
}) {
    start := time.Now()
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 50; j++ {
                entry := LogEntry{
                    Level:   "INFO",
                    Context: fmt.Sprintf("req-%d", id),
                    Message: fmt.Sprintf("Processing item %d", j),
                }
                logger.Log(entry)
            }
        }(i)
    }

    wg.Wait()
    logger.Close()
    elapsed := time.Since(start)
    fmt.Printf("%s logger completed in %s\n", loggerType, elapsed)
}

func main() {
    runBenchmark("Naive", NewNaiveLogger("naive.log"))
    runBenchmark("Mutex", NewMutexLogger("mutex.log"))
    runBenchmark("Channel", NewChannelLogger("channel.log"))
}