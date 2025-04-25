
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

    fmt.Printf("Write time: %v
", writeTime)
    fmt.Printf("Read time: %v
", readTime)
    fmt.Printf("Blocks: %d, Block size: %d, Total: %d MB

", numBlocks, blockSize, totalSize/1024/1024)
}