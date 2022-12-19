package v1

// 生成 proto
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination -I=../../../../ --go_out=paths=source_relative:../ ./*.proto

// 生成 proto grpc
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination -I=../../../../ --go-grpc_out=paths=source_relative:../ ./*.proto

// 生成 proto http
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination -I=../../../../ --go-http_out=paths=source_relative:../ ./*.proto

// 生成 proto errors
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination -I=../../../../ --go-errors_out=paths=source_relative:../ ./*.proto

// 生成 openapi doc
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination -I=../../../../ --openapiv2_out ../ --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true ./*.proto
