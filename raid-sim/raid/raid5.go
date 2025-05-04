package raid

import (
    "errors"
    "raid-sim/disk"
)

type RAID5 struct {
    disks []*disk.Disk
}

func NewRAID5(disks []*disk.Disk) (*RAID5, error) {
    if len(disks) < 3 {
        return nil, errors.New("RAID 5 needs at least 3 disks")
    }
    return &RAID5{disks: disks}, nil
}

func (r *RAID5) Write(blockNum int, data []byte) error {
    n := len(r.disks)
    stripe := blockNum / (n - 1)
    offset := blockNum % (n - 1)
    parityDisk := stripe % n

    dataDiskIndex := func(i int) int {
        if i >= parityDisk {
            return i + 1
        }
        return i
    }(offset)

    if err := r.disks[dataDiskIndex].WriteBlock(stripe, data); err != nil {
        return err
    }

    parity := make([]byte, len(data))
    for i := 0; i < n-1; i++ {
        idx := i
        if idx >= parityDisk {
            idx++
        }
        d, _ := r.disks[idx].ReadBlock(stripe)
        for j := range d {
            parity[j] ^= d[j]
        }
    }
    return r.disks[parityDisk].WriteBlock(stripe, parity)
}

func (r *RAID5) Read(blockNum int) ([]byte, error) {
    n := len(r.disks)
    stripe := blockNum / (n - 1)
    offset := blockNum % (n - 1)
    parityDisk := stripe % n

    dataDiskIndex := func(i int) int {
        if i >= parityDisk {
            return i + 1
        }
        return i
    }(offset)

    return r.disks[dataDiskIndex].ReadBlock(stripe)
}
