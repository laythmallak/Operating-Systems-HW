
    "errors"
    "raid-sim/disk"
    "bytes"
)

type RAID4 struct {
    disks []*disk.Disk
}

func NewRAID4(disks []*disk.Disk) (*RAID4, error) {
    if len(disks) < 3 {
        return nil, errors.New("RAID 4 needs at least 3 disks")
    }
    return &RAID4{disks: disks}, nil
}

func (r *RAID4) Write(blockNum int, data []byte) error {
    dataDiskCount := len(r.disks) - 1
    diskIndex := blockNum % dataDiskCount
    blockIndex := blockNum / dataDiskCount

    parityDisk := r.disks[len(r.disks)-1]
    if err := r.disks[diskIndex].WriteBlock(blockIndex, data); err != nil {
        return err
    }

    parity := make([]byte, len(data))
    for i := 0; i < dataDiskCount; i++ {
        d, _ := r.disks[i].ReadBlock(blockIndex)
        for j := range d {
            parity[j] ^= d[j]
        }
    }
    return parityDisk.WriteBlock(blockIndex, parity)
}

func (r *RAID4) Read(blockNum int) ([]byte, error) {
    dataDiskCount := len(r.disks) - 1
    diskIndex := blockNum % dataDiskCount
    blockIndex := blockNum / dataDiskCount
    return r.disks[diskIndex].ReadBlock(blockIndex)
}