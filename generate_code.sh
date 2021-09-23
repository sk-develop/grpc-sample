#!/bin/sh

# 自動生成する元となるprotoファイルはproto/helloworld.protoです
# ⭐注意⭐：事前にdir作成しておかないとエラーでます
protoc proto/helloworld.proto \
       --go-grpc_out=api/helloworld \
       --go_out=api/helloworld

# 実行前のdirectory
# .
# ├── api
# │  └── helloworld
# ├── generate_code.sh
# └── proto
#    └── helloworld.proto