#!/bin/sh

# 自動生成する元となるprotoファイルはproto/helloworld.protoです
# ⭐注意⭐：事前にdir作成しておかないとエラーでます
protoc proto//helloworld/helloworld.proto \
       --go-grpc_out=api/helloworld \
       --go_out=api/helloworld \
       # --go_opt=module=$PREFIXというフラグ
       # $PREFIXに書いた部分のディレクトリが省略され, 期待する位置にファイルが生成される
       --go_opt=module=github.com/sk-develop/grpc-sample/proto/helloworld
       # --go_opt=module=github.com/jun06t/grpc-sample/go-package-option/after \
       #                github.com/sk-develop/grpc-sample/proto/helloworld

# -Iオプションは.protoファイルでimportするファイルのPATHです。 --proto_pathの短縮形です。
# https://christina04.hatenablog.com/entry/protoc-usage

# --go_opt=module=$PREFIXというフラグ
# .
# ├── api
# │  └── helloworld
# │     └── github.com
# │        └── sk-develop
# │           └── grpc-sample
# │              └── api
# │                 └── helloworld
# │                    ├── helloworld.pb.go
# │                    └── helloworld_grpc.pb.go

# 実行前のdirectory
# .
# ├── api
# │  └── helloworld
# ├── generate_code.sh
# └── proto
#    └── helloworld
#       └── helloworld.proto