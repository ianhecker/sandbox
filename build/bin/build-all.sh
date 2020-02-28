#!/bin/bash
PLATFORMS=("linux_amd64" "windows_amd64")

BINARY_NAME=$1
VERSION=$2
TARGET=$3
ERROR_FMT="Error in build script: %s. Exiting\n"

if [[ -z "$BINARY_NAME" ]]; then
	printf "$ERROR_FMT" "Binary name not set"
	exit 1
fi
if [[ -z "$VERSION" ]]; then
	printf "$ERROR_FMT" "Version not set"
	exit 1
fi
if [[ -z "$TARGET" ]]; then
	printf "$ERROR_FMT" "Target not set"
	exit 1
fi

for PLATFORM in "${PLATFORMS[@]}"
do
	split=(${PLATFORM//_/ })
	goos=${split[0]}
	goarch=${split[1]}
	
	env GOOS=$goos GOARCH=$goarch \
	go build \
	-o bin/$BINARY_NAME-$goos-$goarch-$VERSION \
	main.go
done