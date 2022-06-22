#!/usr/bin/env bash

# generate huge file first
if [ -e dist/generate.bin ]
then
  ./dist/generate.bin
else
  go build -o "dist/generate.bin" ./cmd/generate_main.go
  ./dist/generate.bin
fi

go test -v -coverpkg=./... -coverprofile=coverage.out
cat coverage.out | grep -v "tools/" > tmp.out
cat tmp.out | grep -v "main" > cover.out
go tool cover -func=cover.out

