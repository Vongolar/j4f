package gate

type serializeMode = uint8

const (
	_ serializeMode = iota
	serJSON
	serProtobuf
)
