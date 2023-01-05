package proto

// 生成 proto
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination/proto/ -I=../../../../ -I=../../../../comment/service/v1/proto/ -I=../../../../content/service/v1/proto/ -I=../../../../user/service/v1/proto/ -I=../../../../file/service/v1/proto/ --go_out=paths=source_relative:../ ./*.proto

// 生成 proto http
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination/proto/ -I=../../../../ -I=../../../../comment/service/v1/proto/ -I=../../../../content/service/v1/proto/ -I=../../../../user/service/v1/proto/ -I=../../../../file/service/v1/proto/ --go-http_out=paths=source_relative:../ ./*.proto

// 生成 proto errors
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination/proto/ -I=../../../../ -I=../../../../comment/service/v1/proto/ -I=../../../../content/service/v1/proto/ -I=../../../../user/service/v1/proto/ -I=../../../../file/service/v1/proto/ --go-errors_out=paths=source_relative:../ ./*.proto

// 生成 openapi doc
//go:generate protoc -I=. -I=../../../../../third_party -I=../../../../pagination/proto/ -I=../../../../ -I=../../../../comment/service/v1/proto/ -I=../../../../content/service/v1/proto/ -I=../../../../user/service/v1/proto/ -I=../../../../file/service/v1/proto/ --openapiv2_out ../ --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true ./*.proto
