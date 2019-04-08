APP?=goserialread
PROJECT?=github.com/enzo1920/go_dummy_serial

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')


GOOS?=linux
GOARCH?=amd64

GOOSW32?=windows
GOARCHW32?=386
APPVER32?=32

GOOSW64?=windows
GOARCHW64?=amd64
APPVER64?=64



clean:
			rm -f ${APP}
build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}
		
		
buildwin32: clean
	CGO_ENABLED=0 GOOS=${GOOSW32} GOARCH=${GOARCHW32} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}${APPVER32}.exe
		
buildwin64: clean
	CGO_ENABLED=0 GOOS=${GOOSW64} GOARCH=${GOARCHW64} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}${APPVER64}.exe
run: build
			./${APP}
test:
			go test -v -race ./...