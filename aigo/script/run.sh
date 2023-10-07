#!/bin/bash
# shellcheck disable=SC2164
cd "$1" || exit
echo "Run in $1"
go build;
./aigo;
rm -rf aigo;

cd "$2" || exit
rm -rf go.mod
go mod init "$3"
echo "go mod tidy"
go mod tidy
echo "go fmt"
go fmt ./...