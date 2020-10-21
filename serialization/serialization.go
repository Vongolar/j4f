package serialization

//SerializateType 序列化方式
type SerializateType int

const (
	_ SerializateType = iota
	//JSON json
	JSON
	//Protobuf google protobuf
	Protobuf
)
