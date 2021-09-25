FROM golang:1.17.1-alpine3.14 AS build

ENV GO111MODULE=on
WORKDIR /go/src/github.com/sk-develop/grpc-sample/hello-api
COPY ./hello-api/go.mod ./hello-api/go.sum ./
RUN go mod download

COPY ./hello-api/. ./
RUN go build

FROM alpine:3.14
COPY --from=build /go/src/github.com/sk-develop/grpc-sample/hello-api/hello-api /usr/local/bin/

EXPOSE 9090
CMD ["hello-api"]