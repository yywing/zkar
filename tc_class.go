package javaserialize

type TCClass struct {
	ClassPointer *TCClassPointer
}

func (c *TCClass) ToBytes() []byte {
	var bs = []byte{JAVA_TC_CLASS}
	bs = append(bs, c.ClassPointer.ToBytes()...)
	return bs
}

func readTCClass(stream *ObjectStream) (*TCClass, error) {
	var classes []*TCClassDesc // current unused
	var class = new(TCClass)
	var err error

	_, _ = stream.ReadN(1)
	class.ClassPointer, err = readTCClassPointer(stream, classes)
	if err != nil {
		return nil, err
	}

	stream.AddReference(class)
	return class, nil
}