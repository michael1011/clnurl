PKG := github.com/michael1011/clnurl

GO_BIN := ${GOPATH}/bin

GOTEST := CGO_ENABLED=1 GO111MODULE=on go test -v
GOBUILD := CGO_ENABLED=1 GO111MODULE=on go build -v
GOINSTALL := CGO_ENABLED=1 GO111MODULE=on go install -v
GOLIST := go list -deps $(PKG)/... | grep '$(PKG)'| grep -v '/vendor/'

COMMIT := $(shell git log --pretty=format:'%h' -n 1)
LDFLAGS := -ldflags "-X $(PKG)/build.Commit=$(COMMIT) -w -s"

LINT_PKG := github.com/golangci/golangci-lint/cmd/golangci-lint
LINT_BIN := $(GO_BIN)/golangci-lint
LINT = $(LINT_BIN) run -v --timeout 5m

GREEN := "\\033[0;32m"
NC := "\\033[0m"

define print
	echo $(GREEN)$1$(NC)
endef

default: build

#
# Dependencies
#

$(LINT_BIN):
	@$(call print, "Fetching linter")
	go install $(LINT_PKG)@latest

#
# Tests
#

unit:
	@$(call print, "Running unit tests")
	$(GOLIST) | $(XARGS) env $(GOTEST)

#
# Building
#

build-frontend:
	@$(call print, "Building frontend")
	npm run export
	rm -r router/out 2> /dev/null || echo > /dev/null
	cp -r out router/out

build-backend:
	@$(call print, "Building backend")
	$(GOBUILD) -o bin/clnurl $(LDFLAGS) $(PKG)/plugin

build-breez:
	@$(call print, "Building Breez")
	$(GOBUILD) -o bin/breez $(LDFLAGS) $(PKG)/breez-cli

build:
	$(MAKE) build-frontend
	$(MAKE) build-backend
	$(MAKE) build-breez

#
# Utils
#

fmt:
	@$(call print, "Formatting source")
	gofmt -l -s -w .

lint: $(LINT_BIN)
	@$(call print, "Linting source")
	$(LINT)
	npm run lint

.PHONY: build
