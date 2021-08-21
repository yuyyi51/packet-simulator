package simple_packet

type DataEntry struct {
	offset int64
	size   int64
}

type ReceiveDataSorter struct {
	sort       map[int64]*DataEntry
	readOffset int64
}

func NewReceiveDataSorter() *ReceiveDataSorter {
	sort := &ReceiveDataSorter{
		sort: make(map[int64]*DataEntry),
	}
	return sort
}

func (sort *ReceiveDataSorter) onReceiveData(offset, size int64) {
	entry := &DataEntry{
		offset: offset,
		size:   size,
	}
	sort.sort[offset] = entry
}

func (sort *ReceiveDataSorter) readData() *DataEntry {
	data, ok := sort.sort[sort.readOffset]
	if !ok {
		return nil
	}
	delete(sort.sort, sort.readOffset)
	sort.readOffset = data.offset + data.size
	return data
}

func (sort *ReceiveDataSorter) hasData() bool {
	_, ok := sort.sort[sort.readOffset]
	return ok
}
