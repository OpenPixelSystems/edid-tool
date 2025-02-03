package interfaces

type BlockType int
type DescriptorType byte
type DataBlockType byte

type EDIDBlock interface {
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
	Type() BlockType
}

type DisplayDescriptor interface {
	Type() DescriptorType
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
}

type DataBlock interface {
	Type() DataBlockType
	Len() int
	ToBytes() []byte
	FromBytes([]byte) (int, error)
	Print()
}
