package serialize

type SerializeMode = uint8

const (
	_ SerializeMode = iota
	JSON
	Protobuf
)

func Invaild(mode int) bool {
	return mode >= int(JSON) && mode <= int(Protobuf)
}
