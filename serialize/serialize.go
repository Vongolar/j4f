package serialize

type SerializeMode = uint8

const (
	_ SerializeMode = iota
	JSON
	Protobuf
)
