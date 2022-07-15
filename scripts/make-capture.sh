#!/usr/bin/env bash
p=$(dirname $0)
cd ../cmd/games/capture-flag || exit 1
go build -o ../../../bin/capture-flag
echo builded capture-flag
cd example || exit 1
go build -o ../../../../bin/capture-flag-example
echo builded capture-flag-example
