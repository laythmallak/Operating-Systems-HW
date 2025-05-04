package benchmark

import (
    "fmt"
    "math/rand"
    "raid-sim/raid"
    "time"
)

func RunBenchmark(r raid.RAID, blockSize, totalSize int) {
    numBlocks := totalSize / blockSize
    sample := make([]byte, blockSize)
    rand.Read(sample)

    start := time.Now()
    for i := 0; i < numBlocks; i++ {
        r.Write(i, sample)
    }
    writeTime := time.Since(start)

    start = time.Now()
    for i := 0; i < numBlocks; i++ {
        r.Read(i)
    }
    readTime := time.Since(start)

    fmt.Printf("Write time: %v\n", writeTime)
    fmt.Printf("Read time: %v\n", readTime)
    fmt.Printf("Blocks: %d, Block size: %d, Total: %d MB\n\n", numBlocks, blockSize, totalSize/1024/1024)
}
