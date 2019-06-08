GIT_HASH=`git rev-parse HEAD`
GIT_TAG=`git tag --points-at HEAD`
OUTPUT=bin
BINARY_LINUX=github-releaser-travis
BINARY_WIN=github-releaser-travis.exe
BINARY_MAC=github-releaser-travis-mac
BUILD_FLAGS=-ldflags="-s -w -X main.buildTag=$(GIT_TAG)"

build-all: build-win build-linux build-mac

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o ./$(OUTPUT)/$(BINARY_WIN) ./cmd/github-release-cli/travis.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o ./$(OUTPUT)/$(BINARY_LINUX) ./cmd/github-release-cli/travis.go

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o ./$(OUTPUT)/$(BINARY_MAC) ./cmd/github-release-cli/travis.go

test:
	go test -v ./...
