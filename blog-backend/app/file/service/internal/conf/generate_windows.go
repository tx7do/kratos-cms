//go:build windows

//go:generate protoc --proto_path=. --proto_path=../../../../../third_party --go_out=paths=source_relative:. ./*.proto

package conf
