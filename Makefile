BINARY = time-manage
GOOS = linux
GOARCH = amd64
CGO_ENABLED = 0
VERSION = 1.0.0

build:
	@export GOOS=${GOOS}; \
  	export GOARCH=${GOARCH}; \
			export CGO_ENABLED=${CGO_ENABLED}; \
	go build -o ${BINARY}-${GOOS}-${GOARCH}-${VERSION} -ldflags="-s -w" ./main.go
