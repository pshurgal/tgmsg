#!/bin/bash -ex

goos=$1
goarch=$2

artifact_name="tgmsg-${goos}-${goarch}"
if [ "${goos}" == "windows" ]; then
    artifact_name="${artifact_name}.exe"
fi

GOOS=${goos} GOARCH=${goarch} go build -ldflags "-s -w" -o ${artifact_name}
