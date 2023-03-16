#!/bin/bash
# shellcheck disable=SC2164
go build;
./gogen-cli;
rm -rf gogen-cli;