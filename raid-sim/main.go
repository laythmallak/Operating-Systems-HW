package main

import (
    "fmt"
    "os"
    "raid-sim/raid"
    "raid-sim/disk"
    "raid-sim/benchmark"
)

const (
    blockSize = 4096
    totalSize = 100 * 1024 * 1024 // 100MB
)

func main() {
    // Setup disks
    var disks []*disk.Disk
    for i := 0; i < 5; i++ {
        d, err := disk.NewDisk(fmt.Sprintf("data/disk%d.dat", i), blockSize)
        if err != nil {
            panic(err)
        }
        disks = append(disks, d)
    }

    // RAID setups
    raid0, _ := raid.NewRAID0(disks)
    raid1, _ := raid.NewRAID1(disks)
    raid4, _ := raid.NewRAID4(disks)
    raid5, _ := raid.NewRAID5(disks)

    fmt.Println("Benchmarking RAID 0")
    benchmark.RunBenchmark(raid0, blockSize, totalSize)

    fmt.Println("Benchmarking RAID 1")
    benchmark.RunBenchmark(raid1, blockSize, totalSize)

    fmt.Println("Benchmarking RAID 4")
    benchmark.RunBenchmark(raid4, blockSize, totalSize)

    fmt.Println("Benchmarking RAID 5")
    benchmark.RunBenchmark(raid5, blockSize, totalSize)
}
