PROJECT := rename4plex
BUILD_TIME ?= $(shell date +%d.%m.%Y-%H%M)
GIT_COMMIT_HASH := $(shell git rev-parse --short HEAD)
VERSION ?= $(GIT_COMMIT_HASH)@$(BUILD_TIME)
BUILD_NAME ?= $(PROJECT)-$(BUILD_TIME)

GO_MAIN := cmd/rename4plex/main.go

CGO_0_BUILD := CGO_ENABLED=0 go build

RELEASE_LDFLAGS := "-X main.version=$(VERSION) -s -w"
RELEASE_BUILD := $(CGO_0_BUILD) -ldflags $(RELEASE_LDFLAGS) -v

DEV_LDFLAGS := "-X main.version=dev.$(VERSION) -X main.configPath=./tmp/config.yml"
DEV_BUILD := $(CGO_0_BUILD) -ldflags $(DEV_LDFLAGS) -v

.PHONY: clean build build-dev build-all install uninstall purge test release

clean:
	rm -rf _build/ release/

build:
	$(RELEASE_BUILD) -o $(BUILD_NAME) $(GO_MAIN)

build-dev:
	$(DEV_BUILD) -o $(PROJECT) $(GO_MAIN)

build-all:
	- make clean
	mkdir _build
	GOOS=linux   GOARCH=amd64 $(RELEASE_BUILD) -o _build/$(PROJECT)-linux-amd64 $(GO_MAIN)
	GOOS=linux   GOARCH=arm   $(RELEASE_BUILD) -o _build/$(PROJECT)-linux-arm $(GO_MAIN)
	GOOS=linux   GOARCH=arm64 $(RELEASE_BUILD) -o _build/$(PROJECT)-linux-arm64 $(GO_MAIN)
	GOOS=windows GOARCH=amd64 $(RELEASE_BUILD) -o _build/$(PROJECT)-windows-amd64.exe $(GO_MAIN)
	cd _build; sha256sum * > sha256sums.txt

install: build
ifeq ($(shell id -u), 0)
	mv ./$(BUILD_NAME) /usr/local/bin/$(PROJECT)
else
	-mkdir $(HOME)/bin
	chmod 700 ./ $(BUILD_NAME)
	mv ./$(BUILD_NAME) $(HOME)/bin/$(PROJECT)
endif
	$(PROJECT) init

uninstall:
ifeq ($(shell id -u), 0)
	rm -f /usr/local/bin/$(PROJECT)
else
	rm -f $(HOME)/bin/$(PROJECT)
endif

purge: uninstall
	rm -rf $(HOME)/.config/$(PROJECT)

test:
	go test -v -cover ./...

release: clean build-all
	mkdir release
	cp _build/* release
	cd release; sha256sum --quiet --check sha256sums.txt