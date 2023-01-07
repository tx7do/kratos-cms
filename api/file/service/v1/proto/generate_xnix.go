//go:build !windows
// +build !windows

// 生成 proto
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../pagination/proto/ --proto_path=../../../../ --go_out=paths=source_relative:../ ./*.proto"

// 生成 proto grpc
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../pagination/proto/ --proto_path=../../../../ --go-grpc_out=paths=source_relative:../ ./*.proto"

// 生成 proto http
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../pagination/proto/ --proto_path=../../../../ --go-http_out=paths=source_relative:../ ./*.proto"

// 生成 proto errors
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../pagination/proto/ --proto_path=../../../../ --go-errors_out=paths=source_relative:../ ./*.proto"

// 生成 openapi doc
//go:generate sh -c "protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../pagination/proto/ --proto_path=../../../../ --openapiv2_out ../ --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true ./*.proto"

package v1
