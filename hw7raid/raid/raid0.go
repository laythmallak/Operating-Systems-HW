
    "errors"
    "raid-sim/disk"
)

type RAID0 struct {
    disks []*disk.Disk
}

func NewRAID0(disks []*disk.Disk) (*RAID0, error) {
    if len(disks) < 2 {
        return nil, errors.New("RAID 0 needs at least 2 disks")
    }
    return &RAID0{disks: disks}, nil
}

func (r *RAID0) Write(blockNum int, data []byte) error {
    diskIndex := blockNum % len(r.disks)
    return r.disks[diskIndex].WriteBlock(blockNum/len(r.disks), data)
}

func (r *RAID0) Read(blockNum int) ([]byte, error) {
    diskIndex := blockNum % len(r.disks)
    return r.disks[diskIndex].ReadBlock(blockNum/len(r.disks))
}