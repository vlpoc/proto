#!/usr/bin/env bash

set -x

protoc -I=. --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	authpb/auth.proto \
	exec/exec.proto

#	ds/ds.proto \