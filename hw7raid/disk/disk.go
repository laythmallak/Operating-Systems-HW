
    "os"
)

type Disk struct {
    file      *os.File
    blockSize int
}

func NewDisk(path string, blockSize int) (*Disk, error) {
    os.MkdirAll("data", 0755)
    f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        return nil, err
    }
    return &Disk{file: f, blockSize: blockSize}, nil
}

func (d *Disk) WriteBlock(blockNum int, data []byte) error {
    offset := int64(blockNum) * int64(d.blockSize)
    _, err := d.file.WriteAt(data, offset)
    d.file.Sync()
    return err
}

func (d *Disk) ReadBlock(blockNum int) ([]byte, error) {
    offset := int64(blockNum) * int64(d.blockSize)
    buf := make([]byte, d.blockSize)
    _, err := d.file.ReadAt(buf, offset)
    return buf, err
}