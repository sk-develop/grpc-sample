# 最も簡単な gPPC サーバコンテナ <!-- omit in toc -->

## 目次 <!-- omit in toc -->

- [目的:gRPC サーバー構築方法, 手順のマスター](#目的grpc-サーバー構築方法-手順のマスター)
- [gRPC サーバー構築参考記事](#grpc-サーバー構築参考記事)
  - [開発時の注意点](#開発時の注意点)
    - [設定や path をソースコードを置く repository に合わせる](#設定や-path-をソースコードを置く-repository-に合わせる)
    - [dir 構成](#dir-構成)
- [必要なコマンド](#必要なコマンド)
  - [確認コマンド](#確認コマンド)

## 目的:gRPC サーバー構築方法, 手順のマスター

## gRPC サーバー構築参考記事

- [gRPC Introduction Using Golang: Build a Unary Service Method - Part 1](https://youtu.be/YudT0nHvkkE)
  - [Git](https://github.com/tech-with-moss/go-usermgmt-grpc/blob/main/usermgmt_server/usermgmt_server.go)
- [「 gRPC Web 」で gRPC 実践！ Go と gRPC で WebAPI を作ってみよう！！](https://youtu.be/hlyNZoaXvqU)
  - [Git](https://github.com/yassun-youtube/grpc-web-sample/blob/master/api/server.go)

### 開発時の注意点

#### 設定や path をソースコードを置く repository に合わせる

```txt
go mod init github.com/sk-develop/grpc-sample/hello-api を行う際は、
push する github repository に合わせる。
そして、hello.proto も
option go_package = "github.com/sk-develop/grpc-sample/hello-api;hello_api";
と合わせる
server を作って、pb ファイルを読み込む際も
pb "github.com/sk-develop/grpc-sample/hello-api/hello-proto"
と合わせて、一応~/go/pkg/mod/配下(Module mode でのコンパイラの import pacakage の参照先)を確認する
```

#### dir 構成

```txt
sh 実行直前
.
├── generate_code.sh
└── hello-api
├── go.mod
├── go.sum
└── hello-proto
└──hello.proto

sh 実行後
.
├── generate_code.sh
└── hello-api
├── go.mod
├── go.sum
├── hello-proto
├── hello.pb.go
├── hello.proto
└── hello_grpc.pb.go

目的達成後、最終的に以下となる
.
├── Dockerfile
├── generate_code.sh
└── hello-api
   ├── go.mod
   ├── go.sum
   ├── hello-proto
   │  ├── hello.pb.go
   │  ├── hello.proto
   │  └── hello_grpc.pb.go
   └── main.go
```

## 必要なコマンド

### 確認コマンド

```sh
  grpc_cli call localhost:9090 hello.Greeter.SayHello 'name: "sk"'
  grpc_cli ls localhost:9090 hello.Greeter
  grpc_cli ls localhost:9090 hello.Greeter -l
```
