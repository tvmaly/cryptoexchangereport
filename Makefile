BIN=$(GOPATH)/bin
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_INSTALL=$(GO_CMD) install -v
GO_CLEAN=$(GO_CMD) clean
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_CMD) get -d -v -u
GO_VET=$(GO_CMD) vet
GO_FMT=$(GO_CMD) fmt
GO_GET=$(GO_CMD) get
DEP_CMD=$(BIN)/dep


.PHONY: all

GOFLAGS ?= $(GOFLAGS:)

all: deps web scripts

deps: $(DEP_CMD)

$(DEP_CMD):
	$(GO_GET) -u github.com/golang/dep/cmd/dep

web:server/website.go
	$(GO_BUILD) server/website.go

scripts: deps
	$(GO_BUILD) scripts/compile_scss.go
