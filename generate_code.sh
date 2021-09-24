#!/bin/sh

protoc \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative hello-proto/hello.proto \

# 引用:https://www.youtube.com/watch?v=YudT0nHvkkE&t=1468s&ab_channel=TechandBeyondWithMoss