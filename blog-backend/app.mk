GOPATH:=$(shell go env GOPATH)

KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

APP_VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "go-kratos/" $$0 ":0.1.0"}')

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

.PHONY: init
# initialize develop environment
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/google/gnostic@latest
	#@go install github.com/google/wire/cmd/wire@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: dep
# download dependencies of module
dep:
	@go mod download

.PHONY: vendor
# create vendor
vendor:
	@go mod vendor

.PHONY: build
# build application
build:
	mkdir -p bin/ && go build -ldflags "-X main.Service.Version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
.PHONY: clean
clean:
	@go clean
	rm --force "coverage.out"

.PHONY: docker
# build docker image
docker:
	cd ../../.. && @docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(APP_DOCKER_IMAGE) .

.PHONY: conf
# generate config struct code
conf:
	@go generate ./internal/conf/...

.PHONY: ent
# generate ent code
ent:
	@go generate ./internal/data/ent/...

.PHONY: wire
# generate wire code
wire:
	@go run github.com/google/wire/cmd/wire ./cmd/server

.PHONY: api
# generate api code
api:
	$(if $(IS_WINDOWS), cd ..\..\..\api\ & .\generate.bat, cd ../../../api/ & ./generate.sh)

.PHONY: run
# run application
run:
	go run ./cmd/server -conf ./configs

.PHONY: test
# run tests
test:
	@go test ./...

.PHONY: cover
# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

.PHONY: vet
# run static analysis
vet:
	@go vet

.PHONY: lint
# run lint
lint:
	@golangci-lint run

.PHONY: all
# build all
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
