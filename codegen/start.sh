#!/bin/bash
# shellcheck disable=SC2164
cd /Users/hb.li/Documents/royce/star-royce/echo-shopping/scripts/codegen;
go build;
./codegen;
rm -rf codegen;