#!/usr/bin/env bash

mkdir -p dist/

go build -o "dist/main.bin" ./main.go
go build -o "dist/generate.bin" ./cmd/generate_main.go