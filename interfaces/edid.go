package interfaces

type EDIDBlock interface {
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
}

type DisplayDescriptor interface {
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
}
