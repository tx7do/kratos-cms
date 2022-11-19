GOPATH:=$(shell go env GOPATH)

KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

APP_VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "go-kratos/" $$0 ":0.1.0"}')

.PHONY: init
# 初始化环境
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/google/wire/cmd/wire@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/google/gnostic@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: dep
# 下载依赖库
dep:
	@go mod download

.PHONY: vendor
# 创建依赖库
vendor:
	@go mod vendor

.PHONY: build
# 构建生成二进制可执行文件
build:
	mkdir -p bin/ && go build -ldflags "-X main.Service.Version=$(APP_VERSION)" -o ./bin/ ./...

# 清理中间目标文件
.PHONY: clean
clean:
	@go clean
	rm --force "coverage.out"

.PHONY: run
# 直接运行程序
run:
	go run ./cmd/server -conf ./configs

.PHONY: docker
# 构建docker镜像
docker:
	cd ../../.. && @docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(APP_DOCKER_IMAGE) .

.PHONY: test
# 执行单元测试
test:
	@go test ./...

.PHONY: cover
# 执行覆盖率测试
cover:
	@go test -v ./... -coverprofile=coverage.out

.PHONY: vet
# 执行代码静态检查
vet:
	@go vet

.PHONY: conf
# 生成Protobuf配置代码
conf:
	@go generate ./internal/conf/...

.PHONY: ent
# 生成Entgo代码
ent:
	@go generate ./internal/data/ent/...

.PHONY: wire
# 生成Wire代码
wire:
	@wire ./cmd/server

.PHONY: api
# 生成所有的Protobuf API代码
api:
	@go generate ../../../api/...

.PHONY: lint
# 执行代码检查
lint:
	@golangci-lint run

.PHONY: all
# 执行所有
all: api wire conf build test

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
