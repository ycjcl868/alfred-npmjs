#!/bin/sh

set -o errexit
set -o nounset

archive="alfred-npmjs-${VERSION}.alfredworkflow"

echo "Building go binary:"
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/alfred-npmjs" .

echo ""
echo "Crearing archive:"
(
    envsubst >.workflow/info.plist <./info.plist.template
    cd ./.workflow || exit
    zip -r "../${archive}" ./*
)

echo ""
echo "Build completed: \"${archive}\""
