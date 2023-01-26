.PHONY: help wire conf ent build api all

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

CURRENT_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

SRCS_MK := $(foreach dir, app, $(wildcard $(dir)/*/*/Makefile))

wire:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make wire;\
    )

conf:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make conf;\
    )

ent:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make ent;\
    )

build:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make build;\
    )

api:
	buf generate && \
	$(if $(IS_WINDOWS), rmdir /s/q .\gen\api\go\third_party, rm -fr ./gen/api/go/third_party)

all:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make app;\
    )

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
