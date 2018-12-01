BIN_DIR := $(GOPATH)/bin

GOLINT_CMD := ${BIN_DIR}/golint
GOCYCLO_CMD := ${BIN_DIR}/gocyclo
DEADCODE_CMD := ${BIN_DIR}/deadcode
GOLINT := $(shell command -v ${GOLINT_CMD} 2> /dev/null)
GOCYCLO := $(shell command -v ${GOCYCLO_CMD} 2> /dev/null)
DEADCODE := $(shell command -v ${DEADCODE_CMD} 2> /dev/null)

OUTPUT := _output

#LDFLAGS = -X github.com/innoxchain/ixstorage/cmd/ixclient.Version=`git rev-parse --short HEAD` \
		  -X github.com/innoxchain/ixstorage/cmd/ixclient.BuildTime=`date +%Y-%m-%d`

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
	@go build -o ${OUTPUT}/ixclient ./cmd/ixclient
	#@go build -ldflags ${LDFLAGS} -o ${OUTPUT}/ixclient ./cmd/ixclient

install: build
	@echo "Installing ixclient binary to '$(GOPATH)/bin/ixclient'"
	@mkdir -p $(GOPATH)/bin && cp $(PWD)/_output/ixclient $(GOPATH)/bin/ixclient
	@echo "Installation successful."

.PHONY: clean
clean:
	@rm $(GOPATH)/bin/ixclient
	@rm -rf ${OUTPUT}