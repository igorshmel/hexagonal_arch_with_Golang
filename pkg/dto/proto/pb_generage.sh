#!/bin/bash
set -e

currentDir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)

mkdir -p "$currentDir/../pb"
cd "$currentDir/../pb"
rm -rf ./*

cd "$currentDir"

protoc \
    --go_out="$currentDir/../pb" --go_opt=paths=source_relative \
    --go-grpc_out="$currentDir/../pb" --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
		file_producer.proto user.proto schema.proto