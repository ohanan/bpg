#!/usr/bin/env bash
work_dir=$(pwd)
proto_exec=$work_dir/protoc
if [ ! -f "$proto_exec" ]; then
  proto_file=$work_dir/protoc.zip
  curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.2/protoc-21.2-osx-universal_binary.zip --output "$proto_file" --silent
  unzip -q "$proto_file" -d "$work_dir"/tmp
  mv "$work_dir"/tmp/bin/protoc "$proto_exec"
  rm -rf "$proto_file" "$work_dir"/tmp
fi
if [ "$(protoc-gen-go --version 2>&1)" != "protoc-gen-go v1.28.0" ]; then
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
fi
if [ "$(protoc-gen-go-grpc --version 2>&1)" != "protoc-gen-go v1.28.0" ]; then
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
fi
protoc --go_out=../pkg/ --go_opt=paths=source_relative --go-grpc_out=../pkg/ \
  --go-grpc_opt=paths=source_relative --proto_path=../proto bgp.proto game.proto bot.proto
