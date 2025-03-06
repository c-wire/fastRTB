#!/bin/bash

# Generate Go code from protobuf definitions
protoc -I=. \
       --go_out=go \
       --go_opt=paths=source_relative \
       --go-grpc_out=go \
       --go-grpc_opt=paths=source_relative \
       fastrtb/*.proto fastrtb/formats/*.proto