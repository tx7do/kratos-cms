//go:build !windows
// +build !windows

// 生成 proto
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../third_party --go_out=paths=source_relative:../ ./*.proto"

// 生成 openapi doc
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../third_party --openapiv2_out ../ ./*.proto --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true"

package proto
