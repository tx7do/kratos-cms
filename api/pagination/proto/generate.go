package proto

// 生成 proto
//go:generate protoc -I=. --go_out=paths=source_relative:../ ./*.proto

// 生成 openapi doc
//go:generate protoc -I=. --openapiv2_out ../ ./*.proto --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true
