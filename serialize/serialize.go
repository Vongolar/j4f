package serialize

type SerializeMode = uint8

const (
	_ SerializeMode = iota
	JSON
	Protobuf
)

func VaildSerializeMode(mode int) bool {
	return mode >= int(JSON) && mode <= int(Protobuf)
}
