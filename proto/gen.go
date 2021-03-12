package proto

//go:generate protoc --proto_path=./ --go_out=../ command.proto errorCode.proto common/log.proto
