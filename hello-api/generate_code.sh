#!/bin/sh

protoc \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative hello-proto/hello.proto \

# 引用:https://www.youtube.com/watch?v=YudT0nHvkkE&t=1468s&ab_channel=TechandBeyondWithMoss

# 以下のdirの状態でsh実行

# 注意!
# go mod init github.com/sk-develop/grpc-sampleを行う際は、
# pushするgithub repositoryに合わせる。
# そして、hello.protoも
# option go_package = "github.com/sk-develop/grpc-sample;grpc_sample";
# と合わせる
# serverを作って、pbファイルを読み込む際も
# pb "github.com/sk-develop/grpc-sample/hello-proto"
# と合わせて、一応~/go/pkg/mod/配下(Module modeでのコンパイラのimport pacakageの参照先)を確認する
# 「~/go/pkg/mod/github.com/sk-develop/grpc-sample@v0.0.0-202109241204」

# sh実行直前
# .
# ├── generate_code.sh
# ├── go.mod
# ├── go.sum
# ├── hello-proto
# │  └── hello.proto

# sh実行後
# .
# ├── generate_code.sh
# ├── go.mod
# ├── go.sum
# ├── hello-proto
# │  ├── hello.pb.go
# │  ├── hello.proto
# │  └── hello_grpc.pb.go

# 目的達成後、最終的に以下となる
# .
# ├── generate_code.sh
# ├── go.mod
# ├── go.sum
# ├── hello-proto
# │  ├── hello.pb.go
# │  ├── hello.proto
# │  └── hello_grpc.pb.go
# └── server
#    └── hello.go