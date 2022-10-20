GOPATH:=$(shell go env GOPATH)

KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

APP_VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "go-kratos/" $$0 ":0.1.0"}')

# 初始化环境
.PHONY: init
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/google/wire/cmd/wire@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/google/gnostic@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 下载依赖库
.PHONY: dep
dep:
	@go mod download

# 创建依赖库
.PHONY: vendor
vendor:
	@go mod vendor

# 构建生成二进制可执行文件
.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Service.Version=$(APP_VERSION)" -o ./bin/ ./...

# 清理中间目标文件
.PHONY: clean
clean:
	@go clean
	rm --force "coverage.out"

# 直接运行程序
.PHONY: run
run:
	go run ./cmd/server -conf ./configs

# 构建docker镜像
.PHONY: docker
docker:
	cd ../../.. && @docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(APP_DOCKER_IMAGE) .

# 执行单元测试
.PHONY: test
test:
	@go test ./...

# 执行覆盖率测试
.PHONY: cover
cover:
	@go test -v ./... -coverprofile=coverage.out

# 执行代码静态检查
.PHONY: vet
vet:
	@go vet

# 生成Protobuf配置代码
.PHONY: conf
conf:
	@go generate ./internal/conf/...

# 生成Entgo代码
.PHONY: ent
ent:
	@go generate ./internal/data/ent/...

# 生成Wire代码
.PHONY: wire
wire:
	@wire ./cmd/server

# 生成所有的Protobuf API代码
.PHONY: api
api:
	@go generate ../../../api/...

# 执行代码检查
.PHONY: lint
lint:
	@golangci-lint run

# 执行所有
.PHONY: all
all: api wire conf build test

# 显示帮助信息
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make init: init environment'
	@echo "  make dep: download dependency libs"
	@echo "  make run: run the program"
	@echo "  make build: build binary program"
	@echo "  make docker: build docker images"
	@echo "  make clean: clean object"
	@echo "  make test: run the unit test"
	@echo "  make cover: run the coverage check"
	@echo "  make lint: run the lint check"
	@echo "  make vet: run the static analysis check"
	@echo "  make api: generate Protobuf API code"
	@echo "  make wire: generate Wire code"
	@echo "  make ent: generate Entgo code"
	@echo "  make conf: generate Protobuf config code"

.DEFAULT_GOAL := help
