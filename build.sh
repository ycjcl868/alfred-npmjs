#!/bin/sh

set -o errexit

echo "Building go binary:"
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/alfred-npmjs" ./src

if [[ -z "${VERSION}" ]]; then
  echo "Build go binary completed"
else
  archive="alfred-npmjs-${VERSION}.alfredworkflow"
  echo ""
  echo "Crearing archive:"
  (
      envsubst >.workflow/info.plist <./info.plist.template
      cd ./.workflow || exit
      zip -r "../${archive}" ./*
  )

  echo ""
  echo "Build completed: \"${archive}\""
fi
