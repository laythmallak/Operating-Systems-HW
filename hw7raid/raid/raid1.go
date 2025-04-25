
    "errors"
    "raid-sim/disk"
)

type RAID1 struct {
    disks []*disk.Disk
}

func NewRAID1(disks []*disk.Disk) (*RAID1, error) {
    if len(disks) < 2 {
        return nil, errors.New("RAID 1 needs at least 2 disks")
    }
    return &RAID1{disks: disks[:2]}, nil
}

func (r *RAID1) Write(blockNum int, data []byte) error {
    for _, d := range r.disks {
        if err := d.WriteBlock(blockNum, data); err != nil {
            return err
        }
    }
    return nil
}

func (r *RAID1) Read(blockNum int) ([]byte, error) {
    return r.disks[0].ReadBlock(blockNum)
}