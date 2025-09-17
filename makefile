APP := colo
PKG := .
DIST := dist

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
LDFLAGS := -s -w -X 'main.version=$(VERSION)' -X 'main.commit=$(COMMIT)'

.PHONY: all build clean

all: build

build:
	@echo "[build] linux/amd64 → $(DIST)/$(APP)"
	@mkdir -p $(DIST)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o $(DIST)/$(APP) $(PKG)

clean:
	@rm -rf $(DIST)
