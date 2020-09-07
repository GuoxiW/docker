# Alpine Linux 是一个社区开发的面向安全应用的轻量级Linux发行版，适合用来做Docker镜像
FROM golang:1-alpine as build

RUN apk add --no-cache git

# Golang 包管理工具
RUN go get -u github.com/Masterminds/glide

# FLO 的 Golang 语言全节点
RUN git clone https://github.com/bitspill/flod /go/src/github.com/bitspill/flod
WORKDIR /go/src/github.com/bitspill/flod
RUN glide install
RUN go install . ./cmd/...


FROM alpine:latest

# 构建参数用于设置环境变量
ARG network=mainnet

COPY --from=build /go/bin/flod /flod/bin/flod
COPY ./${network}.conf /flod/flod.conf

VOLUME ["/root/.flod/"]
ENTRYPOINT ["/flod/bin/flod", "--configfile=/flod/flod.conf"]
