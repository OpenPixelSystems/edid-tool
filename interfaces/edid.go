package interfaces

type BlockType int

type EDIDBlock interface {
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
	Type() BlockType
}

type DisplayDescriptor interface {
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
}
