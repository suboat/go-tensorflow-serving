#!/bin/bash

# base
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/util/*.proto
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/config/*.proto
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/sources/storage_path/*.proto
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/core/*.proto
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/example/*.proto

# apis
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/apis/*.proto
protoc -I. -I ${GOPATH}/src/github.com/tensorflow/tensorflow --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tensorflow_serving/apis/internal/*.proto

# format
go fmt ./...
