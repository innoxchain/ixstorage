ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

GOBIN := $(GOPATH)/bin

GOLINT_CMD := ${GOBIN}/golint
GOCYCLO_CMD := ${GOBIN}/gocyclo
DEADCODE_CMD := ${GOBIN}/deadcode
GOLINT := $(shell command -v ${GOLINT_CMD} 2> /dev/null)
GOCYCLO := $(shell command -v ${GOCYCLO_CMD} 2> /dev/null)
DEADCODE := $(shell command -v ${DEADCODE_CMD} 2> /dev/null)

GIT_VERSION=$(shell git describe --tags --abbrev=0)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_BUILD_TIME=$(shell date +%Y-%m-%d)

OUTPUT := _output

LDFLAGS = "-X github.com/innoxchain/ixstorage/build.Version=$(GIT_VERSION) \
-X github.com/innoxchain/ixstorage/build.Commit=$(GIT_COMMIT) \
-X github.com/innoxchain/ixstorage/build.Branch=$(GIT_BRANCH) \
-X github.com/innoxchain/ixstorage/build.BuildTime=$(GIT_BUILD_TIME)"

pkgs = ./build ./cmd/ixclient ./pkg/apps/ixclient

all: build

dependencies:
ifndef GOLINT
	@echo "Installing golint" && go get -u golang.org/x/lint/golint
endif
ifndef GOCYCLO
	@echo "Installing gocyclo" && go get -u github.com/fzipp/gocyclo
endif
ifndef DEADCODE
	@echo "Installing deadcode" && go get -u github.com/remyoudompheng/go-misc/deadcode
endif

verify: dependencies lint check_cyclo vet check_deadcode

.PHONY: lint
lint: dependencies
	@echo "Running $@"
	@${GOLINT_CMD} -min_confidence=1.0 -set_exit_status $(pkgs)

.PHONY: check_cyclo
check_cyclo: dependencies
	@echo "Running $@"
	@${GOCYCLO_CMD} -over 20 cmd
	@${GOCYCLO_CMD} -over 20 pkg

.PHONY: vet
vet:
	@echo "Running $@"
	@go tool vet cmd
	@go tool vet pkg

.PHONY: check_deadcode
check_deadcode: dependencies
	@echo "Running $@"
	@${DEADCODE_CMD} $(pkgs) || true

test: 
	@go test -v $(pkgs)

build: verify
	@echo "Building ixclient binary to './_output/ixclient'"
	@go build -ldflags ${LDFLAGS} -o ${OUTPUT}/ixclient ./cmd/ixclient

install: build
	@echo "Installing ixclient binary to '$(GOPATH)/_output/ixclient'"
	@mkdir -p $(GOPATH)/bin && cp $(PWD)/_output/ixclient $(GOPATH)/bin/ixclient
	@echo "Installation successful."

.PHONY: clean
clean:
	@rm $(GOPATH)/bin/ixclient
	@rm -rf ${OUTPUT}