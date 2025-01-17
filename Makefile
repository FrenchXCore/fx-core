#!/usr/bin/make -f

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo 'unknown')
GIT_TAGS := $(shell git describe --tags --always 2>/dev/null || echo 'unknown')
VERSION := $(GIT_BRANCH)-$(GIT_TAGS)
COMMIT := $(shell git log -1 --format='%H' 2>/dev/null || echo 'unknown')

LEDGER_ENABLED ?= true
BUILDDIR ?= $(CURDIR)/build
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)
GOPATH ?= '$(HOME)/go'
STATIK = $(GOPATH)/bin/statik

export GO111MODULE = on

ifeq ($(OS),Windows_NT)
  BINARYNAME := fxcored.exe
else
  BINARYNAME := fxcored
endif

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

#ifeq (badgerdb,$(findstring badgerdb,$(FX_BUILD_OPTIONS)))
#  build_tags += badgerdb
#endif

#ifeq (cleveldb,$(findstring cleveldb,$(FX_BUILD_OPTIONS)))
#  build_tags += gcc cleveldb muslc
#endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

comma:= ,
empty:=
space:= $(empty) $(empty)
build_tags_comma_sep := $(subst $(space),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep) \
		  -X github.com/cosmos/cosmos-sdk/version.Name=fxcore \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=fxcored

ifeq (,$(findstring nostrip,$(FX_BUILD_OPTIONS)))
  ldflags += -w -s
endif

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags '$(build_tags)' -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(FX_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# Check for debug option
ifeq (debug,$(findstring debug,$(FX_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif

###############################################################################
###                                  Build                                  ###
###############################################################################

all: build lint test

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify
	go mod tidy
	@echo "--> Download go modules to local cache"
	go mod download

build: go.sum
	go build -mod=readonly -v $(BUILD_FLAGS) -o $(BUILDDIR)/bin/$(BINARYNAME) ./cmd/fxcored
	@echo "--> Done building."

build-win:
	@$(MAKE) build

build-linux:
	@GOOS=linux GOARCH=amd64 $(MAKE) build

INSTALL_DIR := $(shell go env GOPATH)/bin
install: build $(INSTALL_DIR)
	mv $(BUILDDIR)/bin/fxcored $(shell go env GOPATH)/bin/fxcored
	@echo "--> Run \"fxcored start\" or \"$(shell go env GOPATH)/bin/fxcored start\" to launch fxcored."

$(INSTALL_DIR):
	@echo "Folder $(INSTALL_DIR) does not exist"
	mkdir -p $@

docker:
	@echo "--> Building fxcore docker image"
	docker build --progress plain -t ghcr.io/functionx/fx-core:latest .

run-local: install
	@./local-node.sh init

.PHONY: build build-win install docker go.sum run-local

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	@echo "--> Running linter"
	@which golangci-lint > /dev/null || echo "\033[91m install golangci-lint ...\033[0m" && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run -v --go=1.19 --out-format=tab

format: format-goimports
	golangci-lint run --fix --out-format=tab --issues-exit-code=0

format-goimports:
	@go install github.com/incu6us/goimports-reviser/v3@latest
	@find . -name '*.go' -type f -not -path './build*' -not -name 'statik.go' -not -name '*.pb.go' -not -name '*.pb.gw.go' -exec goimports-reviser -use-cache -rm-unused {} \;

lint-shell:
	# install shellcheck > https://github.com/koalaman/shellcheck
	grep -r '^#!/usr/bin/env bash' --exclude-dir={node_modules,build}  . | cut -d: -f1 | xargs shellcheck

format-shell:
	# install shfmt > https://github.com/mvdan/sh
	go install mvdan.cc/sh/v3/cmd/shfmt@v3.6.0
	grep -r '^#!/usr/bin/env bash' --exclude-dir={node_modules,build}  . | cut -d: -f1 | xargs shfmt -l -w -i 2

.PHONY: format lint format-goimports lint-shell

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test:
	@echo "--> Running tests"
	go test -mod=readonly ./...

test-count:
	go test -mod=readonly -cpu 1 -count 1 -cover ./... | grep -v 'types\|cli\|no test files'

.PHONY: test

###############################################################################
###                                Protobuf                                 ###
###############################################################################
protoVer=0.11.2
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)
containerProtoGenSwagger=$(PROJECT_NAME)-proto-gen-swagger-$(protoVer)
containerProtoFmt=$(PROJECT_NAME)-proto-fmt-$(protoVer)
containerProtoFork=$(PROJECT_NAME)-proto-fork-$(protoVer)

proto-all:
	@$(MAKE) proto-format
	@$(MAKE) proto-gen

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --rm --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./contrib/protoc/format.sh; fi

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen:v0.7 \
		sh ./contrib/protoc/gen.sh; fi
	@go mod tidy

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./contrib/protoc/swagger-gen.sh; fi

proto-fork:
	@echo "Forking Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFork}$$"; then docker rm $(containerProtoFork); fi
	@docker run --rm --name $(containerProtoFork) -e BUF_NAME=${BUF_NAME} -e BUF_TOKEN=${BUF_TOKEN} -e BUF_ORG=${BUF_ORG} -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./contrib/protoc/fork.sh

.PHONY: proto-format proto-gen proto-swagger-gen proto-fork

statik: $(STATIK)
$(STATIK):
	@echo "Installing statik..."
	@go install github.com/rakyll/statik@latest

update-swagger-docs: proto-swagger-gen statik
	$(GOPATH)/bin/statik -src=docs/swagger-ui -dest=docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
	perl -pi -e "print \"host: fx-rest.functionx.io\nschemes:\n  - https\n\" if $$.==6 " ./docs/swagger-ui/swagger.yaml

.PHONY: statik update-swagger-docs

###############################################################################
###                               Contracts                                 ###
###############################################################################

contract-abigen:
	@./contract/compile.sh

contract-publish:
	@./solidity/release.sh

.PHONY: contract-abigen contract-publish

###############################################################################
###                                Releasing                                ###
###############################################################################

PACKAGE_NAME := github.com/functionx/fx-core/v5
GOLANG_CROSS_VERSION := v1.19
release-dry-run:
	docker run --rm --privileged -e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v ${GOPATH}/pkg:/go/pkg \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release --clean --skip-publish --snapshot

release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run --rm --privileged -e CGO_ENABLED=1 --env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release --clean --skip-validate --release-notes ./release-note.md

.PHONY: release-dry-run release