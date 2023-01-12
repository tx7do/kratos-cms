//go:build windows

// 生成 proto
//go:generate protoc --proto_path=. --proto_path=../../../third_party --go_out=paths=source_relative:../ ./*.proto

// 生成 openapi doc
//go:generate protoc --proto_path=. --proto_path=../../../third_party --openapiv2_out ../ ./*.proto --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true

package proto
